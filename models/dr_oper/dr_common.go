package dr_oper

import (
	"opms/utils"
	"strconv"
	"time"
	"fmt"
	"bytes"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    "golang.org/x/crypto/ssh"
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
	Is_Switch   int    `orm:"column(is_switch);"`
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
					d.is_switch
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

//获取Oracle容灾列表
func ListOracleDr(condArr map[string]string, page int, offset int) (num int64, err error, dr []Dr) {
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
				 and d.asset_type = 1
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

//统计oracle容灾数量
func CountOracleDrConfig(condArr map[string]string) int64 {
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
	cond = cond.And("asset_type", 1)
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

func GetIsShiftIP(bs_id int) (int) {
	var is_shift int =-1

	sql := `select is_shift from pms_dr_config where bs_id = ?`
	o := orm.NewOrm()
	_ = o.Raw(sql, bs_id).QueryRow(&is_shift)
	return is_shift 	 	
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
	Asset_Type      int    `orm:"column(asset_type);"`
	Connect         int    `orm:"column(connect);"`
	Instance_name   string `orm:"column(instance_name);"`
	Db_Name         string `orm:"column(db_name);"`
	Host            string `orm:"column(host);"`
	Port            string `orm:"column(port);"`
	Role            string `orm:"column(inst_role);"`
	Version         string `orm:"column(version);"`
	Open_Mode       string `orm:"column(open_mode);"`
	Flashback_On    string `orm:"column(flashback_on);"`
	Flashback_Usage string `orm:"column(flashback_usage);"`
	Created         int64  `orm:"column(created);"`
}

func GetOracleBasicInfo(db_id int) (OracleInstance, error) {
	var ora_instance OracleInstance

	sql := `select c.id, c.asset_type, s.connect, c.instance_name, c.db_name, c.host, c.port, s.inst_role, version, open_mode, flashback_on, flashback_usage, s.created 
			from pms_asset_config c, pms_oracle_status s 
			where c.id = s.db_id and s.db_id = ?`
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
	Mrp_Status   string    `orm:"column(mrp_status);"`
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



type MySQLInstance struct {
	Id              int    `orm:"pk;column(id);"`
	Asset_Type      int    `orm:"column(asset_type);"`
	Connect         int    `orm:"column(connect);"`
	Instance_name   string `orm:"column(instance_name);"`
	Db_Name         string `orm:"column(db_name);"`
	Host            string `orm:"column(host);"`
	Port            string `orm:"column(port);"`
	Role            string `orm:"column(role);"`
	Version         string `orm:"column(version);"`
	Created         int64  `orm:"column(created);"`
}

func GetMySQLBasicInfo(db_id int) (MySQLInstance, error) {
	var asset MySQLInstance

	sql := `select c.id, c.asset_type, s.connect, c.instance_name, c.db_name, c.host, c.port, s.role, version, s.created 
			from pms_asset_config c, pms_mysql_status s 
			where c.id = s.db_id  and s.db_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&asset)

	return asset, err
}

type DrMySQLPrimary struct {
	Id             	 int    `orm:"pk;column(id);"`
	Dr_Id          	 int    `orm:"column(dr_id);"`
	Db_Id		   	 int    `orm:"column(db_id);"`
	Read_Only        string `orm:"column(read_only);"`
	Gtid_Mode  		 string `orm:"column(gtid_mode);"`
	M_Binlog_File    string `orm:"column(master_binlog_file);"`
	M_Binlog_Pos     string `orm:"column(master_binlog_pos);"`
	M_Binlog_Space   int64	`orm:"column(master_binlog_space);"`
	Created        	 int64  `orm:"column(created);"`
}

func GetDrMySqlPrimaryInfo(db_id int) (DrMySQLPrimary, error) {
	var dis_pri DrMySQLPrimary

	sql := `select id, dr_id, db_id, read_only, gtid_mode, master_binlog_file, master_binlog_pos, master_binlog_space, created
			from pms_dr_mysql_p where db_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&dis_pri)

	return dis_pri, err
}


type DrMySQLStandby struct {
	Id        	 			int    `orm:"pk;column(id);"`
	Dr_Id       		 	int    `orm:"column(dr_id);"`
	DB_Id        			int    `orm:"column(db_id);"`
	Read_Only        		string `orm:"column(read_only);"`
	Gtid_Mode  		 		string `orm:"column(gtid_mode);"`
	M_Server  		 		string `orm:"column(master_server);"`
	M_Port  		 		string `orm:"column(master_port);"`
	Slave_IO_Run  		 	string `orm:"column(slave_io_run);"`
	Slave_SQL_Run  		 	string `orm:"column(slave_sql_run);"`
	Delay  		 			int 	`orm:"column(delay);"`
	C_Binlog_File  		 	string `orm:"column(current_binlog_file);"`
	C_Binlog_Pos  		 	string `orm:"column(current_binlog_pos);"`
	M_Binlog_File    		string `orm:"column(master_binlog_file);"`
	M_Binlog_Pos     		string `orm:"column(master_binlog_pos);"`
	M_Binlog_Space   		int64  `orm:"column(master_binlog_space);"`
	Created      			int64  `orm:"column(created);"`
}

func GetDrMySqlStandbyInfo(db_id int) (DrMySQLStandby, error) {
	var dis_sta DrMySQLStandby

	sql := `select id, dr_id, db_id, read_only, gtid_mode, master_server, master_port, slave_io_run, slave_sql_run, delay, 
				current_binlog_file, current_binlog_pos, master_binlog_file, master_binlog_pos, master_binlog_space, created
			from pms_dr_mysql_s where db_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&dis_sta)

	return dis_sta, err
}

type MSSqlInstance struct {
	Id              int    `orm:"pk;column(id);"`
	Asset_Type      int    `orm:"column(asset_type);"`
	Connect         int    `orm:"column(connect);"`
	Instance_name   string `orm:"column(instance_name);"`
	Db_Name         string `orm:"column(db_name);"`
	Host            string `orm:"column(host);"`
	Port            string `orm:"column(port);"`
	Role            string `orm:"column(role);"`
	Version         string `orm:"column(version);"`
	Created         int64  `orm:"column(created);"`
}

func GetMSSqlBasicInfo(db_id int) (MSSqlInstance, error) {
	var asset MSSqlInstance

	sql := `select c.id, c.asset_type, s.connect, c.instance_name, c.db_name, c.host, c.port, s.role, version, s.created 
			from pms_asset_config c, pms_mssql_status s 
			where c.id = s.db_id  and s.db_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&asset)

	return asset, err
}

type DrMSSqlPrimary struct {
	Id             	 int    `orm:"pk;column(id);"`
	Dr_Id          	 int    `orm:"column(dr_id);"`
	Db_Id		   	 int    `orm:"column(db_id);"`
	Database_Id        	int `orm:"column(database_id);"`
	Db_Name  		 	string `orm:"column(db_name);"`
	Role    			int 	`orm:"column(role);"`
	State     			int 	`orm:"column(state);"`
	State_Desc   		string	`orm:"column(state_desc);"`
	Safety_Level   		int		`orm:"column(safety_level);"`
	Partner_Name   		string	`orm:"column(partner_name);"`
	Partner_Instance   	string	`orm:"column(partner_instance);"`
	Failover_Lsn   		int64	`orm:"column(failover_lsn);"`
	Connection_Timeout  int		`orm:"column(connection_timeout);"`
	Redo_Queue   		int		`orm:"column(redo_queue);"`
	End_Of_Log_Lsn   	int64	`orm:"column(end_of_log_lsn);"`
	Replication_Lsn   	int64	`orm:"column(replication_lsn);"`
	Created        	 	int64  `orm:"column(created);"`
}

func GetDrMSSqlPrimaryInfo(db_id int) (DrMSSqlPrimary, error) {
	var dis_pri DrMSSqlPrimary

	sql := `select id, dr_id, db_id, database_id, db_name, role, state, state_desc, safety_level, partner_name, partner_instance, 
				failover_lsn, connection_timeout, redo_queue, end_of_log_lsn, replication_lsn, created
			from pms_dr_mssql_p where db_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&dis_pri)

	return dis_pri, err
}


type DrMSSqlStandby struct {
	Id             	 int    `orm:"pk;column(id);"`
	Dr_Id          	 int    `orm:"column(dr_id);"`
	Db_Id		   	 int    `orm:"column(db_id);"`
	Database_Id        	int `orm:"column(database_id);"`
	Db_Name  		 	string `orm:"column(db_name);"`
	Master_Server  		string `orm:"column(master_server);"`
	Master_Port  		string `orm:"column(master_port);"`
	Role    			int 	`orm:"column(role);"`
	State     			int 	`orm:"column(state);"`
	State_Desc   		string	`orm:"column(state_desc);"`
	Safety_Level   		int		`orm:"column(safety_level);"`
	Partner_Name   		string	`orm:"column(partner_name);"`
	Partner_Instance   	string	`orm:"column(partner_instance);"`
	Failover_Lsn   		int64	`orm:"column(failover_lsn);"`
	Connection_Timeout  int		`orm:"column(connection_timeout);"`
	Redo_Queue   		int		`orm:"column(redo_queue);"`
	End_Of_Log_Lsn   	int64	`orm:"column(end_of_log_lsn);"`
	Replication_Lsn   	int64	`orm:"column(replication_lsn);"`
	Created        	 	int64  `orm:"column(created);"`
}

func GetDrMSSqlStandbyInfo(db_id int) (DrMSSqlStandby, error) {
	var dis_sta DrMSSqlStandby

	sql := `select id, dr_id, db_id, database_id, db_name, master_server, master_port, role, state, state_desc, safety_level, partner_name, partner_instance, 
				failover_lsn, connection_timeout, redo_queue, end_of_log_lsn, replication_lsn, created
			from pms_dr_mssql_s where db_id = ?`
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





//切换IP
func SwitchIPs(op_id int64, dr_id int, asset_type int, pri_id int, sta_id int) int{
	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "切换IP开始")

	//Get switch ips
	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "开始获取需要切换的IP")
	ips,err := GetIps(dr_id)
	if err != nil || ips == ""{
		Update_OP_Reason(op_id, "获取需要切换的IP失败")
		Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "获取需要切换的IP失败")
		utils.LogDebug("GetIps failed: " + err.Error())
		return -1
	}

	//get network card
	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "开始获取绑定的目标端网卡")
	s_card,err := GetStaNetcard(dr_id)
	if err != nil || s_card == ""{
		Update_OP_Reason(op_id, "获取绑定的目标端网卡失败")
		Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "获取绑定的目标端网卡失败")
		utils.LogDebug("Get network card failed: " + err.Error())
		return -1
	}

	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "开始源端解绑IP")
	result := Unbind_IPs(pri_id, ips)
	if result == -1{
		Update_OP_Reason(op_id, "源端解绑IP失败")
		Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "源端解绑IP失败")
		utils.LogDebug("Unbind ips failed")
		return -1
	}
	
	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "开始目标端绑定IP")
	result = Bind_IPs(sta_id, ips, s_card)
	if result == -1{
		Update_OP_Reason(op_id, "目标端绑定IP失败")
		Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "目标端绑定IP失败")
		utils.LogDebug("Bind ips failed")
		return -1
	}

	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "切换IP成功")
	return 1
}

//切换IP
func FailoverIPs(op_id int64, dr_id int, asset_type int, sta_id int) int{
	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "绑定IP开始")

	//Get switch ips
	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "开始获取需要绑定的IP")
	ips,err := GetIps(dr_id)
	if err != nil || ips == ""{
		Update_OP_Reason(op_id, "获取需要绑定的IP失败")
		Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "获取需要绑定的IP失败")
		utils.LogDebug("GetIps failed: " + err.Error())
		return -1
	}

	//get network card
	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "开始获取绑定的目标端网卡")
	s_card,err := GetStaNetcard(dr_id)
	if err != nil || s_card == ""{
		Update_OP_Reason(op_id, "获取绑定的目标端网卡失败")
		Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "获取绑定的目标端网卡失败")
		utils.LogDebug("Get network card failed: " + err.Error())
		return -1
	}

	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "开始目标端绑定IP")
	result := Bind_IPs(sta_id, ips, s_card)
	if result == -1{
		Update_OP_Reason(op_id, "目标端绑定IP失败")
		Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "目标端绑定IP失败")
		utils.LogDebug("Bind ips failed")
		return -1
	}

	Log_OP_Process(op_id, dr_id, asset_type, "SWITCHOVER", "绑定IP成功")
	return 1
}

func GetIps(dr_id int) (string, error){
	var ips string
	sql := `select shift_vips from pms_dr_config where bs_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, dr_id).QueryRow(&ips)
	if err != nil {
		utils.LogDebug("GetIps failed: " + err.Error())
		return ips, err
	}
	return ips, nil
}

func GetStaNetcard(dr_id int) (string, error){
	var netcard string
	sql := `select trim(CASE is_switch
				WHEN 0 THEN network_s
				ELSE network_p
			END) as netcard
			from pms_dr_config
			where bs_id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, dr_id).QueryRow(&netcard)
	if err != nil {
		utils.LogDebug("GetStaNetcard failed: " + err.Error())
		return netcard, err
	}
	return netcard, nil
}


type OsInfo struct {
	DB_Id        	int    	`orm:"pk;column(db_id);"`
	Host       	 	string 	`orm:"column(host);"`
	Port       	 	string 	`orm:"column(os_port);"`
	Os_Type     	string  `orm:"column(os_type);"`
	Os_Protocol     string 	`orm:"column(os_protocol);"`
	Os_Username   	string  `orm:"column(os_username);"`
	Os_Password   	string 	`orm:"column(os_password);"`
}

func Unbind_IPs(db_id int, ips string) int{
	var result int =1
	var osinfo OsInfo
	sql := `select id, host, os_port, os_type, os_protocol, os_username, os_password  from pms_asset_config where id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&osinfo)

	utils.LogDebugf("pri_id: %d", db_id)
	utils.LogDebugf("Host: %s", osinfo.Host)
	
	// 建立SSH客户端连接
	host := fmt.Sprintf("%s:%s", osinfo.Host, osinfo.Port)
	utils.LogDebugf("SSH host: %s", host)
    client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
        User:            osinfo.Os_Username,
        Auth:            []ssh.AuthMethod{ssh.Password(osinfo.Os_Password)},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    })
    if err != nil {
        utils.LogDebugf("SSH dial error: %s", err.Error())
		result = -1
    }

	iplist := strings.Split(ips, ",")
	if len(iplist) > 0{
		for _, ip := range iplist {
			//ip := "192.168.153.79"
			//get netcard
			//ifconfig | awk '/192.168.153.79/{print a;}{a=$0}' | awk '{print $1}' | awk -F':' '{print $1":"$2}'
			command := fmt.Sprintf(`ifconfig | awk '/%s/{print a;}{a=$0}' | awk '{print $1}' | awk -F':' '{print $1":"$2}'`,ip)
			utils.LogDebugf("command: %s", command)

			out, err := runCommand(client, command)
			if err != nil {
				utils.LogDebugf("runCommand error: %s", err.Error())
				result = -1
				continue
			}else{
				netcard := strings.Replace(out, "\n", "", -1)
				utils.LogDebugf("netcard: %s", netcard)
				if netcard == ""{
					utils.LogDebugf("IP %s was not bind, skip", ip)
					result = -1
					continue
				}else{
					//ip_cmd = "ifconfig ens33:1 down"
					command = fmt.Sprintf(`ifconfig %s down`, netcard)
					utils.LogDebugf("command: %s", command)

					out, err = runCommand(client, command)
					if err != nil {
						utils.LogDebugf("runCommand error: %s", err.Error())
					}
					utils.LogDebugf("out: %s", out)

					//check ip down
					command = fmt.Sprintf(`ifconfig | grep %s | wc -l`, ip)
					utils.LogDebugf("command: %s", command)

					out, err = runCommand(client, command)
					if err != nil {
						utils.LogDebugf("Check ip down error: %s", err.Error())
					}else{
						checkvalue := strings.Replace(out, "\n", "", -1)
						utils.LogDebugf("checkvalue: %s", checkvalue)
						if checkvalue != "0"{
							utils.LogDebugf("Check ip down failed")
							result = -1
							continue
						}
					}
				}
			}
		}
	}else{
		utils.LogDebug("No switch ip exists")
		return -1
	}
	return result
}

