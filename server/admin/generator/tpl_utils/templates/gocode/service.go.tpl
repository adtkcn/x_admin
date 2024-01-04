package {{{ .ModuleName }}}

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"gorm.io/gorm"
)
import "x_admin/core"

type I{{{ title (toCamelCase .EntityName) }}}Service interface {
	List(page request.PageReq, listReq {{{ title (toCamelCase .EntityName) }}}ListReq) (res response.PageResp, e error)
	ListAll() (res []{{{ title (toCamelCase .EntityName) }}}Resp, e error)

	Detail(id int) (res {{{ title (toCamelCase .EntityName) }}}Resp, e error)
	Add(addReq {{{ title (toCamelCase .EntityName) }}}AddReq) (e error)
	Edit(editReq {{{ title (toCamelCase .EntityName) }}}EditReq) (e error)
	Del(id int) (e error)
}
var Service=New{{{ title (toCamelCase .EntityName) }}}Service()
//New{{{ title (toCamelCase .EntityName) }}}Service 初始化
func New{{{ title (toCamelCase .EntityName) }}}Service() *{{{ toCamelCase .EntityName }}}Service {
	db := core.GetDB()
	return &{{{ toCamelCase .EntityName }}}Service{db: db}
}

//{{{ toCamelCase .EntityName }}}Service {{{ .FunctionName }}}服务实现类
type {{{ toCamelCase .EntityName }}}Service struct {
	db *gorm.DB
}

//List {{{ .FunctionName }}}列表
func (service {{{ toCamelCase .EntityName }}}Service) List(page request.PageReq, listReq {{{ title (toCamelCase .EntityName) }}}ListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.{{{ title (toCamelCase .EntityName) }}}{})
	{{{- range .Columns }}}
	{{{- if .IsQuery }}}
	{{{- $queryOpr := index $.ModelOprMap .QueryType }}}
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
	{{{- if contains .AllFields "is_delete" }}}
	dbModel = dbModel.Where("is_delete = ?", 0)
	{{{- end }}}
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.{{{ title (toCamelCase .EntityName) }}}
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []{{{ title (toCamelCase .EntityName) }}}Resp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}
//ListAll {{{ .FunctionName }}}列表
func (service {{{ toCamelCase .EntityName }}}Service) ListAll() (res []{{{ title (toCamelCase .EntityName) }}}Resp, e error) {
	var objs []model.{{{ title (toCamelCase .EntityName) }}}
	
	err := service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

//Detail {{{ .FunctionName }}}详情
func (service {{{ toCamelCase .EntityName }}}Service) Detail(id int) (res {{{ title (toCamelCase .EntityName) }}}Resp, e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
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

//Add {{{ .FunctionName }}}新增
func (service {{{ toCamelCase .EntityName }}}Service) Add(addReq {{{ title (toCamelCase .EntityName) }}}AddReq) (e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	response.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

//Edit {{{ .FunctionName }}}编辑
func (service {{{ toCamelCase .EntityName }}}Service) Edit(editReq {{{ title (toCamelCase .EntityName) }}}EditReq) (e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", editReq.Id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

//Del {{{ .FunctionName }}}删除
func (service {{{ toCamelCase .EntityName }}}Service) Del(id int) (e error) {
	var obj model.{{{ title (toCamelCase .EntityName) }}}
	err := service.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
    // 删除
    {{{- if contains .AllFields "is_delete" }}}
    obj.IsDelete = 1
    err = service.db.Save(&obj).Error
    e = response.CheckErr(err, "Del Save err")
    {{{- else }}}
    err = service.db.Delete(&obj).Error
    e = response.CheckErr(err, "Del Delete err")
    {{{- end }}}
	return
}
