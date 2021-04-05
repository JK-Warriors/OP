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
	Session_Total 		int    `orm:"column(sessions);"`
	Repl	 			string  `orm:"column(repl);"`
	Repl_Delay	 		string  `orm:"column(repl_delay);"`
	Tablespace	 		string  `orm:"column(tablespace);"`
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
	
	num, _ := qs.SetCond(cond).Count()
	return num
}
