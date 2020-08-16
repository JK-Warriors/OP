package mssql

import (
	"database/sql"
	"log"
	//"time"
	//"context"
	//"opms/monitor/utils"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/godror/godror"
	//"github.com/xormplus/xorm"
)
func GetVersion(db *sql.DB) string{
	var version string
	var s_version string
	sql := `SELECT @@VERSION`
	err := db.QueryRow(sql).Scan(&version)
	if err != nil {
		log.Printf("GetVersion failed: %s", err.Error())
	}else{
		s_version = version[21:25]
	}

	return s_version
}

func GetUptime(db *sql.DB) string{
	var uptime string
	sql := `SELECT CONVERT(varchar, sqlserver_start_time, 120) as time_restart FROM sys.dm_os_sys_info`
	err := db.QueryRow(sql).Scan(&uptime)
	if err != nil {
		log.Printf("GetUptime failed: %s", err.Error())
	}
	return uptime
}


func GetVariables(db *sql.DB, matrix_name string) string{
	var matrix_value string
	sql := `select @@` + matrix_name
	err := db.QueryRow(sql).Scan(&matrix_value)
	if err != nil {
		log.Printf("GetVariables for %s failed: %s", matrix_name, err.Error())
	}
	return matrix_value
}

func GetProcesses(db *sql.DB) int{
	var total_processes int
	sql := `SELECT count(*) FROM [master].[dbo].[sysprocesses] WHERE [DBID] > 0`
	err := db.QueryRow(sql).Scan(&total_processes)
	if err != nil {
		log.Printf("GetProcesses failed: %s", err.Error())
		return -1
	}
	return total_processes
}

func GetProcessesRunning(db *sql.DB) int{
	var running_processes int
	sql := `SELECT COUNT(*) FROM [master].[dbo].[sysprocesses] WHERE [DBID] >0 AND status !='SLEEPING' AND status !='BACKGROUND'`
	err := db.QueryRow(sql).Scan(&running_processes)
	if err != nil {
		log.Printf("GetProcessesRunning failed: %s", err.Error())
		return -1
	}
	return running_processes
}

func GetProcessesWaits(db *sql.DB) int{
	var wait_processes int
	sql := `SELECT COUNT(*) FROM [master].[dbo].[sysprocesses] WHERE [DBID] >0 AND status ='SUSPENDED' AND waittime >2`
	err := db.QueryRow(sql).Scan(&wait_processes)
	if err != nil {
		log.Printf("GetProcessesWaits failed: %s", err.Error())
		return -1
	}
	return wait_processes
}