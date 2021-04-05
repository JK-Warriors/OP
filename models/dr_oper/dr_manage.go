package dr_oper

import (
	"opms/utils"
	. "opms/models/dbconfig"
)


func GetTransferStatus(sta_id int) string{
	var status_img string
	
	dbconf,_ :=GetDBconfig(sta_id)
	if(dbconf.Dbtype == 1){
		sta_dr, err := GetStandbyDrInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetStandbyDrInfo failed: " + err.Error())
		}
	
		if(sta_dr.Mrp_Status == "ACTIVE" && sta_dr.Delay_Mins < 60 ){
			status_img = "/static/img/health_transfer.gif"
		}else if(sta_dr.Mrp_Status == "ACTIVE" && sta_dr.Delay_Mins > 60 ){
			status_img = "/static/img/trans_alarm.png"
		}else{
			status_img = "/static/img/trans_error.png"
		}
	}else if(dbconf.Dbtype == 2){
		sta_dr, err := GetDrMySqlStandbyInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetDrMySqlStandbyInfo failed: " + err.Error())
		}
		
		if(sta_dr.Slave_IO_Run == "Yes" && sta_dr.Slave_SQL_Run == "Yes" && sta_dr.Delay < 60 ){
			status_img = "/static/img/health_transfer.gif"
		}else if(sta_dr.Slave_IO_Run == "Yes" && sta_dr.Slave_SQL_Run == "Yes" && sta_dr.Delay > 60 ){
			status_img = "/static/img/trans_alarm.png"
		}else{
			status_img = "/static/img/trans_error.png"
		}
	

	}else if(dbconf.Dbtype == 3){
		sta_dr, err := GetDrMSSqlStandbyInfo(sta_id)
		if err != nil {
			utils.LogDebug("GetDrMSSqlStandbyInfo failed: " + err.Error())
		}
		
		if(sta_dr.State == 4 ){
			status_img = "/static/img/health_transfer.gif"
		}else{
			status_img = "/static/img/trans_error.png"
		}

	}
	
	return status_img
}