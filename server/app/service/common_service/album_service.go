package common_service

import (
	"path"
	"x_admin/app/model/common_model"
	"x_admin/app/schema/common_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var AlbumService = NewAlbumService()

// NewAlbumService 初始化
func NewAlbumService() *albumService {
	db := core.GetDB()
	return &albumService{db: db}
}

// albumService 相册服务实现类
type albumService struct {
	db *gorm.DB
}

// AlbumList 相册文件列表。
// 文件信息（uri/ext/size）不存于 Album 表，统一经 file_hash_id 关联 x_common_file_hash 取得。
// 采用「先查 album 分页 + 再批量查 file_hash」两段式，避免 JOIN 列映射、COUNT 干扰等坑。
func (albSrv albumService) AlbumList(adminId string, page request.PageReq, listReq common_schema.CommonAlbumListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	query := albSrv.db.Model(&common_model.Album{}).Where("admin_id = ?", adminId)
	if listReq.Cid != "" {
		query = query.Where("cid = ?", listReq.Cid)
	}
	if listReq.Name != "" {
		query = query.Where("name like ?", "%"+listReq.Name+"%")
	}
	// 扩展名过滤走 file_hash 子查询（album 表不存 ext）
	if len(listReq.Ext) > 0 {
		query = query.Where("file_hash_id IN (?)",
			albSrv.db.Model(&common_model.CommonFileHash{}).Select("id").Where("ext IN ?", listReq.Ext))
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "Album列表总数获取失败"); e != nil {
		return
	}

	// 分页数据
	var albums []common_model.Album
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&albums).Error
	if e = response.CheckErr(err, "Album列表获取失败"); e != nil {
		return
	}

	// 批量取关联的文件哈希，内存组装 uri/ext/size
	hashMap := map[string]common_model.CommonFileHash{}
	if len(albums) > 0 {
		ids := make([]string, 0, len(albums))
		for _, a := range albums {
			ids = append(ids, a.FileHashId)
		}
		var hashes []common_model.CommonFileHash
		if err = albSrv.db.Where("id IN ?", ids).Find(&hashes).Error; err != nil {
			e = response.CheckErr(err, "文件信息查询失败")
			return
		}
		for _, h := range hashes {
			hashMap[h.ID] = h
		}
	}

	albumResps := make([]common_schema.CommonAlbumListResp, 0, len(albums))
	for _, alb := range albums {
		albumResps = append(albumResps, buildAlbumListResp(alb, hashMap[alb.FileHashId]))
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    albumResps,
	}, nil
}

// buildAlbumListResp 由相册行 + 关联文件哈希组装列表返回。
// uri/ext/size 来自 x_common_file_hash（经 file_hash_id），Album 自身不再冗余存储。
func buildAlbumListResp(alb common_model.Album, hash common_model.CommonFileHash) common_schema.CommonAlbumListResp {
	return common_schema.CommonAlbumListResp{
		ID:         alb.ID,
		Cid:        alb.Cid,
		Name:       alb.Name,
		Path:       hash.FilePath,
		Uri:        path.Join(config.FileConfig.UploadPrefix, hash.FilePath),
		Ext:        hash.Ext,
		Size:       util.ServerUtil.GetFmtSize(uint64(hash.FileSize)),
		CreateTime: alb.CreateTime,
		UpdateTime: alb.UpdateTime,
	}
}

// AlbumRename 相册文件重命名（更新 Album 表的 name）
func (albSrv albumService) AlbumRename(id string, name string) (e error) {
	var album common_model.Album
	err := albSrv.db.Where("id = ?", id).First(&album).Error
	if e = response.CheckDBNotRecord(err, "文件丢失！"); e != nil {
		return
	}
	err = albSrv.db.Model(&common_model.Album{}).
		Where("id = ?", id).
		UpdateColumn("name", name).Error
	e = response.CheckErr(err, "AlbumRename err")
	return
}

// AlbumMove 相册文件移动
func (albSrv albumService) AlbumMove(ids []string, cid string) (e error) {
	var albums []common_model.Album
	err := albSrv.db.Where("id in ?", ids).Find(&albums).Error
	if e = response.CheckErr(err, "AlbumMove Find err"); e != nil {
		return
	}
	if len(albums) == 0 {
		return response.AssertArgumentError.SetMessage("文件丢失！")
	}
	if cid != "" {
		err = albSrv.db.Where("id = ?", cid).First(&common_model.AlbumCate{}).Error
		if e = response.CheckDBNotRecord(err, "类目已不存在！"); e != nil {
			return
		}
		if e = response.CheckErr(err, "AlbumMove First err"); e != nil {
			return
		}
	}
	err = albSrv.db.Model(&common_model.Album{}).Where("id in ?", ids).UpdateColumn("cid", cid).Error
	e = response.CheckErr(err, "AlbumMove UpdateColumn err")
	return
}

