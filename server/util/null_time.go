package util

import (
	"time"

	"github.com/adtkcn/x_null"
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

// ParseTime 时间转时间戳
func (t nullTimeUtil) ParseTime(date time.Time) x_null.Time {
	return x_null.Time{
		Val:   &date,
		Exist: true,
	}
}

// NowTime 当前时间
func (t nullTimeUtil) Now() x_null.Time {
	now := time.Now()
	return x_null.Time{
		Val:   &now,
		Exist: true,
	}
}

// 今日0点
func (t nullTimeUtil) TodayZero() x_null.Time {
	now := time.Now()
	todayZero := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return x_null.Time{
		Val:   &todayZero,
		Exist: true,
	}
}

// DecodeTime 时间解码
// func (t nullTimeUtil) DecodeTime(value any) (any, error) {
// 	tt, e := t.Parse(value)
// 	return tt, e
// }

// ToUnix 时间戳转时间戳
// func (t nullTimeUtil) ToUnix(date any) int64 {
// 	switch v := date.(type) {
// 	case string:
// 		if v == "" {
// 			return 0
// 		}
// 		tt, e := time.Parse(t.TimeFormat, v)
// 		if e != nil {
// 			return 0
// 		}
// 		return time.Time(tt).Unix()
// 	case time.Time:
// 		return v.Unix()
// 	default:
// 		return 0
// 	}
// }

// Parse 时间戳转时间
// func (t nullTimeUtil) Parse(value any) (x_null.Time, error) {
// 	switch v := value.(type) {
// 	case string:
// 		tt, e := t.ParseString(v)
// 		return tt, e
// 	case time.Time:
// 		tt := t.ParseTime(v)
// 		return tt, nil
// 	case x_null.Time:
// 		return v, nil
// 	default:
// 		return t.Null(), errors.New("时间格式错误")
// 	}
// }

// ParseString 时间字符串转时间戳
// func (t nullTimeUtil) ParseString(date string) (x_null.Time, error) {
// 	tt, e := time.Parse(t.TimeFormat, date)
// 	if e != nil {
// 		return t.Null(), e
// 	}
// 	return x_null.Time{
// 		Val:   &tt,
// 		Exist: true,
// 	}, nil
// }

// Null 返回空时间
// func (t nullTimeUtil) Null() x_null.Time {
// 	return x_null.Time{
// 		Val:   nil,
// 		Exist: false,
// 	}
// }
