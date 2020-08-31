package dr_oper

import (
	//"fmt"
	"opms/utils"
	//"time"
	//"reflect"
	"github.com/xormplus/xorm"

)

func SlaveToMaster(op_id int64, bs_id int, dsn_p string, dsn_s string, slave_id int) int {
	result := -1
	mdb, err := xorm.NewEngine("mysql", dsn_p)
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("%s: %s", dsn_p, err.Error())
		return -1
	}
	defer mdb.Close()
	err = mdb.Ping()
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("Ping database failed: %s", err.Error())
		return -1
	}

	//lock tables
	utils.LogDebugf("Lock tables for database first")
	_, err = mdb.Exec("flush tables with read lock")
	if err != nil {
		Update_OP_Reason(op_id, "锁表失败")
		utils.LogDebugf("Lock tables error: %s", err.Error())
		return -1
	}

	
	sdb, err := xorm.NewEngine("mysql", dsn_s)
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("%s: %s", dsn_s, err.Error())
		return -1
	}
	defer sdb.Close()
	err = sdb.Ping()
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("Ping database failed: %s", err.Error())
		return -1
	}

	//get database role, 1: master; 2:slave
	role := GetDatabaseRole(sdb)
	Log_OP_Process(op_id, bs_id, 2, "SWITCHOVER", "获取数据库角色成功")
	utils.LogDebugf("The current database role is: %d", role)

	if role == 2{
		// Check prerequisites
		check_result := CheckSlaveDelay(mdb, sdb)
		if (check_result > 0){
			// can switch now
			utils.LogDebugf("Now we are going to switch database %d to master.", slave_id)
			Log_OP_Process(op_id, bs_id, 2, "SWITCHOVER", "正在将从库切换成主库...")
			_, err = sdb.Exec("stop slave io_thread")
			if err != nil {
				Update_OP_Reason(op_id, "停止io_thread进程失败")
				utils.LogDebugf("Stop slave io_thread error: %s", err.Error())
				result = -1
			}

			_, err = sdb.Exec("stop slave")
			if err != nil {
				Update_OP_Reason(op_id, "停止slave进程失败")
				utils.LogDebugf("Stop slave error: %s", err.Error())
				result = -1
			}

			
			_, err = sdb.Exec("reset slave all")
			if err != nil {
				Update_OP_Reason(op_id, "重置slave失败")
				utils.LogDebugf("Reset slave all error: %s", err.Error())
				result = -1
			}
	
			utils.LogDebug("Switchover slave to master successfully.")
			Log_OP_Process(op_id, bs_id, 2, "SWITCHOVER", "从库已经成功切换成主库")
			result=0
		}else{
			Update_OP_Reason(op_id, "校验切换条件失败")
			Log_OP_Process(op_id, bs_id, 2, "SWITCHOVER", "校验切换条件失败，取消切换")
			utils.LogDebug("You can not switchover a master database to master!")
			result = -1
		}
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		Log_OP_Process(op_id, bs_id, 2, "SWITCHOVER", "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		utils.LogDebug("Check prerequisites failed!")
		result = -1
	}
	return result
}

func RebuildReplication(op_id int64, bs_id int, dsn_p string, dsn_s string, slave_id int) int {
	utils.LogDebug("Rebuild replication in progress...")
	var result int = -1

	mdb, err := xorm.NewEngine("mysql", dsn_p)
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("%s: %s", dsn_p, err.Error())
		return -1
	}
	defer mdb.Close()
	err = mdb.Ping()
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("Ping database failed: %s", err.Error())
		return -1
	}

	sdb, err := xorm.NewEngine("mysql", dsn_s)
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("%s: %s", dsn_s, err.Error())
		return -1
	}
	defer sdb.Close()
	err = sdb.Ping()
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("Ping database failed: %s", err.Error())
		return -1
	}

	var master_file string
	var master_pos string
	masterstatus, err := sdb.QueryString("show master status")
	if err != nil {
		Update_OP_Reason(op_id, "获取master状态失败")
		utils.LogDebugf("[Info] Get master status error: %s", err.Error())
		return -1
	}
	if(len(masterstatus) == 1){
		for key,value := range masterstatus[0] {
			if(key == "File"){
				master_file = value
			}else if (key == "Position"){
				master_pos = value
			}
		}

	}
    
    // unlock tables
	utils.LogDebugf("Unlock tables for database first")
	_, err = mdb.Exec(`unlock tables`)
	if err != nil {
		Update_OP_Reason(op_id, "解锁tables失败")
		utils.LogDebugf("Unlock tables error: %s", err.Error())
		result = -1
	}
	
	_, err = mdb.Exec("stop slave io_thread")
	if err != nil {
		Update_OP_Reason(op_id, "停止io_thread进程失败")
		utils.LogDebugf("Stop slave io_thread error: %s", err.Error())
		result = -1
	}
	_, err = mdb.Exec("stop slave")
	if err != nil {
		Update_OP_Reason(op_id, "停止slave失败")
		utils.LogDebugf("Stop slave error: %s", err.Error())
		result = -1
	}
	_, err = mdb.Exec("reset slave all")
	if err != nil {
		Update_OP_Reason(op_id, "重置slave失败")
		utils.LogDebugf("Reset slave all error: %s", err.Error())
		result = -1
	}

	utils.LogDebugf("[Info] master_file: %s", master_file)
	utils.LogDebugf("[Info] master_pos: %s", master_pos)
	sql,_ := GetChangeMasterCmd(slave_id, master_file, master_pos)
	utils.LogDebugf("Change master command: %s", sql)
	_, err = mdb.Exec(sql)
	if err != nil {
		Update_OP_Reason(op_id, "重建复制关系失败")
		utils.LogDebugf("Rebuild replication error: %s", err.Error())
		result = -1
	}
	
	_, err = mdb.Exec("start slave")
	if err != nil {
		Update_OP_Reason(op_id, "开启slave进程失败")
		utils.LogDebugf("Start slave error: %s", err.Error())
		result = -1
	}
	
	//check result
	var io_running string = "No"
	var sql_running string = "No"
	slavestatus, err := mdb.QueryString("show slave status")
	if err != nil {
		Update_OP_Reason(op_id, "查看slave进程状态失败")
		utils.LogDebugf("[Info] Get slave status error: %s", err.Error())
		return -1
	}

	if(len(slavestatus) == 1){
		for key,value := range slavestatus[0] {
			if(key == "Slave_IO_Running"){
				io_running = value
			}else if (key == "Slave_SQL_Running"){
				sql_running = value
			}
		}
	}

	utils.LogDebugf("[Info] io_running: %s", io_running)
	utils.LogDebugf("[Info] sql_running: %s", sql_running)
	if((io_running == "Yes" || io_running == "Connecting") && sql_running == "Yes"){
		result = 0
	}

	return result
}



