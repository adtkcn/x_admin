package systemService

import (
	"errors"
	"x_admin/app/schema/systemSchema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"
	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var MenuService = NewSystemAuthMenuService()

// NewSystemAuthMenuService 初始化
func NewSystemAuthMenuService() *systemAuthMenuService {
	db := core.GetDB()
	return &systemAuthMenuService{db: db}
}

// systemAuthMenuService 系统菜单服务实现类
type systemAuthMenuService struct {
	db *gorm.DB
}

// SelectMenuByAdminId 根据管理员ID获取菜单
func (menuSrv systemAuthMenuService) SelectMenuByAdminId(adminId string) (menuList []*systemSchema.SystemAuthMenuResp, e error) {
	var menuIds = []string{}
	if adminId != config.AdminConfig.SuperAdminId {
		roleIds, err := AdminRoleService.GetRoleIdsByAdminId(adminId)
		if err != nil {
			return nil, err
		}
		if len(roleIds) == 0 {
			return menuList, errors.New("用户未绑定角色")
		}
		if menuIds, e = PermService.SelectMenuIdsByRoleIds(roleIds); e != nil {
			return
		}
		if len(menuIds) == 0 {
			return menuList, errors.New("角色未绑定菜单")
		}
	}

	chain := menuSrv.db.Where("menu_type in ? AND is_disable = ?", []string{"M", "C"}, 0)
	if adminId != config.AdminConfig.SuperAdminId {
		chain = chain.Where("id in ?", menuIds)
	}
	var menus []system_model.SystemAuthMenu
	err := chain.Order("menu_sort desc, id").Find(&menus).Error
	if e = response.CheckErr(err, "查询菜单失败"); e != nil {
		return
	}
	var menuResps []*systemSchema.SystemAuthMenuResp
	convert_util.Copy(&menuResps, menus)

	trees := util.ListToTree(menuResps, "0")

	return trees, nil
}

// List 菜单列表
func (menuSrv systemAuthMenuService) List() (res interface{}, e error) {
	var menus []system_model.SystemAuthMenu
	err := menuSrv.db.Order("menu_sort desc, id").Find(&menus).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	var menuResps []systemSchema.SystemAuthMenuResp
	convert_util.Copy(&menuResps, menus)
	return menuResps, nil
}

// Detail 菜单详情
func (menuSrv systemAuthMenuService) Detail(id string) (res systemSchema.SystemAuthMenuResp, e error) {
	var menu system_model.SystemAuthMenu
	err := menuSrv.db.Where("id = ?", id).First(&menu).Error
	if e = response.CheckDBNotRecord(err, "菜单已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, menu)
	return
}

func (menuSrv systemAuthMenuService) Add(addReq systemSchema.SystemAuthMenuAddReq) (e error) {
	var menu system_model.SystemAuthMenu
	convert_util.Copy(&menu, addReq)
	err := menuSrv.db.Create(&menu).Error
	if e = response.CheckErr(err, "添加失败"); e != nil {
		return
	}
	// TODO 清除角色缓存
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
	return
}

func (menuSrv systemAuthMenuService) Edit(editReq systemSchema.SystemAuthMenuEditReq) (e error) {
	// 检查菜单是否存在
	var menu system_model.SystemAuthMenu
	err := menuSrv.db.Where("id = ?", editReq.ID).First(&menu).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("菜单已不存在")
		}
		return response.CheckErr(err, "查询菜单失败")
	}

	convert_util.Copy(&menu, editReq)
	result := menuSrv.db.Model(&menu).Select("*").Updates(menu)
	if result.Error != nil {
		return response.CheckErr(result.Error, "编辑失败")
	}

	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
	return
}

// Del 删除菜单
func (menuSrv systemAuthMenuService) Del(id string) (e error) {
	// 检查是否有子菜单
	if r := menuSrv.db.Where("pid = ?", id).Limit(1).Find(&system_model.SystemAuthMenu{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("请先删除子菜单再操作！")
	}

	result := menuSrv.db.Where("id = ?", id).Delete(&system_model.SystemAuthMenu{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("菜单已不存在")
	}

	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
	return
}
