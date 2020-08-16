package mssql

import (
	"strings"
	"opms/controllers"
	. "opms/models/mssql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//SQLServer状态管理
type ManageMssqlController struct {
	controllers.BaseController
}

func (this *ManageMssqlController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mssql-status-manage") {
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

	countMssql := CountMssql(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countMssql)
	_, _, mssql := ListMssqlStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["mssql"] = mssql
	this.Data["countMssql"] = countMssql

	this.TplName = "mssql/index.tpl"
}


