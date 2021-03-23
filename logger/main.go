package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//定义LogLevel类型
type LogLevel uint16

//定义日志级对应的序号，类型是LogLevel
const (
	UNKNOW LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FAULT
)

//定义Logger 结构体
type Logger struct {
	level LogLevel
}

//根据字符串获取到对应的日志级别序号
func pasreLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "falt":
		return FAULT, nil
	default:
		err := errors.New("WARNING:无效的日志级别")
		return UNKNOW, err
	}
}

//获取日志级别，并返回Logger类型
func NewLog(s string) Logger {
	loglevel, err := pasreLogLevel(s)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	return Logger{
		level: loglevel,
	}
}

//根据日志级别序号打印日志
func logHandle(msg string, level LogLevel, n string, l Logger) {
	if l.level <= level {
		now := time.Now()
		fmt.Printf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), n, msg)
	}
}

//BEBUG级别日志函数，传入日志字符串，调用打印日志函数logHandle
func (l Logger) Debug(msg string) {
	logHandle(msg, DEBUG, "DEUG", l)
}

//INFO级别日志函数，传入日志字符串，调用打印日志函数logHandle
func (l Logger) Info(msg string) {
	logHandle(msg, INFO, "INFO", l)

}

//WARNING级别日志函数，传入日志字符串，调用打印日志函数logHandle
func (l Logger) Warning(msg string) {
	logHandle(msg, WARNING, "WARNING", l)
}

//ERROR级别日志函数，传入日志字符串，调用打印日志函数logHandle
func (l Logger) Error(msg string) {
	logHandle(msg, ERROR, "ERROR", l)
}

//FAULT级别日志函数，传入日志字符串，调用打印日志函数logHandle
func (l Logger) Fault(msg string) {
	logHandle(msg, FAULT, "FALT", l)
}

func main() {
	//日志级别
	// DEBUG
	// INFO
	// WARNING
	// ERROR
	// FALT

	log := NewLog("infos")
	for {
		log.Debug("this is test log DEBUG ........")
		log.Info("this is test log INFO .......")
		log.Warning("this is test log WARNING .....")
		log.Error("this is test log ERROR .....")
		log.Fault("this is test log FALT ......")
		time.Sleep(time.Second * 2)
	}
}
