package asset

import (
	"strings"
	"opms/controllers"
	. "opms/models/asset"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//资产状态管理
type ManageAssetController struct {
	controllers.BaseController
}

func (this *ManageAssetController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "asset-status-manage") {
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

	countAssets := CountAsset(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countAssets)
	_, _, asset := ListAssetStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["asset"] = asset
	this.Data["countAssets"] = countAssets

	this.TplName = "asset/index.tpl"
}
