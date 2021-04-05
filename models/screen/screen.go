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

	sql := `select count(1) from pms_asset_status where asset_type = ? and repl > -1`
	_ = o.Raw(sql, asset_type).QueryRow(&num)
	return num
}

func GetDRWarning(asset_type int) (num int) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select count(1) from pms_asset_status where asset_type = ? and repl > -1 and repl_delay > 600 and  repl_delay <= 3600`
	_ = o.Raw(sql, asset_type).QueryRow(&num)
	return num
}

func GetDRCritical(asset_type int) (num int) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select count(1) from pms_asset_status where asset_type = ? and repl = -1 and repl_delay > 3600`
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
			and c.status = 1
			and c.is_delete = 0
			and c.show_on_screen = 1
			group by alias`
	_, _ = o.Raw(sql).QueryRows(&dbtime)

	return dbtime
}

type Tbs_Rate struct {
	Tbs_Name string  `orm:"column(tbs_name);"`
	Rate     float32 `orm:"column(max_rate);"`
}

func GetTablespace() (tbs_rate []Tbs_Rate) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select CONCAT(o.alias, ': ', o.tablespace_name) tbs_name, o.max_rate 
			from pms_oracle_tablespace o, pms_asset_config c
				where o.db_id = c.id
				and c.status = 1
				and c.is_delete = 0
				and c.show_on_screen = 1
				order by max_rate desc`
	_, _ = o.Raw(sql).QueryRows(&tbs_rate)

	return tbs_rate
}

type MetricValue struct {
	Db_Id int    `orm:"column(db_id);"`
	Time  string `orm:"column(time);"`
	Value int    `orm:"column(value);"`
}

func GetMetrixValueX(metric_name string) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, timestamp as time, value
			from pms_metric_data m, (select * from pms_asset_config where status = 1 and is_delete = 0 and show_on_screen = 1 limit 1) c
			where m.metric = ?
			and m.db_id = c.id
			and m.timestamp > FROM_UNIXTIME(UNIX_TIMESTAMP() - 60*24, '%Y-%m-%d %H:%i:%S')
			order by db_id, time`
	_, _ = o.Raw(sql, metric_name).QueryRows(&metricvalue)

	return metricvalue
}

func GetMetrixValueY(metric_name string) (metricvalue []MetricValue) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select db_id, timestamp as time, value
			from pms_metric_data m, (select * from pms_asset_config where status = 1 and is_delete = 0 and show_on_screen = 1 limit 1) c
			where m.metric = ?
			and m.db_id = c.id
			and c.status = 1
			and c.is_delete = 0
			and c.show_on_screen = 1
			and m.timestamp > FROM_UNIXTIME(UNIX_TIMESTAMP() - 60*24, '%Y-%m-%d %H:%i:%S')
			order by db_id, time`
	_, _ = o.Raw(sql, metric_name).QueryRows(&metricvalue)

	return metricvalue
}
