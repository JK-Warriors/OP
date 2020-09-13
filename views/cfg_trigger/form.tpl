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
        <li> <a href="/config/db/manage">配置中心</a> </li>
        <li> <a href="/config/cfg_trigger/manage">告警配置</a> </li>
        <li class="active"> {{if gt .triconf.Id 0}}编辑{{else}}新增{{end}}告警 </li>
      </ul>
      <!-- <div class="pull-right"><a href="/config/db/add" class="btn btn-success">+添加资产</a></div>-->
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="triconfig-form">
                <header><b> 基本信息 </b></header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>资产类型</label>
                  <div class="col-sm-10">
                    <select id="asset_type" name="asset_type" class="form-control">
                      <option value="">请选择类型</option>
                      <option value="1" {{if eq 1 .triconf.Asset_Type}}selected{{end}}>Oracle</option>
                      <option value="2" {{if eq 2 .triconf.Asset_Type}}selected{{end}}>MySQL</option>
                      <option value="3" {{if eq 3 .triconf.Asset_Type}}selected{{end}}>SQLServer</option>
                      <option value="99" {{if eq 99 .triconf.Asset_Type}}selected{{end}}>OS</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>告警类型</label>
                  <div class="col-sm-10">
                    <input type="text" id="trigger_type" name="trigger_type" readonly="readonly" value="{{.triconf.Trigger_Type}}" class="form-control">
                  </div>
                </div>
                <div id="div_severity" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>级别</label>
                  <div class="col-sm-10">
                    <select id="severity" name="severity" class="form-control">
                      <option value="">请选择协议</option>
                      <option value="Critical" {{if eq "Critical" .triconf.Severity}}selected{{end}}>Critical</option>
                      <option value="Warning" {{if eq "Warning" .triconf.Severity}}selected{{end}}>Warning</option>
                      <option value="Info" {{if eq "Info" .triconf.Severity}}selected{{end}}>Info</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>表达式</label>
                  <div class="col-sm-10">
                    <input type="text" id="expression" name="expression"  value="{{.triconf.Expression}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>描述</label>
                  <div class="col-sm-10">
                    <input type="text" id="description" name="description"  value="{{.triconf.Description}}" class="form-control">
                  </div>
                </div>
                
                <div id="div_mode" class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>恢复</label>
                  <div class="col-sm-10">
                    <select id="recovery_mode" name="recovery_mode" class="form-control">
                      <option value="">请选择协议</option>
                      <option value="1" {{if eq 1 .triconf.Recovery_Mode}}selected{{end}}>是</option>
                      <option value="0" {{if eq 0 .triconf.Recovery_Mode}}selected{{end}}>否</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>恢复表达式</label>
                  <div class="col-sm-10">
                    <input type="text" id="recovery_expression" name="recovery_expression"  value="{{.triconf.Recovery_Expression}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>恢复描述</label>
                  <div class="col-sm-10">
                    <input type="text" id="recovery_description" name="recovery_description"  value="{{.triconf.Recovery_Description}}" class="form-control">
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" id="id" name="id" value="{{.triconf.Id}}">
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



    $('#triconfig-form').validate({
        ignore:'',        
		    rules : {
			      //username:{required: true},
			      role:{required: true},
        },
        messages : {
			      //username:{required: '请填写用户名'},
			      role:{required: '请选择角色'}, 
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       		setTimeout(function(){window.location.href="/config/cfg_trigger/manage"}, 1000);
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
