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
	_, _, ora := ListOracle(condArr, page, offset)

	
	core_db := GetCoreDb()

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["dbs"] = dbs
	this.Data["countDBs"] = countDBs
	this.Data["ora"] = ora
	this.Data["core_db"] = core_db

	this.TplName = "cfg_screen/index.tpl"
}

type AjaxSaveScreenController struct {
	controllers.BaseController
}

func (this *AjaxSaveScreenController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-screen-manage") {
		this.Abort("401")
	}

	ids := this.GetString("ids")
	core_db := this.GetString("core_db")
	//utils.LogDebug(ids)
	var err error
	if "" == ids {
		err = RemoveAllShowOnScreen(ids)
	} else {
		err = SaveShowOnScreen(ids)
	}

	if core_db != "" {
		err = SaveCoreDbOnScreen(core_db)
	}

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "大屏显示配置更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "大屏显示配置更改失败"}
	}
	this.ServeJSON()
}
