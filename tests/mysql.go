package mysql

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(192.168.210.240:3306)/mysql"

	//建立连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		utils.LogDebugf("Open Connection failed: %s", err.Error())
	}
	defer db.Close()

	
	err = db.Ping()
	if err != nil {
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
			log.Printf("%s: %s", key, val)
	
		}
		
	}
}