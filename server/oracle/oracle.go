package oracle

import (
	"database/sql"
	"log"
	"opms/server/utils"
	"time"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/godror/godror"
	"github.com/xormplus/xorm"
)

func GenerateOracleStats(wg *sync.WaitGroup, mysql *xorm.Engine, db_id int, host string, port string, alias string) {
	//Get Dsn
	dsn, err := GetDsn(mysql, db_id, 1)
	P, err := godror.ParseConnString(dsn)

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	connect := 1
	//get instance info 
	inst_num := Get_Instance(db, "instance_number")
	inst_name := Get_Instance(db, "instance_name")
	inst_role := Get_Instance(db, "instance_role")
	inst_status := Get_Instance(db, "status")
	version := Get_Instance(db, "version")
	startup_time := Get_Instance(db, "startup_time")
	host_name := Get_Instance(db, "host_name")
	archiver := Get_Instance(db, "archiver")

	//get database info
	db_name := Get_Database(db, "name")
	db_role := Get_Database(db, "database_role")
	open_mode := Get_Database(db, "open_mode")
	protection_mode := Get_Database(db, "protection_mode")
	flashback_on := Get_Database(db, "flashback_on")

	
	//get session
	var sessions int
	sql := `select count(1) from v$session where status = 'ACTIVE'`
	err = db.QueryRow(sql).Scan(&sessions)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	//get flashback_usage
	var flashback_usage int
	sql = `select sum(nvl(percent_space_used,0)) from v$flash_recovery_area_usage`
	err = db.QueryRow(sql).Scan(&flashback_usage)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		flashback_usage = 0
	}

	// storage result
	session := mysql.NewSession()
	defer session.Close()
	// add Begin() before any action
	err = session.Begin()
	//storage stats into pms_asset_status
	//move old data to history table
	sql = `insert into pms_asset_status_his select * from pms_asset_status where asset_id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	sql = `delete from pms_asset_status where asset_id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	sql = `insert into pms_asset_status(asset_id, asset_type, host, port, alias, role, version, connect, sessions, created) 
						values(?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.Exec(sql, db_id, 1, host, port, alias, db_role, version, connect, sessions, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	//storage stats into pms_oracle_status
	sql = `insert into pms_oracle_status_his select * from pms_oracle_status where db_id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	sql = `delete from pms_oracle_status where db_id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	sql = `insert into pms_oracle_status(db_id, connect, inst_num, inst_name, inst_role, inst_status, version, startup_time, host_name, archiver, db_name, db_role, open_mode, protection_mode, flashback_on, flashback_usage, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.Exec(sql, db_id, connect, inst_num, inst_name, inst_role, inst_status, version, startup_time, host_name, archiver, db_name, db_role, open_mode, protection_mode, flashback_on, flashback_usage, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	// add Commit() after all actions
	err = session.Commit()
	(*wg).Done()

}

func Get_Instance(db *sql.DB, matrix_name string) string{
	var matrix_value string
	sql := `select ` + matrix_name + ` from v$instance`
	err := db.QueryRow(sql).Scan(&matrix_value)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		return ""
	}
	return matrix_value
}

func Get_Database(db *sql.DB, matrix_name string) string{
	var matrix_value string
	sql := `select ` + matrix_name + ` from v$database`
	err := db.QueryRow(sql).Scan(&matrix_value)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		return ""
	}
	return matrix_value
}

func GetDsn(db *xorm.Engine, db_id int, db_type int) (string, error) {
	var dsn string
	var sql string
	if db_type == 1 {
		sql = `select concat("oracle://",username,":",password ,"@" , host , ":" , port , "/" , instance_name , "?sysdba=1") as dsn 
				from pms_db_config where id = ? and db_type = ?`
	} else if db_type == 2 {
		sql = `select host from pms_db_config where id = ? and db_type = ?`
	} else if db_type == 3 {
		sql = `select host from pms_db_config where id = ? and db_type = ?`
	} else {
		sql = `select host from pms_db_config where id = ? and db_type = ?`
	}

	_, err := db.SQL(sql, db_id, db_type).Get(&dsn)
	if err != nil {
		log.Fatal(err)
	}

	return dsn, err
}
