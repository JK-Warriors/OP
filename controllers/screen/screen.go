package screen

import (
	"opms/controllers"
	asset "opms/models/asset"
	alert "opms/models/alarm"
	
	. "opms/models/screen"
)

// 聚合大屏
type ManageScreenController struct {
	controllers.BaseController
}

func (this *ManageScreenController) Get() {
	//权限检测
	//if !strings.Contains(this.GetSession("userPermission").(string), "screen-manage") {
		//this.Abort("401")
	//}
	lastchecktime := GetLastCheckTime()
	this.Data["lastchecktime"] = lastchecktime
	
	countAsset, _, assets := asset.ListAllDBStatus()
	this.Data["assets"] = assets
	this.Data["countAsset"] = countAsset

	countAlert, _, alerts := alert.ListAllAlerts()
	this.Data["alerts"] = alerts
	this.Data["countAlert"] = countAlert
	
	_, _, alertgroup := alert.ListAlertGroup()
	this.Data["alertgroup"] = alertgroup

	_, _, dbscore := GetDBScore()
	this.Data["dbscore"] = dbscore

	dr_ora_normal := GetDRNormal(1)
	dr_ora_warning := GetDRWarning(1)
	dr_ora_critical := GetDRCritical(1)
	this.Data["dr_ora_normal"] = dr_ora_normal
	this.Data["dr_ora_warning"] = dr_ora_warning
	this.Data["dr_ora_critical"] = dr_ora_critical

	dr_mysql_normal := GetDRNormal(2)
	dr_mysql_warning := GetDRWarning(2)
	dr_mysql_critical := GetDRCritical(2)
	this.Data["dr_mysql_normal"] = dr_mysql_normal
	this.Data["dr_mysql_warning"] = dr_mysql_warning
	this.Data["dr_mysql_critical"] = dr_mysql_critical
	
	dr_mssql_normal := GetDRNormal(3)
	dr_mssql_warning := GetDRWarning(3)
	dr_mssql_critical := GetDRCritical(3)
	this.Data["dr_mssql_normal"] = dr_mssql_normal
	this.Data["dr_mssql_warning"] = dr_mssql_warning
	this.Data["dr_mssql_critical"] = dr_mssql_critical
	

	this.TplName = "screen/index.tpl"
}
