package fabu_service

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"uuid"

	"x_admin/app/model/fabu_model"
	"x_admin/app/schema/fabu_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
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

func (s fabuWgtService) Upload(file *multipart.FileHeader, versionId string) (e error) {
	var version fabu_model.FabuAppVersion
	if e = response.CheckErr(s.db.Where("id = ?", versionId).First(&version).Error, "版本不存在"); e != nil {
		return
	}
	pkgUUID := uuid.NewV7().String()
	pkgDir := filepath.Join(config.FileConfig.UploadDirectory, "fabu", "wgt")
	if e = response.CheckErr(os.MkdirAll(pkgDir, 0755), "创建目录失败"); e != nil {
		return
	}
	pkgPath := filepath.Join(pkgDir, pkgUUID+".wgt")
	if e = response.CheckErr(saveUploadedFile(file, pkgPath), "保存失败"); e != nil {
		return
	}

	ver, versionCode, err := parseWgtManifest(pkgPath)
	if err != nil {
		e = response.AssertArgumentError.SetMessage("解析wgt失败: " + err.Error())
		return
	}
	md5Str, _ := fileMd5(pkgPath)
	wgt := fabu_model.FabuWgt{
		AppId:       version.AppId,
		VersionId:   versionId,
		Version:     ver,
		VersionCode: versionCode,
		DownloadUrl: "/api/fabu/static/fabu/wgt/" + pkgUUID + ".wgt",
		Md5:         md5Str, Size: file.Size,
	}
	if e = response.CheckErr(s.db.Create(&wgt).Error, "创建wgt失败"); e != nil {
		return
	}
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
