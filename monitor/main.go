package main

import (
	"fmt"
	"log"
	oracle "opms/monitor/oracle"
	mssql "opms/monitor/mssql"
	mysql "opms/monitor/mysql"
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
		
		//wg.Add(1)
		//GatherDisasterRecoveryStats(&wg)

		log.Println("循环结束！")

		time.Sleep(1 * time.Minute)
	}
	wg.Wait()

	//fmt.Println("cpu", x, time.Now().UnixNano()-start, "ns")
}


type Asset struct {
    Id 				int				`xorm:"int 'id'"`
    Asset_Type 		int		    	`xorm:"int 'asset_type'"`
    Host 			string		    `xorm:"int 'host'"`
    Port 			string		    `xorm:"int 'port'"`
    Alias 			string		    `xorm:"int 'alias'"`
    Inst_Name 		string		    `xorm:"int 'instance_name'"`
    Db_Name 		string		    `xorm:"int 'db_name'"`
    Username 		string			`xorm:"int 'username'"`
    Passwd 			string			`xorm:"int 'password'"`
}

func GatherAssetStats(wg *sync.WaitGroup) int {
	var assets []Asset
	sql := `select id, asset_type, host, port, alias, instance_name, db_name, username, password from pms_asset_config where status = 1 and is_delete = 0`
	err := db.SQL(sql).Find(&assets)
	if err != nil {
		log.Fatal(err)
	}

    //log.Print("Gather Asset Stats start.")
	var wga sync.WaitGroup
	for _, v := range assets {
		if v.Asset_Type == 1{
			wga.Add(1)
			go oracle.GenerateOracleStats(&wga, db, v.Id, v.Host, v.Port, v.Alias)
		}else if v.Asset_Type == 2{
			wga.Add(1)
			go mysql.GenerateMySQLStats(&wga, db, v.Id, v.Host, v.Port, v.Alias)
		}else if v.Asset_Type == 3{
			wga.Add(1)
			go mssql.GenerateMssqlStats(&wga, db, v.Id, v.Host, v.Port, v.Alias)
		}else if v.Asset_Type == 99{
			wga.Add(1)
			go mos.GenerateOSStats(&wga, db, v.Id, v.Host, v.Port, v.Alias)
		}
	}
	wga.Wait()

	(*wg).Done()

	//log.Print("Gather Asset Stats finished.")
	
	return 0
}


func GatherDisasterRecoveryStats(wg *sync.WaitGroup) int {
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
					d.is_shift
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
