package dbconfig

import (
	//"fmt"
	"opms/models"
	//"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Dbconfigs struct {
	Id           int    `orm:"pk;column(id);"`
	Dbtype       int    `orm:"column(db_type);"`
	Host         string `orm:"column(host);"`
	Port         int    `orm:"column(port);"`
	Alias        string `orm:"column(alias);"`
	InstanceName string `orm:"column(instance_name);"`
	Dbname       string `orm:"column(db_name);"`
	Username     string `orm:"column(username);"`
	Password     string `orm:"column(password);"`
	Bs_Id        int    `orm:"column(bs_id);"`
	Role         int    `orm:"column(role);"`
	Status       int    `orm:"column(status);"`
	IsDelete     int    `orm:"column(is_delete);"`
	Retention    int    `orm:"column(retention);"`
	Created      int64  `orm:"column(created);"`
	Updated      int64  `orm:"column(updated);"`
}

func (this *Dbconfigs) TableName() string {
	return models.TableName("db_config")
}
func init() {
	orm.RegisterModel(new(Dbconfigs))
}

//添加数据库
func AddDBconfig(upd Dbconfigs) error {
	o := orm.NewOrm()
	o.Using("default")
	dbconf := new(Dbconfigs)

	dbconf.Dbtype = upd.Dbtype
	dbconf.Host = upd.Host
	dbconf.Port = upd.Port
	dbconf.Alias = upd.Alias
	dbconf.InstanceName = upd.InstanceName
	dbconf.Dbname = upd.Dbname
	dbconf.Username = upd.Username
	dbconf.Password = upd.Password
	dbconf.Bs_Id = upd.Bs_Id
	dbconf.Role = upd.Role
	dbconf.Status = 1
	dbconf.IsDelete = 0
	dbconf.Created = time.Now().Unix()
	_, err := o.Insert(dbconf)
	return err
}

//修改数据库配置信息
func UpdateDBconfig(id int, upd Dbconfigs) error {
	var dbconf Dbconfigs
	o := orm.NewOrm()
	dbconf, err := GetDBconfig(id)
	if err == nil {
		dbconf.Dbtype = upd.Dbtype
		dbconf.Host = upd.Host
		dbconf.Port = upd.Port
		dbconf.Alias = upd.Alias
		dbconf.InstanceName = upd.InstanceName
		dbconf.Dbname = upd.Dbname
		dbconf.Username = upd.Username
		dbconf.Password = upd.Password
		dbconf.Bs_Id = upd.Bs_Id
		dbconf.Role = upd.Role
		dbconf.Updated = time.Now().Unix()

		_, err = o.Update(&dbconf)
	}
	return err
}

//得到数据库配置信息
func GetDBconfig(id int) (Dbconfigs, error) {
	var dbconf Dbconfigs
	var err error
	o := orm.NewOrm()

	dbconf = Dbconfigs{Id: id}
	err = o.Read(&dbconf)

	if err == orm.ErrNoRows {
		return dbconf, nil
	}
	return dbconf, err
}

//获取数据库类型
func GetDBtype(id int) string {
	var db_type string

	if id == 1 {
		db_type = "Oracle"
	} else if id == 2 {
		db_type = "MySQL"
	} else if id == 3 {
		db_type = "SQLServer"
	}
	return db_type
}

//获取数据库配置列表
func ListDBconfig(condArr map[string]string, page int, offset int) (num int64, err error, dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("db_config"))
	cond := orm.NewCondition()

	if condArr["dbtype"] != "" {
		cond = cond.And("db_type", condArr["dbtype"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}

	cond = cond.And("is_delete", 0)

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("id")
	nums, errs := qs.Limit(offset, start).All(&dbconf)
	return nums, errs, dbconf
}

func ListAllDBconfig() (dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("db_config"))
	cond := orm.NewCondition()

	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&dbconf)
	return dbconf
}

func ListPrimaryDBconfig(bs_id int) (dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("db_config"))
	cond := orm.NewCondition()

	cond = cond.And("bs_id", bs_id)
	cond = cond.And("role", 1)
	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&dbconf)
	return dbconf
}

func ListStandbyDBconfig(bs_id int) (dbconf []Dbconfigs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("db_config"))
	cond := orm.NewCondition()

	cond = cond.And("bs_id", bs_id)
	cond = cond.And("role", 2)
	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&dbconf)
	return dbconf
}

//统计数量
func CountDBconfig(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("db_config"))
	cond := orm.NewCondition()

	if condArr["dbtype"] != "" {
		cond = cond.And("db_type", condArr["dbtype"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["alias"] != "" {
		cond = cond.And("alias__icontains", condArr["alias"])
	}
	cond = cond.And("status", 1)
	cond = cond.And("is_delete", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteDBconfig(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("db_config") + " WHERE id IN(" + ids + ")").Exec()
	return err
}

//更改数据库状态
func ChangeDBconfigStatus(id int, status int) error {
	o := orm.NewOrm()

	dbconf := Dbconfigs{Id: id}
	err := o.Read(&dbconf, "id")
	if nil != err {
		return err
	} else {
		dbconf.Status = status
		_, err := o.Update(&dbconf)
		return err
	}
}
