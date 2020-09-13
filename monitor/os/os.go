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

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	gs "github.com/soniah/gosnmp"
	
)

func GenerateLinuxStats(wg *sync.WaitGroup, mysql *xorm.Engine, os_id int, host string, port int, alias string) {
	//连接字符串
    snmp := &gs.GoSNMP{
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
		GatherLinuxDiskInfo(mysql , os_id, host, alias)
		GatherLinuxDiskIOInfo(mysql , os_id, host, alias)
		GatherLinuxNetInfo(snmp, mysql , os_id, host, alias)
		GatherLinuxBasicInfo(snmp, mysql , os_id, host, alias)
		AlertConnect(mysql, os_id)
		
	}

	(*wg).Done()

}

func GatherLinuxBasicInfo(snmp *gs.GoSNMP, mysql *xorm.Engine, os_id int, host string, alias string) error{

	connect := 1
	oids := []string{".1.3.6.1.2.1.1.5.0",			//SNMPv2-MIB::sysName.0		hostname
					 ".1.3.6.1.2.1.1.1.0",			//SNMPv2-MIB::sysDescr.0	kernel
					 ".1.3.6.1.2.1.25.1.2.0",		//HOST-RESOURCES-MIB::hrSystemDate.0	
					 ".1.3.6.1.2.1.25.1.1.0",		//HOST-RESOURCES-MIB::hrSystemUptime.0	
					 ".1.3.6.1.2.1.25.1.6.0",		//HOST-RESOURCES-MIB::hrSystemProcesses.0	
					 ".1.3.6.1.4.1.2021.10.1.3.1",		//UCD-SNMP-MIB::laLoad.1	The 1,5 and 10 minute load averages (one per row).
					 ".1.3.6.1.4.1.2021.10.1.3.2",		//UCD-SNMP-MIB::laLoad.2
					 ".1.3.6.1.4.1.2021.10.1.3.3",		//UCD-SNMP-MIB::laLoad.3
					 ".1.3.6.1.4.1.2021.11.9.0",		//UCD-SNMP-MIB::ssCpuUser.0
					 ".1.3.6.1.4.1.2021.11.10.0",		//UCD-SNMP-MIB::ssCpuSystem.0
					 ".1.3.6.1.4.1.2021.4.3.0",			//UCD-SNMP-MIB::memTotalSwap.0
					 ".1.3.6.1.4.1.2021.4.4.0",			//UCD-SNMP-MIB::memAvailSwap.0
					 ".1.3.6.1.4.1.2021.4.5.0",			//UCD-SNMP-MIB::memTotalReal.0
					 ".1.3.6.1.4.1.2021.4.6.0",			//UCD-SNMP-MIB::memAvailReal.0
					 ".1.3.6.1.4.1.2021.4.11.0",		//UCD-SNMP-MIB::memTotalFree.0
					 ".1.3.6.1.4.1.2021.4.13.0",		//UCD-SNMP-MIB::memShared.0
					 ".1.3.6.1.4.1.2021.4.14.0",		//UCD-SNMP-MIB::memBuffer.0
					 ".1.3.6.1.4.1.2021.4.15.0",		//UCD-SNMP-MIB::memCached.0
					}
	
    result, err := GetSnmpValueByOids(snmp, oids)
    if err != nil {
		utils.LogDebugf("GetSnmpValueByOids err: %s", err.Error())
		return err
    }else{
		hostname := result[0]
		kernel := result[1]
		system_date, err := GetSystemDate(snmp)
		system_uptime, err := GetUptime(snmp)

		process, err := strconv.Atoi(result[4])
		if err != nil {
			process = -1
		}

		load_1, err := strconv.ParseFloat(result[5],64)
		if err != nil {
			load_1 = -1.00
		}

		load_5, err := strconv.ParseFloat(result[6],64)
		if err != nil {
			load_5 = -1.00
		}

		load_15, err := strconv.ParseFloat(result[7],64)
		if err != nil {
			load_15 = -1.00
		}


		cpu_user_time, err := strconv.Atoi(result[8])
		if err != nil {
			cpu_user_time = -1
		}

		cpu_system_time, err := strconv.Atoi(result[9])
		if err != nil {
			cpu_system_time = -1
		}

		cpu_idle_time := -1
		if cpu_user_time != -1 && cpu_system_time == -1 {
			cpu_idle_time = 100 - cpu_user_time - cpu_system_time
		}

		swap_total, err := strconv.Atoi(result[10])
		if err != nil {
			swap_total = -1
		}
		swap_avail, err := strconv.Atoi(result[11])
		if err != nil {
			swap_avail = -1
		}
		mem_total, err := strconv.Atoi(result[12])
		if err != nil {
			mem_total = -1
		}
		mem_avail, err := strconv.Atoi(result[13])
		if err != nil {
			mem_avail = -1
		}
		mem_free, err := strconv.Atoi(result[14])
		if err != nil {
			mem_free = -1
		}
		mem_shared, err := strconv.Atoi(result[15])
		if err != nil {
			mem_shared = -1
		}
		mem_buffered, err := strconv.Atoi(result[16])
		if err != nil {
			mem_buffered = -1
		}
		mem_cached, err := strconv.Atoi(result[17])
		if err != nil {
			mem_cached = -1
		}

		mem_available := -1
		if mem_total != -1 && mem_free == -1 && swap_avail == -1 {
			mem_available = mem_total - mem_free + swap_avail
		}

		mem_usage_rate := -1.00
		if mem_available != -1 && mem_total == -1 {
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
		_, err = mysql.Exec(sql, os_id, host, alias, connect, hostname, kernel, system_date, system_uptime, process, load_1, load_5, load_15, cpu_user_time, cpu_system_time, cpu_idle_time, swap_total, swap_avail,mem_total,mem_avail,mem_free,mem_shared, mem_buffered, mem_cached, mem_usage_rate, mem_available, disk_io_reads_total, disk_io_writes_total, net_in_bytes_total, net_out_bytes_total, time.Now().Unix())
		if err != nil {
			log.Printf("%s: %s", sql, err.Error())
			err = session.Rollback()
			return err
		}
	
		// add Commit() after all actions
		err = session.Commit()
		
		return err

	}

}

func GatherLinuxDiskInfo(mysql *xorm.Engine, os_id int, host string, alias string) error{
	s_cmd := `/usr/bin/snmpdf -v1 -c public ` + host + `|grep -E "/"|grep -vE "/boot"|grep -vE "DVD"`
	//fmt.Println(s_cmd)
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

	//storage stats into pms_os_disk
	MoveToHistory(mysql, "pms_os_disk", "os_id", os_id)

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		
		s := strings.Fields(line)
		fmt.Println(s, len(s))
		mounted :=s[0]
		total_size :=s[1]
		used_size :=s[2]
		avail_size :=s[3]
		used_rate :=s[4]

		sql := `insert into pms_os_disk(os_id, host, alias, mounted, total_size, used_size, avail_size, used_rate, created) 
				values(?,?,?,?,?,?,?,?,?)`
		_, err = mysql.Exec(sql, os_id, host, alias, mounted, total_size, used_size, avail_size, used_rate, time.Now().Unix())
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

func GatherLinuxDiskIOInfo(mysql *xorm.Engine, os_id int, host string, alias string) error{
	s_cmd := `/usr/bin/snmptable -v1 -c public ` + host + ` diskIOTable |grep -ivE "ram|loop|md|SNMP table|diskIOIndex" | grep -v '^$'`

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



