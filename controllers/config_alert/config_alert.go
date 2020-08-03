package config_alert

import (
	"strings"
	"opms/controllers"
	. "opms/models/asset"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//告警配置管理
type ManageAlertController struct {
	controllers.BaseController
}

func (this *ManageAlertController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-alert-manage") {
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

	condArr := make(map[string]string)

	countAssets := CountAsset(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countAssets)
	_, _, asset := ListAssetStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["asset"] = asset

	this.TplName = "config-alert/index.tpl"
}
