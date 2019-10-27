package mylogger

import (
	"fmt"
	"os"
	"time"
)

// 网终端内打印日志
type ConsoleLogger struct {
	level Level
}

// 构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger{
	logLevel := parseLogLevel(levelStr)
	cl := &ConsoleLogger{
		level:logLevel,
	}
	return cl
}

// 封装公共方法
func (c *ConsoleLogger) log(level Level,format string, args ...interface{}) {
	// 日志格式 : 时间 文件:行号 函数名 日志级别 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName,line,funName := getCallerInfo(3)
	msg := fmt.Sprintf(format,args...) // 处理用户要记录的日志
	levelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s] [%s:%d] [%s] [%s] %s",nowStr,fileName,line,funName,levelStr,msg)
	fmt.Fprintln(os.Stdout,logMsg)
}

// Debug方法
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel,format, args...)
}

// INFO方法
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel,format, args...)
}

// WARN方法
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel,format, args...)
}

// ERROR方法
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel,format, args...)
}

// FATAL方法
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel,format, args...)
}

// 终端输出不需要关闭,但是接口方法需要实现
func (c *ConsoleLogger) Close(){
}


