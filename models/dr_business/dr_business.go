package dr_business

import (
	//"fmt"
	"opms/models"
	//"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Business struct {
	Id       int    `orm:"pk;column(id);"`
	BsName   string `orm:"column(bs_name);"`
	IsDelete int    `orm:"column(is_delete);"`
	Created  int64  `orm:"column(created);"`
	Updated  int64  `orm:"column(updated);"`
}

func (this *Business) TableName() string {
	return models.TableName("dr_business")
}
func init() {
	orm.RegisterModel(new(Business))
}

//添加业务系统
func AddBusiness(bs Business) error {
	o := orm.NewOrm()
	o.Using("default")
	bsconf := new(Business)

	bsconf.BsName = bs.BsName
	bsconf.IsDelete = 0
	bsconf.Created = time.Now().Unix()
	_, err := o.Insert(bsconf)
	return err
}

//验证系统名称是否已经存在
func CheckNameExist(name string) int {
	o := orm.NewOrm()
	var li_count int32
	_ = o.Raw("select count(1) from pms_dr_business WHERE bs_name =?", name).QueryRow(&li_count)

	if li_count == 0 {
		return 0
	}
	return 1
}

//修改业务系统
func UpdateBusiness(id int, bs Business) error {
	var bsconf Business
	o := orm.NewOrm()
	bsconf = Business{Id: id}

	bsconf.BsName = bs.BsName
	bsconf.Updated = time.Now().Unix()

	_, err := o.Update(&bsconf)
	return err
}

//得到业务系统
func GetBusiness(id int) (Business, error) {
	var bsconf Business
	var err error
	o := orm.NewOrm()

	bsconf = Business{Id: id}
	err = o.Read(&bsconf)

	if err == orm.ErrNoRows {
		return bsconf, nil
	}
	return bsconf, err
}

func GetBusinessName(id int) string {
	var bs_name string
	o := orm.NewOrm()

	err := o.Raw("select bs_name from pms_dr_business where id = ?", id).QueryRow(&bs_name)

	if err == orm.ErrNoRows {
		return ""
	}
	return bs_name
}

//获取业务系统列表
func ListBusiness(condArr map[string]string, page int, offset int) (num int64, err error, bsconf []Business) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("dr_business"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("bs_name__icontains", condArr["search_name"])
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
	nums, errs := qs.Limit(offset, start).All(&bsconf)
	return nums, errs, bsconf
}

func ListAllBusiness() (bsconf []Business) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("dr_business"))
	cond := orm.NewCondition()

	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&bsconf)
	return bsconf
}

//统计数量
func CountBusiness(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("dr_business"))
	cond := orm.NewCondition()

	if condArr["bs_name"] != "" {
		cond = cond.And("bs_name__icontains", condArr["bs_name"])
	}
	cond = cond.And("is_delete", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteBusiness(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("dr_business") + " WHERE id IN(" + ids + ")").Exec()
	return err
}
