package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

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

func QueryMQInfo(url string) (data DTSItems, err error) {
	body, err := GetRequestUrl(url)
	var reqDtsInfo DTSItems
	err = json.Unmarshal(body, &reqDtsInfo)
	if err != nil {
		return
	}
	return reqDtsInfo, nil
}
