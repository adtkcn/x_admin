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

// AlbumList 相册文件列表
func (albSrv albumService) AlbumList(adminId string, page request.PageReq, listReq common_schema.CommonAlbumListReq) (res response.PageResp, e error) {

	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	albumModel := albSrv.db.Model(&common_model.Album{})

	albumModel = albumModel.Where("admin_id = ?", adminId)

	if listReq.Cid != "" && listReq.Cid != "0" {
		albumModel = albumModel.Where("cid = ?", listReq.Cid)
	}
	if listReq.Name != "" {
		albumModel = albumModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if len(listReq.Ext) > 0 {
		albumModel = albumModel.Where("ext in ?", listReq.Ext)
	}

	// 总数
	var count int64
	err := albumModel.Count(&count).Error
	if e = response.CheckErr(err, "Album列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var albums []common_model.Album
	err = albumModel.Limit(limit).Offset(offset).Order("id desc").Find(&albums).Error
	if e = response.CheckErr(err, "Album列表获取失败"); e != nil {
		return
	}
	albumResps := []common_schema.CommonAlbumListResp{}
	convert_util.Copy(&albumResps, albums)
	// TODO: engine默认local
	engine := "local"
	for i := 0; i < len(albumResps); i++ {
		if engine == "local" {
			albumResps[i].Path = path.Join(config.FileConfig.PublicPrefix, albums[i].Uri)
		} else {
			// TODO: 其他engine
		}
		albumResps[i].Uri = util.UrlUtil.ToAbsoluteUrl(albums[i].Uri)
		albumResps[i].Size = util.ServerUtil.GetFmtSize(uint64(albums[i].Size))
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    albumResps,
	}, nil
}

// AlbumRename 相册文件重命名
func (albSrv albumService) AlbumRename(id string, name string) (e error) {
	var album common_model.Album
	err := albSrv.db.Where("id = ?", id).First(&album).Error
	if e = response.CheckDBNotRecord(err, "文件丢失！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "AlbumRename First err"); e != nil {
		return
	}
	album.Name = name
	err = albSrv.db.Save(&album).Error
	e = response.CheckErr(err, "AlbumRename Save err")
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

// AlbumAdd 相册文件新增
func (albSrv albumService) AlbumAdd(addReq common_schema.CommonAlbumAddReq) (res string, e error) {
	var alb common_model.Album

	convert_util.Copy(&alb, addReq)
	err := albSrv.db.Create(&alb).Error
	if e = response.CheckErr(err, "相册添加失败"); e != nil {
		return
	}
	return alb.ID, nil
}

// AlbumDel 相册文件删除
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
	cate.IsDelete = 1
	cate.DeleteTime = util.NullTimeUtil.Now()
	err = albSrv.db.Save(&cate).Error
	e = response.CheckErr(err, "CateDel Save err")
	return
}
