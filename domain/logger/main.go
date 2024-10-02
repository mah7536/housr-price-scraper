package logger

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

const (
	// 時間格式
	TIME_LAYOUT = "2006-01-02 15:04:05"

	// logger 顏色格式
	InfoColor    = "\033[1;34m%s %v [Info] : \033[0m%s"
	NoticeColor  = "\033[1;36m%s %v [Note] : \033[0m%s"
	WarningColor = "\033[1;33m%s %v [Warn] : \033[0m%s"
	ErrorColor   = "\033[1;31m%s %v [Error] :\033[0m%s"
	DebugColor   = "\033[0;36m%s %v [Debug] :\033[0m%s"
)

func Now() string {
	return time.Now().Format(TIME_LAYOUT)
}

func Info(info ...interface{}) {
	log(InfoColor, info...)
}

func Notice(info ...interface{}) {
	log(NoticeColor, info...)
}

func Warn(info ...interface{}) {
	log(WarningColor, info...)
}

func Error(info ...interface{}) {
	log(ErrorColor, info...)
}

func Debug(info ...interface{}) {
	log(DebugColor, info...)
}

func log(colorType string, info ...interface{}) {
	_, fromFile, fromLine, _ := runtime.Caller(2)
	fmt.Printf(colorType, Now(), fromFile+":"+strconv.Itoa(fromLine), fmt.Sprintln(info...))
}
