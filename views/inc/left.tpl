<!-- left side start-->
<div class="left-side sticky-left-side">
  <!--logo and iconic logo start-->
  
  <!--logo and iconic logo end-->
  <div class="left-side-inner">
    <!-- visible to small devices only -->
    <!-- <div class="visible-xs hidden-sm hidden-md hidden-lg">
      <div class="media logged-user"> <img alt="{{.LoginUsername}}" src="{{getAvatar .LoginAvatar}}" class="media-object">
        <div class="media-body">
          <h4><a href="/system/user/show/{{.LoginUserid}}">{{.LoginUsername}}</a></h4>
          <span>OPMS系统</span> </div>
      </div>
      <h5 class="left-nav-title">控制台</h5>
      <ul class="nav nav-pills nav-stacked custom-nav">
        <li><a href="/system/user/profile"><i class="fa fa-user"></i> <span>个人设置</span></a></li>
        <li><a href="/logout"><i class="fa fa-sign-out"></i> <span>退出</span></a></li>
      </ul>
    </div> -->
    <!--sidebar nav start-->


 <!-- new left-menu -->
 <div class="navbar-default sidebar navbarstyle" role="navigation">
  <div class="logo"> <a href="/"><img src="/static/img/logo-left.png" alt="DRM管理系统">DRM管理系统</a>  </div>
  <div class="logo-icon text-center"> <a href="/"><img src="/static/img/logo_icon.png" style="width:40px;" alt="DRM">DRM</a> </div>
  
  <!--<ul class="nav nav-pills nav-stacked custom-nav js-left-nav">-->
  <div class="sidebar-nav navbar-collapse">
  <ul class="nav" id="side-menu1">
      {{range $index, $elem := .leftNav}}
        {{if eq 1 $elem.IsShow}}
        <li>
              <a href="{{$elem.Url}}" data-icon="{{$elem.Icon}}" data-title="{{$elem.Name}}" {{if eq 1 $elem.IsActive}}class="active"{{end}} class="pointer" data-id="{{$elem.Id}}">
                <i class="fa {{$elem.Icon}}"></i> <span>{{$elem.Name}}</span>
              </a>
        </li>
        {{end}}
      {{end}}
  </ul>

  <!--
  <ul class="nav" id="side-menu">
        <li>
          <a href="#">
            <i class="iconfont icon-UI_icon_zonghe"></i> <b>UI demo</b> <span class="fa arrow"></span>
          </a>
          <ul class="nav nav-second-level">
            <li><a href="/demo/base"><b>基础表格页</b></a></li>
            <li><a href="/demo/index" ><b>基础UI</b></a></li>
            <li><a href="/demo/form"><b>表单</b></a></li>
            <li><a href="/demo/base">index</a></li>
          </ul>
        </li>
        <li>
          <a href="#">
            <i class="fa fa-home"></i> <b>一级菜单</b>
          </a>
        </li>
        <li>
          <a href="#">
            <i class="fa fa-user"></i>  <b>一级菜单</b> <span class="fa arrow"></span>
          </a>
          <ul class="nav nav-second-level">
            <li><a href="#"><b>二级菜单</b></a></li>
            <li><a href="#"  class="active"><b>二级菜单(激活)</b></a></li>
            <li><a href="#"><b>二级菜单</b></a></li>
          </ul>
        </li>
      </ul>-->
  </div>
  <!--sidebar nav end-->
  
</div>


  </div>
</div>
<!-- left side end-->


<script src="/static/js/layui/layui.js"></script>
<script>
    //JavaScript代码区域
    layui.use(['element','jquery','layer'], function(){

        //tips
        $(".pp-nav-childs").find('a').hover(function(){
            layer.tips($(this).attr('data-title'), $(this),{time:1000});
        });
    });
</script>
