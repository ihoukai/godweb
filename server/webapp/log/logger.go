package log

import (
	"fmt"
	"github.com/go-ozzo/ozzo-log"
	"time"
)

// NewDefaultLogger 创建默认的日志
func newDefaultLogger() *log.Logger {
	logger := log.NewLogger()
	logger.CallStackDepth = 10
	t1 := log.NewConsoleTarget()
	t1.MaxLevel = log.LevelDebug
	logger.Targets = append(logger.Targets, t1)
	return logger
}

// NewLogger 创建日志
func newLogger(debug bool, ProcessID string, Logdir string) *log.Logger {
	logger := log.NewLogger()
	logger.CallStackDepth = 10
	//pid.nohup.log
	//pid.access.log
	//pid.error.log
	if debug {
		t1 := log.NewConsoleTarget()
		logger.Targets = append(logger.Targets, t1)
	} else {
		t2 := log.NewFileTarget()
		t2.FileName = fmt.Sprintf("%s/%s.error.log", Logdir, ProcessID)
		t2.MaxLevel = log.LevelWarning
		t3 := log.NewFileTarget()
		t3.FileName = fmt.Sprintf("%s/%s.access.log", Logdir, ProcessID)
		t3.MaxLevel = log.LevelDebug
		logger.Targets = append(logger.Targets, t2, t3)
	}
	logger = getLogger(logger, "godnet")
	return logger
}

func getLogger(logger *log.Logger, category string) *log.Logger {
	return logger.GetLogger(category, func(l *log.Logger, e *log.Entry) string {
		if e.Level <= log.LevelWarning {
			return fmt.Sprintf("%v [%v] %v %v", e.Time.Format(time.RFC3339), e.Level, e.Message, e.CallStack)
		}
		return fmt.Sprintf("%v [%v][%v] %v", e.Time.Format(time.RFC3339), e.Level, e.Category, e.Message)
	})
}
