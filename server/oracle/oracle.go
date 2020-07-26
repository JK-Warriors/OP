package oracle

import (
	"database/sql"
	"log"
	"opms/server/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/godror/godror"
	"github.com/xormplus/xorm"
)

func GenerateOracleBasic(mysql *xorm.Engine, P godror.ConnectionParams, db_id int) {

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	//get db_role, open_mode, flashback_on
	var db_role, open_mode, flashback_on string
	sql := `select database_role, open_mode, flashback_on from v$database`
	err = db.QueryRow(sql).Scan(&db_role, &open_mode, &flashback_on)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	//get version
	var version string
	sql = `select version from v$instance`
	err = db.QueryRow(sql).Scan(&version)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	//get flashback_usage
	var flashback_usage string
	sql = `select sum(percent_space_used) from v$flash_recovery_area_usage`
	err = db.QueryRow(sql).Scan(&flashback_usage)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	// storage result
	session := mysql.NewSession()
	defer session.Close()
	// add Begin() before any action
	err = session.Begin()
	//move old data to history table
	sql = `insert into pms_db_status_his select * from pms_db_status where id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		session.Rollback()
		return
	}

	sql = `delete from pms_db_status where id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		session.Rollback()
		return
	}

	sql = `insert into pms_db_status(id, db_type, connect, role, version, open_mode, flashback_on, flashback_usage, created) 
						values(?,?,?,?,?,?,?,?,?)`

	_, err = mysql.Exec(sql, db_id, 1, 1, db_role, version, open_mode, flashback_on, flashback_usage, time.Now().Unix())

	if err != nil {
		log.Printf("%s: %w", sql, err)
		session.Rollback()
		return
	}
	// add Commit() after all actions
	err = session.Commit()

}

func GeneratePrimary(mysql *xorm.Engine, P godror.ConnectionParams, db_id int) {

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
	}
	defer db.Close()

	sql := `select dest_id, transmit_mode, thread, sequence, archived, applied, current_scn, curr_db_time
			from (select t.dest_id,
						transmit_mode,
						thread# as thread,
						sequence#+1 as sequence,
						archived,
						applied,
						current_scn,
						to_char(scn_to_timestamp(current_scn), 'yyyy-mm-dd hh24:mi:ss') curr_db_time,
						row_number() over(partition by thread# order by sequence# desc) rn
					from v$archived_log t, v$archive_dest a, v$database d
					where t.dest_id = a.dest_id
					and t.dest_id = 2)
			where rn = 1`
	rows, err := db.Query(sql)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}
	defer rows.Close()

	for rows.Next() {
		var dest_id, thread, sequence int
		var transmit_mode, archived, applied, curr_db_time string
		var current_scn int64

		if err = rows.Scan(&dest_id, &transmit_mode, &thread, &sequence, &archived, &applied, &current_scn, &curr_db_time); err != nil {
			log.Println(err.Error())
		}
		// log.Printf("dest_id: %d", dest_id)
		// log.Printf("thread: %d", thread)
		// log.Printf("sequence: %d", sequence)
		// log.Printf("transmit_mode: %s", transmit_mode)
		// log.Printf("archived: %s", archived)
		// log.Printf("applied: %s", applied)
		// log.Printf("current_scn: %d", current_scn)
		// log.Printf("curr_db_time: %s", curr_db_time)

		//get archived delay
		var archived_delay int
		sql = `select count(1) from v$archived_log where dest_id = :1 and thread# = :2 and archived= 'NO'`
		err = db.QueryRow(sql, dest_id, thread).Scan(&archived_delay)
		if err != nil {
			log.Printf("%s: %w", sql, err)
		}

		//get applied delay
		var applied_delay int
		sql = `select count(1) from v$archived_log where dest_id = :1 and thread# = :2 and applied= 'NO'`
		err = db.QueryRow(sql, dest_id, thread).Scan(&applied_delay)
		if err != nil {
			log.Printf("%s: %w", sql, err)
		}

		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err := session.Begin()
		//move old data to history table
		sql = `insert into pms_dr_pri_status_his select * from pms_dr_pri_status where db_id = ?`
		_, err = mysql.Exec(sql, db_id)
		if err != nil {
			log.Printf("%s: %w", sql, err)
			session.Rollback()
			return
		}

		sql = `delete from pms_dr_pri_status where db_id = ?`
		_, err = mysql.Exec(sql, db_id)
		if err != nil {
			log.Printf("%s: %w", sql, err)
			session.Rollback()
			return
		}

		sql = `insert into pms_dr_pri_status(db_id, dest_id, transmit_mode, thread, sequence, curr_scn, curr_db_time, archived_delay, applied_delay, created) 
						values(?,?,?,?,?,?,?,?,?,?)`

		_, err = mysql.Exec(sql, db_id, dest_id, transmit_mode, thread, sequence, current_scn, curr_db_time, archived_delay, applied_delay, time.Now().Unix())

		if err != nil {
			log.Printf("%s: %w", sql, err)
			session.Rollback()
			return
		}
		// add Commit() after all actions
		err = session.Commit()

	}

}

