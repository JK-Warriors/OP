package dr_oper

import (
	"opms/utils"
	"strconv"
	"time"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Dr struct {
	Id          int    `orm:"pk;column(id);"`
	Bs_Name     string `orm:"column(bs_name);"`
	Db_Id_P     int    `orm:"column(db_id_p);"`
	Db_Type_P   int    `orm:"column(db_type_p);"`
	Host_P      string `orm:"column(host_p);"`
	Port_P      int    `orm:"column(port_p);"`
	Alias_P     string `orm:"column(alias_p);"`
	Inst_Name_P string `orm:"column(inst_name_p);"`
	Db_Name_P   string `orm:"column(db_name_p);"`
	Db_Id_S     int    `orm:"column(db_id_s);"`
	Db_Type_S   int    `orm:"column(db_type_s);"`
	Host_S      string `orm:"column(host_s);"`
	Port_S      int    `orm:"column(port_s);"`
	Alias_S     string `orm:"column(alias_s);"`
	Inst_Name_S string `orm:"column(inst_name_s);"`
	Db_Name_S   string `orm:"column(db_name_s);"`
	Is_Shift    int    `orm:"column(is_shift);"`
}

//获取容灾列表
func ListDr(condArr map[string]string, page int, offset int) (num int64, err error, dr []Dr) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select d.bs_id as id, 
					d.bs_name,
					d.db_id_p, 
					pp.asset_type as db_type_p,
					pp.host as host_p,
					pp.port as port_p, 
					pp.alias as alias_p, 
					pp.instance_name as inst_name_p, 
					pp.db_name as db_name_p, 
					d.db_id_s, 
					ps.asset_type as db_type_s,
					ps.host as host_s,
					ps.port as port_s, 
					ps.alias as alias_s, 
					ps.instance_name as inst_name_s, 
					ps.db_name as db_name_s, 
					d.is_shift
				from pms_dr_config d
				left join pms_asset_config pp on d.db_id_p = pp.id
				left join pms_asset_config ps on d.db_id_s = ps.id
				where d.is_delete = 0
				 and d.status = 1`

	if condArr["search_name"] != "" {
		sql = sql + " and (b.bs_name like '%" + condArr["search_name"] + "%')"
	}

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	sql = sql + " order by id"
	sql = sql + " limit " + strconv.Itoa(offset) + " offset " + strconv.Itoa(start)
	nums, err := o.Raw(sql).QueryRows(&dr)
	if err != nil {
		utils.LogDebug("Get ListDr failed:" + err.Error())
	}
	//utils.LogDebugf("%+v", dr)
	return nums, err, dr
}

//统计数量
func CountDrconfig(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("pms_dr_config")
	cond := orm.NewCondition()

	if condArr["asset_type"] != "" {
		cond = cond.And("asset_type", condArr["asset_type"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	cond = cond.And("status", 1)
	cond = cond.And("is_delete", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

func CheckDrConfig(bs_id int) (int, error) {
	var cfg_count int

	sql := `select count(1) from pms_dr_config where bs_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, bs_id).QueryRow(&cfg_count)
	return cfg_count, err
}

func GetAssetType(bs_id int) (int) {
	var asset_type int =-1

	sql := `select asset_type from pms_dr_config where bs_id = ?`
	o := orm.NewOrm()
	_ = o.Raw(sql, bs_id).QueryRow(&asset_type)
	return asset_type 	 	
}

