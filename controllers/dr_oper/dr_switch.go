package dr_oper

import (
	"encoding/json"
	"log"
	"opms/controllers"
	"opms/lib/exception"
	"strconv"

	. "opms/models/dr_oper"
	. "opms/models/users"
	"opms/utils"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/godror/godror"
)

//业务切换管理
type ManageDrSwitchController struct {
	controllers.BaseController
}

func (this *ManageDrSwitchController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch") {
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

	countDr := CountDrconfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDr)
	_, _, dr := ListDr(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["dr"] = dr
	this.Data["countDr"] = countDr

	userid, _ := this.GetSession("userId").(int64)
	user, _ := GetUser(userid)
	this.Data["user"] = user

	this.TplName = "dr_oper/switch-index.tpl"
}

//业务大屏
type ScreenDrSwitchController struct {
	controllers.BaseController
}

func (this *ScreenDrSwitchController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
		this.Abort("401")
	}

	userid, _ := this.GetSession("userId").(int64)
	user, _ := GetUser(userid)
	this.Data["user"] = user

	idstr := this.Ctx.Input.Param(":id")
	bs_id, _ := strconv.Atoi(idstr)
	//灾备配置检查
	cfg_count, err := CheckDrConfig(bs_id)
	if cfg_count == 0 {
		//没有配置容灾库
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "该系统没有配置容灾库"}
		this.ServeJSON()
		return
	}

	pri_id, err := GetPrimaryDBId(bs_id)
	if err != nil {
		utils.LogDebug("GetPrimaryDBID failed: " + err.Error())
	}
	utils.LogDebug("GetPrimaryDBID successfully.")
	sta_id, err := GetStandbyDBId(bs_id)
	if err != nil {
		utils.LogDebug("GetStandbyDBID failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDBID successfully.")

	pri_basic, err := GetOracleBasicInfo(pri_id)
	sta_basic, err := GetOracleBasicInfo(sta_id)
	pri_dr, err := GetPrimaryDrInfo(pri_id)
	sta_dr, err := GetStandbyDrInfo(sta_id)

	this.Data["pri_basic"] = pri_basic
	this.Data["sta_basic"] = sta_basic
	this.Data["pri_dr"] = pri_dr
	this.Data["sta_dr"] = sta_dr

	this.TplName = "dr_oper/screen-oracle.tpl"
}

type AjaxDrSwitchoverController struct {
	controllers.BaseController
}

