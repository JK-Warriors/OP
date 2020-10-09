package cfg_trigger

import (
	"strings"
	"strconv"

	"opms/controllers"
	. "opms/models/cfg_trigger"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//告警配置管理
type ManageTriggerController struct {
	controllers.BaseController
}

func (this *ManageTriggerController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-alarm-manage") {
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

	countTriggers := CountTriggers(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countTriggers)
	_, _, triggers := ListTriggers(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["triggers"] = triggers
	this.Data["countTriggers"] = countTriggers

	this.TplName = "cfg_trigger/index.tpl"
}


type EditTriggerController struct {
	controllers.BaseController
}

func (this *EditTriggerController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-alarm-manage") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)

	triconf, err := GetTriggerConfig(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["triconf"] = triconf


	this.TplName = "cfg_trigger/form.tpl"
}


func (this *EditTriggerController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-alarm-manage") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}
	idstr := this.GetString("id")
	//utils.LogDebug(idstr)
	id, _ := strconv.Atoi(idstr)

	//idstr := this.GetString("asset_id")
	//asset_id, _ := strconv.Atoi(idstr)
	
	typestr := this.GetString("asset_type")
	asset_type, _ := strconv.Atoi(typestr)

	trigger_type := this.GetString("trigger_type")
	severity := this.GetString("severity")
	expression := this.GetString("expression")
	description := this.GetString("description")

	
	modestr := this.GetString("recovery_mode")
	recovery_mode, _ := strconv.Atoi(modestr)
	recovery_expression := this.GetString("recovery_expression")
	recovery_description := this.GetString("recovery_description")


	var triconf Trigger

	triconf.Id = id
	triconf.Asset_Type = asset_type
	triconf.Trigger_Type = trigger_type
	triconf.Severity = severity
	triconf.Expression = expression
	triconf.Description = description
	
	triconf.Recovery_Mode = recovery_mode
	triconf.Recovery_Expression = recovery_expression
	triconf.Recovery_Description = recovery_description
	

	err := UpdateTriggerConfig(id, triconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "修改告警配置成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "修改告警配置失败"}
	}
	this.ServeJSON()
}