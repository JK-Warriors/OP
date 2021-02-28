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
                  <span>数据库活跃会话数</span>
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
        <div class="pt_right" style="position: absolute; right: 0;">
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



<script type="text/javascript">

//圆环
window.onload = function () {
  new Progress().renderOne('canvas0', 100, 7, 60)
  new Progress().renderOne('canvas1', 100, 7, 50)
  //new Progress().renderOne('canvas2', 100, 7, 80)
  //new Progress().renderOne('canvas3', 100, 7, 75)
};


// .pt_2
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.pie'), 'chalk')
  // 配置
  var option = {
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)',
    },
    legend: {
      orient: 'vertical',
      left: 10,
      data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
    },
    series: [
      {
        name: 'DB Time',
        type: 'pie',
        radius: ['50%', '70%'],
        avoidLabelOverlap: false,
        label: {
          show: false,
          position: 'center',
        },
        emphasis: {
          label: {
            show: true,
            fontSize: '30',
            fontWeight: 'bold',
          },
        },
        labelLine: {
          show: false,
        },
        data: [{{range $k,$v := $.db_time}}
                { value: {{$v.Value}}, name: '{{$v.Alias}}' },
              {{end}}
        ],
        label: {
          color: 'rgba(255, 255, 255, 0.3)',
          formatter: '{d}%',
        },
        labelLine: {
          lineStyle: {
            color: 'rgba(255, 255, 255, 0.3)',
          },
          smooth: 0.2,
          length: 10,
          length2: 20,
        },
      },
    ],
  }
  myChart.setOption(option)
})()


// total sessions
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.pie2'), 'chalk')
  // 配置
  var option = {
    title: {
      text: '',
    },
    tooltip: {
      trigger: 'axis',
    },
     legend: {
          data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
     },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true,
    },
    toolbox: {
      feature: {},
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [{{range $k,$v := $.total_session_x}}{{$v.Time}},{{end}}],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {{range $k1,$v := .screen_db}}
      {
        name: {{$v.Alias}},
        type: 'line',
        data: [{{range $k2,$a := $.total_session_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
        ],
      },
      {{end}}
    ],
  }
  myChart.setOption(option)
})()


// active sessions
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.ct1'), 'chalk')
  // 配置
  var option = {
    title: {
      text: '',
    },
    tooltip: {
      trigger: 'axis',
    },
     legend: {
          data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
     },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      top:'6%',
      containLabel: true,
    },
    toolbox: {
      feature: {},
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [{{range $k,$v := $.active_session_x}}{{$v.Time}},{{end}}],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {{range $k1,$v := .screen_db}}
      {
        name: {{$v.Alias}},
        type: 'line',
        data: [{{range $k2,$a := $.active_session_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
        ],
      },
      {{end}}
    ],
  }
  myChart.setOption(option)
})()


// ct2 QPS
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.ct2'), 'chalk')
  // 配置
  var option = {
    title: {
      text: '',
    },
    tooltip: {
      trigger: 'axis',
    },
     legend: {
          data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
     },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      top:'6%',
      containLabel: true,
    },
    toolbox: {
      feature: {},
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [{{range $k,$v := $.qps_x}}{{$v.Time}},{{end}}],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {{range $k1,$v := .screen_db}}
      {
        name: {{$v.Alias}},
        type: 'line',
        data: [{{range $k2,$a := $.qps_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
        ],
      },
      {{end}}
    ],
  }
  myChart.setOption(option)
})()

// ct3 TPS
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.ct3'), 'chalk')
  // 配置
  var option = {
    title: {
      text: '',
    },
    tooltip: {
      trigger: 'axis',
    },
     legend: {
          data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
     },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      top:'6%',
      containLabel: true,
    },
    toolbox: {
      feature: {},
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [{{range $k,$v := $.tps_x}}{{$v.Time}},{{end}}],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {{range $k1,$v := .screen_db}}
      {
        name: {{$v.Alias}},
        type: 'line',
        data: [{{range $k2,$a := $.tps_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
        ],
      },
      {{end}}
    ],
  }
  myChart.setOption(option)
})()

// ct4 redo
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.ct4'), 'chalk')
  // 配置
  var option = {
    title: {
      text: '',
    },
    tooltip: {
      trigger: 'axis',
    },
     legend: {
          data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
     },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      top:'6%',
      containLabel: true,
    },
    toolbox: {
      feature: {},
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [{{range $k,$v := $.redo_x}}{{$v.Time}},{{end}}],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {{range $k1,$v := .screen_db}}
      {
        name: {{$v.Alias}},
        type: 'line',
        data: [{{range $k2,$a := $.redo_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
        ],
      },
      {{end}}
    ],
  }
  myChart.setOption(option)
})()

// ct5 buffer cache hit
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.ct5'), 'chalk')
  // 配置
  var option = {
    title: {
      text: '',
    },
    tooltip: {
      trigger: 'axis',
    },
     legend: {
          data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
     },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      top:'6%',
      containLabel: true,
    },
    toolbox: {
      feature: {},
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [{{range $k,$v := $.bch_x}}{{$v.Time}},{{end}}],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {{range $k1,$v := .screen_db}}
      {
        name: {{$v.Alias}},
        type: 'line',
        data: [{{range $k2,$a := $.bch_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
        ],
      },
      {{end}}
    ],
  }
  myChart.setOption(option)
})()

// ct6 tablespace 
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.ct6'), 'chalk')
  // 配置
  var option = {
    //    legend: {},
    grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        top:'6%',
        containLabel: true,
      },
  tooltip: {},
  dataset: {
  },
  xAxis: {type: 'category',
          data: ['Matcha Latte', 'Milk Tea', 'Cheese Cocoa', 'Walnut Brownie']
  },
  yAxis: {},
  // Declare several bar series, each will be mapped
  // to a column of dataset.source by default.
  series: [
    {
        type: 'bar',
        name: '2015',
        data: [89.3, 92.1, 94.4, 85.4]
    },
    {
        type: 'bar',
        name: '2016',
        data: [95.8, 89.4, 91.2, 76.9]
    },
    {
        type: 'bar',
        name: '2017',
        data: [97.7, 83.1, 92.5, 78.1]
    }
  ]
  }
  myChart.setOption(option)
})()

// ct7 容灾RPO
;(function () {
    // 返回对象
    var myChart = echarts.init(document.querySelector('.ct7'), 'chalk')
    // 配置
    var option = {
        title: {
          text: '',
        },
        tooltip: {
          trigger: 'axis',
        },
        legend: {
          data: [{{range $k,$v := $.screen_db}}{{$v.Alias}},{{end}}],
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true,
        },
        toolbox: {
          feature: {},
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: [{{range $k,$v := $.alert_x}}{{$v.Time}},{{end}}],
        },
        yAxis: {
          type: 'value',
        },
        series: [
                {{range $k1,$v := .screen_db}}
                {
                  name: {{$v.Alias}},
                  type: 'line',
                  data: [{{range $k2,$a := $.alert_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
                  ],
                },
                {{end}}
        ],
      }
    myChart.setOption(option)
  })()
</script>

</html>

