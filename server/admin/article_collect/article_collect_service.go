package article_collect

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"

	"gorm.io/gorm"
)

type IArticleCollectService interface {
	List(page request.PageReq, listReq ArticleCollectListReq) (res response.PageResp, e error)
	Detail(id int) (res ArticleCollectResp, e error)
	Add(addReq ArticleCollectAddReq) (e error)
	Edit(editReq ArticleCollectEditReq) (e error)
	Del(id int) (e error)
}

// NewArticleCollectService 初始化
func NewArticleCollectService(db *gorm.DB) IArticleCollectService {
	return &articleCollectService{db: db}
}

// articleCollectService 文章收藏服务实现类
type articleCollectService struct {
	db *gorm.DB
}

// List 文章收藏列表
func (Service articleCollectService) List(page request.PageReq, listReq ArticleCollectListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := Service.db.Model(&model.ArticleCollect{})
	if listReq.UserId > 0 {
		dbModel = dbModel.Where("user_id = ?", listReq.UserId)
	}
	if listReq.ArticleId > 0 {
		dbModel = dbModel.Where("article_id = ?", listReq.ArticleId)
	}
	dbModel = dbModel.Where("is_delete = ?", 0)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.ArticleCollect
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []ArticleCollectResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// Detail 文章收藏详情
func (Service articleCollectService) Detail(id int) (res ArticleCollectResp, e error) {
	var obj model.ArticleCollect
	err := Service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, obj)
	return
}

// Add 文章收藏新增
func (Service articleCollectService) Add(addReq ArticleCollectAddReq) (e error) {
	var obj model.ArticleCollect
	response.Copy(&obj, addReq)
	err := Service.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 文章收藏编辑
func (Service articleCollectService) Edit(editReq ArticleCollectEditReq) (e error) {
	var obj model.ArticleCollect
	err := Service.db.Where("id = ? AND is_delete = ?", editReq.Id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = Service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 文章收藏删除
func (Service articleCollectService) Del(id int) (e error) {
	var obj model.ArticleCollect
	err := Service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 删除
	obj.IsDelete = 1
	err = Service.db.Save(&obj).Error
	e = response.CheckErr(err, "Del Save err")
	return
}
