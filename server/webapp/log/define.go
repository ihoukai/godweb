package log

// ILog 日志模块
type ILog interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
	Warning(format string, a ...interface{})
	GetLogger(category string) ILog
}
