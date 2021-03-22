// Copyright 2021 YWY, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"io"
	"opms/modules/oramon/common"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/ziutek/mymysql/mysql"
)

//Log logger of project
var Log *logs.BeeLogger

func main() {
	// parse config file
	var confFile string
	flag.StringVar(&confFile, "c", "myMon.cfg", "myMon configure file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	if *version {
		fmt.Println(fmt.Sprintf("%10s: %s", "Version", Version))
		os.Exit(0)
	}

	conf, err := common.NewConfig(confFile)
	if err != nil {
		fmt.Printf("NewConfig Error: %s\n", err.Error())
		return
	}
	if conf.Base.LogDir != "" {
		err = os.MkdirAll(conf.Base.LogDir, 0755)
		if err != nil {
			fmt.Printf("MkdirAll Error: %s\n", err.Error())
			return
		}
	}

	// init log and other necessary
	Log = common.MyNewLogger(conf, common.CompatibleLog(conf))

	db, err := common.NewMySQLConnection(conf)
	if err != nil {
		fmt.Printf("NewMySQLConnection Error: %s\n", err.Error())
		return
	}
	defer func() { _ = db.Close() }()

	// start...
	Log.Info("Oracle Monitor for opms")
	go timeout()
	err = fetchData(conf, db)
	if err != nil && err != io.EOF {
		Log.Error("Error: %s", err.Error())
	}

	/*
		cfg := flag.String("c", "cfg.json", "configuration file")
		version := flag.Bool("v", false, "show version")
		check := flag.Bool("check", false, "check collector")

		flag.Parse()

		if *version {
			fmt.Printf("Open-Falcon %s version %s, build %s\n", BinaryName, Version, GitCommit)
			os.Exit(0)
		}

		if *check {
			funcs.CheckCollector()
			os.Exit(0)
		}

		g.ParseConfig(*cfg)

		if g.Config().Debug {
			g.InitLog("debug")
		} else {
			g.InitLog("info")
		}

		g.InitRootDir()
		g.InitLocalIp()
		g.InitRpcClients()

		funcs.BuildMappers()

		go cron.InitDataHistory()

		cron.ReportAgentStatus()
		cron.SyncMinePlugins()
		cron.SyncBuiltinMetrics()
		cron.SyncTrustableIps()
		cron.Collect()
	*/
	select {}

}

func timeout() {
	time.AfterFunc(TimeOut*time.Second, func() {
		Log.Error("Execute timeout")
		os.Exit(1)
	})
}

func fetchData(conf *common.Config, db mysql.Conn) (err error) {

	return
}
