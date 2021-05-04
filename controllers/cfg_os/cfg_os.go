package cfg_os

import (
	"fmt"
	"opms/controllers"
	. "opms/models/dbconfig"
	. "opms/models/cfg_os"
	"strconv"
	"strings"
	//"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//操作系统管理
type ManageOSConfigController struct {
	controllers.BaseController
}

func (this *ManageOSConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-os-manage") {
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

	countOS := CountOSconfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countOS)
	_, _, osconf := ListOSconfig(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["osconf"] = osconf
	this.Data["countOS"] = countOS

	this.TplName = "cfg_os/config-index.tpl"
}

//添加配置信息
type AddOSConfigController struct {
	controllers.BaseController
}

func (this *AddOSConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-os-add") {
		this.Abort("401")
	}

	var osconf Dbconfigs
	this.Data["osconf"] = osconf

	this.TplName = "cfg_os/config-form.tpl"
}

func (this *AddOSConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-os-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权新增"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	alias := this.GetString("alias")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写别名"}
		this.ServeJSON()
		return
	}

	os_type, _ := this.GetInt("os_type")
	os_protocol := this.GetString("os_protocol")
	os_port,_ := this.GetInt("os_port")
	os_username := this.GetString("os_username")
	os_password := this.GetString("os_password")
	

	alert_mail,_ := this.GetInt("alert_mail")
	alert_wechat,_  := this.GetInt("alert_wechat")
	alert_sms,_  := this.GetInt("alert_sms")

	var osconf Dbconfigs

	osconf.Dbtype = 99
	osconf.Host = host
	osconf.Alias = alias
	osconf.Ostype = os_type
	osconf.OsProtocol = os_protocol
	osconf.OsPort = os_port
	osconf.OsUsername = os_username
	osconf.OsPassword = os_password
	osconf.Alert_Mail = alert_mail
	osconf.Alert_WeChat = alert_wechat
	osconf.Alert_SMS = alert_sms

	count := CheckOsExists(osconf) 
	if count == 0{
		err := AddOSconfig(osconf)
		
		if err == nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作系统信息添加成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "操作系统信息添加失败"}
		}
	}else{
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "操作系统的IP已经存在"}
	}

	this.ServeJSON()
}

//修改资产配置信息
type EditOSConfigController struct {
	controllers.BaseController
}

func (this *EditOSConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-os-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	osconf, err := GetOSconfig(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["osconf"] = osconf


	this.TplName = "cfg_os/config-form.tpl"
}

func (this *EditOSConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-os-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}

	id, _ := this.GetInt("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户参数出错"}
		this.ServeJSON()
		return
	}


	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	alias := this.GetString("alias")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写别名"}
		this.ServeJSON()
		return
	}

	os_type, _ := this.GetInt("os_type")
	os_protocol := this.GetString("os_protocol")
	os_port,_ := this.GetInt("os_port")
	os_username := this.GetString("os_username")
	os_password := this.GetString("os_password")

	alert_mail,_ := this.GetInt("alert_mail")
	alert_wechat,_ := this.GetInt("alert_wechat")
	alert_sms,_ := this.GetInt("alert_sms")

	var osconf Dbconfigs

	osconf.Host = host
	osconf.Alias = alias
	osconf.Ostype = os_type
	osconf.OsProtocol = os_protocol
	osconf.OsPort = os_port
	osconf.OsUsername = os_username
	osconf.OsPassword = os_password
	osconf.Alert_Mail = alert_mail
	osconf.Alert_WeChat = alert_wechat
	osconf.Alert_SMS = alert_sms


	count := CheckOsExists(osconf) 
	if count == 0{
		err := UpdateOSconfig(id, osconf)
		if err == nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作系统配置信息修改成功", "id": fmt.Sprintf("%d", id)}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "操作系统配置信息修改失败"}
		}
	}else{
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "操作系统IP已经存在"}
	}
	this.ServeJSON()
}

//配置状态更改异步操作
type AjaxStatusOSConfigController struct {
	controllers.BaseController
}

func (this *AjaxStatusOSConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-os-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	id, _ := this.GetInt("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作系统"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 3 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeOSconfigStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作系统状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "操作系统状态更改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteOSConfigController struct {
	controllers.BaseController
}

func (this *AjaxDeleteOSConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-os-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权删除"}
		this.ServeJSON()
		return
	}
	ids := this.GetString("ids")
	if "" == ids {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的记录"}
		this.ServeJSON()
		return
	}

	err := DeleteOSconfig(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

type AjaxConnectOSConfigController struct {
	controllers.BaseController
}

func (this *AjaxConnectOSConfigController) Post() {
	//权限检测
	host := this.GetString("host")
	protocol := this.GetString("protocol")
	if(protocol == ""){
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测试连接失败，请选择协议！"}
		this.ServeJSON()
		return
	}

	port := this.GetString("port")
	username := this.GetString("username")
	password := this.GetString("password")

	err := CheckOSConnect(host, port, protocol, username, password)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "测试连接成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测试连接失败: " + err.Error()}
	}
	this.ServeJSON()
}
