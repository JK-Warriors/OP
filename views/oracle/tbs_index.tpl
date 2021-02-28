<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8" />
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section>
  {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content">
    <!-- header section start-->
    <div class="header-section">
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--search start-->
      <!--search end-->
      {{template "inc/user-info.tpl" .}}
    </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <!-- <h3> 日志管理 </h3>-->
      <ul class="breadcrumb pull-left">
        <li><a href="/oracle/status/manage">Oracle</a></li>
        <li class="active">表空间列表</li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <!-- 主体内容 开始 -->
            <div class="searchdiv">
              <div class="search-form">
                <div class="form-inline">
                  <div class="form-group">
                    <form action="/oracle/tbs/manage" method="get">
                    <input type="text" name="alias" placeholder="请输入别名" class="form-control" value="{{.condArr.alias}}"/>
                    <input type="text" name="host" placeholder="请输入主机" class="form-control" value="{{.condArr.host}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/oracle/tbs/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 表空间列表 / 总数：{{.CountTbs}}
                <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
                <!--a href="javascript:;" class="fa fa-times"></a-->
                </span> 
              </header>
              <div class="panel-body">
                <section id="unseen">
                  <form id="user-form-list">
                    <table class="table table-bordered table-striped table-condensed">
                      <thead>
                        <tr>
                          <th>别名</th>
                          <th>主机</th>
                          <th>表空间名</th>
                          <th>状态</th>
                          <th>管理方式</th>
                          <th>总大小(MB)</th>
                          <th>已使用大小(MB)</th>
                          <th>使用率</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .tbs}}
                        <tr>
                          <td>{{$v.Alias}}</td>
                          <td>{{$v.Host}}</td>
                          <td>{{$v.Tablespace_Name}}</td>
                          <td>{{$v.Status}}</td>
                          <td>{{$v.Management}}</td>
                          <td>{{$v.Total_Size}}</td>
                          <td>{{$v.Used_Size}}</td>
                          <td>{{$v.Max_Rate}}</td>
                        </tr>
                      {{end}}
                      </tbody>
                    </table>
                  </form>
                  {{template "inc/page.tpl" .}}
                </section>
              </div>
            </section>
          <!-- 主体内容 结束 -->
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
<script>
   
</script>
</body>
</html>
