package mysql

import (
	"database/sql"
	"fmt"
	"log"

	//"time"
	//"context"
	//"opms/monitor/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	//"github.com/xormplus/xorm"
)

func GetVersion(db *sql.DB) string {
	var version string
	sql := `select version()`
	err := db.QueryRow(sql).Scan(&version)
	if err != nil {
		log.Printf("GetVersion failed: %s", err.Error())
	}

	return version
}

func GetRole(db *sql.DB) string {
	// var version string
	sql := `show slave status`
	_, err := db.Query(sql)
	if err != nil {
		log.Printf("GetRole failed: %s", err.Error())
		return "MASTER"
	}

	return "SLAVE"
}

func GetGlobalStatus(db *sql.DB) (map[string]string, error) {
	sql := `SHOW GLOBAL STATUS`
	globalStatusRows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer globalStatusRows.Close()

	var key string
	var val string
	var statusItems = make(map[string]string)
	for globalStatusRows.Next() {
		if err := globalStatusRows.Scan(&key, &val); err != nil {
			//return nil, err
		}

		statusItems[key] = val
		//log.Printf("%s: %s", key, val)

	}

	return statusItems, nil
}

func GetGlobalVariables(db *sql.DB) (map[string]string, error) {
	sql := `SHOW GLOBAL variables`
	variablesRows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer variablesRows.Close()

	var key string
	var val string
	var variablesItems = make(map[string]string)
	for variablesRows.Next() {
		if err := variablesRows.Scan(&key, &val); err != nil {
			//return nil, err
		}

		variablesItems[key] = val
		//log.Printf("%s: %s", key, val)

	}

	return variablesItems, nil
}

func GetGlobalVariable(db *sql.DB, variable string) (string, error) {
	var key, value string
	sql := fmt.Sprintf(`SHOW GLOBAL variables like '%s'`, variable)
	err := db.QueryRow(sql).Scan(&key, &value)
	if err != nil {
		log.Printf("GetGlobalVariable failed: %s", err.Error())
		return "", err
	}

	return value, nil
}

func GetProcessWaits(db *sql.DB) int {
	var count int
	sql := `select count(1) from information_schema.processlist where state <> '' and user <> 'repl' and time > 2`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("GetProcessWaits failed: %s", err.Error())
	}
	return count
}

func GetMasterLogSpace(db *sql.DB) (int, error) {
	sql := `SHOW MASTER LOGS`
	variablesRows, err := db.Query(sql)
	if err != nil {
		return 0, err
	}
	defer variablesRows.Close()

	var name string
	var total_space, space int = 0, 0
	for variablesRows.Next() {
		if err := variablesRows.Scan(&name, &space); err != nil {
			//return nil, err
		}

		total_space = total_space + space
		//log.Printf("%s: %s", key, val)
	}

	return total_space, nil
}

func GetMasterStatus(db *sql.DB) (string, string, error) {
	var file, position, binlog_do_db, binlog_ignore_db, executed_gtid_set string

	sql := `SHOW MASTER STATUS`
	err := db.QueryRow(sql).Scan(&file, &position, &binlog_do_db, &binlog_ignore_db, &executed_gtid_set)
	if err != nil {
		return "", "", err
	}

	return file, position, nil
}

func GetQPS(mysql *xorm.Engine, db_id int) int {
	var curr_questions int
	sql := `select questions from pms_mysql_status where db_id = ? `
	_, err := mysql.SQL(sql, db_id).Get(&curr_questions)
	if err != nil {
		log.Printf("GetQPS failed: %s", err.Error())
		curr_questions = 0
	}

	var curr_created int
	sql = `select created from pms_mysql_status where db_id = ? `
	_, err = mysql.SQL(sql, db_id).Get(&curr_created)
	if err != nil {
		log.Printf("GetQPS failed: %s", err.Error())
		curr_created = 0
	}

	var last_questions int
	sql = `select questions from pms_mysql_status_his where db_id = ? order by id desc limit 1`
	_, err = mysql.SQL(sql, db_id).Get(&last_questions)
	if err != nil {
		log.Printf("GetQPS failed: %s", err.Error())
		last_questions = 0
	}

	var last_created int
	sql = `select created from pms_mysql_status_his where db_id = ? order by id desc limit 1 `
	_, err = mysql.SQL(sql, db_id).Get(&last_created)
	if err != nil {
		log.Printf("GetQPS failed: %s", err.Error())
		last_created = 0
	}

	if last_created > 0 && curr_created > 0 && last_questions > 0 && curr_questions > 0 && curr_created > last_created && curr_questions >= last_questions {
		qps := (curr_questions - last_questions) / (curr_created - last_created)
		return qps
	} else {
		return -1
	}
}

func GetTPS(mysql *xorm.Engine, db_id int) int {
	var curr_transactions int
	sql := `select com_commit + com_rollback from pms_mysql_status where db_id = ? `
	_, err := mysql.SQL(sql, db_id).Get(&curr_transactions)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		curr_transactions = 0
	}

	var curr_created int
	sql = `select created from pms_mysql_status where db_id = ? `
	_, err = mysql.SQL(sql, db_id).Get(&curr_created)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		curr_created = 0
	}

	var last_transactions int
	sql = `select com_commit + com_rollback from pms_mysql_status_his where db_id = ? order by id desc limit 1`
	_, err = mysql.SQL(sql, db_id).Get(&last_transactions)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		last_transactions = 0
	}

	var last_created int
	sql = `select created from pms_mysql_status_his where db_id = ? order by id desc limit 1 `
	_, err = mysql.SQL(sql, db_id).Get(&last_created)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		last_created = 0
	}

	if last_created > 0 && curr_created > 0 && last_transactions > 0 && curr_transactions > 0 && curr_created > last_created && curr_transactions >= last_transactions {
		tps := (curr_transactions - last_transactions) / (curr_created - last_created)
		return tps
	} else {
		return -1
	}
}

func GetLogPerSecond(mysql *xorm.Engine, db_id int) int {
	var curr_innodb_log int
	sql := `select innodb_log from pms_mysql_status where db_id = ? `
	_, err := mysql.SQL(sql, db_id).Get(&curr_innodb_log)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		curr_innodb_log = 0
	}

	var curr_created int
	sql = `select created from pms_mysql_status where db_id = ? `
	_, err = mysql.SQL(sql, db_id).Get(&curr_created)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		curr_created = 0
	}

	var last_innodb_log int
	sql = `select innodb_log from pms_mysql_status_his where db_id = ? order by id desc limit 1`
	_, err = mysql.SQL(sql, db_id).Get(&last_innodb_log)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		last_innodb_log = 0
	}

	var last_created int
	sql = `select created from pms_mysql_status_his where db_id = ? order by id desc limit 1 `
	_, err = mysql.SQL(sql, db_id).Get(&last_created)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		last_created = 0
	}

	if last_created > 0 && curr_created > 0 && last_innodb_log > 0 && curr_innodb_log > 0 && curr_created > last_created && curr_innodb_log >= last_innodb_log {
		log_per_sec := (curr_innodb_log - last_innodb_log) / 1024 / 1024 / (curr_created - last_created)
		return log_per_sec
	} else {
		return -1
	}
}