// AlbumAddFromFileRef 把已登记的文件（x_common_file_hash）挂载到相册分类：
// 相册行只记录 file_hash_id 关联，文件信息（uri/ext/hash/size）经关联查询获得，不冗余存储。
func (albSrv albumService) AlbumAddFromFileRef(fileHashId, fileName, cid, adminId string) (resp common_schema.CommonAlbumListResp, e error) {
	// 校验分类存在（cid 为空 表示未归类，跳过校验）
	var err error
	if cid != "" {
		var category common_model.AlbumCate
		err = albSrv.db.Where("id = ?", cid).First(&category).Error
		if e = response.CheckDBNotRecord(err, "相册分类不存在"); e != nil {
			return
		}
	}
	// 取文件哈希记录（上传时登记），仅用于校验存在性
	var hash common_model.CommonFileHash
	err = albSrv.db.Where("id = ?", fileHashId).First(&hash).Error
	if e = response.CheckDBNotRecord(err, "文件不存在或已过期"); e != nil {
		return
	}
	// 新建相册行：仅挂载关联与元信息，不冗余复制文件字段
	var alb common_model.Album
	alb.Cid = cid
	alb.AdminId = adminId
	alb.FileHashId = hash.ID
	alb.Name = fileName
	err = albSrv.db.Create(&alb).Error
	if e = response.CheckErr(err, "相册添加失败"); e != nil {
		return
	}
	// 返回时填充文件信息（来自关联的 file_hash）
	resp = buildAlbumListResp(alb, hash)
	return
}

// AlbumDel 相册文件删除（软删相册行；物理文件由 CleanOrphanFiles 基于访问时间/冷热清理）
func (albSrv albumService) AlbumDel(ids []string) (e error) {
	var albums []common_model.Album
	err := albSrv.db.Where("id in ?", ids).Find(&albums).Error
	if e = response.CheckErr(err, "相册文件查找失败"); e != nil {
		return
	}
	if len(albums) == 0 {
		return response.AssertArgumentError.SetMessage("文件丢失！")
	}
	err = albSrv.db.Model(&common_model.Album{}).Where("id in ?", ids).Updates(
		common_model.Album{IsDelete: 1, DeleteTime: util.NullTimeUtil.Now()}).Error
	e = response.CheckErr(err, "相册文件删除失败")
	return
}

// CateList 相册分类列表
func (albSrv albumService) CateList(adminId string, listReq common_schema.CommonCateListReq) (mapList []common_schema.CommonCateListResp, e error) {

	var cates []common_model.AlbumCate
	cateModel := albSrv.db.Order("id desc")

	cateModel = cateModel.Where("admin_id = ?", adminId)
	if listReq.Name != "" {
		cateModel = cateModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	err := cateModel.Find(&cates).Error
	if e = response.CheckErr(err, "分类列表获取失败"); e != nil {
		return
	}
	cateResps := []common_schema.CommonCateListResp{}
	convert_util.Copy(&cateResps, cates)
	return cateResps, nil
}

// CateAdd 分类新增
func (albSrv albumService) CateAdd(adminId string, addReq common_schema.CommonCateAddReq) (e error) {

	var cate common_model.AlbumCate
	// 查询分类是否存在
	albSrv.db.Where("admin_id = ? AND pid=? AND name = ?", adminId, addReq.Pid, addReq.Name).First(&cate)
	if cate.ID != "" {
		return response.AssertArgumentError.SetMessage("分类已存在！")
	}

	convert_util.Copy(&cate, addReq)
	cate.AdminId = adminId

	err := albSrv.db.Create(&cate).Error
	e = response.CheckErr(err, "分类添加失败")
	return
}

// CateRename 分类重命名
func (albSrv albumService) CateRename(id string, name string) (e error) {
	var cate common_model.AlbumCate
	err := albSrv.db.Where("id = ?", id).First(&cate).Error
	if e = response.CheckDBNotRecord(err, "分类已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "重命名失败"); e != nil {
		return
	}
	var cate2 common_model.AlbumCate
	albSrv.db.Where("admin_id = ? AND pid=? AND name = ? AND id <> ?", cate.AdminId, cate.Pid, name, cate.ID).First(&cate2)
	if cate2.ID != "" {
		return response.AssertArgumentError.SetMessage("分类“" + name + "”已存在！")
	}

	cate.Name = name
	err = albSrv.db.Save(&cate).Error
	e = response.CheckErr(err, "分类重命名失败")
	return
}

// CateDel 分类删除
func (albSrv albumService) CateDel(id string) (e error) {
	var cate common_model.AlbumCate
	err := albSrv.db.Where("id = ?", id).First(&cate).Error
	if e = response.CheckDBNotRecord(err, "分类已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待删除数据查找失败"); e != nil {
		return
	}
	r := albSrv.db.Where("cid = ?", id).Limit(1).Find(&common_model.Album{})
	if e = response.CheckErr(r.Error, "分类数据使用失败"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("当前分类正被使用中,不能删除！")
	}
	err = albSrv.db.Delete(&cate).Error
	e = response.CheckErr(err, "分类删除失败")
	return
}
