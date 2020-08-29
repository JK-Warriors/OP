package dr_oper

import (
	"database/sql"
	"fmt"
	"opms/utils"
	"time"

)

func SlaveToMaster(op_id int64, bs_id int, dsn string) int {
	result := -1
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		utils.LogDebugf("%s: %s", dsn, err.Error())
	}
	defer db.Close()

	//get database role, 1: master; 2:slave
	role, _ := GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "获取主库角色成功")
	utils.LogDebugf("The current database role is: %s", role)

	if role == 2{
		// get master status
		// get slave status
		if (1==1){
			// can switch now
			utils.LogDebugf("Now we are going to switch database %d to master.", sta_id)
			Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "正在将从库切换成主库...")
			_, err = mysql.Exec("stop slave io_thread")
			if err != nil {
				utils.LogDebugf("Stop slave io_thread error: %s", err.Error())
			}

			_, err = mysql.Exec("stop slave")
			if err != nil {
				utils.LogDebugf("Stop slave error: %s", err.Error())
			}

			
			_, err = mysql.Exec("reset slave all")
			if err != nil {
				utils.LogDebugf("Reset slave all error: %s", err.Error())
			}
	
			utils.Info("Switchover slave to master successfully.")
			Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "从库已经成功切换成主库")
			result=0
		}
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		Log_OP_Process(op_id, bs_id, 1, "SWITCHOVER", "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		utils.LogDebug("You can not switchover a master database to master!")
		result = -1
	}
	return result
}

func RebuildReplication(op_id int64, bs_id int, dsn string) int {
	
	return result
}


func FailoverToMaster(op_id int64, bs_id int, dsn string) int {
	result := -1
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		utils.LogDebugf("%s: %s", dsn, err.Error())
	}
	defer db.Close()

	//get database role, 1: master; 2:slave
	role, _ := GetDatabaseRole(db)
	Log_OP_Process(op_id, bs_id, 1, "FAILOVER", "获取主库角色成功")
	utils.LogDebugf("The current database role is: %s", role)

	if role == 2{
		// failover to master
		utils.LogDebugf("Now we are going to switch database %d to master.", sta_id)
		Log_OP_Process(op_id, bs_id, 1, "FAILOVER", "正在将从库切换成主库...")
		_, err = mysql.Exec("stop slave io_thread")
		if err != nil {
			utils.LogDebugf("Stop slave io_thread error: %s", err.Error())
		}

		_, err = mysql.Exec("stop slave")
		if err != nil {
			utils.LogDebugf("Stop slave error: %s", err.Error())
		}

		
		_, err = mysql.Exec("reset slave all")
		if err != nil {
			utils.LogDebugf("Reset slave all error: %s", err.Error())
		}

		utils.Info("Switchover slave to master successfully.")
		Log_OP_Process(op_id, bs_id, 1, "FAILOVER", "从库已经成功切换成主库")
		result=0
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		Log_OP_Process(op_id, bs_id, 1, "FAILOVER", "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		utils.LogDebug("You can not switchover a master database to master!")
		result = -1
	}

	return result
}


func GetDatabaseRole(db *sql.DB) int{
	return 2
}
