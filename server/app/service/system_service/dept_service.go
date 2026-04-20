package system_service

import (
	"errors"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/system_schema"
	"x_admin/core"
	"x_admin/core/response"

	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var DeptService = NewSystemAuthDeptService()

// NewSystemAuthDeptService 初始化
func NewSystemAuthDeptService() *systemAuthDeptService {
	db := core.GetDB()
	return &systemAuthDeptService{db: db}
}

// systemAuthDeptService 系统部门服务实现类
type systemAuthDeptService struct {
	db *gorm.DB
}

// All 部门所有
func (service systemAuthDeptService) All() (res []system_schema.SystemAuthDeptResp, e error) {
	var depts []system_model.SystemAuthDept
	err := service.db.Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []system_schema.SystemAuthDeptResp{}
	convert_util.Copy(&res, depts)
	return
}

// List 部门列表
func (service systemAuthDeptService) List(listReq system_schema.SystemAuthDeptListReq) (deptResps []system_schema.SystemAuthDeptResp, e error) {
	deptModel := service.db
	if listReq.Name != "" {
		deptModel = deptModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.IsStop >= 0 {
		deptModel = deptModel.Where("is_stop = ?", listReq.IsStop)
	}
	var depts []system_model.SystemAuthDept
	err := deptModel.Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	convert_util.Copy(&deptResps, depts)
	return
}

// Detail 部门详情
func (service systemAuthDeptService) Detail(id string) (res system_schema.SystemAuthDeptResp, e error) {
	var dept system_model.SystemAuthDept
	err := service.db.Where("id = ?", id).First(&dept).Error
	if e = response.CheckDBNotRecord(err, "部门已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, dept)
	return
}

// Add 部门新增
func (service systemAuthDeptService) Add(addReq system_schema.SystemAuthDeptAddReq) (e error) {
	if addReq.Pid == "" {
		r := service.db.Where("pid = ?", "").Limit(1).Find(&system_model.SystemAuthDept{})
		if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
			return
		}
		if r.RowsAffected > 0 {
			return response.AssertArgumentError.SetMessage("顶级部门只允许有一个!")
		}
	}
	var dept system_model.SystemAuthDept
	convert_util.Copy(&dept, addReq)
	err := service.db.Create(&dept).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 部门编辑
func (service systemAuthDeptService) Edit(editReq system_schema.SystemAuthDeptEditReq) (e error) {
	var dept system_model.SystemAuthDept
	err := service.db.Where("id = ?", editReq.ID).First(&dept).Error
	// 校验
	if e = response.CheckDBNotRecord(err, "部门不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	if dept.Pid == "" && editReq.Pid != "" {
		return response.AssertArgumentError.SetMessage("顶级部门不能修改上级!")
	}
	if editReq.ID == editReq.Pid {
		return response.AssertArgumentError.SetMessage("上级部门不能是自己!")
	}
	// 更新
	convert_util.Copy(&dept, editReq)
	err = service.db.Model(&dept).Select("*").Updates(dept).Error
	e = response.CheckErr(err, "编辑失败")
	return
}

// Del 部门删除
func (service systemAuthDeptService) Del(id string) (e error) {
	// 检查是否有子部门
	if r := service.db.Where("pid = ?", id).Limit(1).Find(&system_model.SystemAuthDept{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("请先删除子级部门!")
	}

	// 检查是否有管理员使用
	if r := service.db.Where("dept_id = ?", id).Limit(1).Find(&system_model.SystemAuthAdmin{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该部门已被管理员使用,请先移除!")
	}

	result := service.db.Where("id = ?", id).Delete(&system_model.SystemAuthDept{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("部门不存在")
	}
	return
}
