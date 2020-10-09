// 自调用函数
;(function () {
  // 封装函数
  var setFont = function () {
    // 获取html元素
    var html = document.documentElement
    // var html = document.querySelector('html');
    // 获取宽度
    var width = html.clientWidth
    // 如果小于1024，那么就按1024
    if (width < 1024) {
      width = 1024
    }
    // 如果大于1920，那么就按1920
    if (width > 1920) {
      width = 1920
    }
    // 计算
    var fontSize = width / 80 + 'px'
    // 设置给html
    html.style.fontSize = fontSize
  }
  setFont()
  // onresize：改变大小事件
  window.onresize = function () {
    setFont()
  }
})()

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
      data: ['直接访问', '邮件营销', '联盟广告', '视频广告', '搜索引擎'],
    },
    series: [
      {
        name: '访问来源',
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
        data: [
          { value: 335, name: '直接访问' },
          { value: 310, name: '邮件营销' },
          { value: 234, name: '联盟广告' },
          { value: 135, name: '视频广告' },
          { value: 1548, name: '搜索引擎' },
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
// 折线
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
      data: ['邮件营销', '联盟广告', '视频广告', '直接访问', '搜索引擎'],
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
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {
        name: '邮件营销',
        type: 'line',
        stack: '总量',
        data: [120, 132, 101, 134, 90, 230, 210],
      },
      {
        name: '联盟广告',
        type: 'line',
        stack: '总量',
        data: [220, 182, 191, 234, 290, 330, 310],
      },
      {
        name: '视频广告',
        type: 'line',
        stack: '总量',
        data: [150, 232, 201, 154, 190, 330, 410],
      },
      {
        name: '直接访问',
        type: 'line',
        stack: '总量',
        data: [320, 332, 301, 334, 390, 330, 320],
      },
    ],
  }
  myChart.setOption(option)
})()

// ct1
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
    // legend: {
    //   data: ['邮件营销', '联盟广告', '视频广告', '直接访问', '搜索引擎'],
    // },
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
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {
        name: '邮件营销',
        type: 'line',
        stack: '总量',
        data: [120, 132, 101, 134, 90, 230, 210],
      },
      {
        name: '联盟广告',
        type: 'line',
        stack: '总量',
        data: [220, 182, 191, 234, 290, 330, 310],
      },
      {
        name: '视频广告',
        type: 'line',
        stack: '总量',
        data: [150, 232, 201, 154, 190, 330, 410],
      },
      {
        name: '直接访问',
        type: 'line',
        stack: '总量',
        data: [320, 332, 301, 334, 390, 330, 320],
      },
    ],
  }
  myChart.setOption(option)
})()

// ct2
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
    // legend: {
    //   data: ['邮件营销', '联盟广告', '视频广告', '直接访问', '搜索引擎'],
    // },
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
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {
        name: '邮件营销',
        type: 'line',
        stack: '总量',
        data: [120, 132, 101, 134, 90, 230, 210],
      },
      {
        name: '联盟广告',
        type: 'line',
        stack: '总量',
        data: [220, 182, 191, 234, 290, 330, 310],
      },
      {
        name: '视频广告',
        type: 'line',
        stack: '总量',
        data: [150, 232, 201, 154, 190, 330, 410],
      },
      {
        name: '直接访问',
        type: 'line',
        stack: '总量',
        data: [320, 332, 301, 334, 390, 330, 320],
      },
    ],
  }
  myChart.setOption(option)
})()

// ct3
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
    // legend: {
    //   data: ['邮件营销', '联盟广告', '视频广告', '直接访问', '搜索引擎'],
    // },
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
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {
        name: '邮件营销',
        type: 'line',
        stack: '总量',
        data: [120, 132, 101, 134, 90, 230, 210],
      },
      {
        name: '联盟广告',
        type: 'line',
        stack: '总量',
        data: [220, 182, 191, 234, 290, 330, 310],
      },
      {
        name: '视频广告',
        type: 'line',
        stack: '总量',
        data: [150, 232, 201, 154, 190, 330, 410],
      },
      {
        name: '直接访问',
        type: 'line',
        stack: '总量',
        data: [320, 332, 301, 334, 390, 330, 320],
      },
    ],
  }
  myChart.setOption(option)
})()

