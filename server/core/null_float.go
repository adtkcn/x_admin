package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"x_admin/util/convert_util"
)

// 支持前端传递null，int，float，string类型和不传值
// 前端传1，“1”都可以，都转换为float64类型: NullFloat{Float: 1.0, Exist: true}
// 前端null值: NullFloat{Float: nil, Exist: true}
// 前端没传值: NullFloat{Float: nil, Exist: false}
type NullFloat struct {
	Val   *float64
	Exist bool // 是否有值
}

func DecodeFloat(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		return NullFloat{Val: nil, Exist: false}, nil
	case NullFloat:
		return v, nil
	default:
		result, err := convert_util.ToFloat64(value)
		if err != nil {
			return NullFloat{Val: nil, Exist: false}, err
		}
		return NullFloat{Val: &result, Exist: true}, nil
	}
}

// gorm实现Scanner
func (f *NullFloat) Scan(value interface{}) error {

	result, err := convert_util.ToFloat64(value)
	if err != nil {
		return err
	}
	f.Val, f.Exist = &result, true
	return nil
}

// gorm实现 Valuer
func (f NullFloat) Value() (driver.Value, error) {
	if !f.Exist {
		return nil, nil
	}
	v := f.Val
	if v == nil {
		return nil, nil
	}
	return *v, nil
}

func (f NullFloat) String() string {
	if f.Exist {
		return strconv.FormatFloat(*f.Val, 'f', -1, 64)
	} else {
		return ""
	}
}

func (i *NullFloat) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// 实现gin框架的参数绑定接口
func (i *NullFloat) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// 实现json序列化接口
func (f NullFloat) MarshalJSON() ([]byte, error) {
	if f.Exist {
		return json.Marshal(f.Val)
	} else {
		return json.Marshal(nil)
	}
}

// 实现json反序列化接口
func (f *NullFloat) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case nil:
		f.Exist = true
		return nil
	case int64:
		f64 := float64(v)
		f.Val = &f64
		f.Exist = true
		return nil
	case float64:
		f.Val = &v
		f.Exist = true
		return nil
	case string:
		if v == "" {
			f.Val = nil
			f.Exist = true
			return nil
		}
		num, err := strconv.ParseFloat(v, 64)
		if err == nil {
			f.Val = &num
			f.Exist = true
		} else {
			f.Exist = false
		}
		return err

	default:
		return fmt.Errorf("不能将类型 %T 转换为 float64, 值为 %v", v, v)
	}
}

func (i *NullFloat) SetValue(value float64) {
	i.Val = &value
	i.Exist = true
}
func (i *NullFloat) SetNull() {
	i.Val = nil
	i.Exist = true
}
func (i *NullFloat) IsExists() bool {
	return i.Exist
}
func (i *NullFloat) GetValue() *float64 {
	return i.Val
}
