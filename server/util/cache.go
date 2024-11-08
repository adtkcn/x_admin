package util

import (
	"errors"
	"strconv"
)

// var CacheUtil = toolsUtil{}

type CacheUtil struct {
	Name string
}

// 设置缓存
func (c CacheUtil) SetCache(key interface{}, obj interface{}) bool {
	str, e := ToolsUtil.ObjToJson(obj)
	if e != nil {
		return false
	}
	var cacheKey string
	switch k := key.(type) {
	case int:
		cacheKey = strconv.Itoa(k)
	case string:
		cacheKey = k
	default:
		return false
	}
	return RedisUtil.HSet(c.Name, cacheKey, str, 3600)
}

// 获取缓存
func (c CacheUtil) GetCache(key interface{}, obj interface{}) error {
	var cacheKey string
	switch k := key.(type) {
	case int:
		cacheKey = strconv.Itoa(k)
	case string:
		cacheKey = k
	default:
		return errors.New("缓存key无效")
	}
	str := RedisUtil.HGet(c.Name, cacheKey)
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
func (c CacheUtil) RemoveCache(key interface{}) bool {
	var cacheKey []string
	switch k := key.(type) {
	case int:
		cacheKey = append(cacheKey, strconv.Itoa(k))
	case string:
		cacheKey = append(cacheKey, k)
	//  判断是slice
	case []int:
		for _, v := range k {
			cacheKey = append(cacheKey, strconv.Itoa(v))
		}
	case []string:
		cacheKey = k
	default:
		return false
	}
	return RedisUtil.HDel(c.Name, cacheKey...)
}
