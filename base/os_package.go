package main

import (
	"fmt"
	"os"
	"path"
)

func main(){
	filename := "/Users/zzw/Desktop/git/golang/server.log"
	fileObj, _ := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	old_name := fileObj.Name()
	dir_str := path.Dir(old_name)
	name_str := path.Base(old_name)
	fmt.Printf("目录：%s,文件名：%s",dir_str,name_str)
	new_name := fmt.Sprintf("%s_kkkk",old_name)
	fmt.Println("old",old_name)
	fmt.Println("new",new_name)
	err := os.Rename(old_name,new_name)
	if err !=nil{
		fmt.Println("lllllllllllll")
	}
}