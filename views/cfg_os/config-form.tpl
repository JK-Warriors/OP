<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <!-- <h3> 组织管理 {{template "users/nav.tpl" .}}</h3>-->
      <ul class="breadcrumb pull-left">
        <li> <a href="/config/db/manage">配置中心</a> </li>
        <li> <a href="/config/os/manage">操作系统配置</a> </li>
        <li class="active"> {{if gt .osconf.Id 0}}编辑{{else}}新增{{end}}操作系统 </li>
      </ul>
      <div class="pull-right"><a href="/config/os/add" class="btn btn-success">+添加</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="osconfig-form">
                <header><b> 基本信息 </b></header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主机IP</label>
                  <div class="col-sm-10">
                    <input type="text" id="host" name="host"  value="{{.osconf.Host}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>别名</label>
                  <div class="col-sm-10">
                    <input type="text" id="alias" name="alias"  value="{{.osconf.Alias}}" class="form-control">
                  </div>
                </div>
                <div id="div_os_type" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主机类型</label>
                  <div class="col-sm-10">
                    <select id="os_type" name="os_type" class="form-control">
                      <option value="">请选择主机类型</option>
                      <option value="1" {{if eq 1 .osconf.Ostype}}selected{{end}}>Linux</option>
                      <option value="2" {{if eq 2 .osconf.Ostype}}selected{{end}}>Windows</option>
                      <option value="3" {{if eq 3 .osconf.Ostype}}selected{{end}}>AIX</option>
                      <option value="4" {{if eq 4 .osconf.Ostype}}selected{{end}}>HP-Unix</option>
                      <option value="5" {{if eq 5 .osconf.Ostype}}selected{{end}}>Solaris</option>
                    </select>
                  </div>
                </div>
                <div id="div_os_protocol" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主机协议</label>
                  <div class="col-sm-10">
                    <select id="os_protocol" name="os_protocol" class="form-control">
                      <!--<option value="ssh" {{if eq "ssh" .osconf.OsProtocol}}selected{{end}}>ssh</option>-->
                      <!--<option value="telnet" {{if eq "telnet" .osconf.OsProtocol}}selected{{end}}>telnet</option>-->
                      <option value="snmp" {{if eq "snmp" .osconf.OsProtocol}}selected{{end}}>snmp</option>
                      <option value="winrm" {{if eq "winrm" .osconf.OsProtocol}}selected{{end}}>winrm</option>
                    </select>
                  </div>
                </div>
                <div id="div_os_port" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主机端口</label>
                  <div class="col-sm-10">
                    <input type="text" id="os_port" name="os_port"  value="{{.osconf.OsPort}}" class="form-control" placeholder="请填写主机端口">
                  </div>
                </div>
                <div id="div_os_username" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>主机用户名</label>
                  <div class="col-sm-10">
                    <input type="text" id="os_username" name="os_username"  value="{{.osconf.OsUsername}}" class="form-control" placeholder="请填写主机用户名">
                  </div>
                </div>
                <div id="div_os_password" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>主机密码</label>
                  <div class="col-sm-10">
                    <input type="password" id="os_password" name="os_password"  value="{{.osconf.OsPassword}}" class="form-control" placeholder="请填写主机密码">
                  </div>
                </div>
                
                <div id="div_alarm" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>告警发送 </label>
                  <div>
                    <input type="checkbox" value="1" name="alert_mail" {{if eq 1 .osconf.Alert_Mail}}checked="checked"{{end}}>邮件
                    <input type="checkbox" value="1" name="alert_wechat" {{if eq 1 .osconf.Alert_WeChat}}checked="checked"{{end}}>微信
                    <input type="checkbox" value="1" name="alert_sms" {{if eq 1 .osconf.Alert_SMS}}checked="checked"{{end}}>短信
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" id="id" name="id" value="{{.osconf.Id}}">
                    <button type="button" onclick="checkConnect()" class="btn btn-primary">主机连接测试</button>
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script src="/static/js/jquery-ui-1.10.3.min.js"></script>
<script>
    $(function() {// 初始化内容
        id =  {{.osconf.Id}};
        if (id && id > 0){
        }else{
          $("#os_protocol").attr("value","snmp");
          $("#os_port").attr("value","161");
        }
    });  

    $("#os_protocol").change(function(){
        if($("#os_protocol").val() == "snmp"){
            $("#os_port").attr("value","161");
        }else if($("#os_protocol").val() == "winrm"){
            $("#os_port").attr("value","5985");
        }
    });
    
    function checkConnect(){
        var host = $("#host").val();
        var protocol = $("#os_protocol").val();
        var port = $("#os_port").val();
        var username = $("#os_username").val();
        var password = $("#os_password").val();
        
        $.ajax({url: "/config/os/ajax/connect",
                type: "POST",
								data: {"host":host,"protocol":protocol, "port":port,"username":username,"password":password,},
                success: function (data) {
                    dialogInfo(data.message)
                    setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 3000);
                    if (data.code) {
                    }
                    else {
                    }
                }
        });
    }

    $('#osconfig-form').validate({
        ignore:'',        
		    rules : {
			      host:{required: true},
			      alias:{required: true},
			      os_type:{required: true},
			      os_protocol:{required: true},
			      os_port:{required: true},
        },
        messages : {
			      host:{required: '请填写主机IP'}, 
			      alias:{required: '请填写主机别名'}, 
			      os_type:{required: '请选择主机类型'}, 
			      os_protocol:{required: '请选择连接协议'}, 
			      os_port:{required: '请填写协议端口'}, 
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       		setTimeout(function(){window.location.href="/config/os/manage"}, 1000);
                    } else {
                        setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });
</script>
</body>
</html>
