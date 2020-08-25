package main

import (
	"context"
	//"crypto/tls"
	"database/sql"
	"fmt"
	//"io/ioutil"
	//"net/http"
	"os"
	"path"
	"strconv"
	//"strings"
	"time"
	//"bytes"
	//"reflect"
	"opms/agent/mysqld_exporter/collector"
	"opms/agent/mysqld_exporter/com"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/percona/exporter_shared"
	"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	//"github.com/prometheus/common/expfmt"
	"gopkg.in/alecthomas/kingpin.v2"
	//"gopkg.in/ini.v1"
	//"gopkg.in/yaml.v2"

	
	"github.com/pythonsite/yamlConfig"
)

// System variable params formatting.
// See: https://github.com/go-sql-driver/mysql#system-variables
const (
	sessionSettingsParam = `log_slow_filter=%27tmp_table_on_disk,filesort_on_disk%27`
	timeoutParam         = `lock_wait_timeout=%d`
)

var (
	showVersion = kingpin.Flag(
		"version",
		"Print version information.",
	).Default("false").Bool()
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address to listen on for web interface and telemetry.",
	).Default(":9104").String()
	metricPath = kingpin.Flag(
		"web.telemetry-path",
		"Path under which to expose metrics.",
	).Default("/metrics").String()
	timeoutOffset = kingpin.Flag(
		"timeout-offset",
		"Offset to subtract from timeout in seconds.",
	).Default("0.25").Float64()
	configMycnf = kingpin.Flag(
		"config.my-cnf",
		"Path to .my.cnf file to read MySQL credentials from.",
	).Default(path.Join(os.Getenv("HOME"), ".my.cnf")).String()
	exporterLockTimeout = kingpin.Flag(
		"exporter.lock_wait_timeout",
		"Set a lock_wait_timeout on the connection to avoid long metadata locking.",
	).Default("2").Int()
	exporterLogSlowFilter = kingpin.Flag(
		"exporter.log_slow_filter",
		"Add a log_slow_filter to avoid slow query logging of scrapes. NOTE: Not supported by Oracle MySQL.",
	).Default("false").Bool()
	exporterGlobalConnPool = kingpin.Flag(
		"exporter.global-conn-pool",
		"Use global connection pool instead of creating new pool for each http request.",
	).Default("true").Bool()
	exporterMaxOpenConns = kingpin.Flag(
		"exporter.max-open-conns",
		"Maximum number of open connections to the database. https://golang.org/pkg/database/sql/#DB.SetMaxOpenConns",
	).Default("3").Int()
	exporterMaxIdleConns = kingpin.Flag(
		"exporter.max-idle-conns",
		"Maximum number of connections in the idle connection pool. https://golang.org/pkg/database/sql/#DB.SetMaxIdleConns",
	).Default("3").Int()
	exporterConnMaxLifetime = kingpin.Flag(
		"exporter.conn-max-lifetime",
		"Maximum amount of time a connection may be reused. https://golang.org/pkg/database/sql/#DB.SetConnMaxLifetime",
	).Default("1m").Duration()
	collectAll = kingpin.Flag(
		"collect.all",
		"Collect all metrics.",
	).Default("false").Bool()

	dsn string
)


