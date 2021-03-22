package oracle

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"opms/monitor/utils"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/godror/godror"
	"github.com/xormplus/xorm"
)

func GenerateOracleStats(wg *sync.WaitGroup, mysql *xorm.Engine, db_id int, host string, port int, alias string) {
	//Get Dsn
	dsn, err := GetDsn(mysql, db_id, 1)
	P, err := godror.ParseConnString(dsn)

	db, err := sql.Open("godror", P.StringWithPassword())
	defer db.Close()
	if err != nil {
		utils.LogDebugf("%s: %s", P.StringWithPassword(), err.Error())

	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		utils.LogDebugf("PingContext: %s", err.Error())
		MoveToHistory(mysql, "pms_asset_status", "asset_id", db_id)
		MoveToHistory(mysql, "pms_oracle_status", "db_id", db_id)

		sql := `insert into pms_asset_status(asset_id, asset_type, host, port, alias, connect, created) 
				values(?,?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, db_id, 1, host, port, alias, -1, time.Now().Unix())
		if err != nil {
			log.Printf("%s: %w", sql, err)
		}

		sql = `insert into pms_oracle_status(db_id, host, port, alias, connect, created) 
		values(?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, db_id, host, port, alias, -1, time.Now().Unix())
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
		}

		AlertConnect(mysql, db_id)
	} else {
		log.Println("ping succeeded")
		//if !ok {
		//	log.Println("ping succeeded after deadline!")
		//}
		//get oracle basic infomation
		GatherBasicInfo(db, mysql, db_id, host, port, alias)
		AlertBasicInfo(mysql, db_id)

		GatherRedo(db, mysql, db_id, host, port, alias)
		GatherDbTime(db, mysql, db_id, host, port, alias)
		GatherMetricValue(db, mysql, db_id, host, port, alias)

		//get tablespace
		GatherTablespaces(db, mysql, db_id, host, port, alias)
		AlertTablespaces(mysql, db_id)

		//get asm diskgroup
		GatherDiskgroups(db, mysql, db_id, host, port, alias)
		AlertDiskgroups(mysql, db_id)
	}

	(*wg).Done()

}

