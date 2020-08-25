package com

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)




func ErrorConnection(dbid int, db *sql.DB) error {
	var db_id int
	curr_time := time.Now().Unix()

	err := db.QueryRow("SELECT `id` FROM pms_db_status where `id` = ?", dbid).Scan(&db_id)
	if err == sql.ErrNoRows {
		db.Exec("insert into pms_db_status (`id`, asset_type, alias, connect, created) select `id`, asset_type, alias, -1, ? from pms_asset_config where id = ?", curr_time, dbid)
	} else if err != nil {
		db.Exec("update pms_db_status set connect=-1, role='', version = '', updated = ? where id = ?", curr_time, dbid)
	}

	return err
}