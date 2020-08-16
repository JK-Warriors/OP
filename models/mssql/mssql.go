package mssql

import (
	//"fmt"
	"opms/models"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Mssql struct {
	Id       					int `orm:"pk;column(id);"`
	Db_Id   					int `orm:"column(db_id);"`
	Host   						string `orm:"column(host);"`
	Port   						string `orm:"column(port);"`
	Alias   					string `orm:"column(alias);"`
	Connect   					int `orm:"column(connect);"`
	Role   						string `orm:"column(role);"`
	Uptime   					string `orm:"column(uptime);"`
	Version   					string `orm:"column(version);"`
	Lock_timeout   				string `orm:"column(lock_timeout);"`
	Trancount   				string `orm:"column(trancount);"`
	Max_Connections   			string `orm:"column(max_connections);"`
	Processes   				string `orm:"column(processes);"`
	Processes_Running   		string `orm:"column(processes_running);"`
	Processes_Waits 			string `orm:"column(processes_waits);"`
	Created  					int64  `orm:"column(created);"`
}


func (this *Mssql) TableName() string {
	return models.TableName("mssql_status")
}


func init() {
	orm.RegisterModel(new(Mssql))
}


//获取sqlserver状态列表
func ListMssqlStatus(condArr map[string]string, page int, offset int) (num int64, err error, mssql []Mssql) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("mssql_status"))
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
	nums, errs := qs.Limit(offset, start).All(&mssql)
	return nums, errs, mssql
}


//统计数量
func CountMssql(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("mssql_status"))
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