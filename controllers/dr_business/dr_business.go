package dr_business

import (
	"fmt"
	"opms/controllers"
	. "opms/models/dr_business"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//业务系统管理
type ManageBusinessController struct {
	controllers.BaseController
}

func (this *ManageBusinessController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-business-manage") {
		this.Abort("401")
	}

	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	search_name := this.GetString("search_name")
	condArr := make(map[string]string)
	condArr["search_name"] = search_name

	countBs := CountBusiness(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countBs)
	_, _, bsconf := ListBusiness(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["bsconf"] = bsconf
	this.Data["countBs"] = countBs

	this.TplName = "dr_business/index.tpl"
}

//添加业务系统
type AddBusinessController struct {
	controllers.BaseController
}

func (this *AddBusinessController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-business-add") {
		this.Abort("401")
	}

	var bsconf Business
	this.Data["bsconf"] = bsconf

	this.TplName = "dr_business/form.tpl"
}

func (this *AddBusinessController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-business-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权新增"}
		this.ServeJSON()
		return
	}

	bs_name := this.GetString("bs_name")
	if "" == bs_name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写业务系统名称"}
		this.ServeJSON()
		return
	}

	var bsconf Business

	bsconf.BsName = bs_name
	is_exist := CheckNameExist(bs_name)
	if 1 == is_exist {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "业务系统名称已经存在，请换一个名称"}
	} else {
		err := AddBusiness(bsconf)

		if err == nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "业务系统添加成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "业务系统添加失败"}
		}
	}

	this.ServeJSON()
}

//修改业务系统
type EditBusinessController struct {
	controllers.BaseController
}

func (this *EditBusinessController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-business-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	bsconf, err := GetBusiness(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["bsconf"] = bsconf
}

func (this *EditBusinessController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-business-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}
	idstr := this.GetString("bs_id")
	//utils.LogDebug(idstr)
	id, _ := strconv.Atoi(idstr)

	bs_name := this.GetString("bs_name")
	if "" == bs_name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写业务系统名称"}
		this.ServeJSON()
		return
	}

	var bsconf Business

	bsconf.BsName = bs_name

	err := UpdateBusiness(id, bsconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "业务系统信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "业务系统信息修改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteBusinessController struct {
	controllers.BaseController
}

func (this *AjaxDeleteBusinessController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-business-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权删除"}
		this.ServeJSON()
		return
	}
	ids := this.GetString("ids")
	if "" == ids {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的记录"}
		this.ServeJSON()
		return
	}

	err := DeleteBusiness(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
