package dbconfig

import (
	"database/sql"
	"fmt"
	"net"
	"opms/models"
	"opms/models/cfg_trigger"
	"opms/utils"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
	"golang.org/x/crypto/ssh"
)

type Dbconfigs struct {
	Id             int    `orm:"pk;column(id);"`
	Dbtype         int    `orm:"column(asset_type);"`
	Host           string `orm:"column(host);"`
	Protocol       string `orm:"column(protocol);"`
	Port           int    `orm:"column(port);"`
	Alias          string `orm:"column(alias);"`
	InstanceName   string `orm:"column(instance_name);"`
	Dbname         string `orm:"column(db_name);"`
	Username       string `orm:"column(username);"`
	Password       string `orm:"column(password);"`
	Role           int    `orm:"column(role);"`
	Ostype         int    `orm:"column(os_type);"`
	OsProtocol     string `orm:"column(os_protocol);"`
	OsPort         string `orm:"column(os_port);"`
	OsUsername     string `orm:"column(os_username);"`
	OsPassword     string `orm:"column(os_password);"`
	Status         int    `orm:"column(status);"`
	Display_Order         int    `orm:"column(display_order);"`
	IsDelete       int    `orm:"column(is_delete);"`
	Show_On_Screen int    `orm:"column(show_on_screen);"`
	Retention      int    `orm:"column(retention);"`
	Alert_Mail     int    `orm:"column(alert_mail);"`
	Alert_WeChat   int    `orm:"column(alert_wechat);"`
	Alert_SMS      int    `orm:"column(alert_sms);"`
	Created        int64  `orm:"column(created);"`
	Updated        int64  `orm:"column(updated);"`
}

func (this *Dbconfigs) TableName() string {
	return models.TableName("asset_config")
}
func init() {
	orm.RegisterModel(new(Dbconfigs))
}

//添加数据库
func AddDBconfig(upd Dbconfigs) error {
	o := orm.NewOrm()
	o.Using("default")
	dbconf := new(Dbconfigs)

	dbconf.Dbtype = upd.Dbtype
	dbconf.Host = upd.Host
	dbconf.Protocol = upd.Protocol
	dbconf.Port = upd.Port
	dbconf.Alias = upd.Alias
	dbconf.InstanceName = upd.InstanceName
	dbconf.Dbname = upd.Dbname
	dbconf.Username = upd.Username
	dbconf.Password = upd.Password
	dbconf.Role = upd.Role
	dbconf.Ostype = upd.Ostype
	dbconf.OsProtocol = upd.OsProtocol
	dbconf.OsPort = upd.OsPort
	dbconf.OsUsername = upd.OsUsername
	dbconf.OsPassword = upd.OsPassword
	dbconf.Display_Order = upd.Display_Order
	dbconf.Alert_Mail = upd.Alert_Mail
	dbconf.Alert_WeChat = upd.Alert_WeChat
	dbconf.Alert_SMS = upd.Alert_SMS
	dbconf.Status = 1
	dbconf.IsDelete = 0
	dbconf.Show_On_Screen = 0
	dbconf.Created = time.Now().Unix()
	id, err := o.Insert(dbconf)

	cfg_trigger.AddAssetTriggers(id, upd.Dbtype)

	return err
}

//修改数据库配置信息
func UpdateDBconfig(id int, upd Dbconfigs) error {
	var dbconf Dbconfigs
	o := orm.NewOrm()
	dbconf, err := GetDBconfig(id)
	if err == nil {
		dbconf.Dbtype = upd.Dbtype
		dbconf.Host = upd.Host
		dbconf.Protocol = upd.Protocol
		dbconf.Port = upd.Port
		dbconf.Alias = upd.Alias
		dbconf.InstanceName = upd.InstanceName
		dbconf.Dbname = upd.Dbname
		dbconf.Username = upd.Username
		dbconf.Password = upd.Password
		dbconf.Role = upd.Role
		dbconf.Ostype = upd.Ostype
		dbconf.OsProtocol = upd.OsProtocol
		dbconf.OsPort = upd.OsPort
		dbconf.OsUsername = upd.OsUsername
		dbconf.OsPassword = upd.OsPassword
		dbconf.Display_Order = upd.Display_Order
		dbconf.Alert_Mail = upd.Alert_Mail
		dbconf.Alert_WeChat = upd.Alert_WeChat
		dbconf.Alert_SMS = upd.Alert_SMS
		dbconf.Updated = time.Now().Unix()

		_, err = o.Update(&dbconf)
	}
	return err
}

