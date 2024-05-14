package post

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/system_model"

	"gorm.io/gorm"
)

// type ISystemAuthPostService interface {
// 	All() (res []SystemAuthPostResp, e error)
// 	List(page request.PageReq, listReq SystemAuthPostListReq) (res response.PageResp, e error)
// 	Detail(id uint) (res SystemAuthPostResp, e error)
// 	Add(addReq SystemAuthPostAddReq) (e error)
// 	Edit(editReq SystemAuthPostEditReq) (e error)
// 	Del(id uint) (e error)
// }

var Service = NewSystemAuthPostService()

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
func (service systemAuthPostService) All() (res []SystemAuthPostResp, e error) {
	var posts []system_model.SystemAuthPost
	err := service.db.Order("sort desc, id desc").Find(&posts).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []SystemAuthPostResp{}
	response.Copy(&res, posts)
	return
}

// List 岗位列表
func (service systemAuthPostService) List(page request.PageReq, listReq SystemAuthPostListReq) (res response.PageResp, e error) {
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
	postResps := []SystemAuthPostResp{}
	response.Copy(&postResps, posts)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    postResps,
	}, nil
}

// Detail 部门详情
func (service systemAuthPostService) Detail(id uint) (res SystemAuthPostResp, e error) {
	var post system_model.SystemAuthPost
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&post).Error
	if e = response.CheckErrDBNotRecord(err, "岗位不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	response.Copy(&res, post)
	return
}

// Add 部门新增
func (service systemAuthPostService) Add(addReq SystemAuthPostAddReq) (e error) {
	r := service.db.Where("(code = ? OR name = ?) AND is_delete = ?", addReq.Code, addReq.Name, 0).Limit(1).Find(&system_model.SystemAuthPost{})
	if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该岗位已存在!")
	}
	var post system_model.SystemAuthPost
	response.Copy(&post, addReq)
	err := service.db.Create(&post).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 部门编辑
func (service systemAuthPostService) Edit(editReq SystemAuthPostEditReq) (e error) {
	var post system_model.SystemAuthPost
	err := service.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&post).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "部门不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	r := service.db.Where("(code = ? OR name = ?) AND id != ? AND is_delete = ?", editReq.Code, editReq.Name, editReq.ID, 0).Limit(1).Find(&system_model.SystemAuthPost{})
	if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该岗位已存在!")
	}
	// 更新
	response.Copy(&post, editReq)
	err = service.db.Model(&post).Select("*").Updates(post).Error
	e = response.CheckErr(err, "编辑失败")
	return
}

// Del 部门删除
func (service systemAuthPostService) Del(id uint) (e error) {
	var post system_model.SystemAuthPost
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&post).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "岗位不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待删除数据查找失败"); e != nil {
		return
	}
	r := service.db.Where("post_id = ? AND is_delete = ?", id, 0).Limit(1).Find(&system_model.SystemAuthAdmin{})
	if e = response.CheckErr(r.Error, "Del Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该岗位存在管理员,请先移除!")
	}

	err = service.db.Delete(&post).Error
	e = response.CheckErr(err, "Del Save err")
	return
}
