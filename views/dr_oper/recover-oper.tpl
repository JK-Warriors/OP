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
        <li><a href="/operation/dr_recover/manage">容灾操作</a></li>
        <li><a href="/operation/dr_recover/manage">误删除恢复</a></li>
        <li class="active">恢复详情</li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <!-- 主体内容 开始 -->
            <section class="panel">
              <header class="panel-heading">
                <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
                <!--a href="javascript:;" class="fa fa-times"></a-->
                </span> 
              </header>
              <div class="panel-body">
                    <form id="recover_oper_form" class="form-horizontal adminex-form" action="" method="get">
					<div class="form-group">
						<label class="col-sm-2 col-sm-2 control-label"><span>*</span>恢复方式</label>
						<div class="col-sm-10">
							<select id="fb_method" onchange="fb_method_change(this)" class="form-control">
							<option value="">请选择恢复方式</option>
								  <option value="1"> 按快照点</option>
								  <option value="2"> 按快照时间</option>
							</select>
						</div>
                	</div>
					
					<div id="div_point" class="form-group">
						<label class="col-sm-2 col-sm-2 control-label"><span></span>快照点</label>
						<div class="col-sm-10">
							<select id="fb_point" class="form-control">
							<option value="">请选择快照点</option>
									{{range $k,$v :=.restore_point }}
										<option value="{{$v}}">{{$v}}</option>
									{{end}}
							</select>
						</div>
					</div>

					<div id="div_time" class="form-group">
						<label class="col-sm-2 col-sm-2 control-label"><span></span>快照时间</label>
						<div class="col-sm-10">
				    		<input id="fb_time" step=1 type="datetime-local" class="form-control">
						</div>
                	</div>
					
					<div class="form-group">
					<label class="col-lg-2 col-sm-2 control-label"></label>
					<div class="col-lg-10">
						<button type="button" onclick="checkUser(this)" class="btn btn-primary">开始恢复</button>
					</div>
					</div>
					</form>
                
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
  
  <div id="div_layer" style="display:none" ></div>
</section>
{{template "inc/foot.tpl" .}}    

<script src="/static/js/busy-load/app.min.js"></script>
<link href="/static/js/busy-load/app.min.css" rel="stylesheet"/>
<script>
var user_pwd = {{.user.Password}} ;
var query_url="/operation/dr_switch/process";
var bs_id = {{.bs_id}} ;
var db_id = {{.db_id}} ;
var	op_type = "STARTFLASHBACK";

// 初始化内容
$(function(){
	$("#div_point").hide();
	$("#div_time").hide();
});


function fb_method_change(e){
	if(e.value == 1){
		$("#div_point").show();
		$("#div_time").hide();
	}
	else{
		$("#div_time").show();
		$("#div_point").hide();
	}
}



