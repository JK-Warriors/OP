package mysql

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



func Update_Item_Tips(mysql *xorm.Engine, db_id int, field string, value string, trigger_type string, level int, severity string){
	var value_tips string
	curr_time := time.Now().Format("2006-01-02 15:04:05")

	field_tips := field + "_tips"
	if value == "-1"{
		value_tips = "no data"
	}else{
		value_tips = fmt.Sprintf("Value: %s\nLevel: %s\nTime: %s", value, severity, curr_time)

	}

	sql := fmt.Sprintf(`update pms_asset_status set %s=%d, %s='%s' where asset_id = ?`, field, level, field_tips, value_tips)
	
	_, err := mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}
}

func Check_Item_Status(mysql *xorm.Engine, db_id int, trigger_type string, item_name string, item_value string){
	if item_value == "-1"{
		Update_Item_Tips(mysql, db_id, item_name, item_value, trigger_type, -1, "--")

	}else{
		tri_critical, err := common.GetTrigger(mysql, db_id, trigger_type, "Critical")
		if err == nil {
			result := CheckExpression(tri_critical, item_value)
			if result == 1{
				Update_Item_Tips(mysql, db_id, item_name, item_value, trigger_type, 3, "Critical")
				return 
			}
		}
	
		tri_warning, err := common.GetTrigger(mysql, db_id, trigger_type, "Warning")
		if err == nil {
			result := CheckExpression(tri_warning, item_value)
			if result == 1{
				Update_Item_Tips(mysql, db_id, item_name, item_value, trigger_type, 2, "Warning")
				return
			}
		}
	
		Update_Item_Tips(mysql, db_id, item_name, item_value, trigger_type, 1, "OK")
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

func GatherDbStatus(mysql *xorm.Engine, db_id int){
	log.Printf("Update tips in asset_status: %d", db_id)
	
	connect := GetColumnValue(mysql, "connect", db_id)
	Check_Item_Status(mysql, db_id, "connect", "connect", connect)
	// need add more

	session_total := GetColumnValue(mysql, "threads_connected", db_id)
	Check_Item_Status(mysql, db_id, "session_total", "sessions", session_total)

	session_actives := GetColumnValue(mysql, "threads_running", db_id)
	Check_Item_Status(mysql, db_id, "session_actives", "actives", session_actives)

	session_waits := GetColumnValue(mysql, "threads_waits", db_id)
	Check_Item_Status(mysql, db_id, "session_waits", "waits", session_waits)

	processes := GetColumnValue(mysql, "threads_connected", db_id)
	Check_Item_Status(mysql, db_id, "processes", "process", processes)

	db_role := GetColumnValue(mysql, "role", db_id)
	repl_status := "-1"
	repl_delay := "-1"
	if db_role == "SLAVE"{
		repl_status = GetReplicationStatus(mysql, db_id)
		repl_delay = GetReplicationDelay(mysql, db_id)
	}
	Check_Item_Status(mysql, db_id, "repl", "repl", repl_status)
	Check_Item_Status(mysql, db_id, "repl_delay", "repl_delay", repl_delay)
	
	Check_Item_Status(mysql, db_id, "tablespace", "tablespace", "-1")
	
}


func GetColumnValue(mysql *xorm.Engine, column string, db_id int) string{
	var value string = "-1"

	sql := `select ` + column + ` from pms_mysql_status where db_id = ?`

	_, err := mysql.SQL(sql, db_id).Get(&value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return value
}
