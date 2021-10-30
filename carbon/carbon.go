package carbon

import (
	"fmt"
	"time"
)

type Carbon struct {
}

func New() *Carbon {
	return new(Carbon)
}

// Now 获取当前时间
func (this *Carbon) Now() time.Time {
	return time.Now()
}

// Parse 字符串转时间
func (this *Carbon) Parse(value string, layout string) time.Time {
	if len(layout) == 0 {
		layout = "2006-01-02 15:04:05"
	}
	timeStamp, _ := time.ParseInLocation(layout, value, time.Local)
	return timeStamp
}

// ParseTimeStamp 通过时间戳获取时间
func (this *Carbon) ParseTimeStamp(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// GetStartOfDay 获取一天的开始时间
func (this *Carbon) GetStartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// GetEndOfDay 获取一天的结束时间
func (this *Carbon) GetEndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
}

// GetStartOfMonth 获取一个月的开始时间
func (this *Carbon) GetStartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

// GetEndOfMonth 获取一个月的结束时间
func (this *Carbon) GetEndOfMonth(date time.Time) time.Time {
	lastDay := this.GetStartOfMonth(date).AddDate(0, 1, -1)
	return time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, date.Location())
}

// 获取一年的开始时间
func (this *Carbon) GetStartOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), 1, 1, 0, 0, 0, 0, date.Location())
}

// GetEndOfYear 获取一年的结束时间
func (this *Carbon) GetEndOfYear(date time.Time) time.Time {
	lastDay := this.GetStartOfYear(date).AddDate(1, 0, -1)
	return time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, date.Location())
}

// SubDay 减去指定天数
func (this *Carbon) SubDay(date time.Time, day int) time.Time {
	return date.AddDate(0, 0, -day)
}

// FormatSeconds 格式化时间
func (this *Carbon) FormatSeconds(seconds int64) string {
	if seconds >= 86400 {
		return fmt.Sprintf("%d天%s", seconds / 86400, this.ParseTimeStamp(seconds - 8 * 3600).Format("15:04:05"))
	}
	return this.ParseTimeStamp(seconds - 8 * 3600).Format("15:04:05")
}