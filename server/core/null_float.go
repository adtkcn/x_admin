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
	// case float64:
	// 	f := v
	// 	return NullFloat{Float: &f, Valid: true}, nil
	// case float32:
	// 	f := float64(v)
	// 	return NullFloat{Float: &f, Valid: true}, nil
	// case int:
	// 	f := float64(v)
	// 	return NullFloat{Float: &f, Valid: true}, nil
	// case int64:
	// 	f := float64(v)
	// 	return NullFloat{Float: &f, Valid: true}, nil
	// case string:
	// 	if v == "" {
	// 		return NullFloat{Float: nil, Valid: false}, nil
	// 	}
	// 	f, err := strconv.ParseFloat(v, 64)
	// 	return NullFloat{Float: &f, Valid: true}, err
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
	// switch v := value.(type) {
	// case float64:
	// 	f.Float, f.Valid = &v, true
	// case float32:
	// 	// 直接用float64(float32(v)), 会丢失精度
	// 	val, _ := strconv.ParseFloat(fmt.Sprintf("%f", v), 64)
	// 	f.Float, f.Valid = &val, true
	// 	// 匹配所有int
	// case uint, uint8, uint16, uint32, uint64, int8, int16, int, int32, int64:
	// 	val := float64(v)
	// 	f.Float, f.Valid = &val, true
	// case string:
	// 	if v == "" {
	// 		f.Float, f.Valid = nil, false
	// 		return nil
	// 	}
	// 	val, err := strconv.ParseFloat(v, 64)
	// 	if err != nil {
	// 		f.Float, f.Valid = nil, false
	// 		return err
	// 	}
	// 	f.Float, f.Valid = &val, true
	// 	return err
	// case nil:
	// 	f.Float, f.Valid = nil, false
	// 	return nil
	// }

	// return nil
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
