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
      <!--search start-->
      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!--body wrapper start-->
    <!-- page heading start-->
    <!-- <div class="page-heading">
      <ul class="breadcrumb pull-left">
        <li class="active"><a href="/">首页</a></li>
      </ul>
    </div>-->
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
                    <form action="/" method="get">
                    <input type="text" name="alias" placeholder="请输入别名" class="form-control" value="{{.condArr.alias}}"/>
                    <input type="text" name="host" placeholder="请输入主机" class="form-control" value="{{.condArr.host}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 资产状态列表 / 总数：{{.countDb}}
                <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
                <!--a href="javascript:;" class="fa fa-times"></a-->
                </span> 
              </header>
              <div class="panel-body">
                <section id="unseen">
                  <form id="user-form-list">
                    <table class="table table-bordered table-striped table-condensed">
                      <thead>
                        <tr style="font-size: 13px;">
                        <th colspan="5"><center>数据库</center></th>
                        <th colspan="8"><center>状态</center></th>
                        <th colspan="5"><center>操作系统</center></th>
                        <!--<th></th>-->
                        </tr>
                        <tr>
                          <th>类型</th>
                          <th>别名</th>
                          <th>主机</th>
                          <th>角色</th>
                          <th>版本</th>
                          <th>连接</th>
                          <th>会话</th>
                          <th>活动</th>
                          <th>等待</th>
                          <th>进程数</th>
                          <th>同步</th>
                          <th>延时</th>
                          <th>表空间</th>
                          <th>负载</th>
                          <th>CPU</th>
                          <th>内存</th>
                          <th>网络</th>
                          <th>磁盘</th>
                          <!--<th>图表</th>-->
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .db}}
                        <tr>
                          <td>{{ getAssetImage $v.Asset_Type | str2html}}</td>
                          <td>{{$v.Alias}}</td>
                          <td>{{$v.Host}}</td>
                          <td>{{ getDbRoleImage $v.Role | str2html}}</td>
                          <td>{{$v.Version}}</td>
                          <td>{{checkDbStatusLevel $v.Connect $v.Connect_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Sess_Total $v.Sess_Total_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Sess_Actives $v.Sess_Actives_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Sess_Waits $v.Sess_Waits_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Process $v.Process_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Repl $v.Repl_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Repl_Delay $v.Repl_Delay_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Tablespace $v.Tablespace_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Load $v.Load_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Cpu $v.Cpu_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Memory $v.Memory_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.IO $v.IO_Tips | str2html}}</td>
                          <td>{{checkDbStatusLevel $v.Net $v.Net_Tips | str2html}}</td>
                          <!--<td>-</td>-->
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
</body>
</html>
