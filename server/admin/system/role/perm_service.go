package role

import (
	"strconv"
	"strings"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"
	"x_admin/util"

	"gorm.io/gorm"
)

type ISystemAuthPermService interface {
	SelectMenuIdsByRoleId(roleId uint) (menuIds []uint, e error)
	CacheRoleMenusByRoleId(roleId uint) (e error)
	BatchSaveByMenuIds(roleId uint, menuIds string, db *gorm.DB) (e error)
	BatchDeleteByRoleId(roleId uint, db *gorm.DB) (e error)
	BatchDeleteByMenuId(menuId uint) (e error)
}

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
func (service systemAuthPermService) SelectMenuIdsByRoleId(roleId uint) (menuIds []uint, e error) {
	var role system_model.SystemAuthRole
	err := service.db.Where("id = ? AND is_disable = ?", roleId, 0).Limit(1).First(&role).Error
	if e = response.CheckErr(err, "角色不存在"); e != nil {
		return []uint{}, e
	}
	var perms []system_model.SystemAuthPerm
	err = service.db.Where("role_id = ?", role.ID).Find(&perms).Error
	if e = response.CheckErr(err, "查询角色的菜单失败"); e != nil {
		return []uint{}, e
	}
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	return
}

// CacheRoleMenusByRoleId 缓存角色菜单
func (service systemAuthPermService) CacheRoleMenusByRoleId(roleId uint) (e error) {
	var perms []system_model.SystemAuthPerm
	err := service.db.Where("role_id = ?", roleId).Find(&perms).Error
	if e = response.CheckErr(err, "查询角色的菜单失败"); e != nil {
		return
	}
	var menuIds []uint
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
	util.RedisUtil.HSet(config.AdminConfig.BackstageRolesKey, strconv.FormatUint(uint64(roleId), 10), strings.Join(menuArray, ","), 0)
	return
}

// BatchSaveByMenuIds 批量写入角色菜单
func (service systemAuthPermService) BatchSaveByMenuIds(roleId uint, menuIds string, db *gorm.DB) (e error) {
	if menuIds == "" {
		return
	}
	if db == nil {
		db = service.db
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []system_model.SystemAuthPerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.ParseUint(menuIdStr, 10, 32)
			perms = append(perms, system_model.SystemAuthPerm{ID: util.ToolsUtil.MakeUuid(), RoleId: roleId, MenuId: uint(menuId)})
		}
		txErr := tx.Create(&perms).Error
		var te = response.CheckErr(txErr, "BatchSaveByMenuIds Create in tx err")
		return te
	})
	e = response.CheckErr(err, "BatchSaveByMenuIds Transaction err")
	return
}

// BatchDeleteByRoleId 批量删除角色菜单(根据角色ID)
func (service systemAuthPermService) BatchDeleteByRoleId(roleId uint, db *gorm.DB) (e error) {
	if db == nil {
		db = service.db
	}
	err := db.Delete(&system_model.SystemAuthPerm{}, "role_id = ?", roleId).Error
	e = response.CheckErr(err, "BatchDeleteByRoleId Delete err")
	return
}

// BatchDeleteByMenuId 批量删除角色菜单(根据菜单ID)
func (service systemAuthPermService) BatchDeleteByMenuId(menuId uint) (e error) {
	err := service.db.Delete(&system_model.SystemAuthPerm{}, "menu_id = ?", menuId).Error
	e = response.CheckErr(err, "BatchDeleteByMenuId Delete err")
	return
}
