package core

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"x_admin/util/convert_util"
)

// Float类型别名，支持前端传递null，float64, string类型
// 忽略前端null值，接收""值时，返回0
type NullFloat struct {
	Float *float64
	Valid bool
}

func DecodeFloat(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		return NullFloat{Float: nil, Valid: false}, nil
	case NullFloat:
		return v, nil
	default:
		result, err := convert_util.ToFloat64(value)
		if err != nil {
			return NullFloat{Float: nil, Valid: false}, err
		}
		return NullFloat{Float: &result, Valid: true}, nil
	}
}

// gorm实现Scanner
func (f *NullFloat) Scan(value interface{}) error {
	f.Valid = false

	result, err := convert_util.ToFloat64(value)
	if err != nil {
		return err
	}
	f.Float, f.Valid = &result, true
	return nil
}

// gorm实现 Valuer
func (f NullFloat) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	v := f.Float
	if v == nil {
		return nil, nil
	}
	return *v, nil
}

func (f NullFloat) String() string {
	if f.Valid {
		return strconv.FormatFloat(*f.Float, 'f', -1, 64)
	} else {
		return ""
	}
}

// 实现json序列化接口
func (f NullFloat) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float)
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
	case int64:
		f64 := float64(v)
		f.Float = &f64
		f.Valid = true
		return nil
	case float64:
		f.Float = &v
		f.Valid = true
		return nil
	case string:
		if v == "" {
			f.Float = nil
			f.Valid = true
			return nil
		}
		num, err := strconv.ParseFloat(v, 64)
		if err == nil {
			f.Float = &num
			f.Valid = true
		} else {
			f.Valid = false
		}
		return err
	case nil:
		f.Valid = false
	default:
		f.Valid = false
	}

	return nil
}
