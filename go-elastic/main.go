package main

import (
	"fmt"
	"time"
	
	"go-elastic/esutil"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Home string `json:"home"`
	Create_At time.Time `json:"create_at"`
}

func main() {
	indexName := "libra"
	// type_doc := "test"
	// tf := "2006-01-02 15:04:05.000"
	// data := person{Name:"demo",Age:29,Home:"aaaa",Create_At:time.Now()}
	// // put1,err := esutil.ESClient.Index().Index(indexName).Type(type_doc).BodyJson(data).Do(context.Background())
	// // if err != nil {
	// // 	fmt.Printf("写入es 失败")
	// // }
	// // fmt.Printf("Indexed libra %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	//
	// savedata := esutil.ToES{IndexName:indexName,TypeDoc:type_doc,Data:data}
	// // savedata.Save()
	// select {
	// }
	
	err := esutil.DeleteEsIndex(indexName)
	if err != nil{
		fmt.Printf("delete failed err :%s",err.Error() )
	}
	
	
	
}
