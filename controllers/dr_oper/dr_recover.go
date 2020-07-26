package dr_oper

import (
	"log"
	"opms/controllers"
	"opms/lib/exception"
	"strconv"

	. "opms/models/dr_business"
	. "opms/models/dbconfig"
	. "opms/models/dr_oper"
	. "opms/models/users"
	"opms/utils"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/godror/godror"
)

//业务系统管理
type ManageDrRecoverController struct {
	controllers.BaseController
}

func (this *ManageDrRecoverController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-sync") {
		this.Abort("401")
	}

	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	search_name := this.GetString("search_name")
	condArr := make(map[string]string)
	condArr["search_name"] = search_name

	countBs := CountBusiness(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countBs)
	_, _, dr := ListDr(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["dr"] = dr
	this.Data["countBs"] = countBs

	userid, _ := this.GetSession("userId").(int64)
	user, _ := GetUser(userid)
	this.Data["user"] = user

	this.TplName = "dr_oper/recover-index.tpl"
}

type OperDrRecoverController struct {
	controllers.BaseController
}

func (this *OperDrRecoverController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-sync") {
		this.Abort("401")
	}

	idstr := this.Ctx.Input.Param(":id")
	bs_id, _ := strconv.Atoi(idstr)

	//灾备配置检查
	cfg_count, err := CheckDrConfig(bs_id)
	if cfg_count == 0 {
		//没有配置容灾库
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "该系统没有配置容灾库"}
		//this.ServeJSON()
		//return
	}

	// 获取db id
	db_id, err := GetStandbyDBId(bs_id)
	if err != nil {
		utils.LogDebug("GetStandbyDBID failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDBID successfully.")

	// 获取db type
	db_type := GetDBtypeByDBId(db_id)

	dsn, err := GetDsn(db_id, db_type)
	if err != nil {
		utils.LogDebug("GetStandbyDsn failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDsn successfully.")

	p_db, err := godror.ParseConnString(dsn)
	if err != nil {
		utils.LogDebugf("%s: %w", dsn, err)
	}

	restore_point, err := GetRestorePointName(p_db)
	this.Data["restore_point"] = restore_point

	userid, _ := this.GetSession("userId").(int64)
	user, _ := GetUser(userid)
	this.Data["user"] = user

	this.Data["bs_id"] = bs_id
	this.Data["db_id"] = db_id

	this.TplName = "dr_oper/recover-oper.tpl"
}

type AjaxDrFlashbackController struct {
	controllers.BaseController
}

func (this *AjaxDrFlashbackController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }
	bs_id, _ := this.GetInt("bs_id")
	db_id, _ := this.GetInt("db_id")
	fb_method, _ := this.GetInt("fb_method")
	fb_point := this.GetString("fb_point")
	fb_time := this.GetString("fb_time")
	op_type := "STARTFLASHBACK"

	//灾备配置检查
	cfg_count, err := CheckDrConfig(bs_id)
	if cfg_count == 0 {
		//没有配置容灾库
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "该系统没有配置容灾库"}
		this.ServeJSON()
		return
	}

	// 获取db type
	db_type := GetDBtypeByDBId(db_id)

	// Get dsn
	dsn_s, err := GetDsn(db_id, db_type)
	if err != nil {
		utils.LogDebug("GetDsn failed: " + err.Error())
	}
	utils.LogDebug("GetDsn successfully.")

	on_process, err := GetOnProcess(bs_id)
	if on_process == 1 {
		utils.LogDebug("There is another opration on process.")

		this.Data["json"] = map[string]interface{}{"code": 0, "message": "有另外一个操作正在进行中"}
		this.ServeJSON()
		return
	} else {
		exception.Try(func() {

			utils.LogDebug("操作加锁")
			OperationLock(bs_id, op_type)

			utils.LogDebug("初始化切换实例")
			op_id := utils.SnowFlakeId()
			Init_OP_Instance(op_id, bs_id, db_type, op_type)
			//db_type = 5
			utils.LogDebug("正式开始恢复快照任务")
			if db_type == 1 { //oracle
				p_sta, err := godror.ParseConnString(dsn_s)
				if err != nil {
					utils.LogDebugf("%s: %w", dsn_s, err)
				}

				utils.LogDebug("开始恢复快照...")
				s_result := OraStartFlashback(op_id, bs_id, fb_method, fb_point, fb_time, p_sta)
				utils.LogDebug("恢复快照结束")

				if s_result == 1 {
					utils.LogDebug("更新恢复结果")
					Update_OP_Result(op_id, 1)
				} else {
					utils.LogDebug("备库恢复快照失败，更新恢复结果")
					Update_OP_Result(op_id, -1)
				}

				OperationUnlock(bs_id, op_type)
			} else if db_type == 2 { //mysql
				//OraPrimaryToStandby
				//OraStandbyToPrimary

			} else if db_type == 3 { //sqlserver
				//OraPrimaryToStandby
				//OraStandbyToPrimary

			}

		}).Catch(1, func(e exception.Exception) {
			log.Println(e.Id, e.Msg)
		}).Catch(2, func(e exception.Exception) {
			log.Println(e.Id, e.Msg)
		}).Catch(-1, func(e exception.Exception) {
			log.Println(e.Id, e.Msg)
		}).Finally(func() {
			OperationUnlock(bs_id, op_type)
		})

		this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作完成"}
		this.ServeJSON()
	}
}

