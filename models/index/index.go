package index

import (
	//"fmt"
	"opms/models"
	// . "opms/models/asset"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DbStatus struct {
	Id       			int `orm:"pk;column(id);"`
	Asset_Id   			int `orm:"column(asset_id);"`
	Asset_Type   		int `orm:"column(asset_type);"`
	Host   				string `orm:"column(host);"`
	Port   				string `orm:"column(port);"`
	Alias   			string `orm:"column(alias);"`
	Role   				string `orm:"column(role);"`
	Version   			string `orm:"column(version);"`
	Connect   			int `orm:"column(connect);"`
	Connect_Tips   		string `orm:"column(connect_tips);"`
	Sess_Total 			int    `orm:"column(sessions);"`
	Sess_Total_Tips 	string    `orm:"column(sessions_tips);"`
	Sess_Actives 		int    		`orm:"column(actives);"`
	Sess_Actives_Tips 	string    	`orm:"column(actives_tips);"`
	Sess_Waits 			int    		`orm:"column(waits);"`
	Sess_Waits_Tips 	string    	`orm:"column(waits_tips);"`
	Process 			int    		`orm:"column(process);"`
	Process_Tips 		string    	`orm:"column(process_tips);"`
	Repl	 			int  		`orm:"column(repl);"`
	Repl_Tips 			string    	`orm:"column(repl_tips);"`
	Repl_Delay	 		int  		`orm:"column(repl_delay);"`
	Repl_Delay_Tips 	string    	`orm:"column(repl_delay_tips);"`
	Tablespace	 		int  		`orm:"column(tablespace);"`
	Tablespace_Tips 	string    	`orm:"column(tablespace_tips);"`
	Diskgroup	 		int  		`orm:"column(diskgroup);"`
	Diskgroup_Tips 		string    	`orm:"column(diskgroup_tips);"`
	Flashback_Space	 		int  		`orm:"column(flashback_space);"`
	Flashback_Space_Tips 	string    	`orm:"column(flashback_space_tips);"`
	Load	 				int  		`orm:"column(load);"`
	Load_Tips 				string    	`orm:"column(load_tips);"`
	Cpu	 					int  		`orm:"column(cpu);"`
	Cpu_Tips 				string    	`orm:"column(cpu_tips);"`
	Memory	 				int  		`orm:"column(memory);"`
	Memory_Tips 			string    	`orm:"column(memory_tips);"`
	IO	 					int  		`orm:"column(io);"`
	IO_Tips 				string    	`orm:"column(io_tips);"`
	Net	 					int  		`orm:"column(net);"`
	Net_Tips 				string    	`orm:"column(net_tips);"`
	Score		 		string  `orm:"column(score);"`
	Created  			int64  `orm:"column(created);"`
}

func (this *DbStatus) TableName() string {
	return models.TableName("asset_status")
}
func init() {
	orm.RegisterModel(new(DbStatus))
}

//获取db状态列表
func ListDbStatus(condArr map[string]string, page int, offset int) (num int64, err error, db []DbStatus) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_status"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	cond = cond.And("asset_type__lt", 99)

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("asset_id")
	nums, errs := qs.Limit(offset, start).All(&db)
	return nums, errs, db
}


//统计db数量
func CountDb(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("asset_status"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}

	cond = cond.And("asset_type__lt", 99)

	num, _ := qs.SetCond(cond).Count()
	return num
}
