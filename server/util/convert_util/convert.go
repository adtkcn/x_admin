package convert_util

import (
	"encoding/json/v2"
	"fmt"
	"strconv"

	"github.com/duke-git/lancet/v2/convertor"

	"github.com/jinzhu/copier"
)

func ToFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case float32:
		return strconv.ParseFloat(fmt.Sprintf("%f", v), 64)
	case []uint8:
		return strconv.ParseFloat(string(v), 64)
	default:
		return convertor.ToFloat(value)
	}
}
func ToInt64(value any) (int64, error) {
	return convertor.ToInt(value)
}
func ToString(value any) string {
	return convertor.ToString(value)
}

// StructToMap 使用 JSON 中转，确保调用 MarshalJSON
func StructToMap(v any) (map[string]any, error) {
	// 第一步：序列化为 JSON
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	// 第二步：反序列化为 map
	var result map[string]any
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// StructsToMaps 将结构体转换成Map列表
func StructsToMaps(from any) ([]map[string]any, error) {
	// 第一步：序列化为 JSON
	data, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	// 第二步：反序列化为 map
	var result []map[string]any
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AnyToAny 将map类型转换成结构体
func AnyToAny(from any, to any) (err error) {
	jsonData, err := json.Marshal(from)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, to)
	return err
}

func Copy(toValue any, fromValue any) any {
	if err := copier.Copy(toValue, fromValue); err != nil {
		// core.Logger.Errorf("Copy err: err=[%+v]", err)
		panic(err)
	}
	return toValue
}
