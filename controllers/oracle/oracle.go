package oracle

import (
	"strings"
	"opms/controllers"
	. "opms/models/oracle"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//Oracle状态管理
type ManageOracleController struct {
	controllers.BaseController
}

func (this *ManageOracleController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oracle-status-manage") {
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

	countOracle := CountOracle(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countOracle)
	_, _, oracle := ListOracleStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["oracle"] = oracle
	this.Data["countOracle"] = countOracle

	this.TplName = "oracle/index.tpl"
}


//Oracle表空间管理
type ManageOracleTbsController struct {
	controllers.BaseController
}

func (this *ManageOracleTbsController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oracle-tbs-manage") {
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

	CountTbs := CountTablespaces(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, CountTbs)
	_, _, tbs := ListTablespaces(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["tbs"] = tbs
	this.Data["CountTbs"] = CountTbs

	this.TplName = "oracle/tbs_index.tpl"
}


//Oracle磁盘组管理
type ManageOracleAsmController struct {
	controllers.BaseController
}

func (this *ManageOracleAsmController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oracle-asm-manage") {
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

	CountDiskgroup := CountDiskgroups(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, CountDiskgroup)
	_, _, diskgroups := ListDiskgroups(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["diskgroups"] = diskgroups
	this.Data["CountDiskgroup"] = CountDiskgroup

	this.TplName = "oracle/asm_index.tpl"
}
