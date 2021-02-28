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
	Asset_Id int    `orm:"column(asset_id);"`
	Alias    string `orm:"column(alias);"`
	Score    int    `orm:"column(score);"`
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

type DbTime struct {
	Alias string `orm:"column(alias);"`
	Value int    `orm:"column(value);"`
}

func GetDBTime() (dbtime []DbTime) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select alias, sum(db_time) as value
			from pms_oracle_db_time o, pms_asset_config c
			where o.db_id = c.id
			and db_id in(101, 102)
			group by alias`
	_, _ = o.Raw(sql).QueryRows(&dbtime)

	return dbtime
}

type MetricValue struct {
	Db_Id int    `orm:"column(db_id);"`
	Time  string `orm:"column(time);"`
	Value int    `orm:"column(value);"`
}

func GetActiveSessionX(db_id int) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, from_unixtime(created) as time, session_actives as value
			from pms_oracle_status_his
			where db_id in(101)
			and created > UNIX_TIMESTAMP() - 3600*24*100
			order by db_id, created`
	_, _ = o.Raw(sql).QueryRows(&metricvalue)

	return metricvalue
}

func GetActiveSessionY(db_id int) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, from_unixtime(created) as time, session_actives as value
			from pms_oracle_status_his
			where db_id in(101,102)
			and created > UNIX_TIMESTAMP() - 3600*24*100
			order by db_id, created`
	_, _ = o.Raw(sql).QueryRows(&metricvalue)

	return metricvalue
}

func GetTotalSessionX(db_id int) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, from_unixtime(created) as time, session_total as value
			from pms_oracle_status_his
			where db_id in(101)
			and created > UNIX_TIMESTAMP() - 3600*24*100
			order by db_id, created`
	_, _ = o.Raw(sql).QueryRows(&metricvalue)

	return metricvalue
}

func GetTotalSessionY(db_id int) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, from_unixtime(created) as time, session_total as value
			from pms_oracle_status_his
			where db_id in(101,102)
			and created > UNIX_TIMESTAMP() - 3600*24*100
			order by db_id, created`
	_, _ = o.Raw(sql).QueryRows(&metricvalue)

	return metricvalue
}

func GetRedoX(db_id int) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, key_time as time, redo_log as value
			from pms_oracle_redo
			where db_id in(101)
			and created > UNIX_TIMESTAMP() - 3600*24*100
			order by db_id, created`
	_, _ = o.Raw(sql).QueryRows(&metricvalue)

	return metricvalue
}

func GetRedoY(db_id int) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, key_time as time, redo_log as value
			from pms_oracle_redo
			where db_id in(101,102)
			and created > UNIX_TIMESTAMP() - 3600*24*100
			order by db_id, created`
	_, _ = o.Raw(sql).QueryRows(&metricvalue)

	return metricvalue
}

func GetMetrixValueX(db_id int, metric_name string) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, key_time as time, value
			from pms_item_data
			where metric_name = ?
			and db_id in(101)
			and ns > UNIX_TIMESTAMP() - 3600*24*7
			order by db_id, time`
	_, _ = o.Raw(sql, metric_name).QueryRows(&metricvalue)

	return metricvalue
}

func GetMetrixValueY(db_id int, metric_name string) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, key_time as time, value
			from pms_item_data
			where metric_name = ?
			and db_id in(101,102)
			and ns > UNIX_TIMESTAMP() - 3600*24*7
			order by db_id, time`
	_, _ = o.Raw(sql, metric_name).QueryRows(&metricvalue)

	return metricvalue
}
