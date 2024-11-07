package core

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"x_admin/util/convert_util"
)

// int类型别名，支持前端传递null，int，string类型
// 忽略前端null值，接收""值时，返回0
type NullInt struct {
	Int   *int64
	Valid bool
}

//	func EncodeInt(value any) any {
//		switch v := value.(type) {
//		case NullInt:
//			if v.Valid {
//				return *v.Int
//			} else {
//				return nil
//			}
//		case map[string]any:
//			if v["Int"] != nil {
//				val := v["Int"]
//				switch i := val.(type) {
//				case *int:
//					return *i
//				case *int64:
//					return *i
//				case *string:
//					return *i
//				default:
//					return nil
//				}
//				// return val
//			}
//		}
//		return nil
//	}
func DecodeInt(value any) (any, error) {
	switch v := value.(type) {
	// case int:
	// 	i := int64(v)
	// 	return NullInt{Int: &i, Valid: true}, nil
	// case int64:
	// 	return NullInt{Int: &v, Valid: true}, nil
	// case string:
	// 	if v == "" {
	// 		return NullInt{Int: nil, Valid: false}, nil
	// 	}
	// 	i, err := strconv.ParseInt(v, 10, 64)
	// 	return NullInt{Int: &i, Valid: true}, err
	case nil:
		return NullInt{Int: nil, Valid: false}, nil
	case NullInt:
		return v, nil
	default:
		result, err := convert_util.ToInt64(value)
		if err != nil {
			return NullInt{Int: nil, Valid: false}, err
		}
		return NullInt{Int: &result, Valid: true}, nil
	}

}

// gorm实现Scanner
func (i *NullInt) Scan(value interface{}) error {
	i.Valid = false
	if value == nil {
		return nil
	}
	v := value.(int64)
	i.Int, i.Valid = &v, true
	return nil
}

// gorm实现 Valuer
func (i NullInt) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	v := i.Int
	if v == nil {
		return nil, nil
	}
	return *v, nil
}
func (i NullInt) String() string {
	if i.Valid {
		return strconv.FormatInt(*i.Int, 10)
	} else {
		return ""
	}
}

// 实现json序列化接口
func (i NullInt) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int)
	} else {
		return json.Marshal(nil)
	}
}

// 实现json反序列化接口
func (i *NullInt) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case int64:
		i.Int = &v
		i.Valid = true
		return nil
	case float64:
		i64 := int64(v)
		i.Int = &i64
		i.Valid = true
		return nil
	case string:
		if v == "" {
			i.Int = nil
			i.Valid = true
			return nil
		}
		num, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			i.Int = &num
			i.Valid = true
		} else {
			i.Valid = false
		}
		return err
	case nil:
		i.Valid = false
	default:
		i.Valid = false
	}

	return nil
}
