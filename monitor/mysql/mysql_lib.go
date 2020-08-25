package mysql

import (
	"database/sql"
	"log"
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

func GetProcessWaits(db *sql.DB) int{
	var count int
	sql := `select count(1) from information_schema.processlist where state <> '' and user <> 'repl' and time > 2`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("GetProcessWaits failed: %s", err.Error())
	}
	return count
}


