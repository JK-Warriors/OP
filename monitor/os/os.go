package os

import (
	"log"
	"opms/monitor/utils"
	"time"
	"sync"
	"strconv"

	"bufio"
	"io"
	"os/exec"
	"strings"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	gs "github.com/soniah/gosnmp"
	
)
var snmp *gs.GoSNMP

func GenerateLinuxStats(wg *sync.WaitGroup, mysql *xorm.Engine, os_id int, host string, port int, alias string) {
	//连接字符串
    snmp = &gs.GoSNMP{
        Target:    host,
        Port:      uint16(port),
        Community: "public",
        Version:   gs.Version2c,
        Timeout:   time.Duration(2) * time.Second,      
    }
    err := snmp.Connect()
	defer snmp.Conn.Close()
	//log.Printf("%v", snmp)
	
	_, err = GetSystemDate(snmp)
    if err != nil {
		utils.LogDebugf("Connect %s failed: %s", alias, err.Error())
		MoveToHistory(mysql, "pms_os_status", "os_id", os_id)

		sql := `insert into pms_os_status(os_id, host, alias, connect, created) 
		values(?,?,?,?,?)`
		_, err = mysql.Exec(sql, os_id, host, alias, -1, time.Now().Unix())
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
		}

		AlertConnect(mysql, os_id)
    }else {
		log.Println("connect succeeded")
		
		//get os basic infomation
		GatherLinuxDiskInfo(snmp, mysql , os_id, host, alias)
		GatherLinuxDiskIOInfo(mysql , os_id, host, alias)
		GatherLinuxNetInfo(snmp, mysql , os_id, host, alias)
		GatherLinuxBasicInfo(snmp, mysql , os_id, host, alias, port)
		
		GatherOSStatus(mysql, os_id, host)

		AlertConnect(mysql, os_id)
		
	}

	(*wg).Done()

}

