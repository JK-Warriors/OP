package dr_oper

import (
	"database/sql"
	"fmt"
	"opms/utils"

	_ "github.com/denisenkom/go-mssqldb"
)

func SwitchMirror(op_id int64, bs_id int, dsn_p string, dsn_s string, db_name string) int {
	result := -1
	
	//建立主体连接
	mdb, err := sql.Open("mssql", dsn_p)
	if err != nil {
		Update_OP_Reason(op_id, "连接主体库失败")
		utils.LogDebugf("Open SQLServer failed: %s", err.Error())
		return -1
	}
	defer mdb.Close()

	err = mdb.Ping()
	if err != nil {
		Update_OP_Reason(op_id, "Ping主体库失败")
		utils.LogDebugf("Ping SQLServer failed: %s", err.Error())
		return -1
	}

	//建立镜像连接
	sdb, err := sql.Open("mssql", dsn_s)
	if err != nil {
		Update_OP_Reason(op_id, "连接镜像库失败")
		utils.LogDebugf("Open SQLServer failed: %s", err.Error())
		return -1
	}
	defer sdb.Close()

	err = sdb.Ping()
	if err != nil {
		Update_OP_Reason(op_id, "Ping镜像库失败")
		utils.LogDebugf("Ping SQLServer failed: %s", err.Error())
		return -1
	}

	//get database role, 1: 主体; 2:镜像; NULL:数据库不可访问或未镜像
	role, err := GetMirrorRole(mdb, db_name)
	if err != nil {
		Update_OP_Reason(op_id, "获取镜像角色失败")
		utils.LogDebugf("Get mirror role failed: %s", err.Error())
		return -1
	}
	Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "获取镜像角色成功")
	utils.LogDebugf("The current mirror role is: %d", role)

	state, err := GetMirrorState(mdb, db_name)
	if err != nil {
		Update_OP_Reason(op_id, "获取镜像状态失败")
		utils.LogDebugf("Get mirror state failed: %s", err.Error())
		return -1
	}
	Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "获取镜像状态成功")
	utils.LogDebugf("The current mirror state is: %d", state)

	if role == 1 {
		Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "验证镜像角色成功")
		utils.LogDebug("Now we are going to switch mirror.")
		if state == 4 {
			//设置镜像传输模式为高安全模式
			utils.LogDebug("SET SAFETY FULL...")
			sql := fmt.Sprintf(`ALTER DATABASE %s SET SAFETY FULL`, db_name)
			_, err = mdb.Exec(sql)
			if err != nil {
				Update_OP_Reason(op_id, "切换成高安全模式失败")
				utils.LogDebugf("Alter database set safety full failed: %s", err.Error())
				return -1
			}else{
				Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "主库已经成功切换成高安全模式")
			}
    
			utils.LogDebug("SET PARTNER FAILOVER begin...")
            //切换镜像
			sql = fmt.Sprintf(`ALTER DATABASE %s SET PARTNER FAILOVER`, db_name)
			_, err = mdb.Exec(sql)
			if err != nil {
				Update_OP_Reason(op_id, "切换镜像失败")
				utils.LogDebugf("Alter database set partner failover failed: %s", err.Error())
				return -1
			}

			role, err := GetMirrorRole(mdb, db_name)
			if err != nil {
				Update_OP_Reason(op_id, "获取切换后镜像角色失败")
				utils.LogDebugf("Get mirror role after switch failed: %s", err.Error())
				return -1
			}

			if role == 2{
				Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "主体已经成功切换成镜像")
				utils.LogDebug("SET PARTNER FAILOVER successfully.")
				
                //设置镜像传输模式为高性能模式
				sql = fmt.Sprintf(`ALTER DATABASE %s SET SAFETY OFF`, db_name)
				_, err = sdb.Exec(sql)
				if err != nil {
					Update_OP_Reason(op_id, "切换高性能模式失败")
					utils.LogDebugf("Alter database set safety off failed: %s", err.Error())
					return -1
				}
        
				Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "镜像已经成功切换成高性能模式")
				utils.LogDebug("SET SAFETY OFF successfully.")
                
                result=0
			}else{
				Update_OP_Reason(op_id, "校验切换后镜像角色失败")
				Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "主体切换成镜像失败")
				utils.LogDebug("SET PARTNER FAILOVER failed.")
                result=-1
			}
		}else{
			Update_OP_Reason(op_id, "验证镜像状态失败，当前镜像状态不是'已同步'")
			Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "验证镜像状态失败，当前镜像状态不是'已同步'")
			utils.LogDebug("Check mirror state failed!")
			result = -1
		}
	}else {
		Update_OP_Reason(op_id, "验证镜像角色失败，当前镜像不是主体，不能切换镜像")
		Log_OP_Process(op_id, bs_id, 3, "SWITCHOVER", "验证镜像角色失败，当前镜像不是主体，不能切换镜像")
		utils.LogDebug("Check mirror role failed!")
		result = -1
	}

	return result
}


func GetMirrorRole(db *sql.DB, db_name string) (int, error){
	var role int
	sql := `select m.mirroring_role
			from sys.database_mirroring m, sys.databases d
			where M.mirroring_guid is NOT NULL
				AND m.database_id = d.database_id
				AND d.name = ?`
	err := db.QueryRow(sql, db_name).Scan(&role)
	if err != nil {
		utils.LogDebugf("GetMirrorRole failed: %s", err.Error())
		return -1, err
	}

	return role, nil
}

func GetMirrorState(db *sql.DB, db_name string) (int, error){
	var state int =-1
	// 0:已挂起; 1:与其他伙伴断开; 2:正在同步; 3:挂起故障转移; 4:已同步; 5:伙伴未同步; 6:伙伴已同步;
	sql := `select m.mirroring_state
			from sys.database_mirroring m, sys.databases d
			where M.mirroring_guid is NOT NULL
				AND m.database_id = d.database_id
				AND d.name = ?`
	err := db.QueryRow(sql, db_name).Scan(&state)
	if err != nil {
		utils.LogDebugf("GetMirrorState failed: %s", err.Error())
		return -1, err
	}

	return state, nil
}