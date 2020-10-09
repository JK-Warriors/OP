package alert

import (
	"log"
	//"time"
	//"context"
	//"opms/monitor/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

func GetSendMailGlobal(mysql *xorm.Engine) string{
	var value string
	sql := `select value from pms_global_options where id = 'send_alert_mail' `
	_, err := mysql.SQL(sql).Get(&value)
	if err != nil {
		log.Printf("GetSendMailGlobal failed: %s", err.Error())
	}

	return value
}

func GetSMTPHost(mysql *xorm.Engine) string{
	var value string
	sql := `select value from pms_global_options where id = 'smtp_host' `
	_, err := mysql.SQL(sql).Get(&value)
	if err != nil {
		log.Printf("GetSMTPHost failed: %s", err.Error())
	}

	return value
}

func GetSMTPPort(mysql *xorm.Engine) string{
	var value string
	sql := `select value from pms_global_options where id = 'smtp_port' `
	_, err := mysql.SQL(sql).Get(&value)
	if err != nil {
		log.Printf("GetSMTPPort failed: %s", err.Error())
	}

	return value
}

func GetSMTPUser(mysql *xorm.Engine) string{
	var value string
	sql := `select value from pms_global_options where id = 'smtp_user' `
	_, err := mysql.SQL(sql).Get(&value)
	if err != nil {
		log.Printf("GetSMTPUser failed: %s", err.Error())
	}

	return value
}

func GetSMTPPassword(mysql *xorm.Engine) string{
	var value string
	sql := `select value from pms_global_options where id = 'smtp_pass' `
	_, err := mysql.SQL(sql).Get(&value)
	if err != nil {
		log.Printf("GetSMTPPassword failed: %s", err.Error())
	}

	return value
}

func GetSendFrom(mysql *xorm.Engine) string{
	var value string
	sql := `select value from pms_global_options where id = 'mailfrom' `
	_, err := mysql.SQL(sql).Get(&value)
	if err != nil {
		log.Printf("GetSendFrom failed: %s", err.Error())
	}

	return value
}

func GetMailContent(mysql *xorm.Engine, alert_id int) string{
	var value string
	sql := `select concat(from_unixtime(a.created), ' [', a.severity, '] [', ac.host, ':', ac.port,' (', ac.alias, ')', '] ', a.message) 
			from pms_alerts a, pms_asset_config ac 
			where a.asset_id = ac.id 
			and a.id = ? `
	_, err := mysql.SQL(sql, alert_id).Get(&value)
	if err != nil {
		log.Printf("GetMailContent failed: %s", err.Error())
	}

	return value
}



