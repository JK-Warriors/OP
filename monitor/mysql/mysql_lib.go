package mysql

import (
	"database/sql"
	"log"
	"fmt"
	//"time"
	//"context"
	//"opms/monitor/utils"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/xormplus/xorm"
)


	   

func GetVersion(db *sql.DB) string{
	var version string
	sql := `select version()`
	err := db.QueryRow(sql).Scan(&version)
	if err != nil {
		log.Printf("GetVersion failed: %s", err.Error())
	}

	return version
}

func GetGlobalStatus(db *sql.DB) (map[string]string, error){
	sql :=`SHOW GLOBAL STATUS`
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

func GetGlobalVariables(db *sql.DB) (map[string]string, error){
	sql :=`SHOW GLOBAL variables`
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

func GetGlobalVariable(db *sql.DB, variable string) (string, error){
	var key, value string
	sql :=fmt.Sprintf(`SHOW GLOBAL variables like '%s'`, variable)
	err := db.QueryRow(sql).Scan(&key, &value)
	if err != nil {
		log.Printf("GetGlobalVariable failed: %s", err.Error())
		return "", err
	}

	return value, nil
}

func GetProcessWaits(db *sql.DB) int{
	var count int
	sql := `select count(1) from information_schema.processlist where state <> '' and user <> 'repl' and time > 2`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("GetProcessWaits failed: %s", err.Error())
	}
	return count
}

func GetMasterLogSpace(db *sql.DB) (int, error){
	sql :=`SHOW MASTER LOGS`
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


func GetMasterStatus(db *sql.DB) (string, string, error){
	var file, position, binlog_do_db, binlog_ignore_db, executed_gtid_set string

	sql :=`SHOW MASTER STATUS`
	err := db.QueryRow(sql).Scan(&file, &position, &binlog_do_db, &binlog_ignore_db, &executed_gtid_set)
	if err != nil {
		return "", "", err
	}
	
	return file, position, nil
}

