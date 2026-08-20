package util

import (
	"encoding/json"
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

	return RedisUtil.Set(c.Name+":"+field, str, 3600)
}

// 获取缓存
func (c CacheUtil) GetCache(field string, obj any) error {

	str := RedisUtil.Get(c.Name + ":" + field)
	if str == "" {
		return errors.New("获取缓存失败")
	}
	if err := json.Unmarshal([]byte(str), obj); err != nil {
		return errors.New("解析缓存失败")
	}
	return nil
}

// 删除缓存-支持批量删除
func (c CacheUtil) RemoveCache(fields ...string) bool {
	if len(fields) == 0 {
		return true
	}
	var keys []string
	for _, field := range fields {
		keys = append(keys, c.Name+":"+field)
	}

	return RedisUtil.Del(keys...)
}
