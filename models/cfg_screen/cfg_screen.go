package cfg_screen

import (
	//"fmt"
	"opms/models"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Dbconfig struct {
	Id           int    `orm:"pk;column(id);"`
	Dbtype       int    `orm:"column(asset_type);"`
	Host         string `orm:"column(host);"`
	Protocol     string `orm:"column(protocol);"`
	Port         int    `orm:"column(port);"`
	Alias        string `orm:"column(alias);"`
	InstanceName string `orm:"column(instance_name);"`
	Dbname       string `orm:"column(db_name);"`
	Username     string `orm:"column(username);"`
	Password     string `orm:"column(password);"`
	Role         int    `orm:"column(role);"`
	Ostype       int    `orm:"column(os_type);"`
	OsProtocol   string `orm:"column(os_protocol);"`
	OsPort       string `orm:"column(os_port);"`
	OsUsername   string `orm:"column(os_username);"`
	OsPassword   string `orm:"column(os_password);"`
	Status       int    `orm:"column(status);"`
	IsDelete     int    `orm:"column(is_delete);"`
	Retention    int    `orm:"column(retention);"`
	Alert_Mail   int    `orm:"column(alert_mail);"`
	Alert_WeChat int    `orm:"column(alert_wechat);"`
	Alert_SMS    int    `orm:"column(alert_sms);"`
	Created      int64  `orm:"column(created);"`
	Updated      int64  `orm:"column(updated);"`
}

func ListDB(condArr map[string]string, page int, offset int) (num int64, err error, dbconfig []Dbconfig) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("bs_name__icontains", condArr["search_name"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("id")
	nums, errs := qs.Limit(offset, start).All(&dbconfig)
	return nums, errs, dbconfig
}

//统计数量
func CountDB(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	if condArr["bs_name"] != "" {
		cond = cond.And("bs_name__icontains", condArr["bs_name"])
	}

	num, _ := qs.SetCond(cond).Count()
	return num
}
