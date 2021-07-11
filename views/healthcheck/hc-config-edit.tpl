<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}

</head>
<body class="sticky-header">
<section> 
    <div class="layui-tab layui-tab-brief" lay-filter="tab_hc">
      <ul class="layui-tab-title">
        <li class="layui-this">巡检计划</li>
      </ul>
      <div class="layui-tab-content" style="height: 400px;">
        <div class="layui-tab-item layui-show">
          <div class="layui-form-item">
            <label class="layui-form-label"><span class="layui-btn layui-btn-primary">定时巡检</span></label>
            <div class="layui-input-block">
		          <input name="status" type="checkbox" data-size="small"> 
              <input name="open" type="checkbox" lay-skin="switch" lay-filter="switchTest" lay-text="ON|OFF">
            </div>
            <div class="layui-input-block">
              <input type="checkbox" name="close" lay-skin="switch" lay-text="ON|OFF">
            </div>
          </div>

          <div class="layui-form-item">
            <label class="layui-form-label"><span class="layui-btn layui-btn-primary">巡检时间</span></label>
            <div class="layui-input-block">
              <input type="text" class="layui-input" id="task_time" placeholder="请填写crontab表达式">
            </div>
          </div>
          
            
            <div class="form-group">
              <label class="col-lg-2 col-sm-2 control-label"></label>
              <div class="col-lg-10">
                <button type="submit" class="btn btn-primary">提 交</button>
              </div>
            </div>
          </div>

        </div>
        
      </div>
    </div>
</section>

<script>

</script>
</body>
</html>
