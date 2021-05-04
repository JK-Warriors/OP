package cfg_screen

import (
	//"fmt"
	"opms/models"
	"opms/utils"
	dbconfig "opms/models/dbconfig"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func ListDB(condArr map[string]string, page int, offset int) (num int64, err error, dbconfig []dbconfig.Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
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

func ListOracle(condArr map[string]string, page int, offset int) (num int64, err error, dbconfig []dbconfig.Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}

	if condArr["asset_type"] != "" {
		cond = cond.And("asset_type__exact", 1)
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

	if condArr["db_name"] != "" {
		cond = cond.And("db_name__icontains", condArr["db_name"])
	}

	num, _ := qs.SetCond(cond).Count()
	return num
}

func GetCoreDb() int {
	o := orm.NewOrm()
	var core_db int

	err := o.Raw("select CASE WHEN LENGTH(trim(value)) > 0 THEN value else -1 END from pms_global_options where id='core_db' ").QueryRow(&core_db)
	if err != nil {
		utils.LogDebug("GetCoreDb failed: " + err.Error())
		core_db = -1
	}
	return core_db
}

func SaveShowOnScreen(ids string) error {
	o := orm.NewOrm()

	_, err := o.Raw("update pms_asset_config set show_on_screen = 1 WHERE id IN(" + ids + ")").Exec()
	return err
}

func RemoveAllShowOnScreen(ids string) error {
	o := orm.NewOrm()

	_, err := o.Raw("update pms_asset_config set show_on_screen = 0 WHERE show_on_screen = 1").Exec()
	return err
}

func SaveCoreDbOnScreen(id string) error {
	o := orm.NewOrm()

	_, err := o.Raw("update pms_global_options set value = " + id + " WHERE id = 'core_db' ").Exec()
	return err
}