func Bind_IPs(db_id int, ips string, netcard string) int{
	var result int = 1

	var osinfo OsInfo
	sql := `select id, host, os_port, os_type, os_protocol, os_username, os_password  from pms_asset_config where id = ?`
	o := orm.NewOrm()
	err := o.Raw(sql, db_id).QueryRow(&osinfo)

	utils.LogDebugf("pri_id: %d", db_id)
	utils.LogDebugf("Host: %s", osinfo.Host)
	
	// 建立SSH客户端连接
	host := fmt.Sprintf("%s:%s", osinfo.Host, osinfo.Port)
	utils.LogDebugf("SSH host: %s", host)
    client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
        User:            osinfo.Os_Username,
        Auth:            []ssh.AuthMethod{ssh.Password(osinfo.Os_Password)},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    })
    if err != nil {
        utils.LogDebugf("SSH dial error: %s", err.Error())
		result = -1
    }

	iplist := strings.Split(ips, ",")
	if len(iplist) > 0{
		for _, ip := range iplist {
			//get netmask
			command := fmt.Sprintf(`ifconfig -a %s | grep 'netmask' | awk '{print $4}'`, netcard)
			utils.LogDebugf("command: %s", command)
			out, err := runCommand(client, command)
			if err != nil {
				utils.LogDebugf("get netmask error: %s", err.Error())
				result = -1
				continue
			}
			netmask := strings.Replace(out, "\n", "", -1)

			//get gateway
			command = fmt.Sprintf(`route -n | grep %s | awk '{print $2}' | grep -v '0.0.0.0'`, netcard)
			utils.LogDebugf("command: %s", command)
			out, err = runCommand(client, command)
			if err != nil {
				utils.LogDebugf("get gateway error: %s", err.Error())
				result = -1
				continue
			}
			gateway := strings.Replace(out, "\n", "", -1)

			//get netcard bind ip count
			//ifconfig | grep ens33 | wc -l
			command = fmt.Sprintf(`ifconfig | grep %s | wc -l`, netcard)
			utils.LogDebugf("command: %s", command)

			out, err = runCommand(client, command)
			if err != nil {
				utils.LogDebugf("runCommand error: %s", err.Error())
				result = -1
				continue
			}
			
			bindcount,err := strconv.Atoi(strings.Replace(out, "\n", "", -1))
			if err != nil {
				utils.LogDebugf("Get netcard bind ip count error: %s", err.Error())
				result = -1
				continue
			}
			utils.LogDebugf("bindcount: %s", bindcount)
			
			//bind ip
			command = fmt.Sprintf(`ifconfig %s:%d %s netmask %s`, netcard, bindcount, ip, netmask)
			utils.LogDebugf("command: %s", command)

			out, err = runCommand(client, command)
			if err != nil {
				utils.LogDebugf("runCommand error: %s", err.Error())
				result = -1
				continue
			}

			//arp
			command = fmt.Sprintf(`arping -U -c 1 -I %s -s %s %s`, netcard, ip, gateway)
			utils.LogDebugf("command: %s", command)

			out, err = runCommand(client, command)
			if err != nil {
				utils.LogDebugf("runCommand error: %s", err.Error())
				result = -1
				continue
			}

			//check ip bind
			command = fmt.Sprintf(`ifconfig | grep %s | wc -l`, ip)
			utils.LogDebugf("command: %s", command)

			out, err = runCommand(client, command)
			if err != nil {
				utils.LogDebugf("Check ip bind error: %s", err.Error())
			}else{
				checkvalue := strings.Replace(out, "\n", "", -1)
				utils.LogDebugf("checkvalue: %s", checkvalue)
				if checkvalue == "0"{
					utils.LogDebugf("Check ip bind failed")
					result = -1
					continue
				}
			}
		}
	}else{
		utils.LogDebug("No switch ip exists")
		return -1
	}
	
	return result
}

func runCommand(client *ssh.Client, command string) (stdout string, err error) {
	session, err := client.NewSession()
	if err != nil {
		//log.Print(err)
		return
	}
	defer session.Close()

	var buf bytes.Buffer
	session.Stdout = &buf
	err = session.Run(command)
	if err != nil {
		//log.Print(err)
		return
	}
	stdout = string(buf.Bytes())

	return
}