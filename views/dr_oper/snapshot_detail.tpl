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
        <li><a href="/operation/dr_manage/list">容灾快照</a></li>
        <li class="active">容灾详细</li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
	<div class="pull-right">
		<button name="startsnapshot" class="btn btn-primary" type="button" value="StartSnapshot" onclick="checkUser(this)" data-id="{{.dr_id}}"> <i class="fa fa-reset"></i> 进入快照 </button>
		<button name="stopsnapshot" class="btn btn-danger" type="button" value="StopSnapshot" onclick="checkUser(this)" data-id="{{.dr_id}}"> <i class="fa fa-reset"></i> 退出快照 </button>
	</div>

	<div style="padding: 19px; " >
		<div style='padding: 20px 120px 0px 60px; height:150px; overflow:hidden'>
			<div style='float:left; height:100px; width:280px;'>
			<div><label name="pri_host" class="control-label" for="">IP：{{ .pri_config.Host }}</label></div>
			<div><label name="pri_dbname" class="control-label" for="">数据库名：{{ .pri_config.Instance_name }}</label></div>
			<div><label name="pri_dbstatus" class="control-label" for="">实例状态：{{ .pri_config.Open_Mode }}</label></div>
			<div><label name="pri_port" class="control-label" for="">监听端口：{{ .pri_config.Port }}</label></div>
			</div>
			<div style='float:right; height:100px; width:280px;'>
			<div><label name="sta_host" class="control-label" for="">IP：{{ .sta_config.Host }}</label></div>
			<div><label name="sta_dbname" class="control-label" for="">数据库名：{{ .sta_config.Instance_name }}</label></div>
			<div><label name="sta_dbstatus" class="control-label" for="">实例状态：{{ .sta_config.Open_Mode }}</label></div>
			<div><label name="sta_port" class="control-label" for="">监听端口：{{ .sta_config.Port }}</label></div>
			</div>
		</div>
	</div>

	<div style='padding: 5px 0px 0px 200px; height:150px;'>
		<div style="float:left;"><img {{if eq -1 .pri_config.Connect }}src="/static/img/connect_error.png"{{else}}src="/static/img/primary_db.png"{{end}} /></div> 

			<div style="float:left;">
			<div><label style='padding: 0px 0px 0px 120px;' class="control-label" for="">序列：{{ .sta_dr.Sequence }}</label></div>
			<div><img src="{{ getTransferStatus .sta_dr.DB_Id}}"> </div>
			</div> 
			
			<div style="float:left;"><img {{if eq -1 .sta_config.Connect }}src="/static/img/connect_error.png"{{else}}src="/static/img/standby_db.png"{{end}}/></div> 
			</div>


			<div style="float:left; width:340px; height:30px; border:0px solid red;">
			</div>
			<div id="mrp_warning" style="float:left; width:400px; height:30px; border:0px solid red; color:red; ">
				<label id="lb_warning" class="control-label" style="font-size:18px;color:red; padding: 5px 0px 0px 20px;"></label>
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
<script>
var mylay = null;
var oTimer = null; 

    
var user_pwd = {{.user.Password}} ;
var div_layer = document.getElementById("div_layer");
var query_url="";
var bs_id = -1;

function checkUser(e){
    bs_id = $(e).attr('data-id');
    
		if(e.value == "StartSnapshot"){
			_message = "确认要开始进入快照吗？";
      target_url = "/operation/dr_snapshot/startsnapshot";
			op_type = "STARTSNAPSHOT";
		}
		else if(e.value == "StopSnapshot"){
			_message = "确认要开始退出快照吗？";
      target_url = "/operation/dr_snapshot/stopsnapshot";
			op_type = "STOPSNAPSHOT";
		}
		else{
			return;
		}	
    

		bootbox.prompt({
		    title: "请确认密码",
		    inputType: 'password',
			buttons: {confirm: {label: "确认"}, cancel: {label: "取消"} },
		    callback: function (result) {
		    	if(result)
		    	{ 
		        if (md5(result) == user_pwd)
		        {
					bootbox.dialog({
						message: _message,
						buttons: {
							cancel: {
								label: '取消',
								className: 'btn-default',
								callback: function () {
								}
							},
							ok: {
								label: '确定',
								className: 'btn-danger',
								callback: function(){
									$.ajax({url: target_url,
											type: "POST",
											data: {"bs_id":bs_id,"asset_type":1},
											success: function (json) {
												//回调函数，判断提交返回的数据执行相应逻辑
												if (json.code == 0) {
                            bootbox.alert({
                              message: json.message,
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
		        		
		        		if(json.op_type == "STARTSNAPSHOT"){
                    if(json.op_reason == 'null'){
                      error_message = "进入快照失败，详细原因请查看相关日志";
                    }else{
                      error_message = "进入快照失败，原因是：" + json.op_reason;
                    }
							
							      ok_message = "进入快照成功";
		        		}else if(json.op_type == "STOPSNAPSHOT"){
                    if(json.op_reason == 'null'){
                      error_message = "退出快照失败，详细原因请查看相关日志";
                    }else{
                      error_message = "退出快照失败，原因是：" + json.op_reason;
                    }
                    
                    ok_message = "退出快照成功";
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
