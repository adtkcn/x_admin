package {{{ .ModuleName }}}

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"gorm.io/gorm"
	"x_admin/util"
	"x_admin/util/excel2"
)

var {{{ toUpperCamelCase .EntityName }}}Service=New{{{ toUpperCamelCase .EntityName }}}Service()
var cacheUtil = util.CacheUtil{
	Name: {{{ toUpperCamelCase .EntityName }}}Service.Name,
}

// New{{{ toUpperCamelCase .EntityName }}}Service 初始化
func New{{{ toUpperCamelCase .EntityName }}}Service() *{{{ toCamelCase .EntityName }}}Service {
	return &{{{ toCamelCase .EntityName }}}Service{
		db:   core.GetDB(),
		Name: "{{{ toCamelCase .EntityName }}}",
	}
}

//{{{ toCamelCase .EntityName }}}Service {{{ .FunctionName }}}服务实现类
type {{{ toCamelCase .EntityName }}}Service struct {
	db *gorm.DB
	Name string
}



// List {{{ .FunctionName }}}列表
func (service {{{ toCamelCase .EntityName }}}Service) GetModel(listReq {{{ toUpperCamelCase .EntityName }}}ListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.{{{ toUpperCamelCase .EntityName }}}{})
	{{{- range .Columns }}}
	{{{- if .IsQuery }}}
	{{{- $queryOpr := index $.ModelOprMap .QueryType }}}
		{{{- if eq .HtmlType "datetime" }}}
			if listReq.{{{ toUpperCamelCase .ColumnName }}}Start!= nil {
				dbModel = dbModel.Where("{{{ .ColumnName }}} >= ?", *listReq.{{{ toUpperCamelCase .ColumnName }}}Start)
			}
			if listReq.{{{ toUpperCamelCase .ColumnName }}}End!= nil {
				dbModel = dbModel.Where("{{{ .ColumnName }}} <= ?", *listReq.{{{ toUpperCamelCase .ColumnName }}}End)
			}
		{{{- else }}}
			{{{- if and (eq .GoType "string") (eq $queryOpr "like") }}}
			if listReq.{{{ toUpperCamelCase .ColumnName }}}!= nil {
				dbModel = dbModel.Where("{{{ .ColumnName }}} like ?", "%"+*listReq.{{{ toUpperCamelCase .ColumnName }}}+"%")
			}
			{{{- else }}}
			if listReq.{{{ toUpperCamelCase .ColumnName }}}!= nil {
				dbModel = dbModel.Where("{{{ .ColumnName }}} = ?", *listReq.{{{ toUpperCamelCase .ColumnName }}})
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
func (service {{{ toCamelCase .EntityName }}}Service) List(page request.PageReq, listReq {{{ toUpperCamelCase .EntityName }}}ListReq) (res response.PageResp, e error) {
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
	var modelList []model.{{{ toUpperCamelCase .EntityName }}}
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []{{{ toUpperCamelCase .EntityName }}}Resp{}
	util.ConvertUtil.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}
// ListAll {{{ .FunctionName }}}列表
func (service {{{ toCamelCase .EntityName }}}Service) ListAll(listReq {{{ toUpperCamelCase .EntityName }}}ListReq) (res []{{{ toUpperCamelCase .EntityName }}}Resp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.{{{ toUpperCamelCase .EntityName }}}

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	util.ConvertUtil.Copy(&res, modelList)
	return res, nil
}

// Detail {{{ .FunctionName }}}详情
func (service {{{ toCamelCase .EntityName }}}Service) Detail({{{ toUpperCamelCase .PrimaryKey }}} int) (res {{{ toUpperCamelCase .EntityName }}}Resp, e error) {
	var obj = model.{{{ toUpperCamelCase .EntityName }}}{}
	err := cacheUtil.GetCache({{{ toUpperCamelCase .PrimaryKey }}}, &obj)
	if err != nil {
		err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", {{{ toUpperCamelCase .PrimaryKey }}}{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
	
		{{{- range .Columns }}}
		{{{- if and .IsEdit (contains (slice "image" "avatar" "logo" "img") .GoField) }}}
		res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
		{{{- end }}}
		{{{- end }}}
		cacheUtil.SetCache(obj.{{{ toUpperCamelCase .PrimaryKey }}}, obj)
	}

	util.ConvertUtil.Copy(&res, obj)
	return
}

// Add {{{ .FunctionName }}}新增
func (service {{{ toCamelCase .EntityName }}}Service) Add(addReq {{{ toUpperCamelCase .EntityName }}}AddReq) (createId int,e error) {
	var obj model.{{{ toUpperCamelCase .EntityName }}}
	util.ConvertUtil.StructToStruct(addReq,&obj)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return 0,e
	}
	cacheUtil.SetCache(obj.{{{ toUpperCamelCase .PrimaryKey }}}, obj)
	createId = obj.{{{ toUpperCamelCase .PrimaryKey }}}
	return
}

