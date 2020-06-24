package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

const flinkUrl = "https://flinklog.yimidida.com/"

type FlinkJobs struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type JobInfo struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

func GetRequest(url string) (result []byte, err error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	resp, err := myClient.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return body, nil
}

func main() {
	//flinkClusterList := map[string]string{
	//	"阿里云FlinK":  "https://flinklog.yimidida.com/",
	//	"纪蕴机房Flink": "http://10.200.3.26:8081/",
	//}
	//for c, v := range flinkClusterList {
	//
	//}

	//resultByte, err := GetRequest(flinkUrl + "jobs")
	//if err != nil {
	//	os.Exit(0)
	//}
	//var flinkJobs FlinkJobs
	//err = json.Unmarshal([]byte(resultByte), &flinkJobs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var alarmMsg string
	//for _, v := range flinkJobs.Jobs {
	//	var jobInfo JobInfo
	//	jobResultByte, _ := GetRequest(flinkUrl + "jobs/" + v.Id)
	//	err = json.Unmarshal(jobResultByte, &jobInfo)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	nameSlice := strings.Split(jobInfo.Name, " ")
	//	topic := nameSlice[len(nameSlice)-1]
	//	if jobInfo.State == "RUNNING" {
	//		msg := fmt.Sprintf("Job: %s 状态: %s\n", topic, jobInfo.State)
	//		alarmMsg += msg
	//	}
	//}
	//fmt.Printf("------>", alarmMsg)
}