function checkUser(e){
		var fb_method = $('#fb_method option:selected').val();
		var fb_point = $('#fb_point option:selected').val();
		var fb_time = $('#fb_time').val();
		var check_message = "";
		
		if(fb_method == "" || typeof(fb_method) == "undefined"){
			bootbox.alert({
					message: "请选择恢复方式!",
					buttons: {
								ok: {
									label: '确定',
									className: 'btn-success'
								}
							}
				});
			return false;
		}
		//空值检查
		if(fb_method == "1" && (fb_point == "" || typeof(fb_point) == "undefined")){
				bootbox.alert({
		        		message: "请选择快照点!",
		        		buttons: {
							        ok: {
							            label: '确定',
							            className: 'btn-success'
							        }
							    }
		        	});
		        	
		    return false;
		}
		if(fb_method == "2" && (fb_time == "" || typeof(fb_time) == "undefined")){
				bootbox.alert({
		        		message: "请选择快照时间!",
		        		buttons: {
							        ok: {
							            label: '确定',
							            className: 'btn-success'
							        }
							    }
		        	});
		        	
		    return false;
		}
		
		if(fb_method == 1){
			check_message = "确定需要恢复到快照点" + fb_point + "吗？"
		}else{
			fb_time = fb_time.replace(/T/i, " ");
			check_message = "确定需要恢复到快照时间" + fb_time + "吗？"
		}

		bootbox.prompt({
		    title: "请输入管理员密码!",
		    inputType: 'password',
		    callback: function (result) {
		    	if(result)
		    	{ 
		        if (md5(result) == user_pwd)
		        { 
							bootbox.dialog({
							    message: check_message,
							    buttons: {
							        ok: {
							            label: '确定',
							            className: 'btn-danger',
											callback: function(){
												$.ajax({
												url: "/operation/dr_recover/flashback",
												data: {"bs_id":bs_id,"db_id":db_id,"fb_method":fb_method,"fb_point":fb_point,"fb_time":fb_time},
												type: "POST",
												success: function (data) {
													//回调函数，判断提交返回的数据执行相应逻辑
													if (data.Success) {
													}
													else {
													}
												}
												});
		                						
												$('#div_layer').html("");			//初始化div
												mylay = layer.open({
													type: 1,
													skin: 'layui-layer-demo layblack', //样式类名
													closeBtn: 0, //不显示关闭按钮
													anim: 1,
													title: '详细过程',
													area: ['450px', '240px'],
													shadeClose: false, //开启遮罩关闭
													content: $('#div_layer')
												});
		                						oTimer = setInterval("queryHandle(query_url, bs_id, op_type)",2000);
		                        }
							        },
							        cancel: {
							            label: '取消',
							            className: 'btn-default',
							            callback: function () {
		                      }
							        }
							    }
							});
		        }
		        else
		        {
		        	bootbox.alert({
		        		message: "密码不对，请确认后重新尝试!",
		        		buttons: {
							        ok: {
							            label: '确定',
							            className: 'btn-success'
							        }
							    }
		        	});
		        }
		      }
		
		    }
		});

}

  
function queryHandle(url, bs_id, op_type){
    $.post(url, {"bs_id":bs_id, "op_type":op_type}, function(json){
        if(json.on_process == 0){
			if(json.op_type != ""){
				//alert(json.op_result);
				if(json.op_type == "STARTFLASHBACK"){
					if(json.op_reason == 'null'){
						error_message = "恢复快照失败，详细原因请查看相关日志";
					}else{
						error_message = "恢复快照失败，原因是：" + json.op_reason;
					}
						
					ok_message = "恢复快照成功";
				}else if(json.op_type == "STOPFLASHBACK"){
					if(json.op_reason == 'null'){
						error_message = "恢复同步失败，详细原因请查看相关日志";
					}else{
						error_message = "恢复同步失败，原因是：" + json.op_reason;
					}
				
					ok_message = "恢复同步成功";
				}
			
				if(json.op_result == '-1'){
					bootbox.alert({
					message: error_message,
					buttons: {
							ok: {
							label: '确定',
							className: 'btn-success'
							}
						},
						callback: function () {
						window.location.reload();
					}
					});
						
					if(mylay!=null){
					layer.close(mylay);
					}
					clearInterval(oTimer);
								
				}else if(json.op_result == '1'){
					bootbox.alert({
						message: ok_message,
						buttons: {
							ok: {
								label: '确定',
								className: 'btn-success'
							}
							},
						callback: function () {
							window.location.reload();
						}
					});
						
					if(mylay!=null){
					layer.close(mylay);
					}
					clearInterval(oTimer); 
				}else{
					if(mylay!=null){
					layer.close(mylay);
					}
					clearInterval(oTimer); 
				}
			}
        }else if(json.on_process == -1){
            bootbox.alert({
                message: "该系统没有配置容灾库",
                buttons: {
                      ok: {
                        label: '确定',
                        className: 'btn-success'
                      }
                    },
                  callback: function () {
                    window.location.reload();
                }
            });
                    
            if(mylay!=null){
              layer.close(mylay);
            }
            clearInterval(oTimer); 
        }

		if(mylay!=null){
			if(json.json_process == "null"){
						$("#div_layer").append("<p>" + json.json_process + "</p>");
			}else{
				localJson = $.parseJSON(json.json_process);
				//alert(localJson);
				$("#div_layer").empty();
				$.each(localJson,function(idx,item){   
				//alert("Time:"+item.Time+",Process_desc:"+item.Process_desc);   
					$("#div_layer").append("<p>" + item.Time + ": " + item.Process_desc + "</p>");
				});  
      		}

      		$(".layui-layer-content").scrollTop($(".layui-layer-content")[0].scrollHeight);
    	}
    },'json');  
}

</script>
</body>
</html>
