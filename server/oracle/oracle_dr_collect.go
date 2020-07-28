package oracle

import (
	"database/sql"
	"log"
	"opms/server/utils"
	"opms/server/common"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/godror/godror"
	"github.com/xormplus/xorm"
)

func GenerateOracleDrStats(wg *sync.WaitGroup, mysql *xorm.Engine, dis common.Dr) {
	var pri_id int
	var sta_id int
	if dis.Is_Shift == 0 {
		pri_id = dis.Db_Id_P
		sta_id = dis.Db_Id_S
	} else {
		pri_id = dis.Db_Id_S
		sta_id = dis.Db_Id_P

	}

	dsn_p, err := GetDsn(mysql, pri_id, 1)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	dsn_s, err := GetDsn(mysql, sta_id, 1)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	p_pri, _ := godror.ParseConnString(dsn_p)
	p_sta, _ := godror.ParseConnString(dsn_s)


	GeneratePrimary(mysql, p_pri, pri_id)
	GenerateStandby(mysql, p_pri, p_sta, sta_id)
	
	(*wg).Done()
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
