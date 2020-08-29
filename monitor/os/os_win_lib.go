package os

import (
	"fmt"
	"bytes"
	"strings"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/masterzen/winrm"

)

func GetWinHostname(client *winrm.Client) (string, error){
	var value string

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic os get CSName /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "CSName="){
				value := strings.Replace(s[0], "CSName=", "",1)
				return value, err
			}
		}
	}

	return value, nil
}

func GetWinKernel(client *winrm.Client) (string, error){
	var value string

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic os get caption /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Split(stdout.String(),"=")
		fmt.Println(s, len(s))

		if len(s) == 2{
			value := strings.Trim(strings.Replace(strings.Replace(s[1], "\n", "", -1), "\r", "", -1), " ")
			fmt.Println(len(value))

			return value, err
		}
	}

	return value, nil
}


func GetWinSystemDate(client *winrm.Client) (string, error){
	var value string

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic os get LocalDateTime /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "LocalDateTime="){
				value := strings.Replace(s[0], "LocalDateTime=", "",1)
				if(strings.Contains(value, ".")){
					ss := strings.Split(value, ".")
					return ss[0], err
				}else{
					return value, err
				}
			}
		}
	}

	return value, nil
}


func GetWinUptime(client *winrm.Client) (string, error){
	var value string

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic os get LastBootUpTime /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "LastBootUpTime="){
				value := strings.Replace(s[0], "LastBootUpTime=", "",1)
				if(strings.Contains(value, ".")){
					ss := strings.Split(value, ".")
					return ss[0], err
				}else{
					return value, err
				}
			}
		}
	}

	return value, nil
}


func GetWinProcess(client *winrm.Client) (int, error){
	var value int =-1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic process get CommandLine /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		value := strings.Count(stdout.String(),"CommandLine=")
		return value, err
	}

	return value, nil
}


func GetWinIdleCPU(client *winrm.Client) (string, error){
	var value string

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic path Win32_PerfFormattedData_PerfOS_Processor where Name="_Total" get PercentIdleTime  /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "PercentIdleTime="){
				value := strings.Replace(s[0], "PercentIdleTime=", "",1)
				return value, err
			}
		}
	}

	return value, nil
}


func GetWinMemoryFree(client *winrm.Client) (int64, error){
	var value int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic os get FreePhysicalMemory /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "FreePhysicalMemory="){
				size := strings.Replace(s[0], "FreePhysicalMemory=", "",1)

				value, err := strconv.ParseInt(size, 10, 64)
				return value, err
			}
		}
	}

	return value, nil
}


func GetWinMemoryTotal(client *winrm.Client) (int64, error){
	var value int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic os get TotalVisibleMemorySize /format:list`)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return value, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "TotalVisibleMemorySize="){
				size := strings.Replace(s[0], "TotalVisibleMemorySize=", "",1)

				value, err := strconv.ParseInt(size, 10, 64)
				return value, err
			}
		}
	}

	return value, nil
}


func GetDriverSize(client *winrm.Client, driver string) (int64, error){
	var driversize int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic LogicalDisk where DeviceID="%s" get Size  /format:list`, driver)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return driversize, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "Size="){
				size := strings.Replace(s[0], "Size=", "",1)
				//fmt.Printf("%s ", size)

				driversize, err := strconv.ParseInt(size, 10, 64)
				return driversize, err
			}
		}
	}

	return driversize, nil
}

func GetDriverFreeSize(client *winrm.Client, driver string) (int64, error){
	var freesize int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic LogicalDisk where DeviceID="%s" get FreeSpace  /format:list`, driver)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return freesize, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "FreeSpace="){
				size := strings.Replace(s[0], "FreeSpace=", "",1)
				//fmt.Printf("%s ", size)

				freesize, err := strconv.ParseInt(size, 10, 64)
				return freesize, err
			}
		}
	}

	return freesize, nil
}


func GetDiskReadBytesPerSec(client *winrm.Client, driver string) (int64, error){
	var readbytes int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic path Win32_PerfFormattedData_PerfDisk_LogicalDisk where Name="%s" get DiskReadBytesPersec /format:list`, driver)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return readbytes, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "DiskReadBytesPersec="){
				size := strings.Replace(s[0], "DiskReadBytesPersec=", "",1)
				//fmt.Printf("%s ", size)

				readbytes, err := strconv.ParseInt(size, 10, 64)
				return readbytes, err
			}
		}
	}

	return readbytes, nil
}


func GetDiskWriteBytesPerSec(client *winrm.Client, driver string) (int64, error){
	var writebytes int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic path Win32_PerfFormattedData_PerfDisk_LogicalDisk where Name="%s" get DiskWriteBytesPersec /format:list`, driver)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return writebytes, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "DiskWriteBytesPersec="){
				size := strings.Replace(s[0], "DiskWriteBytesPersec=", "",1)
				//fmt.Printf("%s ", size)

				writebytes, err := strconv.ParseInt(size, 10, 64)
				return writebytes, err
			}
		}
	}

	return writebytes, nil
}


func GetBytesReceivedPerSec(client *winrm.Client, net string) (int64, error){
	var in_bytes int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic path Win32_PerfFormattedData_Tcpip_NetworkInterface where Name="%s" get BytesReceivedPersec /format:list`, net)
	//fmt.Println(cmd)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return in_bytes, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "BytesReceivedPersec="){
				size := strings.Replace(s[0], "BytesReceivedPersec=", "",1)
				//fmt.Printf("%s ", size)

				in_bytes, err := strconv.ParseInt(size, 10, 64)
				return in_bytes, err
			}
		}
	}

	return in_bytes, nil
}

func GetBytesSentPerSec(client *winrm.Client, net string) (int64, error){
	var out_bytes int64 = -1

	var stdout, stderr bytes.Buffer
	cmd := fmt.Sprintf(`wmic path Win32_PerfFormattedData_Tcpip_NetworkInterface where Name="%s" get BytesSentPersec /format:list`, net)
	_, err := client.Run(cmd, &stdout, &stderr)
	if nil != err {
		return out_bytes, err
	}else{
		//fmt.Println(stdout.String())
		s := strings.Fields(stdout.String())
		//fmt.Println(s, len(s))

		if len(s) == 1{
			if strings.Contains(s[0], "BytesSentPersec="){
				size := strings.Replace(s[0], "BytesSentPersec=", "",1)
				//fmt.Printf("%s ", size)

				out_bytes, err := strconv.ParseInt(size, 10, 64)
				return out_bytes, err
			}
		}
	}

	return out_bytes, nil
}