package daycounter

import (
	"fmt"
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
		formatStr    string
	)
	if timeStr == "now" {
		timeObj = time.Now()
		return timeObj
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	switch {
	case len(timeStr) == dayStrLen:
		formatStr = "2006/01/02"
	case len(timeStr) == hourStrLen:
		formatStr = "2006/01/02-15"
	case len(timeStr) == minuteStrLen:
		formatStr = "2006/01/02-15:04"
	case len(timeStr) == secondStrLen:
		formatStr = "2006/01/02-15:04:05"
	default:
		fmt.Println("format error")
		return timeObj
	}
	timeObj, err := time.ParseInLocation(formatStr, timeStr, loc)
	if err != nil {
		fmt.Println("Format error, try again", err)
	}
	return timeObj
}

// DuringDay 天数差
func (t *TimeCount) DuringDay(t2 *TimeCount) int {
	var durTime time.Duration = t2.FromNow - t.FromNow
	var durDay int = int(durTime.Hours() / 24)
	return durDay
}

// DuringHour 小时差
func (t *TimeCount) DuringHour(t2 *TimeCount) int {
	var durTime time.Duration = t2.FromNow - t.FromNow
	var durHour int = int(durTime.Hours())
	return durHour
}

// DuringMinute 分钟差
func (t *TimeCount) DuringMinute(t2 *TimeCount) int {
	var durTime time.Duration = t2.FromNow - t.FromNow
	var durMinute int = int(durTime.Minutes())
	return durMinute
}

// DuringSecond 秒差
func (t *TimeCount) DuringSecond(t2 *TimeCount) int {
	var durTime time.Duration = t2.FromNow - t.FromNow
	var durSecond int = int(durTime.Seconds())
	return durSecond
}

// DuringDetail 详细时间差
func (t *TimeCount) DuringDetail(t2 *TimeCount) string {
	var (
		durTime time.Duration = t2.FromNow - t.FromNow
		minute  int           = int(durTime.Minutes()) % 60
		hour    int           = int(durTime.Hours()) % 24
		day     int           = int(durTime.Hours()) / 24
		second  int           = int(durTime.Seconds()) % 60
	)
	var detail string = fmt.Sprintf("%v days %v hours %v minute %v seconds", day, hour, minute, second)
	return detail
}

// AfterSeconds 获得若干秒之后的时间
func (t *TimeCount) AfterSeconds(second int) time.Time {
	var counts time.Duration = time.Duration(second)
	var target time.Time = t.Time.Add(counts * time.Second)
	return target
}

// AfterMinutes 获得若干分钟后的时间
func (t *TimeCount) AfterMinutes(minute int) time.Time {
	var counts time.Duration = time.Duration(minute)
	var target time.Time = t.Time.Add(counts * time.Minute)
	return target
}

// AfterHours 获得若干小时后的时间
func (t *TimeCount) AfterHours(hour int) time.Time {
	var counts time.Duration = time.Duration(hour)
	var target time.Time = t.Time.Add(counts * time.Hour)
	return target
}

// AfterDays 获得若干天后的时间
func (t *TimeCount) AfterDays(day int) time.Time {
	var hour int = day * 24
	return t.AfterHours(hour)
}
