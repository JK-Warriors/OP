package main

import (
	"fmt"
	"log"
	oracle "opms/monitor/oracle"
	mssql "opms/monitor/mssql"
	mysql "opms/monitor/mysql"
	alert "opms/monitor/alert"
	mos "opms/monitor/os"

	"os"
	"runtime"
	"sync"
	"time"
	"opms/monitor/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pythonsite/yamlConfig"
	"github.com/xormplus/xorm"
)

var db *xorm.Engine

type ServerConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
	Timeout  int
}

func InitSql() {
	currdir, _ := os.Getwd()
	yamlFile := currdir + "/etc/config.yml"

	config := yamlConfig.ConfigEngine{}
	err := config.Load(yamlFile)
	if err != nil {
		log.Fatalln("Config load error:", err)
	}

	serverconf := ServerConfig{}
	config.GetStruct("mysql", &serverconf)
	//fmt.Printf("%v", res)
	//utils.LogInfof("%v", res)

	host := serverconf.Host
	port := serverconf.Port
	user := serverconf.Username
	passwd := serverconf.Password
	dbname := serverconf.Dbname

	//引入xorm引擎
	db, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))

	if err != nil {
		log.Fatal(err)
	}

	db.ShowSQL(true)

	log.Println(db)
	//	db.NewSession().SqlMapClient().Execute()
	//log.Println(db.GetSqlMap("json_category-16-17"))
}

func main() {
	InitSql()

	num := runtime.NumCPU()
	DoTasks(num)

}

func DoTasks(x int) {
	runtime.GOMAXPROCS(x)
	var wg sync.WaitGroup
	//start := time.Now().UnixNano()

	for {
		log.Println("循环开始！")
		wg.Add(1)
		GatherAssetStats(&wg)
		
		wg.Add(1)
		GatherDisasterRecoveryStats(&wg)

		wg.Add(1)
		AlertMedia(&wg)
		//log.Println("循环结束！")

		time.Sleep(1 * time.Minute)
	}
	wg.Wait()

	//fmt.Println("cpu", x, time.Now().UnixNano()-start, "ns")
}


type Asset struct {
    Id 				int				`xorm:"int 'id'"`
    Asset_Type 		int		    	`xorm:"int 'asset_type'"`
    Host 			string		    `xorm:"varchar 'host'"`
    Protocol 		string		    `xorm:"varchar 'protocol'"`
    Port 			int		    	`xorm:"int 'port'"`
    Alias 			string		    `xorm:"varchar 'alias'"`
    Inst_Name 		string		    `xorm:"varchar 'instance_name'"`
    Db_Name 		string		    `xorm:"varchar 'db_name'"`
    Username 		string			`xorm:"varchar 'username'"`
    Passwd 			string			`xorm:"varchar 'password'"`
    Os_Type 		string			`xorm:"int 'os_type'"`
    Os_Protocol 	string			`xorm:"varchar 'os_protocol'"`
    Os_Port 		int				`xorm:"int 'os_port'"`
    Os_Username 	string			`xorm:"varchar 'os_username'"`
    Os_Password 	string			`xorm:"varchar 'os_password'"`
    Is_Alert 		int				`xorm:"int 'is_alert'"`
}

func GatherAssetStats(wg *sync.WaitGroup) int {
	//添加异常处理
	defer func() {
		if err := recover(); err != nil{
		   // 出现异常，继续
		   log.Printf("Error: %v", err)
		   (*wg).Done()
		}
	}()


	var assets []Asset
	sql := `select id, asset_type, host, protocol, port, alias, instance_name, db_name, username, password,
			os_type, os_protocol, os_port, os_username, os_password, is_alert
			from pms_asset_config where status = 1 and is_delete = 0`
	err := db.SQL(sql).Find(&assets)
	if err != nil {
		log.Fatal(err)
	}


    //log.Print("Gather Asset Stats start.")
	var wga sync.WaitGroup
	for _, v := range assets {
		if v.Asset_Type == 1{
			wga.Add(1)
			go oracle.GenerateOracleStats(&wga, db, v.Id, v.Host, v.Port, v.Alias, v.Is_Alert)
		}else if v.Asset_Type == 2{
			wga.Add(1)
			go mysql.GenerateMySQLStats(&wga, db, v.Id, v.Host, v.Port, v.Alias, v.Is_Alert)
		}else if v.Asset_Type == 3{
			wga.Add(1)
			go mssql.GenerateMssqlStats(&wga, db, v.Id, v.Host, v.Port, v.Alias, v.Is_Alert)
		}else if v.Asset_Type == 99{
			wga.Add(1)
			if v.Os_Protocol == "snmp" {		
				go mos.GenerateLinuxStats(&wga, db, v.Id, v.Host, v.Os_Port, v.Alias, v.Is_Alert)
			}else if v.Os_Protocol == "winrm" {
				go mos.GenerateWinStats(&wga, db, v.Id, v.Host, v.Os_Port, v.Alias, v.Os_Username, v.Os_Password, v.Is_Alert)
			}
		}
	}
	wga.Wait()

	(*wg).Done()

	//log.Print("Gather Asset Stats finished.")
	
	return 0
}


