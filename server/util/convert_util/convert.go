package convert_util

import (
	"encoding/json/v2"
	"fmt"
	"strconv"

	"reflect"

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
func StructsToMaps[T any](from []T) (data []map[string]any, err error) {
	for _, v := range from {
		// 忽略错误
		m, err := StructToMap(v)
		if err != nil {
			return nil, err
		}
		data = append(data, m)
	}
	return data, nil
}

// ShallowStructToMap 将结构体转换成map,浅转换
func ShallowStructToMap(from any) map[string]any {
	m := make(map[string]any)
	v := reflect.ValueOf(from)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		m[field.Name] = value
	}

	return m
}

// ShallowStructsToMaps 将结构体列表转换成Map列表,浅转换
func ShallowStructsToMaps(from any) (data []map[string]any) {
	var objList []any
	err := copier.Copy(&objList, from)
	if err != nil {
		return nil
	}
	for _, v := range objList {
		data = append(data, ShallowStructToMap(v))
	}
	return data
}

// MapToStruct 将map类型转换成结构体
func MapToStruct(from any, to any) (err error) {
	// err = mapstructure.WeakDecode(from, to) // 需要tag:mapstructure

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
