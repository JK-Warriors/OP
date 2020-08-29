package os

import (
	"log"
	//"reflect"
	"os/exec"
	"bufio"
	"io"
	"strings"
	"fmt"
	"strconv"
	
	"opms/monitor/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
  gs "github.com/soniah/gosnmp"
)

func GetSystemDate(snmp *gs.GoSNMP) (string, error){
	oids := []string{".1.3.6.1.2.1.25.1.2.0"}		//HOST-RESOURCES-MIB::hrSystemDate.0
	
	var sysdate string

	result, err := snmp.Get(oids)
    if err != nil {
        utils.LogDebugf("GetSystemDate err: %s", err.Error())
		return "-1", err
	}
	
	fmt.Println(result)

    for _, v := range result.Variables {
		//fmt.Println(reflect.TypeOf(v.Value)) 
		if(nil == v.Value){
			return "-1", nil
		}else{
			sysdate = fmt.Sprintf("%s", v.Value)
		}
	}  

	return sysdate, nil
}



func GetUptime(snmp *gs.GoSNMP) (string, error){
	oids := []string{".1.3.6.1.2.1.25.1.1.0"}		//HOST-RESOURCES-MIB::hrSystemUptime.0	
	
	var uptime string

	result, err := snmp.Get(oids)
    if err != nil {
        utils.LogDebugf("GetUptime err: %s", err.Error())
		return "-1", err
	}
	
    for _, v := range result.Variables {
		uptime = fmt.Sprintf("%d", gs.ToBigInt(v.Value))
		//fmt.Println(reflect.TypeOf(uptime)) 
	}  

	return uptime, nil
}


func GetNetDescr(host string) (map[string]string, error){
	var net = make(map[string]string)
	
	s_cmd := `/usr/bin/snmpwalk -v1 -c public ` + host + ` IF-MIB::ifDescr | grep -ivE "lo|sit0" `

	cmd := exec.Command("/bin/bash", "-c", s_cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("GetNetDescrerror: %s", err.Error())
		return nil, err
	}
	cmd.Start()

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}

		s := strings.Fields(line)
		
		
		var idx string
		if(s[0] != ""){
			kv := strings.Split(s[0], ".")
			if len(kv) == 2 {
				idx = kv[1]
			}
		}
		
		name := s[3]
		
		net[idx] = name
		//fmt.Println(idx, name)
		
	}

	return net, nil
}

func GetNetBytesIn(host string, netid string) (int64, error){
	s_cmd := `/usr/bin/snmpwalk -v1 -c public ` + host + ` IF-MIB::ifInOctets.` + netid + `| awk '{print $NF}'  `
	//fmt.Println(s_cmd)

	cmd := exec.Command("/bin/bash", "-c", s_cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("GetNetInBytes: %s", err.Error())
		return -1, err
	}
	cmd.Start()

	var netbytes int64= 0
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		
		line = strings.Replace(line, "\n", "", -1)
		
		netbytes, err = strconv.ParseInt(line,10,64)
		if(nil != err){
			log.Printf("GetNetInBytes: %s", err.Error())
		}
	}

	return netbytes, nil
}

func GetNetBytesOut(host string, netid string) (int64, error){
	s_cmd := `/usr/bin/snmpwalk -v1 -c public ` + host + ` IF-MIB::ifOutOctets.` + netid + `| awk '{print $NF}'  `

	cmd := exec.Command("/bin/bash", "-c", s_cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("GetNetBytesOut: %s", err.Error())
		return -1, err
	}
	cmd.Start()

	var netbytes int64= 0
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		
		line = strings.Replace(line, "\n", "", -1)
		
		netbytes, err = strconv.ParseInt(line,10,64)
		if(nil != err){
			log.Printf("GetNetBytesOut: %s", err.Error())
		}
	}

	return netbytes, nil
}

func GetSnmpValueByOids(snmp *gs.GoSNMP, oids []string) (map[int]string, error){
	valuestr := make(map[int]string)

	result, err := snmp.Get(oids)
    if err != nil {
        utils.LogDebugf("GetSnmpValueByOids err: %s", err.Error())
		return nil, err
	}
	
    for i, v := range result.Variables {
        //fmt.Printf("%d. oid: %s ", i, v.Name)
        switch v.Type{
        case gs.OctetString:
			//fmt.Printf("string: %s\n", v.Value)
			valuestr[i] = fmt.Sprintf("%s", v.Value)
		case gs.Uinteger32:
			//fmt.Printf("number: %d\n", gs.ToBigInt(v.Value))
			valuestr[i] = fmt.Sprintf("%d", gs.ToBigInt(v.Value))
        default:
            //fmt.Printf("number: %d\n", gs.ToBigInt(v.Value))
			valuestr[i] = fmt.Sprintf("%s", v.Value)
		}
		
	}  
	
	return valuestr, nil
}


