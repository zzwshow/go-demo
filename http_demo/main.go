package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//"rocketmq_check_url":
//"eds_url": "http://10.200.2.55:8099/consumer/groupList.query"
//"dts_url": "http://10.200.2.55:8180/consumer/groupList.query"
//"canal_url": "http://172.16.0.7:11000/offset"

const canalUrl = "http://canalmq.ymdd.tech/"

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return
}

type Result struct {
	data interface{}
}

func main() {
	result := new(Result)
	err := getJson(canalUrl, result)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	fmt.Println(result)

}
