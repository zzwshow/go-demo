package main

import (
	"exporter/collector"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func init() {
	// 注册自身采集器
	if err := collector.InitConf();err != nil{
		log.Fatal(err)
	}
	prometheus.MustRegister(collector.NewNodeCollector())
}

func main() {
	http.Handle(collector.Conf.Uri, promhttp.Handler())
	fmt.Println("success")
	if err := http.ListenAndServe(collector.Conf.Host, nil); err != nil {
		fmt.Printf("Error occur when start server %v", err)
		os.Exit(1)
	}
}