func GatherDisasterRecoveryStats(wg *sync.WaitGroup) int {
	//添加异常处理
	defer func() {
		if err := recover(); err != nil{
		   // 出现异常，继续
		   log.Printf("Error: %v", err)
		   (*wg).Done()
		}
	}()


	var dr []common.Dr
	sql := `select d.bs_id as id, 
					d.bs_name,
					d.asset_type,
					d.db_id_p, 
					pp.host as host_p,
					pp.port as port_p, 
					pp.alias as alias_p, 
					pp.instance_name as inst_name_p, 
					d.db_id_s, 
					ps.host as host_s,
					ps.port as port_s, 
					ps.alias as alias_s, 
					ps.instance_name as inst_name_s, 
					d.db_name, 
					d.is_shift, 
					d.is_switch,
					d.is_alert
				from pms_dr_config d
				left join pms_asset_config pp on d.db_id_p = pp.id
				left join pms_asset_config ps on d.db_id_s = ps.id
				where d.status = 1 and d.is_delete = 0`

	err := db.SQL(sql).Find(&dr)
	if err != nil {
		log.Fatal(err)
	}

	var wgb sync.WaitGroup
	for _, v := range dr {
		if v.Asset_Type == 1 {
			log.Println("获取Oracle容灾数据开始！")
			wgb.Add(1)
			go oracle.GenerateOracleDrStats(&wgb, db, v)
		}else if v.Asset_Type == 2 {
			log.Println("获取MySQL容灾数据开始！")
			wgb.Add(1)
			go mysql.GenerateMySQLDrStats(&wgb, db, v)
		}else if v.Asset_Type == 3 {
			log.Println("获取SQLServer容灾数据开始！")
			wgb.Add(1)
			go mssql.GenerateMssqlDrStats(&wgb, db, v)
		}
	}

	wgb.Wait()
	
	(*wg).Done()

	return 0
}


type Alert struct {
    Id 				int				`xorm:"int 'id'"`
    Asset_Id 		int		    	`xorm:"int 'asset_id'"`
    Name 			string		    `xorm:"varchar 'name'"`
    Severity 		string		    `xorm:"varchar 'severity'"`
    Templateid 		int		    	`xorm:"int 'templateid'"`
    Subject 		string		    `xorm:"varchar 'subject'"`
    Message 		string		    `xorm:"varchar 'message'"`
    Status 			int				`xorm:"int 'status'"`
    Send_Mail 					int				`xorm:"int 'send_mail'"`
    Send_Mail_List 				string			`xorm:"varchar 'send_mail_list'"`
    Send_Mail_Status 			int				`xorm:"int 'send_mail_status'"`
    Send_Mail_Retries 			int				`xorm:"int 'send_mail_retries'"`
    Send_Mail_Error 			string			`xorm:"varchar 'send_mail_error'"`
    Send_WeChat 				int				`xorm:"int 'send_wechat'"`
    Send_WeChat_Status 			int				`xorm:"int 'send_wechat_status'"`
    Send_WeChat_Retries 		int				`xorm:"int 'send_wechat_retries'"`
    Send_WeChat_Error 			string			`xorm:"varchar 'send_wechat_error'"`
    Send_SMS 					int				`xorm:"int 'send_sms'"`
    Send_SMS_List 				string			`xorm:"varchar 'send_sms_list'"`
    Send_SMS_Status 			int				`xorm:"int 'send_sms_status'"`
    Send_SMS_Retries 			int				`xorm:"int 'send_sms_retries'"`
    Send_SMS_Error 				string			`xorm:"varchar 'send_sms_error'"`
    Created 					int				`xorm:"int 'created'"`
}

func AlertMedia(wg *sync.WaitGroup) int {
	//添加异常处理
	defer func() {
		if err := recover(); err != nil{
		   // 出现异常，继续
		   log.Printf("Error: %v", err)
		   (*wg).Done()
		}
	}()

	
	var alerts []Alert
	sql := `select * from pms_alerts where status = 1`
	err := db.SQL(sql).Find(&alerts)
	if err != nil {
		log.Fatal(err)
	}

    //log.Print("AlertMedia start.")
	var wga sync.WaitGroup
	for _, v := range alerts {
		if v.Send_Mail == 1 && v.Send_Mail_Status == 0 {
			wga.Add(1)
			go alert.AlertEMail(&wga, db, v.Id, v.Send_Mail_Retries, v.Send_Mail_List, v.Subject, v.Created)
		}else if v.Send_WeChat == 1 && v.Send_WeChat_Status == 0 {
			wga.Add(1)
			go alert.AlertWeChat(&wga, db, v.Id, v.Send_Mail_Retries, v.Subject, v.Created)
		}else if v.Send_SMS == 1 && v.Send_SMS_Status == 0 {
			wga.Add(1)
			go alert.AlertSMS(&wga, db, v.Id, v.Send_Mail_Retries, v.Send_Mail_List, v.Subject, v.Created)
		}
	}
	wga.Wait()

	(*wg).Done()

	//log.Print("AlertMedia finished.")
	
	return 0
}