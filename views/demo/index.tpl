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
          <!--toggle button start-->
          <a class="toggle-btn"><i class="fa fa-bars"></i></a>
          <!--toggle button end-->
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
            <li class="active">二级面包屑</li>
          </ul>
        </div>
        <!-- page heading end-->
        <!--body wrapper start-->
        <div class="wrapper">
          <div class="row">
            <div class="col-sm-12">
              <!-- 主体内容 开始 -->
              <div class="row">
                <div class="col-sm-12">
                  <div class="ibox">
                    <div class="ibox-title">颜色规范</div>
                    <div class="ibox-content">
                      <div class="democolor">
                        <p>
                          <span class="primary1">█primary1</span>
                          <span class="primary2">█primary2</span>
                          <span class="primary3">█primary3</span>
                        </p>
                        <p>
                          <span class="danger1">█danger1</span>
                          <span class="danger2">█danger2</span>
                          <span class="danger3">█danger3</span>
                        </p>
                        <p>
                          <span class="warning1">█warning1</span>
                          <span class="warning2">█warning2</span>
                          <span class="warning3">█warning3</span>
                        </p>
                        <p>
                          <span class="success1">█success1</span>
                          <span class="success2">█success2</span>
                          <span class="success3">█success3</span>
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col-sm-6">
                  <div class="ibox">
                    <div class="ibox-title">
                      <h5>排版</h5>
                    </div>
                    <div class="ibox-content">
                      <div>
                        <h4>标题</h4>
                        <div class="input-group">
                          <input type="text" class="form-control" />
                          <span class="input-group-btn">
                            <button type="button" class="btn btn-primary">
                              搜索
                            </button>
                          </span>
                        </div>
                        <h1>
                          标题 1
                          <small>二级标题</small>
                        </h1>
                        <h2>
                          标题 2
                          <small>二级标题</small>
                        </h2>
                        <h3>
                          标题 3
                          <small>二级标题</small>
                        </h3>
                        <h4>
                          标题 4
                          <small>二级标题</small>
                        </h4>
                        <h5>
                          标题 5
                          <small>二级标题</small>
                        </h5>
                        <h6>
                          标题 6
                          <small>二级标题</small>
                        </h6>
                      </div>
                      <div>
                        <h4>强调</h4>
                        <p class="text-muted">
                          每个人都有一个死角， 自己走不出来，别人也闯不进去。
                        </p>
                        <p class="text-primary">
                          每个人都有一个死角， 自己走不出来，别人也闯不进去。
                        </p>
                        <p class="text-success">
                          每个人都有一个死角， 自己走不出来，别人也闯不进去。
                        </p>
                        <p class="text-info">
                          每个人都有一个死角， 自己走不出来，别人也闯不进去。
                        </p>
                        <p class="text-warning">
                          每个人都有一个死角， 自己走不出来，别人也闯不进去。
                        </p>
                        <p class="text-danger">
                          每个人都有一个死角， 自己走不出来，别人也闯不进去。
                        </p>
                      </div>
                      <div>
                        <h4>引用</h4>
                        <blockquote>
                          <p>
                            大弦嘈嘈如急雨，小弦切切如私雨。二十四桥明月夜，玉人何处教吹箫。
                          </p>
                          <small
                            ><strong>杜牧</strong>-<cite
                              title=""
                              data-original-title=""
                              >《寄扬州韩绰判官》</cite
                            ></small
                          >
                        </blockquote>
                      </div>
                      <div>
                        <h4>对齐</h4>
                        <p class="text-left">左对齐文本</p>
                        <p class="text-center">居中文本</p>
                        <p class="text-right">右对齐文本</p>
                      </div>
                      <div>
                        <h4>分组列表</h4>
                        <ul class="list-group">
                          <li class="list-group-item">
                            <span class="badge badge-primary">16</span>
                            每个人都有一个死角， 自己走不出来，别人也闯不进去。
                          </li>
                          <li class="list-group-item ">
                            <span class="badge badge-info">12</span>
                            我把最深沉的秘密放在那里。
                          </li>
                          <li class="list-group-item">
                            <span class="badge badge-danger">10</span>
                            你不懂我，我不怪你。
                          </li>
                          <li class="list-group-item">
                            <span class="badge badge-success">10</span>
                            每个人都有一道伤口， 或深或浅，盖上布，以为不存在。
                          </li>
                        </ul>
                        <div class="list-group">
                          <a class="list-group-item active" href="#">
                            <h3 class="list-group-item-heading">
                              广西白龙洞题壁诗
                            </h3>

                            <p class="list-group-item-text">
                              挺身登峻岭，举目照遥空。毁佛崇天帝，移民复古风。临军称将勇，玩洞羡诗雄。剑气冲星斗，文光射日虹。
                            </p>
                          </a>

                          <a class="list-group-item" href="#">
                            <h3 class="list-group-item-heading">丑奴儿</h3>

                            <p class="list-group-item-text">
                              沉思十五年中事，才也纵横，泪也纵横，双负箫心与剑名。春来没个关心梦，自忏飘零，不信飘零，请看床头金字经。
                            </p>
                          </a>

                          <a class="list-group-item" href="#">
                            <h3 class="list-group-item-heading">夜登金山</h3>

                            <p class="list-group-item-text">
                              楼台两岸水相连，江北江南镜里天。芦管玉箫齐送夜，一声飞断月如烟。
                            </p>
                          </a>
                        </div>
                      </div>
                      <div>
                        <h4>Wells</h4>
                        <div class="well">
                          <h3>
                            默认效果
                          </h3>
                          把 Well 用在元素上，能有嵌入（inset）的的简单效果。
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col-sm-6">
                  <div class="ibox ">
                    <div class="ibox-title">
                      <h5>按钮</h5>
                    </div>
                    <div class="ibox-content">
                      <div>
                        <button type="button" class="btn btn-w-m btn-default">
                          btn-default
                        </button>
                        <button type="button" class="btn btn-w-m btn-primary">
                          btn-primary
                        </button>
                        <button type="button" class="btn btn-w-m btn-success">
                          btn-success
                        </button>
                        <button type="button" class="btn btn-w-m btn-info">
                          btn-info
                        </button>
                        <button type="button" class="btn btn-w-m btn-warning">
                          btn-warning
                        </button>
                        <button type="button" class="btn btn-w-m btn-danger">
                          btn-danger
                        </button>
                        <button type="button" class="btn btn-w-m btn-white">
                          btn-white
                        </button>
                        <button type="button" class="btn btn-w-m btn-link">
                          btn-link
                        </button>
                      </div>
                      <div>
                        <h4>按钮尺寸</h4>
                        <p>
                          可以通过添加class的值为<code>.btn-lg</code>,
                          <code>.btn-sm</code>, or
                          <code>.btn-xs</code>来修改按钮的大小
                        </p>
                        <p>
                          <button type="button" class="btn btn-primary btn-lg">
                            大按钮
                          </button>
                          <button type="button" class="btn btn-default btn-lg">
                            大按钮
                          </button>
                          <br />
                          <button type="button" class="btn btn-primary">
                            默认按钮
                          </button>
                          <button type="button" class="btn btn-default">
                            默认按钮
                          </button>
                          <br />
                          <button type="button" class="btn btn-primary btn-sm">
                            小按钮
                          </button>
                          <button type="button" class="btn btn-default btn-sm">
                            小按钮
                          </button>
                          <br />
                          <button type="button" class="btn btn-primary btn-xs">
                            Mini按钮
                          </button>
                          <button type="button" class="btn btn-default btn-xs">
                            Mini按钮
                          </button>
                        </p>
                        <p>
                          要使用线性按钮，可添加class<code>.btn-block</code>或<code
                            >.btn-outline</code
                          >
                        </p>

                        <h4>线性按钮</h4>
                        <p>
                          <button
                            type="button"
                            class="btn btn-outline btn-default"
                          >
                            默认
                          </button>
                          <button
                            type="button"
                            class="btn btn-outline btn-primary"
                          >
                            主要
                          </button>
                          <button
                            type="button"
                            class="btn btn-outline btn-success"
                          >
                            成功
                          </button>
                          <button
                            type="button"
                            class="btn btn-outline btn-info"
                          >
                            信息
                          </button>
                          <button
                            type="button"
                            class="btn btn-outline btn-warning"
                          >
                            警告
                          </button>
                          <button
                            type="button"
                            class="btn btn-outline btn-danger"
                          >
                            危险
                          </button>
                          <button
                            type="button"
                            class="btn btn-outline btn-link"
                          >
                            链接
                          </button>
                        </p>
                        <h3 class="font-bold">块级按钮</h3>
                        <p>
                          <button
                            type="button"
                            class="btn btn-block btn-outline btn-primary"
                          >
                            这是一个块级按钮
                          </button>
                        </p>
                      </div>
                      <!-- 下拉按钮 -->
                      <div>
                        <h4>下拉按钮</h4>
                        <div class="btn-group">
                          <button
                            data-toggle="dropdown"
                            class="btn btn-primary dropdown-toggle"
                          >
                            操作 <span class="caret"></span>
                          </button>
                          <ul class="dropdown-menu">
                            <li><a href="#">置顶</a></li>
                            <li>
                              <a href="#" class="font-bold">修改</a>
                            </li>
                            <li><a href="#">禁用</a></li>
                            <li class="divider"></li>
                            <li><a href="#">删除</a></li>
                          </ul>
                        </div>
                      </div>
                      <!-- 按钮组 -->
                      <div>
                        <h4>按钮组</h4>
                        <div class="btn-group">
                          <button class="btn btn-white" type="button">
                            左
                          </button>
                          <button class="btn btn-primary" type="button">
                            中
                          </button>
                          <button class="btn btn-white" type="button">
                            右
                          </button>
                        </div>
                        <br />
                        <br />
                        <div class="btn-group">
                          <button type="button" class="btn btn-white">
                            <i class="fa fa-chevron-left"></i>
                          </button>
                          <button class="btn btn-white">1</button>
                          <button class="btn btn-white  active">2</button>
                          <button class="btn btn-white">3</button>
                          <button class="btn btn-white">4</button>
                          <button type="button" class="btn btn-white">
                            <i class="fa fa-chevron-right"></i>
                          </button>
                        </div>
                      </div>

                      <!-- 图标按钮 -->
                      <div>
                        <h4>图标按钮</h4>
                        <button class="btn btn-primary " type="button">
                          <i class="fa fa-check"></i>&nbsp;提交
                        </button>
                        <button class="btn btn-success " type="button">
                          <i class="fa fa-upload"></i>&nbsp;&nbsp;<span
                            class="bold"
                            >上传</span
                          >
                        </button>
                        <button class="btn btn-info " type="button">
                          <i class="fa fa-paste"></i> 编辑
                        </button>
                        <button class="btn btn-warning " type="button">
                          <i class="fa fa-warning"></i>
                          <span class="bold">警告</span>
                        </button>
                      </div>
                      <!-- demo -->
                      <div>
                        <h4>demo</h4>
                      </div>
                    </div>
                  </div>
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
    {{template "inc/foot.tpl" .}}
  </body>
</html>
