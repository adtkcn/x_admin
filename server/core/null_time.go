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
	Val   *time.Time
	Exist bool
	// Format string
}

// func DecodeNulLTime(value any) (any, error) {
// 	switch v := value.(type) {
// 	case string:
// 		if v == "" {
// 			return NullTime{
// 				Val:   nil,
// 				Exist: true,
// 			}, nil
// 		}
// 		tt, e := time.ParseInLocation(TimeFormat, v, time.Local)
// 		if e != nil {
// 			return NullTime{}, e
// 		}
// 		return NullTime{
// 			Val:   &tt,
// 			Exist: true,
// 		}, nil
// 	case time.Time:
// 		return NullTime{
// 			Val:   &v,
// 			Exist: true,
// 		}, nil

//		default:
//			return NullTime{}, errors.New("时间格式错误")
//		}
//	}
//
// 读取数据gorm调用
func (t *NullTime) Scan(v any) error {
	switch val := v.(type) {
	case nil:
		t.Val = nil
		t.Exist = true
		return nil
	case string:
		tt, err := time.ParseInLocation(TimeFormat, val, time.Local)
		if err != nil {
			return err
		}
		t.Val = &tt
		t.Exist = true
		return nil
	case time.Time:
		tt := val.Format(TimeFormat)
		if tt == "0001-01-01 00:00:00" {
			t.Val = nil
			t.Exist = true
		} else {
			t.Val = &val
			t.Exist = true
		}
		return nil
	default:
		return fmt.Errorf("不能将类型 %T 转换为 time.Time, 值为 %v", v, v)
	}
	// return fmt.Errorf("NullTime cant convert %s", v)
}

// 写入数据库gorm调用
func (t NullTime) Value() (driver.Value, error) {
	timeStr := t.String()
	if timeStr == "" {
		return nil, nil
	}
	return timeStr, nil
}
func (t NullTime) String() string {
	if !t.Exist {
		return ""
	}
	if t.Val == nil {
		return ""
	}
	tt := *t.Val
	return tt.Format(TimeFormat)
}

// MarshalJSON 将NullTime类型的时间转化为JSON字符串格式
// 返回转化后的JSON字符串和错误信息
func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Exist {
		if t.Val == nil {
			return json.Marshal(nil)
		}
		tt := *t.Val
		tStr := tt.Format(TimeFormat)
		return json.Marshal(tStr)
	} else {
		return json.Marshal(nil)
	}
}
func (i *NullTime) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// 实现gin框架的参数绑定接口
func (i *NullTime) UnmarshalParam(param string) error {
	return i.Scan(param)
}

func (t *NullTime) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}
	if date == "" {
		*t = NullTime{
			Val:   nil,
			Exist: true,
		}
		return nil
	}
	tt, err := time.ParseInLocation(TimeFormat, date, time.Local)
	if err != nil {
		return err
	}
	*t = NullTime{
		Val:   &tt,
		Exist: true,
	}
	return nil
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
func (i *NullTime) SetValue(value time.Time) {
	i.Val = &value
	i.Exist = true
}
func (i *NullTime) SetNull() {
	i.Val = nil
	i.Exist = true
}

func (i *NullTime) GetValue() *time.Time {
	return i.Val
}
func (i *NullTime) ValueOr(v time.Time) time.Time {
	if i.Val == nil {
		return v
	}
	return *i.Val
}
func (i *NullTime) ValueOrZero() time.Time {
	if i.Val == nil {
		return time.Time{}
	}
	return *i.Val
}

// go to json时omitempty标签是否忽略该字段
func (i NullTime) IsZero() bool {
	return !i.Exist
}
func (i *NullTime) IsExists() bool {
	return i.Exist
}

// IsExistsAndNull 存在且为null
func (i *NullTime) IsExistsAndNull() bool {
	return i.Exist && i.Val == nil
}
