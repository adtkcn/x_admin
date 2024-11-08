package article_collect

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

type IArticleCollectService interface {
	List(page request.PageReq, listReq ArticleCollectListReq) (res response.PageResp, e error)
	Detail(id int) (res ArticleCollectResp, e error)
	Add(addReq ArticleCollectAddReq) (e error)
	Edit(editReq ArticleCollectEditReq) (e error)
	Del(id int) (e error)
}

var Service = NewArticleCollectService()

// NewArticleCollectService 初始化
func NewArticleCollectService() ArticleCollectService {
	db := core.GetDB()
	return ArticleCollectService{db: db}
}

// ArticleCollectService 文章收藏服务实现类
type ArticleCollectService struct {
	db *gorm.DB
}

// List 文章收藏列表
func (service ArticleCollectService) List(page request.PageReq, listReq ArticleCollectListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.ArticleCollect{})
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
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.ArticleCollect
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	result := []ArticleCollectResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}
