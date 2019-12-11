package main

import "go-elastic/esutil"

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Home string `json:"home"`
}

func main() {
	indexName := "libra"
	type_doc := "test"
	data := person{Name:"张志伟",Age:29,Home:"河南"}
	// put1,err := esutil.ESClient.Index().Index(indexName).Type(type_doc).BodyJson(person_1).Do(context.Background())
	// if err != nil{
	// 	fmt.Printf("写入es 失败")
	// }
	// fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	
	savedata := esutil.ToES{IndexName:indexName,TypeDoc:type_doc,Data:data}
	savedata.Save()
	select {
	
	}
}
