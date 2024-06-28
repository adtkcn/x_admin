package {{{ .ModuleName }}}

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"gorm.io/gorm"
)

var Service=New{{{ title (toCamelCase .EntityName) }}}Service()
// New{{{ title (toCamelCase .EntityName) }}}Service 初始化
func New{{{ title (toCamelCase .EntityName) }}}Service() *{{{ toCamelCase .EntityName }}}Service {
	db := core.GetDB()
	return &{{{ toCamelCase .EntityName }}}Service{db: db}
}

//{{{ toCamelCase .EntityName }}}Service {{{ .FunctionName }}}服务实现类
type {{{ toCamelCase .EntityName }}}Service struct {
	db *gorm.DB
}

// List {{{ .FunctionName }}}列表
func (service {{{ toCamelCase .EntityName }}}Service) GetModel(listReq {{{ title (toCamelCase .EntityName) }}}ListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.{{{ title (toCamelCase .EntityName) }}}{})
	{{{- range .Columns }}}
	{{{- if .IsQuery }}}
	{{{- $queryOpr := index $.ModelOprMap .QueryType }}}
		{{{- if eq .HtmlType "datetime" }}}
			if listReq.{{{ title (toCamelCase .ColumnName) }}}Start != "" {
				dbModel = dbModel.Where("{{{ .ColumnName }}} >= ?", listReq.{{{ title (toCamelCase .ColumnName) }}}Start)
			}
			if listReq.{{{ title (toCamelCase .ColumnName) }}}End != "" {
				dbModel = dbModel.Where("{{{ .ColumnName }}} <= ?", listReq.{{{ title (toCamelCase .ColumnName) }}}End)
			}
		{{{- else }}}
			{{{- if and (eq .GoType "string") (eq $queryOpr "like") }}}
			if listReq.{{{ title (toCamelCase .ColumnName) }}} != "" {
				dbModel = dbModel.Where("{{{ .ColumnName }}} like ?", "%"+listReq.{{{ title (toCamelCase .ColumnName) }}}+"%")
			}
			{{{- else }}}
			if listReq.{{{ title (toCamelCase .ColumnName) }}} {{{ if eq .GoType "string" }}}!= ""{{{ else }}}> 0{{{ end }}} {
				dbModel = dbModel.Where("{{{ .ColumnName }}} = ?", listReq.{{{ title (toCamelCase .ColumnName) }}})
			}
			{{{- end }}}
		{{{- end }}}
    {{{- end }}}
    {{{- end }}}
	{{{- if contains .AllFields "is_delete" }}}
	dbModel = dbModel.Where("is_delete = ?", 0)
	{{{- end }}}
	return dbModel
}
// List {{{ .FunctionName }}}列表
func (service {{{ toCamelCase .EntityName }}}Service) List(page request.PageReq, listReq {{{ title (toCamelCase .EntityName) }}}ListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	dbModel := service.GetModel(listReq)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.{{{ title (toCamelCase .EntityName) }}}
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []{{{ title (toCamelCase .EntityName) }}}Resp{}
	response.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}
// ListAll {{{ .FunctionName }}}列表
func (service {{{ toCamelCase .EntityName }}}Service) ListAll(listReq {{{ title (toCamelCase .EntityName) }}}ListReq) (res []{{{ title (toCamelCase .EntityName) }}}Resp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.{{{ title (toCamelCase .EntityName) }}}

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	response.Copy(&res, modelList)
	return res, nil
}

// Detail {{{ .FunctionName }}}详情
func (service {{{ toCamelCase .EntityName }}}Service) Detail(id int) (res {{{ title (toCamelCase .EntityName) }}}Resp, e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "获取详情失败"); e != nil {
		return
	}
	response.Copy(&res, obj)
	{{{- range .Columns }}}
    {{{- if and .IsEdit (contains (slice "image" "avatar" "logo" "img") .GoField) }}}
    res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
    {{{- end }}}
    {{{- end }}}
	return
}

// Add {{{ .FunctionName }}}新增
func (service {{{ toCamelCase .EntityName }}}Service) Add(addReq {{{ title (toCamelCase .EntityName) }}}AddReq) (e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	response.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return e
	}
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit {{{ .FunctionName }}}编辑
func (service {{{ toCamelCase .EntityName }}}Service) Edit(editReq {{{ title (toCamelCase .EntityName) }}}EditReq) (e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", editReq.Id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = service.db.Model(&obj).Select("*").Updates(obj).Error
	e = response.CheckErr(err, "更新失败")
	return
}

// Del {{{ .FunctionName }}}删除
func (service {{{ toCamelCase .EntityName }}}Service) Del(id int) (e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询数据失败"); e != nil {
		return
	}
    // 删除
    {{{- if contains .AllFields "is_delete" }}}
    obj.IsDelete = 1
	{{{- if contains .AllFields "delete_time" }}}
	obj.DeleteTime = core.NowTime()
	{{{- end }}}
    err = service.db.Save(&obj).Error
    e = response.CheckErr(err, "删除失败")
    {{{- else }}}
    err = service.db.Delete(&obj).Error
    e = response.CheckErr(err, "删除失败")
    {{{- end }}}
	return
}

// ExportFile {{{ .FunctionName }}}导出
func (service {{{ toCamelCase .EntityName }}}Service) ExportFile(listReq {{{ title (toCamelCase .EntityName) }}}ListReq) (res []{{{ title (toCamelCase .EntityName) }}}Resp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.{{{ title (toCamelCase .EntityName) }}}
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []{{{ title (toCamelCase .EntityName) }}}Resp{}
	response.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service {{{ toCamelCase .EntityName }}}Service) ImportFile(importReq []{{{ title (toCamelCase .EntityName) }}}Resp) (e error) {
	var importData []model.{{{ title (toCamelCase .EntityName) }}}
	response.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}