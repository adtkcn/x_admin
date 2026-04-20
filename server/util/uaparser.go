package util

import (
	"strings"

	"github.com/ua-parser/uap-go/uaparser"
)

type uaUtils struct {
	uaParser *uaparser.Parser
}
type UaInfo struct {
	OsName         string // 操作系统名称:Windows
	OsVersion      string // 操作系统版本:10.0
	BrowserName    string // 浏览器名称:Edge
	BrowserVersion string // 浏览器版本:91.0.4472.124
	DeviceName     string // 设备名称: 设备类型/系列+型号
	DeviceBrand    string // 设备品牌:Apple
}

func (u uaUtils) Parse(ua string) UaInfo {
	result := u.uaParser.Parse(ua)
	osName := result.Os.Family
	osVersion := result.Os.Major + "." + result.Os.Minor + "." + result.Os.Patch
	browserName := result.UserAgent.Family
	browserVersion := result.UserAgent.Major + "." + result.UserAgent.Minor + "." + result.UserAgent.Patch
	deviceName := strings.Join([]string{result.Device.Family, result.Device.Model}, " ")
	deviceBrand := result.Device.Brand
	return UaInfo{
		OsName:         osName,
		OsVersion:      osVersion,
		BrowserName:    browserName,
		BrowserVersion: browserVersion,
		DeviceName:     deviceName,
		DeviceBrand:    deviceBrand,
	}
}

var UAUtils = uaUtils{
	uaParser: uaparser.NewFromSaved(),
}