func GetPrimaryDBId(bs_id int) (int, error) {
	var pri_dbid int

	sql := `select CASE is_switch
				WHEN 0 THEN db_id_p
				ELSE db_id_s
			END as pri_dbid
			from pms_dr_config
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
			from pms_dr_config
			where bs_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, bs_id).QueryRow(&sta_dbid)
	return sta_dbid, err
}

func GetDsn(db_id int, asset_type int) (string, error) {
	var dsn string
	var sql string
	if asset_type == 1 {
		sql = `select concat("oracle://",username,":",password ,"@" , host , ":" , port , "/" , instance_name , "?sysdba=1") as dsn 
				from pms_asset_config where id = ? and asset_type = ?`
	} else if asset_type == 2 {
		sql = `select concat(username,":",password,"@tcp(",host,":",port,")/",db_name,"?charset=utf8") from pms_asset_config where id = ? and asset_type = ?`
	} else if asset_type == 3 {
		sql = `select concat("server=",host,"\\",instance_name,";port",port,";database=",case db_name when "" then "master" end,";user id=",username,";password=",password,";encrypt=disable") from pms_asset_config where id = ? and asset_type = ?`
	} else {
		sql = `select "" from pms_asset_config where id = ? and asset_type = ?`
	}
	
	o := orm.NewOrm()
	err := o.Raw(sql, db_id, asset_type).QueryRow(&dsn)
	return dsn, err
}


func OperationLock(bs_id int, op_type string) error {
	utils.LogDebug("Lock the process status in pms_dr_config.")
	o := orm.NewOrm()
	var sql string
	if op_type == "SWITCHOVER" {
		sql = `update pms_dr_config set on_process = 1, on_switchover = 1 where bs_id= ?`
	} else if op_type == "FAILOVER" {
		sql = `update pms_dr_config set on_process = 1, on_failover = 1 where bs_id= ?`
	} else if op_type == "STARTSYNC" {
		sql = `update pms_dr_config set on_process = 1, on_startsync = 1 where bs_id= ?`
	} else if op_type == "STOPSYNC" {
		sql = `update pms_dr_config set on_process = 1, on_stopsync = 1 where bs_id= ?`
	} else if op_type == "STARTREAD" {
		sql = `update pms_dr_config set on_process = 1, on_startread = 1 where bs_id= ?`
	} else if op_type == "STOPREAD" {
		sql = `update pms_dr_config set on_process = 1, on_stopread = 1 where bs_id= ?`
	} else if op_type == "STARTSNAPSHOT" {
		sql = `update pms_dr_config set on_process = 1, on_startsnapshot = 1 where bs_id= ?`
	} else if op_type == "STOPSNAPSHOT" {
		sql = `update pms_dr_config set on_process = 1, on_stopsnapshot = 1 where bs_id= ?`
	} else if op_type == "STARTFLASHBACK" {
		sql = `update pms_dr_config set on_process = 1, on_startflashback = 1 where bs_id= ?`
	} else if op_type == "STOPFLASHBACK" {
		sql = `update pms_dr_config set on_process = 1, on_stopflashback = 1 where bs_id= ?`
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
	utils.LogDebug("Unlock the process status in pms_dr_config.")
	if op_type == "SWITCHOVER" {
		sql = `update pms_dr_config set on_process = 0, on_switchover = 0 where bs_id= ?`
	} else if op_type == "FAILOVER" {
		sql = `update pms_dr_config set on_process = 0, on_failover = 0 where bs_id= ?`
	} else if op_type == "STARTSYNC" {
		sql = `update pms_dr_config set on_process = 0, on_startsync = 0 where bs_id= ?`
	} else if op_type == "STOPSYNC" {
		sql = `update pms_dr_config set on_process = 0, on_stopsync = 0 where bs_id= ?`
	} else if op_type == "STARTREAD" {
		sql = `update pms_dr_config set on_process = 0, on_startread = 1 where bs_id= ?`
	} else if op_type == "STOPREAD" {
		sql = `update pms_dr_config set on_process = 0, on_stopread = 1 where bs_id= ?`
	} else if op_type == "STARTSNAPSHOT" {
		sql = `update pms_dr_config set on_process = 0, on_startsnapshot = 0 where bs_id= ?`
	} else if op_type == "STOPSNAPSHOT" {
		sql = `update pms_dr_config set on_process = 0, on_stopsnapshot = 0 where bs_id= ?`
	} else if op_type == "STARTFLASHBACK" {
		sql = `update pms_dr_config set on_process = 0, on_startflashback = 0 where bs_id= ?`
	} else if op_type == "STOPFLASHBACK" {
		sql = `update pms_dr_config set on_process = 0, on_stopflashback = 0 where bs_id= ?`
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

func Init_OP_Instance(op_id int64, bs_id int, asset_type int, op_type string) error {
	o := orm.NewOrm()
	var sql string
	//将之前的操作记录移入his表
	MoveOpRecordToHis(bs_id, op_type)

	//开始新的操作初始化
	utils.LogDebugf("Initialize opration instance for business %d.", bs_id)

	sql = `insert into pms_opration(id, bs_id, asset_type, op_type, created) values(?, ?, ?, ?, ?)`
	_, err := o.Raw(sql, op_id, bs_id, asset_type, op_type, time.Now().Unix()).Exec()
	if err == nil {
		utils.LogDebug("Init the opration successfully.")
	} else {
		utils.LogDebug("Init the opration failed: " + err.Error())
	}
	return err
}

func Log_OP_Process(op_id int64, bs_id int, asset_type int, op_type string, process_desc string) error {
	o := orm.NewOrm()
	var sql string

	sql = `insert into pms_op_process(op_id, bs_id, asset_type, process_type, process_desc, created) values (?, ?, ?, ?, ?, ?)`
	_, err := o.Raw(sql, op_id, bs_id, asset_type, op_type, process_desc, time.Now().Unix()).Exec()

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

	utils.LogDebug("Update switch flag in pms_dr_config in progress...")

	// get current switch flag
	sql = `select is_switch from pms_dr_config where bs_id= ?`
	err := o.Raw(sql, bs_id).QueryRow(&is_switch)

	utils.LogDebug("The current switch flag is: " + strconv.Itoa(is_switch))

	if is_switch == 0 {
		sql = `update pms_dr_config set is_switch = 1 where bs_id = ?`
	} else {
		sql = `update pms_dr_config set is_switch = 0 where bs_id = ?`
	}

	_, err = o.Raw(sql, bs_id).Exec()
	if err == nil {
		utils.LogDebug("Update switch flag in pms_dr_config successfully.")
	} else {
		utils.LogDebug("Update switch flag in pms_dr_config failed: " + err.Error())
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

	sql := `select on_process from pms_dr_config where bs_id =?`
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
// 			from pms_dr_config where bs_id =?`
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

type OracleInstance struct {
	Id              int    `orm:"pk;column(id);"`
	Asset_Type         int    `orm:"column(asset_type);"`
	Connect         int    `orm:"column(connect);"`
	Instance_name   string `orm:"column(instance_name);"`
	Db_Name         string `orm:"column(db_name);"`
	Host            string `orm:"column(host);"`
	Role            string `orm:"column(role);"`
	Version         string `orm:"column(version);"`
	Open_Mode       string `orm:"column(open_mode);"`
	Flashback_On    string `orm:"column(flashback_on);"`
	Flashback_Usage string `orm:"column(flashback_usage);"`
	Created         int64  `orm:"column(created);"`
}

func GetOracleBasicInfo(db_id int) (OracleInstance, error) {
	var ora_instance OracleInstance

	sql := `select c.id, c.asset_type, s.connect, c.instance_name, c.db_name, c.host, s.role, version, open_mode, flashback_on, flashback_usage, s.created 
			from pms_asset_config c, pms_db_status s 
			where c.id = s.id and s.id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&ora_instance)

	return ora_instance, err
}

type DrPrimary struct {
	DB_Id          int    `orm:"pk;column(db_id);"`
	Check_Seq      string `orm:"column(check_seq);"`
	Dest_Id        int    `orm:"column(dest_id);"`
	Transmit_Mode  string `orm:"column(transmit_mode);"`
	Thread         int    `orm:"column(thread);"`
	Sequence       int    `orm:"column(sequence);"`
	Curr_Scn       int64  `orm:"column(curr_scn);"`
	Curr_Db_Time   string `orm:"column(curr_db_time);"`
	Archived_delay int    `orm:"column(archived_delay);"`
	Applied_delay  int    `orm:"column(applied_delay);"`
	Created        int64  `orm:"column(created);"`
}

func GetPrimaryDrInfo(db_id int) (DrPrimary, error) {
	var dis_pri DrPrimary

	sql := `select db_id, check_seq, dest_id, transmit_mode, thread, sequence, curr_scn, curr_db_time, archived_delay, applied_delay, created
			from pms_dr_pri_status where db_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&dis_pri)

	return dis_pri, err
}

