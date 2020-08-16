package mysql

import (
	//"fmt"
	"opms/models"
	"strconv"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Mysql struct {
	Id       			int `orm:"pk;column(id);"`
	Db_Id   			int `orm:"column(db_id);"`
	Host   				string `orm:"column(host);"`
	Port   				string `orm:"column(port);"`
	Alias   			string `orm:"column(alias);"`
	Connect   			int `orm:"column(connect);"`
	Role   				string `orm:"column(role);"`
	Version   			string `orm:"column(version);"`
	Uptime   			string `orm:"column(uptime);"`
	Max_Connections   			string `orm:"column(max_connections);"`
	Max_Connect_Errors   		string `orm:"column(max_connect_errors);"`
	Open_Files_Limit   			string `orm:"column(open_files_limit);"`
	Open_Files   				string `orm:"column(open_files);"`
	Table_Open_Cache   			string `orm:"column(table_open_cache);"`
	Open_Tables   				string `orm:"column(open_tables);"`
	Threads_Connected   		string `orm:"column(threads_connected);"`
	Threads_Running 			string `orm:"column(threads_running);"`
	Threads_Waits   			string `orm:"column(threads_waits);"`
	Key_Buffer_Size   			string `orm:"column(key_buffer_size);"`
	Sort_Buffer_Size 			string `orm:"column(sort_buffer_size);"`
	Join_Buffer_Size   			string `orm:"column(join_buffer_size);"`
	Key_Blocks_Unused   		string `orm:"column(key_blocks_unused);"`
	Key_Blocks_Used 			string `orm:"column(key_blocks_used);"`
	Key_Blocks_Not_Flushed   	string `orm:"column(key_blocks_not_flushed);"`
	Key_Blocks_Used_Rate   		string `orm:"column(key_blocks_used_rate);"`
	Key_Buffer_Read_Rate 		string `orm:"column(key_buffer_read_rate);"`
	Key_Buffer_Write_Rate   	string `orm:"column(key_buffer_write_rate);"`
	
	Created  					int64  `orm:"column(created);"`
}

func (this *Mysql) TableName() string {
	return models.TableName("mysql_status")
}



func init() {
	orm.RegisterModel(new(Mysql))
}


//获取Mysql状态列表
func ListMysqlStatus(condArr map[string]string, page int, offset int) (num int64, err error, mysql []Mysql) {
	o := orm.NewOrm()
	o.Using("default")
	sql := `select id, db_id, host, port, alias, connect, role, version, max_connections, max_connect_errors, open_files_limit, open_files, table_open_cache, open_tables, uptime, threads_connected, threads_running, threads_waits, 
				key_buffer_size, sort_buffer_size, join_buffer_size, key_blocks_unused, key_blocks_used, key_blocks_not_flushed, key_blocks_used_rate, key_buffer_read_rate, key_buffer_write_rate, created
				from pms_mysql_status where 1=1`

	if condArr["host"] != "" {
		sql = sql + " and (host like '%" + condArr["host"] + "%')"
	}
	if condArr["alias"] != "" {
		sql = sql + " and (alias like '%" + condArr["alias"] + "%')"
	}

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	
	sql = sql + " order by db_id"
	sql = sql + " limit " + strconv.Itoa(offset) + " offset " + strconv.Itoa(start)
	nums, errs := o.Raw(sql).QueryRows(&mysql)
	return nums, errs, mysql
}


//统计数量
func CountMysql(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("mysql_status"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}


