package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
)

//"aliyun_voice":
//"REGION": "cn-hangzhou"
//"PRODUCT_NAME": "Dyvmsapi"
//"DOMAIN": "dyvmsapi.aliyuncs.com"
//"ACCESS_KEY_ID": "LTAI4FghuSDX1pzpKqCX2cL1"
//"ACCESS_KEY_SECRET": "zhV5IJMdyMyWodrLGrfXGRyAIegu24"
//"CalledShowNumber": "01086466077"
//"TtsCode": "TTS_172743869"

func main() {
	client, err := dyvmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FghuSDX1pzpKqCX2cL1", "zhV5IJMdyMyWodrLGrfXGRyAIegu24")

	request := dyvmsapi.CreateSingleCallByTtsRequest()
	request.Scheme = "https"
	request.CalledShowNumber = "01086466077"
	request.CalledNumber = "18839399820"
	request.TtsCode = "TTS_172743869"
	request.TtsParam = `{"serverid":"172点16点0点1","fault":"启动失败了"}`
	response, err := client.SingleCallByTts(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
