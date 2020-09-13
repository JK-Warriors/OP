package alarm

import (
	//"fmt"
	"opms/models"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Alert struct {
	Id       				string `orm:"pk;column(Id);"`
	Asset_Id       			string `orm:"column(asset_id);"`
	Name   					string `orm:"column(name);"`
	Severity   				string `orm:"column(severity);"`
	TemplateId   			string `orm:"column(templateid);"`
	MediaTypeId   			string `orm:"column(mediatypeid);"`
	Sendto   				string `orm:"column(sendto);"`
	Subject   				string `orm:"column(subject);"`
	Message   				string `orm:"column(message);"`
	Status   				string `orm:"column(status);"`
	Retries   				string `orm:"column(retries);"`
	Error   				string `orm:"column(error);"`
	AlertType   			string `orm:"column(alerttype);"`
	Created   				int64 `orm:"column(created);"`
}


func (this *Alert) TableName() string {
	return models.TableName("alerts")
}
func init() {
	orm.RegisterModel(new(Alert))
}


//获取告警列表
func ListAlerts(condArr map[string]string, page int, offset int) (num int64, err error, alerts []Alert) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("alerts"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("name__icontains", condArr["search_name"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	//qs = qs.OrderBy("name")
	nums, errs := qs.Limit(offset, start).All(&alerts)
	return nums, errs, alerts
}


//统计数量
func CountAlerts(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("alerts"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("name__icontains", condArr["search_name"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}

