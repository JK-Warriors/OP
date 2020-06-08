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
      <!-- <h3> 日志管理 </h3>-->
      <ul class="breadcrumb pull-left">
        <li> <a href="/system/log/manage">日志管理</a> </li>
        <li class="active"> 系统日志 </li>
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
                  <!--search start-->
                  <form action="/system/log/manage" method="get">
                    <input type="text" class="form-control" name="username" placeholder="请输入用户名" value="{{.condArr.username}}"/>
                    <input type="text" class="form-control" name="ip" placeholder="请输入IP" value="{{.condArr.ip}}"/>
                    <button type="submit" class="btn btn-primary"><i class="fa fa-search"></i>搜索</button>
                    <a href="/system/log/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                  </form>
                  <!--search end-->
                </div>
                <div class="pull-right">
                </div>
              </div>
            </div>
          </div>

          <section class="panel mail-box">
            <header class="panel-heading"> 日志列表 / 总数：{{.countLog}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body mail-box-info">
			        <a href="javascript:;" id="deletelogs" class="btn btn-sm btn-primary">批量删除</a>
              <section class="mail-list" style="margin-top:6px;">
                <form id="log-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th><input type="checkbox" class="checkboxbtn"></th>
                        <th>用户名</th>
                        <th>标题</th>
                        <th style="width:200px">Url</th>
                        <th>IP</th>
                        <th style="width:300px">浏览器</th>
                        <th>创建时间</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                        {{range $k,$v := .syslogs}}
                        <tr>
                            <!--<li class="list-group-item"> <span class="pull-left chk">
                              <input type="checkbox" class="checked" value="{{$v.Id}}">
                              </span> <a class="thumb pull-left" href="/system/user/show/{{$v.Userid}}"> <img src="{{getAvatarUserid $v.Userid}}" style="width:22px;"> </a> <a href="{{$v.Url}}"> <small class="pull-right text-muted">{{getDateMH $v.Created}}</small> <strong>{{getRealname $v.Userid}}</strong>&nbsp;&nbsp;{{$v.Title}}</span> </a> 
                            </li>-->
                            <td><input type="checkbox" class="checked" value="{{$v.Id}}">
                            </td>
                            <td>{{$v.Username}}</td>
                            <td>{{$v.Title}}</td>
                            <td style="width:200px">{{$v.Url}}</td>
                            <td>{{$v.Ip}}</td>
                            <!--<td style="width:300px">{{substr $v.Useragent 0 42}}</td>-->
                            <td style="width:300px">{{$v.Useragent}}</td>
                            <td>{{GetDateMHS $v.Created}}</td>

                            <td><div class="btn-group">
                                <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                                <ul class="dropdown-menu">
                                  <li><a href="javascript:;" class="deletelog" data-op="delete" data-id="{{$v.Id}}">删除</a></li>
                                </ul>
                              </div>
                              </td>
                        </tr>
                        {{else}}
                        <tr>
                          <td colspan="8" class="text-center">当前无日志记录</td>
                        </tr>
                        {{end}}
                    </tbody>
                    
                  </table>
                </form>
                {{template "inc/page.tpl" .}} </section>
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
$(function(){
	//全选　
	$('.checkboxbtn').click(function(){
		$('#log-form-list').find("input[type='checkbox']").prop('checked', $(this).is(':checked'));   							 
	});
	
	$('#deletelogs').on('click', function(){
	
		var ck = $('.checked:checked');
		if(ck.length <= 0) { dialogInfo('至少选择一个复选框'); return false; }
		
		var str = '';
		$.each(ck, function(i, n){
			str += n['value']+',';
		});
		str = str.substring(0, str.length - 1)

		layer.confirm('您确定要删除吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
      $.post('/system/log/ajax/delete', {ids:str},function(data){
        dialogInfo(data.message)
        if (data.code) {
          setTimeout(function(){ window.location.reload(); }, 1000);
        } else {
          setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
        }			
      },'json');
		});
	});

  
	$('.deletelog').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');
    
		layer.confirm('您确定要删除吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
			$.post('/system/log/ajax/delete', {ids:id},function(data){
				dialogInfo(data.message)
				if (data.code) {
					setTimeout(function(){ window.location.reload() }, 1000);
				} else {
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
				}
				
			},'json');
		});
	});

})
</script>
</body>
</html>
