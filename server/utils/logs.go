package utils

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger

//运行方式
var runmode string

var level string = "7"

func init() {
	consoleLogs = logs.NewLogger(1)
	consoleLogs.EnableFuncCallDepth(true)
	consoleLogs.SetLogFuncCallDepth(4)
	consoleLogs.SetLogger(logs.AdapterConsole)
	consoleLogs.Async() //异步

	fileLogs = logs.NewLogger(10000)
	fileLogs.EnableFuncCallDepth(true)
	fileLogs.SetLogFuncCallDepth(4)

	fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/main.log",
		"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
		"level":`+level+`,
		"daily":true,
		"maxdays":10}`)
	fileLogs.Async() //异步
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}
func LogEmergency(v interface{}) {
	log("emergency", "%s", v)
}
func LogAlert(v interface{}) {
	log("alert", "%s", v)
}
func LogCritical(v interface{}) {
	log("critical", "%s", v)
}
func LogError(v interface{}) {
	log("error", "%s", v)
}
func LogWarning(v interface{}) {
	log("warning", "%s", v)
}
func LogNotice(v interface{}) {
	log("notice", "%s", v)
}
func LogInfo(v interface{}) {
	log("info", "%s", v)
}
func LogDebug(v interface{}) {
	log("debug", "%s", v)
}

func LogEmergencyf(format string, v ...interface{}) {
	log("emergency", format, v...)
}
func LogAlertf(format string, v ...interface{}) {
	log("alert", format, v...)
}
func LogCriticalf(format string, v ...interface{}) {
	log("critical", format, v...)
}
func LogErrorf(format string, v ...interface{}) {
	log("error", format, v...)
}
func LogWarningf(format string, v ...interface{}) {
	log("warning", format, v...)
}
func LogNoticef(format string, v ...interface{}) {
	log("notice", format, v...)
}
func LogInfof(format string, v ...interface{}) {
	log("info", format, v...)
}
func LogDebugf(format string, v ...interface{}) {
	log("debug", format, v...)
}

//Log 输出日志
func log(level, format string, v ...interface{}) {
	//format := "%s"
	output := fmt.Sprintf(format, v...)
	if level == "" {
		level = "debug"
	}
	if runmode == "dev" || runmode == "pro" {
		switch level {
		case "emergency":
			fileLogs.Emergency(output)
		case "alert":
			fileLogs.Alert(output)
		case "critical":
			fileLogs.Critical(output)
		case "error":
			fileLogs.Error(output)
		case "warning":
			fileLogs.Warning(output)
		case "notice":
			fileLogs.Notice(output)
		case "info":
			fileLogs.Info(output)
		case "debug":
			fileLogs.Debug(output)
		default:
			fileLogs.Debug(output)
		}
	}
	switch level {
	case "emergency":
		consoleLogs.Emergency(output)
	case "alert":
		consoleLogs.Alert(output)
	case "critical":
		consoleLogs.Critical(output)
	case "error":
		consoleLogs.Error(output)
	case "warning":
		consoleLogs.Warning(output)
	case "notice":
		consoleLogs.Notice(output)
	case "info":
		consoleLogs.Info(output)
	case "debug":
		consoleLogs.Debug(output)
	default:
		consoleLogs.Debug(output)
	}
}
