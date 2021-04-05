package asset

import (
	//"fmt"
	"opms/models"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Asset struct {
	Id       			int `orm:"pk;column(id);"`
	Asset_Type   		int `orm:"column(asset_type);"`
	Host   				string `orm:"column(host);"`
	Port   				string `orm:"column(port);"`
	Alias   			string `orm:"column(alias);"`
	Connect   			int `orm:"column(connect);"`
	Db_Name   			string `orm:"column(db_name);"`
	Inst_Name   		string `orm:"column(inst_name);"`
	Inst_Role   		string `orm:"column(inst_role);"`
	Inst_Status   		string `orm:"column(inst_status);"`
	Db_Role   			string `orm:"column(db_role);"`
	Open_Mode   		string `orm:"column(open_mode);"`
	Protection_Mode   	string `orm:"column(protection_mode);"`
	Host_Name   		string `orm:"column(host_name);"`
	Startup_Time   		string `orm:"column(startup_time);"`
	Uptime   			string `orm:"column(uptime);"`
	Version   			string `orm:"column(version);"`
	Archiver 			string `orm:"column(archiver);"`
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

// func (this *Asset) TableName() string {
// 	return models.TableName("asset_status")
// }
// func init() {
// 	orm.RegisterModel(new(Asset))
// }


//获取资产状态列表
func ListAssetStatus(condArr map[string]string, page int, offset int) (num int64, err error, asset []Asset) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_status"))
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
	nums, errs := qs.Limit(offset, start).All(&asset)
	return nums, errs, asset
}


//统计数量
func CountAsset(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("asset_status"))
	cond := orm.NewCondition()

	if condArr["bs_name"] != "" {
		cond = cond.And("bs_name__icontains", condArr["bs_name"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}


func ListAllDBStatus() (num int64, err error, asset []Asset) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select s.* from pms_asset_status s, pms_asset_config c where s.asset_id = c.id and s.asset_type < 99 order by c.display_order`
	nums, errs := o.Raw(sql).QueryRows(&asset)

	return nums, errs, asset
}
