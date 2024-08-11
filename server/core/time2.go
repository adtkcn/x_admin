package core

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DateFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"

// 注解:时间类型从time.Time改为string的原因是struct转interface{}时，丢弃了部分信息，导致导出excel时，时间格式不对
// TsTime 自定义时间格式
type TsTime string

// 通过时间字符串生成时间戳
//
//	func ToUnix(date string) int64 {
//		if date == "" {
//			return 0
//		}
//		tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
//		return tt.Unix()
//	}
func ParseTimeToTsTime(date time.Time) TsTime {
	return TsTime(date.Format(TimeFormat))
}

func NowTime() TsTime {
	return TsTime(time.Now().Format(TimeFormat))
}

func (tst *TsTime) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}
	tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
	*tst = TsTime(
		tt.Format(TimeFormat),
	)
	return nil
}

// MarshalJSON 将TsTime类型的时间转化为JSON字符串格式
// 返回转化后的JSON字符串和错误信息
func (tst TsTime) MarshalJSON() ([]byte, error) {
	// tt := time.Time(tst.Str).Format(TimeFormat)
	tt, _ := time.Parse(TimeFormat, tst.String())
	str := tt.Format(TimeFormat)

	return json.Marshal(str)
}

// 写入数据库gorm调用
func (t TsTime) Value() (driver.Value, error) {
	// timeStr := t.String()
	// if timeStr == "0001-01-01 00:00:00" {
	// 	return nil, nil
	// }
	return t.String(), nil
}

// 读取数据gorm调用
func (t *TsTime) Scan(v any) error {
	// pt, err := time.ParseInLocation("2006-01-02 15:04:05", v.(time.Time).String(), time.Local)
	if _, ok := v.(time.Time); ok {
		*t = TsTime(v.(time.Time).Format(TimeFormat))
		return nil
	} else {
		*t = "0001-01-01 00:00:00"
		return nil
	}
}

func (t TsTime) String() string {
	return string(t)
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
