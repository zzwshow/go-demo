package main

import "time"

// StrToTime 字符串转datetime
func StrToTime(s string) time.Time {
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	timeLayout := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(timeLayout, s, nyc)
	return t
}

// StrToTime 字符串转date
func StrToDateTime(s string) time.Time {
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	timeLayout := "2006-01-02"
	t, _ := time.ParseInLocation(timeLayout, s, nyc)
	return t
}

func main() {

}
