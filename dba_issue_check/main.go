package dba_issue_check

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
根据ops 编号获取ops状态
*/

type DBAResponseStatus struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

const DBA_WEBSQL = "https://dba.yimidida.com/workflow_api/"
const opsCode = "ops-1321"

// const headers = `{"User-Agent": "Mozilla/5.0"}`

func DBAOpsStatusCheck(opsCode string) bool {
	opsID := strings.Split(opsCode, "-")[1]
	fmt.Println(opsID)
	queryKey := fmt.Sprintf("workflow_id=%d", opsID)
	httpClient := &http.Client{}
	httpClient.Timeout = time.Second * 5
	fmt.Println(DBA_WEBSQL)
	fmt.Println(opsID)
	req, err := http.NewRequest("GET", DBA_WEBSQL, strings.NewReader(queryKey))
	if err != nil {
		fmt.Println("连接失败... error : ", err)
		return false
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("从 dba.yimidida.com 获取ops 状态失败... error : ", err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取 dba.yimidida.com 的ops 状态失败... error : ", err)
		return false
	}
	fmt.Println(string(body))
	return true
}

func TestCheckOpsStatus(opsCode string) bool {
	opsID := strings.Split(opsCode, "-")[1]
	queryUrlPara := fmt.Sprintf("workflow_api/?workflow_id=%s", opsID)
	url := fmt.Sprintf("https://dba.yimidida.com/%s", queryUrlPara)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("连接 dba.yimidida.com  失败... error : ", err)
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var dbaRS DBAResponseStatus
	err = json.Unmarshal(body, &dbaRS)
	if err != nil {
		fmt.Println("获取 dba.yimidida.com ops 状态 解析返回值失败... error: ", err)
		return false
	}
	fmt.Println(dbaRS.Status)
	if dbaRS.Status != "workflow_finish" {
		return false
	}
	return true
}

func main() {
	b := TestCheckOpsStatus(opsCode)
	fmt.Println(b)
}
