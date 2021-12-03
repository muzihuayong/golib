package libtime

import (
	"time"
)

var myTime *MyTime

// 定义一些常用的格式
const (
	FORMAT_STANDARD       = "2006-01-02 15:04:05"
	FORMAT_STANDARD_SLASH = "2006/01/02 15:04:05"
	FORMAT_DATE           = "2006-01-02"
	FORMAT_DATE_SLASH     = "2006/01/02"
	FORMAT_DAYTIME        = "15:04:05"
)

func init() {
	myTime = NewTime()
}

// SetLocation 设置当前时区(8为东8区)
func SetTimeZoneOffset(hour int) {
	myTime.SetTimeZoneOffset(hour)
}

// NowTimestamp 获取当前时间戳
func NowTimestamp() int64 {
	return myTime.NowTimestamp()
}

// Date 将时间戳解析成对应格式的时间格式
func Date(format string, ts int64) string {
	return myTime.FormatTimestamp(format, ts)
}

// FormatTimestamp 将时间戳解析成对应格式的时间格式
func FormatTimestamp(format string, ts int64) string {
	return myTime.FormatTimestamp(format, ts)
}

// Format 跟 PHP 中 date 类似的使用方式，如果 ts 没传递，则使用当前时间
func Format(format string, ts ...time.Time) string {
	return myTime.Format(format, ts...)
}

// StrToLocalTime 将字符串格式的本地时间表示为Go语言的时间类型
func StrToLocalTime(value string) (time.Time, error) {
	return myTime.StrToLocalTime(value)
}

// StrToTime 将时间字符串转换为Go的时间类型
func StrToTime(value string) (time.Time, error) {
	return myTime.StrToTime(value)
}

// StrToTimestamp 将字符串转为时间戳
func StrToTimestamp(value string) int64 {
	return myTime.StrToTimestamp(value)
}

// StartTimestampOfDay 当天零点的时间戳
func StartTimestampOfDay(value string) int64 {
	ts, err := myTime.StrToTime(value)
	if err != nil {
		return 0
	}
	ts = time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, myTime.location)
	return ts.Unix()
}

// StartTimestampOfDayTime 当天零点的时间戳
func StartTimestampOfDayTime(ts time.Time) int64 {
	ts = time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, myTime.location)
	return ts.Unix()
}
