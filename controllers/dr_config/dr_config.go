package dr_config

import (
	"fmt"
	"opms/controllers"
	. "opms/models/dbconfig"
	. "opms/models/dr_config"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//容灾组管理
type ManageDrController struct {
	controllers.BaseController
}

func (this *ManageDrController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-dr-manage") {
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

	host := this.GetString("host")
	condArr := make(map[string]string)
	condArr["host"] = host

	countDr := CountDrConfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDr)
	_, _, drconf := ListDrConfig(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["drconf"] = drconf
	this.Data["countDr"] = countDr

	this.TplName = "dr_config/index.tpl"
}

//新增容灾组
type AddDrController struct {
	controllers.BaseController
}

func (this *AddDrController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-dr-manage") {
		this.Abort("401")
	}
	
	var drconf DrConfig
	this.Data["drconf"] = drconf
	
	pridbconf := ListPrimaryDBconfig()
	this.Data["pridbconf"] = pridbconf

	stadbconf := ListStandbyDBconfig()
	this.Data["stadbconf"] = stadbconf

	this.Data["dest_list"] = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	this.TplName = "dr_config/form.tpl"
}


func (this *AddDrController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-dr-manage") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}
	bs_name := this.GetString("bs_name")
	
	typestr := this.GetString("asset_type")
	asset_type, _ := strconv.Atoi(typestr)

	idstr := this.GetString("db_id_p")
	db_id_p, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_dest_p")
	db_dest_p, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_id_s")
	db_id_s, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_dest_s")
	db_dest_s, _ := strconv.Atoi(idstr)

	db_name := this.GetString("db_name")

	idstr = this.GetString("fb_retention")
	fb_retention, err := strconv.Atoi(idstr)
	if err != nil {
		fb_retention = -1
	}

	var is_shift int
	idstr = this.GetString("is_shift")
	if idstr == "on" {
		is_shift = 1
	} else {
		is_shift = 0
	}

	shift_vips := this.GetString("shift_vips")
	network_p := this.GetString("network_p")
	network_s := this.GetString("network_s")

	var drconf DrConfig

	drconf.Bs_Name = bs_name
	drconf.Asset_Type = asset_type
	drconf.Db_Id_P = db_id_p
	drconf.Db_Dest_P = db_dest_p
	drconf.Db_Id_S = db_id_s
	drconf.Db_Dest_S = db_dest_s
	drconf.Db_Name = db_name
	drconf.Fb_Retention = fb_retention
	drconf.Is_Shift = is_shift
	drconf.Shift_Vips = shift_vips
	drconf.Network_P = network_p
	drconf.Network_S = network_s

	err = AddDrConfig(drconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "新增容灾业务成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "新增容灾业务失败"}
	}
	this.ServeJSON()
}


//修改业务系统
type EditDrController struct {
	controllers.BaseController
}

func (this *EditDrController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-dr-manage") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)

	drconf, err := GetDrConfig(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["drconf"] = drconf

	pridbconf := ListPrimaryDBconfig()
	this.Data["pridbconf"] = pridbconf

	stadbconf := ListStandbyDBconfig()
	this.Data["stadbconf"] = stadbconf

	this.Data["dest_list"] = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	this.TplName = "dr_config/form.tpl"
}


func (this *EditDrController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-dr-manage") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}
	idstr := this.GetString("bs_id")
	//utils.LogDebug(idstr)
	bs_id, _ := strconv.Atoi(idstr)

	bs_name := this.GetString("bs_name")
	
	typestr := this.GetString("asset_type")
	asset_type, _ := strconv.Atoi(typestr)

	idstr = this.GetString("db_id_p")
	db_id_p, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_dest_p")
	db_dest_p, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_id_s")
	db_id_s, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_dest_s")
	db_dest_s, _ := strconv.Atoi(idstr)

	db_name := this.GetString("db_name")

	idstr = this.GetString("fb_retention")
	fb_retention, err := strconv.Atoi(idstr)
	if err != nil {
		fb_retention = 0
	}

	var is_shift int
	idstr = this.GetString("is_shift")
	if idstr == "on" {
		is_shift = 1
	} else {
		is_shift = 0
	}

	shift_vips := this.GetString("shift_vips")
	network_p := this.GetString("network_p")
	network_s := this.GetString("network_s")

	var drconf DrConfig

	drconf.Bs_Name = bs_name
	drconf.Asset_Type = asset_type
	drconf.Db_Id_P = db_id_p
	drconf.Db_Dest_P = db_dest_p
	drconf.Db_Id_S = db_id_s
	drconf.Db_Dest_S = db_dest_s
	drconf.Db_Name = db_name
	drconf.Fb_Retention = fb_retention
	drconf.Is_Shift = is_shift
	drconf.Shift_Vips = shift_vips
	drconf.Network_P = network_p
	drconf.Network_S = network_s

	//ldc, err := GetDrConfig(bs_id)
	//utils.LogDebug(ldc.Db_Id_P)

	err = UpdateDrConfig(bs_id, drconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "修改容灾配置成功", "id": fmt.Sprintf("%d", bs_id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "修改容灾配置失败"}
	}
	this.ServeJSON()
}

//容灾配置状态更改异步操作
type AjaxStatusDrConfigController struct {
	controllers.BaseController
}

func (this *AjaxStatusDrConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-dr-manage") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	id, _ := this.GetInt("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择容灾组"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 3 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}
	
	err := ChangeDrconfigStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "容灾组状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "容灾组更改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteDrConfigController struct {
	controllers.BaseController
}

func (this *AjaxDeleteDrConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-dr-manage") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权删除"}
		this.ServeJSON()
		return
	}
	ids := this.GetString("ids")
	if "" == ids {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的记录"}
		this.ServeJSON()
		return
	}

	err := DeleteDBconfig(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}