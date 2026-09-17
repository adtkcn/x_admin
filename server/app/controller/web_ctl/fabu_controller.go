package web_ctl

import (
	"strconv"

	"x_admin/app/model/fabu_model"
	"x_admin/app/service/fabu_service"
	"x_admin/core"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// FabuController 公开下载/安装相关接口
type FabuController struct{}

// Plist 返回 iOS 安装 manifest
func (h FabuController) Plist(c *gin.Context) {
	appId := c.Param("appId")
	versionId := c.Param("versionId")
	app, version, err := fabu_service.VersionService.PlistData(appId, versionId)
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	xmlStr, err := fabu_service.PlistService.Build(app, version)
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	c.Data(200, "application/xml", []byte(xmlStr))
}

// Download 按短链下载(自增计数后跳转)
func (h FabuController) Download(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	app, err := fabu_service.AppService.FindByShortUrl(shortUrl)
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	url, err := fabu_service.VersionService.DownloadInfo(app.ID, "")
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	fabu_service.VersionService.Count(app.ID, "")
	c.Redirect(302, url)
}

// Count 按应用/版本计数后跳转下载
func (h FabuController) Count(c *gin.Context) {
	appId := c.Param("appId")
	versionId := c.Param("versionId")
	url, err := fabu_service.VersionService.DownloadInfo(appId, versionId)
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	fabu_service.VersionService.Count(appId, versionId)
	c.Redirect(302, url)
}

// CheckUpdate 版本检查更新
func (h FabuController) CheckUpdate(c *gin.Context) {
	bundleId := c.Query("bundle_id")
	clientCode, _ := strconv.Atoi(c.Query("version_code"))

	var app fabu_model.FabuApp
	if err := fabu_service.AppService.FindByBundleId(bundleId, &app); err != nil {
		response.JSON(c, nil, err)
		return
	}
	if app.CurrentVersionId == "" {
		response.JSON(c, gin.H{"update": false}, nil)
		return
	}
	var version fabu_model.FabuAppVersion
	if err := core.GetDB().Where("id = ?", app.CurrentVersionId).First(&version).Error; err != nil {
		response.JSON(c, gin.H{"update": false}, nil)
		return
	}
	if version.VersionCode <= clientCode {
		response.JSON(c, gin.H{"update": false}, nil)
		return
	}
	response.JSON(c, gin.H{
		"update":       true,
		"version":      version.Version,
		"version_code": version.VersionCode,
		"download_url": version.DownloadUrl,
		"install_url":  version.InstallUrl,
		"update_mode":  version.UpdateMode,
	}, nil)
}
