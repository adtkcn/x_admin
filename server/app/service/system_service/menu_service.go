package system_service

import (
	"errors"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/system_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
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
func (menuSrv systemAuthMenuService) SelectMenuByAdminId(adminId string) (menuList []*system_schema.SystemAuthMenuResp, e error) {
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
	var menuResps []*system_schema.SystemAuthMenuResp
	convert_util.Copy(&menuResps, menus)

	trees := util.ListToTree(menuResps, "")

	return trees, nil
}

// List 菜单列表
func (menuSrv systemAuthMenuService) List() (res []system_schema.SystemAuthMenuResp, e error) {
	var menus []system_model.SystemAuthMenu
	err := menuSrv.db.Order("menu_sort desc, id").Find(&menus).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	var menuResps []system_schema.SystemAuthMenuResp
	convert_util.Copy(&menuResps, menus)
	return menuResps, nil
}

// Detail 菜单详情
func (menuSrv systemAuthMenuService) Detail(id string) (res system_schema.SystemAuthMenuResp, e error) {
	var menu system_model.SystemAuthMenu
	err := menuSrv.db.Where("id = ?", id).First(&menu).Error
	if e = response.CheckDBErr(err, "菜单已不存在!", "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, menu)
	return
}

func (menuSrv systemAuthMenuService) Add(addReq system_schema.SystemAuthMenuAddReq) (e error) {
	var menu system_model.SystemAuthMenu
	convert_util.Copy(&menu, addReq)
	err := menuSrv.db.Create(&menu).Error
	if e = response.CheckErr(err, "添加失败"); e != nil {
		return
	}
	return
}

// getMenuDescendantIds 获取某菜单的所有后代菜单ID（用于防止将上级设为自身或子孙）
func (menuSrv systemAuthMenuService) getMenuDescendantIds(id string) (descendantIds []string, e error) {
	var menus []system_model.SystemAuthMenu
	if e = menuSrv.db.Select("id, pid").Find(&menus).Error; e != nil {
		return
	}
	childrenMap := make(map[string][]string)
	for _, m := range menus {
		childrenMap[m.Pid] = append(childrenMap[m.Pid], m.ID)
	}
	// 广度优先收集所有后代
	queue := []string{id}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, childId := range childrenMap[current] {
			descendantIds = append(descendantIds, childId)
			queue = append(queue, childId)
		}
	}
	return
}

func (menuSrv systemAuthMenuService) Edit(editReq system_schema.SystemAuthMenuEditReq) (e error) {
	// 检查菜单是否存在
	var menu system_model.SystemAuthMenu
	err := menuSrv.db.Where("id = ?", editReq.ID).First(&menu).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("菜单已不存在")
		}
		return response.CheckErr(err, "查询菜单失败")
	}

	// 禁止将上级菜单设为自身或自身的子孙，避免形成环路
	if editReq.Pid != "" {
		if editReq.Pid == editReq.ID {
			return response.AssertArgumentError.SetMessage("上级菜单不能是自己!")
		}
		descendantIds, dErr := menuSrv.getMenuDescendantIds(editReq.ID)
		if dErr != nil {
			return dErr
		}
		for _, dId := range descendantIds {
			if dId == editReq.Pid {
				return response.AssertArgumentError.SetMessage("上级菜单不能是自己的子级!")
			}
		}
	}

	convert_util.Copy(&menu, editReq)
	result := menuSrv.db.Model(&menu).Select("*").Updates(menu)
	if result.Error != nil {
		return response.CheckErr(result.Error, "编辑失败")
	}

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

	return
}

// Sort 菜单拖拽排序（仅同级），按传入 id 顺序重排 menu_sort 字段（单条 SQL 批量更新）
func (menuSrv systemAuthMenuService) Sort(ids []string) (e error) {
	if len(ids) == 0 {
		return
	}
	var menus []system_model.SystemAuthMenu
	if e = response.CheckErr(menuSrv.db.Where("id in ?", ids).Find(&menus).Error, "查询菜单失败"); e != nil {
		return
	}
	if len(menus) != len(ids) {
		return response.AssertArgumentError.SetMessage("拖拽数据异常，存在无效菜单!")
	}
	// 列表排序为 menu_sort desc，ids[0] 排最前赋予最大 menu_sort 值；用 CASE 单条 SQL 批量更新
	caseExpr := "CASE id"
	args := make([]interface{}, 0, len(ids)*2)
	for i, id := range ids {
		caseExpr += " WHEN ? THEN ?"
		args = append(args, id, len(ids)-i)
	}
	caseExpr += " ELSE menu_sort END"
	if err := menuSrv.db.Model(&system_model.SystemAuthMenu{}).
		Where("id in ?", ids).
		Update("menu_sort", gorm.Expr(caseExpr, args...)).Error; err != nil {
		return response.CheckErr(err, "排序更新失败")
	}
	return
}
