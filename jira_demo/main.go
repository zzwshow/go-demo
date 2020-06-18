package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

type JiraClient struct {
	Url      string
	Username string
	Password string
}

func (js *JiraClient) GetJiraConObj() (jiraClient *jira.Client, err error) {
	tp := jira.BasicAuthTransport{
		Username: js.Username,
		Password: js.Password,
	}
	jiraClient, err = jira.NewClient(tp.Client(), js.Url)
	if err != nil {
		return nil, err
	}
	return
}

func (js *JiraClient) ExitsIssueByOps(issueId string) bool {
	jiraCon, err := js.GetJiraConObj()
	if err != nil {
		fmt.Println("jira url 连接失败...")
	}
	iss, res, err := jiraCon.Issue.Get(issueId, nil)
	if err != nil && res.Status != "200" {
		fmt.Println("获取 问题信息 failed。。。")
		return false
	}
	fmt.Printf("issue: %#v", iss.Fields.Status.StatusCategory.Name)
	return true
}

// func (js *JiraClient) IsEndFinish(issueId string) bool {
// 	jiraCon, err := js.GetJiraConObj()
// 	if err != nil {
// 		fmt.Println("jira url 连接失败...")
// 	}
// 	issue, info, err := jiraCon.Issue.Get(issueId, nil)
// 	if err != nil {
// 		return false
// 	}
// 	fmt.Println("info_status:  ", info.Close)
// 	fmt.Println("info:  ", issue)
// 	return true
// }

func main() {
	jiraClient := new(JiraClient)
	jiraClient.Url = "https://jira.uce.cn/"
	jiraClient.Username = "git"
	jiraClient.Password = "git.ymdd.com"
	// jiraCon, err := jiraClient.GetJiraConObj()
	// if err != nil {
	// 	fmt.Println("jira url 连接失败...")
	// }

	res := jiraClient.ExitsIssueByOps("ops-2828")
	if res {
		fmt.Println("ok")
		return
	}
	fmt.Println("no")

	// b := jiraClient.IsEndFinish("ops-2828")
	// if b {
	// 	fmt.Println(" 又ok")
	// 	return
	// }

	fmt.Println(" 又on")
	// fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)

}
