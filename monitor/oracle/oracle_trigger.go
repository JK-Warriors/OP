package oracle

import (
	//"database/sql"
	"log"
	"strings"
	//"time"
	//"context"
	"opms/monitor/utils"
	"opms/monitor/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/Knetic/govaluate"
)


func ExecTriggers(mysql *xorm.Engine, db_id int, trigger_type string, item_name string, item_value string){

	triconf := common.GetTriggers(mysql, db_id, trigger_type)
	for _, tri := range triconf {
		utils.LogDebugf("Alert status: %s", tri.Status)
		if tri.Status == 1 {
			exp := strings.Replace(tri.Expression, "{ItemValue}", item_value, -1)
			
			expression, err := govaluate.NewEvaluableExpression(exp)
			if err != nil {
				utils.LogDebugf("govaluate: %s", err.Error())
				return
			}
	
			//parameters := make(map[string]interface{}, 8)
			//parameters["{ItemValue}"] = itemvalue;
	
			result, err := expression.Evaluate(nil)
			if err != nil {
				utils.LogDebugf("Expression error: %s", err.Error())
				return
			}
			
			utils.LogDebugf("Expression result: %v", result)
			if result == true {
				utils.LogDebugf("Add %s alert for db %d", trigger_type, db_id)
				common.AddAlert(mysql, db_id, item_name, item_value, tri)
			}else{
				// recover
				log.Printf("Recovery_Mode: %v", tri.Recovery_Mode)
				if tri.Recovery_Mode == 1 {
					exp = strings.Replace(tri.Recovery_Expression, "{ItemValue}", item_value, -1)
			
					expression, err = govaluate.NewEvaluableExpression(exp)
					if err != nil {
						utils.LogDebugf("govaluate: %s", err.Error())
						return
					}
			
					result, err := expression.Evaluate(nil)
					if err != nil {
						utils.LogDebugf("ExecTriggers: %s", err.Error())
						return
					}

					if result == true {
						common.AddRecoveryAlert(mysql, db_id, item_name, item_value, tri)
					}
				}else{
					utils.LogDebugf("There is no recovery mode for this trigger")
				}
			}


		}else{
			utils.LogDebugf("Alert status is disable")
		}
	}
	
}


func AlertBasicInfo(mysql *xorm.Engine, db_id int){
	connect := GetConnect(mysql, db_id)
	utils.LogDebugf("AlertConnect: %s", connect)
	ExecTriggers(mysql, db_id, "connect", "", connect)

	restart := GetRestart(mysql, db_id)
	utils.LogDebugf("AlertRestart: %s", restart)
	ExecTriggers(mysql, db_id, "restart", "", restart)

	mrpstatus := GetMrpStatus(mysql, db_id)
	utils.LogDebugf("Alert mrp status: %s", mrpstatus)
	ExecTriggers(mysql, db_id, "mrp_status", "", mrpstatus)

	dgdelay := GetDgDelay(mysql, db_id)
	utils.LogDebugf("Alert dg delay: %s", dgdelay)
	ExecTriggers(mysql, db_id, "repli_delay", "", dgdelay)
}

func AlertConnect(mysql *xorm.Engine, db_id int){
	connect := GetConnect(mysql, db_id)
	utils.LogDebugf("AlertConnect: %s", connect)
	ExecTriggers(mysql, db_id, "connect", "", connect)
}

type Tablespace struct{
	Name   				string `xorm:"varchar(200) 'tablespace_name' "`
	Size   				string `xorm:"varchar(200) 'used_size' "`
	Rate   				string `xorm:"varchar(200) 'max_rate' "`
}
func AlertTablespaces(mysql *xorm.Engine, db_id int){
	var tablespace []Tablespace
	sql := `select tablespace_name, used_size, max_rate from pms_oracle_tablespace where db_id = ?`

	err := mysql.SQL(sql, db_id).Find(&tablespace)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	for _, tbs := range tablespace{
		log.Printf("Tablespace name: %s", tbs.Name)
		log.Printf("Tablespace rate: %s", tbs.Rate)
		ExecTriggers(mysql, db_id, "tablespace", tbs.Name, tbs.Rate)
	}
}

type Diskgroup struct{
	Name   				string `xorm:"varchar(200) 'diskgroup_name' "`
	Size   				string `xorm:"varchar(200) 'used_mb' "`
	Rate   				string `xorm:"varchar(200) 'used_rate' "`
}
func AlertDiskgroups(mysql *xorm.Engine, db_id int){
	var diskgroup []Diskgroup
	sql := `select diskgroup_name, (total_mb - free_mb) as used_mb, used_rate from pms_oracle_diskgroup where db_id = ?`

	err := mysql.SQL(sql, db_id).Find(&diskgroup)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	for _, dg := range diskgroup{
		log.Printf("Asm diskgroup name: %s", dg.Name)
		log.Printf("Asm diskgroup rate: %s", dg.Rate)
		ExecTriggers(mysql, db_id, "asm_diskgroup", dg.Name, dg.Rate)
	}
}

func GetConnect(mysql *xorm.Engine, db_id int) string{
	var connect string = "-1"

	sql := `select connect from pms_oracle_status where db_id = ?`

	_, err := mysql.SQL(sql, db_id).Get(&connect)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return connect
}

func GetRestart(mysql *xorm.Engine, db_id int) string{
	var restart string = "-1"

	var last_startup string
	sql := `select startup_time from pms_oracle_status where db_id = ?`
	_, err := mysql.SQL(sql, db_id).Get(&last_startup)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	var second_startup string
	sql = `select startup_time from pms_oracle_status_his where db_id = ? order by id desc limit 1 `
	_, err = mysql.SQL(sql, db_id).Get(&second_startup)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	if last_startup == second_startup {
		restart = "1"
	}

	return restart
}

func GetMrpStatus(mysql *xorm.Engine, db_id int) string{
	var mrp_status string = "-1"

	sql := `select mrp_status from pms_dr_sta_status where db_id = ?`

	_, err := mysql.SQL(sql, db_id).Get(&mrp_status)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return mrp_status
}


func GetDgDelay(mysql *xorm.Engine, db_id int) string{
	var dg_delay string = "-1"

	sql := `select delay_mins from pms_dr_sta_status where db_id = ?`

	_, err := mysql.SQL(sql, db_id).Get(&dg_delay)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}
	log.Printf("dg_delay: %s", dg_delay)

	return dg_delay
}

