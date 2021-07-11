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



// 告警信息滚动
;(function () {
  // 滚动复制一份
  $('.monitor .marquee').each(function () {
    // 拿到了marquee里面的所有row

    var rows = $(this).children().clone()
    // 追加进去
    //$(this).append(rows)
    // 父.append(子)==>子.appendTo(父)
    // $('ul').append($('<li>li</li>'));==>$('<li>li</li>').appendTo('ul');
  })
})();



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