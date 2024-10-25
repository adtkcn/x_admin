package util

import (
	"time"
	"x_admin/core"
)

var NullTimeUtil = nullTimeUtil{
	DateFormat: "2006-01-02",
	TimeFormat: "2006-01-02 15:04:05",
}

// arrayUtil 数组工具类
type nullTimeUtil struct {
	DateFormat string
	TimeFormat string
}

// DecodeTime 时间解码
func (t nullTimeUtil) DecodeTime(value any) (any, error) {
	tt, e := t.Parse(value)
	return tt, e
}

// EncodeTime 时间编码
func (t nullTimeUtil) EncodeTime(value any) any {
	var str = ""
	switch v := value.(type) {
	case map[string]interface{}:
		val := v["Time"]
		switch tt := val.(type) {
		case *string:
			ttt, _ := t.Parse(*tt)
			str = ttt.String()
		case *time.Time:
			if tt != nil {
				ttt, _ := t.Parse(*tt)
				str = ttt.String()
			} else {
				str = ""
			}
		}

	}
	return str
}

// ToUnix 时间戳转时间戳
func (t nullTimeUtil) ToUnix(date any) int64 {
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

// Parse 时间戳转时间
func (t nullTimeUtil) Parse(value interface{}) (core.NullTime, error) {
	switch v := value.(type) {
	case string:
		tt, e := t.ParseString(v)
		return tt, e
	case time.Time:
		tt := t.ParseTime(v)
		return tt, nil
	case core.NullTime:
		return v, nil
	default:
		return t.Null(), nil
	}
}

// ParseTime 时间转时间戳
func (t nullTimeUtil) ParseTime(date time.Time) core.NullTime {
	return core.NullTime{
		Time:  &date,
		Valid: true,
	}
}

// ParseString 时间字符串转时间戳
func (t nullTimeUtil) ParseString(date string) (core.NullTime, error) {
	tt, e := time.Parse(t.TimeFormat, date)
	if e != nil {
		return t.Null(), e
	}
	return core.NullTime{
		Time:  &tt,
		Valid: true,
	}, nil
}

// NowTime 当前时间
func (t nullTimeUtil) Null() core.NullTime {
	return core.NullTime{
		Time:  nil,
		Valid: true,
	}
}

// NowTime 当前时间
func (t nullTimeUtil) Now() core.NullTime {
	now := time.Now()
	return core.NullTime{
		Time:  &now,
		Valid: true,
	}
}
