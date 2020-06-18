package main

type MarkdownMsgStruct struct {
	Title     string `json:"title"`
	Source    string `json:"source"`
	ChatId    string `json:"chat_id"`
	AlarmType string `json:"alarm_type"` // 类型：基础资源，应用监控，业务监控
	Env       string `json:"env"`        // 环境: libra 内定义的几种环境
	Business  string `json:"business"`   // 业务线: 告警来自哪个业务线
	Module    string `json:"module"`     // 模块: 模块名称
	Content   string `json:"content"`    // 内容: 具体的告警信息
	Level     string `json:"level"`      // 告警级别：1: fatal 2: error 3:Warning 4: info
	AlarmTime string `json:"alarm_time"` // 告警时间
}

func main() {

}
