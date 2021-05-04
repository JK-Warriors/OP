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
        <li class="layui-this">上下班巡检</li>
        <li>全面巡检</li>
      </ul>
      <div class="layui-tab-content" style="height: 400px;">
        <div class="layui-tab-item layui-show">
          <div class="layui-form-item">
            <label class="layui-form-label"><span class="layui-btn layui-btn-primary">上下班巡检</span></label>
            <div class="layui-input-block">
              <input type="checkbox" checked="" name="open" lay-skin="switch" lay-filter="switchTest" lay-text="ON|OFF">
            </div>
          </div>
          
          <div class="layui-form-item">
            <label class="layui-form-label"><span class="layui-btn layui-btn-primary">上班巡检时间</span></label>
            <div class="layui-input-block">
              <input type="text" class="layui-input" id="onduty_time">
            </div>
          </div>

          <div class="layui-form-item">
            <label class="layui-form-label"><span class="layui-btn layui-btn-primary">下班巡检时间</span></label>
            <div class="layui-input-block">
              <input type="text" class="layui-input" id="offduty_time">
            </div>
          </div>
          
          <div class="layui-form-item">
            <label class="layui-form-label"><span class="layui-btn layui-btn-primary">巡检周期</span></label>
            <div class="layui-input-block">
              <input type="checkbox" checked="" name="check_all" title="全选">
            </div>
            <div class="layui-input-block">
              <input type="checkbox" checked="" name="day" lay-skin="primary" title="周一">
              <input type="checkbox" checked="" name="day" lay-skin="primary" title="周二">
              <input type="checkbox" checked="" name="day" lay-skin="primary" title="周三">
              <input type="checkbox" checked="" name="day" lay-skin="primary" title="周四">
              <input type="checkbox" checked="" name="day" lay-skin="primary" title="周五">
              <input type="checkbox" checked="" name="day" lay-skin="primary" title="周六">
              <input type="checkbox" checked="" name="day" lay-skin="primary" title="周日">
            </div>
            
            <div class="form-group">
              <label class="col-lg-2 col-sm-2 control-label"></label>
              <div class="col-lg-10">
                <button type="submit" class="btn btn-primary">提 交</button>
              </div>
            </div>
          </div>

        </div>
        
        <div class="layui-tab-item">
          <div class="layui-form-item">
            <label class="layui-form-label"><span class="layui-btn layui-btn-primary">全面巡检</span></label>
            <div class="layui-input-block">
              <input type="checkbox" checked="" name="open" lay-skin="switch" lay-filter="switchTest" lay-text="ON|OFF">
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
layui.use('element', function(){
  var $ = layui.jquery
  ,element = layui.element; //Tab的切换功能，切换事件监听等，需要依赖element模块
});

layui.use('laydate', function(){
  var laydate = layui.laydate;
  
  //执行一个laydate实例
  laydate.render({
    elem: '#onduty_time' //指定元素
    ,type: 'time'
  });

  laydate.render({
    elem: '#offduty_time' //指定元素
    ,type: 'time'
  });

});

</script>
</body>
</html>
