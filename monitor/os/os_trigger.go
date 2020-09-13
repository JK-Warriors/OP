package os

import (
	//"database/sql"
	"log"
	"strings"
	//"time"
	//"context"
	//"opms/monitor/utils"
	"opms/monitor/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/Knetic/govaluate"
)


func ExecTriggers(mysql *xorm.Engine, db_id int, trigger_type string, item_name string, item_value string){

	triconf := common.GetTriggers(mysql, db_id, trigger_type)
	for _, tri := range triconf {
		log.Printf("Alert status: %s", tri.Status)
		if tri.Status == 1 {
			exp := strings.Replace(tri.Expression, "{ItemValue}", item_value, -1)
			
			expression, err := govaluate.NewEvaluableExpression(exp)
			if err != nil {
				log.Printf("govaluate: %s", err.Error())
				return
			}
	
			//parameters := make(map[string]interface{}, 8)
			//parameters["{ItemValue}"] = itemvalue;
	
			result, err := expression.Evaluate(nil)
			if err != nil {
				log.Printf("Expression error: %s", err.Error())
				return
			}
			
			log.Printf("Expression result: %v", result)
			log.Printf("TemplateId: %s", tri.TemplateId)
			if result == true {
				common.AddAlert(mysql, db_id, item_name, item_value, tri)
			}else{
				// recover
				log.Printf("Recovery_Mode: %v", tri.Recovery_Mode)
				if tri.Recovery_Mode == 1 {
					exp = strings.Replace(tri.Recovery_Expression, "{ItemValue}", item_value, -1)
			
					expression, err = govaluate.NewEvaluableExpression(exp)
					if err != nil {
						log.Printf("govaluate: %s", err.Error())
						return
					}
			
					result, err := expression.Evaluate(nil)
					if err != nil {
						log.Printf("ExecTriggers: %s", err.Error())
						return
					}

					if result == true {
						common.AddRecoveryAlert(mysql, db_id, item_name, item_value, tri)
					}
				}else{
					log.Printf("there is no recovery mode for this trigger")
				}

			}


		}else{
			log.Printf("Alert status is disable")
		}
	}
	
}

/*
func AlertBasicInfo(mysql *xorm.Engine, db_id int){
	connect := GetConnect(mysql, db_id)
	log.Printf("AlertConnect: %s", connect)
	ExecTriggers(mysql, db_id, "connect", "", connect)

}
*/

func AlertConnect(mysql *xorm.Engine, db_id int){
	connect := GetConnect(mysql, db_id)
	log.Printf("AlertConnect: %d", connect)
	ExecTriggers(mysql, db_id, "connect", "", connect)
}


func GetConnect(mysql *xorm.Engine, db_id int) string{
	var connect string = "-1"

	sql := `select connect from pms_asset_status where asset_id = ?`

	_, err := mysql.SQL(sql, db_id).Get(&connect)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	return connect
}



