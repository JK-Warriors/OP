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
        <li><a href="/config/db/manage">配置中心</a></li>
        <li class="active">全局配置</li>
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
                    <form action="/config/config_global/manage" method="get">
                    <input type="text" name="search_name" placeholder="请输入名称" class="form-control" value="{{.condArr.search_name}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/config/config_global/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 全局配置列表 / 总数：{{.countOptions}}
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
                          <th>名称</th>
                          <th>值</th>
                          <th>描述</th>
                          <th>操作</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .options}}
                        <tr>
                          <td>{{$v.Name}}</td>
                          <td contentEditable="true">{{$v.Value}}</td>
                          <td>{{$v.Description}}</td>
                          <td>
                            <a href="javascript:;" class="table_btn table_btn_icon" onclick="edit_config(this)" data-id="{{$v.Id}}" data-name="{{$v.Name}}" data-value="{{$v.Value}}">
                              <i class="iconfont icon-btn_edit"></i>编辑
                            </a>
                          </td>
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
      
  <form id="config-form">
    <div id="config_box" class="layui_drm">
      <div class="layercontent">
        <!-- layer content start -->
        <div class="form-horizontal adminex-form">
          <div class="form-group">
            <label id="name" class="col-xs-2  control-label"></label>
            <div class="col-xs-10">
              <input type="hidden" id="id" name="id" value="" class="form-control"/>
              <input type="text" id="value" name="value" value="" class="form-control"/>
            </div>
          </div>
        </div>
        <!-- layer content end -->
      </div>
      <!-- layer foot start -->
      <div class="layerfoot">
        <button type="submit" class="btn btn-primary">提 交</button>
      </div>
      <!-- layer content end -->
    </div>
  </form>
      
</section>

{{template "inc/foot.tpl" .}}    
<script>

    function edit_config(e){
      var id = e.getAttribute("data-id");
      var name = e.getAttribute("data-name");
      var value = e.getAttribute("data-value");

      $("#name").html(name);
      $("#id").attr("value",id);
      $("#value").attr("value",value);

      layer.open({
        type: 1,
        closeBtn: true,
        shift: 2,
        title: '编辑配置',
        area: ['600px', '30%'],
        offset: ['180px'],
        shadeClose: true,
        content: $('#config_box')
      })
    }

    $('#config-form').validate({
        ignore:'',     

        submitHandler:function(form) {
            var id = $("#id").val()
            var value = $("#value").val()
            
            $(form).ajaxSubmit({
                type:'POST',
                url: '/config/config_global/ajax/save', 
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       setTimeout(function(){window.location.href="/config/config_global/manage"}, 1000);
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
