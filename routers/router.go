package routers

import (
	"opms/controllers/business"
	"opms/controllers/dbconfig"
	"opms/controllers/demo"
	"opms/controllers/disaster_config"
	"opms/controllers/disaster_oper"
	"opms/controllers/logs"
	"opms/controllers/messages"
	"opms/controllers/roles"
	"opms/controllers/screen"
	"opms/controllers/users"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &users.MainController{})

	//用户
	beego.Router("/login", &users.LoginUserController{})
	beego.Router("/logout", &users.LogoutUserController{})
	beego.Router("/system/user/manage", &users.ManageUserController{})
	beego.Router("/system/user/ajax/status", &users.AjaxStatusUserController{})
	beego.Router("/system/user/add", &users.AddUserController{})
	beego.Router("/system/user/edit/:id", &users.EditUserController{})
	beego.Router("/system/user/ajax/delete", &users.AjaxDeleteUserController{})
	beego.Router("/system/user/avatar", &users.AvatarUserController{})
	beego.Router("/system/user/ajax/search", &users.AjaxSearchUserController{}) //搜索用户名匹配
	beego.Router("/system/user/show/:id", &users.ShowUserController{})
	beego.Router("/system/user/profile", &users.EditUserProfileController{})
	beego.Router("/system/user/password", &users.EditUserPasswordController{})
	beego.Router("/system/user/ajax/reset_passwd", &users.AjaxResetPasswordController{})

	beego.Router("/system/my/manage", &users.ShowUserController{})

	//消息
	beego.Router("/system/message/manage", &messages.ManageMessageController{})
	beego.Router("/system/message/ajax/delete", &messages.AjaxDeleteMessageController{})
	beego.Router("/system/message/ajax/status", &messages.AjaxStatusMessageController{})

	//角色
	beego.Router("/system/role/manage", &roles.ManageRoleController{})
	beego.Router("/system/role/ajax/delete", &roles.AjaxDeleteRoleController{})
	beego.Router("/system/role/add", &roles.FormRoleController{})
	beego.Router("/system/role/edit/:id", &roles.FormRoleController{})
	//角色成员
	beego.Router("/system/role/user/:id", &roles.ManageRoleUserController{})
	beego.Router("/system/role/user/add/:id", &roles.FormRoleUserController{})
	beego.Router("/system/role/user/ajax/delete", &roles.AjaxDeleteRoleUserController{})

	//角色权限
	beego.Router("/system/role/permission/:id", &roles.ManageRolePermissionController{})
	beego.Router("/system/role/permission/ajax/delete", &roles.AjaxDeleteRolePermissionController{})

	//权限
	beego.Router("/system/permission/manage", &roles.ManagePermissionController{})
	beego.Router("/system/permission/ajax/delete", &roles.AjaxDeletePermissionController{})
	beego.Router("/system/permission/add", &roles.FormPermissionController{})
	beego.Router("/system/permission/edit/:id", &roles.FormPermissionController{})

	//日志
	beego.Router("/system/log/manage", &logs.ManageLogController{})
	beego.Router("/system/log/ajax/delete", &logs.AjaxDeleteLogController{})

	//业务系统配置
	beego.Router("/config/business/manage", &business.ManageBusinessController{})
	beego.Router("/config/business/add", &business.AddBusinessController{})
	beego.Router("/config/business/edit", &business.EditBusinessController{})
	beego.Router("/config/business/ajax/delete", &business.AjaxDeleteBusinessController{})

	//数据库配置
	beego.Router("/config/db/manage", &dbconfig.ManageDBConfigController{})
	beego.Router("/config/db/add", &dbconfig.AddDBConfigController{})
	beego.Router("/config/db/edit/:id", &dbconfig.EditDBConfigController{})
	beego.Router("/config/db/ajax/status", &dbconfig.AjaxStatusDBConfigController{})
	beego.Router("/config/db/ajax/delete", &dbconfig.AjaxDeleteDBConfigController{})
	beego.Router("/config/db/ajax/connect", &dbconfig.AjaxConnectDBConfigController{})

	//容灾配置
	beego.Router("/config/disaster/manage", &disaster_config.ManageDisasterController{})
	beego.Router("/config/disaster/edit/:id", &disaster_config.EditDisasterController{})

	//操作
	beego.Router("/operation/disaster_switch/manage", &disaster_oper.ManageDisasterSwitchController{})
	//beego.Router("/operation/disaster_switch/view/:id", &disaster_oper.ViewDisasterSwitchController{})
	beego.Router("/operation/disaster_switch/screen/:id", &disaster_oper.ScreenDisasterSwitchController{})
	beego.Router("/operation/disaster_switch/switchover", &disaster_oper.AjaxDisasterSwitchoverController{})
	beego.Router("/operation/disaster_switch/failover", &disaster_oper.AjaxDisasterFailoverController{})
	beego.Router("/operation/disaster_switch/process", &disaster_oper.AjaxDisasterProcessController{})
	beego.Router("/operation/disaster_active/manage", &disaster_oper.ManageDisasterActiveController{})
	beego.Router("/operation/disaster_active/startread", &disaster_oper.AjaxDisasterStartReadController{})
	beego.Router("/operation/disaster_active/stopread", &disaster_oper.AjaxDisasterStopReadController{})
	beego.Router("/operation/disaster_snapshot/manage", &disaster_oper.ManageDisasterSnapshotController{})
	beego.Router("/operation/disaster_snapshot/startsnapshot", &disaster_oper.AjaxDisasterStartSnapshotController{})
	beego.Router("/operation/disaster_snapshot/stopsnapshot", &disaster_oper.AjaxDisasterStopSnapshotController{})
	beego.Router("/operation/disaster_sync/manage", &disaster_oper.ManageDisasterSyncController{})
	beego.Router("/operation/disaster_sync/startsync", &disaster_oper.AjaxDisasterStartSyncController{})
	beego.Router("/operation/disaster_sync/stopsync", &disaster_oper.AjaxDisasterStopSyncController{})
	beego.Router("/operation/disaster_recover/manage", &disaster_oper.ManageDisasterRecoverController{})
	beego.Router("/operation/disaster_recover/oper/:id", &disaster_oper.OperDisasterRecoverController{})
	beego.Router("/operation/disaster_recover/flashback", &disaster_oper.AjaxDisasterFlashbackController{})
	beego.Router("/operation/disaster_recover/recover", &disaster_oper.AjaxDisasterRecoverController{})

	//大屏
	beego.Router("/screen/manage", &screen.ManageScreenController{})

	//UI demo
	beego.Router("/demo/index", &demo.DemoController{})
	beego.Router("/demo/form", &demo.FormController{})
	beego.Router("/demo/base", &demo.BaseController{})
	beego.Router("/demo/dashboard", &demo.DashboardController{})
	beego.Router("/demo/dgscreen", &demo.DgscreenController{})
}
