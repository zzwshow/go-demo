package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
	"sync"
)

//var hostname string

type NodeCollector struct {
	goroutinesDesc *prometheus.Desc //Gauge
	threadsDesc    *prometheus.Desc //Gauge
	mutex          sync.Mutex
}

//初始化采集器
func NewNodeCollector() prometheus.Collector {
	//host, _ := host.Info()
	//hostname = host.Hostname
	return &NodeCollector{
		goroutinesDesc: prometheus.NewDesc(
			"goroutines_num",
			"协程数.",
			nil, nil),
		threadsDesc: prometheus.NewDesc(
			"threads_num",
			"线程数",
			nil, nil),
	}
}

// Describe returns all descriptions of the collector.
//实现采集器Describe接口
func (n *NodeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.goroutinesDesc
	ch <- n.threadsDesc
}

// Collect returns the current state of all metrics of the collector.
//实现采集器Collect接口,真正采集动作
func (n *NodeCollector) Collect(ch chan<- prometheus.Metric) {
	n.mutex.Lock()
	ch <- prometheus.MustNewConstMetric(n.goroutinesDesc, prometheus.GaugeValue, float64(runtime.NumGoroutine()))

	num, _ := runtime.ThreadCreateProfile(nil)
	ch <- prometheus.MustNewConstMetric(n.threadsDesc, prometheus.GaugeValue, float64(num))
	n.mutex.Unlock()
}
