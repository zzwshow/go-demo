package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	mqConsumptionGroupDelay = prometheus.GaugeOpts{prometheus.GaugeOpts{
		Name: "RocketMQ Consumption Group Delay diffTotal",
		Help: "This value indicates the number of consumption stacks",
	}}
	mqConsumptionGroupTPS = prometheus.GaugeOpts{prometheus.GaugeOpts{
		Name: "RocketMQ Consumption group per second TPS",
		Help: "This value indicates the number of Consumption per second",
	}}
	mqConsumptionNumber = prometheus.GaugeOpts{prometheus.GaugeOpts{
		Name: "RocketMQ Number of consumers",
		Help: "This value identifies the number of consumers",
	}}
)

// ===

type DTSDataSlice struct {
	Group        string `json:"group"`
	Version      string `json:"version"`
	Count        int    `json:"count"`
	ConsumeType  string `json:"consumeType"`
	MessageModel string `json:"messageModel"`
	ConsumeTps   int    `json:"consumeTps"`
	DiffTotal    int    `json:"diffTotal"`
}

type DTSItems struct {
	Status int            `json:"status"`
	Data   []DTSDataSlice `json:"data"`
}

var myClient = &http.Client{Timeout: 30 * time.Second}

func GetRequestUrl(url string) (body []byte, err error) {
	resp, err := myClient.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return
}

func queryMQInfo(url string) (delayTotal, tps, number int, err error) {
	body, err := GetRequestUrl(url)
	var reqDtsInfo DTSItems
	err = json.Unmarshal(body, &reqDtsInfo)
	if err != nil {
		content := fmt.Sprintf(" request DTS mq console failed. url: %v", url)
		fmt.Println(content)
		return
	}

}

func main() {

}
