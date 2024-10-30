package util

import (
	"fmt"
	"net"
	"strings"
	"x_admin/core"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var IpUtil = initIpUtil()

// serverUtil IP工具类
type ipUtil struct {
	Searcher *xdb.Searcher
}

// GetHostIp 获取本地主机名
func (su ipUtil) GetHostIp() (ip string) {
	conn, err := net.Dial("udp", "114.114.114.114:80")
	if err != nil {
		core.Logger.Errorf("GetHostIp Dial err: err=[%+v]", err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

type Region struct {
	Country  string
	Province string
	City     string
	Operator string
}

func (ipUtils *ipUtil) Parse(ip string) Region {

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
	fmt.Println(regionInfo)
	return regionInfo
}
func initIpUtil() *ipUtil {
	ip_util := ipUtil{}

	var dbPath = "resources/ip/ip2region.xdb"
	// 创建完全基于内存的查询对象。
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return &ip_util
	}
	// 并发使用，用整个 xdb 缓存创建的 searcher 对象可以安全用于并发。
	searcher, err := xdb.NewWithBuffer(cBuff)

	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return &ip_util
	}
	fmt.Printf("创建完全基于内存的查询对象。")
	ip_util.Searcher = searcher
	return &ip_util
}
