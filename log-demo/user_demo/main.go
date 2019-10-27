package main

import "log-demo/mylogger"

var logger mylogger.Logger

func main()  {
	logger = mylogger.NewFileLogger("debug","./","server.log")  //网日志文件内写
	//logger = mylogger.NewConsoleLogger("debug")  // 终端输出
	defer logger.Close()
	//for {
	//	logger.Debug("这是一条日志!!!")
	//	logger.Info("这是一条日志!!! 值: %d",1000)
	//	logger.Error("这是一条日志!!! 值: %d",1000)
	//	logger.Fatal("这是一条日志!!! 值: %d",1000)
	//}
	logger.Debug("这是一条日志!!!")
	logger.Info("这是一条日志!!! 值: %d",1000)
	logger.Warn("这是一条日志!!! 值: %d",1000)
	logger.Error("这是一条日志!!! 值: %d",1000)
	logger.Fatal("这是一条日志!!! 值: %d",1000)

}
