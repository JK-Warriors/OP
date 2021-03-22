package mssql

import (
	"database/sql"
	"log"

	//"time"
	//"context"
	//"opms/monitor/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	//"github.com/godror/godror"
	//"github.com/xormplus/xorm"
)

func GetVersion(db *sql.DB) string {
	var version string
	var s_version string
	sql := `SELECT @@VERSION`
	err := db.QueryRow(sql).Scan(&version)
	if err != nil {
		log.Printf("GetVersion failed: %s", err.Error())
	} else {
		s_version = version[21:25]
	}

	return s_version
}

func GetUptime(db *sql.DB) string {
	var uptime string
	sql := `SELECT CONVERT(varchar, sqlserver_start_time, 120) as time_restart FROM sys.dm_os_sys_info`
	err := db.QueryRow(sql).Scan(&uptime)
	if err != nil {
		log.Printf("GetUptime failed: %s", err.Error())
	}
	return uptime
}

func GetVariables(db *sql.DB, matrix_name string) string {
	var matrix_value string
	sql := `select @@` + matrix_name
	err := db.QueryRow(sql).Scan(&matrix_value)
	if err != nil {
		log.Printf("GetVariables for %s failed: %s", matrix_name, err.Error())
	}
	return matrix_value
}

func GetTotalSessions(db *sql.DB) int {
	var total_sessions int
	sql := `select count(1) from sys.dm_exec_sessions`
	err := db.QueryRow(sql).Scan(&total_sessions)
	if err != nil {
		log.Printf("GetTotalSessions failed: %s", err.Error())
		return -1
	}
	return total_sessions
}

func GetActiveSessions(db *sql.DB) int {
	var total_sessions int
	sql := `select count(1) from sys.dm_exec_sessions where status = 'running' `
	err := db.QueryRow(sql).Scan(&total_sessions)
	if err != nil {
		log.Printf("GetActiveSessions failed: %s", err.Error())
		return -1
	}
	return total_sessions
}

func GetProcesses(db *sql.DB) int {
	var total_processes int
	sql := `SELECT count(*) FROM [master].[dbo].[sysprocesses] WHERE [DBID] > 0`
	err := db.QueryRow(sql).Scan(&total_processes)
	if err != nil {
		log.Printf("GetProcesses failed: %s", err.Error())
		return -1
	}
	return total_processes
}

func GetProcessesRunning(db *sql.DB) int {
	var running_processes int
	sql := `SELECT COUNT(*) FROM [master].[dbo].[sysprocesses] WHERE [DBID] >0 AND status !='SLEEPING' AND status !='BACKGROUND'`
	err := db.QueryRow(sql).Scan(&running_processes)
	if err != nil {
		log.Printf("GetProcessesRunning failed: %s", err.Error())
		return -1
	}
	return running_processes
}

func GetProcessesWaits(db *sql.DB) int {
	var wait_processes int
	sql := `SELECT COUNT(*) FROM [master].[dbo].[sysprocesses] WHERE [DBID] >0 AND status ='SUSPENDED' AND waittime >2`
	err := db.QueryRow(sql).Scan(&wait_processes)
	if err != nil {
		log.Printf("GetProcessesWaits failed: %s", err.Error())
		return -1
	}
	return wait_processes
}

func GetBytesWritten(db *sql.DB) int {
	var bytes_written int
	sql := `select SUM(BytesWritten) from (select * from sysfiles where fileid = 2) aa cross apply fn_virtualfilestats(null,aa.fileid) eqp `
	err := db.QueryRow(sql).Scan(&bytes_written)
	if err != nil {
		log.Printf("GetBytesWritten failed: %s", err.Error())
		return -1
	}
	return bytes_written
}

func GetQPS(db *sql.DB) int {
	var cntr_value int
	sql := `select cntr_value from sys.dm_os_performance_counters where counter_name='Batch Requests/sec' `
	err := db.QueryRow(sql).Scan(&cntr_value)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		cntr_value = -1
	}
	return cntr_value
}

func GetTPS(db *sql.DB) int {
	var cntr_value int
	sql := `select cntr_value from sys.dm_os_performance_counters where counter_name= 'Transactions/sec' and object_name='SQLServer:Databases' and instance_name = '_Total' `
	err := db.QueryRow(sql).Scan(&cntr_value)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		cntr_value = -1
	}
	return cntr_value
}

func GetBufferCacheHit(db *sql.DB) int {
	var count int
	sql := `select cast(cast((a.cntr_value * 1.0 / b.cntr_value) * 100 as int) as varchar(20)) as buffercachehitratio
			from (select * from sys.dm_os_performance_counters where counter_name = 'Buffer cache hit ratio') a
			cross join (select * from sys.dm_os_performance_counters where counter_name = 'Buffer cache hit ratio base') b `
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		count = 0
	}
	return count
}

func GetLogPerSecond(mysql *xorm.Engine, db_id int) int {
	var curr_bytes_written int
	sql := `select bytes_written from pms_mssql_status where db_id = ? `
	_, err := mysql.SQL(sql, db_id).Get(&curr_bytes_written)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		curr_bytes_written = 0
	}

	var curr_created int
	sql = `select created from pms_mssql_status where db_id = ? `
	_, err = mysql.SQL(sql, db_id).Get(&curr_created)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		curr_created = 0
	}

	var last_bytes_written int
	sql = `select bytes_written from pms_mssql_status_his where db_id = ? order by id desc limit 1`
	_, err = mysql.SQL(sql, db_id).Get(&last_bytes_written)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		last_bytes_written = 0
	}

	var last_created int
	sql = `select created from pms_mssql_status_his where db_id = ? order by id desc limit 1 `
	_, err = mysql.SQL(sql, db_id).Get(&last_created)
	if err != nil {
		log.Printf("GetTPS failed: %s", err.Error())
		last_created = 0
	}

	if last_created > 0 && curr_created > 0 && last_bytes_written > 0 && curr_bytes_written > 0 && curr_created > last_created && curr_bytes_written >= last_bytes_written {
		log_per_sec := (curr_bytes_written - last_bytes_written) / 1024 / 1024 / (curr_created - last_created)
		return log_per_sec
	} else {
		return -1
	}
}
