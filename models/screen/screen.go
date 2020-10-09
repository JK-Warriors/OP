package screen

import (
	//"fmt"
	//"opms/models"
	//"opms/utils"

	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func GetLastCheckTime() (lastchecktime string) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select from_unixtime(max(created), '%Y年%m月%d日 %H:%i:%S') from pms_asset_status`
	_ = o.Raw(sql).QueryRow(&lastchecktime)
	return lastchecktime
}

type DBScore struct {
    Asset_Id 		int		    	`orm:"column(asset_id);"`
    Alias 			string			`orm:"column(alias);"`
    Score 			int				`orm:"column(score);"`
}
func GetDBScore() (num int64, err error, dbscore []DBScore) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select asset_id, alias, 50 as score from pms_asset_status where connect > 0 limit 4`
	nums, errs := o.Raw(sql).QueryRows(&dbscore)

	return nums, errs, dbscore
}

func GetDRNormal(asset_type int) (num int) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select count(1) from pms_asset_status where asset_type = ? and repl = -1`
	_ = o.Raw(sql, asset_type).QueryRow(&num)
	return num
}

func GetDRWarning(asset_type int) (num int) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select count(1) from pms_asset_status where asset_type = ? and repl = -1`
	_ = o.Raw(sql, asset_type).QueryRow(&num)
	return num
}

func GetDRCritical(asset_type int) (num int) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select count(1) from pms_asset_status where asset_type = ? and repl = -1`
	_ = o.Raw(sql, asset_type).QueryRow(&num)
	return num
}