// scrapers lists all possible collection methods and if they should be enabled by default.
var scrapers = map[collector.Scraper]bool{
	collector.ScrapeGlobalStatus{}:                        true,
	collector.ScrapeGlobalVariables{}:                     false,
	collector.ScrapeSlaveStatus{}:                         false,
	collector.ScrapeProcesslist{}:                         false,
	collector.ScrapeTableSchema{}:                         false,
	collector.ScrapeInfoSchemaInnodbTablespaces{}:         false,
	collector.ScrapeInnodbMetrics{}:                       false,
	collector.ScrapeAutoIncrementColumns{}:                false,
	collector.ScrapeBinlogSize{}:                          false,
	collector.ScrapePerfTableIOWaits{}:                    false,
	collector.ScrapePerfIndexIOWaits{}:                    false,
	collector.ScrapePerfTableLockWaits{}:                  false,
	collector.ScrapePerfEventsStatements{}:                false,
	collector.ScrapePerfEventsWaits{}:                     false,
	collector.ScrapePerfFileEvents{}:                      false,
	collector.ScrapePerfFileInstances{}:                   false,
	collector.ScrapeUserStat{}:                            false,
	collector.ScrapeClientStat{}:                          false,
	collector.ScrapeTableStat{}:                           false,
	collector.ScrapeQueryResponseTime{}:                   false,
	collector.ScrapeEngineTokudbStatus{}:                  false,
	collector.ScrapeEngineInnodbStatus{}:                  false,
	collector.ScrapeHeartbeat{}:                           false,
	collector.ScrapeInnodbCmp{}:                           false,
	collector.ScrapeInnodbCmpMem{}:                        false,
	collector.ScrapeCustomQuery{Resolution: collector.HR}: false,
	collector.ScrapeCustomQuery{Resolution: collector.MR}: false,
	collector.ScrapeCustomQuery{Resolution: collector.LR}: false,
	collector.NewStandardGo():                             false,
	collector.NewStandardProcess():                        false,
}

// TODO Remove
var scrapersHr = map[collector.Scraper]struct{}{
	collector.ScrapeGlobalStatus{}:                        {},
	collector.ScrapeInnodbMetrics{}:                       {},
	collector.ScrapeCustomQuery{Resolution: collector.HR}: {},
}

// TODO Remove
var scrapersMr = map[collector.Scraper]struct{}{
	collector.ScrapeSlaveStatus{}:                         {},
	collector.ScrapeProcesslist{}:                         {},
	collector.ScrapePerfEventsWaits{}:                     {},
	collector.ScrapePerfFileEvents{}:                      {},
	collector.ScrapePerfTableLockWaits{}:                  {},
	collector.ScrapeQueryResponseTime{}:                   {},
	collector.ScrapeEngineInnodbStatus{}:                  {},
	collector.ScrapeInnodbCmp{}:                           {},
	collector.ScrapeInnodbCmpMem{}:                        {},
	collector.ScrapeCustomQuery{Resolution: collector.MR}: {},
}

// TODO Remove
var scrapersLr = map[collector.Scraper]struct{}{
	collector.ScrapeGlobalVariables{}:                     {},
	collector.ScrapeTableSchema{}:                         {},
	collector.ScrapeAutoIncrementColumns{}:                {},
	collector.ScrapeBinlogSize{}:                          {},
	collector.ScrapePerfTableIOWaits{}:                    {},
	collector.ScrapePerfIndexIOWaits{}:                    {},
	collector.ScrapePerfFileInstances{}:                   {},
	collector.ScrapeUserStat{}:                            {},
	collector.ScrapeTableStat{}:                           {},
	collector.ScrapePerfEventsStatements{}:                {},
	collector.ScrapeClientStat{}:                          {},
	collector.ScrapeInfoSchemaInnodbTablespaces{}:         {},
	collector.ScrapeEngineTokudbStatus{}:                  {},
	collector.ScrapeHeartbeat{}:                           {},
	collector.ScrapeCustomQuery{Resolution: collector.LR}: {},
}

 
type ServerConfig struct {
	Host 		string
	Port  		int
	Username  	string
	Password  	string
	Dbname  	string
	Timeout  	int
}


func GetDsn() string{
	
    currdir, _ := os.Getwd()
	yamlFile := currdir + "/etc/config.yml"

    config := yamlConfig.ConfigEngine{}
    err := config.Load(yamlFile)
	if err != nil {
		log.Fatalln("Config load error:", err)
	}

	serverconf := ServerConfig{}
    res := config.GetStruct("mysql",&serverconf)
    //fmt.Printf("%v",res)
	log.Debugln(res)
	host := serverconf.Host
	port := strconv.Itoa(serverconf.Port)
	username := serverconf.Username
	password := serverconf.Password
	dbname := serverconf.Dbname

	dsn := username + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?timeout=5s"
	//fmt.Println("dsn: " + dsn)
	log.Debugln(dsn)
	return dsn
}


