package util

import (
	"x_admin/core"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
)

var ConvertUtil = convertUtil{}

// convertUtil 转换工具
type convertUtil struct{}

// StructsToMaps 将结构体转换成Map列表
func (c convertUtil) StructsToMaps(from interface{}) (data []map[string]interface{}) {
	var objList []interface{}
	err := copier.Copy(&objList, from)
	if err != nil {
		core.Logger.Errorf("convertUtil.StructsToMaps err: err=[%+v]", err)
		return nil
	}
	for _, v := range objList {
		// data = append(data, structs.Map(v))
		data = append(data, c.StructToMap(v))
	}
	return data
}

// StructToMap 结构体转换成map
func (c convertUtil) StructToMap(from interface{}) map[string]interface{} {
	// var m = map[string]interface{}{}
	// mapstructure.Decode(from, &m) //mapstructure

	m, _ := convertor.StructToMap(from) // 需要tag:json

	return m
}

// MapToStruct 将map弱类型转换成结构体
func (c convertUtil) MapToStruct(from map[string]interface{}, to interface{}) (err error) {
	err = mapstructure.WeakDecode(from, to) // 需要tag:mapstructure
	return err
}

// StructToStruct 将结构体from弱类型转换成结构体to,需要tag:json,mapstructure
func (c convertUtil) StructToStruct(from interface{}, to interface{}) (err error) {
	m := c.StructToMap(from)
	err = c.MapToStruct(m, to)

	return err
}