//得到数据库配置信息
func GetDBconfig(id int) (Dbconfigs, error) {
	var dbconf Dbconfigs
	var err error
	o := orm.NewOrm()

	dbconf = Dbconfigs{Id: id}
	err = o.Read(&dbconf)

	if err == orm.ErrNoRows {
		return dbconf, nil
	}
	return dbconf, err
}

//根据ID获取数据库类型
func GetDBtypeByDBId(id int) int {
	var asset_type int

	o := orm.NewOrm()
	o.Using("default")

	sql := `select asset_type from pms_asset_config where id = ?`
	err := o.Raw(sql, id).QueryRow(&asset_type)
	if err != nil {
		utils.LogDebug("GetDBtypeByDBId failed: " + err.Error())
		return -1
	}
	return asset_type
}

//根据类型ID获取数据库类型
func GetDBtype(id int) string {
	var asset_type string

	if id == 1 {
		asset_type = "Oracle"
	} else if id == 2 {
		asset_type = "MySQL"
	} else if id == 3 {
		asset_type = "SQLServer"
	} else if id == 99 {
		asset_type = "OS"
	}
	return asset_type
}

//根据资产ID获取资产描述
func GetDBDesc(id int) string {
	o := orm.NewOrm()
	o.Using("default")
	var db_desc string

	sql := `select concat(host, ':', port, ' (' , alias, ')') from pms_asset_config where id = ?`
	err := o.Raw(sql, id).QueryRow(&db_desc)
	if err != nil {
		utils.LogDebug("GetDBDesc failed: " + err.Error())
		return ""
	}

	return db_desc
}

//根据资产ID获取资产别名
func GetDBAlias(id int) string {
	o := orm.NewOrm()
	o.Using("default")
	var db_alias string

	sql := `select alias from pms_asset_config where id = ?`
	err := o.Raw(sql, id).QueryRow(&db_alias)
	if err != nil {
		utils.LogDebug("GetDBAlias failed: " + err.Error())
		return ""
	}

	return db_alias
}

//获取数据库配置列表
func ListDBconfig(condArr map[string]string, page int, offset int) (num int64, err error, dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	if condArr["asset_type"] != "" {
		cond = cond.And("asset_type", condArr["asset_type"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}

	cond = cond.And("is_delete", 0)

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("display_order")
	nums, errs := qs.Limit(offset, start).All(&dbconf)
	return nums, errs, dbconf
}

func ListAllDBconfig() (dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&dbconf)
	return dbconf
}

func ListScreenDBconfig() (dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")

	sql := `select * from pms_asset_config where is_delete = 0 and show_on_screen = 1 `
	_, _ = o.Raw(sql).QueryRows(&dbconf)

	return dbconf
}

func ListPrimaryDBconfig() (dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	cond = cond.And("role", 1)
	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&dbconf)
	return dbconf
}

func ListStandbyDBconfig() (dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	cond = cond.And("role", 2)
	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&dbconf)
	return dbconf
}