type AjaxDrRecoverController struct {
	controllers.BaseController
}

func (this *AjaxDrRecoverController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }
	bs_id, _ := this.GetInt("bs_id")
	db_type, _ := this.GetInt("db_type")
	op_type := "STOPFLASHBACK"

	//灾备配置检查
	cfg_count, err := CheckDrConfig(bs_id)
	if cfg_count == 0 {
		//没有配置容灾库
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "该系统没有配置容灾库"}
		this.ServeJSON()
		return
	}

	sta_id, err := GetStandbyDBId(bs_id)
	if err != nil {
		utils.LogDebug("GetStandbyDBID failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDBID successfully.")

	dsn_s, err := GetDsn(sta_id, db_type)
	if err != nil {
		utils.LogDebug("GetStandbyDsn failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDsn successfully.")

	on_process, err := GetOnProcess(bs_id)
	if on_process == 1 {
		utils.LogDebug("There is another opration on process.")

		this.Data["json"] = map[string]interface{}{"code": 0, "message": "有另外一个操作正在进行中"}
		this.ServeJSON()
		return
	} else {
		exception.Try(func() {

			utils.LogDebug("操作加锁")
			OperationLock(bs_id, op_type)

			utils.LogDebug("初始化切换实例")
			op_id := utils.SnowFlakeId()
			Init_OP_Instance(op_id, bs_id, db_type, op_type)
			//db_type = 5
			utils.LogDebug("正式开始恢复同步任务")
			if db_type == 1 { //oracle
				p_sta, err := godror.ParseConnString(dsn_s)
				if err != nil {
					utils.LogDebugf("%s: %w", dsn_s, err)
				}

				utils.LogDebug("正在恢复同步...")
				s_result := OraRecover(op_id, bs_id, p_sta)
				utils.LogDebug("恢复同步结束")

				if s_result == 1 {
					utils.LogDebug("更新恢复结果")
					Update_OP_Result(op_id, 1)
				} else {
					utils.LogDebug("备库恢复同步失败，更新恢复结果")
					Update_OP_Result(op_id, -1)
				}

				OperationUnlock(bs_id, op_type)
			} else if db_type == 2 { //mysql
				//OraPrimaryToStandby
				//OraStandbyToPrimary

			} else if db_type == 3 { //sqlserver
				//OraPrimaryToStandby
				//OraStandbyToPrimary

			}

		}).Catch(1, func(e exception.Exception) {
			log.Println(e.Id, e.Msg)
		}).Catch(2, func(e exception.Exception) {
			log.Println(e.Id, e.Msg)
		}).Catch(-1, func(e exception.Exception) {
			log.Println(e.Id, e.Msg)
		}).Finally(func() {
			OperationUnlock(bs_id, op_type)
		})

		this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作完成"}
		this.ServeJSON()
	}

}
