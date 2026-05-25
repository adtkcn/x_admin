package system_service

import (
	"x_admin/app/model/system_model"
	"x_admin/core"

	"gorm.io/gorm"
)

var AdminRoleService = NewAdminRoleService()

func NewAdminRoleService() *adminRoleService {
	db := core.GetDB()
	return &adminRoleService{db: db}
}

type adminRoleService struct {
	db *gorm.DB
}

// GetRoleIdsByAdminId 根据用户ID获取角色ID列表
func (s *adminRoleService) GetRoleIdsByAdminId(adminId string) ([]string, error) {
	var adminRoles []system_model.SystemAuthAdminRole
	err := s.db.Where("admin_id = ?", adminId).Find(&adminRoles).Error
	if err != nil {
		return nil, err
	}
	var roleIds []string
	for _, ar := range adminRoles {
		roleIds = append(roleIds, ar.RoleId)
	}
	return roleIds, nil
}

// GetRolesByAdminId 根据用户ID获取角色列表
func (s *adminRoleService) GetRolesByAdminId(adminId string) ([]system_model.SystemAuthRole, error) {
	// 获取角色id列表
	roleIds, err := s.GetRoleIdsByAdminId(adminId)
	if err != nil {
		return nil, err
	}
	if len(roleIds) == 0 {
		return []system_model.SystemAuthRole{}, nil
	}
	var roles []system_model.SystemAuthRole
	err = s.db.Where("id in ?", roleIds).Find(&roles).Error
	return roles, err
}

// SaveAdminRoles 保存用户角色关联
func (s *adminRoleService) SaveAdminRoles(adminId string, roleIds []string, tx *gorm.DB) error {
	if tx == nil {
		tx = s.db
	}
	return tx.Transaction(func(t *gorm.DB) error {
		if err := t.Where("admin_id = ?", adminId).Delete(&system_model.SystemAuthAdminRole{}).Error; err != nil {
			return err
		}
		if len(roleIds) == 0 {
			return nil
		}
		var adminRoles []system_model.SystemAuthAdminRole
		for _, roleId := range roleIds {
			adminRoles = append(adminRoles, system_model.SystemAuthAdminRole{
				AdminId: adminId,
				RoleId:  roleId,
			})
		}
		return t.Create(&adminRoles).Error
	})
}

// DeleteByAdminId 删除用户所有角色关联
func (s *adminRoleService) DeleteByAdminId(adminId string) error {
	return s.db.Where("admin_id = ?", adminId).Delete(&system_model.SystemAuthAdminRole{}).Error
}

// DeleteByRoleId 删除角色所有用户关联
func (s *adminRoleService) DeleteByRoleId(roleId string) error {
	return s.db.Where("role_id = ?", roleId).Delete(&system_model.SystemAuthAdminRole{}).Error
}
