<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>磐石智能监控平台</title>
    <link rel="stylesheet" href="/static/css/screen.css" />
  </head>
  <body>
    <div class="pagemain">
      <div class="maintit">
        <div class="time1 time">
          当前时间：
          <span id="Timer"></span>
        </div>
        <div class="time2 time">
          最新检测时间：
          <span>{{.lastchecktime}}</span>
        </div>
        <div class="webname"><a href="/">磐石智能监控平台</a></div>
      </div>

      <div class="p_content cf">
        <div class="pt_1">
          <div class="inner">
            {{range $k,$v := .assets}}
            <div class="item">
              <img src="/static/img/icon-normal.png" alt="" />
              <p>{{$v.Alias}}</p>
            </div>
            {{end}}
          </div>
        </div>
        <div class="pt_left">
          <div class="pt_2">
            <div class="tit">
              <span>过去7天数据库繁忙程度</span>
            </div>
            <div class="pie"></div>

            <div class="tit">
              <span>数据库会话数</span>
            </div>
            <div class="pie2"></div>
            <div class="canvasdiv">
              <div class="tit">
                <span>数据库健康度</span>
              </div>
              <ul class="cul cf">
                {{range $k,$v := .dbscore}}
                <li class="cli">
                  <canvas id="canvas{{$k}}"></canvas>
                  <p>{{$v.Alias}}</p>
                </li>
                {{end}}
              </ul>
            </div>
          </div>
        </div>
        <div class="pt_mid">
          <div class="pt_4">
            <div class="midcharts cf">
              <li class="c1">
                <div class="tit">
                  <span>Active session</span>
                </div>
                <div class="ct1 cc"></div>
              </li>
              <li class="c2">
                <div class="tit">
                  <span>性能指标QPS</span>
                </div>
                <div class="ct2 cc"></div>
              </li>
              <li class="c3">
                <div class="tit">
                  <span>性能指标TPS</span>
                </div>
                <div class="ct3 cc"></div>
              </li>
              <li class="c4">
                <div class="tit">
                  <span>每小时日志量</span>
                </div>
                <div class="ct4 cc"></div>
              </li>
              <li class="c5">
                <div class="tit">
                  <span>Buffer Cache命中率</span>
                </div>
                <div class="ct5 cc"></div>
              </li>
              <li class="c6">
                <div class="tit">
                  <span>表空间使用率</span>
                </div>
                <div class="ct6 cc"></div>
              </li>
            </div>
            <!--监控-->
            <div class="monitor">
              <div class="tit">
                <span>告警信息</span>
              </div>
              <ul class="mbow">
                <li>
                  <i></i>
                  一级：红色
                </li>
                <li>
                  <i></i>
                  二级：紫色
                </li>
                <li>
                  <i></i>
                  三级：蓝色
                </li>
              </ul>
              <div class="inner">
                <div class="content">
                  <div class="marquee-view">
                    <div class="marquee">
                      {{range $k,$v := .alerts}}
                      <div class="row">
                        <span class="span1">{{GetDateMHS $v.Created}}</span>
                        <span class="span2">[{{$v.Severity}}] [{{getDBDesc $v.Asset_Id}}] {{$v.Message}}</span>
                      </div>
                      {{end}}
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="pt_right">
          <div class="pt_5">
            <div class="c8">
              <div class="tit">
                <span>过去7天历史告警统计</span>
              </div>
              <div class="c8line">
                <table>
                  <colgroup>
                    <col style="width: 1%" />
                    <col />
                    <col style="width: 1%" />
                  </colgroup>
                  {{range $k,$v := .alertgroup}}
                  <tr>
                    <td>{{getDBAlias $v.Asset_Id}}</td>
                    <td>
                      <div class="bar">
                        <div class="barline" style="width: {{$v.Rate}}%"></div>
                      </div>
                    </td>
                    <td>{{$v.Count}}</td>
                  </tr>
                  {{end}}
                  
                </table>
              </div>
            </div>
            <div class="c7">
              <div class="tit">
                <span>容灾RPO</span>
              </div>
              <div class="ct7"></div>
            </div>
            <div class="c9">
              <div class="tit">
                <span>容灾状态</span>
              </div>
              <ul class="c9grid cf">
                <li>
                  <h4>Oracle</h4>
                  <div class="cbox">
                    <p class="box_num">{{.dr_ora_normal}}</p>
                    <p class="box_text">正常</p>
                  </div>
                  <div class="cbox box2">
                    <p class="box_num">{{.dr_ora_warning}}</p>
                    <p class="box_text">告警</p>
                  </div>
                  <div class="cbox box3">
                    <p class="box_num">{{.dr_ora_critical}}</p>
                    <p class="box_text">异常</p>
                  </div>
                </li>
                <li>
                  <h4>MySQL</h4>
                  <div class="cbox">
                    <p class="box_num">{{.dr_mysql_normal}}</p>
                    <p class="box_text">正常</p>
                  </div>
                  <div class="cbox box2">
                    <p class="box_num">{{.dr_mysql_warning}}</p>
                    <p class="box_text">告警</p>
                  </div>
                  <div class="cbox box3">
                    <p class="box_num">{{.dr_mysql_critical}}</p>
                    <p class="box_text">异常</p>
                  </div>
                </li>
                <li>
                  <h4>SQLServer</h4>
                  <div class="cbox">
                    <p class="box_num">{{.dr_mssql_normal}}</p>
                    <p class="box_text">正常</p>
                  </div>
                  <div class="cbox box2">
                    <p class="box_num">{{.dr_mssql_warning}}</p>
                    <p class="box_text">告警</p>
                  </div>
                  <div class="cbox box3">
                    <p class="box_num">{{.dr_mssql_critical}}</p>
                    <p class="box_text">异常</p>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <div></div>
    </div>

    <script type="text/javascript" src="/static/js/jquery.min.js"></script>
    <script type="text/javascript" src="/static/js/echarts.min.js"></script>
    <script type="text/javascript" src="/static/js/Progress.js"></script>
    <script type="text/javascript" src="/static/js/chalk.js"></script>
    <script type="text/javascript" src="/static/js/screen.js"></script>
  </body>
</html>
