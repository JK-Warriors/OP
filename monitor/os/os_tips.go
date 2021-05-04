package os

import (
	//"database/sql"
	"log"
	"strings"
	"time"
	"fmt"
	//"context"
	//"opms/monitor/utils"
	"opms/monitor/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/Knetic/govaluate"
)



func Update_Item_Tips(mysql *xorm.Engine, host string, field string, value string, trigger_type string, level int, severity string){
	var value_tips string
	curr_time := time.Now().Format("2006-01-02 15:04:05")

	field_tips := field + "_tips"
	if value == "-1"{
		value_tips = "no data"
	}else{
		value_tips = fmt.Sprintf("Value: %s\nLevel: %s\nTime: %s", value, severity, curr_time)

	}

	sql := fmt.Sprintf("update pms_asset_status set `%s`=%d, `%s`='%s' where host = ?", field, level, field_tips, value_tips)
	
	_, err := mysql.Exec(sql, host)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}
}

func Check_Item_Status(mysql *xorm.Engine, os_id int, host string, trigger_type string, item_name string, item_value string){
	if item_value == "-1"{
		Update_Item_Tips(mysql, host, item_name, item_value, trigger_type, -1, "--")

	}else{
		tri_critical, err := common.GetTrigger(mysql, os_id, trigger_type, "Critical")
		if err == nil {
			result := CheckExpression(tri_critical, item_value)
			if result == 1{
				Update_Item_Tips(mysql, host, item_name, item_value, trigger_type, 3, "Critical")
				return 
			}
		}
	
		tri_warning, err := common.GetTrigger(mysql, os_id, trigger_type, "Warning")
		if err == nil {
			result := CheckExpression(tri_warning, item_value)
			if result == 1{
				Update_Item_Tips(mysql, host, item_name, item_value, trigger_type, 2, "Warning")
				return
			}
		}
	
		Update_Item_Tips(mysql, host, item_name, item_value, trigger_type, 1, "OK")
	}

}


func CheckExpression(tri common.Trigger, item_value string) int{
	if tri.Status == 1 {
		exp := strings.Replace(tri.Expression, "{ItemValue}", item_value, -1)
		
		expression, err := govaluate.NewEvaluableExpression(exp)
		if err != nil {
			log.Printf("govaluate: %s", err.Error())
			return 0
		}

		result, err := expression.Evaluate(nil)
		if err != nil {
			log.Printf("Expression error: %s", err.Error())
			return 0
		}
		
		log.Printf("Expression result: %v", result)
		if result == true {
			return 1
		}else{
			return 0
		}
	}else{
		return 0
	}

}

func GatherOSStatus(mysql *xorm.Engine, os_id int, host string){
	log.Printf("Update tips in asset_status: %s", host)

	load_1 := GetColumnValue(mysql, "load_1", os_id)
	Check_Item_Status(mysql, os_id, host, "load", "load", load_1)

	cpu_used := GetCpuUsage(mysql, os_id)
	Check_Item_Status(mysql, os_id, host, "cpu", "cpu", cpu_used)

	mem_usage_rate := GetColumnValue(mysql, "mem_usage_rate", os_id)
	Check_Item_Status(mysql, os_id, host, "memory", "memory", mem_usage_rate)

	io_total := GetIOTotal(mysql, os_id)
	Check_Item_Status(mysql, os_id, host, "io", "io", io_total)

	net_total := GetNetTotal(mysql, os_id)
	Check_Item_Status(mysql, os_id, host, "net", "net", net_total)
}


func GetColumnValue(mysql *xorm.Engine, column string, db_id int) string{
	var value string = "-1"

	sql := `select ` + column + ` from pms_os_status where os_id = ?`

	_, err := mysql.SQL(sql, db_id).Get(&value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return value
}

func GetCpuUsage(mysql *xorm.Engine, os_id int) string{
	var value string = "-1"

	sql := `select cpu_user_time + cpu_system_time from pms_os_status where os_id = ?`

	_, err := mysql.SQL(sql, os_id).Get(&value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return value
}


func GetIOTotal(mysql *xorm.Engine, os_id int) string{
	var value string = "-1"

	sql := `select disk_io_reads_total + disk_io_writes_total from pms_os_status where os_id = ?`

	_, err := mysql.SQL(sql, os_id).Get(&value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return value
}


func GetNetTotal(mysql *xorm.Engine, os_id int) string{
	var value string = "-1"

	sql := `select net_in_bytes_total + net_out_bytes_total from pms_os_status where os_id = ?`

	_, err := mysql.SQL(sql, os_id).Get(&value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return value
}