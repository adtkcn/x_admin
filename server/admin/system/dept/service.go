package dept

import (
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"

	"gorm.io/gorm"
)

type ISystemAuthDeptService interface {
	All() (res []SystemAuthDeptResp, e error)
	List(listReq SystemAuthDeptListReq) (mapList []interface{}, e error)
	Detail(id uint) (res SystemAuthDeptResp, e error)
	Add(addReq SystemAuthDeptAddReq) (e error)
	Edit(editReq SystemAuthDeptEditReq) (e error)
	Del(id uint) (e error)
}

var Service = NewSystemAuthDeptService()

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
func (deptSrv systemAuthDeptService) All() (res []SystemAuthDeptResp, e error) {
	var depts []system_model.SystemAuthDept
	err := deptSrv.db.Where("pid > ? AND is_delete = ?", 0, 0).Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []SystemAuthDeptResp{}
	response.Copy(&res, depts)
	return
}

// List 部门列表
func (deptSrv systemAuthDeptService) List(listReq SystemAuthDeptListReq) (deptResps []SystemAuthDeptResp, e error) {
	deptModel := deptSrv.db.Where("is_delete = ?", 0)
	if listReq.Name != "" {
		deptModel = deptModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.IsStop >= 0 {
		deptModel = deptModel.Where("is_stop = ?", listReq.IsStop)
	}
	var depts []system_model.SystemAuthDept
	err := deptModel.Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	// deptResps = []SystemAuthDeptResp{}
	response.Copy(&deptResps, depts)
	// mapList = util.ArrayUtil.ListToTree(
	// util.ConvertUtil.StructsToMaps(deptResps), "id", "pid", "children")
	return
}

// Detail 部门详情
func (deptSrv systemAuthDeptService) Detail(id uint) (res SystemAuthDeptResp, e error) {
	var dept system_model.SystemAuthDept
	err := deptSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dept).Error
	if e = response.CheckErrDBNotRecord(err, "部门已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, dept)
	return
}

// Add 部门新增
func (deptSrv systemAuthDeptService) Add(addReq SystemAuthDeptAddReq) (e error) {
	if addReq.Pid == 0 {
		r := deptSrv.db.Where("pid = ? AND is_delete = ?", 0, 0).Limit(1).Find(&system_model.SystemAuthDept{})
		if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
			return
		}
		if r.RowsAffected > 0 {
			return response.AssertArgumentError.Make("顶级部门只允许有一个!")
		}
	}
	var dept system_model.SystemAuthDept
	response.Copy(&dept, addReq)
	err := deptSrv.db.Create(&dept).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 部门编辑
func (deptSrv systemAuthDeptService) Edit(editReq SystemAuthDeptEditReq) (e error) {
	var dept system_model.SystemAuthDept
	err := deptSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&dept).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "部门不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	if dept.Pid == 0 && editReq.Pid > 0 {
		return response.AssertArgumentError.Make("顶级部门不能修改上级!")
	}
	if editReq.ID == editReq.Pid {
		return response.AssertArgumentError.Make("上级部门不能是自己!")
	}
	// 更新
	response.Copy(&dept, editReq)
	err = deptSrv.db.Model(&dept).Updates(dept).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 部门删除
func (deptSrv systemAuthDeptService) Del(id uint) (e error) {
	var dept system_model.SystemAuthDept
	err := deptSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dept).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "部门不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	if dept.Pid == 0 {
		return response.AssertArgumentError.Make("顶级部门不能删除!")
	}
	r := deptSrv.db.Where("pid = ? AND is_delete = ?", id, 0).Limit(1).Find(&system_model.SystemAuthDept{})
	if e = response.CheckErr(r.Error, "Del Find dept err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("请先删除子级部门!")
	}
	r = deptSrv.db.Where("dept_id = ? AND is_delete = ?", id, 0).Limit(1).Find(&system_model.SystemAuthAdmin{})
	if e = response.CheckErr(r.Error, "Del Find admin err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该部门已被管理员使用,请先移除!")
	}
	dept.IsDelete = 1
	err = deptSrv.db.Save(&dept).Error
	e = response.CheckErr(err, "Del Save err")
	return
}
