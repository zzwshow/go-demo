package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// --------------- 创建告警组 ------------------------
type AlarmGroupProfil struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
	ChatID   string   `json:"chatid"`
}

// 构造map string
func generateJsonObj(name, owner, chatid string, userlist []string) (alarmProfile string, err error) {
	alarmProfil := &AlarmGroupProfil{
		Owner:    owner,
		Name:     name,
		UserList: userlist,
		ChatID:   chatid,
	}
	data, err := json.Marshal(alarmProfil)
	return string(data), err
}

func main() {
	data, err := generateJsonObj("测试组", "ym050646", "xxxxxxxxxxx", []string{"ym050646", "ym050446"})
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println("data: ", data)
}
