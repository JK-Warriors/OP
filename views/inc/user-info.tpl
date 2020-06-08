 <!--toggle button start-->
 
 <div class="header-menu">
   <li {{if or (eq "/" .url_first_part) (eq "" .url_first_part)}}class="active"{{end}}><a href="/">首页</a></li>
   <li {{if eq "/config" .url_first_part}}class="active"{{end}}><a href="/config/business/manage">配置</a></li>
   <li {{if eq "/business" .url_first_part}}class="active"{{end}}><a href="/business/manage">业务系统</a></li>
   <li {{if eq "/operation" .url_first_part}}class="active"{{end}}><a href="/operation/disaster_switch/manage">操作</a></li>
   <li {{if eq "/screen" .url_first_part}}class="active"{{end}}><a href="/screen/manage">大屏</a></li>
   <li {{if eq "/system" .url_first_part}}class="active"{{end}}><a href="/system/user/manage">系统管理</a></li>
 </div>
 <!--toggle button end-->

<!--notification menu start -->
<div class="menu-right">
  <ul class="notification-menu">
  <li>
                    <a href="#" class="btn btn-default dropdown-toggle info-number" data-toggle="dropdown">
                        <i class="fa fa-envelope-o"></i>
						{{if gt .countTopMessage 0}}
                        <span class="badge">{{.countTopMessage}}</span>
						{{end}}
                    </a>
                    <div class="dropdown-menu dropdown-menu-head pull-right">
                        <h5 class="title">你有 {{.countTopMessage}} 最新信息</h5>
                        <ul class="dropdown-list normal-list">
						{{range $k,$v := .topMessages}}
                            <li>
                                <a href="{{$v.Url}}" class="js-msg-status" data-id="{{$v.Id}}">
                                    <span class="thumb"><img src="{{getAvatarUserid $v.Userid}}"></span>
                                        <span class="desc">
                                          <span class="name">{{getRealname $v.Userid}}<!--span class="badge badge-success">new</span--></span>
                                          <span class="msg">{{getMessageType $v.Type}}了{{getMessageSubtype $v.Subtype}}&nbsp;&nbsp;{{$v.Title}}</span>
                                        </span>
                                </a>
                            </li>
							 {{else}}
			                  <li class="text-center">目前还没有最新消息</li>
			                  {{end}}
                            <li class="new"><a href="/system/message/manage">查看更多</a></li>
                        </ul>
                    </div>
                </li>
  
    <li> <a href="javascript:;" class="btn btn-default dropdown-toggle" data-toggle="dropdown"> <img src="{{getAvatar .LoginAvatar}}" alt="{{.LoginUsername}}" /> {{.LoginUsername}} <span class="caret"></span> </a>
      <ul class="dropdown-menu dropdown-menu-usermenu pull-right">
        <!--<li><a href="/system/user/show/{{.LoginUserid}}"><i class="fa fa-user"></i> 个人主页</a></li>-->
        <li><a href="/system/user/profile"><i class="fa fa-cog"></i> 基本资料</a></li>
		    <li><a href="/system/user/avatar"><i class="fa fa-camera"></i> 更换头像</a></li>
		    <li><a href="/system/user/password"><i class="fa fa-cog"></i> 更换密码</a></li>
		    <!--<li><a href="/my/task"><i class="fa fa-th-list"></i> 我的任务</a></li>-->		
        <li><a href="/logout"><i class="fa fa-sign-out"></i> 退出</a></li>
      </ul>
    </li>
  </ul>
</div>
<!--notification menu end -->
