package main

import (
	"fmt"
	"log"
	oracle "opms/server/oracle"
	"os"
	"runtime"
	"sync"
	"time"

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

	for {
		oracle.Oracle_Collect(db)
		time.Sleep(1 * time.Minute)
	}
}

func DoTask(wg *sync.WaitGroup) int {
	n := 2
	for i := 0; i < 20; i++ {
		for j := 0; j < 1000; j++ {
			if n > 10000 {
				n = n - 10000
			} else {
				n++
			}
		}
	}
	(*wg).Done()
	return n
}
func DoTasks(x int) {
	runtime.GOMAXPROCS(x)
	var wg sync.WaitGroup
	start := time.Now().UnixNano()
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go DoTask(&wg)
	}

	wg.Wait()
	fmt.Println("cpu", x, time.Now().UnixNano()-start, "ns")
}
