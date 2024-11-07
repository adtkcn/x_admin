package convert

import (
	"fmt"
	"strconv"

	"reflect"

	"github.com/duke-git/lancet/v2/convertor"

	"github.com/fatih/structs"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
)

var ConvertUtil = convertUtil{}

// convertUtil 转换工具
type convertUtil struct{}

func (c convertUtil) ToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64, float32:
		return strconv.ParseFloat(fmt.Sprintf("%f", v), 64)
	default:
		return convertor.ToFloat(value)
	}

	// return 0, nil
}

// StructToMap 结构体转换成map,深度转换
func (c convertUtil) StructToMap(from interface{}) map[string]interface{} {
	// var m = map[string]interface{}{}
	// mapstructure.Decode(from, &m) //深度转换所有结构体

	m := structs.Map(from) // 需要tag:structs，深度转换
	return m
}

// StructsToMaps 将结构体转换成Map列表
func (c convertUtil) StructsToMaps(from interface{}) (data []map[string]interface{}) {
	var objList []interface{}
	err := copier.Copy(&objList, from)
	if err != nil {
		// core.Logger.Errorf("convertUtil.StructsToMaps err: err=[%+v]", err)
		return nil
	}
	for _, v := range objList {
		data = append(data, c.StructToMap(v))
	}
	return data
}

// ShallowStructToMap 将结构体转换成map,浅转换
func (c convertUtil) ShallowStructToMap(from interface{}) map[string]interface{} {
	m := make(map[string]interface{})
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
func (c convertUtil) ShallowStructsToMaps(from interface{}) (data []map[string]interface{}) {
	var objList []interface{}
	err := copier.Copy(&objList, from)
	if err != nil {
		// core.Logger.Errorf("convertUtil.StructsToMaps err: err=[%+v]", err)
		return nil
	}
	for _, v := range objList {
		data = append(data, c.ShallowStructToMap(v))
	}
	return data
}

// MapToStruct 将map弱类型转换成结构体
func (c convertUtil) MapToStruct(from interface{}, to interface{}) (err error) {
	err = mapstructure.WeakDecode(from, to) // 需要tag:mapstructure
	return err
}

// StructToStruct 将结构体from弱类型转换成结构体to
func (c convertUtil) StructToStruct(from interface{}, to interface{}) (err error) {
	m := c.StructToMap(from)
	err = c.MapToStruct(m, to)

	return err
}

func (c convertUtil) Copy(toValue interface{}, fromValue interface{}) interface{} {
	if err := copier.Copy(toValue, fromValue); err != nil {
		// core.Logger.Errorf("Copy err: err=[%+v]", err)
		panic("SystemError")
	}
	return toValue
}
