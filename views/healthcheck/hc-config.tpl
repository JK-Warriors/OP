<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section> 
  {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->
      <!--search end-->
      {{template "inc/user-info.tpl" .}} 
    </div>
    <!-- header section end-->
    <!-- page heading start-->
    
    <div class="page-heading">
      <!--<h3> 组织管理 {{template "users/nav.tpl" .}}</h3>-->
      <ul class="breadcrumb pull-left">
        <li> <a href="/config/dr_business/manage">健康巡检</a> </li>
        <li class="active"> 巡检配置 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <div class="searchdiv">
            <div class="search-form">
              <div class="form-inline">
                <div class="form-group">
                  <form action="/healthcheck/config/manage" method="get">
                    <input type="text" class="form-control" name="host" placeholder="请输入IP" value="{{.condArr.host}}"/>
                    <input type="text" class="form-control" name="alias" placeholder="请输入别名" value="{{.condArr.alias}}"/>
                    <button type="submit" class="btn btn-primary"><i class="fa fa-search"></i>搜索</button>
                    <a href="/healthcheck/config/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                  </form>
                </div>
              </div>
            </div>
          </div>
            <section class="panel">
              <header class="panel-heading"> 数据库列表 / 总数：{{.countDb}}
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
                          <th>类型</th>
                          <th>别名</th>
                          <th>主机</th>
                          <th></th>
                          <!--<th>图表</th>-->
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .db}}
                        <tr>
                          <td>{{ getAssetImage $v.Dbtype | str2html}}</td>
                          <td>{{$v.Alias}}</td>
                          <td>{{$v.Host}}</td>
                          <td><button type="button" id="btnConfig" class="btn btn-primary"> 配置</button></td>
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

    $("#btnConfig").click(function () {
        layer.open({
            type: 2,
            title: '健康巡检',
            shadeClose: false,
            maxmin: false,
            area: ['800px', '500px'],
            content: ["/healthcheck/config/edit",'no'],
            end: function () {
                //关闭时做的事情
                //oDataTable.ajax.reload();
            }
        });
    });


</script>
</body>
</html>
