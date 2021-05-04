package dr_oper

import (
	"opms/controllers"
	"strconv"
	"fmt"

	. "opms/models/dr_oper"
	. "opms/models/users"
	. "opms/models/dr_config"
	. "opms/models/os"
	. "opms/models/oracle"
	"opms/utils"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//业务切换管理
type ListDrController struct {
	controllers.BaseController
}

func (this *ListDrController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-manage-list") {
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

	this.TplName = "dr_oper/manage_list.tpl"
}

//容灾详细
type DetailDrController struct {
	controllers.BaseController
}

func (this *DetailDrController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-manage-detail") {
		this.Abort("401")
	}

	userid, _ := this.GetSession("userId").(int64)
	user, _ := GetUser(userid)
	this.Data["user"] = user
	
	idstr := this.Ctx.Input.Param(":id")
	dr_id, _ := strconv.Atoi(idstr)

	asset_type :=GetAssetType(dr_id)

	pri_id, err := GetPrimaryDBId(dr_id)
	if err != nil {
		utils.LogDebug("GetPrimaryID failed: " + err.Error())
	}
	utils.LogDebug(fmt.Sprintf("GetPrimaryID %d successfully.", pri_id))
	sta_id, err := GetStandbyDBId(dr_id)
	if err != nil {
		utils.LogDebug("GetStandbyID failed: " + err.Error())
	}
	utils.LogDebug(fmt.Sprintf("GetStandbyID %d successfully.", sta_id))


	if(asset_type == 1){
		// type is oracle
		pri_config, err := GetOracleBasicInfo(pri_id)
		if err != nil {
			utils.LogDebug("GetOracleBasicInfo failed: " + err.Error())
		}

		sta_config, err := GetOracleBasicInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetOracleBasicInfo failed: " + err.Error())
		}

		pri_dr, err := GetPrimaryDrInfo(pri_id)
		if err != nil {
			utils.LogDebug("GetPrimaryDrInfo failed: " + err.Error())
		}

		sta_dr, err := GetStandbyDrInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetStandbyDrInfo failed: " + err.Error())
		}

		this.Data["pri_config"] = pri_config
		this.Data["sta_config"] = sta_config
		this.Data["pri_dr"] = pri_dr
		this.Data["sta_dr"] = sta_dr
		this.Data["dr_id"] = dr_id
		this.Data["asset_type"] = asset_type

		this.TplName = "dr_oper/manage_detail_oracle.tpl"
	}else if(asset_type == 2){
		// type is mysql
		pri_config, err := GetMySQLBasicInfo(pri_id)
		if err != nil {
			utils.LogDebug("GetMySQLBasicInfo failed: " + err.Error())
		}

		sta_config, err := GetMySQLBasicInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetMySQLBasicInfo failed: " + err.Error())
		}

		pri_dr, err := GetDrMySqlPrimaryInfo(pri_id)
		if err != nil {
			utils.LogDebug("GetMySQLPrimaryInfo failed: " + err.Error())
		}

		sta_dr, err := GetDrMySqlStandbyInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetMySQLStandbyInfo failed: " + err.Error())
		}

		this.Data["pri_config"] = pri_config
		this.Data["sta_config"] = sta_config
		this.Data["pri_dr"] = pri_dr
		this.Data["sta_dr"] = sta_dr
		this.Data["dr_id"] = dr_id
		this.Data["asset_type"] = asset_type
		

		this.TplName = "dr_oper/manage_detail_mysql.tpl"

	}else if(asset_type == 3){
		// type is sqlserver
		pri_config, err := GetMSSqlBasicInfo(pri_id)
		if err != nil {
			utils.LogDebug("GetMSSqlBasicInfo failed: " + err.Error())
		}

		sta_config, err := GetMSSqlBasicInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetMSSqlBasicInfo failed: " + err.Error())
		}

		pri_dr, err := GetDrMSSqlPrimaryInfo(pri_id)
		if err != nil {
			utils.LogDebug("GetDrMSSqlPrimaryInfo failed: " + err.Error())
		}

		sta_dr, err := GetDrMSSqlStandbyInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetDrMSSqlStandbyInfo failed: " + err.Error())
		}

		this.Data["pri_config"] = pri_config
		this.Data["sta_config"] = sta_config
		this.Data["pri_dr"] = pri_dr
		this.Data["sta_dr"] = sta_dr
		this.Data["dr_id"] = dr_id
		this.Data["asset_type"] = asset_type

		this.TplName = "dr_oper/manage_detail_mssql.tpl"
	}
	


}



//业务大屏
type ScreenDrViewController struct {
	controllers.BaseController
}

func (this *ScreenDrViewController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-screen-view") {
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

	drconf,_ := GetDrConfig(bs_id)

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

	//get
	redo, err := GetOracleRedo(pri_id)

	os_id := GetPrimaryOSId(pri_id)
	cpu_rate := GetCpuRateByOsId(os_id)
	mem_rate := GetMemRateByOsId(os_id)
	swap_rate := GetSwapRateByOsId(os_id)
	disk_rate := GetDiskRateByOsId(os_id)
	inode_rate := GetInodeRateByOsId(os_id)
	process_rate := GetProcessRateByDbId(pri_id)

	this.Data["pri_basic"] = pri_basic
	this.Data["sta_basic"] = sta_basic
	this.Data["pri_dr"] = pri_dr
	this.Data["sta_dr"] = sta_dr
	this.Data["drconf"] = drconf
	this.Data["redo"] = redo
	
	this.Data["cpu_rate"] = cpu_rate
	this.Data["mem_rate"] = mem_rate
	this.Data["swap_rate"] = swap_rate
	this.Data["disk_rate"] = disk_rate
	this.Data["inode_rate"] = inode_rate
	this.Data["process_rate"] = process_rate

	this.TplName = "dr_oper/screen-oracle.tpl"
}
