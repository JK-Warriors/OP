package disaster_oper

import (
	"database/sql"
	"fmt"
	"opms/lib/oracle"
	"opms/utils"
	"time"

	"github.com/godror/godror"
)

func OraPrimaryToStandby(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1
	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "获取主库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get switchover status
	switch_status, _ := oracle.GetSwitchoverStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "获取主库切换状态成功")
	utils.LogDebug("The current database switchover status is: " + switch_status)

	// get database version
	version, _ := oracle.GetDatabaseVersion(db)
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "获取主库版本成功")
	utils.LogDebugf("The current database version is: %d", version)

	// get gap count
	gap_count, _ := oracle.GetGapCount(db)
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "获取主库gap数量成功")
	utils.LogDebugf("The current gap count is: %d", gap_count)

	if role == "PRIMARY" {
		Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "验证主库角色成功")
		utils.LogDebug("Now we are going to switch database %d to physical standby.")

		if switch_status == "TO STANDBY" || switch_status == "SESSIONS ACTIVE" || switch_status == "FAILED DESTINATION" || (switch_status == "RESOLVABLE GAP" && gap_count == 0) {

			utils.LogDebug("Switchover to physical standby...")
			Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "正在将主库切换成备库，可能会花费几分钟时间，请耐心等待...")

			if _, err = db.Exec("alter database commit to switchover to physical standby with session shutdown"); err != nil {
				utils.LogDebugf("Commit to switchover to physical standby failed: %w", err)
			}
			oracle.ShutdownImmediate(P)
			oracle.StartupMount(P)

			// 获取oracle连接
			db2, err := sql.Open("godror", P.StringWithPassword())
			if err != nil {
				utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
			}
			defer db2.Close()

			if version > 10 {
				utils.LogDebug("Alter standby database to open read only in progress...")
				Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "正在将当前数据库启动到只读状态...")

				if _, err = db2.Exec("alter database open"); err != nil {
					utils.LogDebug("alter database open failed: " + err.Error())
				}
				if _, err = db2.Exec("alter database recover managed standby database disconnect from session"); err != nil {
					utils.LogDebug("Recover managed standby database failed: " + err.Error())
				}

				open_mode, _ := oracle.GetOpenMode(db2)
				if open_mode == "READ ONLY" || open_mode == "READ ONLY WITH APPLY" {
					Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "当前数据库已经成功启动到只读状态")
					utils.LogDebug("Alter standby database to open successfully.")
				} else {
					Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "当前数据库已经成功启动到只读状态")
					utils.LogDebug("Start MRP process failed!")
				}
			} else {
				if _, err = db2.Exec("alter database recover managed standby database disconnect from session"); err != nil {
					utils.LogDebug("Recover managed standby database failed: " + err.Error())
				}
			}

			role, _ := oracle.GetDatabaseRole(db2)
			if role == "PHYSICAL STANDBY" {
				Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "主库已经成功切换成备库")
				utils.LogDebug("Switchover to physical standby successfully.")
				return 1
			} else {
				Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "主库切换备库失败")
				utils.LogDebug("Switchover to physical standby failed.")
				return -1
			}
		} else {
			Update_OP_Reason(op_id, "验证数据库切换状态失败")
			result = -1
		}

	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，当前数据库不是主库，不能切换到备库")
		Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "验证数据库角色失败，当前数据库不是主库，不能切换到备库")
		utils.LogDebug("You can not switchover a standby database to standby standby!")
		result = -1
	}
	return result
}

func OraStandbyToPrimary(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1
	var msg string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	utils.LogDebug("Switchover database to primary in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "获取备库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get switchover status
	switch_status, _ := oracle.GetSwitchoverStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "获取备库切换状态成功")
	utils.LogDebug("The current database switchover status is: " + switch_status)

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "验证备库角色成功")
		utils.LogDebug("Now we are going to switch database to primary.")
		if switch_status == "NOT ALLOWED" {
			msg = fmt.Sprintf("数据库状态为 %s，无法进行切换", switch_status)
			utils.LogDebug(msg)
			Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", msg)
			Update_OP_Reason(op_id, msg)
			return -1
		} else if switch_status == "SWITCHOVER PENDING" || switch_status == "SWITCHOVER LATENT" {
			if _, err = db.Exec("alter database recover managed standby database disconnect from session"); err != nil {
				utils.LogDebug("Recover managed standby database failed: " + err.Error())
			}
			time.Sleep(5 * time.Second)
			to_primary(op_id, bs_id, P)
		} else if switch_status == "TO PRIMARY" || switch_status == "SESSIONS ACTIVE" {
			to_primary(op_id, bs_id, P)
		}

		// 重新切换后数据库角色
		db2, err := sql.Open("godror", P.StringWithPassword())
		if err != nil {
			utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
		}
		defer db2.Close()

		db_role, _ := oracle.GetDatabaseRole(db2)
		if db_role == "PRIMARY" {
			Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "备库已经成功切换成主库")
			utils.LogDebug("Switchover standby database to primary successfully.")
			return 1
		} else {
			Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "备库切换主库失败")
			utils.LogDebug("Switchover standby database to primary failed.")
			return -1
		}
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，当前数据库无法切换到主库")
		utils.LogDebug("You can not switchover primary database to primary!")
	}
	return result
}

func to_primary(op_id int64, bs_id int, P godror.ConnectionParams) {
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "正在将备库切换成主库，可能会花费几分钟时间，请耐心等待...")
	utils.LogDebug("Switchover standby database to primary...")

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	if _, err = db.Exec("alter database commit to switchover to primary with session shutdown"); err != nil {
		utils.LogDebug("Switchover to primary failed: " + err.Error())
	}

	oracle.ShutdownImmediate(P)
	oracle.StartupOpen(P)

	db2, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db2.Close()

	if _, err = db2.Exec("alter system archive log current"); err != nil {
		utils.LogDebug("Alter system archive log current failed: " + err.Error())
	}
}
