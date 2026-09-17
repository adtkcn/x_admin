package fabu_service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"uuid"

	"x_admin/util"

	"x_admin/app/model/fabu_model"
	"x_admin/app/schema/fabu_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"github.com/phinexdaz/ipapk"
	"gorm.io/gorm"
)

var VersionService = NewFabuVersionService()

func NewFabuVersionService() *fabuVersionService {
	return &fabuVersionService{db: core.GetDB()}
}

type fabuVersionService struct {
	db *gorm.DB
}

func (s fabuVersionService) List(req fabu_schema.FabuVersionListReq) (res map[string]any, e error) {
	var versions []fabu_model.FabuAppVersion
	var total int64
	chain := s.db.Model(&fabu_model.FabuAppVersion{}).Where("app_id = ?", req.AppId)
	if e = response.CheckErr(chain.Count(&total).Error, "统计版本失败"); e != nil {
		return
	}
	pageNo, pageSize := req.PageNo, req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if e = response.CheckErr(chain.Order("create_time desc").Offset((pageNo-1)*pageSize).Limit(pageSize).Find(&versions).Error, "查询版本失败"); e != nil {
		return
	}
	var list []fabu_schema.FabuVersionResp
	convert_util.Copy(&list, versions)
	res = map[string]any{"count": total, "lists": list}
	return
}

// Upload 上传安装包：解析(ipapk) -> 重编码图标 -> 落盘 -> 建/追加应用与版本
func (s fabuVersionService) Upload(file *multipart.FileHeader) (res fabu_schema.FabuUploadResp, e error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	platform := "android"
	if ext == ".ipa" {
		platform = "ios"
	}

	tmpDir := filepath.Join(config.FileConfig.UploadDirectory, "fabu", "tmp")
	if e = response.CheckErr(os.MkdirAll(tmpDir, 0755), "创建临时目录失败"); e != nil {
		return
	}
	tmpPath := filepath.Join(tmpDir, uuid.NewV7().String()+ext)
	if e = response.CheckErr(saveUploadedFile(file, tmpPath), "保存上传文件失败"); e != nil {
		return
	}

	info, err := ipapk.NewAppParser(tmpPath)
	if err != nil {
		os.Remove(tmpPath)
		e = response.AssertArgumentError.SetMessage("解析安装包失败: " + err.Error())
		return
	}

	// 重编码图标（iOS 优化 PNG 已由库解码为 image.Image）
	iconURL := ""
	if info.Icon != nil {
		var buf bytes.Buffer
		if encErr := png.Encode(&buf, info.Icon); encErr == nil {
			iconUUID := uuid.NewV7().String()
			iconPath := filepath.Join(config.FileConfig.UploadDirectory, "fabu", "icon", iconUUID+".png")
			if mkErr := os.MkdirAll(filepath.Dir(iconPath), 0755); mkErr == nil {
				if wErr := os.WriteFile(iconPath, buf.Bytes(), 0644); wErr == nil {
					iconURL = "/api/fabu/static/fabu/icon/" + iconUUID + ".png"
				}
			}
		}
	}

	// 落盘包体
	pkgUUID := uuid.NewV7().String()
	pkgDir := filepath.Join(config.FileConfig.UploadDirectory, "fabu", "pkg", platform)
	if e = response.CheckErr(os.MkdirAll(pkgDir, 0755), "创建包目录失败"); e != nil {
		os.Remove(tmpPath)
		return
	}
	pkgPath := filepath.Join(pkgDir, pkgUUID+ext)
	if e = response.CheckErr(os.Rename(tmpPath, pkgPath), "移动包体失败"); e != nil {
		return
	}

	md5Str, _ := fileMd5(pkgPath)
	versionCode, _ := strconv.Atoi(info.Build)
	downloadUrl := "/api/fabu/static/fabu/pkg/" + platform + "/" + pkgUUID + ext

	// 按 BundleId + Platform 查重（iOS/Android 通常共用包名，需区分），无则新建应用
	var app fabu_model.FabuApp
	var isNewApp bool
	err = s.db.Where("bundle_id = ? AND platform = ?", info.BundleId, platform).First(&app).Error
	if err == gorm.ErrRecordNotFound {
		app = fabu_model.FabuApp{
			Name: info.Name, Platform: platform, BundleId: info.BundleId,
			BundleName: info.BundleId, Version: info.Version, VersionCode: versionCode,
			Icon: iconURL, ShortUrl: genShortUrl(),
		}
		if e = response.CheckErr(s.db.Create(&app).Error, "创建应用失败"); e != nil {
			return
		}
		isNewApp = true
	} else if err != nil {
		e = response.CheckErr(err, "查询应用失败")
		return
	}

	// 创建版本
	version := fabu_model.FabuAppVersion{
		AppId: app.ID, Version: info.Version, VersionCode: versionCode,
		Size: file.Size, Md5: md5Str, DownloadUrl: downloadUrl,
		InstallUrl: downloadUrl, Released: false, UpdateMode: 0, Gray: false,
	}
	if e = response.CheckErr(s.db.Create(&version).Error, "创建版本失败"); e != nil {
		return
	}
	if platform == "ios" {
		version.InstallUrl = "/api/web/fabu/plist/" + app.ID + "/" + version.ID
		if e = response.CheckErr(s.db.Model(&version).Update("install_url", version.InstallUrl).Error, "更新安装地址失败"); e != nil {
			return
		}
	}

	res = fabu_schema.FabuUploadResp{AppId: app.ID, VersionId: version.ID, IsNewApp: isNewApp}
	return
}

