package controllers

import (
	//"opms/initial"

	"fmt"
	. "opms/models/roles"
	. "opms/models/messages"
	//"opms/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	//UserInfo string
	UserId   int64
	Username string
	UserAvatar   string
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
		tmp := strings.Split((this.GetSession("userLogin")).(string), "||")

		userid, _ := strconv.Atoi(tmp[0])
		longid := int64(userid)
		this.Data["LoginUserid"] = longid
		this.Data["LoginUsername"] = tmp[1]
		this.Data["LoginAvatar"] = tmp[2]

		this.UserId = longid
		this.Username = tmp[1]
		this.UserAvatar = tmp[2]

		//this.Data["PermissionModel"] = this.GetSession("userPermissionModel")
		//this.Data["PermissionModelc"] = this.GetSession("userPermissionModelc")

		//消息
		msgcondArr := make(map[string]string)
		msgcondArr["touserid"] = fmt.Sprintf("%d", longid)
		msgcondArr["view"] = "1"
		countTopMessage := CountMessages(msgcondArr)
		_, _, topMessages := ListMessages(msgcondArr, 1, 10)
		this.Data["topMessages"] = topMessages
		this.Data["countTopMessage"] = countTopMessage

		//fmt.Println(this.GetSession("userGroupid").(string))
		//左侧导航
		//_, _, leftNav := ListGroupsUserPermission(this.GetSession("userRoleid").(string))
		_, _, leftNav := ListRoleUserPermission(this.GetSession("userRoleid").(string))
		this.Data["leftNav"] = leftNav
		
	}
	this.Data["IsLogin"] = this.IsLogin
}

