package config_global

import (
	//"fmt"
	"opms/models"
	//"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GlobalOption struct {
	Id       			string `orm:"pk;column(Id);"`
	Name       			string `orm:"column(name);"`
	Value   			string `orm:"column(value);"`
	Description   		string `orm:"column(description);"`
}

func (this *GlobalOption) TableName() string {
	return models.TableName("global_options")
}
func init() {
	orm.RegisterModel(new(GlobalOption))
}


//获取全局配置列表
func ListGlobalOptions(condArr map[string]string, page int, offset int) (num int64, err error, options []GlobalOption) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("global_options"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("name__icontains", condArr["search_name"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	//qs = qs.OrderBy("name")
	nums, errs := qs.Limit(offset, start).All(&options)
	return nums, errs, options
}


//统计数量
func CountOptions(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("global_options"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("name__icontains", condArr["search_name"])
	}
	
	num, _ := qs.SetCond(cond).Count()
	return num
}

func SaveGlobalConfig(id string, value string) error {
	o := orm.NewOrm()

	g_option := GlobalOption{Id: id}
	err := o.Read(&g_option, "Id")
	if nil != err {
		return err
	} else {
		g_option.Value = value
		_, err := o.Update(&g_option)
		return err
	}
}
