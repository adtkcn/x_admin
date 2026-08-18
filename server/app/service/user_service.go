package service

import (
	"errors"
	"x_admin/app/model/user_model"
	"x_admin/app/schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var UserService = NewUserService()

func NewUserService() *userService {
	return &userService{db: core.GetDB()}
}

type userService struct {
	db *gorm.DB
}

func (s userService) List(page request.PageReq, listReq schema.UserListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	dbModel := s.db.Model(&user_model.User{})
	if listReq.Keyword.IsExistsAndNotNull() && listReq.Keyword.ValueOrZero() != "" {
		kw := "%" + listReq.Keyword.ValueOrZero() + "%"
		dbModel = dbModel.Where("email LIKE ? OR nickname LIKE ? OR phone LIKE ?", kw, kw, kw)
	}
	if listReq.Status.IsExistsAndNotNull() && listReq.Status.ValueOrZero() != "" {
		dbModel = dbModel.Where("status = ?", listReq.Status.ValueOrZero())
	}
	if listReq.CreateTimeStart.IsExistsAndNotNull() && listReq.CreateTimeStart.ValueOrZero() != "" {
		dbModel = dbModel.Where("create_time >= ?", listReq.CreateTimeStart.ValueOrZero())
	}
	if listReq.CreateTimeEnd.IsExistsAndNotNull() && listReq.CreateTimeEnd.ValueOrZero() != "" {
		dbModel = dbModel.Where("create_time <= ?", listReq.CreateTimeEnd.ValueOrZero())
	}
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	var modelList []user_model.User
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []schema.UserResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{PageNo: page.PageNo, PageSize: page.PageSize, Count: count, Lists: result}, nil
}

func (s userService) Detail(Id string) (res schema.UserResp, e error) {
	var obj user_model.User
	err := s.db.Where("id = ?", Id).First(&obj).Error
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "获取详情失败"); e != nil {
		return
	}
	convert_util.Copy(&res, obj)
	return
}

func (s userService) Edit(editReq schema.UserEditReq) (e error) {
	// 先确认记录存在（RowsAffected==0 不代表不存在）
	var exist int64
	if err := s.db.Model(&user_model.User{}).Where("id = ?", editReq.Id).Count(&exist).Error; err != nil {
		return response.CheckErr(err, "编辑失败")
	}
	if exist == 0 {
		return errors.New("数据不存在")
	}
	updates := map[string]any{}
	if editReq.Nickname.IsExistsAndNotNull() {
		updates["nickname"] = editReq.Nickname.ValueOrZero()
	}
	if editReq.Avatar.IsExistsAndNotNull() {
		updates["avatar"] = editReq.Avatar.ValueOrZero()
	}
	if editReq.Phone.IsExistsAndNotNull() {
		updates["phone"] = editReq.Phone.ValueOrZero()
	}
	if editReq.PhoneCode.IsExistsAndNotNull() {
		updates["phone_code"] = editReq.PhoneCode.ValueOrZero()
	}
	if editReq.Status.IsExistsAndNotNull() {
		updates["status"] = editReq.Status.ValueOrZero()
	}
	if len(updates) == 0 {
		return nil
	}
	result := s.db.Model(&user_model.User{}).Where("id = ?", editReq.Id).Updates(updates)
	if result.Error != nil {
		return response.CheckErr(result.Error, "编辑失败")
	}
	return
}

// Disable 禁用/启用用户（0正常 1禁用）
func (s userService) Disable(disableReq schema.UserDisableReq) (e error) {
	// 先确认记录存在（RowsAffected==0 不代表不存在）
	var exist int64
	if err := s.db.Model(&user_model.User{}).Where("id = ?", disableReq.Id).Count(&exist).Error; err != nil {
		return response.CheckErr(err, "操作失败")
	}
	if exist == 0 {
		return errors.New("数据不存在")
	}
	result := s.db.Model(&user_model.User{}).Where("id = ?", disableReq.Id).Update("status", disableReq.Status)
	if result.Error != nil {
		return response.CheckErr(result.Error, "操作失败")
	}
	return
}

// Kick 踢人下线（自增 token_version 使旧 token 失效）
func (s userService) Kick(Id string) (e error) {
	// 先确认记录存在（RowsAffected==0 不代表不存在）
	var exist int64
	if err := s.db.Model(&user_model.User{}).Where("id = ?", Id).Count(&exist).Error; err != nil {
		return response.CheckErr(err, "操作失败")
	}
	if exist == 0 {
		return errors.New("数据不存在")
	}
	result := s.db.Model(&user_model.User{}).Where("id = ?", Id).Update("token_version", gorm.Expr("token_version + 1"))
	if result.Error != nil {
		return response.CheckErr(result.Error, "操作失败")
	}
	return
}
