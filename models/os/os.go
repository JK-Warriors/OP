package mssql

import (
	//"fmt"
	"opms/models"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OS struct {
	Id       					int `orm:"pk;column(id);"`
	Os_Id   					int `orm:"column(os_id);"`
	Host   						string `orm:"column(host);"`
	Alias   					string `orm:"column(alias);"`
	Connect   					int `orm:"column(connect);"`
	Hostname   					string `orm:"column(hostname);"`
	Kernel   					string `orm:"column(kernel);"`
	System_Date   				string `orm:"column(system_date);"`
	System_Uptime   			string `orm:"column(system_uptime);"`
	Process   					string `orm:"column(process);"`
	Load_1   					string `orm:"column(load_1);"`
	Load_5   					string `orm:"column(load_5);"`
	Load_15   					string `orm:"column(load_15);"`
	Cpu_User_Time 				string `orm:"column(cpu_user_time);"`
	Cpu_System_Time 			string `orm:"column(cpu_system_time);"`
	Cpu_Idle_Time 				string `orm:"column(cpu_idle_time);"`
	Swap_Total 					string `orm:"column(swap_total);"`
	Swap_Avail 					string `orm:"column(swap_avail);"`
	Mem_Total 					string `orm:"column(mem_total);"`
	Mem_Avail 					string `orm:"column(mem_avail);"`
	Mem_Free 					string `orm:"column(mem_free);"`
	Mem_Shared 					string `orm:"column(mem_shared);"`
	Mem_Buffered 				string `orm:"column(mem_buffered);"`
	Mem_Cached 					string `orm:"column(mem_cached);"`
	Mem_Usage_Rate 				string `orm:"column(mem_usage_rate);"`
	Mem_Available 				string `orm:"column(mem_available);"`
	Disk_IO_Reads_Total 		string `orm:"column(disk_io_reads_total);"`
	Disk_IO_Writes_Total 		string `orm:"column(disk_io_writes_total);"`
	Net_In_Bytes_Total 			string `orm:"column(net_in_bytes_total);"`
	Net_Out_Bytes_Total 		string `orm:"column(net_out_bytes_total);"`
	Created  					int64  `orm:"column(created);"`
}

type OSDisk struct {
	Id       					int `orm:"pk;column(id);"`
	Os_Id   					int `orm:"column(os_id);"`
	Host   						string `orm:"column(host);"`
	Alias   					string `orm:"column(alias);"`
	Mounted   					string `orm:"column(mounted);"`
	Total_Size   				string `orm:"column(total_size);"`
	Used_Size   				string `orm:"column(used_size);"`
	Avail_Size   				string `orm:"column(avail_size);"`
	Used_Rate   				string `orm:"column(used_rate);"`
	Created  					int64  `orm:"column(created);"`
}

type OSDiskIO struct {
	Id       					int `orm:"pk;column(id);"`
	Os_Id   					int `orm:"column(os_id);"`
	Host   						string `orm:"column(host);"`
	Alias   					string `orm:"column(alias);"`
	Fdisk   					string `orm:"column(fdisk);"`
	Disk_IO_Reads   			string `orm:"column(disk_io_reads);"`
	Disk_IO_Writes   			string `orm:"column(disk_io_writes);"`
	Created  					int64  `orm:"column(created);"`
}

func (this *OS) TableName() string {
	return models.TableName("os_status")
}

func (this *OSDisk) TableName() string {
	return models.TableName("os_disk")
}

func (this *OSDiskIO) TableName() string {
	return models.TableName("os_diskio")
}


func init() {
	orm.RegisterModel(new(OS))
	orm.RegisterModel(new(OSDisk))
	orm.RegisterModel(new(OSDiskIO))
}


//获取OS状态列表
func ListOSStatus(condArr map[string]string, page int, offset int) (num int64, err error, os []OS) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("os_status"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("os_id")
	nums, errs := qs.Limit(offset, start).All(&os)
	return nums, errs, os
}


//统计数量
func CountOS(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("os_status"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}



//获取OS Disk状态列表
func ListDiskStatus(condArr map[string]string, page int, offset int) (num int64, err error, disk []OSDisk) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("os_disk"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("os_id")
	nums, errs := qs.Limit(offset, start).All(&disk)
	return nums, errs, disk
}

//统计数量
func CountDisk(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("os_disk"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}


//获取OS Disk IO状态列表
func ListDiskIOStatus(condArr map[string]string, page int, offset int) (num int64, err error, diskio []OSDiskIO) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("os_diskio"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("os_id")
	nums, errs := qs.Limit(offset, start).All(&diskio)
	return nums, errs, diskio
}

//统计数量
func CountDiskIO(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("os_diskio"))
	cond := orm.NewCondition()

	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}