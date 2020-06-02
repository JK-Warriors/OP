package disaster_oper

import (
	"database/sql"
	"opms/lib/oracle"
	"opms/utils"
	"time"

	"github.com/godror/godror"
	errors "golang.org/x/xerrors"
)

func OraStartSync(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1
	var exec_command string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebug(errors.Errorf("%s: %w", P.StringWithPassword(), err))
	}
	defer db.Close()

	utils.LogDebug("Start Sync database in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get sync_status
	sync_status, _ := oracle.GetSyncStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "获取数据库同步进程状态成功")
	utils.LogDebug("获取数据库同步进程状态成功")

	// get standby redo count
	sta_redo_count, _ := oracle.GetstandbyRedoLog(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "获取数据库备用在线日志个数")
	if sta_redo_count > 0 {
		exec_command = "alter database recover managed standby database using current logfile disconnect from session"
	} else {
		exec_command = "alter database recover managed standby database disconnect from session"
	}

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "验证数据库角色成功")
		utils.LogDebug("Now we are going to start synchronization.")
		if sync_status > 0 {
			utils.LogDebug("The synchronization process is already active...")
			Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "数据库同步进程已经开启")
			Update_OP_Reason(op_id, "数据库同步进程已经开启")
			return -1
		} else {
			if _, err = db.Exec(exec_command); err != nil {
				utils.LogDebug("Recover managed standby database failed: " + err.Error())
			}
			Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "数据库同步进程启动成功")
		}

		// 再次获取数据库同步进程
		time.Sleep(3 * time.Second)
		sync_status, _ := oracle.GetSyncStatus(db)
		if sync_status > 0 {
			Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "备库已经处于同步状态")
			utils.LogDebug("Start synchronization successfully.")
			return 1
		} else {
			Log_OP_Process(op_id, bs_id, 1, "STARTSYNC", "备库开启同步失败")
			utils.LogDebug("Start synchronization failed.")
			return -1
		}
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，无法开启同步进程")
		utils.LogDebug("Check database role failed. You can not start synchronization!")
	}
	return result
}

func OraStopSync(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebug(errors.Errorf("%s: %w", P.StringWithPassword(), err))
	}
	defer db.Close()

	utils.LogDebug("Stop Sync database in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPSYNC", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get sync_status
	sync_status, _ := oracle.GetSyncStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPSYNC", "获取数据库同步进程状态成功")
	utils.LogDebug("获取数据库同步进程状态成功")

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STOPSYNC", "验证数据库角色成功")
		utils.LogDebug("Now we are going to stop synchronization.")
		if sync_status > 0 {
			if _, err = db.Exec("alter database recover managed standby database cancel"); err != nil {
				utils.LogDebug("Cancel recover managed standby database failed: " + err.Error())
			}
			Log_OP_Process(op_id, bs_id, 1, "STOPSYNC", "数据库同步进程停止成功")

		} else {
			utils.LogDebug("The synchronization process is already stop...")
			Log_OP_Process(op_id, bs_id, 1, "STOPSYNC", "数据库同步进程已经停止")
			Update_OP_Reason(op_id, "数据库同步进程已经停止")
			return -1
		}

		// 再次获取数据库同步进程
		time.Sleep(3 * time.Second)
		sync_status, _ := oracle.GetSyncStatus(db)
		if sync_status > 0 {
			Log_OP_Process(op_id, bs_id, 1, "STOPSYNC", "备库停止同步失败")
			utils.LogDebug("Stop synchronization failed.")
			return -1
		} else {
			Log_OP_Process(op_id, bs_id, 1, "STOPSYNC", "备库已经处于停止同步状态")
			utils.LogDebug("Stop synchronization successfully.")
			return 1
		}
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，无法停止同步进程")
		utils.LogDebug("Check database role failed. You can not stop synchronization!")
	}
	return result
}
