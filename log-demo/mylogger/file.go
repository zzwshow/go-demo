package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志


// FileLogger 文件日志结构体
type FileLogger struct {
	level Level   //使用方需要打印的日志级别
	maxSize int64
	fileName string
	filePath string
	file *os.File
	errFile *os.File

}


// 构造函数
func NewFileLogger(levelStr string,fileName,filePath string) *FileLogger{
	logLevel := parseLogLevel(levelStr)
	f1 := &FileLogger{
		fileName: fileName,
		filePath: filePath,
		level:logLevel,
		maxSize: 10 * 1024 * 1024, // 设置日志文件大小10M
	}
	f1.initFile()  // 根据文件名和路径初始化文件句柄,并赋值给结构体,并返回
	return f1
}

// 创建文件句柄,赋值给结构体
func (f *FileLogger) initFile(){
	logName := path.Join(f.filePath,f.fileName)
	// 打开文件
	fileObj, err := os.OpenFile(logName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil{
		panic(fmt.Errorf("打开日志记录文件:%s 失败!!",logName))
	}
	f.file = fileObj
	// 打开错误日志记录的日志文件
	errFileName := fmt.Sprintf("error_%s",logName)
	errfileObj, err := os.OpenFile(errFileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil{
		panic(fmt.Errorf("打开日志记录文件:%s 失败!!",logName))
	}
	f.errFile = errfileObj
}

// 判断日志文件是否需要拆分
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo,_ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize // 当文件大于设定值时,返回true ,需要拆分了
}


// 切割日志的方法
func (f *FileLogger) splitLogFile(ff *os.File) *os.File {
	// 判断被写入的文件大小是否大于maxSize ,如果大于了设定的大小就切割一份
	nowStr := time.Now().Format("2006-01-02_15-04-05")
	fileName := ff.Name() //拿到文件完整路径
	backupName := fmt.Sprintf("%s_%s",nowStr,fileName)
	ff.Close()
	// 备份文件
	os.Rename(fileName,backupName)
	// 新建文件
	fileObj, err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil{
		panic(fmt.Errorf("打开日志记录文件:%s 失败!!",fileName))
	}
	return fileObj
}


// 封装写入日志的公共方法
func (f *FileLogger) log(level Level,format string, args ...interface{}) {
	if f.level > level{
		return
	}
	// 日志格式 : 时间 文件:行号 函数名 日志级别 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName,line,funName := getCallerInfo(3)
	msg := fmt.Sprintf(format,args...) // 处理用户要记录的日志
	levelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s] [%s:%d] [%s] [%s] %s",nowStr,fileName,line,funName,levelStr,msg)
	// 写之前做日志切割检查
	if f.checkSplit(f.file){
		f.file = f.splitLogFile(f.file)
	}
	fmt.Fprintln(f.file,logMsg)
	//将大于error 级别的日志在写一份到错误日志文件中
	if level >= ErrorLevel{
		if f.checkSplit(f.errFile){
			f.errFile = f.splitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile,logMsg)
	}
}

// Debug方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel,format, args...)
}

// INFO方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel,format, args...)
}

// WARN方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel,format, args...)
}

// ERROR方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel,format, args...)
}

// FATAL方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel,format, args...)
}

// 关闭文件句柄
func (f *FileLogger) Close(){
	f.file.Close()
	f.errFile.Close()
}