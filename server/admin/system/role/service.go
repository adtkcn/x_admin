package role

import (
	"strconv"
	"strings"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/system_model"
	"x_admin/util"
	"x_admin/util/convert_util"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

type ISystemAuthRoleService interface {
	All() (res []SystemAuthRoleSimpleResp, e error)
	List(page request.PageReq) (res response.PageResp, e error)
	Detail(id uint) (res SystemAuthRoleResp, e error)
	Add(addReq SystemAuthRoleAddReq) (e error)
	Edit(editReq SystemAuthRoleEditReq) (e error)
	Del(id uint) (e error)
}

var Service = NewSystemAuthRoleService()

// NewSystemAuthRoleService 初始化
func NewSystemAuthRoleService() ISystemAuthRoleService {
	db := core.GetDB()
	return &systemAuthRoleService{db: db}
}

// systemAuthRoleService 系统角色服务实现类
type systemAuthRoleService struct {
	db *gorm.DB
}

// All 角色所有
func (roleSrv systemAuthRoleService) All() (res []SystemAuthRoleSimpleResp, e error) {
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
	err = roleModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	var roleResp []SystemAuthRoleResp
	convert_util.Copy(&roleResp, roles)
	for i := 0; i < len(roleResp); i++ {
		roleResp[i].Menus = []uint{}
		roleResp[i].Member = roleSrv.getMemberCnt(roleResp[i].ID)
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    roleResp,
	}, nil
}

// Detail 角色详情
func (roleSrv systemAuthRoleService) Detail(id uint) (res SystemAuthRoleResp, e error) {
	var role system_model.SystemAuthRole
	err := roleSrv.db.Where("id = ?", id).Limit(1).First(&role).Error
	if e = response.CheckErrDBNotRecord(err, "角色已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, role)
	res.Member = roleSrv.getMemberCnt(role.ID)
	res.Menus, _ = PermService.SelectMenuIdsByRoleId(role.ID)
	return
}

// getMemberCnt 根据角色ID获取成员数量
func (roleSrv systemAuthRoleService) getMemberCnt(roleId uint) (count int64) {
	roleSrv.db.Model(&system_model.SystemAuthAdmin{}).Where(
		"role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return
}

// Add 新增角色
func (roleSrv systemAuthRoleService) Add(addReq SystemAuthRoleAddReq) (e error) {
	var role system_model.SystemAuthRole
	if r := roleSrv.db.Where("name = ?", strings.Trim(addReq.Name, " ")).Limit(1).First(&role); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("角色名称已存在!")
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
func (roleSrv systemAuthRoleService) Edit(editReq SystemAuthRoleEditReq) (e error) {
	err := roleSrv.db.Where("id = ?", editReq.ID).Limit(1).First(&system_model.SystemAuthRole{}).Error
	if e = response.CheckErrDBNotRecord(err, "角色已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	var role system_model.SystemAuthRole
	if r := roleSrv.db.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).Limit(1).First(&role); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("角色名称已存在!")
	}
	role.ID = editReq.ID
	roleMap := structs.Map(editReq)
	delete(roleMap, "ID")
	delete(roleMap, "MenuIds")
	roleMap["Name"] = strings.Trim(editReq.Name, " ")
	// 事务
	err = roleSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&role).Updates(roleMap).Error
		var te error
		if te = response.CheckErr(txErr, "编辑角色失败"); te != nil {
			return te
		}
		// 删除角色菜单
		if te = PermService.BatchDeleteByRoleId(editReq.ID, tx); te != nil {
			return te
		}
		// 重新保存角色菜单
		if te = PermService.BatchSaveByMenuIds(editReq.ID, editReq.MenuIds, tx); te != nil {
			return te
		}
		return nil
	})
	e = PermService.CacheRoleMenusByRoleId(editReq.ID)
	if e != nil {
		return e
	}
	e = response.CheckErr(err, "Edit Transaction err")
	return
}

// Del 删除角色
func (roleSrv systemAuthRoleService) Del(id uint) (e error) {
	err := roleSrv.db.Where("id = ?", id).Limit(1).First(&system_model.SystemAuthRole{}).Error
	if e = response.CheckErrDBNotRecord(err, "角色已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待删除数据查找失败"); e != nil {
		return
	}
	if r := roleSrv.db.Where("role = ? AND is_delete = ?", id, 0).Limit(1).Find(&system_model.SystemAuthAdmin{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("角色已被管理员使用,请先移除!")
	}
	// 事务
	err = roleSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&system_model.SystemAuthRole{}, "id = ?", id).Error
		var te error
		if te = response.CheckErr(txErr, "Del Delete in tx err"); te != nil {
			return te
		}
		if te = PermService.BatchDeleteByRoleId(id, tx); te != nil {
			return te
		}
		util.RedisUtil.HDel(config.AdminConfig.BackstageRolesKey, strconv.FormatUint(uint64(id), 10))
		return nil
	})
	e = response.CheckErr(err, "Del Transaction err")
	return
}
