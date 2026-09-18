package web_ctl

import (
	"strconv"

	"x_admin/app/service/fabu_service"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// FabuController 公开下载/安装相关接口
type FabuController struct{}

// requestBaseURL 拼接当前站点绝对前缀（scheme://host），供 iOS plist 使用
func requestBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host
}

// Plist 返回 iOS 安装 manifest
func (h FabuController) Plist(c *gin.Context) {
	appId := c.Param("appId")
	versionId := c.Param("versionId")
	app, version, err := fabu_service.VersionService.PlistData(appId, versionId)
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	xmlStr, err := fabu_service.PlistService.Build(app, version, requestBaseURL(c))
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	c.Data(200, "application/xml", []byte(xmlStr))
}

// AppInfo 按短链返回下载页展示信息（公开，无需登录）
func (h FabuController) AppInfo(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	res, err := fabu_service.VersionService.DownloadPageInfo(shortUrl)
	response.JSON(c, res, err)
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

// CheckUpdate 版本检查更新：优先全量包，其次当前版本热更包，判定逻辑见 VersionService.CheckUpdate
func (h FabuController) CheckUpdate(c *gin.Context) {
	clientCode, _ := strconv.Atoi(c.Query("version_code"))
	wgtCode, _ := strconv.Atoi(c.Query("wgt_code"))
	res, err := fabu_service.VersionService.CheckUpdate(c.Query("bundle_id"), c.Query("platform"), clientCode, wgtCode)
	response.JSON(c, res, err)
}
