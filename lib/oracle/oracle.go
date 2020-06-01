package oracle

import (
	"context"
	"database/sql"
	"log"
	"opms/utils"
	"strconv"
	"time"

	"github.com/godror/godror"
	errors "golang.org/x/xerrors"
)

// Startup database mount.
func StartupMount(P godror.ConnectionParams) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if !(P.IsSysDBA || P.IsSysOper) {
		P.IsSysDBA = true
	}
	if !P.IsPrelim {
		P.IsPrelim = true
	}

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		log.Fatal(errors.Errorf("%s: %w", P.StringWithPassword(), err))
	}
	defer db.Close()

	err = godror.Raw(ctx, db, func(oraDB godror.Conn) error {
		utils.LogDebug("Starting database nomount")
		if err = oraDB.Startup(godror.StartupDefault); err != nil {
			return err
		}
		return nil
	})

	// You cannot alter database on the prelim_auth connection.
	// So open a new connection and complete startup, as Startup starts pmon.
	if P.IsPrelim {
		P.IsPrelim = false
	}
	db2, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		return err
	}
	defer db2.Close()

	utils.LogDebug("Mounting database")
	if _, err = db2.Exec("alter database mount"); err != nil {
		return err
	}
	return nil
}

func StartupOpen(P godror.ConnectionParams) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if !(P.IsSysDBA || P.IsSysOper) {
		P.IsSysDBA = true
	}
	if !P.IsPrelim {
		P.IsPrelim = true
	}

	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebug(errors.Errorf("%s: %w", P.StringWithPassword(), err))
	}
	defer db.Close()

	err = godror.Raw(ctx, db, func(oraDB godror.Conn) error {
		utils.LogDebug("Starting database nomount")
		if err = oraDB.Startup(godror.StartupDefault); err != nil {
			return err
		}
		return nil
	})

	// You cannot alter database on the prelim_auth connection.
	// So open a new connection and complete startup, as Startup starts pmon.
	if P.IsPrelim {
		P.IsPrelim = false
	}
	db2, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebug("Open a new connection: " + err.Error())
		return err
	}
	defer db2.Close()

	utils.LogDebug("Mounting database")
	if _, err = db2.Exec("alter database mount"); err != nil {
		utils.LogDebug("alter database mount: " + err.Error())
		return err
	}
	utils.LogDebug("Opening database")
	if _, err = db2.Exec("alter database open"); err != nil {
		utils.LogDebug("alter database open: " + err.Error())
		return err
	}
	return nil
}

// ShutdownMode calls Shutdown to shut down a database.
func ShutdownMode(P godror.ConnectionParams, shutdownMode godror.ShutdownMode) {
	//dsn := "oracle://?sysdba=1" // equivalent to "/ as sysdba"
	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebug(errors.Errorf("%s: %w", P.StringWithPassword(), err))
	}
	defer db.Close()

	if err = Shutdown(db, shutdownMode); err != nil {
		utils.LogDebug(err)
	}
}

func ShutdownImmediate(P godror.ConnectionParams) {
	//dsn := "oracle://?sysdba=1" // equivalent to "/ as sysdba"
	db, err := sql.Open("godror", P.StringWithPassword())
	if err != nil {
		utils.LogDebug(errors.Errorf("%s: %w", P.StringWithPassword(), err))
	}
	defer db.Close()

	if err = Shutdown(db, godror.ShutdownImmediate); err != nil {
		utils.LogDebug(err)
	}
}

func Shutdown(db *sql.DB, shutdownMode godror.ShutdownMode) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := godror.Raw(ctx, db, func(oraDB godror.Conn) error {
		utils.LogDebug(errors.Errorf("Beginning shutdown %v", shutdownMode))
		return oraDB.Shutdown(shutdownMode)
	})
	if err != nil {
		utils.LogDebug(err.Error())
	}
	// If we abort the shutdown process is over immediately.
	if shutdownMode == godror.ShutdownAbort {
		return nil
	}

	utils.LogDebug("Closing database")
	if _, err = db.Exec("alter database close normal"); err != nil {
		utils.LogDebug("alter database close normal: " + err.Error())
	}
	utils.LogDebug("Unmounting database")
	if _, err = db.Exec("alter database dismount"); err != nil {
		utils.LogDebug("alter database dismount: " + err.Error())
	}
	utils.LogDebug("Finishing shutdown")
	return godror.Raw(ctx, db, func(oraDB godror.Conn) error {
		return oraDB.Shutdown(godror.ShutdownFinal)
	})
}

