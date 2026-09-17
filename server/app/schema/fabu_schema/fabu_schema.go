package fabu_schema

// ============ 应用 ============
type FabuAppAddReq struct {
	Name        string `json:"name" form:"name"`
	Platform    string `json:"platform" form:"platform"`
	BundleId    string `json:"bundle_id" form:"bundle_id"`
	BundleName  string `json:"bundle_name" form:"bundle_name"`
	Version     string `json:"version" form:"version"`
	VersionCode int    `json:"version_code" form:"version_code"`
	ShortUrl    string `json:"short_url" form:"short_url"`
	Icon        string `json:"icon" form:"icon"`
}

type FabuAppEditReq struct {
	ID         string `json:"id" form:"id" binding:"required"`
	Name       string `json:"name" form:"name"`
	BundleName string `json:"bundle_name" form:"bundle_name"`
	ShortUrl   string `json:"short_url" form:"short_url"`
	Icon       string `json:"icon" form:"icon"`
}

type FabuAppDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type FabuAppDelReq struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type FabuAppListReq struct {
	Keyword  string `json:"keyword" form:"keyword"`
	PageNo   int    `json:"pageNo" form:"pageNo"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type FabuAppResp struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Platform         string `json:"platform"`
	BundleId         string `json:"bundle_id"`
	BundleName       string `json:"bundle_name"`
	Version          string `json:"version"`
	VersionCode      int    `json:"version_code"`
	ShortUrl         string `json:"short_url"`
	Icon             string `json:"icon"`
	DownloadTimes    int    `json:"download_times"`
	CurrentVersionId string `json:"current_version_id"`
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
	AppId     string `json:"app_id" form:"app_id" binding:"required"`
	ID        string `json:"id" form:"id" binding:"required"`
	UpdateMode int   `json:"update_mode" form:"update_mode"`
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

type FabuWgtResp struct {
	ID          string `json:"id"`
	AppId       string `json:"app_id"`
	VersionId   string `json:"version_id"`
	Version     string `json:"version"`
	VersionCode int    `json:"version_code"`
	DownloadUrl string `json:"download_url"`
	Md5         string `json:"md5"`
	Size        int64  `json:"size"`
	CreateTime  string `json:"create_time"`
}

// ============ 上传返回 ============
type FabuUploadResp struct {
	AppId    string `json:"app_id"`
	VersionId string `json:"version_id"`
	IsNewApp bool   `json:"is_new_app"`
}
