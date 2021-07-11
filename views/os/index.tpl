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
        <li><a href="/os/status/manage">OS</a></li>
        <li class="active">状态列表</li>
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
                    <form action="/os/status/manage" method="get">
                    <input type="text" name="alias" placeholder="请输入别名" class="form-control" value="{{.condArr.alias}}"/>
                    <input type="text" name="host" placeholder="请输入主机" class="form-control" value="{{.condArr.host}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/os/status/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> OS列表 / 总数：{{.countOS}}
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
                          <th>连接状态</th>
                          <th>进程数</th>
                          <th>CPU空闲率</th>
                          <th>总内存</th>
                          <th>可用内存</th>
                          <th>内存使用率 (%)</th>
                          <th>读IO</th>
                          <th>写IO</th>
                          <th>网络进流量</th>
                          <th>网络出流量</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .osList}}
                        <tr>
                          <td>{{$v.Alias}}</td>
                          <td>{{$v.Host}}</td>
                          <td>{{checkDbStatusLevel $v.Connect "" | str2html}}</td>
                          <td>{{$v.Process}}</td>
                          <td>{{$v.Cpu_Idle_Time}}</td>
                          <td>{{$v.Mem_Total}}</td>
                          <td>{{$v.Mem_Available}}</td>
                          <td>{{$v.Mem_Usage_Rate}}</td>
                          <td>{{$v.Disk_IO_Reads_Total}}</td>
                          <td>{{$v.Disk_IO_Writes_Total}}</td>
                          <td>{{$v.Net_In_Bytes_Total}}</td>
                          <td>{{$v.Net_Out_Bytes_Total}}</td>
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
