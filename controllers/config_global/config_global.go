package config_global

import (
	"strings"
	"opms/controllers"
	. "opms/models/config_global"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//全局配置管理
type ManageGlobalController struct {
	controllers.BaseController
}

func (this *ManageGlobalController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-global-manage") {
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

	countOptions := CountOptions(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countOptions)
	_, _, options := ListGlobalOptions(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["options"] = options
	this.Data["countOptions"] = countOptions

	this.TplName = "config_global/index.tpl"
}


//全局配置更改异步操作
type AjaxSaveGlobalController struct {
	controllers.BaseController
}

func (this *AjaxSaveGlobalController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-global-manage") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	id := this.GetString("id")
	value := this.GetString("value")

	err := SaveGlobalConfig(id, value)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "全局配置更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "全局配置更改失败"}
	}
	this.ServeJSON()
}
