package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	// 个人邮件服务器地址
	SMTP_MAIL_HOST = "smtp.126.com"
	SMTP_MAIL_PORT = "25"
	SMTP_MAIL_FROM = "wei3511@126.com"
	SMTP_MAIL_USER = "wei3511@126.com"
	SMTP_MAIL_PWD  = "RKFOPZRRTWJXGGLG"
)

func SendSMTPMail(mailAddress string, subject string, body string) error {
	// 通常身份应该是空字符串，填充用户名.
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	// (text/plain)纯文本 , (text/html)html
	contentType := "Content-Type: text/plain; charset=UTF-8"
	nickname := "SMTPMail"
	msg := []byte("To: " + mailAddress + "\r\nFrom: " + nickname + "<" + SMTP_MAIL_FROM + ">\r\nSubject: " + subject +
		"\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(mailAddress, ",")
	err := smtp.SendMail(fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT), auth, SMTP_MAIL_FROM, sendTo, msg)
	return err
}

func main() {
	err := SendSMTPMail("605882219@qq.com", "hello，这是smtp测试邮件。", "----->的测试邮件")
	if err != nil {
		fmt.Println(err.Error())
	}
}