func (this *AjaxDrSwitchoverController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }

	bs_id, _ := this.GetInt("bs_id")
	op_type := this.GetString("op_type")

	utils.LogDebug(bs_id)
	utils.LogDebug(op_type)

	//灾备配置检查
	cfg_count, err := CheckDrConfig(bs_id)
	if cfg_count == 0 {
		//没有配置容灾库
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "该系统没有配置容灾库"}
		this.ServeJSON()
		return
	}
	asset_type := GetAssetType(bs_id)

	pri_id, err := GetPrimaryDBId(bs_id)
	if err != nil {
		utils.LogDebug("GetPrimaryDBID failed: " + err.Error())
	}
	utils.LogDebug("GetPrimaryDBID successfully.")
	sta_id, err := GetStandbyDBId(bs_id)
	if err != nil {
		utils.LogDebug("GetStandbyDBID failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDBID successfully.")

	// Get dsn
	dsn_p, err := GetDsn(pri_id, asset_type)
	if err != nil {
		utils.LogDebug("GetPrimaryDsn failed: " + err.Error())
	}
	utils.LogDebug("GetPrimaryDsn successfully.")

	dsn_s, err := GetDsn(sta_id, asset_type)
	if err != nil {
		utils.LogDebug("GetStandbyDsn failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDsn successfully.")

	utils.LogDebug(dsn_p)
	utils.LogDebug(dsn_s)

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
			Init_OP_Instance(op_id, bs_id, asset_type, op_type)
			//asset_type = 5
			utils.LogDebug("正式开始切换")
			if asset_type == 1 { //oracle
				p_pri, err := godror.ParseConnString(dsn_p)
				if err != nil {
					utils.LogDebugf("%s: %w", dsn_p, err)
				}

				p_sta, err := godror.ParseConnString(dsn_s)
				if err != nil {
					utils.LogDebugf("%s: %w", dsn_s, err)
				}

				utils.LogDebug("主库开始切换成备库...")
				p_result := OraPrimaryToStandby(op_id, bs_id, p_pri)
				utils.LogDebug("主库切换成备库结束")
				if p_result == 1 {
					utils.LogDebug("备库开始切换成主库...")
					s_result := OraStandbyToPrimary(op_id, bs_id, p_sta)
					utils.LogDebug("备库切换成主库结束")
					if s_result == 1 {
						utils.LogDebug("更新切换标识")
						UpdateSwitchFlag(bs_id)
						utils.LogDebug("更新切换结束")
						Update_OP_Result(op_id, 1)
					} else {
						utils.LogDebug("备库切换主库失败，更新切换结果")
						Update_OP_Result(op_id, -1)
					}
				} else {
					utils.LogDebug("主库切换备库失败，更新切换结果")
					Update_OP_Result(op_id, -1)
				}
				OperationUnlock(bs_id, op_type)
			} else if asset_type == 2 { //mysql
				utils.LogDebug("从库开始切换成主库...")
				p_result := SlaveToMaster(op_id, bs_id, dsn_p, dsn_s, sta_id)
				if p_result == 0 {
					utils.LogDebug("从库切换成主库成功")
					utils.LogDebug("开始重建复制关系...")
					s_result := RebuildReplication(bs_id, dsn_p, dsn_s, sta_id)
					if s_result == 0 {
						utils.LogDebug("重建复制关系成功")
						utils.LogDebug("更新切换标识")
						UpdateSwitchFlag(bs_id)
						utils.LogDebug("更新切换标识结束")
						Update_OP_Result(op_id, 1)
					} else {
						utils.LogDebug("重建复制关系失败，更新切换结果")
						Update_OP_Result(op_id, -1)
					}
				} else {
					utils.LogDebug("从库切换成主库失败，更新切换结果")
					Update_OP_Result(op_id, -1)
				}

				OperationUnlock(bs_id, op_type)
			} else if asset_type == 3 { //sqlserver
				utils.LogDebug("获取镜像库名称...")
				db_name, err := GetMirrorDbname(bs_id)
				if nil == err {
					utils.LogDebugf("镜像库名为：%s", db_name)
					utils.LogDebug("开始切换镜像库...")
					result := SwitchMirror(op_id, bs_id, dsn_p, dsn_s, db_name)
					if result == 0 {
						utils.LogDebug("镜像库切换成功")
						utils.LogDebug("更新切换标识")
						UpdateSwitchFlag(bs_id)
						utils.LogDebug("更新切换标识结束")
						Update_OP_Result(op_id, 1)
					} else {
						utils.LogDebug("镜像库切换失败，更新切换结果")
						Update_OP_Result(op_id, -1)
					}
				}else{
					utils.LogDebug("获取镜像库名称失败")
					Update_OP_Result(op_id, -1)
				}

				OperationUnlock(bs_id, op_type)
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

type AjaxDrFailoverController struct {
	controllers.BaseController
}

func (this *AjaxDrFailoverController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }

	bs_id, _ := this.GetInt("bs_id")
	op_type := this.GetString("op_type")

	utils.LogDebug(bs_id)
	utils.LogDebug(op_type)

	//灾备配置检查
	cfg_count, err := CheckDrConfig(bs_id)
	if cfg_count == 0 {
		//没有配置容灾库
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "该系统没有配置容灾库"}
		this.ServeJSON()
		return
	}
	asset_type := GetAssetType(bs_id)

	sta_id, err := GetStandbyDBId(bs_id)
	if err != nil {
		utils.LogDebug("GetStandbyDBID failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDBID successfully.")

	// Get dsn
	dsn_s, err := GetDsn(sta_id, asset_type)
	if err != nil {
		utils.LogDebug("GetStandbyDsn failed: " + err.Error())
	}
	utils.LogDebug("GetStandbyDsn successfully.")
	//utils.LogDebug(dsn_s)

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
			Init_OP_Instance(op_id, bs_id, asset_type, op_type)
			//asset_type = 5
			utils.LogDebug("正式开始灾难切换")
			if asset_type == 1 { //oracle

				p_sta, err := godror.ParseConnString(dsn_s)
				if err != nil {
					utils.LogDebugf("%s: %w", dsn_s, err)
				}

				utils.LogDebug("备库开始切换成主库...")
				s_result := OraFailoverToPrimary(op_id, bs_id, p_sta)
				utils.LogDebug("备库切换成主库结束")
				if s_result == 1 {
					utils.LogDebug("更新切换标识")
					UpdateSwitchFlag(bs_id)
					utils.LogDebug("更新切换结束")
					Update_OP_Result(op_id, 1)
				} else {
					utils.LogDebug("备库切换主库失败，更新切换结果")
					Update_OP_Result(op_id, -1)
				}

				OperationUnlock(bs_id, op_type)
			} else if asset_type == 2 { //mysql
				utils.LogDebug("从库开始切换成主库...")
				p_result := FailoverToMaster(op_id, bs_id, dsn_s, sta_id)
				if p_result == 0 {
					utils.LogDebug("从库切换成主库成功")
					utils.LogDebug("更新切换标识")
					UpdateSwitchFlag(bs_id)
					utils.LogDebug("更新切换标识结束")
					Update_OP_Result(op_id, 1)
				} else {
					utils.LogDebug("从库切换成主库失败，更新切换结果")
					Update_OP_Result(op_id, -1)
				}

				OperationUnlock(bs_id, op_type)

			} else if asset_type == 3 { //sqlserver
				//Skip

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

type AjaxDrProcessController struct {
	controllers.BaseController
}

func (this *AjaxDrProcessController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }
	//utils.LogDebug("GetDrProcess begin...")
	bs_id, _ := this.GetInt("bs_id")
	op_type := this.GetString("op_type")

	cfg_count, err := CheckDrConfig(bs_id)
	if cfg_count == 0 {
		//没有配置容灾库
		this.Data["json"] = map[string]interface{}{"on_process": -1, "op_type": op_type, "op_result": -1, "op_reason": -1, "json_process": "null"}
		this.ServeJSON()
		return
	}

	on_process, err := GetOnProcess(bs_id)

	//op_type, err := GetCurrentOpType(bs_id)

	op_id, err := GetCurrentOpId(bs_id, op_type)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"on_process": on_process, "op_type": op_type, "op_result": -1, "op_reason": "", "json_process": "null"}
		this.ServeJSON()
		return
	}

	op_result, op_reason, err := GetOpResultById(op_id)
	if err != nil {
		op_result = "-1"
		op_reason = ""
	}

	pro, err := GetOPProcessById(op_id)
	if err != nil {
		utils.LogDebug("获取Process详细步骤失败")
	}

	json_pro, err := json.Marshal(pro)
	if err != nil {
		utils.LogDebug("生成json字符串错误")
	}

	utils.LogDebug(json_pro)

	this.Data["json"] = map[string]interface{}{"on_process": on_process, "op_type": op_type, "op_result": op_result, "op_reason": op_reason, "json_process": string(json_pro)}
	this.ServeJSON()

}
