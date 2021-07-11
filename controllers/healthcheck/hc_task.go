package healthcheck

import (
	"opms/controllers"
	"strings"
	//"log"
	. "opms/models/dbconfig"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

)

//健康巡检
type ManageTaskController struct {
	controllers.BaseController
}

func (this *ManageTaskController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "hc-task-manage") {
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

	alias := this.GetString("alias")
	host := this.GetString("host")
	condArr := make(map[string]string)
	condArr["alias"] = alias
	condArr["host"] = host

	countDb := CountDBconfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDb)
	_, _, db := ListDBconfig(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["db"] = db
	this.Data["countDb"] = countDb


	this.TplName = "healthcheck/hc-status.tpl"
}