package esutil

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
)

var err error

var ESClient *elastic.Client
var servers = []string{"http://localhost:9200/"}
var ctx = context.Background()

var dataChan chan *ToES

// 初始化客户端
func init() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	ESClient, err = elastic.NewClient(
		elastic.SetErrorLog(errorlog),
		elastic.SetURL(servers...),
		elastic.SetGzip(true),
		elastic.SetHealthcheck(true),
		elastic.SetSniff(true), // docker env change to false
	)
	if err != nil {
		fmt.Printf("Access elasticsearch server failed. Reason: %s", err.Error())
		panic(err)
	}
	// 检查各节点连通性
	var index int
	var host string
	for index, host = range servers {
		info, code, err := ESClient.Ping(host).Do(ctx)
		if err != nil {
			// panic(err)
			fmt.Printf("Elasticsearch return with code %d, access failed. Reason: %s\n", err.Error())
		} else {
			fmt.Printf("Elasticsearch return with code %d and version %s\n ", code, info.Version.Number)
		}
	}
	dataChan = make(chan *ToES, 30000)
	go WriteES(dataChan)
	fmt.Printf("Total configuration %d nodes, %d connections succeeded!\n", len(servers), index+1)
}

// 检查索引是否存在，返回bool
func CheckIndexExists(indexName string) bool {
	exists, err := ESClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	return exists
}

// 通用save结构体
type ToES struct {
	IndexName string
	TypeDoc   string
	Data      interface{}
}

// 信息先暂存通道缓冲
func (ste *ToES) Save() {
	dataChan <- ste
}

// 重管道读数据并发写入es
func WriteES(dc chan *ToES) {
	for data := range dataChan {
		fmt.Println(data.IndexName)
		fmt.Println(data.TypeDoc)
		fmt.Println(data.Data)
	}
}

// // 创建索引 (mapping内容用BodyString传入)
// func CreateIndex(indexName string, mapping string) error {
// 	if !CheckIndexExists(indexName) {
// 		_, err := ESClient.CreateIndex(indexName).BodyString(mapping).Do(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// 	return nil
// }