func Context(name string) context.Context {
	return godror.ContextWithTraceTag(context.Background(), godror.TraceTag{Module: "Test" + name})
}

//Get open_mode
func GetOpenMode(db *sql.DB) (string, error) {
	var open_mode string
	var err error
	selectQry := "select open_mode from v$database"
	open_mode, err = GetSingleValue(db, selectQry)
	if err != nil {
		utils.LogDebug("Get open mode failed: " + err.Error())
		return "", err
	}
	return open_mode, nil
}

func GetDatabaseRole(db *sql.DB) (string, error) {
	var db_role string
	var err error
	selectQry := "select database_role from v$database"
	db_role, err = GetSingleValue(db, selectQry)
	if err != nil {
		utils.LogDebug("Get database_role failed: " + err.Error())
		return "", err
	}
	return db_role, nil
}

func GetSwitchoverStatus(db *sql.DB) (string, error) {
	var switch_status string
	var err error
	selectQry := "select switchover_status from v$database"
	switch_status, err = GetSingleValue(db, selectQry)
	if err != nil {
		utils.LogDebug("Get switchover_status failed: " + err.Error())
		return "", err
	}
	return switch_status, nil
}

func GetDatabaseVersion(db *sql.DB) (int, error) {
	var db_version string
	var err error
	selectQry := "select substr(version, 0, instr(version, '.')-1) from v$instance"
	db_version, err = GetSingleValue(db, selectQry)
	if err != nil {
		utils.LogDebug("Get database version failed: " + err.Error())
		return -1, err
	}

	version, err := strconv.Atoi(db_version)
	if err != nil {
		return -1, err
	}

	return version, nil
}

func GetGapCount(db *sql.DB) (int, error) {
	var count string
	var err error
	selectQry := "select count(1) from v$archive_gap"
	count, err = GetSingleValue(db, selectQry)
	if err != nil {
		utils.LogDebug("Get gap count failed: " + err.Error())
		return -1, err
	}

	gap_count, err := strconv.Atoi(count)
	if err != nil {
		return -1, err
	}

	return gap_count, nil
}

func GetSyncStatus(db *sql.DB) (int, error) {
	var count string
	var err error
	selectQry := "select count(1) from gv$session where program like '%(MRP0)' "
	count, err = GetSingleValue(db, selectQry)
	if err != nil {
		utils.LogDebug("Get Sync status failed: " + err.Error())
		return -1, err
	}

	sync_status, err := strconv.Atoi(count)

	if err != nil {
		utils.LogDebug("Get Sync status failed: " + err.Error())
		return -1, err
	}
	return sync_status, nil
}

func GetstandbyRedoLog(db *sql.DB) (int, error) {
	var count string
	var err error
	selectQry := "select count(1) from v$standby_log "
	count, err = GetSingleValue(db, selectQry)
	if err != nil {
		utils.LogDebug("Get standby redo log failed: " + err.Error())
		return -1, err
	}

	sta_redo_count, err := strconv.Atoi(count)

	if err != nil {
		utils.LogDebug("Get standby redo log failed: " + err.Error())
		return -1, err
	}
	return sta_redo_count, nil
}

func GetSingleValue(db *sql.DB, sql string) (string, error) {
	var single_value string

	ctx, cancel := context.WithTimeout(Context("QueryTimeout"), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, sql)

	for rows.Next() {
		if err = rows.Scan(&single_value); err != nil {
			return "", err
		}
	}
	if err = rows.Err(); err != nil {
		return "", err
	}

	return single_value, nil
}
