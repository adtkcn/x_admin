package menu

import (
	"x_admin/admin/system/role"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"
	"x_admin/util"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ISystemAuthMenuService interface {
	SelectMenuByRoleId(c *gin.Context, roleId uint) (mapList []interface{}, e error)
	List() (res []interface{}, e error)
	Detail(id uint) (res SystemAuthMenuResp, e error)
	Add(addReq SystemAuthMenuAddReq) (e error)
	Edit(editReq SystemAuthMenuEditReq) (e error)
	Del(id uint) (e error)
}

var Service = NewSystemAuthMenuService()

// NewSystemAuthMenuService 初始化
func NewSystemAuthMenuService() ISystemAuthMenuService {
	db := core.GetDB()
	return &systemAuthMenuService{db: db}
}

// systemAuthMenuService 系统菜单服务实现类
type systemAuthMenuService struct {
	db *gorm.DB
}

// SelectMenuByRoleId 根据角色ID获取菜单
func (menuSrv systemAuthMenuService) SelectMenuByRoleId(c *gin.Context, roleId uint) (mapList []interface{}, e error) {
	adminId := config.AdminConfig.GetAdminId(c)
	var menuIds []uint
	// 超管
	if adminId == config.AdminConfig.SuperAdminId {
		menuIds = []uint{0}
	} else if menuIds, e = role.PermService.SelectMenuIdsByRoleId(roleId); e != nil {
		return
	}
	if len(menuIds) == 0 {
		menuIds = []uint{0}
	}
	chain := menuSrv.db.Where("menu_type in ? AND is_disable = ?", []string{"M", "C"}, 0)
	if adminId != config.AdminConfig.SuperAdminId {
		chain = chain.Where("id in ?", menuIds)
	}
	var menus []system_model.SystemAuthMenu
	err := chain.Order("menu_sort desc, id").Find(&menus).Error
	if e = response.CheckErr(err, "SelectMenuByRoleId Find err"); e != nil {
		return
	}
	var menuResps []SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	mapList = util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	return
}

// List 菜单列表
func (menuSrv systemAuthMenuService) List() (res []interface{}, e error) {
	var menus []system_model.SystemAuthMenu
	err := menuSrv.db.Order("menu_sort desc, id").Find(&menus).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var menuResps []SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children"), nil
}

// Detail 菜单详情
func (menuSrv systemAuthMenuService) Detail(id uint) (res SystemAuthMenuResp, e error) {
	var menu system_model.SystemAuthMenu
	err := menuSrv.db.Where("id = ?", id).Limit(1).First(&menu).Error
	if e = response.CheckErrDBNotRecord(err, "菜单已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, menu)
	return
}

func (menuSrv systemAuthMenuService) Add(addReq SystemAuthMenuAddReq) (e error) {
	var menu system_model.SystemAuthMenu
	response.Copy(&menu, addReq)
	err := menuSrv.db.Create(&menu).Error
	if e = response.CheckErr(err, "Add Create err"); e != nil {
		return
	}
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
	return
}

func (menuSrv systemAuthMenuService) Edit(editReq SystemAuthMenuEditReq) (e error) {
	var menu system_model.SystemAuthMenu
	err := menuSrv.db.Where("id = ?", editReq.ID).Limit(1).Find(&menu).Error
	if e = response.CheckErrDBNotRecord(err, "菜单已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit Find err"); e != nil {
		return
	}
	response.Copy(&menu, editReq)
	err = menuSrv.db.Model(&menu).Updates(structs.Map(menu)).Error
	if e = response.CheckErr(err, "Edit Updates err"); e != nil {
		return
	}
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
	return
}

// Del 删除菜单
func (menuSrv systemAuthMenuService) Del(id uint) (e error) {
	var menu system_model.SystemAuthMenu
	err := menuSrv.db.Where("id = ?", id).Limit(1).First(&menu).Error
	if e = response.CheckErrDBNotRecord(err, "菜单已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Delete First err"); e != nil {
		return
	}
	r := menuSrv.db.Where("pid = ?", id).Limit(1).Find(&system_model.SystemAuthMenu{})
	err = r.Error
	if e = response.CheckErr(err, "Delete Find by pid err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("请先删除子菜单再操作！")
	}
	err = menuSrv.db.Delete(&menu).Error
	e = response.CheckErr(err, "Delete Delete err")
	return
}
