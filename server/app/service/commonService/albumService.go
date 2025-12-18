package commonService

import (
	"path"
	"x_admin/app/schema/commonSchema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/common_model"
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
func (albSrv albumService) AlbumList(adminId string, page request.PageReq, listReq commonSchema.CommonAlbumListReq) (res response.PageResp, e error) {

	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	albumModel := albSrv.db.Model(&common_model.Album{}).Where("is_delete = ?", 0)

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
	albumResps := []commonSchema.CommonAlbumListResp{}
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
	err := albSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&album).Error
	if e = response.CheckErrDBNotRecord(err, "文件丢失！"); e != nil {
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
	err := albSrv.db.Where("id in ? AND is_delete = ?", ids, 0).Find(&albums).Error
	if e = response.CheckErr(err, "AlbumMove Find err"); e != nil {
		return
	}
	if len(albums) == 0 {
		return response.AssertArgumentError.SetMessage("文件丢失！")
	}
	if cid != "" {
		err = albSrv.db.Where("id = ? AND is_delete = ?", cid, 0).Limit(1).First(&common_model.AlbumCate{}).Error
		if e = response.CheckErrDBNotRecord(err, "类目已不存在！"); e != nil {
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
func (albSrv albumService) AlbumAdd(addReq commonSchema.CommonAlbumAddReq) (res string, e error) {
	var alb common_model.Album
	//var params map[string]interface{}
	//if err := mapstructure.Decode(params, &alb); err != nil {
	//	core.Logger.Errorf("AlbumAdd Decode err: err=[%+v]", err)
	//	return response.SystemError
	//}
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
	err := albSrv.db.Where("id in ? AND is_delete = ?", ids, 0).Find(&albums).Error
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
func (albSrv albumService) CateList(adminId string, listReq commonSchema.CommonCateListReq) (mapList []commonSchema.CommonCateListResp, e error) {

	var cates []common_model.AlbumCate
	cateModel := albSrv.db.Where("is_delete = ?", 0).Order("id desc")
	// if listReq.Type > 0 {
	// 	cateModel = cateModel.Where("type = ?", listReq.Type)
	// }
	cateModel = cateModel.Where("admin_id = ?", adminId)
	if listReq.Name != "" {
		cateModel = cateModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	err := cateModel.Find(&cates).Error
	if e = response.CheckErr(err, "分类列表获取失败"); e != nil {
		return
	}
	cateResps := []commonSchema.CommonCateListResp{}
	convert_util.Copy(&cateResps, cates)
	return cateResps, nil
}

// CateAdd 分类新增
func (albSrv albumService) CateAdd(adminId string, addReq commonSchema.CommonCateAddReq) (e error) {

	var cate common_model.AlbumCate
	// 查询分类是否存在
	albSrv.db.Where("admin_id = ? AND pid=? AND name = ? AND is_delete = ?", adminId, addReq.Pid, addReq.Name, 0).Limit(1).First(&cate)
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
	err := albSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	if e = response.CheckErrDBNotRecord(err, "分类已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "重命名失败"); e != nil {
		return
	}
	var cate2 common_model.AlbumCate
	// 查询分类是否存在
	albSrv.db.Where("admin_id = ? AND pid=? AND name = ? AND is_delete = ? AND id <> ?", cate.AdminId, cate.Pid, name, 0, cate.ID).Limit(1).First(&cate2)
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
	err := albSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	if e = response.CheckErrDBNotRecord(err, "分类已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待删除数据查找失败"); e != nil {
		return
	}
	r := albSrv.db.Where("cid = ? AND is_delete = ?", id, 0).Limit(1).Find(&common_model.Album{})
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
