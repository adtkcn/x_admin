package system_service

import (
	"strings"
	"x_admin/app/model/system_model"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"gorm.io/gorm"
)

var PermService = NewSystemAuthPermService()

// NewSystemAuthPermService 初始化
func NewSystemAuthPermService() *systemAuthPermService {
	db := core.GetDB()
	return &systemAuthPermService{db: db}
}

// systemAuthPermService 系统权限服务实现类
type systemAuthPermService struct {
	db *gorm.DB
}

// SelectMenuIdsByRoleId 根据角色ID获取菜单ID
func (service systemAuthPermService) SelectMenuIdsByRoleId(roleId string) (menuIds []string, e error) {
	var role system_model.SystemAuthRole
	err := service.db.Where("id = ?", roleId).First(&role).Error
	if e = response.CheckErr(err, "角色不存在"); e != nil {
		return []string{}, e
	}
	var perms []system_model.SystemAuthPerm
	err = service.db.Where("role_id = ?", role.ID).Find(&perms).Error
	if e = response.CheckErr(err, "查询角色的菜单失败"); e != nil {
		return []string{}, e
	}
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	return
}

// SelectMenuIdsByRoleIds 根据多个角色ID获取菜单ID
func (service systemAuthPermService) SelectMenuIdsByRoleIds(roleIds []string) (menuIds []string, e error) {
	if len(roleIds) == 0 {
		return []string{}, nil
	}
	var perms []system_model.SystemAuthPerm
	err := service.db.Where("role_id in ?", roleIds).Find(&perms).Error
	if e = response.CheckErr(err, "查询角色的菜单失败"); e != nil {
		return []string{}, e
	}
	menuIdMap := make(map[string]bool)
	for _, perm := range perms {
		menuIdMap[perm.MenuId] = true
	}
	for menuId := range menuIdMap {
		menuIds = append(menuIds, menuId)
	}
	return
}

/**
 * 缓存用户权限(基于多个角色)并返回权限字符串
 * @param adminId 用户ID
 * @param roleIds 角色ID列表
 * @return 权限字符串
 */
func (service systemAuthPermService) CacheAdminPermsByRoleIds(adminId string, roleIds []string) (string, error) {
	if len(roleIds) == 0 {
		util.RedisUtil.HSet(config.AdminConfig.BackstageAdminPermsKey, adminId, "", 0)
		return "", nil
	}
	menuIds, err := service.SelectMenuIdsByRoleIds(roleIds)
	if err != nil {
		return "", err
	}
	if len(menuIds) == 0 {
		util.RedisUtil.HSet(config.AdminConfig.BackstageAdminPermsKey, adminId, "", 0)
		return "", nil
	}
	return service.CacheAdminPermsByMenuIds(adminId, menuIds)
}

/**
 * 缓存用户权限(基于多个菜单)并返回权限字符串
 * @param adminId 用户ID
 * @param menuIds 菜单ID列表
 * @return 权限字符串
 */
func (service systemAuthPermService) CacheAdminPermsByMenuIds(adminId string, menuIds []string) (string, error) {

	if len(menuIds) == 0 {
		util.RedisUtil.HSet(config.AdminConfig.BackstageAdminPermsKey, adminId, "", 0)
		return "", nil
	}
	var menus []system_model.SystemAuthMenu
	err := service.db.Where(
		"is_disable = ? and id in ? and menu_type in ?", 0, menuIds, []string{"C", "A"}).Order(
		"menu_sort, id").Find(&menus).Error
	if err != nil {
		return "", err
	}
	var permArray []string
	permMap := make(map[string]bool)
	for _, menu := range menus {
		if menu.Perms != "" {
			perms := strings.Split(menu.Perms, ",")
			for _, p := range perms {
				p = strings.Trim(p, " ")
				if p != "" && !permMap[p] {
					permMap[p] = true
					permArray = append(permArray, p)
				}
			}
		}
	}
	permsStr := strings.Join(permArray, ",")
	util.RedisUtil.HSet(config.AdminConfig.BackstageAdminPermsKey, adminId, permsStr, 0)
	return permsStr, nil
}

// GetAdminPerms 获取用户缓存的权限列表
func (service systemAuthPermService) GetAdminPerms(adminId string) ([]string, error) {
	permsStr := util.RedisUtil.HGet(config.AdminConfig.BackstageAdminPermsKey, adminId)
	if permsStr == "" {
		// 根据用户ID获取角色ID列表
		roleIds, err := AdminRoleService.GetRoleIdsByAdminId(adminId)
		if err != nil {
			return nil, err
		}
		// 缓存并直接获取返回值，避免二次 Redis 调用
		permsStr, err = service.CacheAdminPermsByRoleIds(adminId, roleIds)
		if err != nil {
			return nil, err
		}
	}
	if permsStr == "" {
		return []string{}, nil
	}
	return strings.Split(permsStr, ","), nil
}

// RemoveAdminPermsCache 移除用户权限缓存
func (service systemAuthPermService) RemoveAdminPermsCache(adminId string) {
	util.RedisUtil.HDel(config.AdminConfig.BackstageAdminPermsKey, adminId)
}

// BatchSaveByMenuIds 批量写入角色和菜单绑定
func (service systemAuthPermService) BatchSaveByMenuIds(roleId string, menuIds string, db *gorm.DB) (e error) {
	if menuIds == "" {
		return
	}
	if db == nil {
		db = service.db
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []system_model.SystemAuthPerm
		for _, menuId := range strings.Split(menuIds, ",") {

			perms = append(perms, system_model.SystemAuthPerm{RoleId: roleId, MenuId: menuId})
		}
		txErr := tx.Create(&perms).Error
		var te = response.CheckErr(txErr, "BatchSaveByMenuIds Create in tx err")
		return te
	})
	e = response.CheckErr(err, "BatchSaveByMenuIds Transaction err")
	return
}

// BatchDeleteByRoleId 批量删除角色菜单(根据角色ID)
func (service systemAuthPermService) BatchDeleteByRoleId(roleId string, db *gorm.DB) (e error) {
	if db == nil {
		db = service.db
	}
	err := db.Delete(&system_model.SystemAuthPerm{}, "role_id = ?", roleId).Error
	e = response.CheckErr(err, "BatchDeleteByRoleId Delete err")
	return
}

// BatchDeleteByMenuId 批量删除角色菜单(根据菜单ID)
func (service systemAuthPermService) BatchDeleteByMenuId(menuId string) (e error) {
	err := service.db.Delete(&system_model.SystemAuthPerm{}, "menu_id = ?", menuId).Error
	e = response.CheckErr(err, "BatchDeleteByMenuId Delete err")
	return
}
