package libtime

import (
	"fmt"
	"strings"
	"time"
)

type MyTime struct {
	// 时区设置
	location *time.Location
}

// NewTime 初始化对象, 默认使用本地时区(使用hours的原因是用默认本地的可以不传参数)
func NewTime(hours ...int) *MyTime {
	mt := &MyTime{}
	if len(hours) == 0 {
		mt.location = time.Local
	} else {
		mt.location = time.FixedZone("CST", hours[0]*3600)
	}
	return mt
}

// SetTimeZoneOffset SetLocation 设置时区(8为东8区)
func (mt *MyTime) SetTimeZoneOffset(hours int) {
	mt.location = time.FixedZone("CST", hours*3600)
}

// SetLocation 直接设置时区位置
func (mt *MyTime) SetLocation(location *time.Location) {
	mt.location = location
}

// NowTimestamp 获取当前时间戳
func (mt *MyTime) NowTimestamp() int64 {
	// 时间戳无时区之分，time.Now().Unix()与之相同
	return time.Now().In(mt.location).Unix()
}

// Date 将时间戳解析成对应格式的时间格式
func (mt *MyTime) Date(format string, ts int64) string {
	return mt.FormatTimestamp(format, ts)
}

// FormatTimestamp 将时间戳解析成对应格式的时间格式
func (mt *MyTime) FormatTimestamp(format string, ts int64) string {
	var tt time.Time
	if ts == 0 {
		tt = time.Now().In(mt.location)
	} else {
		tt = time.Unix(ts, 0).In(mt.location)
	}
	return mt.Format(format, tt)
}

// Format 跟 PHP 中 date 类似的使用方式，如果 ts 没传递，则使用当前时间
func (mt *MyTime) Format(format string, tts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	var tt time.Time
	if len(tts) > 0 {
		tt = tts[0]
	} else {
		tt = time.Now().In(mt.location)
	}
	return tt.Format(format)
}

// StrToLocalTime 将字符串格式的本地时间表示为Go语言的时间类型
func (mt *MyTime) StrToLocalTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, fmt.Errorf("输入错误")
	}
	// 当前系统的时区相比UTC时间的偏移量
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return mt.StrToTime(value)
}

// StrToTime 将时间字符串转换为Go的时间类型
func (mt *MyTime) StrToTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, fmt.Errorf("输入错误")
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var tt time.Time
	var err error
	for _, layout := range layouts {
		tt, err = time.ParseInLocation(layout, value, mt.location)
		if err == nil {
			return tt, nil
		}
	}
	return tt, err
}

// StrToTimestamp 将字符串转为时间戳
func (mt *MyTime) StrToTimestamp(value string) int64 {
	tt, err := mt.StrToTime(value)
	if err == nil {
		return tt.Unix()
	} else {
		return 0
	}
}

// StartTimestampOfDay 当天零点的时间戳
func (mt *MyTime) StartTimestampOfDay(value string) int64 {
	tt, err := mt.StrToTime(value)
	if err != nil {
		return 0
	}
	tt = time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, mt.location)
	return tt.Unix()
}

// StartTimestampOfDayTime 当天零点的时间戳
func (mt *MyTime) StartTimestampOfDayTime(tt time.Time) int64 {
	tt = time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, mt.location)
	return tt.Unix()
}
