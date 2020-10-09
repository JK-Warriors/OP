package dbconfig

import (
	"fmt"
	"opms/controllers"
	. "opms/models/dr_business"
	. "opms/models/dbconfig"
	"strconv"
	"strings"
	//"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//库管理
type ManageDBConfigController struct {
	controllers.BaseController
}

func (this *ManageDBConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-db-manage") {
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

	countDb := CountDBconfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDb)
	_, _, dbconf := ListDBconfig(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["dbconf"] = dbconf
	this.Data["countDb"] = countDb

	this.TplName = "dbconfig/dbconfig-index.tpl"
}

//添加数据库配置信息
type AddDBConfigController struct {
	controllers.BaseController
}

func (this *AddDBConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-db-add") {
		this.Abort("401")
	}
	asset_type, _ := this.GetInt("asset_type")

	var dbconf Dbconfigs
	dbconf.Dbtype = asset_type
	this.Data["dbconf"] = dbconf

	bsconf := ListAllBusiness()
	this.Data["bsconf"] = bsconf

	this.TplName = "dbconfig/dbconfig-form.tpl"
}

func (this *AddDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-db-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权新增"}
		this.ServeJSON()
		return
	}

	asset_type, _ := this.GetInt("asset_type")
	if asset_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选资产类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	protocol := this.GetString("protocol")
	if ("" == protocol && asset_type == 99) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择协议"}
		this.ServeJSON()
		return
	}

	port, _ := this.GetInt("port")
	if port <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写端口"}
		this.ServeJSON()
		return
	}

	username := this.GetString("username")
	if ("" == username && asset_type != 99) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if ("" == password && asset_type != 99) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
		return
	}

	role, _ := this.GetInt("role")
	if (role <= 0 && asset_type != 99) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择角色"}
		this.ServeJSON()
		return
	}

	os_type, _ := this.GetInt("os_type")
	os_protocol := this.GetString("os_protocol")
	os_port := this.GetString("os_port")
	os_username := this.GetString("os_username")
	os_password := this.GetString("os_password")
	
	alert_mail,_ := this.GetInt("alert_mail")
	alert_wechat,_  := this.GetInt("alert_wechat")
	alert_sms,_  := this.GetInt("alert_sms")

	var dbconf Dbconfigs

	dbconf.Dbtype = asset_type
	dbconf.Host = host
	dbconf.Protocol = protocol
	dbconf.Port = port
	dbconf.Alias = this.GetString("alias")
	dbconf.InstanceName = this.GetString("instance_name")
	dbconf.Dbname = this.GetString("db_name")
	dbconf.Username = username
	dbconf.Password = password
	dbconf.Role = role
	dbconf.Ostype = os_type
	dbconf.OsProtocol = os_protocol
	dbconf.OsPort = os_port
	dbconf.OsUsername = os_username
	dbconf.OsPassword = os_password
	dbconf.Alert_Mail = alert_mail
	dbconf.Alert_WeChat = alert_wechat
	dbconf.Alert_SMS = alert_sms

	err := AddDBconfig(dbconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "资产配置信息添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "资产配置信息添加失败"}
	}
	this.ServeJSON()
}

//修改资产配置信息
type EditDBConfigController struct {
	controllers.BaseController
}

func (this *EditDBConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-db-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	dbconf, err := GetDBconfig(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["dbconf"] = dbconf

	bsconf := ListAllBusiness()
	this.Data["bsconf"] = bsconf

	this.TplName = "dbconfig/dbconfig-form.tpl"
}

func (this *EditDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-db-edit") {
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

	asset_type, _ := this.GetInt("asset_type")
	if asset_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择资产类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	protocol := this.GetString("protocol")
	if ("" == protocol && asset_type == 99) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择协议"}
		this.ServeJSON()
		return
	}

	port, _ := this.GetInt("port")
	if port <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写端口"}
		this.ServeJSON()
		return
	}

	username := this.GetString("username")
	if ("" == username && asset_type != 99)  {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if ("" == password && asset_type != 99)  {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
		return
	}


	role, _ := this.GetInt("role")
	if (role <= 0 && asset_type != 99)  {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择角色"}
		this.ServeJSON()
		return
	}

	os_type, _ := this.GetInt("os_type")
	os_protocol := this.GetString("os_protocol")
	os_port := this.GetString("os_port")
	os_username := this.GetString("os_username")
	os_password := this.GetString("os_password")
	
	alert_mail,_ := this.GetInt("alert_mail")
	alert_wechat,_ := this.GetInt("alert_wechat")
	alert_sms,_ := this.GetInt("alert_sms")

	var dbconf Dbconfigs

	dbconf.Dbtype = asset_type
	dbconf.Host = host
	dbconf.Protocol = protocol
	dbconf.Port = port
	dbconf.Alias = this.GetString("alias")
	dbconf.InstanceName = this.GetString("instance_name")
	dbconf.Dbname = this.GetString("db_name")
	dbconf.Username = username
	dbconf.Password = password
	dbconf.Role = role
	dbconf.Ostype = os_type
	dbconf.OsProtocol = os_protocol
	dbconf.OsPort = os_port
	dbconf.OsUsername = os_username
	dbconf.OsPassword = os_password
	dbconf.Alert_Mail = alert_mail
	dbconf.Alert_WeChat = alert_wechat
	dbconf.Alert_SMS = alert_sms


	err := UpdateDBconfig(id, dbconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "资产配置信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "资产配置信息修改失败"}
	}
	this.ServeJSON()
}

//数据库配置状态更改异步操作
type AjaxStatusDBConfigController struct {
	controllers.BaseController
}

func (this *AjaxStatusDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-db-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	id, _ := this.GetInt("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择数据库"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 3 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeDBconfigStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "资产状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "资产状态更改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteDBConfigController struct {
	controllers.BaseController
}

func (this *AjaxDeleteDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-db-delete") {
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

	err := DeleteDBconfig(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

type AjaxConnectDBConfigController struct {
	controllers.BaseController
}

func (this *AjaxConnectDBConfigController) Post() {
	//权限检测
	asset_type := this.GetString("asset_type")
	host := this.GetString("host")
	protocol := this.GetString("protocol")
	if(asset_type == "99" && protocol == ""){
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测试连接失败，请选择协议！"}
		this.ServeJSON()
		return
	}

	port := this.GetString("port")
	username := this.GetString("username")
	password := this.GetString("password")
	inst_name := this.GetString("inst_name")
	db_name := this.GetString("db_name")

	var err error

	if asset_type == "1" {
		err = CheckOracleConnect(host, port, inst_name, username, password)
	} else if asset_type == "2" {
		err = CheckMysqlConnect(host, port, db_name, username, password)
	} else if asset_type == "3" {
		err = CheckSqlserverConnect(host, port, inst_name, db_name, username, password)
	} else if asset_type == "99" {
		err = CheckOSConnect(host, port, protocol, username, password)
	}

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "测试连接成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测试连接失败: " + err.Error()}
	}
	this.ServeJSON()
}
