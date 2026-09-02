package system_service

import (
	"errors"
	"strings"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/system_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var RoleService = NewSystemAuthRoleService()

// NewSystemAuthRoleService 初始化
func NewSystemAuthRoleService() *systemAuthRoleService {
	db := core.GetDB()
	return &systemAuthRoleService{db: db}
}

// systemAuthRoleService 系统角色服务实现类
type systemAuthRoleService struct {
	db *gorm.DB
}

// All 角色所有
func (roleSrv systemAuthRoleService) All() (res []system_schema.SystemAuthRoleSimpleResp, e error) {
	var roles []system_model.SystemAuthRole
	err := roleSrv.db.Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	convert_util.Copy(&res, roles)
	return
}

// List 根据角色ID获取菜单ID
func (roleSrv systemAuthRoleService) List(page request.PageReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	roleModel := roleSrv.db.Model(&system_model.SystemAuthRole{})
	var count int64
	err := roleModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	var roles []system_model.SystemAuthRole
	err = roleSrv.db.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}

	// 批量获取成员数量
	var roleIds []string
	for _, role := range roles {
		roleIds = append(roleIds, role.ID)
	}
	memberCountMap := make(map[string]int64)
	if len(roleIds) > 0 {
		type MemberCount struct {
			RoleId string
			Count  int64
		}
		var memberCounts []MemberCount
		roleSrv.db.Model(&system_model.SystemAuthAdminRole{}).
			Select("role_id, count(*) as count").
			Where("role_id IN ?", roleIds).
			Group("role_id").
			Find(&memberCounts)
		for _, mc := range memberCounts {
			memberCountMap[mc.RoleId] = mc.Count
		}
	}

	var roleResp []system_schema.SystemAuthRoleResp
	convert_util.Copy(&roleResp, roles)
	for i := 0; i < len(roleResp); i++ {
		roleResp[i].Menus = []string{}
		roleResp[i].Member = memberCountMap[roleResp[i].ID]
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    roleResp,
	}, nil
}

// Detail 角色详情
func (roleSrv systemAuthRoleService) Detail(id string) (res system_schema.SystemAuthRoleResp, e error) {
	var role system_model.SystemAuthRole
	err := roleSrv.db.Where("id = ?", id).First(&role).Error
	if e = response.CheckDBErr(err, "角色已不存在!", "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, role)
	res.Member = roleSrv.getMemberCnt(role.ID)
	res.Menus, _ = PermService.SelectMenuIdsByRoleId(role.ID)
	return
}

// getMemberCnt 根据角色ID获取成员数量
func (roleSrv systemAuthRoleService) getMemberCnt(roleId string) (count int64) {
	roleSrv.db.Model(&system_model.SystemAuthAdminRole{}).Where(
		"role_id = ?", roleId).Count(&count)
	return
}

// Add 新增角色
func (roleSrv systemAuthRoleService) Add(addReq system_schema.SystemAuthRoleAddReq) (e error) {
	var role system_model.SystemAuthRole
	if r := roleSrv.db.Where("name = ?", strings.Trim(addReq.Name, " ")).First(&role); r.RowsAffected > 0 {
		return errors.New("角色名称已存在!")
	}
	convert_util.Copy(&role, addReq)
	role.Name = strings.Trim(addReq.Name, " ")
	// 事务
	err := roleSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Create(&role).Error
		var te error
		if te = response.CheckErr(txErr, "Add Create in tx err"); te != nil {
			return te
		}
		te = PermService.BatchSaveByMenuIds(role.ID, addReq.MenuIds, tx)
		return te
	})
	e = response.CheckErr(err, "Add Transaction err")
	return
}

// Edit 编辑角色
func (roleSrv systemAuthRoleService) Edit(editReq system_schema.SystemAuthRoleEditReq) (e error) {
	err := roleSrv.db.Where("id = ?", editReq.ID).First(&system_model.SystemAuthRole{}).Error
	if e = response.CheckDBErr(err, "角色已不存在!", "待编辑数据查找失败"); e != nil {
		return
	}
	var role system_model.SystemAuthRole
	if r := roleSrv.db.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).First(&role); r.RowsAffected > 0 {
		return errors.New("角色名称已存在!")
	}
	role.ID = editReq.ID
	roleMap := map[string]interface{}{
		"Name":      strings.Trim(editReq.Name, " "),
		"Sort":      editReq.Sort,
		"IsDisable": editReq.IsDisable,
		"Remark":    editReq.Remark,
	}
	// 事务
	err = roleSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&role).Updates(roleMap).Error
		var te error
		if te = response.CheckErr(txErr, "编辑角色失败"); te != nil {
			return te
		}
		// 删除角色的菜单
		if te = PermService.BatchDeleteByRoleId(editReq.ID, tx); te != nil {
			return te
		}
		// 重新保存角色的菜单
		if te = PermService.BatchSaveByMenuIds(editReq.ID, editReq.MenuIds, tx); te != nil {
			return te
		}

		// 清空redis角色权限缓存(按前缀批量删除所有管理员权限键)
		util.RedisUtil.DelByPrefix(config.AdminConfig.BackstageAdminPermsKey + ":")
		return nil
	})
	e = response.CheckErr(err, "编辑角色失败")
	return
}

// Del 删除角色
func (roleSrv systemAuthRoleService) Del(id string) (e error) {
	if r := roleSrv.db.Where("role_id = ?", id).Limit(1).Find(&system_model.SystemAuthAdminRole{}); r.RowsAffected > 0 {
		return errors.New("角色已被管理员使用,请先移除!")
	}

	err := roleSrv.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&system_model.SystemAuthRole{}, "id = ?", id)
		if result.Error != nil {
			return response.CheckMysqlErr(result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.New("角色已不存在")
		}

		if te := PermService.BatchDeleteByRoleId(id, tx); te != nil {
			return te
		}
		// 清空redis角色权限缓存(按前缀批量删除所有管理员权限键)
		util.RedisUtil.DelByPrefix(config.AdminConfig.BackstageAdminPermsKey + ":")

		return nil
	})

	return response.CheckErr(err, "删除失败")
}
