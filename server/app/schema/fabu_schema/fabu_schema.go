package fabu_schema

// ============ 应用 ============
type FabuAppDelReq struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type FabuAppListReq struct {
	Keyword  string `json:"keyword" form:"keyword"`
	PageNo   int    `json:"pageNo" form:"pageNo"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type FabuAppResp struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Platform      string `json:"platform"`
	BundleId      string `json:"bundle_id"`
	BundleName    string `json:"bundle_name"`
	ShortUrl      string `json:"short_url"`
	Icon          string `json:"icon"`
	DownloadTimes int    `json:"download_times"`
	CreateTime    string `json:"create_time"`
}

// ============ 版本 ============
type FabuVersionListReq struct {
	AppId    string `json:"app_id" form:"app_id" binding:"required"`
	PageNo   int    `json:"pageNo" form:"pageNo"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type FabuVersionReleaseReq struct {
	AppId string `json:"app_id" form:"app_id" binding:"required"`
	ID    string `json:"id" form:"id" binding:"required"`
}

type FabuVersionGrayReq struct {
	AppId string `json:"app_id" form:"app_id" binding:"required"`
	ID    string `json:"id" form:"id" binding:"required"`
	Gray  bool   `json:"gray" form:"gray"`
}

type FabuVersionUpdateModeReq struct {
	AppId      string `json:"app_id" form:"app_id" binding:"required"`
	ID         string `json:"id" form:"id" binding:"required"`
	UpdateMode int    `json:"update_mode" form:"update_mode"`
}

type FabuVersionDelReq struct {
	AppId string `json:"app_id" form:"app_id" binding:"required"`
	ID    string `json:"id" form:"id" binding:"required"`
}

type FabuVersionResp struct {
	ID            string `json:"id"`
	AppId         string `json:"app_id"`
	Version       string `json:"version"`
	VersionCode   int    `json:"version_code"`
	Size          int64  `json:"size"`
	Md5           string `json:"md5"`
	DownloadUrl   string `json:"download_url"`
	InstallUrl    string `json:"install_url"`
	Released      bool   `json:"released"`
	UpdateMode    int    `json:"update_mode"`
	Gray          bool   `json:"gray"`
	DownloadTimes int    `json:"download_times"`
	CreateTime    string `json:"create_time"`
}

// ============ wgt（版本子表） ============
type FabuWgtListReq struct {
	VersionId string `json:"version_id" form:"version_id" binding:"required"`
	PageNo    int    `json:"pageNo" form:"pageNo"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

type FabuWgtDelReq struct {
	ID string `json:"id" form:"id" binding:"required"`
}

// FabuWgtUploadReq 分片上传完成后，后端按文件引用解析 manifest 并登记 wgt 记录
type FabuWgtUploadReq struct {
	VersionId  string `json:"version_id" form:"version_id" binding:"required"`
	FileHashId string `json:"file_hash_id" form:"file_hash_id" binding:"required"` // 分片上传注册的文件哈希ID
	FileName   string `json:"file_name" form:"file_name" binding:"required"`       // 原始文件名（含 .wgt 扩展名）
}

// FabuWgtReleaseReq 切换热更新包发布状态（released 为目标状态）
type FabuWgtReleaseReq struct {
	ID       string `json:"id" form:"id" binding:"required"`
	Released bool   `json:"released" form:"released"`
}

type FabuWgtResp struct {
	ID          string `json:"id"`
	AppId       string `json:"app_id"`
	VersionId   string `json:"version_id"`
	Version     string `json:"version"`
	VersionCode int    `json:"version_code"`
	DownloadUrl string `json:"download_url"`
	Md5         string `json:"md5"`
	Size        int64  `json:"size"`
	Released    bool   `json:"released"`
	CreateTime  string `json:"create_time"`
}

// ============ 上传返回 ============
type FabuUploadResp struct {
	AppId     string `json:"app_id"`
	VersionId string `json:"version_id"`
	IsNewApp  bool   `json:"is_new_app"`
}

// FabuVersionUploadReq 分片上传完成后，后端按文件引用解析并建应用/版本
type FabuVersionUploadReq struct {
	FileHashId string `json:"file_hash_id" form:"file_hash_id" binding:"required"` // 分片上传注册的文件哈希ID
	FileName   string `json:"file_name" form:"file_name" binding:"required"`       // 原始文件名（含 .ipa/.apk 扩展名）
}

// ============ 检查更新 ============
// FabuCheckUpdateResp 检查更新返回：update=true 时 type 标明更新方式（app 全量包 / wgt 热更包），仅对应数据块有值；
// 示例：{"update":true,"type":"app","app":{...}} / {"update":true,"type":"wgt","wgt":{...}} / {"update":false}
type FabuCheckUpdateResp struct {
	Update bool           `json:"update"`
	Type   string         `json:"type,omitempty"`
	App    *FabuAppUpdate `json:"app,omitempty"`
	Wgt    *FabuWgtUpdate `json:"wgt,omitempty"`
}

// FabuAppUpdate 全量安装包更新信息
type FabuAppUpdate struct {
	Version     string `json:"version"`
	VersionCode int    `json:"version_code"`
	DownloadUrl string `json:"download_url"`
	InstallUrl  string `json:"install_url"`
	UpdateMode  int    `json:"update_mode"`
}

// FabuWgtUpdate 热更新包更新信息
type FabuWgtUpdate struct {
	Version     string `json:"version"`
	VersionCode int    `json:"version_code"`
	DownloadUrl string `json:"download_url"`
	Md5         string `json:"md5"`
	Size        int64  `json:"size"`
}

// ============ 公开下载页 ============
// FabuDownloadResp 下载页展示的应用与当前发布版本信息
type FabuDownloadResp struct {
	AppId         string `json:"app_id"`
	VersionId     string `json:"version_id"`
	Name          string `json:"name"`
	Platform      string `json:"platform"`
	BundleId      string `json:"bundle_id"`
	Icon          string `json:"icon"`
	Version       string `json:"version"`
	VersionCode   int    `json:"version_code"`
	Size          int64  `json:"size"`
	DownloadUrl   string `json:"download_url"`
	InstallUrl    string `json:"install_url"`
	DownloadTimes int    `json:"download_times"`
	HasVersion    bool   `json:"has_version"` // 是否已有发布版本（无则不可下载）
}
