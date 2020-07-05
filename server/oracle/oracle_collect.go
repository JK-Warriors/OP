package oracle

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/godror/godror"
	"github.com/xormplus/xorm"
)

type Disaster struct {
	Id          int    `xorm:"int 'id'"`
	Bs_Name     string `xorm:"varchar(200) 'bs_name'"`
	Db_Id_P     int    `xorm:"int 'db_id_p'"`
	Db_Type_P   int    `xorm:"int 'db_type_p'"`
	Host_P      string `xorm:"varchar(20) 'host_p'"`
	Port_P      int    `xorm:"int 'port_p'"`
	Alias_P     string `xorm:"varchar(200) 'alias_p'"`
	Inst_Name_P string `xorm:"varchar(50) 'inst_name_p'	"`
	Db_Name_P   string `xorm:"varchar(50) 'db_name_p'"`
	Db_Id_S     int    `xorm:"int 'db_id_s'"`
	Db_Type_S   int    `xorm:"int 'db_type_s'"`
	Host_S      string `xorm:"varchar(20) 'host_s'"`
	Port_S      int    `xorm:"int 'port_s'"`
	Alias_S     string `xorm:"varchar(200) 'alias_s'"`
	Inst_Name_S string `xorm:"varchar(50) 'inst_name_s'"`
	Db_Name_S   string `xorm:"varchar(50) 'db_name_s'"`
	Is_Shift    int    `xorm:"int 'is_shift'"`
}

func Oracle_Collect(db *xorm.Engine) {

	var disaster []Disaster
	sql := `select b.id, 
					b.bs_name,
					d.db_id_p, 
					pp.db_type as db_type_p,
					pp.host as host_p,
					pp.port as port_p, 
					pp.alias as alias_p, 
					pp.instance_name as inst_name_p, 
					pp.db_name as db_name_p, 
					d.db_id_s, 
					ps.db_type as db_type_s,
					ps.host as host_s,
					ps.port as port_s, 
					ps.alias as alias_s, 
					ps.instance_name as inst_name_s, 
					ps.db_name as db_name_s, 
					d.is_shift
				from pms_business b 
				left join pms_disaster_config d on d.bs_id = b.id 
				left join pms_db_config pp on d.db_id_p = pp.id
				left join pms_db_config ps on d.db_id_s = ps.id
				where b.is_delete = 0`

	err := db.SQL(sql).Find(&disaster)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range disaster {
		if v.Db_Id_P > 0 {
			if v.Db_Type_P == 1 {
				log.Println("获取Oracle容灾数据开始！")

				if v.Db_Type_P == v.Db_Type_S {
					log.Println("获取容灾数据开始！")
					GenerateDisaster(db, v)
				} else {
					log.Printf("业务系统 %s 里配置的容灾数据库类型不一致！", v.Bs_Name)
				}

				log.Println("获取Oracle容灾数据结束！")
			}
		} else {
			log.Printf("业务系统 %s 没有配置容灾！", v.Bs_Name)
		}
	}

}

func GenerateDisaster(db *xorm.Engine, dis Disaster) {
	var pri_id int
	var sta_id int
	if dis.Is_Shift == 0 {
		pri_id = dis.Db_Id_P
		sta_id = dis.Db_Id_S
	} else {
		pri_id = dis.Db_Id_S
		sta_id = dis.Db_Id_P

	}

	dsn_p, err := GetDsn(db, pri_id, 1)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	dsn_s, err := GetDsn(db, sta_id, 1)
	if err != nil {
		log.Printf("GetDsn failed: %s", err.Error())
	}

	p_pri, _ := godror.ParseConnString(dsn_p)
	p_sta, _ := godror.ParseConnString(dsn_s)

	GenerateOracleBasic(db, p_pri, pri_id)
	GenerateOracleBasic(db, p_sta, sta_id)

	GeneratePrimary(db, p_pri, pri_id)
	GenerateStandby(db, p_pri, p_sta, sta_id)
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