func GatherBasicInfo(db *sql.DB, mysql *xorm.Engine, db_id int, host string, port int, alias string) {

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

	//get sessions
	session_total := GetSessionTotal(db)
	session_actives := GetSessionActive(db)
	session_waits := GetSessionWait(db)

	//get flashback_usage
	flashback_usage := GetFlashbackUsage(db)

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
	_, err = mysql.Exec(sql, db_id, 1, host, port, alias, db_role, version, connect, session_total, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	//storage stats into pms_oracle_status
	MoveToHistory(mysql, "pms_oracle_status", "db_id", db_id)

	sql = `insert into pms_oracle_status(db_id, host, port, alias, connect, inst_num, inst_name, inst_role, inst_status, version, startup_time, host_name, archiver, db_name, db_role, open_mode, protection_mode, session_total, session_actives, session_waits, flashback_on, flashback_usage, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.Exec(sql, db_id, host, port, alias, connect, inst_num, inst_name, inst_role, inst_status, version, startup_time, host_name, archiver, db_name, db_role, open_mode, protection_mode, session_total, session_actives, session_waits, flashback_on, flashback_usage, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	// add Commit() after all actions
	err = session.Commit()
}

func GatherMetricValue(db *sql.DB, mysql *xorm.Engine, db_id int, host string, port int, alias string) {
	timestamp := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	session_total := GetSessionTotal(db)
	session_actives := GetSessionActive(db)
	StorageMetricData(mysql, db_id, "TotalSessions", timestamp, strconv.Itoa(session_total), "GAUGE")
	StorageMetricData(mysql, db_id, "ActiveSessions", timestamp, strconv.Itoa(session_actives), "GAUGE")

	//get QPS
	qps := GetQPS(db)
	StorageMetricData(mysql, db_id, "Queries Per Second", timestamp, strconv.Itoa(qps), "GAUGE")

	//get TPS
	tps := GetTPS(db)
	StorageMetricData(mysql, db_id, "Transactions Per Second", timestamp, strconv.Itoa(tps), "GAUGE")

	//get Buffer Cache Hit
	bchit := GetBufferCacheHit(db)
	StorageMetricData(mysql, db_id, "Buffer Cache Hit", timestamp, strconv.Itoa(bchit), "GAUGE")

	//get Redo
	log_per_sec := GetLogPerSecond(db)
	StorageMetricData(mysql, db_id, "Log Per Second", timestamp, log_per_sec, "GAUGE")
}

// Storage metric data
func StorageMetricData(mysql *xorm.Engine, db_id int, metric string, timestamp string, value string, counterType string) {

	sql := `insert into pms_metric_data(db_id, metric, timestamp, value, counterType) 
			values(?,?,?,?,?)`
	_, err := mysql.Exec(sql, db_id, metric, timestamp, value, counterType)
	if err != nil {
		log.Printf("StorageMetricData -- %s: %s", sql, err.Error())
	}
}

func GatherRedo(db *sql.DB, mysql *xorm.Engine, db_id int, host string, port int, alias string) {
	log.Printf("GatherRedo begin for %d", db_id)
	session := mysql.NewSession()
	defer session.Close()
	err := session.Begin()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sql := `select to_char(first_time, 'yyyy-mm-dd hh24')||':00' key_time,
						trunc(sum(blocks * block_size) / 1024 / 1024) redo
			from v$archived_log
			where to_char(first_time, 'yyyymmddhh24') > to_char(sysdate-7, 'yyyymmddhh24')
			and standby_dest = 'NO'
			group by to_char(first_time, 'yyyy-mm-dd hh24')`
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
	} else {
		defer rows.Close()
		for rows.Next() {
			var key_time string
			var redo int
			err = rows.Scan(&key_time, &redo)
			if err != nil {
				log.Println(err.Error())
			}

			sql = `insert into pms_oracle_redo(db_id, key_time, redo_log, created) 
							values(?,?,?,?)`
			_, err = mysql.Exec(sql, db_id, key_time, redo, time.Now().Unix())
			if err != nil {
				log.Printf("%s: %w", sql, err)
			}
		}
		err = session.Commit()
		log.Printf("GatherRedo end for %d", db_id)
	}
}

func GatherDbTime(db *sql.DB, mysql *xorm.Engine, db_id int, host string, port int, alias string) {
	inst_id := GetCurrentInstanceNumber(db)
	snap_id := GetLastSnapId(db)

	// check snap id exists
	var is_exists int
	sql := `select count(1) from pms_oracle_db_time where db_id = ? and snap_id = ?`
	_, _ = mysql.SQL(sql, db_id, snap_id).Get(&is_exists)
	if is_exists > 0 {
		return
	}

	session := mysql.NewSession()
	defer session.Close()
	err := session.Begin()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sql = fmt.Sprintf(`select snap_id, end_time, dbtime, elapsed, round(dbtime/elapsed, 2) as rate 
			from ( select n.stat_name as name,
					e.snap_id,
					to_char(te.end_interval_time,'yyyy-mm-dd hh24:mi:ss') as end_time,
						round((case when (e.value - b.value) > 0 then e.value - b.value else e.value end) / 1000 / 1000, 2) as dbtime,
						ceil((to_date(to_char(te.end_interval_time,'yyyy-mm-dd hh24:mi:ss'),'yyyy-mm-dd hh24:mi:ss') - 
						to_date(to_char(tb.end_interval_time,'yyyy-mm-dd hh24:mi:ss'),'yyyy-mm-dd hh24:mi:ss'))*86400) as elapsed
					from wrh$_sys_time_model e, wrh$_sys_time_model b, wrh$_stat_name n, wrm$_snapshot tb, wrm$_snapshot te
					where e.stat_id = n.stat_id
					and b.stat_id = n.stat_id
					and b.snap_id = e.snap_id - 1
					and e.snap_id = %d
					and e.snap_id = te.snap_id and e.instance_number = te.instance_number
					and b.snap_id = tb.snap_id and b.instance_number = tb.instance_number
					and e.instance_number=b.instance_number
					and e.instance_number= %d
					and n.stat_name = 'DB time') tmp`, snap_id, inst_id)
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
	} else {
		defer rows.Close()
		for rows.Next() {
			var snap_id, elapsed int
			var end_time string
			var db_time, rate float64
			err = rows.Scan(&snap_id, &end_time, &db_time, &elapsed, &rate)
			if err != nil {
				log.Println(err.Error())
			}

			sql = `insert into pms_oracle_db_time(db_id, snap_id, end_time, db_time, elapsed, rate, created) 
							values(?,?,?,?,?,?,?)`
			_, err = mysql.Exec(sql, db_id, snap_id, end_time, db_time, elapsed, rate, time.Now().Unix())
			if err != nil {
				log.Printf("%s: %w", sql, err)
			}
		}
		err = session.Commit()
	}
}

func GatherTablespaces(db *sql.DB, mysql *xorm.Engine, db_id int, host string, port int, alias string) {

	session := mysql.NewSession()
	defer session.Close()
	err := session.Begin()
	//move old data to history table
	MoveToHistory(mysql, "pms_oracle_tablespace", "db_id", db_id)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sql := `select tpsname,status,mgr,max_size,curr_size, max_used
				from (SELECT d.tablespace_name tpsname,
							d.status status,
							d.segment_space_management mgr,
							TO_CHAR(NVL(trunc(A.maxbytes / 1024 / 1024), 0), '99999990') max_size,
							TO_CHAR(NVL(trunc(a.bytes / 1024 / 1024), 0), '99999990') curr_size,
							TO_CHAR(NVL((a.bytes - NVL(f.bytes, 0)) / a.bytes * 100, 0),
									'990D00') c_used,
							TO_CHAR(NVL((a.bytes - NVL(f.bytes, 0)) / a.maxbytes * 100, 0),
									'990D00') max_used
						FROM sys.dba_tablespaces d,
							(SELECT tablespace_name,
									sum(bytes) bytes,
									SUM(case autoextensible
										when 'NO' then
											BYTES
										when 'YES' then
											MAXBYTES
										else
											null
										end) maxbytes
								FROM dba_data_files
							GROUP BY tablespace_name) a,
							(SELECT tablespace_name,
									SUM(bytes) bytes,
									MAX(bytes) largest_free
								FROM dba_free_space
							GROUP BY tablespace_name) f
					WHERE d.tablespace_name = a.tablespace_name
						AND d.tablespace_name = f.tablespace_name(+))
			order by max_used desc`
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
	} else {
		defer rows.Close()
		for rows.Next() {
			var tbsname, status, mgr, max_size, curr_size, max_used string
			err = rows.Scan(&tbsname, &status, &mgr, &max_size, &curr_size, &max_used)
			if err != nil {
				log.Println(err.Error())
			}

			sql = `insert into pms_oracle_tablespace(db_id, host, port, alias, tablespace_name, status, management, total_size, used_size, max_rate, created) 
							values(?,?,?,?,?,?,?,?,?,?,?)`
			_, err = mysql.Exec(sql, db_id, host, port, alias, tbsname, status, mgr, max_size, curr_size, max_used, time.Now().Unix())
			if err != nil {
				log.Printf("%s: %w", sql, err)
			}
		}
		err = session.Commit()
	}

}

func GatherDiskgroups(db *sql.DB, mysql *xorm.Engine, db_id int, host string, port int, alias string) {

	session := mysql.NewSession()
	defer session.Close()
	err := session.Begin()
	//move old data to history table
	MoveToHistory(mysql, "pms_oracle_diskgroup", "db_id", db_id)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sql := `select name,
				state,
				type,
				total_mb,
				free_mb,
				trunc(((total_mb - free_mb) / total_mb) * 100, 2) used_rate
			from v$asm_diskgroup`
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
	} else {
		defer rows.Close()
		for rows.Next() {
			var group_name, state, group_type, total_mb, free_mb, used_rate string
			err = rows.Scan(&group_name, &state, &group_type, &total_mb, &free_mb, &used_rate)
			if err != nil {
				log.Println(err.Error())
			}

			sql = `insert into pms_oracle_diskgroup(db_id, host, port, alias, diskgroup_name, state, type, total_mb, free_mb, used_rate, created) 
							values(?,?,?,?,?,?,?,?,?,?,?)`
			_, err = mysql.Exec(sql, db_id, host, port, alias, group_name, state, group_type, total_mb, free_mb, used_rate, time.Now().Unix())
			if err != nil {
				log.Printf("%s: %w", sql, err)
			}
		}
		err = session.Commit()
	}

}

func MoveToHistory(mysql *xorm.Engine, table_name string, key_name string, key_value int) {
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
