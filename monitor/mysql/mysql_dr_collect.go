package mysql

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

func GenerateMySQLDrStats(wg *sync.WaitGroup, mysql *xorm.Engine, dis common.Dr) {
	dr_id := dis.Id

	var pri_id int
	var sta_id int
	if dis.Is_Shift == 0 {
		pri_id = dis.Db_Id_P
		sta_id = dis.Db_Id_S
	} else {
		pri_id = dis.Db_Id_S
		sta_id = dis.Db_Id_P
	}

	dsn_p, err := GetDsn(mysql, pri_id, 2)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	dsn_s, err := GetDsn(mysql, sta_id, 2)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	GenerateMaster(mysql, dr_id, pri_id, dsn_p)
	GenerateSlave(mysql, dr_id, sta_id, dsn_s)
	

	log.Println("获取MySQL容灾数据结束！")

	(*wg).Done()
}


func GenerateMaster(mysql *xorm.Engine, dr_id int, db_id int, connStr string) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		utils.LogDebugf("%s: %s", connStr, err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		utils.LogDebugf("DB Ping failed: %s", err.Error())
		return 
	}

	read_only,_ := GetGlobalVariable(db, "read_only")
	gtid_mode,_ := GetGlobalVariable(db, "gtid_mode")
	binlog_space, _ := GetMasterLogSpace(db)

	file, position, err := GetMasterStatus(db)
	if err != nil {
		log.Printf("GetMasterStatus error: %s", err.Error())
		return
	}else{
		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err := session.Begin()
		//move old data to history table
		MoveToHistory(mysql, "pms_dr_mysql_p", "dr_id", dr_id)

		sql := `insert into pms_dr_mysql_p(dr_id, db_id, read_only, gtid_mode, master_binlog_file, master_binlog_pos, master_binlog_space, created) 
						values(?,?,?,?,?,?,?,?)`

		_, err = mysql.Exec(sql, dr_id, db_id, read_only, gtid_mode, file, position, binlog_space, time.Now().Unix())

		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			session.Rollback()
			return
		}
		// add Commit() after all actions
		err = session.Commit()
	}

	

}


func GenerateSlave(mysql *xorm.Engine, dr_id int, db_id int, connStr string) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		utils.LogDebugf("%s: %s", connStr, err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		utils.LogDebugf("DB Ping failed: %s", err.Error())
		return 
	}
	
	read_only,_ := GetGlobalVariable(db, "read_only")
	gtid_mode,_ := GetGlobalVariable(db, "gtid_mode")

	var a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15,a16,a17,a18,a19,a20,a21,a22,a23,a24,a25,a26,a27,a28,a29,a30 sql.NullString
	var a31,a32,a33,a34,a35,a36,a37,a38,a39,a40,a41,a42,a43,a44,a45,a46,a47,a48,a49,a50,a51,a52,a53,a54,a55,a56,a57 sql.NullString
	sql := `show slave status`
	err = db.QueryRow(sql).Scan(&a1,&a2,&a3,&a4,&a5,&a6,&a7,&a8,&a9,&a10,&a11,&a12,&a13,&a14,&a15,&a16,&a17,&a18,&a19,&a20,&a21,&a22,&a23,&a24,&a25,&a26,&a27,&a28,&a29,&a30,
								&a31,&a32,&a33,&a34,&a35,&a36,&a37,&a38,&a39,&a40,&a41,&a42,&a43,&a44,&a45,&a46,&a47,&a48,&a49,&a50,&a51,&a52,&a53,&a54,&a55,&a56,&a57)
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
		MoveToHistory(mysql, "pms_dr_mysql_s", "dr_id", dr_id)

		sql = `insert into pms_dr_mysql_s(dr_id, db_id, read_only, gtid_mode, master_server, master_port, slave_io_run, slave_sql_run, delay, current_binlog_file, current_binlog_pos, master_binlog_file, master_binlog_pos, master_binlog_space, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

		_, err = mysql.Exec(sql, dr_id, db_id, read_only, gtid_mode, a2.String, a4.String, a11.String, a12.String, a33.String, a10.String, a22.String, a6.String, a7.String, 0, time.Now().Unix())

		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			session.Rollback()
			return
		}
		// add Commit() after all actions
		err = session.Commit()
	}

}
