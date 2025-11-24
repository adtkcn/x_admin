package core

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"x_admin/util/convert_util"
)

// 支持前端传递null，int，string类型和不传值
// 前端传1，“1”都可以，都转换为int64类型: NullInt{Int: 1, Valid: true}
// 前端null值: NullInt{Int: nil, Valid: true}
// 前端没传值: NullInt{Int: nil, Valid: false}

type NullInt struct {
	Val   *int64 // 整数或者null
	Valid bool   // 是否有值
}

func DecodeInt(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		return NullInt{Val: nil, Valid: false}, nil
	case NullInt:
		return v, nil
	default:
		result, err := convert_util.ToInt64(value)
		if err != nil {
			return NullInt{Val: nil, Valid: false}, err
		}
		return NullInt{Val: &result, Valid: true}, nil
	}

}

// gorm实现Scanner
func (i *NullInt) Scan(value interface{}) error {
	// 判断int64、string类型
	switch v := value.(type) {
	case nil:
		i.Valid = true
		return nil
	case int64:
		i.Val, i.Valid = &v, true
		return nil
	case string:
		num, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			i.Val = &num
			i.Valid = true
		} else {
			i.Valid = false
		}
		return err
	default:
		return fmt.Errorf("不能将类型 %T 转换为 int64", v)
	}
}

// gorm实现 Valuer
func (i NullInt) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	v := i.Val
	if v == nil {
		return nil, nil
	}
	return *v, nil
}
func (i NullInt) String() string {
	if i.Valid {
		return strconv.FormatInt(*i.Val, 10)
	} else {
		return ""
	}
}

// 实现json序列化接口
func (i NullInt) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Val)
	} else {
		return json.Marshal(nil)
	}
}
func (i *NullInt) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// 实现gin框架的参数绑定接口
func (i *NullInt) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// 实现json反序列化接口,支持 int64, string，null类型，对于float64类型，判断转换前后是否相等，防止精度丢失
func (i *NullInt) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case nil:
		i.Valid = true
	case int64:
		i.Val = &v
		i.Valid = true
		return nil
	case float64:
		i64 := int64(v)
		// 判断转换前后是否相等，防止精度丢失
		if float64(i64) != v {
			i.Valid = false
			return errors.New("int64转换失败，" + fmt.Sprintf("%f", v) + "精度丢失")
		}
		i.Val = &i64
		i.Valid = true
		return nil
	case string:
		if v == "" {
			i.Val = nil
			i.Valid = true
			return nil
		}
		num, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			i.Val = &num
			i.Valid = true
		} else {
			i.Valid = false
		}
		return err

	default:
		return fmt.Errorf("不能将类型 %T 转换为 int64", v)
	}

	return nil
}
func (i *NullInt) SetValue(value int64) {
	i.Val = &value
	i.Valid = true
}
func (i *NullInt) SetNull() {
	i.Val = nil
	i.Valid = true
}
func (i *NullInt) IsValid() bool {
	return i.Valid
}
func (i *NullInt) GetValue() *int64 {
	return i.Val
}
