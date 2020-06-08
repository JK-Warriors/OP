package screen

import (
	"opms/controllers"
)

// 聚合大屏
type ManageScreenController struct {
	controllers.BaseController
}

func (this *ManageScreenController) Get() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-sync") {
	// 	this.Abort("401")
	// }

	this.TplName = "screen/index.tpl"
}
