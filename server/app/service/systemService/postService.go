package systemService

import (
	"errors"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/systemSchema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var PostService = NewSystemAuthPostService()

// NewSystemAuthPostService 初始化
func NewSystemAuthPostService() *systemAuthPostService {
	db := core.GetDB()
	return &systemAuthPostService{db: db}
}

// systemAuthPostService 系统岗位服务实现类
type systemAuthPostService struct {
	db *gorm.DB
}

// All 岗位所有
func (service systemAuthPostService) All() (res []systemSchema.SystemAuthPostResp, e error) {
	var posts []system_model.SystemAuthPost
	err := service.db.Order("sort desc, id desc").Find(&posts).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []systemSchema.SystemAuthPostResp{}
	convert_util.Copy(&res, posts)
	return
}

// List 岗位列表
func (service systemAuthPostService) List(page request.PageReq, listReq systemSchema.SystemAuthPostListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	postModel := service.db.Model(&system_model.SystemAuthPost{})
	if listReq.Code != "" {
		postModel = postModel.Where("code like ?", "%"+listReq.Code+"%")
	}
	if listReq.Name != "" {
		postModel = postModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	// if listReq.IsStop > 0 {
	// 	postModel = postModel.Where("is_stop = ?", listReq.IsStop)
	// }
	// 总数
	var count int64
	err := postModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var posts []system_model.SystemAuthPost
	err = postModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&posts).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	postResps := []systemSchema.SystemAuthPostResp{}
	convert_util.Copy(&postResps, posts)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    postResps,
	}, nil
}

// Detail 部门详情
func (service systemAuthPostService) Detail(id string) (res systemSchema.SystemAuthPostResp, e error) {
	var post system_model.SystemAuthPost
	err := service.db.Where("id = ?", id).First(&post).Error
	if e = response.CheckDBNotRecord(err, "岗位不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, post)
	return
}

// Add 部门新增
func (service systemAuthPostService) Add(addReq systemSchema.SystemAuthPostAddReq) (e error) {
	r := service.db.Where("(code = ? OR name = ?)", addReq.Code, addReq.Name).Limit(1).Find(&system_model.SystemAuthPost{})
	if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该岗位已存在!")
	}
	var post system_model.SystemAuthPost
	convert_util.Copy(&post, addReq)
	err := service.db.Create(&post).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 岗位编辑
func (service systemAuthPostService) Edit(editReq systemSchema.SystemAuthPostEditReq) (e error) {
	// 检查岗位是否存在
	var post system_model.SystemAuthPost
	err := service.db.Where("id = ?", editReq.ID).First(&post).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("岗位不存在")
		}
		return response.CheckErr(err, "查询岗位失败")
	}

	// 检查编码和名称是否重复
	if r := service.db.Where("(code = ? OR name = ?) AND id != ?", editReq.Code, editReq.Name, editReq.ID).Limit(1).Find(&system_model.SystemAuthPost{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该岗位已存在!")
	}

	// 更新
	convert_util.Copy(&post, editReq)
	result := service.db.Model(&post).Select("*").Updates(post)
	if result.Error != nil {
		return response.CheckErr(result.Error, "编辑失败")
	}
	return
}

// Del 岗位删除
func (service systemAuthPostService) Del(id string) (e error) {
	// 检查岗位是否被使用
	if r := service.db.Where("post_id = ?", id).Limit(1).Find(&system_model.SystemAuthAdmin{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该岗位存在管理员,请先移除!")
	}

	result := service.db.Where("id = ?", id).Delete(&system_model.SystemAuthPost{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("岗位不存在")
	}
	return
}
