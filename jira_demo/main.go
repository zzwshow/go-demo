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
	_, _, err = jiraCon.Issue.Get(issueId, nil)
	if err != nil {
		return false
	}
	return true
}

func main() {
	jiraClient := new(JiraClient)
	jiraClient.Url = "https://jira.yimidida.com/"
	jiraClient.Username = "zhangzhiwei"
	jiraClient.Password = "zhiwei123"
	// jiraCon, err := jiraClient.GetJiraConObj()
	// if err != nil {
	// 	fmt.Println("jira url 连接失败...")
	// }

	res := jiraClient.ExitsIssueByOps("ops-3291")
	if res {
		fmt.Println("ok")
		return
	}
	fmt.Println("no")

	// fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)

}
