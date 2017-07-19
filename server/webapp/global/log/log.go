package log

import (
	"github.com/go-ozzo/ozzo-log"
)

var (
	glog ILog
)

type godnetLog struct {
	log *log.Logger
}

// Debug
func (g godnetLog) Debug(format string, a ...interface{}) {
	g.log.Debug(format, a...)
}

// Info
func (g godnetLog) Info(format string, a ...interface{}) {
	g.log.Info(format, a...)
}

// Error
func (g godnetLog) Error(format string, a ...interface{}) {
	g.log.Error(format, a...)
}

// Warning
func (g godnetLog) Warning(format string, a ...interface{}) {
	g.log.Warning(format, a...)
}

// GetLogger 获得日志器
func (g godnetLog) GetLogger(category string) ILog {
	return newGodnetLog(getLogger(g.log, category))
}

func newGodnetLog(log *log.Logger) ILog {
	glog := &godnetLog{}
	glog.log = log
	glog.log.Open()
	return glog
}

// init 系统调用
func init() {
	glog = newGodnetLog(newDefaultLogger())
}

// Init 日志模块初始化
func Init(debug bool, ProcessID string, Logdir string) {
	glog = newGodnetLog(newLogger(debug, ProcessID, Logdir))
}

// SetCustomLog 设置用户自定义日志器
func SetCustomLog(log ILog) {
	glog = log
}

// GetLogger 获得日志器
func GetLogger() ILog {
	return glog
}

// Debug ..
func Debug(format string, a ...interface{}) {
	glog.Debug(format, a...)
}

// Info ..
func Info(format string, a ...interface{}) {
	glog.Info(format, a...)
}

// Error ..
func Error(format string, a ...interface{}) {
	glog.Error(format, a...)
}

// Warning ..
func Warning(format string, a ...interface{}) {
	glog.Warning(format, a...)
}
