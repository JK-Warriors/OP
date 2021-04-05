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
      <div class="layui-tab-content" style="height: 100px;">
        <div class="layui-tab-item layui-show">内容不一样是要有，因为你可以监听tab事件（阅读下文档就是了）
        </div>
        <div class="layui-tab-item">
          <div>
            <div class="layui-form-item">
              <label class="layui-form-label"><span class="layui-btn layui-btn-primary">全面巡检</span></label>
              <div class="layui-input-block">
                <input type="checkbox" checked="" name="open" lay-skin="switch" lay-filter="switchTest" lay-text="ON|OFF">
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

</script>
</body>
</html>
