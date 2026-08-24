package {{{.Domain}}}_service

import (
	"errors"
	"x_admin/app/schema/{{{.Domain}}}_schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/app/model/{{{.Domain}}}_model"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
)

var {{{ toUpperCamelCase .EntityName }}}Service=New{{{ toUpperCamelCase .EntityName }}}Service()

// New{{{ toUpperCamelCase .EntityName }}}Service 初始化
func New{{{ toUpperCamelCase .EntityName }}}Service() *{{{ .EntityName }}}Service {
	return &{{{ .EntityName }}}Service{
		db:   core.GetDB(),
		CacheUtil: util.CacheUtil{
			Name: "{{{ .EntityName }}}",
		},
	}
}

//{{{ .EntityName }}}Service {{{ .FunctionName }}}服务实现类
type {{{ .EntityName }}}Service struct {
	db *gorm.DB
	CacheUtil util.CacheUtil
}



// List {{{ .FunctionName }}}列表
func (service {{{ .EntityName }}}Service) GetModel(listReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}ListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&{{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}{}).Joins("CreatedByUser")
	tableName := core.DBTableName(&{{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}{})
	{{{- range .Columns }}}
	{{{- if .IsQuery }}}
		{{{- $queryOpr := index $.ModelOprMap .QueryType }}}
			{{{- if eq .ColumnName "created_by" }}}
	if listReq.CreatedBy.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".created_by = ?", listReq.CreatedBy.ValueOrZero())
	}
	if listReq.CreatedByNickname.IsExistsAndNotNull() {
		dbModel = dbModel.Where("CreatedByUser.nickname like ?", "%"+listReq.CreatedByNickname.ValueOrZero()+"%")
	}
	if listReq.CreatedByUsername.IsExistsAndNotNull() {
		dbModel = dbModel.Where("CreatedByUser.username like ?", "%"+listReq.CreatedByUsername.ValueOrZero()+"%")
	}
			{{{- else if eq .HtmlType "datetime" }}}
	if listReq.{{{ .GoField }}}Start.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".{{{ .ColumnName }}} >= ?", listReq.{{{ .GoField }}}Start.ValueOrZero())
	}
	if listReq.{{{ .GoField }}}End.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".{{{ .ColumnName }}} <= ?", listReq.{{{ .GoField }}}End.ValueOrZero())
	}
			{{{- else }}}
			{{{- if and (eq .GoType "string") (eq $queryOpr "like") }}}
	if listReq.{{{ .GoField }}}.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".{{{ .ColumnName }}} like ?", "%"+listReq.{{{ .GoField }}}.ValueOrZero()+"%")
	}
			{{{- else }}}
	if listReq.{{{ .GoField }}}.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".{{{ .ColumnName }}} = ?", listReq.{{{ .GoField }}}.ValueOrZero())
	}
			{{{- end }}}
		{{{- end }}}
	{{{- end }}}
    {{{- end }}}
	return dbModel
}
// 获取更新map
func (service {{{ .EntityName }}}Service) GetUpdateMap(editReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}EditReq) map[string]any {
	updateMap := make(map[string]any)
	{{{- range .Columns }}}
	{{{- if .IsEdit }}}
	{{{- if .IsPk }}}
	if editReq.{{{ .GoField }}} !="" {
		updateMap["{{{ .ColumnName }}}"] = editReq.{{{ .GoField }}}
	}
    {{{- else }}}
	if editReq.{{{ .GoField }}}.IsExists() {
		updateMap["{{{ .ColumnName }}}"] = editReq.{{{ .GoField }}}.GetValue()
	}
	{{{- end }}}
	{{{- end }}}
	{{{- end }}}
	return updateMap
}
// List {{{ .FunctionName }}}列表
func (service {{{ .EntityName }}}Service) List(page request.PageReq, listReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}ListReq) (res response.PageResp, e error) {
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
	var modelList []{{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []{{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}
// ListAll {{{ .FunctionName }}}列表
func (service {{{ .EntityName }}}Service) ListAll(listReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}ListReq) (res []{{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []{{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail {{{ .FunctionName }}}详情
func (service {{{ .EntityName }}}Service) Detail({{{ .PrimaryGoField }}} {{{.PrimaryGoType}}}) (res {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp, e error) {
	var obj = {{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}{}
	err := service.CacheUtil.GetCache({{{ .PrimaryGoField }}}, &obj)
	if err != nil {
		err := service.db.Where("{{{ $.PrimaryKey }}} = ?", {{{ .PrimaryGoField }}}).Preload("CreatedByUser").First(&obj).Error
		if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
		service.CacheUtil.SetCache(obj.{{{ .PrimaryGoField }}}, obj)
	}

	convert_util.Copy(&res, obj)
	return
}

// Add {{{ .FunctionName }}}新增
func (service {{{ .EntityName }}}Service) Add(addReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}AddReq, adminId string) (createId {{{.PrimaryGoType}}},e error) {
	var obj {{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}
	convert_util.Copy(&obj, addReq)

 
	{{{- range .Columns }}}
	{{{- if and .IsEdit (eq "CreatedBy" .GoField) }}}
	obj.CreatedBy.SetValue(adminId)
	{{{- end }}}
	{{{- end }}}
	
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return "",e
	}
	service.CacheUtil.SetCache(obj.{{{ .PrimaryGoField }}}, obj)
	createId = obj.{{{ .PrimaryGoField }}}
	return
}

// Edit {{{ .FunctionName }}}编辑
func (service {{{ .EntityName }}}Service) Edit(editReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}EditReq) (e error) {
	var obj {{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?", editReq.{{{ .PrimaryGoField }}}).First(&obj).Error
	// 校验
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}

	updateMap := service.GetUpdateMap(editReq)
	if len(updateMap) == 0 {
		return errors.New("没有可更新的字段")
	}

	err = service.db.Model(&obj).Updates(updateMap).Error
	if e = response.CheckErr(err, "编辑失败"); e != nil {
		return
	}
	service.CacheUtil.RemoveCache(obj.{{{.PrimaryGoField }}})

	return
}

// Del {{{ .FunctionName }}}删除
func (service {{{ .EntityName }}}Service) Del({{{ .PrimaryGoField }}} {{{.PrimaryGoType}}}) (e error) {
	result := service.db.Where("{{{ $.PrimaryKey }}} = ?", {{{ .PrimaryGoField }}}).Delete(&{{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("数据不存在")
	}
	service.CacheUtil.RemoveCache({{{ .PrimaryGoField }}})

	return
}

// DelBatch {{{ .FunctionName }}}-批量删除
func (service {{{ .EntityName }}}Service) DelBatch(Ids []string) (e error) {
	var obj {{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} in (?)", Ids).Delete(&obj).Error
	if err != nil {
		return err
	}
	// 删除缓存
	service.CacheUtil.RemoveCache(Ids...)
	return nil
}

// 获取Excel的列
func (service {{{ .EntityName }}}Service) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
	{{{- range .Columns }}}
	{{{- if and (.IsList) (not .IsPk) }}}
		{{{- if eq .HtmlType "datetime" }}}
	{Name: "{{{.ColumnComment}}}", Key: "{{{ .GoField }}}", Width: 15, Decode: x_null.DecodeTime },
		{{{- else if eq .GoType "int" }}}
	{Name: "{{{.ColumnComment}}}", Key: "{{{ .GoField }}}", Width: 15, Decode: x_null.DecodeInt64},
		{{{- else if eq .GoType "float64" }}}
	{Name: "{{{.ColumnComment}}}", Key: "{{{ .GoField }}}", Width: 15, Decode: x_null.DecodeFloat64},
		{{{- else }}}
	{Name: "{{{.ColumnComment}}}", Key: "{{{ .GoField }}}", Width: 15, Decode: x_null.DecodeString},
		{{{- end }}}
	{{{- end }}}
	{{{- end }}}
	}
	return cols
}

// ExportFile {{{ .FunctionName }}}导出
func (service {{{ .EntityName }}}Service) ExportFile(listReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}ListReq) (res []{{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []{{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []{{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// {{{ .FunctionName }}}导入
func (service {{{ .EntityName }}}Service) ImportFile(importReq []{{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp) (e error) {
	var importData []{{{.Domain}}}_model.{{{ toUpperCamelCase .EntityName }}}
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}