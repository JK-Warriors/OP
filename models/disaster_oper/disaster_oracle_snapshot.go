package disaster_oper

import (
	"database/sql"
	"opms/lib/oracle"
	"opms/utils"

	"github.com/godror/godror"
)

func OraStartSnapshot(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	utils.LogDebug("Start snapshot database in progress...")
	//get database role
	role, err := oracle.GetDatabaseRole(db)
	if err != nil {
		utils.LogDebug("获取数据库角色失败")
		Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "获取数据库角色失败")
		Update_OP_Reason(op_id, "获取数据库角色失败")
		return -1
	}
	Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	//get database version
	db_version, _ := oracle.GetDatabaseVersion(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "获取数据库版本成功")
	utils.LogDebug("获取数据库版本成功")

	// get sync_status
	sync_status, _ := oracle.GetSyncStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "获取数据库同步进程状态成功")
	utils.LogDebug("获取数据库同步进程状态成功")

	// get instance status
	inst_status, _ := oracle.GetInstanceStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "获取数据库实例状态成功")
	utils.LogDebug("获取数据库实例状态成功")

	if db_version <= 10 {
		utils.LogDebug("进入快照失败，10g以下版本不支持数据库快照")
		Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "进入快照失败，10g以下版本不支持数据库快照")
		Update_OP_Reason(op_id, "进入快照失败，10g以下版本不支持数据库快照")
		return -1
	}

	//get flashback status
	// fb_status, _ := oracle.GetFlashbackStatus(db)
	// utils.LogDebug("获取数据库闪回状态成功")
	// if fb_status == "NO" {
	// 	utils.LogDebug("进入快照失败，当前数据库没有开启快照功能")
	// 	Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "进入快照失败，当前数据库没有开启快照功能")
	// 	Update_OP_Reason(op_id, "进入快照失败，当前数据库没有开启快照功能")
	// 	return -1
	// }

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "验证数据库角色成功")
		utils.LogDebug("Now we are going to start snapshot.")

		if inst_status == "MOUNTED" {
			if sync_status > 0 {
				if _, err = db.Exec("alter database recover managed standby database cancel"); err != nil {
					utils.LogDebug("Cancel recover managed standby database failed: " + err.Error())
					Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "备库退出同步进程失败")
					Update_OP_Reason(op_id, "备库退出同步进程失败")
					return -1
				}
				Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "数据库同步进程停止成功")
			}

		} else if inst_status == "OPEN" {
			oracle.ShutdownImmediate(P)
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "关闭备端数据库")
			oracle.StartupMount(P)
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "启动备端数据库到Mount")

			// 再次获取数据库连接
			db, err := sql.Open("godror", P.StringWithPassword())
			if err != nil {
				utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
			}
			defer db.Close()

			// 再次获取数据库实例状态
			inst_status, _ := oracle.GetInstanceStatus(db)
			if inst_status != "MOUNTED" {
				utils.LogDebug("备库启动到mount状态失败: " + err.Error())
				Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "备库启动到mount状态失败")
				Update_OP_Reason(op_id, "备库启动到mount状态失败")
				return -1
			}
		}

		//进入快照
		if _, err = db.Exec("alter database convert to snapshot standby"); err != nil {
			utils.LogDebug("Alter database convert to snapshot standby failed: " + err.Error())
			Update_OP_Reason(op_id, "备库进入快照失败")
			return -1
		} else {
			Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "数据库进入快照成功")
		}

		if _, err = db.Exec("alter database open"); err != nil {
			utils.LogDebug("Alter database open failed: " + err.Error())
			Update_OP_Reason(op_id, "备库打开失败")
			return -1
		} else {
			Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "打开数据库成功")
		}

		//验证结果
		role, _ = oracle.GetDatabaseRole(db)
		if role == "SNAPSHOT STANDBY" {
			Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "备库进入快照成功")
			utils.LogDebug("备库进入快照成功.")
			return 1
		} else {
			utils.LogDebug("备库进入快照失败，当前数据库角色不对.")
			Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "备库进入快照失败，当前数据库角色不对")
			Update_OP_Reason(op_id, "备库进入快照失败，当前数据库角色不对")
			return -1
		}
	} else {
		Log_OP_Process(op_id, bs_id, 1, "STARTSNAPSHOT", "验证数据库角色失败，无法进入快照")
		Update_OP_Reason(op_id, "验证数据库角色失败，无法进入快照")
		utils.LogDebug("验证数据库角色失败，无法进入快照!")
		return -1
	}
	return result
}

