package main

import (
	"exporter/collector"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

func init() {
	// 注册自身采集器
	prometheus.MustRegister(collector.NewNodeCollector())
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":9991", nil); err != nil {
		fmt.Printf("Error occur when start server %v", err)
		os.Exit(1)
	}
}
