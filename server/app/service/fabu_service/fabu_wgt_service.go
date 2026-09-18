package fabu_service

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"x_admin/app/model/fabu_model"
	"x_admin/app/schema/fabu_schema"
	"x_admin/app/service/common_service"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var WgtService = NewFabuWgtService()

func NewFabuWgtService() *fabuWgtService {
	return &fabuWgtService{db: core.GetDB()}
}

type fabuWgtService struct {
	db *gorm.DB
}

func (s fabuWgtService) List(req fabu_schema.FabuWgtListReq) (res map[string]any, e error) {
	var list []fabu_model.FabuWgt
	var total int64
	chain := s.db.Model(&fabu_model.FabuWgt{})
	if req.VersionId != "" {
		chain = chain.Where("version_id = ?", req.VersionId)
	}
	if e = response.CheckErr(chain.Count(&total).Error, "统计wgt失败"); e != nil {
		return
	}
	pageNo, pageSize := req.PageNo, req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if e = response.CheckErr(chain.Order("create_time desc").Offset((pageNo-1)*pageSize).Limit(pageSize).Find(&list).Error, "查询wgt失败"); e != nil {
		return
	}
	var data []fabu_schema.FabuWgtResp
	convert_util.Copy(&data, list)
	res = map[string]any{"count": total, "lists": data}
	return
}

type wgtManifestVersion struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}
type wgtManifest struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Version     wgtManifestVersion `json:"version"`
	Description string             `json:"description"`
}

// Upload 登记分片上传已落盘的 wgt 包：按 file_hash 取回文件 -> 解析 manifest -> 建记录
// 包体已在 x_common_file_hash/存储引擎中，下载地址复用文件哈希 URL，不重复落盘
func (s fabuWgtService) Upload(req fabu_schema.FabuWgtUploadReq) (e error) {
	if strings.ToLower(filepath.Ext(req.FileName)) != ".wgt" {
		return response.AssertArgumentError.SetMessage("仅支持 .wgt 热更新包")
	}
	var version fabu_model.FabuAppVersion
	if e = response.CheckErr(s.db.Where("id = ?", req.VersionId).First(&version).Error, "版本不存在"); e != nil {
		return
	}

	// 定位分片上传落盘的物理文件
	record, err := common_service.FileHashService.FindById(req.FileHashId)
	if err != nil || record.ID == "" {
		return response.AssertArgumentError.SetMessage("文件不存在，请重新上传")
	}

	// 经存储引擎拷回本地临时文件解析 manifest（本地/OSS 通用）
	tmpPath, err := storageToTemp(record.FilePath, ".wgt")
	if err != nil {
		return response.CheckErr(err, "读取wgt失败")
	}
	defer os.Remove(tmpPath)

	ver, versionCode, err := parseWgtManifest(tmpPath)
	if err != nil {
		return response.AssertArgumentError.SetMessage("解析wgt失败: " + err.Error())
	}

	// MD5 分片上传阶段已算过，直接复用注册值
	wgt := fabu_model.FabuWgt{
		AppId:       version.AppId,
		VersionId:   req.VersionId,
		Version:     ver,
		VersionCode: versionCode,
		DownloadUrl: util.UrlUtil.HashUrl(record.ID, req.FileName),
		Md5:         record.FileMd5, Size: record.FileSize,
	}
	if e = response.CheckErr(s.db.Create(&wgt).Error, "创建wgt失败"); e != nil {
		return
	}
	return
}

// LatestReleasedOfVersion 取指定版本下最新发布的热更包（released=1，version_code 倒序兜底）
func (s fabuWgtService) LatestReleasedOfVersion(versionId string) (fabu_model.FabuWgt, bool, error) {
	var wgt fabu_model.FabuWgt
	err := s.db.Where("version_id = ? AND released = ?", versionId, true).Order("version_code DESC").First(&wgt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return wgt, false, nil
	}
	if err != nil {
		return wgt, false, response.CheckErr(err, "查询热更新包失败")
	}
	return wgt, true, nil
}

// Release 切换热更新包发布状态；发布时同版本下其它 wgt 全部取消，保证一个版本仅一个 released=1
func (s fabuWgtService) Release(req fabu_schema.FabuWgtReleaseReq) (e error) {
	var wgt fabu_model.FabuWgt
	if e = response.CheckErr(s.db.Where("id = ?", req.ID).First(&wgt).Error, "热更新包不存在"); e != nil {
		return
	}
	if !req.Released {
		return response.CheckErr(s.db.Model(&fabu_model.FabuWgt{}).Where("id = ?", req.ID).Update("released", false).Error, "取消发布失败")
	}
	e = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&fabu_model.FabuWgt{}).Where("version_id = ? AND id <> ?", wgt.VersionId, req.ID).
			Update("released", false).Error; err != nil {
			return response.CheckErr(err, "发布失败")
		}
		if err := tx.Model(&fabu_model.FabuWgt{}).Where("id = ?", req.ID).Update("released", true).Error; err != nil {
			return response.CheckErr(err, "发布失败")
		}
		return nil
	})
	return
}

func (s fabuWgtService) Del(id string) (e error) {
	if e = response.CheckErr(s.db.Where("id = ?", id).Delete(&fabu_model.FabuWgt{}).Error, "删除wgt失败"); e != nil {
		return
	}
	return
}

func parseWgtManifest(path string) (version string, versionCode int, e error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return "", 0, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == "manifest.json" {
			rc, err := f.Open()
			if err != nil {
				return "", 0, err
			}
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return "", 0, err
			}
			var m wgtManifest
			if err := json.Unmarshal(data, &m); err != nil {
				return "", 0, err
			}
			return m.Version.Name, m.Version.Code, nil
		}
	}
	return "", 0, fmt.Errorf("manifest.json not found")
}
