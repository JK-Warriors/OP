package dbconfig

import (
	"database/sql"
	"fmt"
	"opms/controllers"
	. "opms/models/business"
	. "opms/models/dbconfig"
	"opms/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
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

	dbtype := this.GetString("dbtype")
	host := this.GetString("host")
	alias := this.GetString("alias")
	condArr := make(map[string]string)
	condArr["dbtype"] = dbtype
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
	db_type, _ := this.GetInt64("db_type")
	this.Data["db_type"] = db_type

	var dbconf Dbconfigs
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

	bs_id, _ := this.GetInt("bs_id")

	db_type, _ := this.GetInt("db_type")
	if db_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择数据库类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
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
	if "" == username {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
		return
	}

	role, _ := this.GetInt("role")
	if role <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择角色"}
		this.ServeJSON()
		return
	}

	var dbconf Dbconfigs

	dbconf.Dbtype = db_type
	dbconf.Host = host
	dbconf.Port = port
	dbconf.Alias = this.GetString("alias")
	dbconf.InstanceName = this.GetString("instance_name")
	dbconf.Dbname = this.GetString("db_name")
	dbconf.Username = username
	dbconf.Password = password
	dbconf.Bs_Id = bs_id
	dbconf.Role = role

	err := AddDBconfig(dbconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "数据库配置信息添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "数据库配置信息添加失败"}
	}
	this.ServeJSON()
}

//修改数据库配置信息
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

	db_type, _ := this.GetInt("db_type")
	if db_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择数据库类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
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
	if "" == username {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
		return
	}

	bs_id, _ := this.GetInt("bs_id")

	role, _ := this.GetInt("role")
	if role <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择角色"}
		this.ServeJSON()
		return
	}

	var dbconf Dbconfigs

	dbconf.Dbtype = db_type
	dbconf.Host = host
	dbconf.Port = port
	dbconf.Alias = this.GetString("alias")
	dbconf.InstanceName = this.GetString("instance_name")
	dbconf.Dbname = this.GetString("db_name")
	dbconf.Username = username
	dbconf.Password = password
	dbconf.Bs_Id = bs_id
	dbconf.Role = role

	err := UpdateDBconfig(id, dbconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "数据库配置信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "数据库配置信息修改失败"}
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
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "数据库状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "数据库状态更改失败"}
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
	db_type := this.GetString("db_type")
	host := this.GetString("host")
	port := this.GetString("port")
	username := this.GetString("username")
	password := this.GetString("password")
	inst_name := this.GetString("inst_name")
	db_name := this.GetString("db_name")

	var err error

	if db_type == "1" {
		err = CheckOracleConnect(host, port, inst_name, username, password)
	} else if db_type == "2" {
		err = CheckMysqlConnect(host, port, db_name, username, password)
	} else if db_type == "3" {
		err = CheckSqlserverConnect(host, port, inst_name, db_name, username, password)
	}

	//utils.LogDebug(err)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "测试连接成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "测试连接失败: " + err.Error()}
	}
	this.ServeJSON()
}

func CheckOracleConnect(host string, port string, inst_name string, username string, password string) error {
	con_str := username + "/" + password + "@" + host + ":" + port + "/" + inst_name + "?timeout=5s&readTimeout=6s"
	//db, err := sql.Open("oci8", "sys/oracle@192.168.133.40:1521/orcl?as=sysdba")
	db, err := sql.Open("oci8", con_str)
	defer db.Close()

	_, err = db.Query("select 1 from dual")

	//err_str := fmt.Sprintf("%s", err)

	//ORA-28009: connection as SYS should be as SYSDBA or SYSOPER
	if err != nil {
		if strings.Contains(err.Error(), "ORA-28009") || strings.Contains(err.Error(), "driver: bad connection") {
			con_str = username + "/" + password + "@" + host + ":" + port + "/" + inst_name + "?as=sysdba&timeout=5s&readTimeout=6s"
			db, err = sql.Open("oci8", con_str)
			defer db.Close()

			_, err = db.Query("select 1 from dual")

			if err != nil {
				utils.LogDebug("Open Connection failed: " + err.Error())
			}
		}
	}

	return err
}

func CheckMysqlConnect(host string, port string, db_name string, username string, password string) error {
	//con_str := "root:Aa123456@tcp(192.168.0.101:3306)/?timeout=5s&readTimeout=6s"
	con_str := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db_name + "?timeout=5s&readTimeout=6s"
	db, err := sql.Open("mysql", con_str)
	defer db.Close()

	_, err = db.Query("select 1")
	if err != nil {
		utils.LogDebug("Open Connection failed: " + err.Error())
	}

	return err
}

func CheckSqlserverConnect(host string, port string, inst_name string, db_name string, username string, password string) error {
	//连接字符串
	con_str := fmt.Sprintf("server=%s;port%s;database=%s;user id=%s;password=%s", host, port, db_name, username, password)

	//建立连接
	conn, err := sql.Open("mssql", con_str)
	if err != nil {
		utils.LogDebug("Open Connection failed: " + err.Error())
		return err
	}
	defer conn.Close()

	//产生查询语句的Statement
	stmt, err := conn.Prepare("select 1")
	if err != nil {
		utils.LogDebug("Open Connection failed: " + err.Error())
		return err
	}
	defer stmt.Close()

	//通过Statement执行查询
	rows, err := stmt.Query()
	defer rows.Close()

	if err != nil {
		utils.LogDebug("Open Connection failed: " + err.Error())
	}

	return err
}
