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
        <li> <a href="/user/show/{{.LoginUserid}}">用户权限</a> </li>
        <li> <a href="/user/manage">用户管理</a> </li>
        <li class="active"> 用户 </li>
      </ul>
      <div class="pull-right"><a href="/user/add" class="btn btn-success">+添加用户</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="userprofile-form">
                <header><b> 基本信息 </b></header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>用户名</label>
                  <div class="col-sm-10">
                    <input type="text" name="username"  value="{{.user.Username}}" class="form-control" placeholder="请填写用户名">
                  </div>
                </div>
                <div id="div_passwd" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>密码</label>
                  <div class="col-sm-10">
                    <input type="password" name="password"  value="" class="form-control" placeholder="请填写密码">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>角色</label>
                  <div class="col-sm-10">
                    <select name="role" class="form-control">
                      <option value="">请选择角色</option>
				              {{range .roles}}
                      <option value="{{.Id}}" {{if eq .Id $.userrole.Roleid}}selected{{end}}>{{.Name}}</option>
				              {{end}}
                    </select>
                  </div>
                </div>
                <header> <b>帐号信息</b> </header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>姓名</label>
                  <div class="col-sm-10">
                    <input type="text" name="realname"  value="{{.pro.Realname}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">性别</label>
                  <div class="col-sm-10">
                    <label class="radio-inline">
                    <input type="radio" name="sex" value="1" {{if eq 1 .pro.Sex}}checked{{end}}>
                    男 </label>
                    <label class="radio-inline">
                    <input type="radio" name="sex" value="2" {{if eq 2 .pro.Sex}}checked{{end}}>
                    女 </label>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>生日</label>
                  <div class="col-sm-10">
                    <input type="text" name="birth" id="default-date-picker"  value="{{.pro.Birth}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>邮箱</label>
                  <div class="col-sm-10">
                    <input type="email" name="email"  value="{{.pro.Email}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>手机号</label>
                  <div class="col-sm-10">
                    <input type="number" name="phone" maxlength="11" value="{{.pro.Phone}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">微信号</label>
                  <div class="col-sm-10">
                    <input type="text" name="webchat"  value="{{.pro.Webchat}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">QQ</label>
                  <div class="col-sm-10">
                    <input type="number" name="qq"  value="{{.pro.Qq}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">电话</label>
                  <div class="col-sm-10">
                    <input type="text" name="tel"  value="{{.pro.Tel}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">住址</label>
                  <div class="col-sm-10">
                    <input type="text" name="address"  value="{{.pro.Address}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>紧急联系人</label>
                  <div class="col-sm-10">
                    <input type="text" name="emercontact"  value="{{.pro.Emercontact}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>紧急联系人电话</label>
                  <div class="col-sm-10">
                    <input type="text" name="emerphone"  value="{{.pro.Emerphone}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="{{.pro.Id}}">
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
<script src="/static/js/datepicker-zh-CN.js"></script>
<script>
    $(function(){
      $('#default-date-picker').datepicker('option', $.datepicker.regional['zh-CN']); 	
      $('#default-date-picker').datepicker({
            dateFormat: 'yy-mm-dd',
        changeMonth: true,
        changeYear: true,
        yearRange:'-60:+0'
        });
    });

    $(function(){ 
      var pw = "{{.user.Password}}";
      if(pw!=""){
        $("#div_passwd").hide();
      }
    });
</script>
</body>
</html>
