<!DOCTYPE html>
<html lang="en">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
    <meta name="renderer" content="webkit" />
    <meta
      name="viewport"
      content="width=device-width,initial-scale=1.0,user-scalable=no"
    />
    <title>{{config "String" "globaltitle" ""}}</title>

    <link href="/static/css/dashboard.css" rel="stylesheet" />
    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/js/echarts.min.js"></script>
    <script src="/static/js/chalk.js"></script>

    <script type="text/javascript">
      ;(function(doc, win) {
        var docEl = doc.documentElement,
          resizeEvt =
            'orientationchange' in window ? 'orientationchange' : 'resize',
          recalc = function() {
            var clientWidth = docEl.clientWidth
            if (!clientWidth) return
            if (clientWidth >= 1920) {
              docEl.style.fontSize = '100px' //1rem  = 100px
            } else {
              docEl.style.fontSize = 100 * (clientWidth / 1920) + 'px'
            }
          }
        if (!doc.addEventListener) return
        win.addEventListener(resizeEvt, recalc, false)
        doc.addEventListener('DOMContentLoaded', recalc, false)
      })(document, window)

      $(function() {
        setInterval(function() {
          $('.scroll ul')
            .eq(0)
            .slideUp(400, function() {
              $(this)
                .appendTo($(this).parent())
                .show()
            })
        }, 2000)
      })
    </script>
  </head>

  <body style="visibility: visible;">
    <div class="container-flex">
      <div class="box">
        <div class="pagetit">
          <div class="time1">当前时间：&nbsp;<span id="Timer"></span></div>
          <div class="time2">
            最新检测时间：&nbsp;<span>2020-02-27 13:07:59</span>
          </div>
          <h1><a href="#">DRM监控平台</a></h1>
        </div>
        <div class="datanum">
          <!-- <div class="dtit">数据库的连接状态</div> -->
          <img src="/static/img/bj-1.png" alt="" class="bj-1" />
          <img src="/static/img/bj-2.png" alt="" class="bj-2" />
          <img src="/static/img/bj-3.png" alt="" class="bj-3" />
          <img src="/static/img/bj-4.png" alt="" class="bj-4" />

          <div class="scroll">
            <ul class="cf" style="">
              <li>
                <span>
                  <p>MySQL5.6_51</p>
                  <img src="/static/img/db2.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>MySQL5.6_52</p>
                  <img src="/static/img/db2.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>drmdb</p>
                  <img src="/static/img/db2.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>drmdb_dg</p>
                  <img src="/static/img/db2.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>Primary_115</p>
                  <img src="/static/img/db1.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>Standby_116</p>
                  <img src="/static/img/db2.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>Oracle_236_p</p>
                  <img src="/static/img/db1.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>Oracle237_s</p>
                  <img src="/static/img/db2.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>SQL4</p>
                  <img src="/static/img/db1.png" alt="" />
                </span>
              </li>
              <li>
                <span>
                  <p>SQL204</p>
                  <img src="/static/img/db1.png" alt="" />
                </span>
              </li>
            </ul>
          </div>
        </div>
        <div class="left1">
          <img src="/static/img/bj-1.png" alt="" class="bj-1" />
          <img src="/static/img/bj-2.png" alt="" class="bj-2" />
          <img src="/static/img/bj-3.png" alt="" class="bj-3" />
          <img src="/static/img/bj-4.png" alt="" class="bj-4" />

          <!--- 第一个 --->
          <div class="datarow cf">
            <div class="d1">
              <h1>drmdb</h1>
            </div>
            <div class="d2">
              <h2>性能指数</h2>

              <div id="left2" style="width: 100%;height:100%;"></div>
            </div>
            <div class="d3">
              <h2>Total Sessions 和 Active Sessions</h2>

              <div id="left3" style="width: 100%;height:100%;"></div>
            </div>
            <div class="d4">
              <h2>空间使用率</h2>
              <ul>
                <li>
                  <div class="progress">
                    <div class="progress-value">
                      SYSAUX:<span class="pdata"> 9.72%</span>
                    </div>
                    <div class="progress-bar">
                      <div class="progress-data" style="width: 9.72%;"></div>
                    </div>
                  </div>
                </li>
                <li>
                  <div class="progress">
                    <div class="progress-value">
                      SYSTEM:<span class="pdata"> 1.22%</span>
                    </div>
                    <div class="progress-bar">
                      <div class="progress-data" style="width: 1.22%;"></div>
                    </div>
                  </div>
                </li>
                <li>
                  <div class="progress">
                    <div class="progress-value">
                      UNDOTBS1:<span class="pdata"> 0.08%</span>
                    </div>
                    <div class="progress-bar">
                      <div class="progress-data" style="width: 0.08%;"></div>
                    </div>
                  </div>
                </li>
                <li>
                  <div class="progress">
                    <div class="progress-value">
                      USERS:<span class="pdata"> 0.05%</span>
                    </div>
                    <div class="progress-bar">
                      <div class="progress-data" style="width: 0.05%;"></div>
                    </div>
                  </div>
                </li>
                <li>
                  <div class="progress">
                    <div class="progress-value">
                      UNDOTBS2:<span class="pdata"> 0.03%</span>
                    </div>
                    <div class="progress-bar">
                      <div class="progress-data" style="width: 0.03%;"></div>
                    </div>
                  </div>
                </li>
              </ul>
            </div>
            <div class="d5">
              <h2>每小时日志量</h2>

              <div id="left5" style="width: 100%;height:100%;"></div>
            </div>
          </div>

          <!--- 第二个 --->

          <!--- 第三个 --->
        </div>
        <div class="right1">
          <div class="dtit">核心库指标</div>
          <img src="/static/img/bj-1.png" alt="" class="bj-1" />
          <img src="/static/img/bj-2.png" alt="" class="bj-2" />
          <img src="/static/img/bj-3.png" alt="" class="bj-3" />
          <img src="/static/img/bj-4.png" alt="" class="bj-4" />
          <div class="right11" id="right11"></div>
          <div class="right12" id="right12"></div>
        </div>
        <div class="right2">
          <div class="dtit">容灾状态</div>
          <img src="/static/img/bj-1.png" alt="" class="bj-1" />
          <img src="/static/img/bj-2.png" alt="" class="bj-2" />
          <img src="/static/img/bj-3.png" alt="" class="bj-3" />
          <img src="/static/img/bj-4.png" alt="" class="bj-4" />
          <ul>
            <li>
              <div class="progress">
                <div class="progress-name">Oracle</div>
                <div class="progress-bar">
                  <hr />
                  <hr />
                  <hr />
                  <hr />
                  <!-- color1,color2,color3分别对应正常,告警,异常的颜色 -->
                  <div
                    class="progress-data  cschedule_red"
                    style="height:100%"
                  ></div>
                </div>
              </div>
              <div class="proright">
                <ul class="cf">
                  <li class="co1">
                    <div>
                      <p>0</p>
                      <p>正常</p>
                    </div>
                  </li>
                  <li class="co2">
                    <div>
                      <p>0</p>
                      <p>告警</p>
                    </div>
                  </li>

                  <li class="co3">
                    <div>
                      <p>1</p>
                      <p>异常</p>
                    </div>
                  </li>
                </ul>
              </div>
            </li>
            <li>
              <div class="progress">
                <div class="progress-name">MySQL</div>
                <div class="progress-bar">
                  <hr />
                  <hr />
                  <hr />
                  <hr />
                  <!-- color1,color2,color3分别对应正常,告警,异常的颜色 -->
                  <div class="progress-data " style="height:0%"></div>
                </div>
              </div>
              <div class="proright">
                <ul class="cf">
                  <li class="co1">
                    <div>
                      <p>0</p>
                      <p>正常</p>
                    </div>
                  </li>
                  <li class="co2">
                    <div>
                      <p>0</p>
                      <p>告警</p>
                    </div>
                  </li>

                  <li class="co3">
                    <div>
                      <p>0</p>
                      <p>异常</p>
                    </div>
                  </li>
                </ul>
              </div>
            </li>
            <li>
              <div class="progress">
                <div class="progress-name">SQLServer</div>
                <div class="progress-bar">
                  <hr />
                  <hr />
                  <hr />
                  <hr />
                  <!-- color1,color2,color3分别对应正常,告警,异常的颜色 -->
                  <div
                    class="progress-data  cschedule_red"
                    style="height:100%"
                  ></div>
                </div>
              </div>
              <div class="proright">
                <ul class="cf">
                  <li class="co1">
                    <div>
                      <p>0</p>
                      <p>正常</p>
                    </div>
                  </li>
                  <li class="co2">
                    <div>
                      <p>0</p>
                      <p>告警</p>
                    </div>
                  </li>

                  <li class="co3">
                    <div>
                      <p>1</p>
                      <p>异常</p>
                    </div>
                  </li>
                </ul>
              </div>
            </li>
          </ul>
        </div>
        <div class="foot1">
          <div class="dtit">告警信息</div>

          <img src="/static/img/bj-1.png" alt="" class="bj-1" />
          <img src="/static/img/bj-2.png" alt="" class="bj-2" />
          <img src="/static/img/bj-3.png" alt="" class="bj-3" />
          <img src="/static/img/bj-4.png" alt="" class="bj-4" />
          <div class="finfo">
            <ul>
              <li>
                <span class="circlespan c1"></span>
                一级：红色
              </li>
              <li>
                <span class="circlespan c2"></span>
                二级：黄色
              </li>
              <li>
                <span class="circlespan c3"></span>
                三级（正常）：绿色
              </li>
            </ul>
          </div>
          <div class="fg-box" id="box">
            <ul style="margin-top: -4px;">
              <li class="c-danger">
                <table class="table">
                  <tbody>
                    <tr>
                      <td>2020-04-09 16:18:24</td>
                      <td>mysql</td>
                      <td>MySQL5.6_51</td>
                      <td>mysql server down</td>
                    </tr>
                  </tbody>
                </table>
              </li>
              <li class="c-warning">
                <table class="table">
                  <tbody>
                    <tr>
                      <td>2020-04-09 16:18:24</td>
                      <td>mysql</td>
                      <td>MySQL5.6_51</td>
                      <td>mysql server down</td>
                    </tr>
                  </tbody>
                </table>
              </li>
              <li class="c-success">
                <table class="table">
                  <tbody>
                    <tr>
                      <td>2020-04-09 16:18:24</td>
                      <td>mysql</td>
                      <td>MySQL5.6_51</td>
                      <td>mysql server down</td>
                    </tr>
                  </tbody>
                </table>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <script type="text/javascript">
      var left2 = echarts.init(document.getElementById('left2'), 'chalk')
      var option = {
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          orient: 'vertical',
          data: ['']
        },
        grid: {
          left: '3%',
          right: '6%',
          top: '30px',
          bottom: '0',
          containLabel: true
        },
        color: ['#a4d8cc', '#25f3e6'],
        toolbox: {
          show: false,
          feature: {
            mark: {
              show: true
            },
            dataView: {
              show: true,
              readOnly: false
            },
            magicType: {
              show: true,
              type: ['line', 'bar', 'stack', 'tiled']
            },
            restore: {
              show: true
            },
            saveAsImage: {
              show: true
            }
          }
        },

        calculable: true,
        xAxis: {
          splitLine: { show: false },
          type: 'category',
          boundaryGap: false,
          axisTick: {
            alignWithLabel: true
          },
          axisLine: {
            onZero: false,
            lineStyle: {}
          },
          axisLabel: {
            textStyle: {
              fontSize: '9'
            }
          },
          axisPointer: {
            label: {
              formatter: function(params) {
                return params.value + ' DB Time/Elapsed Time'
              }
            }
          },
          data: [
            '19:00',
            '20:00',
            '21:00',
            '22:00',
            '23:00',
            '00:00',
            '01:00',
            '02:00',
            '03:00',
            '04:00'
          ]
        },
        yAxis: {
          splitLine: { show: false },
          type: 'value',
          axisLabel: {
            textStyle: {
              fontSize: '9'
            }
          }
        },
        series: [
          {
            data: [
              '0.01',
              '0.01',
              '0.01',
              '0.01',
              '0.15',
              '0.01',
              '0.01',
              '0.01',
              '0.01',
              '0.01'
            ],
            type: 'line',
            areaStyle: {}
          }
        ]
      }
      left2.setOption(option)

      var left3 = echarts.init(document.getElementById('left3'), 'chalk')
      var option = {
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          orient: 'vertical',
          data: ['']
        },
        grid: {
          left: '3%',
          right: '6%',
          top: '30px',
          bottom: '0',
          containLabel: true
        },
        color: ['#5793f3', '#675bba'],
        toolbox: {
          show: false,
          feature: {
            mark: {
              show: true
            },
            dataView: {
              show: true,
              readOnly: false
            },
            magicType: {
              show: true,
              type: ['line', 'bar', 'stack', 'tiled']
            },
            restore: {
              show: true
            },
            saveAsImage: {
              show: true
            }
          }
        },

        calculable: true,
        xAxis: {
          splitLine: { show: false },
          type: 'category',
          boundaryGap: false,
          axisTick: {
            alignWithLabel: true
          },
          axisLine: {
            onZero: false,
            lineStyle: {}
          },
          axisLabel: {
            textStyle: {
              fontSize: '9'
            }
          },
          axisPointer: {
            label: {
              formatter: function(params) {
                return params.value + ' 会话'
              }
            }
          },
          data: [
            '21:00',
            '22:00',
            '23:00',
            '00:00',
            '01:00',
            '02:00',
            '03:00',
            '04:00',
            '16:00',
            '16:00'
          ]
        },
        yAxis: {
          splitLine: { show: false },
          type: 'value',
          axisLabel: {
            textStyle: {
              fontSize: '9'
            }
          }
        },
        series: [
          {
            symbol: 'circle', //设定为实心点
            symbolSize: 1, //设定实心点的大小
            name: 'total',
            data: ['49', '49', '49', '49', '48', '49', '49', '33', '50', '51'],
            type: 'line',
            areaStyle: {}
          },
          {
            symbol: 'circle', //设定为实心点
            symbolSize: 1, //设定实心点的大小
            name: 'active',
            data: ['47', '47', '47', '47', '46', '47', '47', '1', '47', '48'],
            type: 'line',
            areaStyle: {}
          }
        ]
      }
      left3.setOption(option)

      var left5 = echarts.init(document.getElementById('left5'), 'chalk')
      var option = {
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          orient: 'vertical',
          data: ['']
        },
        grid: {
          left: '3%',
          right: '6%',
          top: '30px',
          bottom: '0',
          containLabel: true
        },
        color: ['#a4d8cc', '#25f3e6'],
        toolbox: {
          show: false,
          feature: {
            mark: {
              show: true
            },
            dataView: {
              show: true,
              readOnly: false
            },
            magicType: {
              show: true,
              type: ['line', 'bar', 'stack', 'tiled']
            },
            restore: {
              show: true
            },
            saveAsImage: {
              show: true
            }
          }
        },

        calculable: true,
        xAxis: {
          splitLine: { show: false },
          type: 'category',
          boundaryGap: false,
          axisTick: {
            alignWithLabel: true
          },
          axisLine: {
            onZero: false,
            lineStyle: {}
          },
          axisLabel: {
            textStyle: {
              fontSize: '9'
            }
          },
          axisPointer: {
            label: {
              formatter: function(params) {
                return params.value + ' 日志量'
              }
            }
          },
          data: [
            '14:00',
            '15:00',
            '16:00',
            '17:00',
            '18:00',
            '19:00',
            '20:00',
            '21:00',
            '22:00',
            '23:00',
            '00:00',
            '01:00',
            '02:00',
            '03:00',
            '04:00',
            '05:00',
            '06:00',
            '07:00',
            '08:00',
            '09:00',
            '10:00',
            '11:00',
            '12:00',
            '13:00'
          ]
        },
        yAxis: {
          splitLine: { show: false },
          type: 'value',
          axisLabel: {
            textStyle: {
              fontSize: '9'
            }
          }
        },
        series: [
          {
            data: [
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0',
              '0'
            ],
            type: 'line',
            areaStyle: {}
          }
        ]
      }
      left5.setOption(option)

      //中间区域第二行

      //中间区域第三行

      var right11 = echarts.init(document.getElementById('right11'), 'chalk')
      var option = {
        backgroundColor: 'rgba(0,0,0,0)',
        tooltip: {
          trigger: 'item',
          formatter: '{b} <br/>{c} ({d}%)'
        },

        color: [
          '#af89d6',
          '#4ac7f5',
          '#0089ff',
          '#f36f8a',
          '#f5c847',
          '#ff5800',
          '#839557'
        ],
        //color: ['#ff0000', '#ff9600', '#ffff00', '#00ff00', '#00ff96', '#0000ff', '#ff00ff'],
        //[ "#ff5800", "#EAA228", "#4bb2c5", "#839557", "#958c12", "#953579", "#4b5de4", "#d8b83f", "#ff5800", "#0085cc"]
        legend: {
          orient: 'vertical',
          x: 'left',
          textStyle: {
            color: '#ccc'
          },
          data: []
        },
        series: [
          {
            name: '行业占比',
            type: 'pie',
            clockwise: false, //饼图的扇区是否是顺时针排布
            minAngle: 20, //最小的扇区角度（0 ~ 360）
            center: ['55%', '60%'], //饼图的中心（圆心）坐标
            radius: [0, '80%'], //饼图的半径
            avoidLabelOverlap: true, ////是否启用防止标签重叠
            itemStyle: {
              //图形样式
              normal: {
                borderColor: '#1e2239',
                borderWidth: 2
              }
            },
            label: {
              //标签的位置
              normal: {
                show: true,
                position: 'inside', //标签的位置
                formatter: '{d}%',
                textStyle: {
                  color: '#fff'
                }
              },
              emphasis: {
                show: true,
                textStyle: {
                  fontWeight: 'bold'
                }
              }
            },
            data: [
              {
                value: 1223,
                name: '2020-01-09: DB Time'
              },
              {
                value: 1216,
                name: '2020-01-10: DB Time'
              },
              {
                value: 1229,
                name: '2020-01-11: DB Time'
              },
              {
                value: 1761,
                name: '2020-01-12: DB Time'
              },
              {
                value: 1373,
                name: '2020-01-13: DB Time'
              },
              {
                value: 182,
                name: '2020-01-14: DB Time'
              },
              {
                value: 105,
                name: '2020-02-23: DB Time'
              }
            ]
          }
        ]
      }

      right11.setOption(option)

      var right12 = echarts.init(document.getElementById('right12'), 'chalk')
      var option = {
        tooltip: {},
        legend: {
          data: ['']
        },
        radar: {
          // shape: 'circle',
          name: {
            textStyle: {
              color: '#ccc'
            }
          },
          center: ['45%', '50%'],
          radius: 60,
          nameGap: 0,
          indicator: [
            { name: 'CPU', max: 100 },
            { name: '内存', max: 100 },
            { name: 'Swap', max: 100 },
            { name: '磁盘', max: 100 }
          ]
        },
        series: [
          {
            name: '核心主机性能指标（空闲率）',
            type: 'radar',
            areaStyle: {},
            data: [
              {
                value: [87, 19, 98, 19]
              }
            ]
          }
        ]
      }

      right12.setOption(option)

      window.addEventListener('resize', function() {
        left2.resize()
        left3.resize()
        left5.resize()
        left22.resize()
        left23.resize()
        left25.resize()
        left32.resize()
        left33.resize()
        left35.resize()
        right11.resize()
        right12.resize()
      })
    </script>

    <script type="text/javascript">
      $(document).ready(function() {
        // d4进度条
        var getValue = $('.pdata')
        for (var i = 0; i < getValue.length; i++) {
          var get_w = $(getValue[i]).text()
          $(getValue[i])
            .parent()
            .next()
            .find('.progress-data')
            .css('width', get_w)
        }

        var _box = $('#box')
        var _interval = 1000 //刷新间隔时间3秒
        function gdb() {
          $(
            "<p><span class='circlespan c3'></span>2019年4月25日 XXXX</p>"
          ).appendTo('#box')
          $('#box').scrollTop($('#box')[0].scrollHeight)
          /* var _last=$('#box dl dd:last');
        _last.animate({height: '+53px'}, "slow"); */
          setTimeout(function() {
            gdb()
          }, _interval)
        }
        // gdb();
      })

      // 滚动文字
      $(function() {
        setInterval('GetTime()', 1000)
      })
      //获取时间并设置格式
      function GetTime() {
        var mon, day, now, hour, min, ampm, time, str, tz, end, beg, sec
        /*
      mon = new Array("Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug",
              "Sep", "Oct", "Nov", "Dec");
      */
        mon = new Array(
          '一月',
          '二月',
          '三月',
          '四月',
          '五月',
          '六月',
          '七月',
          '八月',
          '九月',
          '十月',
          '十一月',
          '十二月'
        )
        /*
      day = new Array("Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat");
      */
        day = new Array('周日', '周一', '周二', '周三', '周四', '周五', '周六')
        now = new Date()
        hour = now.getHours()
        min = now.getMinutes()
        sec = now.getSeconds()
        if (hour < 10) {
          hour = '0' + hour
        }
        if (min < 10) {
          min = '0' + min
        }
        if (sec < 10) {
          sec = '0' + sec
        }
        $('#Timer').html(
          now.getFullYear() +
            '年' +
            (now.getMonth() + 1) +
            '月' +
            now.getDate() +
            '日' +
            '  ' +
            hour +
            ':' +
            min +
            ':' +
            sec
        )
        //$("#Timer").html(
        //        day[now.getDay()] + ", " + mon[now.getMonth()] + " "
        //                + now.getDate() + ", " + now.getFullYear() + " " + hour
        //                + ":" + min + ":" + sec);
      }
    </script>

    <script type="text/javascript">
      // 滚动效果
      ;(function($) {
        $.fn.myScroll = function(options) {
          var defaults = {
            speed: 40,
            rowHeight: 24
          }

          var opts = $.extend({}, defaults, options),
            intId = []

          function marquee(obj, step) {
            obj.find('ul').animate(
              {
                marginTop: '-=1'
              },
              0,
              function() {
                var s = Math.abs(parseInt($(this).css('margin-top')))
                if (s >= step) {
                  $(this)
                    .find('li')
                    .slice(0, 1)
                    .appendTo($(this))
                  $(this).css('margin-top', 0)
                }
              }
            )
          }

          this.each(function(i) {
            var sh = opts['rowHeight'],
              speed = opts['speed'],
              _this = $(this)
            intId[i] = setInterval(function() {
              if (_this.find('ul').height() <= _this.height()) {
                clearInterval(intId[i])
              } else {
                marquee(_this, sh)
              }
            }, speed)

            _this.hover(
              function() {
                clearInterval(intId[i])
              },
              function() {
                intId[i] = setInterval(function() {
                  if (_this.find('ul').height() <= _this.height()) {
                    clearInterval(intId[i])
                  } else {
                    marquee(_this, sh)
                  }
                }, speed)
              }
            )
          })
        }
      })(jQuery)
      $(function() {
        $('div.fg-box').myScroll({
          speed: 200, //数值越大，速度越慢
          rowHeight: 37 //li的高度
        })
      })
    </script>

    <script type="text/javascript">
      function refresh() {
        window.location.reload()
      }
      setTimeout('refresh()', 60000) //指定60秒刷新一次
    </script>
  </body>
</html>
