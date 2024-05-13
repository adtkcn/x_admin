package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DateFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"

// TsTime 自定义时间格式
type TsTime time.Time
type OnlyRespTsTime time.Time

// func (tst *TsTime) UnmarshalJSON(bs []byte) error {
// 	var date string
// 	err := json.Unmarshal(bs, &date)
// 	if err != nil {
// 		return err
// 	}
// 	tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
// 	*tst = TsTime(tt.Unix())
// 	return nil
// }

// // MarshalJSON 将TsTime类型的时间转化为JSON字符串格式
// // 返回转化后的JSON字符串和错误信息
// func (tst TsTime) MarshalJSON() ([]byte, error) {
// 	tt := time.Unix(int64(tst), 0).Format(TimeFormat)
// 	return json.Marshal(tt)
// }

// 通过时间字符串生成时间戳
func ToUnix(date string) int64 {
	if date == "" {
		return 0
	}
	tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
	return tt.Unix()
}
func NowTime() TsTime {
	return TsTime(time.Now())
}

func (otst OnlyRespTsTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(otst).Format(TimeFormat)
	return json.Marshal(tt)
}

// type TsTime time.Time

func (tst *TsTime) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}
	tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
	*tst = TsTime(tt)
	return nil
}

// MarshalJSON 将TsTime类型的时间转化为JSON字符串格式
// 返回转化后的JSON字符串和错误信息
func (tst TsTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(tst).Format(TimeFormat)
	return json.Marshal(tt)
}

// 写入数据库gorm调用
func (t TsTime) Value() (driver.Value, error) {
	timeStr := t.String()
	if timeStr == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return timeStr, nil
}

// 读取数据gorm调用
func (t *TsTime) Scan(v any) error {
	// pt, err := time.ParseInLocation("2006-01-02 15:04:05", v.(time.Time).String(), time.Local)
	if pt, ok := v.(time.Time); ok {
		*t = TsTime(pt)
		return nil
	}
	return fmt.Errorf("cant convert %s to time", v)
}

func (t TsTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

func (TsTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// 使用 field.Tag、field.TagSettings 获取字段的 tag
	// 查看 https://github.com/go-gorm/gorm/blob/master/schema/field.go 获取全部的选项

	// 根据不同的数据库驱动返回不同的数据类型
	// switch db.Dialector.Name() {
	// case "mysql", "sqlite":
	//   return "JSON"
	// case "postgres":
	//   return "JSONB"
	// }
	// return ""
	return "DATETIME"
}
