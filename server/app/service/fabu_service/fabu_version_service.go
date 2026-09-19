package fabu_service

import (
	"bytes"
	"errors"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"uuid"

	"x_admin/app/model/fabu_model"
	"x_admin/app/schema/fabu_schema"
	"x_admin/app/service/common_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/plugin/storage"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/file_util"

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

// Upload 解析分片上传已落盘的安装包：按 file_hash 取回文件 -> 解析(ipapk) -> 重编码图标 -> 建/追加应用与版本
// 包体在分片上传阶段已存于存储引擎，此处不再重复落盘，仅新增图标与解析元数据
func (s fabuVersionService) Upload(req fabu_schema.FabuVersionUploadReq) (res fabu_schema.FabuUploadResp, e error) {
	ext := strings.ToLower(filepath.Ext(req.FileName))
	if ext != ".ipa" && ext != ".apk" {
		e = response.AssertArgumentError.SetMessage("仅支持 .ipa / .apk 安装包")
		return
	}
	platform := "android"
	if ext == ".ipa" {
		platform = "ios"
	}

	// 定位分片上传落盘的物理文件（file_hash_id -> 存储 key）
	record, err := common_service.FileHashService.FindById(req.FileHashId)
	if err != nil || record.ID == "" {
		e = response.AssertArgumentError.SetMessage("文件不存在，请重新上传")
		return
	}
	engine := storage.GetStorageEngine()

	// 从存储引擎拷回本地临时文件供 ipapk 解析（本地/OSS 通用）
	tmpPath, err := storageToTemp(record.FilePath, ext)
	if err != nil {
		e = response.CheckErr(err, "读取安装包失败")
		return
	}
	defer os.Remove(tmpPath)

	info, err := ipapk.NewAppParser(tmpPath)
	if err != nil {
		e = response.AssertArgumentError.SetMessage("解析安装包失败: " + err.Error())
		return
	}

	// 重编码图标（iOS 优化 PNG 已由库解码为 image.Image），存入存储并登记 x_common_file_hash，走文件流访问
	iconURL := ""
	var iconData []byte
	if info.Icon != nil {
		var buf bytes.Buffer
		if encErr := png.Encode(&buf, info.Icon); encErr == nil {
			iconData = buf.Bytes()
		}
	}
	if len(iconData) > 0 {
		iconMd5, _ := file_util.ReaderMd5(bytes.NewReader(iconData))
		// 秒传：相同图标内容直接复用已有文件哈希记录
		if hit, _ := common_service.FileHashService.FindByMd5(iconMd5); hit.ID != "" {
			iconURL = util.UrlUtil.HashUrl(hit.ID, "icon.png")
		} else {
			iconKey := util.UrlUtil.BuildFileSavePath("icon.png")
			if _, putErr := engine.PutObject(iconKey, bytes.NewReader(iconData), int64(len(iconData))); putErr == nil {
				if iconHashId, creErr := common_service.FileHashService.Create(iconMd5, int64(len(iconData)), iconKey, "png"); creErr == nil && iconHashId != "" {
					iconURL = util.UrlUtil.HashUrl(iconHashId, "icon.png")
				}
			}
		}
	}

	versionCode, _ := strconv.Atoi(info.Build)
	// MD5 与包体大小复用分片上传阶段的注册值，避免对整包重算哈希
	// 包体访问地址复用分片上传注册的文件哈希 URL（/api/uploads/:id/:file_name）
	downloadUrl := util.UrlUtil.HashUrl(record.ID, req.FileName)

	// 按 BundleId + Platform 查重（iOS/Android 通常共用包名，需区分），无则新建应用；
	// 建应用/建版本/iOS 安装地址在同一事务内完成，失败则回滚并清理已落盘图标
	var isNewApp bool
	err = s.db.Transaction(func(tx *gorm.DB) error {
		var app fabu_model.FabuApp
		findErr := tx.Where("bundle_id = ? AND platform = ?", info.BundleId, platform).First(&app).Error
		if errors.Is(findErr, gorm.ErrRecordNotFound) {
			app = fabu_model.FabuApp{
				Name: info.Name, Platform: platform, BundleId: info.BundleId,
				BundleName: info.BundleId, Icon: iconURL, ShortUrl: genShortUrl(),
			}
			if e := tx.Create(&app).Error; e != nil {
				return response.CheckErr(e, "创建应用失败")
			}
			isNewApp = true
		} else if findErr != nil {
			return response.CheckErr(findErr, "查询应用失败")
		}

		// 创建版本
		version := fabu_model.FabuAppVersion{
			AppId: app.ID, Version: info.Version, VersionCode: versionCode,
			Size: record.FileSize, Md5: record.FileMd5, DownloadUrl: downloadUrl,
			InstallUrl: downloadUrl, Released: false, UpdateMode: 0, Gray: false,
		}
		if e := tx.Create(&version).Error; e != nil {
			return response.CheckErr(e, "创建版本失败")
		}
		if platform == "ios" {
			if e := tx.Model(&version).Update("install_url", "/api/web/fabu/plist/"+app.ID+"/"+version.ID).Error; e != nil {
				return response.CheckErr(e, "更新安装地址失败")
			}
		}

		res = fabu_schema.FabuUploadResp{AppId: app.ID, VersionId: version.ID, IsNewApp: isNewApp}
		return nil
	})
	if err != nil {
		e = err
		return
	}
	return
}

func (s fabuVersionService) Release(appId, id string) (e error) {
	e = s.db.Transaction(func(tx *gorm.DB) error {
		var version fabu_model.FabuAppVersion
		if err := tx.Where("id = ? and app_id = ?", id, appId).First(&version).Error; err != nil {
			return response.CheckErr(err, "版本不存在")
		}
		// 同应用下其它版本先取消发布，保证同时只有一个当前发布版本
		if err := tx.Model(&fabu_model.FabuAppVersion{}).Where("app_id = ? and id <> ?", appId, id).Update("released", false).Error; err != nil {
			return response.CheckErr(err, "发布失败")
		}
		if err := tx.Model(&fabu_model.FabuAppVersion{}).Where("id = ?", id).Update("released", true).Error; err != nil {
			return response.CheckErr(err, "发布失败")
		}
		return nil
	})
	return
}

// Cancel 取消发布：读取路径已改为 released=1 查询，仅更新版本表自身
func (s fabuVersionService) Cancel(appId, id string) (e error) {
	if e = response.CheckErr(s.db.Model(&fabu_model.FabuAppVersion{}).
		Where("id = ? and app_id = ?", id, appId).
		Update("released", false).Error, "取消发布失败"); e != nil {
		return
	}
	return
}

func (s fabuVersionService) Gray(appId, id string, gray bool) (e error) {
	if e = response.CheckErr(s.db.Model(&fabu_model.FabuAppVersion{}).
		Where("id = ? and app_id = ?", id, appId).
		Update("gray", gray).Error, "设置灰度失败"); e != nil {
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

// LatestReleased 取应用最新可用版本（released=1 且 version_code 倒序取首条）
func (s fabuVersionService) LatestReleased(appId string) (fabu_model.FabuAppVersion, bool, error) {
	var version fabu_model.FabuAppVersion
	err := s.db.Where("app_id = ? AND released = ?", appId, true).Order("version_code DESC").First(&version).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return version, false, nil
	}
	if err != nil {
		return version, false, response.CheckErr(err, "查询版本失败")
	}
	return version, true, nil
}

// FindByCode 按版本 code 定位应用下的版本记录（用于取客户端当前版本的热更包）
func (s fabuVersionService) FindByCode(appId string, code int) (fabu_model.FabuAppVersion, bool, error) {
	var version fabu_model.FabuAppVersion
	err := s.db.Where("app_id = ? AND version_code = ?", appId, code).First(&version).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return version, false, nil
	}
	if err != nil {
		return version, false, response.CheckErr(err, "查询版本失败")
	}
	return version, true, nil
}

// compareVersionName 按段比较版本号 name（如 1.0.3 vs 1.10），缺失段/非数字段按 0 处理；a>b 返回 1，相等返回 0
func compareVersionName(a, b string) int {
	segNum := func(segs []string, idx int) int {
		if idx >= len(segs) {
			return 0
		}
		v, _ := strconv.Atoi(segs[idx]) // 非法段 Atoi 报错时 v 为 0
		return v
	}
	sa, sb := strings.Split(a, "."), strings.Split(b, ".")
	n := len(sa)
	if len(sb) > n {
		n = len(sb)
	}
	for i := 0; i < n; i++ {
		if x, y := segNum(sa, i), segNum(sb, i); x != y {
			if x > y {
				return 1
			}
			return -1
		}
	}
	return 0
}

// CheckUpdate 客户端检查更新：优先最新已发布全量包（released=1 且 version_code 倒序），
// 无全量更新时再取客户端当前版本下最新发布的热更包；
// wgtVersion 为客户端当前 wgt 资源版本 name（如 1.0.3，来自 plus.runtime.version），
// 缺省时以客户端版本对应 version 的 name 为比较基准，按段比较避免去点 code 的歧义；
// 安卓/iOS 包名可能相同，需传 platform（ios/android）区分应用
func (s fabuVersionService) CheckUpdate(bundleId, platform string, clientCode int, wgtVersion string) (res fabu_schema.FabuCheckUpdateResp, e error) {
	platform = strings.ToLower(platform)
	if platform != "ios" && platform != "android" {
		return res, response.AssertArgumentError.SetMessage("platform 仅支持 ios/android")
	}
	var app fabu_model.FabuApp
	if e = AppService.FindByBundleId(bundleId, platform, &app); e != nil {
		return
	}
	// 优先：全量包更新
	version, ok, e := s.LatestReleased(app.ID)
	if e != nil {
		return
	}
	if ok && version.VersionCode > clientCode {
		return fabu_schema.FabuCheckUpdateResp{
			Update: true, Type: "app",
			App: &fabu_schema.FabuAppUpdate{
				Version:     version.Version,
				VersionCode: version.VersionCode,
				DownloadUrl: version.DownloadUrl,
				InstallUrl:  version.InstallUrl,
				UpdateMode:  version.UpdateMode,
			},
		}, nil
	}
	// 其次：客户端当前 version_code 对应版本下发布的 wgt，基准 name 优先取客户端上报，缺省用宿主版本 name
	cur, curOk, e := s.FindByCode(app.ID, clientCode)
	if e != nil || !curOk {
		return
	}
	baseName := cur.Version
	if wgtVersion != "" {
		baseName = wgtVersion
	}
	wgt, wgtOk, e := WgtService.LatestReleasedOfVersion(cur.ID)
	if e != nil {
		return
	}
	if wgtOk && compareVersionName(wgt.Version, baseName) > 0 {
		res = fabu_schema.FabuCheckUpdateResp{
			Update: true,
			Type:   "wgt",
			Wgt: &fabu_schema.FabuWgtUpdate{
				Version:     wgt.Version,
				VersionCode: wgt.VersionCode,
				DownloadUrl: wgt.DownloadUrl,
				Md5:         wgt.Md5,
				Size:        wgt.Size,
			},
		}
	}
	return
}

// DownloadInfo 返回下载地址（versionId 为空则取最新已发布版本）
func (s fabuVersionService) DownloadInfo(appId, versionId string) (string, error) {
	var version fabu_model.FabuAppVersion
	if versionId == "" {
		v, ok, err := s.LatestReleased(appId)
		if err != nil {
			return "", err
		}
		if !ok {
			return "", response.AssertArgumentError.SetMessage("尚未发布版本")
		}
		version = v
	} else {
		if err := s.db.Where("id = ? and app_id = ?", versionId, appId).First(&version).Error; err != nil {
			return "", response.CheckErr(err, "版本不存在")
		}
	}
	return version.DownloadUrl, nil
}

// DownloadPageInfo 按短链返回下载页所需的应用与最新可用版本信息（公开）
func (s fabuVersionService) DownloadPageInfo(shortUrl string) (res fabu_schema.FabuDownloadResp, e error) {
	app, e := AppService.FindByShortUrl(shortUrl)
	if e != nil {
		return
	}
	res = fabu_schema.FabuDownloadResp{
		AppId: app.ID, Name: app.Name, Platform: app.Platform, BundleId: app.BundleId,
		Icon: app.Icon, DownloadTimes: app.DownloadTimes,
	}
	// 无可用已发布版本：仅返回应用基本信息，前端据 has_version 置灰下载
	version, ok, err := s.LatestReleased(app.ID)
	if err != nil {
		e = err
		return
	}
	if !ok {
		return
	}
	res.VersionId = version.ID
	res.Version = version.Version
	res.VersionCode = version.VersionCode
	res.Size = version.Size
	res.DownloadUrl = version.DownloadUrl
	res.InstallUrl = version.InstallUrl
	res.HasVersion = true
	return
}

// ============ 工具 ============
// storageToTemp 将存储引擎中的对象拷回本地临时文件（本地/OSS 通用），返回临时文件路径，调用方负责删除
func storageToTemp(key, ext string) (string, error) {
	src, err := storage.GetStorageEngine().GetObject(key)
	if err != nil {
		return "", err
	}
	defer src.Close()
	tmpDir := filepath.Join(config.FileConfig.UploadDirectory, "fabu", "tmp")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return "", err
	}
	tmpPath := filepath.Join(tmpDir, uuid.NewV7().String()+ext)
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return "", err
	}
	if _, err = io.Copy(tmpFile, src); err != nil {
		tmpFile.Close()
		os.Remove(tmpPath)
		return "", err
	}
	return tmpPath, tmpFile.Close()
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
