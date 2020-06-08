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
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <!-- <h3> 组织管理 {{template "users/nav.tpl" .}}</h3>-->
      <ul class="breadcrumb pull-left">
        <li> <a href="/system/role/manage">角色管理</a> </li>
        <li class="active"> 角色列表 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
            <div class="search-form">
              <div class="form-inline">
                <div class="form-group">
                  <!--search start-->
                  <form action="/system/role/manage" method="get">
                    <input type="text" class="form-control" name="keywords" placeholder="请输入角色名" value="{{.condArr.keywords}}"/>
                    <button type="submit" class="btn btn-primary"><i class="fa fa-search"></i>搜索</button>
                    <a href="/system/role/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                  </form>
                  <!--search end-->
                </div>
                <div class="pull-right">
                  <a href="/system/role/add" class="btn btn-success" id="add_role"><i class="fa fa-plus"></i> 新增角色</a>
                </div>
              </div>
            </div>


          <section class="panel">
            <header class="panel-heading"> 角色管理 / 总数：{{.countRole}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a> </span> </header>
            <div class="panel-body">
              <table class="table table-bordered table-striped table-condensed">
                <thead>
                  <tr>
                    <th>名称</th>
                    <th class="hidden-phone hidden-xs">描述</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                
                {{range $k,$v := .roles}}
                <tr>
                  <td> {{$v.Name}} </td>
                  <td class="hidden-phone hidden-xs">{{$v.Summary}}</td>
                  <td><div class="btn-group">
                      <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                      <ul class="dropdown-menu">
                        <li><a href="/system/role/permission/{{$v.Id}}">权限</a></li>
                        <!--<li role="separator" class="divider"></li>-->
                        <!--<li><a href="/group/user/{{$v.Id}}">成员</a></li>-->
                        <li role="separator" class="divider"></li>
                        <li><a href="/system/role/edit/{{$v.Id}}">编辑</a></li>
                        <li role="separator" class="divider"></li>
                        <li><a href="javascript:;" class="js-group-delete" data-op="delete" data-id="{{$v.Id}}">删除</a></li>
                      </ul>
                    </div></td>
                </tr>
                {{else}}
                <tr>
                  <td colspan="7">你还没有添加角色</td>
                </tr>
                {{end}}
                </tbody>
                
              </table>
              {{template "inc/page.tpl" .}} </div>
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