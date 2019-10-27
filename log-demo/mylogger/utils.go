package mylogger

import (
	"path"
	"runtime"
)

// 公用工具函数

//获取行号 (skip 跳过多少行)
func getCallerInfo(skip int) (fileName string,line int,funcName string){
	pc,file,line,ok :=runtime.Caller(skip)
	if !ok{
		return
	}
	// file 是个全路径,需要剥离出文件名
	fileName = path.Base(file)
	// 根据pc 获取函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return
}
