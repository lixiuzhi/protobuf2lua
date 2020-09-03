package logger

import (
	"fmt"
)

const(
	ERROR = iota
	WARNING
	INFO
	DEBUG
)

var level = INFO

func SetLevel(levelname string){
	if levelname == "DEBUG"{
		level = DEBUG
	}else if levelname == "INFO"{
		level = INFO
	}else if levelname == "WARNING"{
		level = WARNING
	}else if levelname == "ERROR"{
		level = ERROR
	}
}

func Debug(a ...interface{}){
	if level>=DEBUG{
		msg:=fmt.Sprint(a...)
		fmt.Println("[DEBUG]",msg)
	}
}

func Debugf(format string, a ...interface{}){
	if level>=DEBUG{
		msg:=fmt.Sprintf(format,a...)
		fmt.Println("[DEBUG]",msg)
	}
}

func Info(a ...interface{}){
	if level>=INFO{
		msg:=fmt.Sprint(a...)
		fmt.Println("[INFO]",msg)
	}
}

func Infof(format string, a ...interface{}){
	if level>=INFO{
		msg:=fmt.Sprintf(format,a...)
		fmt.Println("[INFO]",msg)
	}
}

func Warn(a ...interface{}){
	if level>=WARNING{
		msg:=fmt.Sprint(a...)
		fmt.Println("[WARN]",msg)
	}
}

func Warnf(format string, a ...interface{}){
	if level>=WARNING{
		msg:=fmt.Sprintf(format,a...)
		fmt.Println("[WARN]",msg)
	}
}

func Error(a ...interface{}){
	if level>=ERROR{
		msg:=fmt.Sprint(a...)
		fmt.Println("[ERROR]",msg)
	}
}

func Errorf(format string, a ...interface{}){
	if level>=ERROR{
		msg:=fmt.Sprintf(format,a...)
		fmt.Println("[ERROR]",msg)
	}
}