func GetDiskIOReadTotal(mysql *xorm.Engine, os_id int) (int64){
	var total int64= -1
	sql := `select sum(disk_io_reads) from pms_os_diskio where os_id = ?`
	_, err := mysql.SQL(sql, os_id).Get(&total)
	if err != nil {
		log.Printf("GetDiskIOReadTotal failed: %s", err.Error())
		return -1
	}

	return total
}

func GetDiskIOWriteTotal(mysql *xorm.Engine, os_id int) (int64){
	var total int64= -1
	sql := `select sum(disk_io_writes) from pms_os_diskio where os_id = ?`
	_, err := mysql.SQL(sql, os_id).Get(&total)
	if err != nil {
		log.Printf("GetDiskIOWriteTotal failed: %s", err.Error())
		return -1
	}

	return total
}

func GetNetInBytesTotal(mysql *xorm.Engine, os_id int) (int64){
	var total int64= -1
	sql := `select sum(in_bytes) from pms_os_net where os_id = ?`
	_, err := mysql.SQL(sql, os_id).Get(&total)
	if err != nil {
		log.Printf("GetNetInBytesTotal failed: %s", err.Error())
		return -1
	}

	return total
}

func GetNetOutBytesTotal(mysql *xorm.Engine, os_id int) (int64){
	var total int64= -1
	sql := `select sum(out_bytes) from pms_os_net where os_id = ?`
	_, err := mysql.SQL(sql, os_id).Get(&total)
	if err != nil {
		log.Printf("GetNetInBytesTotal failed: %s", err.Error())
		return -1
	}

	return total
}


func GetDiskIOReadsOld(mysql *xorm.Engine, os_id int, fdisk string) (int64){
	var oldvalue int64 = -1
	sql := `select disk_io_reads from pms_os_diskio_tmp where os_id = ? and fdisk = ?`
	_, err := mysql.SQL(sql, os_id, fdisk).Get(&oldvalue)
	if err != nil {
		log.Printf("GetDiskIOReadsOld failed: %s", err.Error())
		return -1
	}

	return oldvalue
}

func GetDiskIOWritesOld(mysql *xorm.Engine, os_id int, fdisk string) (int64){
	var oldvalue int64 = -1
	sql := `select disk_io_writes from pms_os_diskio_tmp where os_id = ? and fdisk = ?`
	_, err := mysql.SQL(sql, os_id, fdisk).Get(&oldvalue)
	if err != nil {
		log.Printf("GetDiskIOWritesOld failed: %s", err.Error())
		return -1
	}

	return oldvalue
}

func GetDiskIOTime(mysql *xorm.Engine, os_id int, fdisk string) (int64){
	var oldvalue int64 = -1
	sql := `select created from pms_os_diskio_tmp where os_id = ? and fdisk = ?`
	_, err := mysql.SQL(sql, os_id, fdisk).Get(&oldvalue)
	if err != nil {
		log.Printf("GetDiskIOTime failed: %s", err.Error())
		return -1
	}

	return oldvalue
}

func GetNetInBytesOld(mysql *xorm.Engine, os_id int, fdisk string) (int64){
	var oldvalue int64 = -1
	sql := `select in_bytes from pms_os_net_tmp where os_id = ? and if_descr = ?`
	_, err := mysql.SQL(sql, os_id, fdisk).Get(&oldvalue)
	if err != nil {
		log.Printf("GetNetInBytesOld failed: %s", err.Error())
		return -1
	}

	return oldvalue
}

func GetNetOutBytesOld(mysql *xorm.Engine, os_id int, fdisk string) (int64){
	var oldvalue int64 = -1
	sql := `select out_bytes from pms_os_net_tmp where os_id = ? and if_descr = ?`
	_, err := mysql.SQL(sql, os_id, fdisk).Get(&oldvalue)
	if err != nil {
		log.Printf("GetNetOutBytesOld failed: %s", err.Error())
		return -1
	}

	return oldvalue
}

func GetNetTime(mysql *xorm.Engine, os_id int, fdisk string) (int64){
	var oldvalue int64 = -1
	sql := `select created from pms_os_net_tmp where os_id = ? and if_descr = ?`
	_, err := mysql.SQL(sql, os_id, fdisk).Get(&oldvalue)
	if err != nil {
		log.Printf("GetNetTime failed: %s", err.Error())
		return -1
	}

	return oldvalue
}



func MoveToHistory(mysql *xorm.Engine, table_name string, key_name string, key_value int){
	sql := `insert into ` + table_name + `_his select * from ` + table_name + ` where ` + key_name + ` = ?`
	_, err := mysql.Exec(sql, key_value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}

	sql = `delete from ` + table_name + ` where ` + key_name + ` = ?`
	_, err = mysql.Exec(sql, key_value)
	if err != nil {
		log.Printf("%s: %s", sql, err.Error())
	}
}
