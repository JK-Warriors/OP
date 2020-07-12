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
        <li> <a href="/config/business/manage">容灾配置</a> </li>
        <li class="active"> 容灾数据库配置 </li>
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
                  <form action="/config/disaster/manage" method="get">
                    <input type="text" class="form-control" name="host" placeholder="请输入IP" value="{{.condArr.host}}"/>
                    <input type="text" class="form-control" name="alias" placeholder="请输入别名" value="{{.condArr.alias}}"/>
                    <button type="submit" class="btn btn-primary"><i class="fa fa-search"></i>搜索</button>
                    <a href="/config/disaster/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                  </form>
                </div>
              </div>
            </div>
          </div>
          <section class="panel">
            <header class="panel-heading"> 业务系统列表 / 总数：{{.countDr}}
              <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span>
            </header>
            
            <div class="panel-body">
              <section id="unseen">
                <form id="user-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr class="text-center">
                        <th>业务系统名</th>
                        <th>主库</th>
                        <th>灾备</th>
                        <th>闪回保留天数</th>
                        <th>漂移IP</th>
                        <th>主库网卡</th>
                        <th>备库网卡</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    {{range $k,$v := .drconf}}
                    <tr>
                      <td>{{getBsName $v.Bs_Id}}</td>
                      <td>{{$v.Db_Id_P}}</td>
                      <td>{{$v.Db_Id_S}}</td>
                      <td>{{$v.Fb_Retention}}</td>
                      <td>{{$v.Shift_Vips}}</td>
                      <td>{{$v.Network_P}}</td>
                      <td>{{$v.Network_S}}</td>
                      <td><a href="/config/disaster/edit/{{$v.Bs_Id}}" class="table_btn">
                            <i class="iconfont icon-xianghujiaohuan"></i>修改
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
    $('.js-dbconfig-single').on('click', function(){
    	var that = $(this);
    	var status = that.attr('data-status')
    	var id = that.attr('data-id');
      $.post('/config/db/ajax/status', { status: status, id: id },function(data){
        dialogInfo(data.message)
        if (data.code) {
          that.attr('data-status', status == 2 ? 1 : 2).text(status == 2 ? '激活' : '禁用').parents('td').prev('td').text(status == 2 ? '禁用' : '激活');
        } else {
          
        }
        setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
      },'json');
    }); 
	
	$('.js-dbconfig-delete').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');

		layer.confirm('您确定要删除吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
			$.post('/config/db/ajax/delete', {ids:id},function(data){
				dialogInfo(data.message)
				if (data.code) {
					setTimeout(function(){ window.location.reload() }, 1000);
				} else {
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
				}
			},'json');
		});
		
	});

</script>
</body>
</html>