func OraStopSnapshot(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1
	var exec_command string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	utils.LogDebug("Stop snapshot database in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get instance status
	inst_status, _ := oracle.GetInstanceStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "获取数据库实例状态成功")
	utils.LogDebug("获取数据库实例状态成功")

	//get database version
	db_version, _ := oracle.GetDatabaseVersion(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "获取数据库版本成功")
	utils.LogDebug("获取数据库版本成功")

	if role == "SNAPSHOT STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "验证数据库角色成功")
		utils.LogDebug("正在准备退出快照...")
		if inst_status == "OPEN" {
			oracle.ShutdownImmediate(P)
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "关闭备端数据库")
			oracle.StartupMount(P)
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "启动备端数据库到Mount")

			// 再次获取数据库连接
			db, err := sql.Open("godror", P.StringWithPassword())
			if err != nil {
				utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
			}
			defer db.Close()

			// 再次获取数据库实例状态
			inst_status, _ := oracle.GetInstanceStatus(db)
			if inst_status != "MOUNTED" {
				utils.LogDebug("备库启动到mount状态失败: " + err.Error())
				Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "备库启动到mount状态失败")
				Update_OP_Reason(op_id, "备库启动到mount状态失败")
				return -1
			}

		}

		//退出快照
		if _, err = db.Exec("alter database convert to physical standby"); err != nil {
			utils.LogDebug("Alter database convert to physical standby failed: " + err.Error())
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "备库退出快照失败")
			Update_OP_Reason(op_id, "备库退出快照失败")
			return -1
		} else {
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "备库退出快照成功")
		}

		//退出快照后，重启实例
		oracle.ShutdownImmediate(P)
		Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "关闭备端数据库")
		oracle.StartupMount(P)
		Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "启动备端数据库到Mount")

		// 11g以上开启到Open状态
		if db_version > 10 {
			if _, err = db.Exec("alter database open"); err != nil {
				utils.LogDebug("Alter database open failed: " + err.Error())
				Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "备库打开失败")
			} else {
				Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "打开数据库成功")
			}
		}

		// get standby redo count
		sta_redo_count, _ := oracle.GetStandbyRedoLog(db)
		Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "获取数据库备用在线日志个数")
		if sta_redo_count > 0 {
			exec_command = "alter database recover managed standby database using current logfile disconnect from session"
		} else {
			exec_command = "alter database recover managed standby database disconnect from session"
		}

		if _, err = db.Exec(exec_command); err != nil {
			utils.LogDebug("Recover managed standby database failed: " + err.Error())
		} else {
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "数据库同步进程启动成功")
		}

		//验证结果
		role, _ = oracle.GetDatabaseRole(db)
		if role == "PHYSICAL STANDBY" {
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "备库退出快照成功")
			utils.LogDebug("备库退出快照成功.")
			return 1
		} else {
			utils.LogDebug("备库退出快照失败，当前数据库角色不对.")
			Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "备库退出快照失败，当前数据库角色不对")
			Update_OP_Reason(op_id, "备库退出快照失败，当前数据库角色不对")
			return -1
		}
	} else {
		Log_OP_Process(op_id, bs_id, 1, "STOPSNAPSHOT", "验证数据库角色失败，无法退出快照")
		Update_OP_Reason(op_id, "验证数据库角色失败，无法退出快照")
		utils.LogDebug("验证数据库角色失败，无法退出快照!")
		return -1
	}
	return result
}
