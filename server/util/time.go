package util

import (
	"errors"
	"time"
	"x_admin/core"
)

var TimeUtil = timeUtil{
	DateFormat: "2006-01-02",
	TimeFormat: "2006-01-02 15:04:05",
}

// arrayUtil 数组工具类
type timeUtil struct {
	DateFormat string
	TimeFormat string
}

// ToUnix 时间戳转时间戳
func (t timeUtil) ToUnix(date any) int64 {
	switch v := date.(type) {
	case string:
		if v == "" {
			return 0
		}
		tt, e := time.Parse(t.TimeFormat, v)
		if e != nil {
			return 0
		}
		return time.Time(tt).Unix()
	case time.Time:
		return v.Unix()
	default:
		return 0
	}
}

// ParseTsTime 时间戳转时间
func (t timeUtil) ParseTsTime(value any) (core.TsTime, error) {
	switch v := value.(type) {
	case string:
		tt, e := t.ParseString(v)
		return tt, e
	case time.Time:
		tt := t.ParseTime(v)
		return tt, nil
	default:
		return t.NowTime(), errors.New("not support type")
	}
}

// ParseTime 时间转时间戳
func (t timeUtil) ParseTime(date time.Time) core.TsTime {
	return core.TsTime(date)
}

// ParseString 时间字符串转时间戳
func (t timeUtil) ParseString(date string) (core.TsTime, error) {
	tt, e := time.Parse(t.TimeFormat, date)
	if e != nil {
		return t.NowTime(), e
	}
	return core.TsTime(tt), nil
}

// NowTime 当前时间
func (t timeUtil) NowTime() core.TsTime {
	return core.TsTime(time.Now())
}

// DecodeTime 时间解码
func (t timeUtil) DecodeTime(value any) (any, error) {
	tt, e := t.ParseTsTime(value)
	return tt, e
}
