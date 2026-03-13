package util

import (
	"errors"
)

// var CacheUtil = toolsUtil{}

type CacheUtil struct {
	Name string
}

// 设置缓存
func (c CacheUtil) SetCache(field string, obj any) bool {
	str, e := ToolsUtil.ObjToJson(obj)
	if e != nil {
		return false
	}

	return RedisUtil.HSet(c.Name, field, str, 3600)
}

// 获取缓存
func (c CacheUtil) GetCache(field string, obj any) error {

	str := RedisUtil.HGet(c.Name, field)
	if str == "" {
		return errors.New("获取缓存失败")
	}
	err := ToolsUtil.JsonToObj(str, &obj)

	if err != nil {
		return errors.New("解析缓存失败")
	}
	return nil
}

// 删除缓存-支持批量删除
func (c CacheUtil) RemoveCache(fields ...string) bool {

	return RedisUtil.HDel(c.Name, fields...)
}