func FailoverToMaster(op_id int64, bs_id int, dsn string, slave_id int) int {
	result := -1
	sdb, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("%s: %s", dsn, err.Error())
		return -1
	}
	defer sdb.Close()
	err = sdb.Ping()
	if err != nil {
		Update_OP_Reason(op_id, "连接数据库失败")
		utils.LogDebugf("Ping database failed: %s", err.Error())
		return -1
	}

	//get database role, 1: master; 2:slave
	role := GetDatabaseRole(sdb)
	Log_OP_Process(op_id, bs_id, 2, "FAILOVER", "获取数据库角色成功")
	utils.LogDebugf("The current database role is: %d", role)

	if role == 2{
		// failover to master
		utils.LogDebugf("Now we are going to switch database %d to master.", slave_id)
		Log_OP_Process(op_id, bs_id, 2, "FAILOVER", "正在将从库切换成主库...")
		_, err = sdb.Exec("stop slave io_thread")
		if err != nil {
			Update_OP_Reason(op_id, "停止io_thread进程失败")
			utils.LogDebugf("Stop slave io_thread error: %s", err.Error())
			result = -1
		}

		_, err = sdb.Exec("stop slave")
		if err != nil {
			Update_OP_Reason(op_id, "停止slave失败")
			utils.LogDebugf("Stop slave error: %s", err.Error())
			result = -1
		}

		
		_, err = sdb.Exec("reset slave all")
		if err != nil {
			Update_OP_Reason(op_id, "重置slave失败")
			utils.LogDebugf("Reset slave all error: %s", err.Error())
			result = -1
		}

		utils.LogDebug("Failover slave to master successfully.")
		Log_OP_Process(op_id, bs_id, 2, "FAILOVER", "从库已经成功切换成主库")
		result=0
	} else {
		Update_OP_Reason(op_id, "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		Log_OP_Process(op_id, bs_id, 2, "FAILOVER", "验证数据库角色失败，当前数据库不是从库，不能切换到主库")
		utils.LogDebug("You can not failover a master database to master!")
		result = -1
	}

	return result
}


func GetDatabaseRole(db *xorm.Engine) int{
	var mastercount int = -1
	_, err := db.SQL("select count(host) from mysql.slave_master_info").Get(&mastercount)
	if err != nil {
		utils.LogDebugf("Info GetDatabaseRole error: %s", err.Error())
		return -1
	}
	utils.LogDebugf("Info GetDatabaseRole: masterCount: %d", mastercount)
	
	if mastercount > 0 {
		return 2
	}else{
		return 1
	}
}

func CheckSlaveDelay(mdb *xorm.Engine, sdb *xorm.Engine) int{
	var checkvalue int =-1

	var master_file string
	var master_pos string
	masterstatus, err := mdb.QueryString("show master status")
	if err != nil {
		utils.LogDebugf("[Info] Get master status error: %s", err.Error())
		return -1
	}
	if(len(masterstatus) == 1){
		for key,value := range masterstatus[0] {
			if(key == "File"){
				master_file = value
			}else if (key == "Position"){
				master_pos = value
			}
		}

	}
	utils.LogDebugf("[Info] master_file: %s", master_file)
	utils.LogDebugf("[Info] master_pos: %s", master_pos)


	var io_running string = "No"
	var sql_running string = "No"
	var log_file string
	var log_pos string
	var sql_delay string
	slavestatus, err := sdb.QueryString("show slave status")
	if err != nil {
		utils.LogDebugf("[Info] Get slave status error: %s", err.Error())
		return -1
	}

	if(len(slavestatus) == 1){
		for key,value := range slavestatus[0] {
			if(key == "Slave_IO_Running"){
				io_running = value
			}else if (key == "Slave_SQL_Running"){
				sql_running = value
			}else if (key == "Master_Log_File"){
				log_file = value
			}else if (key == "Read_Master_Log_Pos"){
				log_pos = value
			}else if (key == "SQL_Delay"){
				sql_delay = value
			}
		}
	}
	
	utils.LogDebugf("[Info] io_running: %s", io_running)
	utils.LogDebugf("[Info] sql_running: %s", sql_running)
	utils.LogDebugf("[Info] log_file: %s", log_file)
	utils.LogDebugf("[Info] log_pos: %s", log_pos)
	utils.LogDebugf("[Info] sql_delay: %s", sql_delay)

	if (io_running == "Yes" && sql_running == "Yes" && sql_delay == "0"){
		if(master_file == log_file && master_pos == log_pos){
			checkvalue = 1
		}else{
			utils.LogDebug("[Info] Check log file and position failed")
		}
	}else{
		utils.LogDebug("[Info] Check io thread failed")
	}

	return checkvalue
}
