package daycounter

import (
	"fmt"
	"strconv"
	"time"
)

// TimeCount 时间计算结构体
type TimeCount struct {
	Time    time.Time
	FromNow time.Duration
}

// NewTimeCount 构造函数，进行构造体初始化
func NewTimeCount(t time.Time) *TimeCount {
	var fromNow time.Duration = t.Sub(time.Now())
	return &TimeCount{
		Time:    t,
		FromNow: fromNow,
	}
}

// InputTime 根据字符串返回时间对象，输入now则返回现在的时间对象。输入格式:2020/01/01-01:02:03
func InputTime(timeStr string) time.Time {
	var (
		dayStrLen    int = len("2006/01/02")
		hourStrLen   int = len("2006/01/02-15")
		minuteStrLen int = len("2006/01/02-15:04")
		secondStrLen int = len("2006/01/02-15:04:05")
		timeObj      time.Time
	)
	if timeStr == "now" {
		timeObj := time.Now()
		return timeObj
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	switch {
	case len(timeStr) == dayStrLen:
		timeObj, err := time.ParseInLocation("2006/01/02", timeStr, loc)
		if err != nil {
			fmt.Println("Format error, try again", err)
		}
		return timeObj
	case len(timeStr) == hourStrLen:
		timeObj, err := time.ParseInLocation("2006/01/02-15", timeStr, loc)
		if err != nil {
			fmt.Println("Format error, try again", err)
		}
		return timeObj
	case len(timeStr) == minuteStrLen:
		timeObj, err := time.ParseInLocation("2006/01/02-15:04", timeStr, loc)
		if err != nil {
			fmt.Println("Format error, try again", err)
		}
		return timeObj
	case len(timeStr) == secondStrLen:
		timeObj, err := time.ParseInLocation("2006/01/02-15:04:05", timeStr, loc)
		if err != nil {
			fmt.Println("Format error, try again", err)
		}
		return timeObj
	default:
		fmt.Println("format error")
		return timeObj
	}
}

// DuringDay 天数差
func (t *TimeCount) DuringDay(t2 *TimeCount) int {
	var durTime = t2.FromNow - t.FromNow
	durDay := int(durTime.Hours() / 24)
	return durDay
}

// DuringHour 小时差
func (t *TimeCount) DuringHour(t2 *TimeCount) int {
	var durTime = t2.FromNow - t.FromNow
	durHour := int(durTime.Hours())
	return durHour
}

// DuringMinute 分钟差
func (t *TimeCount) DuringMinute(t2 *TimeCount) int {
	var durTime = t2.FromNow - t.FromNow
	durMinute := int(durTime.Minutes())
	return durMinute
}

// DuringSecond 秒差
func (t *TimeCount) DuringSecond(t2 *TimeCount) int {
	var durTime = t2.FromNow - t.FromNow
	durSecond := int(durTime.Seconds())
	return durSecond
}

// DuringDetail 详细时间差
func (t *TimeCount) DuringDetail(t2 *TimeCount) string {
	var durTime = t2.FromNow - t.FromNow
	second := strconv.Itoa(int(durTime.Seconds()) % 60)
	minute := strconv.Itoa(int(durTime.Minutes()) % 60)
	hour := strconv.Itoa(int(durTime.Hours()) % 24)
	day := strconv.Itoa(int(durTime.Hours() / 24))
	var detail = fmt.Sprintf("%s days %s hours %s minute %s seconds", day, hour, minute, second)
	return detail
}