func GenerateStandby(mysql *xorm.Engine, p_pri godror.ConnectionParams, p_sta godror.ConnectionParams, db_id int) {
	db, err := sql.Open("godror", p_sta.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", p_sta.StringWithPassword(), err)
	}
	defer db.Close()

	var thread, sequence, block, delay_mins int
	sql := `select ms.thread#,
					ms.sequence#,
					ms.block#,
					ms.delay_mins
				from v$managed_standby  ms
				where ms.process in ('MRP0')
				and ms.sequence# <> 0`
	err = db.QueryRow(sql).Scan(&thread, &sequence, &block, &delay_mins)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	//get apply_rate
	var apply_rate int
	sql = `select avg_apply_rate
			from (select rp.sofar avg_apply_rate
					from v$recovery_progress rp
				where rp.item = 'Average Apply Rate'
				order by start_time desc)
		where rownum < 2`
	err = db.QueryRow(sql).Scan(&apply_rate)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	//get sta_scn
	var sta_scn int64
	sql = `select current_scn from v$database`
	err = db.QueryRow(sql).Scan(&sta_scn)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	//get standby db time by sta_scn
	curr_db_time, err := GetDbtimeBySCN(p_pri, sta_scn)
	if err != nil {
		log.Printf("%w", err)
		curr_db_time = ""
	}

	//get mrp_status
	var mrp_status string
	sql = `select status from gv$session where program like '%(MRP0)'`
	err = db.QueryRow(sql).Scan(&mrp_status)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}

	// storage result
	session := mysql.NewSession()
	defer session.Close()
	// add Begin() before any action
	err = session.Begin()
	//move old data to history table
	sql = `insert into pms_dr_sta_status_his select * from pms_dr_sta_status where db_id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		session.Rollback()
		return
	}

	sql = `delete from pms_dr_sta_status where db_id = ?`
	_, err = mysql.Exec(sql, db_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		session.Rollback()
		return
	}

	sql = `insert into pms_dr_sta_status(db_id, thread, sequence, block, delay_mins, apply_rate, curr_scn, curr_db_time, mrp_status, created) 
						values(?,?,?,?,?,?,?,?,?,?)`

	_, err = mysql.Exec(sql, db_id, thread, sequence, block, delay_mins, apply_rate, sta_scn, curr_db_time, mrp_status, time.Now().Unix())

	if err != nil {
		log.Printf("%s: %w", sql, err)
		session.Rollback()
		return
	}
	// add Commit() after all actions
	err = session.Commit()
}

func GetDbtimeBySCN(P godror.ConnectionParams, scn int64) (string, error) {
	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebugf("%s: %w", P.StringWithPassword(), err)
		return "", err
	}
	defer db.Close()

	//get curr_db_time
	var curr_db_time string
	sql := `select to_char(scn_to_timestamp(:1), 'yyyy-mm-dd hh24:mi:ss') from v$database`
	err = db.QueryRow(sql, scn).Scan(&curr_db_time)
	if err != nil {
		log.Printf("%s: %w", sql, err)
	}
	return curr_db_time, err
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