// Edit {{{ .FunctionName }}}编辑
func (service {{{ toCamelCase .EntityName }}}Service) Edit(editReq {{{ toUpperCamelCase .EntityName }}}EditReq) (e error) {
	var obj model.{{{ toUpperCamelCase .EntityName }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", editReq.{{{ toUpperCamelCase .PrimaryKey }}}{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	util.ConvertUtil.Copy(&obj, editReq)

	err = service.db.Model(&obj).Select("*").Updates(obj).Error
	if e = response.CheckErr(err, "编辑失败"); e != nil {
		return
	}
	cacheUtil.RemoveCache(obj.Id)
	service.Detail(obj.Id)
	return
}

// Del {{{ .FunctionName }}}删除
func (service {{{ toCamelCase .EntityName }}}Service) Del({{{ toUpperCamelCase .PrimaryKey }}} int) (e error) {
	var obj model.{{{ toUpperCamelCase .EntityName }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", {{{ toUpperCamelCase .PrimaryKey }}}{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
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
		obj.DeleteTime = util.NullTimeUtil.Now()
		{{{- end }}}
    err = service.db.Save(&obj).Error
    e = response.CheckErr(err, "删除失败")
    {{{- else }}}
    err = service.db.Delete(&obj).Error
    e = response.CheckErr(err, "删除失败")
    {{{- end }}}
	cacheUtil.RemoveCache(obj.{{{ toUpperCamelCase .PrimaryKey }}})
	return
}

// DelBatch 用户协议-批量删除
func (service {{{ toCamelCase .EntityName }}}Service) DelBatch(Ids []string) (e error) {
	var obj model.{{{ toUpperCamelCase .EntityName }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} in (?)", Ids).Delete(&obj).Error
	if err != nil {
		return err
	}
	// 删除缓存
	cacheUtil.RemoveCache(Ids)
	return nil
}

// 获取Excel的列
func (service {{{ toCamelCase .EntityName }}}Service) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
	{{{- range .Columns }}}
	{{{- if and (.IsList) (not .IsPk) }}}
		{{{- if eq .HtmlType "datetime" }}}
		{Name: "{{{.ColumnComment}}}", Key: "{{{ toUpperCamelCase .GoField }}}", Width: 15, Decode: util.NullTimeUtil.DecodeTime },
		{{{- else if eq .GoType "int" }}}
			{Name: "{{{.ColumnComment}}}", Key: "{{{ toUpperCamelCase .GoField }}}", Width: 15, Decode: core.DecodeInt},
		{{{- else }}}
		{Name: "{{{.ColumnComment}}}", Key: "{{{ toUpperCamelCase .GoField }}}", Width: 15},
		{{{- end }}}
	{{{- end }}}
	{{{- end }}}
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile {{{ .FunctionName }}}导出
func (service {{{ toCamelCase .EntityName }}}Service) ExportFile(listReq {{{ toUpperCamelCase .EntityName }}}ListReq) (res []{{{ toUpperCamelCase .EntityName }}}Resp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.{{{ toUpperCamelCase .EntityName }}}
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []{{{ toUpperCamelCase .EntityName }}}Resp{}
	util.ConvertUtil.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service {{{ toCamelCase .EntityName }}}Service) ImportFile(importReq []{{{ toUpperCamelCase .EntityName }}}Resp) (e error) {
	var importData []model.{{{ toUpperCamelCase .EntityName }}}
	util.ConvertUtil.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}