//统计数量
func CountDBconfig(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("asset_config"))
	cond := orm.NewCondition()

	if condArr["asset_type"] != "" {
		cond = cond.And("asset_type", condArr["asset_type"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	cond = cond.And("status", 1)
	cond = cond.And("is_delete", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteDBconfig(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("asset_config") + " WHERE id IN(" + ids + ")").Exec()
	_, err = o.Raw("DELETE FROM pms_triggers WHERE asset_id IN(" + ids + ")").Exec()

	return err
}

//更改资产状态
func ChangeDBconfigStatus(id int, status int) error {
	o := orm.NewOrm()

	dbconf := Dbconfigs{Id: id}
	err := o.Read(&dbconf, "id")
	if nil != err {
		return err
	} else {
		dbconf.Status = status
		_, err := o.Update(&dbconf)
		return err
	}
}

func CheckOracleConnect(host string, port string, inst_name string, username string, password string) error {
	con_str := username + "/" + password + "@" + host + ":" + port + "/" + inst_name + "?timeout=3s"
	//db, err := sql.Open("oci8", "sys/oracle@192.168.133.40:1521/orcl?as=sysdba")
	db, err := sql.Open("oci8", con_str)
	defer db.Close()

	err = db.Ping()

	//ORA-28009: connection as SYS should be as SYSDBA or SYSOPER
	if err != nil {
		utils.LogDebugf("Open connection as normal failed: %s", err.Error())

		if strings.Contains(err.Error(), "ORA-28009") || strings.Contains(err.Error(), "driver: bad connection") {
			con_str = username + "/" + password + "@" + host + ":" + port + "/" + inst_name + "?as=sysdba&timeout=3s"
			db, err = sql.Open("oci8", con_str)
			defer db.Close()

			err = db.Ping()

			if err != nil {
				utils.LogDebugf("Open connection as sysdba failed: %s", err.Error())
			} else {
				utils.LogDebug("Open connection as sysdba successfully.")

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

	err = db.Ping()
	if err != nil {
		utils.LogDebugf("Open Connection failed: %s", err.Error())
	}

	return err
}

func CheckSqlserverConnect(host string, port string, inst_name string, db_name string, username string, password string) error {
	var dbname string
	if db_name == "" {
		dbname = "master"
	} else {
		dbname = db_name
	}
	//连接字符串
	con_str := fmt.Sprintf("server=%s\\%s;port%s;database=%s;user id=%s;password=%s;encrypt=disable", host, inst_name, port, dbname, username, password)

	//建立连接
	db, err := sql.Open("mssql", con_str)
	if err != nil {
		utils.LogDebugf("Open Connection failed: %s", err.Error())
		return err
	}
	defer db.Close()

	//验证连接
	err = db.Ping()
	if err != nil {
		utils.LogDebugf("Ping sqlserver failed: %s", err.Error())
	}

	return err
}

type TelnetClient struct {
	Host             string
	Port             string
	IsAuthentication bool
	UserName         string
	Password         string
}

const (
	//经过测试，嵌入式设备下，延时大概需要大于300ms
	TIME_DELAY_AFTER_WRITE = 300 //300ms
)

func CheckOSConnect(host string, port string, protocol string, username string, password string) error {
	var err error

	if protocol == "ssh" {
		//dial 获取ssh client
		config := &ssh.ClientConfig{
			Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
			User:            username,
			Auth:            []ssh.AuthMethod{ssh.Password(password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), //不够安全
		}

		addr := fmt.Sprintf("%s:%s", host, port)
		_, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			utils.LogDebugf("SSH dial failed: %s", err.Error())
		}
		return err
	} else {

		telnetClientObj := new(TelnetClient)
		telnetClientObj.Host = host
		telnetClientObj.Port = port
		telnetClientObj.IsAuthentication = true
		telnetClientObj.UserName = username
		telnetClientObj.Password = password

		err = telnetClientObj.Telnet(5)
	}

	return err
}

func (this *TelnetClient) Telnet(timeout int) error {
	addr := fmt.Sprintf("%s:%s", this.Host, this.Port)
	conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout)*time.Second)
	if nil != err {
		utils.LogDebugf("net.DialTimeout, errInfo: %s", err.Error())
		return err
	}
	defer conn.Close()

	err = this.telnetProtocolHandshake(conn)
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake failed: %s", err.Error())
		return err
	}

	return err
}

func (this *TelnetClient) telnetProtocolHandshake(conn net.Conn) error {
	var buf [4096]byte
	n, err := conn.Read(buf[0:])
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Read, errInfo:", err.Error())
		return err
	}
	utils.LogDebug(string(buf[0:n]))

	buf[1] = 252
	buf[4] = 252
	buf[7] = 252
	buf[10] = 252
	utils.LogDebug((buf[0:n]))
	n, err = conn.Write(buf[0:n])
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Write, errInfo:", err.Error())
		return err
	}
	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)

	n, err = conn.Read(buf[0:])
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Read, errInfo:", err.Error())
		return err
	}
	utils.LogDebug(string(buf[0:n]))

	buf[1] = 252
	buf[4] = 251
	buf[7] = 252
	buf[10] = 254
	buf[13] = 252
	utils.LogDebug((buf[0:n]))
	n, err = conn.Write(buf[0:n])
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Write, errInfo:", err.Error())
		return err
	}
	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)

	n, err = conn.Read(buf[0:])
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Read, errInfo:", err.Error())
		return err
	}
	utils.LogDebug(string(buf[0:n]))
	utils.LogDebug((buf[0:n]))

	if false == this.IsAuthentication {
		return nil
	}

	n, err = conn.Write([]byte(this.UserName + "\n"))
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Write, errInfo:", err.Error())
		return err
	}
	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)

	n, err = conn.Read(buf[0:])
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Read, errInfo:", err.Error())
		return err
	}
	utils.LogDebug(string(buf[0:n]))

	n, err = conn.Write([]byte(this.Password + "\n"))
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Write, errInfo:", err.Error())
		return err
	}
	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)

	n, err = conn.Read(buf[0:])
	if nil != err {
		utils.LogDebugf("telnetProtocolHandshake, method: conn.Read, errInfo:", err.Error())
		return err
	}
	utils.LogDebug(string(buf[0:n]))

	return nil
}
