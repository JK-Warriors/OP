package initial

import (
	//"fmt"
	"opms/models/dr_business"
	"opms/models/dbconfig"
	"opms/models/roles"
	"opms/models/users"
	"opms/utils"

	//"time"

	"github.com/astaxie/beego"
)

func InitTemplate() {
	beego.AddFuncMap("getRealname", users.GetRealname)
	beego.AddFuncMap("getAvatarUserid", users.GetAvatarUserid)
	beego.AddFuncMap("getPermissionname", roles.GetPermissiontName)

	beego.AddFuncMap("getDBtype", dbconfig.GetDBtype)
	beego.AddFuncMap("getDBDesc", dbconfig.GetDBDesc)
	beego.AddFuncMap("getDBAlias", dbconfig.GetDBAlias)
	beego.AddFuncMap("getBsName", dr_business.GetBusinessName)
	/*
		beego.AddFuncMap("getNeedsname", projects.GetProjectNeedsName)
		beego.AddFuncMap("getProjectname", projects.GetProjectName)
		beego.AddFuncMap("getLeaveProcess", leaves.ListLeaveApproverProcessHtml)
		beego.AddFuncMap("getExpenseProcess", expenses.ListExpenseApproverProcessHtml)
		beego.AddFuncMap("getBusinesstripProcess", businesstrips.ListBusinesstripApproverProcessHtml)
		beego.AddFuncMap("getGooutProcess", goouts.ListGooutApproverProcessHtml)
		beego.AddFuncMap("getOagoodProcess", oagoods.ListOagoodApproverProcessHtml)
		beego.AddFuncMap("getOvertimeProcess", overtimes.ListOvertimeApproverProcessHtml)
	*/

	beego.AddFuncMap("getDate", utils.GetDate)
	beego.AddFuncMap("getDateMH", utils.GetDateMH)
	beego.AddFuncMap("GetDateMHS", utils.GetDateMHS)
	beego.AddFuncMap("GetDateDiff", utils.GetDateDiff)
	beego.AddFuncMap("GetDateDiffColor", utils.GetDateDiffColor)
	beego.AddFuncMap("getNeedsStatus", utils.GetNeedsStatus)
	beego.AddFuncMap("getNeedsSource", utils.GetNeedsSource)
	beego.AddFuncMap("getNeedsStage", utils.GetNeedsStage)
	beego.AddFuncMap("getTaskStatus", utils.GetTaskStatus)
	beego.AddFuncMap("getTaskType", utils.GetTaskType)
	beego.AddFuncMap("getTestStatus", utils.GetTestStatus)
	beego.AddFuncMap("getLeaveType", utils.GetLeaveType)

	beego.AddFuncMap("getOs", utils.GetOs)
	beego.AddFuncMap("getBrowser", utils.GetBrowser)
	beego.AddFuncMap("getAvatarSource", utils.GetAvatarSource)
	beego.AddFuncMap("getAvatar", utils.GetAvatar)

	beego.AddFuncMap("getEdu", utils.GetEdu)
	beego.AddFuncMap("getWorkYear", utils.GetWorkYear)
	beego.AddFuncMap("getResumeStatus", utils.GetResumeStatus)

	beego.AddFuncMap("getCheckworkType", utils.GetCheckworkType)

	beego.AddFuncMap("getMessageType", utils.GetMessageType)
	beego.AddFuncMap("getMessageSubtype", utils.GetMessageSubtype)

}
