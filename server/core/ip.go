package core

import (
	"fmt"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var IpUtils = new(ipUtils)

type ipUtils struct {
	Searcher *xdb.Searcher
}

func init() {
	var dbPath = "resources/ip/ip2region.xdb"
	// 创建完全基于内存的查询对象。
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return
	}
	// 并发使用，用整个 xdb 缓存创建的 searcher 对象可以安全用于并发。
	searcher, err := xdb.NewWithBuffer(cBuff)

	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}
	fmt.Printf("创建完全基于内存的查询对象。")
	IpUtils.Searcher = searcher
	// defer searcher.Close()

	// region, err := searcher.SearchByStr(ip)
	// if err != nil {
	// 	fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
	// 	return
	// }
	// fmt.Printf("{region: %+v}\n", region)
}

type Region struct {
	Country  string
	Province string
	City     string
	Operator string
}

func (ipUtils *ipUtils) Parse(ip string) Region {
	// var dbPath = "resources/ip/ip2region.xdb"
	// // 创建完全基于内存的查询对象。
	// cBuff, err := xdb.LoadContentFromFile(dbPath)
	// searcher, err := xdb.NewWithBuffer(cBuff)

	// if err != nil {
	// 	fmt.Printf("failed to create searcher: %s\n", err.Error())
	// 	return
	// }
	// defer searcher.Close()
	if ip == "" {
		fmt.Println("输入ip为空")
		return Region{}
	}
	region, err := ipUtils.Searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("解析ip(%s)错误: %s\n", ip, err)
		return Region{}
	}
	//   中国|0|四川省|成都市|电信
	parts := strings.Split(region, "|")
	if len(parts) != 5 {
		fmt.Println("解析ip返回错误")
		return Region{}
	}
	country := parts[0]
	province := parts[2]
	city := parts[3]
	operator := parts[4]
	regionInfo := Region{Country: country, Province: province, City: city, Operator: operator}
	return regionInfo
}
