package core

import (
	"database/sql/driver"
	"encoding/json"
	"x_admin/util/convert_util"
)

// 支持前端传递null，int，string类型和不传值
// 前端传1，“1”都可以，都转换为int64类型: NullString{Int: "1", Exist: true}
// 前端null值: NullString{Int: nil, Exist: true}
// 前端没传值: NullString{Int: nil, Exist: false}
type NullString struct {
	Val   *string //解析行为默认""而不是nil
	Exist bool
}

func DecodeString(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		var s string
		return NullString{Val: &s, Exist: true}, nil
	case NullString:
		return v, nil
	default:
		result := convert_util.ToString(v)
		return NullString{Val: &result, Exist: true}, nil
	}
}

// gorm实现Scanner,支持string, nil类型
func (i *NullString) Scan(value any) error {

	switch v := value.(type) {
	case nil:
		var s string
		i.Val = &s
		i.Exist = true
		return nil
	case string:
		i.Val, i.Exist = &v, true
		return nil

	default:
		result := convert_util.ToString(v)
		i.Val, i.Exist = &result, true
		return nil
		// return fmt.Errorf("类型转换失败，期望string类型，实际类型为%T，值为%v", value, value)
	}
}

// gorm实现 Valuer
func (i NullString) Value() (driver.Value, error) {
	if !i.Exist {
		return nil, nil
	}
	v := i.Val
	if v == nil {
		return nil, nil
	}
	return *v, nil
}

// 实现fmt.Stringer接口
func (i NullString) String() string {
	if i.Val != nil {
		return *i.Val
	} else {
		return ""
	}
}

// 实现json序列化接口
func (i NullString) MarshalJSON() ([]byte, error) {
	if i.Exist {
		return json.Marshal(i.Val)
	} else {
		return json.Marshal(nil)
	}
}

func (i *NullString) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// 实现gin框架的参数绑定接口
func (i *NullString) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// 实现json反序列化接口
func (i *NullString) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case nil:
		var s string
		i.Val = &s
		i.Exist = true

		return nil
	default:
		result := convert_util.ToString(v)
		i.Val = &result
		i.Exist = true

		return nil
	}

}

func (i *NullString) SetValue(value string) {
	i.Val = &value
	i.Exist = true
}
func (i *NullString) SetNull() {
	i.Val = nil
	i.Exist = true
}

func (i *NullString) GetValue() *string {
	return i.Val
}
func (i *NullString) ValueOr(v string) string {
	if i.Val == nil {
		return v
	}
	return *i.Val
}
func (i *NullString) ValueOrZero() string {
	if i.Val == nil {
		return ""
	}
	return *i.Val
}

// go to json时omitempty标签是否忽略该字段
func (i NullString) IsZero() bool {
	return !i.Exist
}

func (i *NullString) IsExists() bool {
	return i.Exist
}

// IsExistsAndNotNull 是否存在且不为空
func (i *NullString) IsExistsAndNotNull() bool {
	return i.Exist && i.Val != nil
}

// IsExistsAndNull 存在且为null
func (i *NullString) IsExistsAndNull() bool {
	return i.Exist && i.Val == nil
}
