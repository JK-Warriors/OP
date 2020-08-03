package oracle

import (
	//"fmt"
	"opms/models"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Oracle struct {
	Id       			int `orm:"pk;column(id);"`
	Db_Id   			int `orm:"column(db_id);"`
	Host   				string `orm:"column(host);"`
	Port   				string `orm:"column(port);"`
	Alias   			string `orm:"column(alias);"`
	Connect   			int `orm:"column(connect);"`
	Inst_Num   			string `orm:"column(inst_num);"`
	Inst_Name   		string `orm:"column(inst_name);"`
	Inst_Role   		string `orm:"column(inst_role);"`
	Inst_Status   		string `orm:"column(inst_status);"`
	Version   			string `orm:"column(version);"`
	Startup_Time   		string `orm:"column(startup_time);"`
	Host_Name   		string `orm:"column(host_name);"`
	Archiver 			string `orm:"column(archiver);"`
	Db_Name   			string `orm:"column(db_name);"`
	Db_Role   			string `orm:"column(db_role);"`
	Open_Mode   		string `orm:"column(open_mode);"`
	Protection_Mode   	string `orm:"column(protection_mode);"`
	Session_Total 		int    `orm:"column(session_total);"`
	Session_Actives 	int    `orm:"column(session_actives);"`
	Session_Waits 		int    `orm:"column(session_waits);"`
	Dg_Stats 			string  `orm:"column(dg_stats);"`
	Dg_Delay 			int    `orm:"column(dg_delay);"`
	Processes 			int    `orm:"column(processes);"`
	Flashback_On 		string `orm:"column(flashback_on);"`
	Flashback_Usage 	string `orm:"column(flashback_usage);"`
	Created  			int64  `orm:"column(created);"`
}

func (this *Oracle) TableName() string {
	return models.TableName("oracle_status")
}


type Tablespace struct {
	Id       			int `orm:"pk;column(id);"`
	Db_Id   			int `orm:"column(db_id);"`
	Host   				string `orm:"column(host);"`
	Port   				string `orm:"column(port);"`
	Alias   			string `orm:"column(alias);"`
	Tablespace_Name   	string `orm:"column(tablespace_name);"`
	Status   			string `orm:"column(status);"`
	Management   		string `orm:"column(management);"`
	Total_Size   		string `orm:"column(total_size);"`
	Used_Size   		string `orm:"column(used_size);"`
	Max_Rate   			string `orm:"column(max_rate);"`
	Created  			int64  `orm:"column(created);"`
}

func (this *Tablespace) TableName() string {
	return models.TableName("oracle_tablespace")
}

type Diskgroup struct {
	Id       			int `orm:"pk;column(id);"`
	Db_Id   			int `orm:"column(db_id);"`
	Host   				string `orm:"column(host);"`
	Port   				string `orm:"column(port);"`
	Alias   			string `orm:"column(alias);"`
	Diskgroup_Name   	string `orm:"column(diskgroup_name);"`
	State   			string `orm:"column(state);"`
	Type   				string `orm:"column(type);"`
	Total_Mb   			string `orm:"column(total_mb);"`
	Free_Mb   			string `orm:"column(free_mb);"`
	Used_Rate   		string `orm:"column(used_rate);"`
	Created  			int64  `orm:"column(created);"`
}
func (this *Diskgroup) TableName() string {
	return models.TableName("oracle_diskgroup")
}

func init() {
	orm.RegisterModel(new(Oracle))
	orm.RegisterModel(new(Tablespace))
	orm.RegisterModel(new(Diskgroup))
}


//获取oracle状态列表
func ListOracleStatus(condArr map[string]string, page int, offset int) (num int64, err error, oracle []Oracle) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("oracle_status"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("db_id")
	nums, errs := qs.Limit(offset, start).All(&oracle)
	return nums, errs, oracle
}


//统计数量
func CountOracle(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("oracle_status"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}


//获取oracle表空间列表
func ListTablespaces(condArr map[string]string, page int, offset int) (num int64, err error, tbs []Tablespace) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("oracle_tablespace"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("db_id")
	nums, errs := qs.Limit(offset, start).All(&tbs)
	return nums, errs, tbs
}

func CountTablespaces(condArr map[string]string) int64 {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("oracle_tablespace"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}



//获取oracle磁盘组列表
func ListDiskgroups(condArr map[string]string, page int, offset int) (num int64, err error, dgs []Diskgroup) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("oracle_diskgroup"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("db_id")
	nums, errs := qs.Limit(offset, start).All(&dgs)
	return nums, errs, dgs
}

func CountDiskgroups(condArr map[string]string) int64 {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("oracle_diskgroup"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}
