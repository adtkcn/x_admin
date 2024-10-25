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

// NullTime 自定义时间格式
type NullTime struct {
	Time   *time.Time
	Valid  bool
	Format string
}

//	func (t *NullTime) IsZero() bool {
//		return t.Valid
//	}
func (t *NullTime) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}
	if date == "" {
		*t = NullTime{
			Time:  nil,
			Valid: false,
		}
		return nil
	}
	tt, _ := time.ParseInLocation(TimeFormat, date, time.Local)
	*t = NullTime{
		Time:  &tt,
		Valid: true,
	}
	return nil
}

// MarshalJSON 将NullTime类型的时间转化为JSON字符串格式
// 返回转化后的JSON字符串和错误信息
func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		if t.Time == nil {
			return json.Marshal(nil)
		}
		tt := *t.Time
		tStr := tt.Format(TimeFormat)
		return json.Marshal(tStr)
	} else {
		return json.Marshal(nil)
	}
}

// 写入数据库gorm调用
func (t NullTime) Value() (driver.Value, error) {
	timeStr := t.String()
	if timeStr == "" {
		return nil, nil
	}
	return timeStr, nil
}

// 读取数据gorm调用
func (t *NullTime) Scan(v any) error {
	// pt, err := time.ParseInLocation("2006-01-02 15:04:05", v.(time.Time).String(), time.Local)
	if pt, ok := v.(time.Time); ok {
		tt := pt.Format(TimeFormat)
		if tt == "0001-01-01 00:00:00" {
			*t = NullTime{
				Time:  &pt,
				Valid: false,
			}
			return nil
		} else {
			*t = NullTime{
				Time:  &pt,
				Valid: true,
			}
		}

		return nil
	}
	return fmt.Errorf("cant convert %s to time", v)
}

func (t NullTime) String() string {
	if !t.Valid {
		return ""
	}
	if t.Time == nil {
		return ""
	}
	tt := *t.Time
	return tt.Format(TimeFormat)
}

func (NullTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
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
