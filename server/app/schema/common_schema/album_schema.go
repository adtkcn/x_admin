package common_schema

import "github.com/adtkcn/x_null"

type CommonUploadImageReq struct {
	Cid string `json:"cid" form:"cid"` // 主键
}

// CommonAlbumListReq 相册文件列表参数
type CommonAlbumListReq struct {
	Cid  string   `json:"cid" form:"cid"`   // 类目ID
	Name string   `json:"name" form:"name"` // 文件名称
	Ext  []string `json:"ext" form:"ext[]"` // 文件扩展

}

// CommonAlbumRenameReq 相册文件重命名参数
type CommonAlbumRenameReq struct {
	ID   string `json:"id" form:"id"`                                     // 主键
	Name string `json:"name" form:"name" binding:"required,min=1,max=30"` // 文件名称
}

// CommonAlbumMoveReq 相册文件移动参数
type CommonAlbumMoveReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 主键
	Cid string   `json:"cid" form:"cid"`                    // 类目ID
}

// CommonAlbumAddFromFileReq 把已登记的文件挂载到相册分类
type CommonAlbumAddFromFileReq struct {
	FileHashId string `json:"file_hash_id" validate:"required"` // 已上传文件的哈希记录ID
	Cid        string `json:"cid" validate:"required"`          // 目标相册分类ID
	FileName   string `json:"file_name" validate:"required"`    // 文件展示名
}

// CommonAlbumDelReq 相册文件删除参数
type CommonAlbumDelReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 主键
}

// CommonCateListReq 相册分类列表参数
type CommonCateListReq struct {
	Name string `json:"name" form:"name"` // 分类名称
}

// CommonCateAddReq 相册分类新增参数
type CommonCateAddReq struct {
	Pid  string `json:"pid" form:"pid" binding:"gte=0"`                   // 父级ID
	Name string `json:"name" form:"name" binding:"required,min=1,max=30"` // 分类名称
}

// CommonCateRenameReq 相册分类重命名参数
type CommonCateRenameReq struct {
	ID   string `json:"id" form:"id" binding:"required,gt=0"`             // 主键
	Name string `json:"name" form:"name" binding:"required,min=1,max=30"` // 分类名称
}

// CommonCateDelReq 相册分类删除参数
type CommonCateDelReq struct {
	ID string `json:"id" form:"id" binding:"required,gt=0"` // 主键
}

// CommonUploadFileResp 上传文件返回信息
// 单文件上传、分片上传、秒传检查（预上传）统一返回此结构
type CommonUploadFileResp struct {
	ID         string `json:"id"`           // 主键（挂载相册后由 addFromFile 返回）
	FileHashId string `json:"file_hash_id"` // 文件哈希记录ID
	Name       string `json:"name"`         // 文件名称
	Uri        string `json:"url"`          // 访问地址（完整可访问 URL）
	Path       string `json:"path"`         // 相对路径
	Ext        string `json:"ext"`          // 文件扩展
	Size       int64  `json:"size"`         // 文件大小
	Instant    bool   `json:"instant"`      // 是否秒传命中（预上传专用）
}

// CommonAlbumListResp 相册文件列表返回信息
type CommonAlbumListResp struct {
	ID         string      `json:"id"`          // 主键
	Cid        string      `json:"cid"`         // 所属类目
	Name       string      `json:"name"`        // 文件名称
	Path       string      `json:"path"`        // 相对路径
	Uri        string      `json:"uri"`         // 文件路径
	Ext        string      `json:"ext"`         // 文件扩展
	Size       string      `json:"size"`        // 文件大小
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}

// CommonCateListResp 相册分类列表返回信息
type CommonCateListResp struct {
	ID         string      `json:"id"`          // 主键
	Pid        string      `json:"pid"`         // 父级ID
	Name       string      `json:"name"`        // 分类名称
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}
