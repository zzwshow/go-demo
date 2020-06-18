package main

import (
	"fmt"
	"os"
	
	"wechat/wechat"
)

func main() {
	// 获取token
	access_token_str, err := wechat.GetAccessToken()
	if err != nil {
		fmt.Println("获取token failed")
		os.Exit(1)
	}
	fmt.Println("token:", access_token_str)
	
	//生成告警内容
	msg, err := wechat.GenerateGroupTextMsg("OpsAlarmTestGroup", "测试告警 忽略...")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		panic(err)
	}
	
	// 发布告警
	err = wechat.SendMsg(access_token_str, msg)
	if err != nil {
		fmt.Println("send failed")
		os.Exit(1)
	}
	fmt.Println("send ok")
	
	fmt.Println(msg)
}
