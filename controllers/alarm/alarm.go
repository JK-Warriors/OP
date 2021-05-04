package alarm

import (
	"strings"
	"opms/controllers"
	. "opms/models/alarm"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//告警管理
type ManageAlarmController struct {
	controllers.BaseController
}

func (this *ManageAlarmController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "alarm-manage") {
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

	CountAlerts := CountAlerts(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, CountAlerts)
	_, _, alerts := ListAlerts(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["alerts"] = alerts
	this.Data["CountAlerts"] = CountAlerts

	this.TplName = "alarm/index.tpl"
}

type AlarmHistoryListController struct {
	controllers.BaseController
}

func (this *AlarmHistoryListController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "alarm-history") {
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

	CountAlerts := CountAlertHistory(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, CountAlerts)
	_, _, alerts := ListAlertHistory(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["alerts"] = alerts
	this.Data["CountAlerts"] = CountAlerts

	this.TplName = "alarm/his_index.tpl"
}