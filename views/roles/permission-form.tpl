<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
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
        <li> <a href="/system/user/manage">系统管理</a> </li>
        <li> <a href="/system/permission/manage">权限管理</a> </li>
        <li class="active"> 权限 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="permission-form">
			
					
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">父节点</label>
                  <div class="col-sm-10">
					<select name="parentid" class="form-control">
                      <option value="">请选择父节点</option>
					{{range .permissions}}
                      <option value="{{.Id}}" {{if eq .Id $.permission.Parentid}}selected{{end}}>{{.Name}}</option>
					{{end}}
                    </select>                  
					</div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>权限名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="name" value="{{.permission.Name}}" class="form-control" placeholder="请输入权限名称，如编辑用户">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>菜单地址</label>
                  <div class="col-sm-10">
                    <input type="text" name="ename" value="{{.permission.Url}}" class="form-control" placeholder="请输入菜单地址，如 /system/user/edit">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">是否显示</label>
                  <div class="col-sm-10">
                    <label class="radio-inline">
                    <input type="radio" name="type" value="1" {{if eq 1 .permission.IsShow}}checked{{end}}>
                    是 </label>    
                    <label class="radio-inline">
                    <input type="radio" name="type" value="0" {{if eq 0 .permission.IsShow}}checked{{end}}>
                    否 </label>                
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">样式</label>
                  <div class="col-sm-10">
                    <input type="text" name="icon" value="{{.permission.Icon}}" class="form-control" placeholder="ICON, 如 user、tasks、suitcase、plane、laptop、home、book">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">排序</label>
                  <div class="col-sm-10">
                    <input type="text" name="sort" value="{{.permission.Sort}}" class="form-control" placeholder="数字">
                  </div>
                </div>                
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" id="permissionid" value="{{.permission.Id}}">
                    <button type="submit" class="btn btn-primary">提交保存</button>
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
</body>
</html>
