package os

import (
	"log"
	"opms/monitor/utils"
	"time"
	"sync"

	"strings"
	"fmt"
	"bytes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/masterzen/winrm"

)

func GenerateWinStats(wg *sync.WaitGroup, mysql *xorm.Engine, os_id int, host string, port int, alias string, username string, password string) {
	//连接字符串
	endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, username, password)
	if err != nil {
		utils.LogDebugf("winrm.NewClient failed: %s", err.Error())
	}

	
	var stdout, stderr bytes.Buffer
	_, err = client.Run("ipconfig /all", &stdout, &stderr)
	if nil != err {
		utils.LogDebugf("Connect %s failed: %s", alias, err.Error())
		MoveToHistory(mysql, "pms_os_status", "os_id", os_id)

		sql := `insert into pms_os_status(os_id, host, alias, connect, created) 
		values(?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, os_id, host, alias, -1, time.Now().Unix())
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
		}

		AlertConnect(mysql, os_id)
	}else {
		log.Println("connect succeeded")
		//log.Printf("%s", stdout.String())
		//log.Printf("%d", code)
		
		//get os basic infomation
		GatherWinDiskInfo(client, mysql, os_id, host, alias)
		GatherWinDiskIOInfo(client, mysql, os_id, host, alias)
		GatherWinNetInfo(client, mysql, os_id, host, alias)
		GatherWinBasicInfo(client, mysql, os_id, host, alias)
		AlertConnect(mysql, os_id)
		
	}

	(*wg).Done()

}

func GatherWinBasicInfo(client *winrm.Client, mysql *xorm.Engine, os_id int, host string, alias string) error{
	connect := 1
	hostname, _ := GetWinHostname(client)
	kernel, _ := GetWinKernel(client)
	system_date, _ := GetWinSystemDate(client)
	system_uptime, _ := GetWinUptime(client)
	process, _ := GetWinProcess(client)

	load_1 := -1
	load_5 := -1
	load_15 := -1
	cpu_user_time := -1
	cpu_system_time := -1
	cpu_idle_time, _ := GetWinIdleCPU(client)

	swap_total := -1
	swap_avail := -1

	mem_free, _ := GetWinMemoryFree(client)
	mem_total, _ := GetWinMemoryTotal(client)
	mem_avail := -1
	mem_shared := -1
	mem_buffered := -1
	mem_cached := -1

	mem_available := mem_total - mem_free
	mem_usage_rate := mem_available*100/mem_total
	
	disk_io_reads_total := GetDiskIOReadTotal(mysql, os_id)
	disk_io_writes_total := GetDiskIOWriteTotal(mysql, os_id)
	net_in_bytes_total := GetNetInBytesTotal(mysql, os_id)
	net_out_bytes_total := GetNetOutBytesTotal(mysql, os_id)
	//fmt.Printf("disk_io_reads_total: %d\n", disk_io_reads_total)
	//fmt.Printf("disk_io_writes_total: %d\n", disk_io_writes_total)
	//fmt.Printf("net_in_bytes_total: %d\n", net_in_bytes_total)
	//fmt.Printf("net_out_bytes_total: %d\n", net_out_bytes_total)

	
	// storage result
	session := mysql.NewSession()
	defer session.Close()
	// add Begin() before any action
	err := session.Begin()

	//storage stats into pms_os_status
	MoveToHistory(mysql, "pms_os_status", "os_id", os_id)

	sql := `insert into pms_os_status(os_id, host, alias, connect, hostname, kernel, system_date, system_uptime, process, load_1, load_5, load_15, cpu_user_time, cpu_system_time, cpu_idle_time, swap_total, swap_avail,mem_total,mem_avail,mem_free,mem_shared, mem_buffered, mem_cached, mem_usage_rate, mem_available, disk_io_reads_total, disk_io_writes_total, net_in_bytes_total, net_out_bytes_total, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.Exec(sql, os_id, host, alias, connect, hostname, kernel, system_date, system_uptime, process, load_1, load_5, load_15, cpu_user_time, cpu_system_time, cpu_idle_time, swap_total, swap_avail,mem_total,mem_avail,mem_free,mem_shared, mem_buffered, mem_cached, mem_usage_rate, mem_available, disk_io_reads_total, disk_io_writes_total, net_in_bytes_total, net_out_bytes_total, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
		return err
	}

	// add Commit() after all actions
	err = session.Commit()

	return nil

}

func GatherWinDiskInfo(client *winrm.Client, mysql *xorm.Engine, os_id int, host string, alias string) error{

	var stdout, stderr bytes.Buffer
	_, err := client.Run("wmic LogicalDisk where DriveType=3 get DeviceID  /format:list", &stdout, &stderr)
	if nil != err {
		return err
	}else{
		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err = session.Begin()

		//storage stats into pms_os_disk
		MoveToHistory(mysql, "pms_os_disk", "os_id", os_id)

		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		for i := 0; i < len(s); i++ {
			if strings.Contains(s[i], "DeviceID="){
				driver := strings.Replace(s[i], "DeviceID=", "",1)
				//fmt.Printf("Driver: %s ", driver)
				total_size, _ := GetDriverSize(client, driver)
				//fmt.Printf("total_size: %d ", total_size)
				
				avail_size, _ := GetDriverFreeSize(client, driver)
				//fmt.Printf("avail_size: %d ", avail_size)

                used_size := total_size - avail_size
                
				used_rate := used_size * 100 / total_size
				
				sql := `insert into pms_os_disk(os_id, host, alias, mounted, total_size, used_size, avail_size, used_rate, created) 
						values(?,?,?,?,?,?,?,?,?)`
				_, err = mysql.Exec(sql, os_id, host, alias, driver, total_size, used_size, avail_size, used_rate, time.Now().Unix())
				if err != nil {
					log.Printf("%s: %s", sql, err.Error())
					err = session.Rollback()
					return err
				}
			}
		}

		// add Commit() after all actions
		err = session.Commit()
		
	}

	return nil
}

func GatherWinDiskIOInfo(client *winrm.Client, mysql *xorm.Engine, os_id int, host string, alias string) error{
	
	var stdout, stderr bytes.Buffer
	_, err := client.Run("wmic LogicalDisk where DriveType=3 get DeviceID  /format:list", &stdout, &stderr)
	if nil != err {
		return err
	}else{
		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err = session.Begin()

		//storage stats into pms_os_diskio
		MoveToHistory(mysql, "pms_os_diskio", "os_id", os_id)

		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		for i := 0; i < len(s); i++ {
			if strings.Contains(s[i], "DeviceID="){
				driver := strings.Replace(s[i], "DeviceID=", "",1)
				//fmt.Printf("Driver: %s ", driver)
				io_reads, _ := GetDiskReadBytesPerSec(client, driver)
				//fmt.Printf("io_reads: %d ", io_reads)
				
				io_writes, _ := GetDiskWriteBytesPerSec(client, driver)
				//fmt.Printf("io_writes: %d ", io_writes)

				
				sql := `insert into pms_os_diskio(os_id, host, alias, fdisk, disk_io_reads, disk_io_writes, created) 
						values(?,?,?,?,?,?,?)`
				_, err = mysql.Exec(sql, os_id, host, alias, driver, io_reads, io_writes, time.Now().Unix())
				if err != nil {
					log.Printf("%s: %s", sql, err.Error())
					err = session.Rollback()
					return err
				}
			}
		}

		// add Commit() after all actions
		err = session.Commit()
		
	}

	return nil
}

func GatherWinNetInfo(client *winrm.Client, mysql *xorm.Engine, os_id int, host string, alias string) error{

	var stdout, stderr bytes.Buffer
	_, err := client.Run("wmic path Win32_PerfFormattedData_Tcpip_NetworkInterface get Name /format:list", &stdout, &stderr)
	if nil != err {
		return err
	}else{
		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err = session.Begin()

		//storage stats into pms_os_net
		MoveToHistory(mysql, "pms_os_net", "os_id", os_id)

		s := strings.Split(stdout.String(),"Name=")
		//fmt.Println(s, len(s))

		for i := 0; i < len(s); i++ {
			if strings.Contains(s[i], "Network"){
				net := strings.Replace(strings.Trim(s[i], " "), "\n", "", -1)
				fmt.Printf("Net: %s\n", net)
				in_bytes, _ := GetBytesReceivedPerSec(client, net)
				fmt.Printf("in_bytes: %d\n", in_bytes)
				
				out_bytes, _ := GetBytesSentPerSec(client, net)
				fmt.Printf("out_bytes: %d\n", out_bytes)

				sql := `insert into pms_os_net(os_id, host, alias, if_descr, in_bytes, out_bytes, created) 
						values(?,?,?,?,?,?,?)`
				_, err = mysql.Exec(sql, os_id, host, alias, net, in_bytes, out_bytes, time.Now().Unix())
				if err != nil {
					log.Printf("%s: %s", sql, err.Error())
					err = session.Rollback()
					return err
				}
			}
		}

		// add Commit() after all actions
		err = session.Commit()
		
	}

	return nil
}



