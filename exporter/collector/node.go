package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type NodeCollector struct {
	consumptionGroupDelayDesc *prometheus.Desc //Gauge
	consumptionGroupTPSDesc   *prometheus.Desc //Gauge
	consumptionNumberDesc     *prometheus.Desc //Gauge
	mutex                     sync.Mutex
}

func NewNodeCollector() prometheus.Collector {
	return &NodeCollector{
		consumptionGroupDelayDesc: prometheus.NewDesc(
			"rocketMQ_consumption_group_delay_diffTotal",
			"This value indicates the number of consumption stacks",
			[]string{"group"}, nil),
		consumptionGroupTPSDesc: prometheus.NewDesc(
			"rocketMQ_consumption_group_per_second_TPS",
			"This value indicates the number of Consumption per second",
			[]string{"group"}, nil),
		consumptionNumberDesc: prometheus.NewDesc(
			"rocketMQ_number_of_consumers",
			"This value identifies the number of consumers", []string{"group"}, nil),
	}
}

func (n *NodeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.consumptionGroupDelayDesc
	ch <- n.consumptionGroupTPSDesc
	ch <- n.consumptionNumberDesc
}

func (n *NodeCollector) Collect(ch chan<- prometheus.Metric) {
	n.mutex.Lock()
	data, err := QueryMQInfo(Conf.MqUrl)
	if err != nil {
		fmt.Printf("request DTS mq console failed. url: %v error:%v", Conf.MqUrl,err)
	}
	if len(data.Data) == 0 {
		return
	}
	for _, i := range data.Data {
		ch <- prometheus.MustNewConstMetric(n.consumptionGroupDelayDesc, prometheus.GaugeValue, float64(i.DiffTotal), i.Group)
		ch <- prometheus.MustNewConstMetric(n.consumptionGroupTPSDesc, prometheus.GaugeValue, float64(i.ConsumeTps), i.Group)
		ch <- prometheus.MustNewConstMetric(n.consumptionNumberDesc, prometheus.GaugeValue, float64(i.Count), i.Group)
	}
	n.mutex.Unlock()
}
