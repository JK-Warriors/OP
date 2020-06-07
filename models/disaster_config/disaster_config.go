package disaster_config

import (
	//"fmt"
	"opms/models"
	"strconv"

	//"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DisasterConfig struct {
	Bs_Id        int    `orm:"pk;column(bs_id);"`
	Db_Id_P      int    `orm:"column(db_id_p);"`
	Db_Dest_P    int    `orm:"column(db_dest_p);"`
	Db_Id_S      int    `orm:"column(db_id_s);"`
	Db_Dest_S    int    `orm:"column(db_dest_s);"`
	Fb_Retention int    `orm:"column(fb_retention);"`
	Is_Shift     int    `orm:"column(is_shift);"`
	Shift_Vips   string `orm:"column(shift_vips);"`
	Network_P    string `orm:"column(network_p);"`
	Network_S    string `orm:"column(network_s);"`
	Created      int64  `orm:"column(created);"`
	Updated      int64  `orm:"column(updated);"`
}

func (this *DisasterConfig) TableName() string {
	return models.TableName("disaster_config")
}
func init() {
	orm.RegisterModel(new(DisasterConfig))
}

//添加容灾配置
func AddDisasterConfig(dc DisasterConfig) error {
	o := orm.NewOrm()
	o.Using("default")
	disasterconf := new(DisasterConfig)

	disasterconf.Bs_Id = dc.Bs_Id
	disasterconf.Db_Id_P = dc.Db_Id_P
	disasterconf.Db_Dest_P = dc.Db_Dest_P
	disasterconf.Db_Id_S = dc.Db_Id_S
	disasterconf.Db_Dest_S = dc.Db_Dest_S
	disasterconf.Fb_Retention = dc.Fb_Retention
	disasterconf.Is_Shift = dc.Is_Shift
	disasterconf.Shift_Vips = dc.Shift_Vips
	disasterconf.Network_P = dc.Network_P
	disasterconf.Network_S = dc.Network_S
	disasterconf.Created = time.Now().Unix()
	_, err := o.Insert(disasterconf)
	return err
}

//修改容灾配置
func UpdateDisasterConfig(id int, dc DisasterConfig) error {
	var disasterconf DisasterConfig
	o := orm.NewOrm()
	disasterconf = DisasterConfig{Bs_Id: id}

	disasterconf.Bs_Id = id
	disasterconf.Db_Id_P = dc.Db_Id_P
	disasterconf.Db_Dest_P = dc.Db_Dest_P
	disasterconf.Db_Id_S = dc.Db_Id_S
	disasterconf.Db_Dest_S = dc.Db_Dest_S
	disasterconf.Fb_Retention = dc.Fb_Retention
	disasterconf.Is_Shift = dc.Is_Shift
	disasterconf.Shift_Vips = dc.Shift_Vips
	disasterconf.Network_P = dc.Network_P
	disasterconf.Network_S = dc.Network_S
	disasterconf.Updated = time.Now().Unix()

	_, err := o.Update(&disasterconf)
	return err
}

//得到容灾信息
func GetDisasterConfig(id int) (DisasterConfig, error) {
	var disasterconf DisasterConfig
	var err error
	o := orm.NewOrm()

	disasterconf = DisasterConfig{Bs_Id: id}
	err = o.Read(&disasterconf)

	if err == orm.ErrNoRows {
		return disasterconf, nil
	}
	return disasterconf, err
}

//获取容灾列表
func ListDisasterConfig(condArr map[string]string, page int, offset int) (num int64, err error, disasterconf []DisasterConfig) {
	o := orm.NewOrm()
	o.Using("default")
	sql := `select b.id as bs_id, d.db_id_p, d.db_dest_p, d.db_id_s, d.db_dest_s, d.fb_retention, d.is_shift, d.shift_vips, d.network_p, d.network_s
			from pms_business b LEFT JOIN pms_disaster_config d on d.bs_id = b.id where 1=1`

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
	nums, errs := o.Raw(sql).QueryRows(&disasterconf)
	return nums, errs, disasterconf
}

//统计容灾数量
func CountDisasterConfig(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("disaster_config"))
	cond := orm.NewCondition()

	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteDisasterConfig(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("disaster_config") + " WHERE bs_id IN(" + ids + ")").Exec()
	return err
}
