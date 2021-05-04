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
        <li class="active">大屏配置</li>
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
                    <form action="/config/screen/manage" method="get">
                    <input type="text" name="alias" placeholder="请输入别名" class="form-control" value="{{.condArr.alias}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/config/screen/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 数据库列表 / 总数：{{.countDBs}}
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
                          <th>大屏显示</th>
                          <th>数据库</th>
                          <th>别名</th>
                          <th>类型</th>
                          <th>状态</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .dbs}}
                        <tr>
                          <td><input type="checkbox" class="checked" value="{{$v.Id}}" {{if eq 1 $v.Show_On_Screen}}checked{{end}}></td>
                          <td>{{$v.Host}}:{{$v.Port}}</td>
                          <td>{{$v.Alias}}</td>
                          <td>{{getDBtype $v.Dbtype}}</td>
                          <td>{{if eq 1 $v.Status}}激活{{else}}禁用{{end}}</td>
                        </tr>
                      {{end}}
                      </tbody>
                    </table>

                  
                  <div class="clearfix"></div>
                  <div class="form-group">
                    <label class="control-label col-md-3 col-sm-3 col-xs-12">大屏显示核心数据库：</label>
                    <div class="col-md-9 col-sm-9 col-xs-12">
                      <select id="core_db" name="core_db" class="form-control">
                      {{range $k,$v := .ora}}
                        <option value={{$v.Id}} {{if eq $.core_db $v.Id}}selected{{end}}>{{$v.Host}}:{{$v.Port}}({{$v.Alias}})</option>
                      {{end}}
                      </select>
                    </div>
                  </div>
                  </form>
			            <a href="javascript:;" id="save" class="btn btn-sm btn-primary">保存</a>
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
	$('#save').on('click', function(){
	
		var ck = $('.checked:checked');
		if(ck.length > 4) { dialogInfo('不能超过4个数据库'); return false; }
		
		var str = '';
		$.each(ck, function(i, n){
			str += n['value']+',';
		});
		str = str.substring(0, str.length - 1)

    var core_db = $('#core_db').val();

    $.post('/config/screen/ajax/save', {ids:str, core_db:core_db},function(data){
      dialogInfo(data.message)
      if (data.code) {
        setTimeout(function(){ window.location.reload(); }, 1000);
      } else {
        setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
      }			
    },'json');

	});
</script>
</body>
</html>
