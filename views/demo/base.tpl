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
            <li><a href="/system/log/manage">一级面包屑</a></li>
            <li class="active">表格页</li>
          </ul>
        </div>
        <!-- page heading end-->
        <!--body wrapper start-->
        <div class="wrapper">
          <div class="row">
            <div class="col-sm-12">
              <!-- 主体内容 开始 -->
              <div class="tablepage">
                <div class="searchdiv">
                  <div class="search-form">
                    <div class="form-inline">
                      <div class="form-group">
                        <select class="form-control " name="account">
                          <option>选项 1</option>
                          <option>选项 2</option>
                          <option>选项 3</option>
                          <option>选项 4</option>
                        </select>
                        <input
                          type="email"
                          placeholder="搜索内容"
                          class="form-control"
                        />
                        <button class="btn btn-default" type="submit">
                          <i class="fa fa-search"></i> 搜索
                        </button>
                      </div>
                    </div>
                  </div>
                  <div class="pull-right">
                    <a href="#" class="btn btn-primary" id="layerdemo">
                      <i class="fa fa-plus"></i> 新增</a
                    >
                  </div>
                </div>
                <div class="tablediv">
                  <div class="choose-data">
                    <ul class="ui-choose" id="uc_01">
                      <li class="selected">全部</li>
                      <li>Oracle</li>
                      <li>MySQL</li>
                      <li>SQLserver</li>
                      <li>Linux</li>
                      <li>Windows</li>
                    </ul>
                  </div>
                  <table class="table table-bordered">
                    <thead>
                      <tr>
                        <th>名称</th>
                        <th>类型</th>
                        <th>IP</th>
                        <th>端口</th>
                        <th>用户名</th>
                        <th>数据库名</th>
                        <th>版本</th>
                        <th>监控状态</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td><a href="#">link</a></td>
                        <td>1</td>
                        <td>1</td>
                        <td>1</td>
                        <td>1</td>
                        <td>1</td>
                        <td>1</td>
                        <td>
                          <div class="switch">
                            <div class="onoffswitch">
                              <input
                                type="checkbox"
                                checked
                                class="onoffswitch-checkbox"
                                id="example1"
                              />
                              <label class="onoffswitch-label" for="example1">
                                <span class="onoffswitch-inner"></span>
                                <span class="onoffswitch-switch"></span>
                              </label>
                            </div>
                          </div>
                        </td>
                        <td>
                          <button class="table_btn">
                            <i class="iconfont icon-xianghujiaohuan"></i>文字
                          </button>
                          <button class="table_btn table_btn_icon">
                            <i class="iconfont icon-btn_edit"></i>编辑
                          </button>
                          <button class="table_btn table_btn_icon">
                            <i class="iconfont icon-iconfontshanchu"></i>删除
                          </button>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>

              <!-- 主体内容 结束 -->
            </div>
          </div>
        </div>
        <!--body wrapper end-->
      </div>
      <!-- main content end-->
    </section>
    <div id="box1" class="layui_drm">
      <div class="layercontent">
        <!-- layer content start -->
        <div class="form-horizontal adminex-form">
          <div class="form-group">
            <label class="col-xs-2  control-label">姓名</label>
            <div class="col-xs-10">
              <input
                type="text"
                name="realname"
                value=""
                class="form-control"
                placeholder="请填写姓名"
              />
            </div>
          </div>
          <div class="form-group">
            <label class="col-xs-2  control-label">姓名</label>
            <div class="col-xs-10">
              <input
                type="text"
                name="realname"
                value=""
                class="form-control"
                placeholder="请填写姓名"
              />
            </div>
          </div>
        </div>
        <!-- layer content end -->
      </div>
      <!-- layer foot start -->
      <div class="layerfoot">
        <button type="submit" class="btn btn-primary">提 交</button>
      </div>
      <!-- layer content end -->
    </div>
    {{template "inc/foot.tpl" .}}
  </body>
  <script>
    // 将所有.ui-choose实例化
    $('.ui-choose').ui_choose()
    // 无序列表单选
    var uc_01 = $('#uc_01').data('ui-choose')
    // 取回已实例化的对象
    uc_01.click = function(index, item) {
      console.log('click', index, item.text())
    }
    uc_01.change = function(index, item) {
      console.log('change', index, item.text())
    }

    //layer
    $(function() {
      $('#layerdemo').click(function() {
        layer.open({
          type: 1,
          closeBtn: true,
          shift: 2,
          area: ['600px', '70%'],
          shadeClose: true,
          content: $('#box1')
        })
      })
    })
  </script>
</html>
