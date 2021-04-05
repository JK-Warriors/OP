package index

import (
	//"log"
	"opms/controllers"
	
	. "opms/models/index"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//主页
type MainController struct {
	controllers.BaseController
}

func (this *MainController) Get() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "alarm-manage") {
	// 	this.Abort("401")
	// }
	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	alias := this.GetString("alias")
	host := this.GetString("host")
	condArr := make(map[string]string)
	condArr["alias"] = alias
	condArr["host"] = host

	countDb := CountDb(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDb)
	_, _, db := ListDbStatus(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["db"] = db
	this.Data["countDb"] = countDb

	this.TplName = "index.tpl"
}