func (s fabuVersionService) Release(appId, id string) (e error) {
	var version fabu_model.FabuAppVersion
	if e = response.CheckErr(s.db.Where("id = ? and app_id = ?", id, appId).First(&version).Error, "版本不存在"); e != nil {
		return
	}
	if e = response.CheckErr(s.db.Model(&fabu_model.FabuApp{}).Where("id = ?", appId).Updates(map[string]any{
		"current_version_id": version.ID, "version": version.Version, "version_code": version.VersionCode,
	}).Error, "发布失败"); e != nil {
		return
	}
	if e = response.CheckErr(s.db.Model(&version).Update("released", true).Error, "发布失败"); e != nil {
		return
	}
	return
}

func (s fabuVersionService) Cancel(appId, id string) (e error) {
	if e = response.CheckErr(s.db.Model(&fabu_model.FabuAppVersion{}).Where("id = ? and app_id = ?", id, appId).Update("released", false).Error, "取消发布失败"); e != nil {
		return
	}
	return
}

func (s fabuVersionService) Gray(appId, id string, gray bool) (e error) {
	if e = response.CheckErr(s.db.Model(&fabu_model.FabuAppVersion{}).Where("id = ? and app_id = ?", id, appId).Update("gray", gray).Error, "设置灰度失败"); e != nil {
		return
	}
	return
}

func (s fabuVersionService) UpdateMode(appId, id string, mode int) (e error) {
	if e = response.CheckErr(s.db.Model(&fabu_model.FabuAppVersion{}).Where("id = ? and app_id = ?", id, appId).Update("update_mode", mode).Error, "设置更新模式失败"); e != nil {
		return
	}
	return
}

func (s fabuVersionService) Del(appId, id string) (e error) {
	if e = response.CheckErr(s.db.Where("id = ? and app_id = ?", id, appId).Delete(&fabu_model.FabuAppVersion{}).Error, "删除版本失败"); e != nil {
		return
	}
	return
}

func (s fabuVersionService) Count(appId, versionId string) (e error) {
	s.db.Model(&fabu_model.FabuApp{}).Where("id = ?", appId).
		UpdateColumn("download_times", gorm.Expr("download_times + 1"))
	if versionId != "" {
		s.db.Model(&fabu_model.FabuAppVersion{}).Where("id = ?", versionId).
			UpdateColumn("download_times", gorm.Expr("download_times + 1"))
	}
	return
}

// PlistData 取应用与版本用于渲染 plist
func (s fabuVersionService) PlistData(appId, versionId string) (app fabu_model.FabuApp, version fabu_model.FabuAppVersion, e error) {
	if e = response.CheckErr(s.db.Where("id = ?", appId).First(&app).Error, "应用不存在"); e != nil {
		return
	}
	if e = response.CheckErr(s.db.Where("id = ? and app_id = ?", versionId, appId).First(&version).Error, "版本不存在"); e != nil {
		return
	}
	return
}

// DownloadInfo 返回下载地址（versionId 为空则取当前发布版本）
func (s fabuVersionService) DownloadInfo(appId, versionId string) (string, error) {
	var version fabu_model.FabuAppVersion
	var err error
	if versionId == "" {
		var app fabu_model.FabuApp
		if err = s.db.Where("id = ?", appId).First(&app).Error; err != nil {
			return "", response.CheckErr(err, "应用不存在")
		}
		if app.CurrentVersionId == "" {
			return "", response.AssertArgumentError.SetMessage("尚未发布版本")
		}
		err = s.db.Where("id = ?", app.CurrentVersionId).First(&version).Error
	} else {
		err = s.db.Where("id = ? and app_id = ?", versionId, appId).First(&version).Error
	}
	if err != nil {
		return "", response.CheckErr(err, "版本不存在")
	}
	return version.DownloadUrl, nil
}

// ============ 工具 ============
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

func fileMd5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// genShortUrl 生成不重复的短链（UUIDv7 前 8 位为毫秒时间戳，并发/同毫秒会撞码，故改用随机并校验唯一）
func genShortUrl() string {
	db := core.GetDB()
	for i := 0; i < 10; i++ {
		code := util.ToolsUtil.RandomString(8)
		var cnt int64
		if e := db.Model(&fabu_model.FabuApp{}).Where("short_url = ?", code).Count(&cnt).Error; e != nil {
			return code
		}
		if cnt == 0 {
			return code
		}
	}
	return util.ToolsUtil.RandomString(8)
}
