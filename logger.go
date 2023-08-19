package logger

import (
	"fmt"
	"os"
)

const (
	FATAL = 0
	ERROR = 1
	WARN  = 2
	INFO  = 3
	DEBUG = 4
)

var Level int = INFO
var pid int = os.Getpid()
var processName string = getProcessName("UPS")
var Identifier string = fmt.Sprintf("%d|%s", pid, processName)
var LogLevelName []string = []string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG"}

var LogMethod []func(level int, format string, v ...interface{}) = []func(level int, format string, v ...interface{}){
	LogStdout,
	LogStdout,
	LogStdout,
	LogStdout,
	LogStdout,
}

type Logger struct {
	Level int
}

func Log(level int, format string, v ...interface{}) {
	if level <= Level {
		LogMethod[level](level, format, v...)
	}
}

func Fatal(format string, v ...interface{}) {
	Log(FATAL, format, v...)
}

func Error(format string, v ...interface{}) {
	Log(ERROR, format, v...)
}

func Warn(format string, v ...interface{}) {
	Log(WARN, format, v...)
}

func Info(format string, v ...interface{}) {
	Log(INFO, format, v...)
}

func Debug(format string, v ...interface{}) {
	Log(DEBUG, format, v...)
}
