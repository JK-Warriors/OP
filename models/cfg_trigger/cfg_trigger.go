package cfg_trigger

import (
	//"fmt"
	"opms/models"
	//"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Trigger struct {
	Id       			int `orm:"pk;column(id);"`
	Asset_Id   			int `orm:"column(asset_id);"`
	Asset_Type   		int `orm:"column(asset_type);"`
	Name   				string `orm:"column(name);"`
	TemplateId   		int `orm:"column(templateid);"`
	Trigger_Type   		string `orm:"column(trigger_type);"`
	Severity   			string `orm:"column(severity);"`
	Expression   		string `orm:"column(expression);"`
	Description   		string `orm:"column(description);"`
	Status   			int `orm:"column(status);"`
	Recovery_Mode   	int `orm:"column(recovery_mode);"`
	Recovery_Expression string `orm:"column(recovery_expression);"`
	Recovery_Description string `orm:"column(recovery_description);"`
	Created  			int64  `orm:"column(created);"`
}


func (this *Trigger) TableName() string {
	return models.TableName("triggers")
}
func init() {
	orm.RegisterModel(new(Trigger))
}


func AddAssetTriggers(asset_id int64, asset_type int) error {
	o := orm.NewOrm()
	o.Using("default")

	sql:=`insert into pms_triggers(asset_id, asset_type, name, templateid, trigger_type, severity, expression, description, status, recovery_mode, recovery_expression, recovery_description, created)
	select ?, asset_type, name, id, trigger_type, severity, expression, description, status, recovery_mode,recovery_expression, recovery_description, ?
	from pms_trigger_template where asset_type = ?`
	
	_, err := o.Raw(sql, asset_id, time.Now().Unix(), asset_type).Exec()
	
	return err
}



func GetTriggerConfig(id int) (Trigger, error) {
	var triconf Trigger
	var err error
	o := orm.NewOrm()

	triconf = Trigger{Id: id}
	err = o.Read(&triconf)

	if err == orm.ErrNoRows {
		return triconf, nil
	}
	return triconf, err
}


//获取资产状态列表
func ListTriggers(condArr map[string]string, page int, offset int) (num int64, err error, trigger []Trigger) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("triggers"))
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

	qs = qs.OrderBy("id")
	nums, errs := qs.Limit(offset, start).All(&trigger)
	return nums, errs, trigger
}


//统计数量
func CountTriggers(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("triggers"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("name__icontains", condArr["search_name"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}


//修改告警配置
func UpdateTriggerConfig(id int, tri Trigger) error {
	var triconf Trigger
	o := orm.NewOrm()
	triconf = Trigger{Id: id}

	triconf.Id = id
	triconf.Asset_Type = tri.Asset_Type
	triconf.Trigger_Type = tri.Trigger_Type
	triconf.Severity = tri.Severity
	triconf.Expression = tri.Expression
	triconf.Description = tri.Description
	
	triconf.Recovery_Mode = tri.Recovery_Mode
	triconf.Recovery_Expression = tri.Recovery_Expression
	triconf.Recovery_Description = tri.Recovery_Description

	_, err := o.Update(&triconf, "Asset_Type", "Trigger_Type", "Severity", "Expression", "Description", "Recovery_Mode", "Recovery_Expression", "Recovery_Description")
	return err
}

