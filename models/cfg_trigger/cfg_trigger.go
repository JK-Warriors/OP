package cfg_trigger

import (
	//"fmt"
	"opms/models"
	//"opms/utils"
	"time"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Trigger struct {
	Id       			int `orm:"pk;column(id);"`
	Asset_Id   			int `orm:"column(asset_id);"`
	Asset_Host   		string `orm:"column(host);"`
	Asset_Port   		string `orm:"column(port);"`
	Asset_Alias   		string `orm:"column(alias);"`
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
	o := orm.NewOrm()

	sql := `select s.*,a.host, a.port, a.alias 
			from pms_triggers s, pms_asset_config a 
			where s.asset_id = a.id 
			  and s.id = ?`

	err := o.Raw(sql, id).QueryRow(&triconf)

	if err == orm.ErrNoRows {
		return triconf, nil
	}
	return triconf, err
}


//获取告警配置列表
func ListTriggers(condArr map[string]string, page int, offset int) (num int64, err error, trigger []Trigger) {
	o := orm.NewOrm()
	o.Using("default")
	sql := `select s.*,a.host, a.port, a.alias 
			from pms_triggers s, pms_asset_config a 
			where s.asset_id = a.id`

	if condArr["search_name"] != "" {
		sql = sql + " and (s.trigger_type like '%" + condArr["search_name"] + "%')"
	}

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	
	sql = sql + " order by s.id"
	sql = sql + " limit " + strconv.Itoa(offset) + " offset " + strconv.Itoa(start)
	nums, errs := o.Raw(sql).QueryRows(&trigger)

	return nums, errs, trigger
}


//统计数量
func CountTriggers(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("triggers"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("trigger_type__icontains", condArr["search_name"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}


//修改告警配置
func UpdateTriggerConfig(id int, tri Trigger) error {
	o := orm.NewOrm()

	sql:=`update pms_triggers set asset_type = ?, 
			trigger_type = ?, 
			severity = ?, 
			expression = ?, 
			description = ?, 
			recovery_mode = ?,  
			recovery_expression = ?, 
			recovery_description = ?, 
			created = ?
		  where id = ?`
	
	_, err := o.Raw(sql, tri.Asset_Type, tri.Trigger_Type, tri.Severity, tri.Expression, tri.Description, tri.Recovery_Mode, tri.Recovery_Expression, tri.Recovery_Description, time.Now().Unix(), id).Exec()
	
	return err
}

