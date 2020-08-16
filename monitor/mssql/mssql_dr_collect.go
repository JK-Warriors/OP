package mssql

import (
	"database/sql"
	"log"
	"opms/monitor/utils"
	"opms/monitor/common"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

func GenerateMssqlDrStats(wg *sync.WaitGroup, mysql *xorm.Engine, dis common.Dr) {
	dr_id := dis.Id
	db_name := dis.Db_Name

	var pri_id int
	var sta_id int
	if dis.Is_Shift == 0 {
		pri_id = dis.Db_Id_P
		sta_id = dis.Db_Id_S
	} else {
		pri_id = dis.Db_Id_S
		sta_id = dis.Db_Id_P
	}

	dsn_p, err := GetDsn(mysql, pri_id, 3)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	dsn_s, err := GetDsn(mysql, sta_id, 3)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	GeneratePrimary(mysql, dr_id, pri_id, dsn_p, db_name)
	GenerateStandby(mysql, dr_id, sta_id, dsn_s, db_name)
	

	log.Println("获取SQLServer容灾数据结束！")

	(*wg).Done()
}


func GeneratePrimary(mysql *xorm.Engine, dr_id int, db_id int, connStr string, db_name string) {
	db, err := sql.Open("mssql", connStr)
	if err != nil {
		utils.LogDebugf("%s: %s", connStr, err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		utils.LogDebugf("DB Ping failed: %s", err.Error())
		return 
	}
	
	var database_id, role, state, safety_level, connection_timeout int
	var name, state_desc, partner_name, partner_instance, redo_queue string
	var failover_lsn , end_of_log_lsn, replication_lsn int64

	sql := `select m.database_id,
					d.name,
					m.mirroring_role,
					m.mirroring_state,
					m.mirroring_state_desc,
					m.mirroring_safety_level,
					m.mirroring_partner_name,
					m.mirroring_partner_instance,
					m.mirroring_failover_lsn,
					m.mirroring_connection_timeout,
					isnull(m.mirroring_redo_queue, -1) as mirroring_redo_queue,
					m.mirroring_end_of_log_lsn,
					m.mirroring_replication_lsn
				from sys.database_mirroring m, sys.databases d
				where m.mirroring_guid is NOT NULL
				AND m.database_id = d.database_id
				and d.name = ?`
	err = db.QueryRow(sql, db_name).Scan(&database_id, &name, &role, &state, &state_desc, &safety_level, &partner_name, &partner_instance, &failover_lsn, &connection_timeout, &redo_queue, &end_of_log_lsn, &replication_lsn)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		return
	}else{
		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err := session.Begin()
		//move old data to history table
		MoveToHistory(mysql, "pms_dr_mssql_p", "db_id", db_id)

		sql = `insert into pms_dr_mssql_p(dr_id, db_id, database_id, db_name, role, state, state_desc, safety_level, partner_name, partner_instance, failover_lsn, connection_timeout, redo_queue, end_of_log_lsn, replication_lsn, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

		_, err = mysql.Exec(sql, dr_id, db_id, database_id, name, role, state, state_desc, safety_level, partner_name, partner_instance, failover_lsn, connection_timeout, redo_queue, end_of_log_lsn, replication_lsn, time.Now().Unix())

		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			session.Rollback()
			return
		}
		// add Commit() after all actions
		err = session.Commit()
	}

	

}

func GenerateStandby(mysql *xorm.Engine, dr_id int, db_id int, connStr string, db_name string) {
	db, err := sql.Open("mssql", connStr)
	if err != nil {
		utils.LogDebugf("%s: %s", connStr, err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		utils.LogDebugf("DB Ping failed: %s", err.Error())
		return 
	}

	
	var database_id, role, state, safety_level, connection_timeout int
	var name, master_server, master_port, state_desc, partner_name, redo_queue, partner_instance string
	var failover_lsn , end_of_log_lsn, replication_lsn int64
	sql := `select m.database_id,
					d.name,
					substring(mirroring_partner_name, 7, charindex(':',mirroring_partner_name,7)-7) as master_server,
					right(mirroring_partner_name, len(mirroring_partner_name) - charindex(':',mirroring_partner_name,7)) as master_port,
					m.mirroring_role,
					m.mirroring_state,
					m.mirroring_state_desc,
					m.mirroring_safety_level,
					m.mirroring_partner_name,
					m.mirroring_partner_instance,
					m.mirroring_failover_lsn,
					m.mirroring_connection_timeout,
					isnull(m.mirroring_redo_queue, -1) as mirroring_redo_queue,
					m.mirroring_end_of_log_lsn,
					m.mirroring_replication_lsn
				from sys.database_mirroring m, sys.databases d
				where m.mirroring_guid is NOT NULL
				AND m.database_id = d.database_id
				and d.name = ?`
	err = db.QueryRow(sql, db_name).Scan(&database_id, &name, &master_server, &master_port, &role, &state, &state_desc, &safety_level, &partner_name, &partner_instance, &failover_lsn, &connection_timeout, &redo_queue, &end_of_log_lsn, &replication_lsn)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		return
	}else{
		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err = session.Begin()
		//move old data to history table
		MoveToHistory(mysql, "pms_dr_mssql_s", "db_id", db_id)

		sql = `insert into pms_dr_mssql_s(dr_id, db_id, database_id, db_name, master_server, master_port, role, state, state_desc, safety_level, partner_name, partner_instance, failover_lsn, connection_timeout, redo_queue, end_of_log_lsn, replication_lsn, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

		_, err = mysql.Exec(sql, dr_id, db_id, database_id, name, master_server, master_port, role, state, state_desc, safety_level, partner_name, partner_instance, failover_lsn, connection_timeout, redo_queue, end_of_log_lsn, replication_lsn, time.Now().Unix())

		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			session.Rollback()
			return
		}
		// add Commit() after all actions
		err = session.Commit()
	}

}
