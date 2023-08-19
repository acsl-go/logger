package logger

import (
	"fmt"
	"time"
)

func LogStdout(level int, format string, v ...interface{}) {
	//ts := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	ts := time.Now().Format("2016-01-02 15:04:05.999")
	fmt.Printf("[%s][%s][%d][%s] "+format+"\n", append([]interface{}{ts, Identifier, pid, LogLevelName[level]}, v...)...)
}
