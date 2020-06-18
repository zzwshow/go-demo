package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type wechatConf struct {
	AgentID    string
	CorpID     string
	Secret     string
	Url_Prefix string
}

const (
	corpid     = "wwfb97e21d01fd4aee"
	corpsecret = "VUVVPxh0PKBT2Vn9la-PxDXKoSvvz1OVL_gCg9Q3Pm4"
	agentid    = "1000002"
)

type JSONToken struct {
	AccessToken string `json:"access_token"`
}

func GetAccessToken() (string, error) {
	getTokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, corpsecret)
	fmt.Println(getTokenUrl)
	c := &http.Client{}
	req, err := c.Get(getTokenUrl)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	var json_str JSONToken
	err = json.Unmarshal([]byte(body), &json_str)
	if err != nil {
		return "", err
	}
	return json_str.AccessToken, nil
}

//  自由组织content 内容
type TextMsg struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Safe    int    `json:"safe"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type Text struct{}

func (t *Text) GenerateGroupTextMsg(chatid, content string) (msg string, err error) {
	info := TextMsg{
		Chatid:  chatid,
		Msgtype: "text",
		Safe:    0,
		Text: struct {
			Content string `json:"content"`
		}{Content: content},
	}
	send_msg, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	return string(send_msg), nil
}

func (t *Text) SendMsg(msg string) error {
	access_token, err := GetAccessToken()
	if err != nil {
		fmt.Println("获取wechat API token 失败....")
	}
	send_url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=%s", access_token)
	client := &http.Client{}
	req, err := http.NewRequest("POST", send_url, bytes.NewBuffer([]byte(msg)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
