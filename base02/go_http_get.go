package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//"rocketmq_check_url":
//"eds_url": "http://10.200.2.55:8099/consumer/groupList.query"
//"dts_url": "http://10.200.2.55:8180/consumer/groupList.query"
//"canal_url": "http://172.16.0.7:11000/offset"

const (
	edsUrl = "https://mq.yimidida.com/consumer/groupList.query"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type DataSlice struct {
	Group             string `json:"group"`
	Version           string `json:"version"`
	Count             int    `json:"count"`
	CONSUME_PASSIVELY string `json:"CONSUME_PASSIVELY"`
	CLUSTERING        string `json:"CLUSTERING"`
	ConsumeTps        int    `json:"consumeTps"`
	DiffTotal         int    `json:"diffTotal"`
}

type Items struct {
	Status int         `json:"status"`
	Data   []DataSlice `json:"data"`
}

func GetRequestUrl(url string) (result string, err error) {
	resp, err := myClient.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var req Items
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Printf("json unmarshal failed error:", err.Error())
		return
	}
	fmt.Println("====", req)
	return string(body), nil
}

func main() {
	jsonStr, err := GetRequestUrl(edsUrl)
	if err != nil {
		os.Exit(0)
	}
	fmt.Printf("------>", jsonStr)
}
