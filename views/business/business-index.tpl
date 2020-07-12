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
        <li><a href="/config/business/manage">配置中心</a></li>
        <li class="active">业务系统配置</li>
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
                    <form action="/config/business/manage" method="get">
                    <input type="text" name="search_name" placeholder="请输入名称" class="form-control" value="{{.condArr.search_name}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/config/business/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
              <div class="pull-right">
                <a href="javascript:;" class="btn btn-success" id="add_business">
                  <i class="fa fa-plus"></i> 新增业务</a>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 业务系统列表 / 总数：{{.countBs}}
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
                          <th>业务系统名称</th>
                          <th>操作</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .bsconf}}
                        <tr>
                          <td>{{$v.BsName}}</td>
                          <td>
                            <a href="javascript:;" class="table_btn table_btn_icon" onclick="edit_bs(this)" data-id="{{$v.Id}}" data-name="{{$v.BsName}}">
                              <i class="iconfont icon-btn_edit"></i>编辑
                            </a>
                            <a href="/config/disaster/edit/{{$v.Id}}" class="table_btn">
                              <i class="iconfont icon-xianghujiaohuan"></i>容灾配置
                            </a>
                            <a href="javascript:;" class="table_btn table_btn_icon" onclick="delete_bs(this)" data-id="{{$v.Id}}">
                              <i class="iconfont icon-iconfontshanchu"></i>
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
      
  <form id="business-form">
    <div id="business_box" class="layui_drm">
      <div class="layercontent">
        <!-- layer content start -->
        <div class="form-horizontal adminex-form">
          <div class="form-group">
            <label class="col-xs-2  control-label">业务系统名称</label>
            <div class="col-xs-10">
              <input type="hidden" id="bs_id" name="bs_id" value="" class="form-control"/>
              <input type="text" id="bs_name" name="bs_name" value="" class="form-control" placeholder="请填写业务系统名称"/>
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
    //layer
    $(function() {
      $('#add_business').click(function() {
        $("#bs_id").attr("value",'');
        $("#bs_name").attr("value",'');

        layer.open({
          type: 1,
          closeBtn: true,
          shift: 2,
          title: '新增业务系统',
          area: ['600px', '30%'],
          offset: ['180px'],
          shadeClose: true,
          content: $('#business_box')
        })
      })
    })

    
    function edit_bs(e){
      var id = e.getAttribute("data-id");
      var bs_name = e.getAttribute("data-name");

      $("#bs_id").attr("value",id);
      $("#bs_name").attr("value",bs_name);

      layer.open({
        type: 1,
        closeBtn: true,
        shift: 2,
        title: '编辑业务系统',
        area: ['600px', '30%'],
        offset: ['180px'],
        shadeClose: true,
        content: $('#business_box')
      })
    }


    function delete_bs(e){
      var id = e.getAttribute("data-id");
      
      layer.confirm('您确定要删除吗？', {
        btn: ['确定','取消'] //按钮
        ,title:"提示"
      }, function(index){
        layer.close(index);
        
        $.post('/config/business/ajax/delete', {ids:id},function(data){
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
                target_url = "/config/business/add";
            }else{
                target_url = "/config/business/edit";
            }
            $(form).ajaxSubmit({
                type:'POST',
                url: target_url, 
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       setTimeout(function(){window.location.href="/config/business/manage"}, 1000);
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
