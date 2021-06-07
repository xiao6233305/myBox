package common

// 日志相关的 都放这个文件里面了  跟日志解耦
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func SetLogLevel(mode string) {
	_ = beego.BeeLogger.SetLogger(`console`, ``)
	//设置日志模式
	switch mode {
	case `dev`:
		beego.SetLevel(logs.LevelDebug)
	case `PRODUCT`:
		beego.SetLevel(logs.LevelCritical)
	default:
		beego.SetLevel(logs.LevelInformational)
	}
}

func Debug(v ...interface{}) {
	beego.Debug(v...)
}

func Error(v ...interface{}) {
	beego.Error(v...)
}

func Alert(v ...interface{}) {
	beego.Alert(v...)
}

func Info(v ...interface{}) {
	beego.Info(v...)
}
func Emergency(v ...interface{}) {
	beego.Emergency(v...)
}
func Critical(v ...interface{}) {
	beego.Critical(v...)
}
func Warning(v ...interface{}) {
	beego.Warning(v...)
}
func Warn(v ...interface{}) {
	beego.Warn(v...)
}
func Notice(v ...interface{}) {
	beego.Notice(v...)
}
func Informational(v ...interface{}) {
	beego.Informational(v...)
}

func Trace(v ...interface{}) {
	beego.Trace(v...)
}