func enabledScrapers(scraperFlags map[collector.Scraper]*bool) (all, hr, mr, lr []collector.Scraper) {
	for scraper, enabled := range scraperFlags {
		if *collectAll || *enabled {
			if _, ok := scrapers[scraper]; ok {
				all = append(all, scraper)
			}
			if _, ok := scrapersHr[scraper]; ok {
				hr = append(hr, scraper)
			}
			if _, ok := scrapersMr[scraper]; ok {
				mr = append(mr, scraper)
			}
			if _, ok := scrapersLr[scraper]; ok {
				lr = append(lr, scraper)
			}
		}
	}

	return all, hr, mr, lr
}

func newDB(dsn string) (*sql.DB, error) {
	// Validate DSN, and open connection pool.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(*exporterMaxOpenConns)
	db.SetMaxIdleConns(*exporterMaxIdleConns)
	db.SetConnMaxLifetime(*exporterConnMaxLifetime)

	return db, nil
}


func init() {
	prometheus.MustRegister(version.NewCollector("mysqld_exporter"))
}

func newHandler(maindb *sql.DB, dbid int, dbinst *sql.DB, metrics collector.Metrics, scrapers []collector.Scraper, defaultGatherer bool) {
		filteredScrapers := scrapers

		//ctx := r.Context()
		ctx, _ := context.WithCancel(context.Background())

		if dbinst == nil || maindb == nil {
			log.Fatalln("Error opening connection to database: ", strconv.Itoa(dbid))
			return
		}

		//var err error
		rows, err := maindb.Query("select key_, label, item_id, value_type from pms_items where obj_id = ?", dbid)
		if err != nil {
			log.Infoln("Get item template error: ", err)
		}

		registry := prometheus.NewRegistry()
		registry.MustRegister(collector.New(ctx, dbinst, metrics, filteredScrapers))

		//count,_ :=prometheus.GetMetricWith("")

		gatherers := prometheus.Gatherers{}
		if defaultGatherer {
			gatherers = append(gatherers, prometheus.DefaultGatherer)
		}
		gatherers = append(gatherers, registry)

		gathering, err := gatherers.Gather()
		if err != nil {
			fmt.Println(err)
			return
		}


		m0 := make(map[string]int)
		m1 := make(map[string]int)
		var(key string
			label string
			item_id	int
			value_type int
		)
		var key_label string
		for rows.Next() { 
			err = rows.Scan(&key, &label, &item_id, &value_type)
			if err != nil {
				log.Infoln(err)
			}
			key_label = key + "_" + label
			m0[key_label] = item_id
			m1[key_label] = value_type
		}
		defer rows.Close()

		
		gather_time := time.Now().Unix()
		fmt.Printf("gather_time: %d\n",gather_time)
		//out := &bytes.Buffer{}
		for _, mf := range gathering {
			//fmt.Println(reflect.TypeOf(mf))   type: io_prometheus_client
			//fmt.Print(mf.GetName() + "\n")
			//fmt.Print(mf.GetHelp() + "\n")
			//fmt.Print(mf.GetMetric())
			//fmt.Print("\n")

			mfname :=  mf.GetName()


			//if mfname == "mysql_up" {
			fmt.Print(mfname + "\n")
			
			mr := mf.GetMetric()
			for _,m := range mr{
				// 获取LabelPair类型值
				ml := m.GetLabel()
				var lvalue string
				for _,mli := range ml{
					//lname := mli.GetName()
					lvalue = mli.GetValue()

					//fmt.Print(lname + "\n")
					fmt.Print(lvalue + "\n")
				}

				k_label := mfname + "_" + lvalue
				fmt.Print("key_label: " + k_label + "\n")
				itemid := m0[k_label]
				val_type := m1[k_label]
				fmt.Print(itemid)
				fmt.Print("\n")
				fmt.Print(val_type)
				fmt.Print("\n")

				if itemid > 0 {
					//根据值的类型取值
					var mr_value float64
					if val_type == 1 {
						// 获取Counter类型值
						mr_value = m.GetCounter().GetValue()
						fmt.Printf("%f\n", mr_value)

					}else if val_type == 2 {
						// 获取Gauge类型值
						mr_value = m.GetGauge().GetValue()
						fmt.Printf("%f\n", mr_value)

					}else if val_type == 3 {
						// 获取Untyped类型值
						mr_value = m.GetUntyped().GetValue()
						fmt.Printf("%f\n", mr_value)
					}else if val_type == 4 {
						// 获取Histogram类型值
						//mh := m.GetSummary()
						//fmt.Print(mh)
						//fmt.Print("\n")
					}else if val_type == 5 {
						// 获取Summary类型值
						//ms := m.GetHistogram()
						//fmt.Print(ms)
						//fmt.Print("\n")
					}

					//保存结果到maindb中
					_, err = maindb.Exec("INSERT INTO pms_item_data(itemid,time,value)VALUES (?,?,?)", itemid, gather_time, mr_value)
					if err != nil{
						fmt.Println("insert failed,",err)
					}
				}

			}

			//}
			

			// if _, err := expfmt.MetricFamilyToText(out, mf); err != nil {
			// 	panic(err)
			// }
		}
		//fmt.Print(out.String())
		fmt.Println("----------")
}




