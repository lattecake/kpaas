package logger

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.StdlibWriter{})
	level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "caller", log.DefaultCaller)
}

func Info(keyvals ...interface{}) {
	_ = level.Info(logger).Log(keyvals...)
}

func Debug(keyvals ...interface{}) {
	_ = level.Debug(logger).Log(keyvals...)
}

func Error(keyvals ...interface{}) {
	_ = level.Error(logger).Log(keyvals...)
}

func Warn(keyvals ...interface{}) {
	_ = level.Warn(logger).Log(keyvals...)
}