type DrStandby struct {
	DB_Id        int    `orm:"pk;column(db_id);"`
	Thread       int    `orm:"column(thread);"`
	Sequence     int    `orm:"column(sequence);"`
	Block        string `orm:"column(block);"`
	Delay_Mins   int    `orm:"column(delay_mins);"`
	Apply_Rate   string `orm:"column(apply_rate);"`
	Curr_Scn     int64  `orm:"column(curr_scn);"`
	Curr_Db_Time string `orm:"column(curr_db_time);"`
	Mrp_Status   int    `orm:"column(mrp_status);"`
	Created      int64  `orm:"column(created);"`
}

func GetStandbyDrInfo(db_id int) (DrStandby, error) {
	var dis_sta DrStandby

	sql := `select db_id, thread, sequence, block, delay_mins, apply_rate, curr_scn, curr_db_time, mrp_status, created 
			from pms_dr_sta_status where db_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&dis_sta)

	return dis_sta, err
}


func GetChangeMasterCmd(db_id int, log_file string, log_pos string) (string, error) {
	var cmd string
	sql := fmt.Sprintf(`select concat("change master to master_host='",host,"',master_port=",port,",master_user='",username, "',master_password='",password, "',master_log_file='","%s","',master_log_pos=","%s") as cmd 
	from pms_asset_config where id = ?`, log_file, log_pos)

	//utils.LogDebugf("[Info] GetChangeMasterCmd: %s", sql)
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&cmd)
	return cmd, err
}

func GetMirrorDbname(dr_id int) (string, error) {
	var dbname string
	sql := `select db_name from pms_dr_config where bs_id = ?`

	o := orm.NewOrm()
	err := o.Raw(sql, dr_id).QueryRow(&dbname)
	return dbname, err
}
