package logger

import (
	"runtime/debug"
)

func Trace(args ...interface{}) {
	debug.PrintStack()
	logger.Trace(args)
}

func Debug(args ...interface{}) {
	debug.PrintStack()
	logger.Debug(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Warning(args ...interface{}) {
	logger.Warning(args)
}
