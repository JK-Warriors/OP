package mysql

import (
	"strings"
	"opms/controllers"
	. "opms/models/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//Mysql状态管理
type ManageMysqlController struct {
	controllers.BaseController
}

func (this *ManageMysqlController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mysql-status-manage") {
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

	countMysql := CountMysql(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countMysql)
	_, _, mysql := ListMysqlStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["mysql"] = mysql
	this.Data["countMysql"] = countMysql

	this.TplName = "mysql/index.tpl"
}

//Mysql资源管理
type ResourceMysqlController struct {
	controllers.BaseController
}

func (this *ResourceMysqlController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mysql-status-manage") {
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

	countMysql := CountMysql(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countMysql)
	_, _, mysql := ListMysqlStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["mysql"] = mysql
	this.Data["countMysql"] = countMysql

	this.TplName = "mysql/resource.tpl"
}

//Mysql键管理
type KeyMysqlController struct {
	controllers.BaseController
}

func (this *KeyMysqlController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mysql-status-manage") {
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

	countMysql := CountMysql(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countMysql)
	_, _, mysql := ListMysqlStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["mysql"] = mysql
	this.Data["countMysql"] = countMysql

	this.TplName = "mysql/key.tpl"
}


