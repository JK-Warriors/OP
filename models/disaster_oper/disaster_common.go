package disaster_oper

import (
	"fmt"
	"opms/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

func CheckDisasterConfig(bs_id int) (int, error) {
	var cfg_count int

	sql := `select count(1) from pms_disaster_config where bs_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, bs_id).QueryRow(&cfg_count)
	return cfg_count, err
}

func GetPrimaryDBId(bs_id int) (int, error) {
	var pri_dbid int

	sql := `select CASE is_switch
				WHEN 0 THEN db_id_p
				ELSE db_id_s
			END as pri_dbid
			from pms_disaster_config
			where bs_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, bs_id).QueryRow(&pri_dbid)
	return pri_dbid, err
}

func GetStandbyDBId(bs_id int) (int, error) {
	var sta_dbid int

	sql := `select CASE is_switch
				WHEN 0 THEN db_id_s
				ELSE db_id_p
			END as sta_dbid
			from pms_disaster_config
			where bs_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, bs_id).QueryRow(&sta_dbid)
	return sta_dbid, err
}

func GetDsn(db_id int, db_type int) (string, error) {
	var dsn string
	var sql string
	if db_type == 1 {
		sql = `select concat("oracle://",username,":",password ,"@" , host , ":" , port , "/" , instance_name , "?sysdba=1") as dsn 
				from pms_db_config where id = ? and db_type = ?`
	} else if db_type == 2 {
		sql = `select host from pms_db_config where id = ? and db_type = ?`
	} else if db_type == 3 {
		sql = `select host from pms_db_config where id = ? and db_type = ?`
	} else {
		sql = `select host from pms_db_config where id = ? and db_type = ?`
	}
	o := orm.NewOrm()
	err := o.Raw(sql, db_id, db_type).QueryRow(&dsn)
	return dsn, err
}

func OperationLock(bs_id int, op_type string) error {
	utils.LogDebug("Lock the process status in pms_disaster_config.")
	o := orm.NewOrm()
	var sql string
	if op_type == "SWITCHOVER" {
		sql = `update pms_disaster_config set on_process = 1, on_switchover = 1 where bs_id= ?`
	} else if op_type == "FAILOVER" {
		sql = `update pms_disaster_config set on_process = 1, on_failover = 1 where bs_id= ?`
	} else if op_type == "STARTSYNC" {
		sql = `update pms_disaster_config set on_process = 1, on_startsync = 1 where bs_id= ?`
	} else if op_type == "STOPSYNC" {
		sql = `update pms_disaster_config set on_process = 1, on_stopsync = 1 where bs_id= ?`
	} else if op_type == "STARTREAD" {
		sql = `update pms_disaster_config set on_process = 1, on_startread = 1 where bs_id= ?`
	} else if op_type == "STOPREAD" {
		sql = `update pms_disaster_config set on_process = 1, on_stopread = 1 where bs_id= ?`
	} else if op_type == "STARTSNAPSHOT" {
		sql = `update pms_disaster_config set on_process = 1, on_startsnapshot = 1 where bs_id= ?`
	} else if op_type == "STOPSNAPSHOT" {
		sql = `update pms_disaster_config set on_process = 1, on_stopsnapshot = 1 where bs_id= ?`
	}

	_, err := o.Raw(sql, bs_id).Exec()
	if err == nil {
		utils.LogDebug("Lock the process status successfully.")
	} else {
		utils.LogDebug("Lock the process status failed: " + err.Error())
	}
	return err
}

func OperationUnlock(bs_id int, op_type string) error {
	o := orm.NewOrm()
	var sql string
	utils.LogDebug("Unlock the process status in pms_disaster_config.")
	if op_type == "SWITCHOVER" {
		sql = `update pms_disaster_config set on_process = 0, on_switchover = 0 where bs_id= ?`
	} else if op_type == "FAILOVER" {
		sql = `update pms_disaster_config set on_process = 0, on_failover = 0 where bs_id= ?`
	} else if op_type == "STARTSYNC" {
		sql = `update pms_disaster_config set on_process = 0, on_startsync = 0 where bs_id= ?`
	} else if op_type == "STOPSYNC" {
		sql = `update pms_disaster_config set on_process = 0, on_stopsync = 0 where bs_id= ?`
	} else if op_type == "STARTREAD" {
		sql = `update pms_disaster_config set on_process = 0, on_startread = 1 where bs_id= ?`
	} else if op_type == "STOPREAD" {
		sql = `update pms_disaster_config set on_process = 0, on_stopread = 1 where bs_id= ?`
	} else if op_type == "STARTSNAPSHOT" {
		sql = `update pms_disaster_config set on_process = 0, on_startsnapshot = 0 where bs_id= ?`
	} else if op_type == "STOPSNAPSHOT" {
		sql = `update pms_disaster_config set on_process = 0, on_stopsnapshot = 0 where bs_id= ?`
	}

	_, err := o.Raw(sql, bs_id).Exec()
	if err == nil {
		utils.LogDebug("Unlock the process status successfully.")
	} else {
		utils.LogDebug("Unlock the process status failed: " + err.Error())
	}
	return err
}

func MoveOpRecordToHis(bs_id int, op_type string) error {
	o := orm.NewOrm()
	var sql string

	//将之前的操作记录移入his表
	sql = `insert into pms_opration_his select * from pms_opration where bs_id = ? and op_type = ? `
	_, err := o.Raw(sql, bs_id, op_type).Exec()
	if err != nil {
		utils.LogDebug("Move opration record to history table failed: " + err.Error())
	}

	sql = `delete from pms_opration where bs_id = ? and op_type = ? `
	_, err = o.Raw(sql, bs_id, op_type).Exec()
	if err != nil {
		utils.LogDebug("Delete opration record failed: " + err.Error())
	}

	sql = `insert into pms_op_process_his select * from pms_op_process where bs_id = ? and process_type = ? `
	_, err = o.Raw(sql, bs_id, op_type).Exec()
	if err != nil {
		utils.LogDebug("Move process record to history table failed: " + err.Error())
	}

	sql = `delete from pms_op_process where bs_id = ? and process_type = ? `
	_, err = o.Raw(sql, bs_id, op_type).Exec()
	if err != nil {
		utils.LogDebug("Delete process record failed: " + err.Error())
	}

	return err
}

func Init_OP_Instance(op_id int64, bs_id int, db_type int, op_type string) error {
	o := orm.NewOrm()
	var sql string
	//将之前的操作记录移入his表
	MoveOpRecordToHis(bs_id, op_type)

	//开始新的操作初始化
	str := fmt.Sprintf("Initialize opration instance for business %d.", bs_id)
	utils.LogDebug(str)

	sql = `insert into pms_opration(id, bs_id, db_type, op_type, created) values(?, ?, ?, ?, ?)`
	_, err := o.Raw(sql, op_id, bs_id, db_type, op_type, time.Now().Unix()).Exec()
	if err == nil {
		utils.LogDebug("Init the opration successfully.")
	} else {
		utils.LogDebug("Init the opration failed: " + err.Error())
	}
	return err
}

func Log_OP_Process(op_id int64, bs_id int, db_type int, op_type string, process_desc string) error {
	o := orm.NewOrm()
	var sql string

	sql = `insert into pms_op_process(op_id, bs_id, db_type, process_type, process_desc, created) values (?, ?, ?, ?, ?, ?)`
	_, err := o.Raw(sql, op_id, bs_id, db_type, op_type, process_desc, time.Now().Unix()).Exec()

	if err == nil {
		utils.LogDebug("Log the process successfully.")
	} else {
		utils.LogDebug("Log the process failed: " + err.Error())
	}

	return err
}

func UpdateSwitchFlag(bs_id int) {
	o := orm.NewOrm()
	var sql string
	var is_switch int

	utils.LogDebug("Update switch flag in pms_disaster_config in progress...")

	// get current switch flag
	sql = `select is_switch from pms_disaster_config where bs_id= ?`
	err := o.Raw(sql, bs_id).QueryRow(&is_switch)

	utils.LogDebug("The current switch flag is: " + strconv.Itoa(is_switch))

	if is_switch == 0 {
		sql = `update pms_disaster_config set is_switch = 1 where bs_id = ?`
	} else {
		sql = `update pms_disaster_config set is_switch = 0 where bs_id = ?`
	}

	_, err = o.Raw(sql, bs_id).Exec()
	if err == nil {
		utils.LogDebug("Update switch flag in pms_disaster_config successfully.")
	} else {
		utils.LogDebug("Update switch flag in pms_disaster_config failed: " + err.Error())
	}
}

func Update_OP_Result(op_id int64, result int) {
	o := orm.NewOrm()
	var sql string

	sql = `update pms_opration set result = ? where id = ?`
	_, _ = o.Raw(sql, result, op_id).Exec()

}

func Update_OP_Reason(op_id int64, reason string) {
	o := orm.NewOrm()
	var sql string

	sql = `update pms_opration set reason = ? where id = ?`
	_, _ = o.Raw(sql, reason, op_id).Exec()

}

func GetOnProcess(bs_id int) (int, error) {
	var on_process int

	sql := `select on_process from pms_disaster_config where bs_id =?`
	o := orm.NewOrm()
	err := o.Raw(sql, bs_id).QueryRow(&on_process)
	if err != nil {
		return -1, err
	}

	return on_process, err
}

// func GetCurrentOpType(bs_id int) (string, error) {
// 	var on_switchover int
// 	var on_failover int
// 	var on_startmrp int
// 	var on_stopmrp int
// 	var on_startsnapshot int
// 	var on_stopsnapshot int
// 	sql := `select on_switchover, on_failover, on_startmrp, on_stopmrp, on_startsnapshot, on_stopsnapshot
// 			from pms_disaster_config where bs_id =?`
// 	o := orm.NewOrm()
// 	err := o.Raw(sql, bs_id).QueryRow(&on_switchover, &on_failover, &on_startmrp, &on_stopmrp, &on_startsnapshot, &on_stopsnapshot)
// 	if err != nil {
// 		return "", err
// 	} else {
// 		if on_switchover == 1 {
// 			return "SWITCHOVER", err
// 		} else if on_failover == 1 {
// 			return "FAILOVER", err
// 		} else if on_startmrp == 1 {
// 			return "MRP_START", err
// 		} else if on_stopmrp == 1 {
// 			return "MRP_STOP", err
// 		} else if on_startsnapshot == 1 {
// 			return "SNAPSHOT_START", err
// 		} else if on_stopsnapshot == 1 {
// 			return "SNAPSHOT_STOP", err
// 		} else {
// 			return "", err
// 		}
// 	}
// }

func GetCurrentOpId(bs_id int, op_type string) (int64, error) {
	var op_id int64
	sql := `select id from pms_opration where bs_id = ? and op_type = ? order by created desc limit 1`
	o := orm.NewOrm()
	err := o.Raw(sql, bs_id, op_type).QueryRow(&op_id)
	return op_id, err
}

type Process struct {
	Time         string
	Process_desc string
}

func GetOPProcessById(op_id int64) ([]*Process, error) {
	var pro []*Process

	sql := `select from_unixtime(created) as time, process_desc from pms_op_process where op_id = ? order by id`
	o := orm.NewOrm()
	_, err := o.Raw(sql, op_id).QueryRows(&pro)

	return pro, err
}

func GetOpResultById(op_id int64) (string, string, error) {
	var result string
	var reason string

	sql := `select result, reason from pms_opration where id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, op_id).QueryRow(&result, &reason)

	return result, reason, err
}
