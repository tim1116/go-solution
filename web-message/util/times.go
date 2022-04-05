package util

import "time"

const Layout = "2006-01-02 15:04:05"

// 类似PHP date()
func Date(format string, unixTime int64) string {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	tm := time.Unix(unixTime, 0).In(cstSh)
	return tm.Format(format)
}
