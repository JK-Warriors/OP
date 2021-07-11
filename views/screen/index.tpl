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
              {{if eq 1 $v.Connect }}
                <div class="item">
                  <img src="/static/img/icon-normal.png" alt="" />
                  <p>{{$v.Alias}}</p>
                </div>
              {{else}}
                <div class="item">
                  <img src="/static/img/icon-abnormal.png" alt="" />
                  <p>{{$v.Alias}}</p>
                </div>
              {{end}}
            {{end}}
          </div>
        </div>
        <div class="pt_left">
          <div class="pt_2">
            <div class="tit">
              <span>过去7天{{getDBAlias .core_db}}库Db Time</span>
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
                  二级：橙色
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
                        <span class="span2 {{if eq "Warning" $v.Severity}}colororange{{else if eq "Info" $v.Severity}}colorblue{{end}}">[{{$v.Severity}}] [{{ $v.Asset_Desc}}] {{$v.Message}}</span>
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
            <div class="c8" style="min-height:170px">
              <div class="tit">
                <span>过去5天历史告警统计</span>
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
                    <td>{{$v.Date}}</td>
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
                <span>容灾RTO</span>
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

function refresh()
{
       window.location.reload();
}
setTimeout('refresh()',60000); //指定60秒刷新一次

//圆环
window.onload = function () {
  new Progress().renderOne('canvas0', 100, 7, 95)
  new Progress().renderOne('canvas1', 100, 7, 90)
  new Progress().renderOne('canvas2', 100, 7, 93)
  new Progress().renderOne('canvas3', 100, 7, 88)
};


// .pt_2
;(function () {
  // 返回对象
  var myChart = echarts.init(document.querySelector('.pie'), 'chalk')
  // 配置
  var option = {
    //backgroundColor: 'rgba(0,0,0,0)',
    //color: ['#af89d6', '#4ac7f5', '#0089ff', '#f36f8a', '#f5c847','#ff5800','#839557'],
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)',
    },
    legend: {
        orient: 'vertical',
        x: 'left',
        textStyle: {
            color: '#ccc'
        },
        data:[]
    },
    series: [
      {
        name: 'DB Time',
        type: 'pie',
        radius: ['50%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: { //图形样式
            normal: {
                borderColor: '#1e2239',
                borderWidth: 2,
            },
        },

        data: [{{range $k,$v := $.db_time}}
                { value: {{$v.Db_Time}}, name: '{{$v.End_Time}}' },
              {{end}}
        ],
           /*
           data: [
          { value: 335, name: '直接访问' },
          { value: 310, name: '邮件营销' },
          { value: 234, name: '联盟广告' },
          { value: 135, name: '视频广告' },
          { value: 1548, name: '搜索引擎' },
        ],
        */
        label: {
            normal: {
                show: true,
                position: 'inside', //标签的位置
                formatter: "{d}%",
                textStyle: {
                    color: '#fff',
                }
            },
            emphasis: {
                show: true,
                textStyle: {
                    fontWeight: 'bold'
                }
            }
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
      top:'10%',
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
      top:'10%',
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
      top:'10%',
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
      top:'10%',
      containLabel: true,
    },
    toolbox: {
      feature: {},
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [{{range $k,$v := $.log_per_sec_x}}{{$v.Time}},{{end}}],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {{range $k1,$v := .screen_db}}
      {
        name: {{$v.Alias}},
        type: 'line',
        data: [{{range $k2,$a := $.log_per_sec_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
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
      top:'12%',
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
          data: [{{range $k,$v := $.tbs}}{{$v.Tbs_Name}},{{end}}],
           axisLabel: {
                             interval:0,
                             rotate:40,
                              formatter: function(value) {
                                 var res = value;
                                 if(res.length > 5) {
                                     res = res.substring(0, 4) + "..";
                                 }
                                 return res;
                             }
                         }
  },
  yAxis: {
        type: 'value'
  },
  visualMap: {
        orient: 'horizontal',
        left:-9999,
        min: 0,
        max: 10,
        // Map the score column to color
        dimension: 0,
        inRange: {
            color: ['#cb31b0', '#28b1ff']
        }
    },
  // Declare several bar series, each will be mapped
  // to a column of dataset.source by default.
  series: [
    {
        type: 'bar',
        data: [{{range $k,$v := $.tbs}}{{$v.Rate}},{{end}}],
       
        //data: [2,2,1,1,0,0,0,0]
    }
  ]
  }
  myChart.setOption(option)
})()

// ct7 容灾RTO
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
          data: [{{range $k,$v := $.dr}}{{$v.Name}},{{end}}],
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
          data: [{{range $k,$v := $.rto_x}}{{$v.Time}},{{end}}],
        },
        yAxis: {
          type: 'value',
        },
        series: [
                {{range $k1,$v := .dr}}
                {
                  name: {{$v.Name}},
                  type: 'line',
                  data: [{{range $k2,$a := $.rto_y}}{{if eq $v.Id $a.Db_Id}}{{$a.Value}},{{end}}{{end}}
                  ],
                },
                {{end}}
        ],
      }
    myChart.setOption(option)
  })()

</script>

</html>

