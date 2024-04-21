package core

import (
	"encoding/json"
	"time"
)

const DateFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"

// TsTime 自定义时间格式
type TsTime int64
type OnlyRespTsTime time.Time

// //TsDate 自定义日期格式
// type TsDate int64
//
//	func (tsd *TsDate) UnmarshalJSON(bs []byte) error {
//		var date string
//		err := json.Unmarshal(bs, &date)
//		if err != nil {
//			return err
//		}
//		tt, _ := time.ParseInLocation(DateFormat, date, time.Local)
//		*tsd = TsDate(tt.Unix())
//		return nil
//	}
//
//	func (tsd TsDate) MarshalJSON() ([]byte, error) {
//		tt := time.Unix(int64(tsd), 0).Format(DateFormat)
//		return json.Marshal(tt)
//	}
//

func (tst *TsTime) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}
	tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
	*tst = TsTime(tt.Unix())
	return nil
}

// MarshalJSON 将TsTime类型的时间转化为JSON字符串格式
// 返回转化后的JSON字符串和错误信息
func (tst TsTime) MarshalJSON() ([]byte, error) {
	tt := time.Unix(int64(tst), 0).Format(TimeFormat)
	return json.Marshal(tt)
}

// 通过时间字符串生成时间戳
func ToUnix(date string) int64 {
	if date == "" {
		return 0
	}
	tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
	return tt.Unix()
}

func (otst OnlyRespTsTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(otst).Format(TimeFormat)
	return json.Marshal(tt)
}
