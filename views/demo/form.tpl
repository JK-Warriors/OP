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
            <li class="active">form</li>
          </ul>
        </div>
        <!-- page heading end-->
        <!--body wrapper start-->
        <div class="wrapper">
          <div class="row">
            <div class="col-sm-12">
              <!-- 主体内容 开始 -->
              <div class="row">
                <div class="col-sm-6">
                  <div class="ibox">
                    <div class="ibox-title">
                      <h5>排版</h5>
                    </div>
                    <div class="ibox-content">
                      <div>
                        <h4>标题</h4>
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
                <div class="col-sm-6">aa</div>
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
