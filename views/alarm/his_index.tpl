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
        <li><a href="/alarm/history/list">告警管理</a></li>
        <li class="active">告警列表</li>
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
                    <form action="/alarm/history/list" method="get">
                    <input type="text" name="search_name" placeholder="请输入名称" class="form-control" value="{{.condArr.search_name}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/alarm/history/list" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 告警列表 / 总数：{{.CountAlerts}}
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
                          <th>级别</th>
                          <th>告警对象</th>
                          <th>名称</th>
                          <th>描述</th>
                          <th>是否发送邮件</th>
                          <th>发送邮件状态</th>
                          <th>是否发送微信</th>
                          <th>发送微信状态</th>
                          <th>是否发送短信</th>
                          <th>发送短信状态</th>
                          <th>时间</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .alerts}}
                        <tr>
                          <td>{{$v.Severity}}</td>
                          <td>{{getDBDesc $v.Asset_Id}}</td>
                          <td>{{$v.Name}}</td>
                          <td>{{$v.Message}}</td>
                          <td>{{$v.Send_Mail}}</td>
                          <td>{{$v.Send_Mail_Status}}</td>
                          <td>{{$v.Send_WeChat}}</td>
                          <td>{{$v.Send_WeChat_Status}}</td>
                          <td>{{$v.Send_SMS}}</td>
                          <td>{{$v.Send_SMS_Status}}</td>
                          <td>{{GetDateMHS $v.Created}}</td>
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
