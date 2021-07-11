package cfg_os

import (
	"fmt"
	"net"
	. "opms/models/dbconfig"
	"opms/models/cfg_trigger"
	"opms/utils"
	"time"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
	"golang.org/x/crypto/ssh"
	gs "github.com/soniah/gosnmp"
)
/*
type OSconfig struct {
	Id             int    `orm:"pk;column(id);"`
	Host           string `orm:"column(host);"`
	Alias          string `orm:"column(alias);"`
	Type           int    `orm:"column(os_type);"`
	Protocol       string `orm:"column(os_protocol);"`
	Port           string `orm:"column(os_port);"`
	Username       string `orm:"column(os_username);"`
	Password       string `orm:"column(os_password);"`
	Status         int    `orm:"column(status);"`
	IsDelete       int    `orm:"column(is_delete);"`
	Alert_Mail     int    `orm:"column(alert_mail);"`
	Alert_WeChat   int    `orm:"column(alert_wechat);"`
	Alert_SMS      int    `orm:"column(alert_sms);"`
	Created        int64  `orm:"column(created);"`
	Updated        int64  `orm:"column(updated);"`
}

func (this *OSconfig) TableName() string {
	return models.TableName("asset_config")
}
*/

//添加操作系统
func AddOSconfig(upd Dbconfigs) error {
	o := orm.NewOrm()
	o.Using("default")
	osconf := new(Dbconfigs)

	osconf.Dbtype = upd.Dbtype
	osconf.Host = upd.Host
	osconf.Alias = upd.Alias
	osconf.Ostype = upd.Ostype
	osconf.OsProtocol = upd.OsProtocol
	osconf.OsPort = upd.OsPort
	osconf.OsUsername = upd.OsUsername
	osconf.OsPassword = upd.OsPassword
	osconf.Alert_Mail = upd.Alert_Mail
	osconf.Alert_WeChat = upd.Alert_WeChat
	osconf.Alert_SMS = upd.Alert_SMS
	osconf.Status = 1
	osconf.IsDelete = 0
	osconf.Created = time.Now().Unix()

	id, err := o.Insert(osconf)
	cfg_trigger.AddAssetTriggers(id, 99)
	return err
}

//修改数据库配置信息
func UpdateOSconfig(id int, upd Dbconfigs) error {
	var osconf Dbconfigs
	o := orm.NewOrm()
	osconf, err := GetOSconfig(id)
	if err == nil {
		osconf.Host = upd.Host
		osconf.Alias = upd.Alias
		osconf.Ostype = upd.Ostype
		osconf.OsProtocol = upd.OsProtocol
		osconf.OsPort = upd.OsPort
		osconf.OsUsername = upd.OsUsername
		osconf.OsPassword = upd.OsPassword
		osconf.Is_Alert = upd.Is_Alert
		osconf.Alert_Mail = upd.Alert_Mail
		osconf.Alert_WeChat = upd.Alert_WeChat
		osconf.Alert_SMS = upd.Alert_SMS
		osconf.Updated = time.Now().Unix()

		_, err = o.Update(&osconf)
	}
	return err
}

func CheckOsExists(upd Dbconfigs) int{
	var sql string
	if upd.Id > 0 {
		sql = "select count(1) from pms_asset_config where asset_type = 99 and host = ? and id != " + strconv.Itoa(upd.Id) 

	}else{
		sql = "select count(1) from pms_asset_config where asset_type = 99 and host = ?"
	}

	var count int
	o := orm.NewOrm()
	err := o.Raw(sql, upd.Host).QueryRow(&count)
	if err != nil{
		return  -1
	}

	return count
}

//得到操作系统配置信息
func GetOSconfig(id int) (Dbconfigs, error) {
	var osconf Dbconfigs
	var err error
	o := orm.NewOrm()

	osconf = Dbconfigs{Id: id}
	err = o.Read(&osconf)

	if err == orm.ErrNoRows {
		return osconf, nil
	}
	return osconf, err
}


//获取操作系统配置列表
func ListOSconfig(condArr map[string]string, page int, offset int) (num int64, err error, osconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("pms_asset_config")
	cond := orm.NewCondition()

	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}

	cond = cond.And("asset_type", 99)
	cond = cond.And("is_delete", 0)

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("id")
	nums, errs := qs.Limit(offset, start).All(&osconf)
	return nums, errs, osconf
}

func ListAllOSconfig() (osconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("pms_asset_config")
	cond := orm.NewCondition()

	cond = cond.And("asset_type", 99)
	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&osconf)
	return osconf
}



//统计数量
func CountOSconfig(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("pms_asset_config")
	cond := orm.NewCondition()

	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	cond = cond.And("asset_type", 99)
	cond = cond.And("is_delete", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteOSconfig(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM pms_asset_config WHERE id IN(" + ids + ")").Exec()
	_, err = o.Raw("DELETE FROM pms_triggers WHERE asset_id IN(" + ids + ")").Exec()

	return err
}

//更改资产状态
func ChangeOSconfigStatus(id int, status int) error {
	o := orm.NewOrm()

	osconf := Dbconfigs{Id: id}
	err := o.Read(&osconf, "id")
	if nil != err {
		return err
	} else {
		osconf.Status = status
		_, err := o.Update(&osconf)
		return err
	}
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
	} else if protocol == "snmp"{
		//连接字符串
		l_port,_ := strconv.Atoi(port)
		snmp := &gs.GoSNMP{
			Target:    host,
			Port:      uint16(l_port),
			Community: "public",
			Version:   gs.Version2c,
			Timeout:   time.Duration(2) * time.Second,      
		}
		err := snmp.Connect()
		defer snmp.Conn.Close()

		oids := []string{".1.3.6.1.2.1.25.1.2.0"}		//HOST-RESOURCES-MIB::hrSystemDate.0
	
		_, err = snmp.Get(oids)
    	if err != nil {
        	utils.LogDebugf("Snmp connect failed: %s", err.Error())
		}
		return err

	}else {
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
