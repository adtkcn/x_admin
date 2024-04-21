package core

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
// 实现自定义的解析逻辑
func (t *TsTime) Bind(ctx *gin.Context) error {
	// 尝试从表单中获取时间字符串
	str := ctx.Query("clientTime") // 对于GET请求使用Query，对于POST请求使用PostForm
	// 对于POST请求中的JSON体，你应该使用 ShouldBindJSON 或其他相关的ShouldBind方法

	// 假设传入的是UNIX时间戳的字符串形式
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	*t = TsTime(i)
	return nil
}

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

func (otst OnlyRespTsTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(otst).Format(TimeFormat)
	return json.Marshal(tt)
}
