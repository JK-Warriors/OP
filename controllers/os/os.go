package os

import (
	"strings"
	"opms/controllers"
	. "opms/models/os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//OS状态管理
type ManageOSController struct {
	controllers.BaseController
}

func (this *ManageOSController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "os-status-manage") {
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

	countOS := CountOS(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countOS)
	_, _, osList := ListOSStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["countOS"] = countOS
	this.Data["osList"] = osList

	this.TplName = "os/index.tpl"
}


type ManageOSDiskController struct {
	controllers.BaseController
}

func (this *ManageOSDiskController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "os-disk-manage") {
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

	countDisk := CountDisk(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDisk)
	_, _, diskList := ListDiskStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["countDisk"] = countDisk
	this.Data["diskList"] = diskList

	this.TplName = "os/disk_index.tpl"
}

type ManageOSDiskIOController struct {
	controllers.BaseController
}

func (this *ManageOSDiskIOController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "os-io-manage") {
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

	countDiskio := CountDiskIO(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDiskio)
	_, _, diskioList := ListDiskIOStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["countDiskio"] = countDiskio
	this.Data["diskioList"] = diskioList

	this.TplName = "os/diskio_index.tpl"
}