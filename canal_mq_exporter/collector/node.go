package collector

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type NodeCollector struct {
	consumptionGroupDelayDesc *prometheus.Desc //Gauge
	mutex                     sync.Mutex
}

func NewNodeCollector() prometheus.Collector {
	return &NodeCollector{
		consumptionGroupDelayDesc: prometheus.NewDesc(
			"rocketMQ_consumption_group_delay_diffTotal",
			"This value indicates the number of consumption stacks",
			[]string{"group"}, nil),
	}
}

var myClient = &http.Client{Timeout: 30 * time.Second}

func getRequestUrl(url string) (body []byte, err error) {
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

func (n *NodeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.consumptionGroupDelayDesc
}

func (n *NodeCollector) Collect(ch chan<- prometheus.Metric) {
	n.mutex.Lock()
	body, err := getRequestUrl(Conf.MqUrl)
	if err != nil {
		fmt.Printf("request DTS mq console failed. url: %v error:%v", Conf.MqUrl, err)
	}
	var reqData map[string]int
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Fatal(err)
	}

	if len(reqData) == 0 {
		return
	}
	for k, v := range reqData {
		ch <- prometheus.MustNewConstMetric(n.consumptionGroupDelayDesc, prometheus.GaugeValue, float64(v), k)
	}
	n.mutex.Unlock()
}
