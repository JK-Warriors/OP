package oracle

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

func GetSessionTotal(db *sql.DB) int {
	var count int
	sql := `select count(1) from v$session`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		count = 0
	}
	return count
}

func GetSessionActive(db *sql.DB) int {
	var count int
	sql := `select count(1) from v$session where status = 'ACTIVE'`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		count = 0
	}
	return count
}

func GetSessionWait(db *sql.DB) int {
	var count int
	sql := `select count(1) from v$session where wait_class != 'Idle'`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		count = 0
	}
	return count
}

func GetFlashbackUsage(db *sql.DB) string {
	var flashback_usage string
	sql := `select sum(nvl(percent_space_used,0)) from v$flash_recovery_area_usage`
	err := db.QueryRow(sql).Scan(&flashback_usage)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		flashback_usage = "0"
	}
	return flashback_usage
}

func GetCurrentInstanceNumber(db *sql.DB) int {
	var inst_num int
	sql := `select instance_number from v$instance`
	err := db.QueryRow(sql).Scan(&inst_num)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		return 0
	}
	return inst_num
}

func GetLastSnapId(db *sql.DB) int {
	var snap_id int
	sql := `select max(snap_id) from wrm$_snapshot`
	err := db.QueryRow(sql).Scan(&snap_id)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		return 0
	}
	return snap_id
}

func Get_Instance(db *sql.DB, matrix_name string) string {
	var matrix_value string
	sql := `select ` + matrix_name + ` from v$instance`
	err := db.QueryRow(sql).Scan(&matrix_value)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		return ""
	}
	return matrix_value
}

func Get_Database(db *sql.DB, matrix_name string) string {
	var matrix_value string
	sql := `select ` + matrix_name + ` from v$database`
	err := db.QueryRow(sql).Scan(&matrix_value)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		return ""
	}
	return matrix_value
}

func GetQPS(db *sql.DB) int {
	var count int
	sql := `select value from v$sysmetric where metric_name in ('Executions Per Sec') where group_id=2`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		count = 0
	}
	return count
}

func GetTPS(db *sql.DB) int {
	var count int
	sql := `select (select value
						from v$sysmetric
					where metric_name in ('User Commits Per Sec')
						and group_id = 2) +
					(select value
						from v$sysmetric
					where metric_name in ('User Rollbacks Per Sec')
						and group_id = 2) as TPS
				from dual`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		count = 0
	}
	return count
}

func GetBufferCacheHit(db *sql.DB) int {
	var count int
	sql := `select round(1 - "physical reads" / ("db block gets" + "consistent gets"),2)*100 as hit
			from (SELECT max(decode(name, 'db block gets', value, null)) as "db block gets",
						max(decode(name, 'consistent gets', value, null)) as "consistent gets",
						max(decode(name, 'physical reads', value, null)) as "physical reads"
					FROM v$sysstat
				WHERE name IN ('db block gets', 'consistent gets', 'physical reads'))`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Printf("%s: %w", sql, err)
		count = 0
	}
	return count
}
