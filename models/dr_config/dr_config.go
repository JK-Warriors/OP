package dr_config

import (
	//"fmt"
	"opms/models"
	"strconv"

	//"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DrConfig struct {
	Bs_Id        int    `orm:"pk;column(bs_id);"`
	Bs_Name   	 string `orm:"column(bs_name);"`
	Asset_Type   int    `orm:"column(asset_type);"`
	Db_Id_P      int    `orm:"column(db_id_p);"`
	Db_Dest_P    int    `orm:"column(db_dest_p);"`
	Db_Id_S      int    `orm:"column(db_id_s);"`
	Db_Dest_S    int    `orm:"column(db_dest_s);"`
	Db_Name	     string `orm:"column(db_name);"`
	Fb_Retention int    `orm:"column(fb_retention);"`
	Is_Shift     int    `orm:"column(is_shift);"`
	Shift_Vips   string `orm:"column(shift_vips);"`
	Network_P    string `orm:"column(network_p);"`
	Network_S    string `orm:"column(network_s);"`
	Is_Switch    int    `orm:"column(is_switch);"`
	Status    	 int    `orm:"column(status);"`
	Is_Alert       int    `orm:"column(is_alert);"`
	Alert_Mail     int    `orm:"column(alert_mail);"`
	Alert_WeChat   int    `orm:"column(alert_wechat);"`
	Alert_SMS      int    `orm:"column(alert_sms);"`
	Is_Delete    int    `orm:"column(is_delete);"`
	Created      int64  `orm:"column(created);"`
	Updated      int64  `orm:"column(updated);"`
}

func (this *DrConfig) TableName() string {
	return models.TableName("dr_config")
}
func init() {
	orm.RegisterModel(new(DrConfig))
}

//添加容灾配置
func AddDrConfig(dc DrConfig) error {
	o := orm.NewOrm()
	o.Using("default")
	drconf := new(DrConfig)

	drconf.Bs_Name = dc.Bs_Name
	drconf.Asset_Type = dc.Asset_Type
	drconf.Db_Id_P = dc.Db_Id_P
	drconf.Db_Dest_P = dc.Db_Dest_P
	drconf.Db_Id_S = dc.Db_Id_S
	drconf.Db_Dest_S = dc.Db_Dest_S
	drconf.Db_Name = dc.Db_Name
	drconf.Fb_Retention = dc.Fb_Retention
	drconf.Is_Shift = dc.Is_Shift
	drconf.Shift_Vips = dc.Shift_Vips
	drconf.Network_P = dc.Network_P
	drconf.Network_S = dc.Network_S
	drconf.Is_Alert = dc.Is_Alert
	drconf.Alert_Mail = dc.Alert_Mail
	drconf.Alert_WeChat = dc.Alert_WeChat
	drconf.Alert_SMS = dc.Alert_SMS
	drconf.Status = 1

	drconf.Created = time.Now().Unix()
	_, err := o.Insert(drconf)
	return err
}

//修改容灾配置
func UpdateDrConfig(id int, dc DrConfig) error {
	var drconf DrConfig
	o := orm.NewOrm()
	drconf = DrConfig{Bs_Id: id}

	drconf.Bs_Id = id
	drconf.Bs_Name = dc.Bs_Name
	drconf.Asset_Type = dc.Asset_Type
	drconf.Db_Id_P = dc.Db_Id_P
	drconf.Db_Dest_P = dc.Db_Dest_P
	drconf.Db_Id_S = dc.Db_Id_S
	drconf.Db_Dest_S = dc.Db_Dest_S
	drconf.Db_Name = dc.Db_Name
	drconf.Fb_Retention = dc.Fb_Retention
	drconf.Is_Shift = dc.Is_Shift
	drconf.Shift_Vips = dc.Shift_Vips
	drconf.Network_P = dc.Network_P
	drconf.Network_S = dc.Network_S
	drconf.Is_Alert = dc.Is_Alert
	drconf.Alert_Mail = dc.Alert_Mail
	drconf.Alert_WeChat = dc.Alert_WeChat
	drconf.Alert_SMS = dc.Alert_SMS
	drconf.Updated = time.Now().Unix()

	_, err := o.Update(&drconf)
	return err
	

}

//得到容灾信息
func GetDrConfig(id int) (DrConfig, error) {
	var drconf DrConfig
	var err error
	o := orm.NewOrm()

	drconf = DrConfig{Bs_Id: id}
	err = o.Read(&drconf)

	if err == orm.ErrNoRows {
		return drconf, nil
	}
	return drconf, err
}

//获取容灾列表
func ListDrConfig(condArr map[string]string, page int, offset int) (num int64, err error, drconf []DrConfig) {
	o := orm.NewOrm()
	o.Using("default")
	sql := `select bs_id, bs_name, asset_type, db_id_p, db_dest_p, db_id_s, db_dest_s, db_name, fb_retention, is_shift, shift_vips, network_p, network_s, is_switch, status, is_delete
			from pms_dr_config where 1=1`

	if condArr["host"] != "" {
		sql = sql + " and (d.db_id_p like '%" + condArr["host"] + "%' or d.db_id_s like '%" + condArr["host"] + "%')"
	}

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	sql = sql + " order by bs_id"
	sql = sql + " limit " + strconv.Itoa(offset) + " offset " + strconv.Itoa(start)
	nums, errs := o.Raw(sql).QueryRows(&drconf)
	return nums, errs, drconf
}

//统计容灾数量
func CountDrConfig(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("dr_config"))
	cond := orm.NewCondition()

	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteDrConfig(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("dr_config") + " WHERE bs_id IN(" + ids + ")").Exec()
	return err
}

//更改资产状态
func ChangeDrconfigStatus(id int, status int) error {
	o := orm.NewOrm()

	drconf := DrConfig{Bs_Id: id}
	err := o.Read(&drconf, "bs_id")
	if nil != err {
		return err
	} else {
		drconf.Status = status
		_, err := o.Update(&drconf)
		return err
	}
}
