package flow_service

import (
	"errors"

	"x_admin/app/model"
	"x_admin/app/schema/flow_schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var TemplateService = NewFlowTemplateService()

// NewFlowTemplateService 初始化
func NewFlowTemplateService() *flowTemplateService {
	db := core.GetDB()
	return &flowTemplateService{db: db}
}

// flowTemplateService 流程模板服务实现类
type flowTemplateService struct {
	db *gorm.DB
}

// List 流程模板列表
func (service flowTemplateService) List(page request.PageReq, listReq flow_schema.FlowTemplateListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.FlowTemplate{})
	if listReq.FlowName != "" {
		dbModel = dbModel.Where("flow_name like ?", "%"+listReq.FlowName+"%")
	}
	if listReq.FlowGroup > 0 {
		dbModel = dbModel.Where("flow_group = ?", listReq.FlowGroup)
	}
	if listReq.FlowRemark != "" {
		dbModel = dbModel.Where("flow_remark = ?", listReq.FlowRemark)
	}
	if listReq.FlowFormData != "" {
		dbModel = dbModel.Where("flow_form_data = ?", listReq.FlowFormData)
	}
	if listReq.FlowProcessData != "" {
		dbModel = dbModel.Where("flow_process_data = ?", listReq.FlowProcessData)
	}
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.FlowTemplate
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	result := []flow_schema.FlowTemplateResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 流程模板列表
func (service flowTemplateService) ListAll() (res []flow_schema.FlowTemplateResp, e error) {
	var modelList []model.FlowTemplate
	err := service.db.Find(&modelList).Error
	if e = response.CheckErr(err, "获取列表失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail 流程模板详情
func (service flowTemplateService) Detail(id string) (res flow_schema.FlowTemplateResp, e error) {
	var obj model.FlowTemplate
	err := service.db.Where("id = ?", id).First(&obj).Error
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, obj)
	return
}

// Add 流程模板新增
func (service flowTemplateService) Add(addReq flow_schema.FlowTemplateAddReq) (e error) {
	var obj model.FlowTemplate
	convert_util.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 流程模板编辑
func (service flowTemplateService) Edit(editReq flow_schema.FlowTemplateEditReq) (e error) {
	var obj model.FlowTemplate
	err := service.db.Where("id = ?", editReq.Id).First(&obj).Error
	// 校验
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	// 更新
	convert_util.Copy(&obj, editReq)
	err = service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "编辑失败")
	return
}

// Del 流程模板删除
func (service flowTemplateService) Del(id string) (e error) {
	result := service.db.Where("id = ?", id).Delete(&model.FlowTemplate{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("数据不存在")
	}
	return
}