func main() {
	// Generate ON/OFF flags for all scrapers.
	scraperFlags := map[collector.Scraper]*bool{}
	for scraper, enabledByDefault := range scrapers {
		f := kingpin.Flag(
			"collect."+scraper.Name(),
			scraper.Help(),
		).Default(strconv.FormatBool(enabledByDefault)).Bool()

		scraperFlags[scraper] = f
	}

	// Parse flags.
	kingpin.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("mysqld_exporter"))
		os.Exit(0)
	}

	

	log.Infoln("Starting mysqld_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	// Get Server Center dsn
	dsn := GetDsn()
	log.Infoln("dsn: ", dsn)

	// Open global connection pool if requested.
	var maindb *sql.DB
	var err error
	if *exporterGlobalConnPool {
		maindb, err = newDB(dsn)
		if err != nil {
			log.Fatalln("Error opening connection to main database:", err)
			return
		}
		
		err = maindb.Ping()
		if err != nil {
			log.Fatalln("Error opening connection to main database:", err)
			return
		}

		defer maindb.Close()
	}
	

	all, _, _, _ := enabledScrapers(scraperFlags)

	var (dbid	int
		 host	string
		 port	int
		 username string
		 password string
		 dbname string
	)
	var mysql_dsn string
	var dbinst *sql.DB
	for {
		log.Infoln("check mysql controller start.")
		rows, err := maindb.Query("select id, host, port, username, password, db_name from pms_asset_config where asset_type = 2 and status = 1 and is_delete = 0")
		if err != nil {
			log.Infoln("No mysql exists.")
		}else{
			for rows.Next() { 
				err = rows.Scan(&dbid, &host, &port, &username, &password, &dbname)
				if err != nil {
					log.Infoln(err)
				}
				mysql_dsn = username + ":" + password + "@(" + host + ":" + strconv.Itoa(port) + ")/" + dbname + "?timeout=5s"
				log.Infoln(mysql_dsn)
				dbinst, err = newDB(mysql_dsn)
				if err != nil {
					log.Errorln("Error opening connection to mysql instance: ", err)
					com.ErrorConnection(dbid, maindb)
					continue
				}else{
					err = dbinst.Ping()
					if err != nil {
						log.Errorln("Error opening connection to mysql instance: ", err)
						com.ErrorConnection(dbid, maindb)
						continue
					}
					newHandler(maindb, dbid, dbinst, collector.NewMetrics(""), all, true)
				}
				defer dbinst.Close()

			}
		}
		defer rows.Close()
		log.Infoln("check mysql controller finished.")
		
		time.Sleep(1 * time.Minute)
	}
	


	
	
}

