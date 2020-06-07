package disaster_oper

import (
	"database/sql"
	"opms/lib/oracle"
	"opms/utils"
	"time"

	"github.com/godror/godror"
)

func OraStartRead(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1
	var exec_command string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	utils.LogDebug("Start Read database in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get sync_status
	sync_status, _ := oracle.GetSyncStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "获取数据库同步进程状态成功")
	utils.LogDebug("获取数据库同步进程状态成功")

	//get database open mode
	open_mode, _ := oracle.GetOpenMode(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "获取数据库打开模式成功")

	// get standby redo count
	sta_redo_count, _ := oracle.GetStandbyRedoLog(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "获取数据库备用在线日志个数")
	if sta_redo_count > 0 {
		exec_command = "alter database recover managed standby database using current logfile disconnect from session"
	} else {
		exec_command = "alter database recover managed standby database disconnect from session"
	}

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "验证数据库角色成功")
		utils.LogDebug("Now we are going to start database read.")
		if open_mode == "MOUNTED" {
			if sync_status > 0 {
				utils.LogDebug("Cancel recover managed standby database...")

				if _, err = db.Exec("alter database recover managed standby database cancel"); err != nil {
					utils.LogDebug("Cancel recover managed standby database failed: " + err.Error())
					Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "关闭数据库同步进程失败")
					Update_OP_Reason(op_id, "关闭数据库同步进程失败")
					return -1
				}
				Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "关闭数据库同步进程")

			}

			if _, err = db.Exec("alter database open"); err != nil {
				utils.LogDebug("alter database open failed: " + err.Error())
				Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "启动数据库到可读状态失败")
				Update_OP_Reason(op_id, "启动数据库到可读状态失败")
				return -1
			}
			Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "启动数据库到可读状态成功")

			if _, err = db.Exec(exec_command); err != nil {
				utils.LogDebug("Recover managed standby database failed: " + err.Error())
				Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "启动数据库同步进程失败")
			}
			Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "启动数据库同步进程成功")

			// 再次获取数据库读写状态
			time.Sleep(3 * time.Second)
			open_mode, _ := oracle.GetOpenMode(db)
			if open_mode == "READ ONLY WITH APPLY" {
				Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "备库开启可读状态成功")
				utils.LogDebug("Start database read successfully.")
				return 1
			} else {
				Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "备库开启可读状态失败")
				utils.LogDebug("Start database read failed.")
				return -1
			}
		} else {
			utils.LogDebug("The database open mode is not mount...")
			Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "数据库没有处于mount状态，无法开启可读")
			Update_OP_Reason(op_id, "数据库没有处于mount状态，无法开启可读")
			return -1
		}

	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，无法开启可读状态")
		utils.LogDebug("Check database role failed. You can not start database read!")
	}
	return result
}

func OraStopRead(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1
	var exec_command string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	utils.LogDebug("Stop Sync database in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	//get database open mode
	open_mode, _ := oracle.GetOpenMode(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "获取数据库打开模式成功")

	// get standby redo count
	sta_redo_count, _ := oracle.GetStandbyRedoLog(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "获取数据库备用在线日志个数")
	if sta_redo_count > 0 {
		exec_command = "alter database recover managed standby database using current logfile disconnect from session"
	} else {
		exec_command = "alter database recover managed standby database disconnect from session"
	}

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "验证数据库角色成功")
		utils.LogDebug("Now we are going to stop database read.")

		if open_mode == "READ ONLY WITH APPLY" {

			oracle.ShutdownImmediate(P)
			Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "关闭备端数据库")
			oracle.StartupMount(P)
			Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "启动备端数据库到Mount")

			db2, err := sql.Open("godror", P.StringWithPassword())
			if err != nil {
				utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
				Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "重新获取数据库连接失败")
				Update_OP_Reason(op_id, "重新获取数据库连接失败")
				return -1
			}
			defer db2.Close()
			Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "重新获取数据库连接成功")

			if _, err = db2.Exec(exec_command); err != nil {
				utils.LogDebug("Recover managed standby database failed: " + err.Error())
				Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "启动数据库同步进程失败")
			}
			Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "启动数据库同步进程成功")

			// 再次获取数据库读写状态
			time.Sleep(3 * time.Second)
			open_mode, _ := oracle.GetOpenMode(db2)
			if open_mode == "MOUNTED" {
				Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "备库停止可读状态成功")
				utils.LogDebug("Stop database read successfully.")
				return 1
			} else {
				Log_OP_Process(op_id, bs_id, 1, "STOPREAD", "备库停止可读状态失败")
				utils.LogDebug("Stop database read failed.")
				return -1
			}
		} else {
			utils.LogDebug("The database open mode is not read mode...")
			Log_OP_Process(op_id, bs_id, 1, "STARTREAD", "数据库没有处于可读状态，无法停止可读")
			Update_OP_Reason(op_id, "数据库没有处于可读状态，无法停止可读")
			return -1
		}

	} else {
		utils.LogDebug("Check database role failed. You can not stop database read!")
		Update_OP_Reason(op_id, "验证数据库角色失败，无法停止同步进程")
	}
	return result
}
