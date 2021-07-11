<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<style>
.form-signin .help-block{color:#a94442;}
</style>
</head><body class="login-body">
<div class="container">
  <form class="form-signin" id="login-form">
    <div class="form-signin-heading text-center">
      <h1 class="sign-title">登录磐石</h1>
      <!--<img src="/static/img/logo.png" alt="磐石数据库容灾系统" style="width:120px;"/>-->
      <font face="微软雅黑" size="6" color="blue" style="width:500px;">磐石容灾管理系统</font> 
      </div>
    <div class="login-wrap">
      <input type="text" class="form-control" name="username" placeholder="请填写用户名" autofocus>
      <input type="password" class="form-control" name="password" placeholder="请填写密码">
      
	<button class="btn btn-lg btn-login btn-block" type="submit"> 登录</button>
    </div>
  </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>
