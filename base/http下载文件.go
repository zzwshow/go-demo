package main

import (
	"net/http"
	"os"
)

func downloadFile(filePath string, url string) error {
	// 获取数据
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

}

func main() {
	fileUrl := "https://mojotv.cn/assets/image/logo01.png"

}
