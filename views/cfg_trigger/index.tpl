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
        <li><a href="/config/cfg_trigger/manage">配置中心</a></li>
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
                    <form action="/config/cfg_trigger/manage" method="get">
                    <input type="text" name="search_name" placeholder="请输入告警类型" class="form-control" value="{{.condArr.search_name}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/config/cfg_trigger/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 告警列表 / 总数：{{.countTriggers}}
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
                          <th>资产主机</th>
                          <th>资产别名</th>
                          <th>资产类型</th>
                          <th>告警类型</th>
                          <th>级别</th>
                          <th>表达式</th>
                          <th>描述</th>
                          <th>状态</th>
                          <th>操作</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .triggers}}
                        <tr>
                          <td>{{$v.Asset_Host}}:{{$v.Asset_Port}}</td>
                          <td>{{$v.Asset_Alias}}</td>
                          <td>{{getDBtype $v.Asset_Type}}</td>
                          <td>{{$v.Trigger_Type}}</td>
                          <td>{{$v.Severity}}</td>
                          <td>{{$v.Expression}}</td>
                          <td>{{$v.Description}}</td>
                          <td>{{if eq 1 $v.Status}}激活{{else}}禁用{{end}}</td>
                          <td><div class="btn-group">
                              <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                              <ul class="dropdown-menu">
                                <li><a href="/config/cfg_trigger/edit/{{$v.Id}}">编辑</a></li>
                                <li role="separator" class="divider"></li>
                                {{if eq 1 $v.Status}}
                                <li><a href="javascript:;" class="js-dbconfig-single" data-id="{{$v.Id}}" data-status="2">禁用</a></li>
                                {{else}}
                                <li><a href="javascript:;" class="js-dbconfig-single" data-id="{{$v.Id}}" data-status="1">激活</a></li>
                                {{end}}
                              </ul>
                            </div></td>
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
    //layer
    function delete_bs(e){
      var id = e.getAttribute("data-id");
      
      layer.confirm('您确定要删除吗？', {
        btn: ['确定','取消'] //按钮
        ,title:"提示"
      }, function(index){
        layer.close(index);
        
        $.post('/config/dr_business/ajax/delete', {ids:id},function(data){
          dialogInfo(data.message)
          if (data.code) {
            setTimeout(function(){ window.location.reload() }, 1000);
          } else {
            setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
          }
        },'json');
      });
    }

    
    $('#business-form').validate({
        ignore:'',        
		    rules : {
			      bs_name:{required: true},
        },
        messages : {
			      bs_name:{required: '请填写业务系统名'},
        },

        submitHandler:function(form) {
            var id = $("#bs_id").val()
            if(id == ""){
                target_url = "/config/dr_business/add";
            }else{
                target_url = "/config/dr_business/edit";
            }
            $(form).ajaxSubmit({
                type:'POST',
                url: target_url, 
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       setTimeout(function(){window.location.href="/config/dr_business/manage"}, 1000);
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
