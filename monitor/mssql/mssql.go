package mssql

import (
	"database/sql"
	"log"
	"opms/monitor/utils"
	"time"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xormplus/xorm"
)

func GenerateMssqlStats(wg *sync.WaitGroup, mysql *xorm.Engine, db_id int, host string, port int, alias string) {
	//Get Dsn

	//连接字符串
	//dsn := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s;;encrypt=disable", ip, port, database, user, password)
	dsn, err := GetDsn(mysql, db_id, 3)
	if err != nil {
		utils.LogDebugf("GetDsn failed: %s", err.Error())
	}
	//建立连接
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		utils.LogDebugf("Open Connection failed: %s", err.Error())
	}
	defer db.Close()

	
	err = db.Ping()
	if err != nil {
		utils.LogDebugf("DB Ping %s failed: %s", alias, err.Error())
		MoveToHistory(mysql, "pms_asset_status", "asset_id", db_id)
		MoveToHistory(mysql, "pms_mssql_status", "db_id", db_id)

		sql := `insert into pms_asset_status(asset_id, asset_type, host, port, alias, connect, created) 
				values(?,?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, db_id, 3, host, port, alias, -1, time.Now().Unix())
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
		}

		sql = `insert into pms_mssql_status(db_id, host, port, alias, connect, created) 
		values(?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, db_id, host, port, alias, -1, time.Now().Unix())
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
		}

		AlertConnect(mysql, db_id)
	} else {
		log.Println("ping succeeded")
		
		//get sqlserver basic infomation
		GatherBasicInfo(db, mysql , db_id, host, port, alias)
		AlertConnect(mysql, db_id)
		
	}

	(*wg).Done()

}

func GatherBasicInfo(db *sql.DB, mysql *xorm.Engine, db_id int, host string, port int, alias string) error{

	connect := 1
	role := 1
	uptime := GetUptime(db)
	version := GetVersion(db)

	processes := GetProcesses(db)
	processes_running := GetProcessesRunning(db)
	processes_waits := GetProcessesWaits(db)
	
	lock_timeout := GetVariables(db,"LOCK_TIMEOUT")
	trancount := GetVariables(db,"TRANCOUNT")
	max_connections := GetVariables(db,"MAX_CONNECTIONS")
	
	// storage result
	session := mysql.NewSession()
	defer session.Close()
	// add Begin() before any action
	err := session.Begin()
	//storage stats into pms_asset_status
	//move old data to history table
	MoveToHistory(mysql, "pms_asset_status", "asset_id", db_id)

	sql := `insert into pms_asset_status(asset_id, asset_type, host, port, alias, role, version, connect, sessions, created) 
						values(?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.Exec(sql, db_id, 1, host, port, alias, role, version, connect, processes, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
		return err
	}

	//storage stats into pms_mssql_status
	MoveToHistory(mysql, "pms_mssql_status", "db_id", db_id)

	sql = `insert into pms_mssql_status(db_id, host, port, alias, connect, role, uptime, version, lock_timeout, trancount, max_connections, processes, processes_running, processes_waits, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.Exec(sql, db_id, host, port, alias, connect, role, uptime, version, lock_timeout, trancount, max_connections, processes, processes_running, processes_waits, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
		return err
	}

	// add Commit() after all actions
	err = session.Commit()
	return err
}


func MoveToHistory(mysql *xorm.Engine, table_name string, key_name string, key_value int){
	sql := `insert into ` + table_name + `_his select * from ` + table_name + ` where ` + key_name + ` = ?`
	_, err := mysql.Exec(sql, key_value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	sql = `delete from ` + table_name + ` where ` + key_name + ` = ?`
	_, err = mysql.Exec(sql, key_value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}
}

func GetDsn(db *xorm.Engine, db_id int, asset_type int) (string, error) {
	var dsn string
	var sql string
	if asset_type == 1 {
		sql = `select concat("oracle://",username,":",password ,"@" , host , ":" , port , "/" , instance_name , "?sysdba=1") as dsn 
				from pms_asset_config where id = ? and asset_type = ?`
	} else if asset_type == 2 {
		sql = `select concat(username,":",password,"@tcp(",host,":",port,")/",db_name,"?charset=utf8") from pms_asset_config where id = ? and asset_type = ?`
	} else if asset_type == 3 {
		sql = `select concat("server=",host,"\\",instance_name,";port",port,";database=",case db_name when "" then "master" end,";user id=",username,";password=",password,";encrypt=disable") from pms_asset_config where id = ? and asset_type = ?`
	} else {
		sql = `select "" from pms_asset_config where id = ? and asset_type = ?`
	}

	_, err := db.SQL(sql, db_id, asset_type).Get(&dsn)
	if err != nil {
		log.Fatal(err)
	}

	return dsn, err
}