func GatherLinuxBasicInfo(snmp *gs.GoSNMP, mysql *xorm.Engine, os_id int, host string, alias string, port int) error{

	connect := 1
	hostname, err := GetHostname(snmp)
	kernel, err := GetKernel(snmp)
	system_date, err := GetSystemDate(snmp)
	system_uptime, err := GetUptime(snmp)
	process, err := GetProcess(snmp)

	var load_1 float64
	var load_2 float64
	var load_3 float64
	oids := []string{".1.3.6.1.4.1.2021.10.1.3.1",			//UCD-SNMP-MIB::laLoad.1
					".1.3.6.1.4.1.2021.10.1.3.2",		//UCD-SNMP-MIB::laLoad.2
					".1.3.6.1.4.1.2021.10.1.3.3",		//UCD-SNMP-MIB::laLoad.3
	}
	result, err := GetSnmpStringByOids(snmp, oids) 
	if err != nil{
		utils.LogDebugf("GetSnmpStringByOids err: %s", err.Error())
		return err
	}else{
		load_1_str := result[0]
		load_1, err = strconv.ParseFloat(load_1_str, 64)
		if err != nil{
			load_1 = -1.0
		}

		load_2_str := result[1]
		load_2, err = strconv.ParseFloat(load_2_str, 64)
		if err != nil{
			load_2 = -1.0
		}

		load_3_str := result[2]
		load_3, err = strconv.ParseFloat(load_3_str, 64)
		if err != nil{
			load_3 = -1.0
		}
	}

	oids = []string{".1.3.6.1.4.1.2021.11.9.0",		//UCD-SNMP-MIB::ssCpuUser.0
					 ".1.3.6.1.4.1.2021.11.10.0",		//UCD-SNMP-MIB::ssCpuSystem.0
					 ".1.3.6.1.4.1.2021.11.11.0",		//UCD-SNMP-MIB::ssCpuIdle
					 ".1.3.6.1.4.1.2021.4.3.0",			//UCD-SNMP-MIB::memTotalSwap.0
					 ".1.3.6.1.4.1.2021.4.4.0",			//UCD-SNMP-MIB::memAvailSwap.0
					 ".1.3.6.1.4.1.2021.4.5.0",			//UCD-SNMP-MIB::memTotalReal.0
					 ".1.3.6.1.4.1.2021.4.6.0",			//UCD-SNMP-MIB::memAvailReal.0
					 ".1.3.6.1.4.1.2021.4.11.0",		//UCD-SNMP-MIB::memTotalFree.0
					 ".1.3.6.1.4.1.2021.4.13.0",		//UCD-SNMP-MIB::memShared.0
					 ".1.3.6.1.4.1.2021.4.14.0",		//UCD-SNMP-MIB::memBuffer.0
					 ".1.3.6.1.4.1.2021.4.15.0",		//UCD-SNMP-MIB::memCached.0
					}
	var cpu_user_time int64 = -1	
	var cpu_system_time int64 = -1	
	var cpu_idle_time int64 = -1	
	var swap_total int64 = -1	
	var swap_avail int64 = -1	
	var mem_total int64 = -1	
	var mem_avail int64 = -1	
	var mem_free int64 = -1	
	var mem_shared int64 = -1	
	var mem_buffered int64 = -1	
	var mem_cached int64 = -1	

	result2, err := GetSnmpInt64ByOids(snmp, oids) 
	if err != nil{
		utils.LogDebugf("GetSnmpInt64ByOids err: %s", err.Error())
		return err
	}else{
		cpu_user_time = result2[0]
		cpu_system_time = result2[1]
		cpu_idle_time = result2[2]
		swap_total = result2[3]
		swap_avail = result2[4]
		mem_total = result2[5]
		mem_avail = result2[6]
		mem_free = result2[7]
		mem_shared = result2[8]
		mem_buffered = result2[9]
		mem_cached = result2[10]
	}

	var mem_available int64 = -1
	if mem_total != -1 && mem_free != -1 && mem_buffered != -1 && mem_cached != -1 {
		mem_available = mem_total - mem_free - mem_buffered - mem_cached
	}

	var mem_usage_rate float64 = -1.00
	if mem_available != -1 && mem_total != -1 {
		mem_usage_rate = float64(mem_available)*100/float64(mem_total)
	}
	
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
	err = session.Begin()

	//storage stats into pms_os_status
	MoveToHistory(mysql, "pms_os_status", "os_id", os_id)

	sql := `insert into pms_os_status(os_id, host, alias, connect, hostname, kernel, system_date, system_uptime, process, load_1, load_5, load_15, cpu_user_time, cpu_system_time, cpu_idle_time, swap_total, swap_avail,mem_total,mem_avail,mem_free,mem_shared, mem_buffered, mem_cached, mem_usage_rate, mem_available, disk_io_reads_total, disk_io_writes_total, net_in_bytes_total, net_out_bytes_total, created) 
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.Exec(sql, os_id, host, alias, connect, hostname, kernel, system_date, system_uptime, process, load_1, load_2, load_3, cpu_user_time, cpu_system_time, cpu_idle_time, swap_total, swap_avail,mem_total,mem_avail,mem_free,mem_shared, mem_buffered, mem_cached, mem_usage_rate, mem_available, disk_io_reads_total, disk_io_writes_total, net_in_bytes_total, net_out_bytes_total, time.Now().Unix())
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
		err = session.Rollback()
		return err
	}

	// add Commit() after all actions
	err = session.Commit()
	
	return err

}

func GatherLinuxDiskInfo(snmp *gs.GoSNMP, mysql *xorm.Engine, os_id int, host string, alias string) error{
	
	oids := ".1.3.6.1.4.1.2021.9.1.1"
	result, err := snmp.WalkAll(oids)
	if err != nil {
		fmt.Printf("Walk Error: %v\n", err)
	}else {
		
		// storage result
		session := mysql.NewSession()
		defer session.Close()
		// add Begin() before any action
		err = session.Begin()

		//storage stats into pms_os_disk
		MoveToHistory(mysql, "pms_os_disk", "os_id", os_id)

		for _, v := range result {
			idx := v.Value.(int)
			
			oids := []string{fmt.Sprintf(".1.3.6.1.4.1.2021.9.1.2.%d", idx),  //UCD-SNMP-MIB::dskPath
							fmt.Sprintf(".1.3.6.1.4.1.2021.9.1.3.%d", idx),		//UCD-SNMP-MIB::dskDevice
			}
			
			var dskPath string
			var dskDevice string
    		result, err := GetSnmpStringByOids(snmp, oids)
			if err != nil {
				utils.LogDebugf("GetSnmpStringByOids err: %s", err.Error())
				return err
			}else{
				dskPath = result[0]
				dskDevice = result[1]
				fmt.Println("dskPath: ", dskPath)
				fmt.Println("dskDevice: ", dskDevice)
			}

			oids = []string{fmt.Sprintf(".1.3.6.1.4.1.2021.9.1.6.%d", idx),   //UCD-SNMP-MIB::dskTotal
							fmt.Sprintf(".1.3.6.1.4.1.2021.9.1.7.%d", idx),   //UCD-SNMP-MIB::dskAvail
							fmt.Sprintf(".1.3.6.1.4.1.2021.9.1.8.%d", idx),   //UCD-SNMP-MIB::dskUsed
							fmt.Sprintf(".1.3.6.1.4.1.2021.9.1.9.%d", idx),   //UCD-SNMP-MIB::dskPercent
							fmt.Sprintf(".1.3.6.1.4.1.2021.9.1.10.%d", idx),   //UCD-SNMP-MIB::dskPercentNode
			}


    		result2, err := GetSnmpInt64ByOids(snmp, oids)
			if err != nil {
				utils.LogDebugf("GetSnmpInt64ByOids err: %s", err.Error())
				return err
			}else{
				dskTotal := result2[0]
				dskAvail := result2[1]
				dskUsed := result2[2]
				dskPercent  := result2[3]
				dskPercentNode  := result2[4]
				
				fmt.Println(reflect.TypeOf(dskPercent) )
				
				fmt.Println("dskTotal: ", dskTotal)
				fmt.Println("dskAvail: ", dskAvail)
				fmt.Println("dskUsed: ", dskUsed)
				fmt.Println("dskPercent: ", dskPercent)
				fmt.Println("dskPercentNode: ", dskPercentNode)
				
				
				sql := `insert into pms_os_disk(os_id, host, alias, mounted, device, total_size, used_size, avail_size, used_rate, node_rate, created) 
				values(?,?,?,?,?,?,?,?,?,?,?)`
				_, err = mysql.Exec(sql, os_id, host, alias, dskPath, dskDevice, dskTotal, dskUsed, dskAvail, dskPercent, dskPercentNode, time.Now().Unix())
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



func GatherLinuxDiskIOInfo(mysql *xorm.Engine, os_id int, host string, alias string) error{
	s_cmd := `/usr/bin/snmptable -v1 -c public ` + host + ` diskIOTable |grep -ivE "ram|loop|md|SNMP table|diskIOIndex|dm-|sr0" | grep -v '^$'`

	cmd := exec.Command("/bin/bash", "-c", s_cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("cmd StdoutPipe error: %s", err.Error())
		return nil
	}
	cmd.Start()
 
	// storage result
	session := mysql.NewSession()
	defer session.Close()
	// add Begin() before any action
	err = session.Begin()

	//storage stats into pms_os_diskio
	MoveToHistory(mysql, "pms_os_diskio", "os_id", os_id)

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}

		s := strings.Fields(line)
		//fmt.Println(s, len(s))
		fdisk :=s[1]
		disk_io_reads,err := strconv.ParseInt(s[4],10,64)
		disk_io_writes,err :=strconv.ParseInt(s[5],10,64)
		
		//计算差值
		disk_io_reads_old := GetDiskIOReadsOld(mysql, os_id, fdisk)
		disk_io_writes_old := GetDiskIOWritesOld(mysql, os_id, fdisk)
		t_old := GetDiskIOTime(mysql, os_id, fdisk)
		
		var disk_io_reads_new = -1
		var disk_io_writes_new = -1
		t_now := time.Now().Unix()
		
		if(disk_io_reads_old > -1 && t_old > -1 && t_now > t_old){
			disk_io_reads_new = int(disk_io_reads - disk_io_reads_old)/int(t_now - t_old)
		}
		if(disk_io_writes_old > -1 && t_old > -1 && t_now > t_old){
			disk_io_writes_new = int(disk_io_writes - disk_io_writes_old)/int(t_now - t_old)
		}

		sql := `insert into pms_os_diskio(os_id, host, alias, fdisk, disk_io_reads, disk_io_writes, created) 
				values(?,?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, os_id, host, alias, fdisk, disk_io_reads_new, disk_io_writes_new, t_now)
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			err = session.Rollback()
			return err
		}
		
		//保留本次获取的值到temp表
		sql = `delete from pms_os_diskio_tmp where os_id = ? and fdisk = ?`
		_, err = mysql.Exec(sql, os_id, fdisk)
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			err = session.Rollback()
			return err
		}
		
		sql = `insert into pms_os_diskio_tmp(os_id, host, alias, fdisk, disk_io_reads, disk_io_writes, created) 
				values(?,?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, os_id, host, alias, fdisk, disk_io_reads, disk_io_writes, t_now)
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			err = session.Rollback()
			return err
		}

	}

	// add Commit() after all actions
	err = session.Commit()
	return nil
}

func GatherLinuxNetInfo(snmp *gs.GoSNMP, mysql *xorm.Engine, os_id int, host string, alias string) error{
	// storage result
	session := mysql.NewSession()
	defer session.Close()
	// add Begin() before any action
	err := session.Begin()

	//storage stats into pms_os_diskio
	MoveToHistory(mysql, "pms_os_net", "os_id", os_id)
	
	nets, _ := GetNetDescr(host)
	
	for k, net := range nets {
		
		in_bytes, _ := GetNetBytesIn(host, k)
		out_bytes, _ := GetNetBytesOut(host, k)
		//fmt.Printf("%s bytes in: %d\n", net, in_bytes)
		//fmt.Printf("%s bytes in: %d\n", net, out_bytes)
		
		
		//计算差值
		in_bytes_old := GetNetInBytesOld(mysql, os_id, net)
		out_bytes_old := GetNetOutBytesOld(mysql, os_id, net)
		t_old := GetNetTime(mysql, os_id, net)
		
		var in_bytes_new = -1
		var out_bytes_new = -1
		t_now := time.Now().Unix()
		
		if(in_bytes_old > -1 && t_old > -1 && t_now > t_old){
			in_bytes_new = int(in_bytes - in_bytes_old)/int(t_now - t_old)
		}
		if(out_bytes_old > -1 && t_old > -1 && t_now > t_old){
			out_bytes_new = int(out_bytes - out_bytes_old)/int(t_now - t_old)
		}
		
		sql := `insert into pms_os_net(os_id, host, alias, if_descr, in_bytes, out_bytes, created) 
				values(?,?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, os_id, host, alias, net, in_bytes_new, out_bytes_new, t_now)
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			err = session.Rollback()
			return err
		}
		
		
		//保留本次获取的值到temp表
		sql = `delete from pms_os_net_tmp where os_id = ? and if_descr = ?`
		_, err = mysql.Exec(sql, os_id, net)
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			err = session.Rollback()
			return err
		}
		
		sql = `insert into pms_os_net_tmp(os_id, host, alias, if_descr, in_bytes, out_bytes, created) 
				values(?,?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, os_id, host, alias, net, in_bytes, out_bytes, t_now)
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			err = session.Rollback()
			return err
		}
	}
	// add Commit() after all actions
	err = session.Commit()

	return nil
}



