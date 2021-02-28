package cfg_screen

import (
	"strings"

	"opms/controllers"

	. "opms/models/cfg_screen"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//告警配置管理
type ManageScreenController struct {
	controllers.BaseController
}

func (this *ManageScreenController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-screen-manage") {
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

	asset_type := this.GetString("asset_type")
	host := this.GetString("host")
	alias := this.GetString("alias")
	condArr := make(map[string]string)
	condArr["asset_type"] = asset_type
	condArr["host"] = host
	condArr["alias"] = alias

	countDBs := CountDB(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDBs)
	_, _, dbs := ListDB(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["dbs"] = dbs
	this.Data["countDBs"] = countDBs

	this.TplName = "cfg_screen/index.tpl"
}

type EditScreenController struct {
	controllers.BaseController
}

func (this *EditScreenController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-screen-manage") {
		this.Abort("401")
	}

	this.TplName = "cfg_screen/form.tpl"
}