// ct4
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
    // legend: {
    //   data: ['邮件营销', '联盟广告', '视频广告', '直接访问', '搜索引擎'],
    // },
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
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {
        name: '邮件营销',
        type: 'line',
        stack: '总量',
        data: [120, 132, 101, 134, 90, 230, 210],
      },
      {
        name: '联盟广告',
        type: 'line',
        stack: '总量',
        data: [220, 182, 191, 234, 290, 330, 310],
      },
      {
        name: '视频广告',
        type: 'line',
        stack: '总量',
        data: [150, 232, 201, 154, 190, 330, 410],
      },
      {
        name: '直接访问',
        type: 'line',
        stack: '总量',
        data: [320, 332, 301, 334, 390, 330, 320],
      },
    ],
  }
  myChart.setOption(option)
})()

// ct5
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
    // legend: {
    //   data: ['邮件营销', '联盟广告', '视频广告', '直接访问', '搜索引擎'],
      
    // },
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
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
    },
    yAxis: {
      type: 'value',
    },
    series: [
        {
            name: '邮件营销',
            type: 'line',
            stack: '总量',
            data: [120, 132, 101, 134, 90, 230, 210],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                ]
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgba(244,62,173,0.4)',
                },
                {
                  offset: 1,
                  color: 'rgba(255,255,255,0)',
                },
              ]),
            },
          },
          {
            name: '联盟广告',
            type: 'line',
            stack: '总量',
            data: [220, 182, 191, 234, 290, 330, 310],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                ]
            },
            areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                  {
                    offset: 0,
                    color: 'rgba(35, 150, 255,0.4)',
                  },
                  {
                    offset: 1,
                    color: 'rgba(255,255,255,0)',
                  },
                ]),
              },
          },
    ],
  }
  myChart.setOption(option)
})()

// ct6
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

// ct7
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
          data: ['邮件营销', '联盟广告', '视频广告', '直接访问', '搜索引擎'],
          
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
          data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
        },
        yAxis: {
          type: 'value',
        },
        series: [
          {
            name: '邮件营销',
            type: 'line',
            stack: '总量',
            data: [120, 132, 101, 134, 90, 230, 210],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                ]
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgba(244,62,173,0.4)',
                },
                {
                  offset: 1,
                  color: 'rgba(255,255,255,0)',
                },
              ]),
            },
          },
          {
            name: '联盟广告',
            type: 'line',
            stack: '总量',
            data: [220, 182, 191, 234, 290, 330, 310],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                ]
            },
            areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                  {
                    offset: 0,
                    color: 'rgba(35, 150, 255,0.4)',
                  },
                  {
                    offset: 1,
                    color: 'rgba(255,255,255,0)',
                  },
                ]),
              },
          },
        ],
      }
    myChart.setOption(option)
  })()
// 告警信息滚动
;(function () {
  // 滚动复制一份
  $('.monitor .marquee').each(function () {
    // 拿到了marquee里面的所有row
    var rows = $(this).children().clone()
    // 追加进去
    $(this).append(rows)
    // 父.append(子)==>子.appendTo(父)
    // $('ul').append($('<li>li</li>'));==>$('<li>li</li>').appendTo('ul');
  })
})();

//圆环
window.onload = function () {
  new Progress().renderOne('canvas0', 100, 7, 60)
  new Progress().renderOne('canvas1', 100, 7, 50)
  new Progress2().renderOne('canvas2', 100, 7, 80)
  new Progress().renderOne('canvas3', 100, 7, 75)
};


 //获取时间并设置格式

 $(function () {
    setInterval("GetTime()", 1000);
  });
 function GetTime() {
    var mon, day, now, hour, min, ampm, time, str, tz, end, beg, sec;
    /*
    mon = new Array("Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug",
            "Sep", "Oct", "Nov", "Dec");
    */
    mon = new Array("一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月",
      "九月", "十月", "十一月", "十二月");
    /*
    day = new Array("Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat");
    */
    day = new Array("周日", "周一", "周二", "周三", "周四", "周五", "周六");
    now = new Date();
    hour = now.getHours();
    min = now.getMinutes();
    sec = now.getSeconds();
    if (hour < 10) {
      hour = "0" + hour;
    }
    if (min < 10) {
      min = "0" + min;
    }
    if (sec < 10) {
      sec = "0" + sec;
    }
    $("#Timer").html(
      now.getFullYear() + "年" + (now.getMonth() + 1) + "月" + now.getDate() + "日" + "  " + hour + ":" + min + ":" + sec
    );
    //$("#Timer").html(
    //        day[now.getDay()] + ", " + mon[now.getMonth()] + " "
    //                + now.getDate() + ", " + now.getFullYear() + " " + hour
    //                + ":" + min + ":" + sec);
  }