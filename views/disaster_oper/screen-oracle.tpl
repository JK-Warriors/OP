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

    <link href="/static/css/dgscreen.css" rel="stylesheet" />
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
      <div class="boxtit">
        <div class="b1">当前时间：&nbsp;<span id="Timer"></span></div>
        <div class="b2">数据平台数据库容灾监控中心</div>
        <div class="b3">
          最新检测时间：&nbsp;<span>2020-04-27 10:10:23</span>

          <a href="#" class="qbtn"
            ><img src="/static/img/quit.png" title="退出"
          /></a>
        </div>
        <div class="b4">
          <h3>灾备状态</h3>
          <div class="m">
            <div class="m2">
              <div class="container ">
                <div class="item-1"></div>
                <div class="item-2"></div>
                <div class="item-3"></div>
                <div class="item-4"></div>
                <div class="item-5"></div>
              </div>
              <p>
                <span class="spancolor">日志应用:</span> 序列: 167209 块号:
                172469
              </p>
              <p><span class="spancolor">传输模式:</span> 异步模式</p>
            </div>
            <div class="m1 ml">
              <p>生产系统</p>
              <img src="/static/img/database.png" alt="" />
              <div class="mtext">
                <li><span>数据库时间:</span> 2020-04-27 10:11:58</li>
                <li><span>实例名:</span> ORCL</li>
                <li><span>IP地址:</span> 10.170.24.159</li>
                <li><span>数据库版本:</span> 11.2.0.4.0</li>
              </div>
            </div>
            <div class="m1 mr">
              <p>灾备系统</p>
              <img src="/static/img/database.png" alt="" />
              <div class="mtext">
                <li><span>数据库时间:</span> 2020-04-27 10:11:54</li>
                <li><span>实例名:</span> ORCL</li>
                <li><span>IP地址:</span> 10.170.24.173</li>
                <li><span>数据库版本:</span> 11.2.0.4.0</li>
              </div>
            </div>
            <div class="m3 ml">
              <div class="mtext">
                <li><span>当前SCN:</span>14736174057981</li>
                <li>
                  <span>线程:</span>1
                  <p class="sp2"><span>序列:</span> 167209</p>
                </li>

                <li><span>状态:</span>读写</li>
                <li style=""><span>生产库快照状态:</span>未启动</li>
                <li style="display: none;"><span>最早快照时间:</span></li>
                <li><span>快照空间使用率:</span>0%</li>
              </div>
            </div>
            <div class="m3 mr">
              <div class="mtext">
                <li><span>当前SCN:</span>14736174057978</li>

                <li>
                  <span>恢复速度:</span>169 KB/sec
                  <p></p>
                </li>
                <li>
                  <span>当前恢复:</span>线程: 1
                  <p class="sp2">序列: 167209</p>
                </li>

                <li><span>状态:</span>实时查询</li>
                <li style="display: none;">
                  <span>容灾库快照状态:</span>未启动
                </li>
                <li style=""><span>最早快照时间:</span>2020-04-22 10:58:47</li>
                <li><span>快照空间使用率:</span>49.95%</li>
              </div>
            </div>
          </div>
        </div>
        <div class="b5">
          <h3>备库信息</h3>
          <div class="m">
            <div class="e1">
              <h4>日志量 (单位:M)</h4>
              <div id="main" style="width: 100%; height: 100%; "></div>
            </div>
            <div class="e2">
              <h4>指标雷达</h4>
              <div id="main2" style="width: 100%; height: 100%; "></div>
            </div>
            <div class="e3">
              <table border="0">
                <tbody>
                  <tr>
                    <td>CPU空闲率</td>
                    <td>99%</td>
                    <td>内存空闲率</td>
                    <td>2%</td>
                  </tr>
                  <tr>
                    <td>Swap空闲率</td>
                    <td>99%</td>
                    <td>磁盘空闲率</td>
                    <td>62%</td>
                  </tr>
                  <tr>
                    <td>Inode空闲率</td>
                    <td></td>
                    <td>进程数</td>
                    <td>277</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
        <div class="b6">
          <h3>备份差异</h3>
          <div class="m">
            <li>
              <span>数据文件写入延时:</span
              ><b class="orange">0天0小时0分4秒 </b>
            </li>
            <li>
              <span>日志传输延时线程:</span
              ><b class="orange"
                >1<span style="display:inline-block;width:3.5em;"></span
              ></b>
            </li>
            <li>
              <span>日志传输延时序列差异:</span
              ><b class="orange"
                >0<span style="display:inline-block;width: 3.5em;"></span
              ></b>
            </li>
            <li>
              <span>日志应用延时线程:</span
              ><b class="orange"
                >1<span style="display:inline-block;width: 3.5em;"></span
              ></b>
            </li>
            <li>
              <span>日志应用延时序列差异:</span
              ><b class="orange"
                >1<span style="display:inline-block;width: 3.5em;"></span
              ></b>
            </li>
          </div>
        </div>
      </div>
    </div>

    <script type="text/javascript">
      var myChart = echarts.init(document.getElementById('main'), 'chalk')
      var option = {
        tooltip: {
          trigger: 'axis'
        },
        grid: {
          x: 25,
          y: 45,
          x2: 2,
          y2: 17,
          borderWidth: 1
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          axisTick: {
            alignWithLabel: true
          },
          axisLine: {
            onZero: false,
            lineStyle: {}
          },
          axisPointer: {
            label: {
              formatter: function(params) {
                return params.value + ' 日志量'
              }
            }
          },
          data: [
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            '',
            ''
          ]
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            symbol: 'circle', //设定为实心点
            symbolSize: 1, //设定实心点的大小
            data: [
              '488',
              '0',
              '0',
              '493',
              '0',
              '490',
              '0',
              '0',
              '0',
              '0',
              '0',
              '508',
              '490',
              '1146',
              '0',
              '0',
              '524',
              '490',
              '0',
              '0',
              '490',
              '0',
              '490',
              '0'
            ],
            type: 'line',
            areaStyle: {}
          }
        ]
      }
      myChart.setOption(option)

      var myChart2 = echarts.init(document.getElementById('main2'), 'chalk')
      var option = {
        tooltip: {},

        radar: {
          name: {
            textStyle: {
              color: '#fff',
              fontSize: 10
            }
          },

          center: ['45%', '50%'],
          radius: 40,
          nameGap: 0,
          indicator: [
            { name: 'CPU', max: 100 },
            { name: 'Swap', max: 100 },
            { name: '内存', max: 100 },
            { name: '磁盘', max: 100 },
            { name: 'Inode', max: 100 },
            { name: '进程', max: 100 }
          ]
        },
        series: [
          {
            name: '容灾库主机性能指标',
            type: 'radar',
            itemStyle: { normal: { areaStyle: { type: 'default' } } },
            symbol: 'circle', //设定为实心点
            symbolSize: 1, //设定实心点的大小
            data: [
              {
                value: [99, 99, 2, 62, 80, 55]
              }
            ]
          }
        ]
      }
      myChart2.setOption(option)
      myChart2.setOption(option)
      window.addEventListener('resize', function() {
        myChart.resize()
        myChart2.resize()
      })
    </script>
    <script type="text/javascript">
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
      function refresh() {
        window.location.reload()
      }
      setTimeout('refresh()', 60000) //指定60秒刷新一次
    </script>
  </body>
</html>
