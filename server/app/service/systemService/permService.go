package systemService

import (
	"strings"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"
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
	err := service.db.Where("id = ? AND is_disable = ?", roleId, 0).Limit(1).First(&role).Error
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

// CacheRoleMenusByRoleId 缓存角色菜单
func (service systemAuthPermService) CacheRoleMenusByRoleId(roleId string) (e error) {
	var perms []system_model.SystemAuthPerm
	err := service.db.Where("role_id = ?", roleId).Find(&perms).Error
	if e = response.CheckErr(err, "查询角色的菜单失败"); e != nil {
		return
	}
	var menuIds []string
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	var menus []system_model.SystemAuthMenu
	err = service.db.Where(
		"is_disable = ? and id in ? and menu_type in ?", 0, menuIds, []string{"C", "A"}).Order(
		"menu_sort, id").Find(&menus).Error
	if e = response.CheckErr(err, "查找角色菜单失败"); e != nil {
		return
	}
	var menuArray []string
	for _, menu := range menus {
		if menu.Perms != "" {
			menuArray = append(menuArray, strings.Trim(menu.Perms, ""))
		}
	}
	util.RedisUtil.HSet(config.AdminConfig.BackstageRolesKey, roleId, strings.Join(menuArray, ","), 0)
	return
}

// BatchSaveByMenuIds 批量写入角色菜单
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

			perms = append(perms, system_model.SystemAuthPerm{ID: util.ToolsUtil.MakeUuidV7(), RoleId: roleId, MenuId: menuId})
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
