package convert_util

import (
	"fmt"
	"strconv"

	"reflect"

	"github.com/duke-git/lancet/v2/convertor"

	"github.com/fatih/structs"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
)

func ToFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case float32:
		return strconv.ParseFloat(fmt.Sprintf("%f", v), 64)
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

// StructToMap 结构体转换成map,深度转换
func StructToMap(from any) map[string]any {
	// var m = map[string]any{}
	// mapstructure.Decode(from, &m) //深度转换所有结构体

	m := structs.Map(from) // 需要tag:structs，深度转换
	return m
}

// StructsToMaps 将结构体转换成Map列表
func StructsToMaps(from any) (data []map[string]any) {
	var objList []any
	err := copier.Copy(&objList, from)
	if err != nil {
		return nil
	}
	for _, v := range objList {
		data = append(data, StructToMap(v))
	}
	return data
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

// MapToStruct 将map弱类型转换成结构体
func MapToStruct(from any, to any) (err error) {
	err = mapstructure.WeakDecode(from, to) // 需要tag:mapstructure
	return err
}

// StructToStruct 将结构体from弱类型转换成结构体to
// func StructToStruct(from any, to any) (err error) {
// 	m := StructToMap(from)
// 	err = MapToStruct(m, to)

// 	return err
// }

func Copy(toValue any, fromValue any) any {
	if err := copier.Copy(toValue, fromValue); err != nil {
		// core.Logger.Errorf("Copy err: err=[%+v]", err)
		panic("SystemError")
	}
	return toValue
}
