<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <!-- <h3> 组织管理 {{template "users/nav.tpl" .}}</h3>-->
      <ul class="breadcrumb pull-left">
        <li> <a href="/config/dr_business/manage">容灾配置</a> </li>
        <li> <a href="/config/dr_config/manage">容灾配置</a> </li>
        <li class="active"> {{if gt .drconf.Bs_Id 0}}编辑{{else}}新增{{end}}容灾组 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="dr_config-form">
                <header><b> 基本信息 </b></header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>容灾组名</label>
                  <div class="col-sm-10">
                    <input type="text" id="bs_name" name="bs_name"  value="{{.drconf.Bs_Name}}" class="form-control">
                  </div>
                </div>
                
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>数据库类型</label>
                  <div class="col-sm-10">
                    <select id="asset_type" name="asset_type" class="form-control">
                      <option value="">请选择主库</option>
                      <option value="1" {{if eq 1 $.drconf.Asset_Type}}selected{{end}}>Oracle</option>
                      <option value="2" {{if eq 2 $.drconf.Asset_Type}}selected{{end}}>MySQL</option>
                      <option value="3" {{if eq 3 $.drconf.Asset_Type}}selected{{end}}>SQLServer</option>
                    </select>
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主库</label>
                  <div class="col-sm-10">
                    <select id="db_id_p" name="db_id_p" class="form-control">
                      <option value="">请选择主库</option>
                      {{range $k,$v := .pridbconf }}
                        <option value="{{$v.Id}}" {{if eq $v.Id $.drconf.Db_Id_P}}selected{{end}}>{{$v.Host}}:{{$v.Port}} ({{$v.Alias}})</option>
                      {{end}}
                    </select>
                  </div>
                </div>
                <div id="div_db_dest_p" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>主库通道</label>
                  <div class="col-sm-10">
                    <select id="db_dest_p" name="db_dest_p" class="form-control">
                      {{range $k := .dest_list }}
                        <option value="{{$k}}" {{if eq $k $.drconf.Db_Dest_P}}selected{{end}}>{{$k}}</option>
                      {{end}}
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>备库</label>
                  <div class="col-sm-10">
                    <select id="db_id_s" name="db_id_s" class="form-control">
                      <option value="">请选择备库</option>
                      {{range $k,$v := .stadbconf }}
                        <option value="{{$v.Id}}" {{if eq $v.Id $.drconf.Db_Id_S}}selected{{end}}>{{$v.Host}}:{{$v.Port}} ({{$v.Alias}})</option>
                      {{end}}
                    </select>
                  </div>
                </div>
                <div id="div_db_dest_s" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>备库通道</label>
                  <div class="col-sm-10">
                    <select id="db_dest_s" name="db_dest_s" class="form-control">
                      {{range $k := .dest_list }}
                        <option value="{{$k}}" {{if eq $k $.drconf.Db_Dest_S}}selected{{end}}>{{$k}}</option>
                      {{end}}
                    </select>
                  </div>
                </div>
                <div id="div_fb_retention" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>闪回天数</label>
                  <div class="col-sm-10">
                    <input type="text" name="fb_retention"  value="{{.drconf.Fb_Retention}}" class="form-control">
                  </div>
                </div>
                
                <div id="div_db_name" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>数据库名</label>
                  <div class="col-sm-10">
                    <input type="text" id="db_name" name="db_name"  value="{{.drconf.Db_Name}}" class="form-control">
                  </div>
                </div>


                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>是否漂移IP</label>
                  <div class="col-sm-10">
                    <input type="checkbox" id="is_shift" name="is_shift">
                  </div>
                </div>
                <div id="div_shift_vips" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>漂移IP</label>
                  <div class="col-sm-10">
                    <input type="text" id="shift_vips" name="shift_vips"  class="form-control" placeholder="请填写漂移IP, 多个IP用逗号分隔">
                  </div>
                </div>
                <div id="div_network_p" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>主库网卡名</label>
                  <div class="col-sm-10">
                    <input type="text" id="network_p" name="network_p"  class="form-control" placeholder="请填写主库网卡名">
                  </div>
                </div>
                <div id="div_network_s" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>备库网卡名称</label>
                  <div class="col-sm-10">
                    <input type="text" id="network_s" name="network_s"  class="form-control" placeholder="请填写备库网卡名">
                  </div>
                </div>
                
                <div id="div_alarm" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>是否发送告警 </label>
                  <div>
                    <input type="checkbox" id="is_alert" value="1" name="is_alert" {{if eq 1 .drconf.Is_Alert}}checked="checked"{{end}}>发送
                  </div>
                </div>

                <div id="div_alarm_media" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span></span>发送媒介 </label>
                  <div>
                    <input type="checkbox" id="alert_mail" value="1" name="alert_mail" {{if eq 1 .drconf.Alert_Mail}}checked="checked"{{end}}>邮件
                    <input type="checkbox" id="alert_wechat" value="1" name="alert_wechat" {{if eq 1 .drconf.Alert_WeChat}}checked="checked"{{end}}>微信
                    <input type="checkbox" id="alert_sms" value="1" name="alert_sms" {{if eq 1 .drconf.Alert_SMS}}checked="checked"{{end}}>短信
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" id="bs_id" name="bs_id" value="{{.drconf.Bs_Id}}">
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
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
<script src="/static/js/jquery-ui-1.10.3.min.js"></script>
<script>
    $(function() {// 初始化内容
        id =  {{.drconf.Bs_Id}};
        is_shift =  {{.drconf.Is_Shift}};
        shift_vips =  {{.drconf.Shift_Vips}};
        network_p =  {{.drconf.Network_P}};
        network_s =  {{.drconf.Network_S}};
        
        if(is_shift == 1){       
            $("input[name='is_shift']").prop("checked",true)  
            $("#div_shift_vips").show();
            $("#div_network_p").show();
            $("#div_network_s").show();
            $("#shift_vips").attr("value",shift_vips);
            $("#network_p").attr("value",network_p);
            $("#network_s").attr("value",network_s);
        }else{
            $("input[name='is_shift']").prop("checked",false)  
            $("#div_shift_vips").hide();
            $("#div_network_p").hide();
            $("#div_network_s").hide();
            $("#shift_vips").attr("value","");
            $("#network_p").attr("value","");
            $("#network_s").attr("value","");
        }

        
        if($("#asset_type option:selected").val() == 3){
            $("#div_db_name").show();
        }else{
            $("#div_db_name").hide();
        }
        if($("#asset_type option:selected").val() == 1){
            $("#div_db_dest_p").show();
            $("#div_db_dest_s").show();
            $("#div_fb_retention").show();
        }else{
            $("#div_db_dest_p").hide();
            $("#div_db_dest_s").hide();
            $("#div_fb_retention").hide();
        }

        if (id && id > 0){
        }else{
          $("#is_alert").attr("checked",true);
        }

        if($('#is_alert').prop('checked')){         
            $("#div_alarm_media").show();
        }else{
            $("#div_alarm_media").hide();
        }
    });  

    $("#asset_type").change(function(){
        if($("#asset_type option:selected").val() == 3){
            $("#div_db_name").show();
        }else{
            $("#div_db_name").hide();
        }
        
        if($("#asset_type option:selected").val() == 1){
            $("#div_db_dest_p").show();
            $("#div_db_dest_s").show();
            $("#div_fb_retention").show();
        }else{
            $("#div_db_dest_p").hide();
            $("#div_db_dest_s").hide();
            $("#div_fb_retention").hide();
        }
    });

    $("#is_shift").change(function(){
        if($('#is_shift').prop('checked')){         
            $("#div_shift_vips").show();
            $("#div_network_p").show();
            $("#div_network_s").show();
            $("#shift_vips").attr("value","");
            $("#network_p").attr("value","");
            $("#network_s").attr("value","");
        }else{
            $("#div_shift_vips").hide();
            $("#div_network_p").hide();
            $("#div_network_s").hide();
            $("#shift_vips").attr("value","");
            $("#network_p").attr("value","");
            $("#network_s").attr("value","");
        }
    });
    
    $("#is_alert").change(function(){
        if($('#is_alert').prop('checked')){         
            $("#div_alarm_media").show();
        }else{
            $("#div_alarm_media").hide();
            $("#alert_mail").attr("value","");
            $("#alert_wechat").attr("value","");
            $("#alert_sms").attr("value","");
        }
    });
       
    $('#dr_config-form').validate({
        ignore:'',        
		    rules : {
			      db_id_p:{required: true},
			      db_id_s:{required: true},
        },
        messages : {
			      db_id_p:{required: '请选择主库'},
			      db_id_s:{required: '请选择备库'}, 
        },

        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       		setTimeout(function(){window.location.href="/config/dr_config/manage"}, 1000);
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
