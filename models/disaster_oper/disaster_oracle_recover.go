package disaster_oper

import (
	"database/sql"
	"fmt"
	"opms/lib/oracle"
	"opms/utils"
	"strings"
	"time"

	"github.com/godror/godror"
)

type RestorePoint struct {
	Name string `orm:"column(name);"`
}

func GetRestorePointName(P godror.ConnectionParams) ([]string, error) {
	var restore_name []string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	//get database restore point
	restore_name, err = oracle.GetRestorePointName(db)
	if err != nil {
		utils.LogDebug("Get restore point failed: " + err.Error())
		return restore_name, err
	}

	return restore_name, nil
}

func OraStartFlashback(op_id int64, bs_id int, fb_method int, fb_point string, fb_time string, P godror.ConnectionParams) int {
	result := -1
	var exec_command string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	utils.LogDebug("Start Flashback database in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get sync_status
	sync_status, _ := oracle.GetSyncStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "获取数据库同步进程状态成功")
	utils.LogDebug("获取数据库同步进程状态成功")

	// get instance status
	inst_status, _ := oracle.GetInstanceStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "获取数据库实例状态成功")
	utils.LogDebug("获取数据库实例状态成功")

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "验证数据库角色成功")
		utils.LogDebug("Now we are going to start flashback.")

		if inst_status == "MOUNTED" {
			if sync_status > 0 {
				if _, err = db.Exec("alter database recover managed standby database cancel"); err != nil {
					utils.LogDebug("Cancel recover managed standby database failed: " + err.Error())
					Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "备库退出同步进程失败")
					Update_OP_Reason(op_id, "备库退出同步进程失败")
					return -1
				}
				Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "数据库同步进程停止成功")
			}

		} else if inst_status == "OPEN" {
			Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "当前数据库处于打开可读状态，需重新启动到Mount状态")
			oracle.ShutdownImmediate(P)
			Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "关闭备端数据库")
			oracle.StartupMount(P)
			Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "启动备端数据库到Mount")

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
				Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "备库启动到mount状态失败")
				Update_OP_Reason(op_id, "备库启动到mount状态失败")
				return -1
			}
		}

		//开始闪回

		if fb_method == 1 {
			exec_command = fmt.Sprintf("flashback database to restore point %s", fb_point)

		} else {
			fb_time = strings.Replace(fb_time, "T", " ", -1)
			exec_command = fmt.Sprintf("flashback database to timestamp to_timestamp('%s','yy-mm-dd hh24:mi:ss')", fb_time)
		}
		utils.LogDebug(exec_command)

		if _, err = db.Exec(exec_command); err != nil {
			utils.LogDebugf("%s failed: %s", exec_command, err.Error())
			Update_OP_Reason(op_id, "执行恢复快照命令失败, "+err.Error())
			return -1
		} else {
			utils.LogDebug("数据库恢复快照成功.")
			Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "数据库恢复快照成功")
		}

		if _, err = db.Exec("alter database open"); err != nil {
			utils.LogDebug("Alter database open failed: " + err.Error())
			Update_OP_Reason(op_id, "备库打开可读失败, "+err.Error())
			return -1
		} else {
			utils.LogDebug("数据库打开可读成功.")
			Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "数据库打开可读成功")
		}

		utils.LogDebug("数据库恢复成功.")
		Log_OP_Process(op_id, bs_id, 1, "STARTFLASHBACK", "数据库恢复成功")
		return 1

	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，无法开启恢复进程")
		utils.LogDebug("Check database role failed. You can not start flashback database!")
	}
	return result
}

func OraRecover(op_id int64, bs_id int, P godror.ConnectionParams) int {
	result := -1
	var exec_command string

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	utils.LogDebug("Start Sync database in progress...")
	//get database role
	role, _ := oracle.GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "获取数据库角色成功")
	utils.LogDebug("The current database role is: " + role)

	// get sync_status
	sync_status, _ := oracle.GetSyncStatus(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "获取数据库同步进程状态成功")
	utils.LogDebug("获取数据库同步进程状态成功")

	// get standby redo count
	sta_redo_count, _ := oracle.GetStandbyRedoLog(db)
	Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "获取数据库备用在线日志个数")
	if sta_redo_count > 0 {
		exec_command = "alter database recover managed standby database using current logfile disconnect from session"
	} else {
		exec_command = "alter database recover managed standby database disconnect from session"
	}

	if role == "PHYSICAL STANDBY" {
		Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "验证数据库角色成功")
		utils.LogDebug("Now we are going to start synchronization.")
		if sync_status > 0 {
			utils.LogDebug("The synchronization process is already active...")
			Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "数据库同步进程已经开启")
			Update_OP_Reason(op_id, "数据库同步进程已经开启")
			return -1
		} else {
			if _, err = db.Exec(exec_command); err != nil {
				utils.LogDebug("Recover managed standby database failed: " + err.Error())
			}
			Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "数据库同步进程启动成功")
		}

		// 再次获取数据库同步进程
		time.Sleep(3 * time.Second)
		sync_status, _ := oracle.GetSyncStatus(db)
		if sync_status > 0 {
			Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "备库已经处于同步状态")
			utils.LogDebug("Start synchronization successfully.")
			return 1
		} else {
			Log_OP_Process(op_id, bs_id, 1, "STOPFLASHBACK", "备库开启同步失败")
			utils.LogDebug("Start synchronization failed.")
			return -1
		}
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，无法开启同步进程")
		utils.LogDebug("Check database role failed. You can not start synchronization!")
	}
	return result
}
