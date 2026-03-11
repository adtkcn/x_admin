package systemService

import (
	"fmt"

	"strings"
	"time"

	"x_admin/app/schema/systemSchema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/system_model"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var AdminService = NewSystemAuthAdminService()

// NewSystemAuthAdminService 初始化
func NewSystemAuthAdminService() *systemAuthAdminService {
	db := core.GetDB()
	return &systemAuthAdminService{db: db}
}

// systemAuthAdminService 系统管理员服务实现类
type systemAuthAdminService struct {
	db *gorm.DB
}

// FindByUsername 根据账号查找管理员
func (adminSrv systemAuthAdminService) FindByUsername(username string) (admin system_model.SystemAuthAdmin, err error) {
	err = adminSrv.db.Where("username = ?", username).First(&admin).Error
	return
}

// Self 当前管理员
func (adminSrv systemAuthAdminService) Self(adminId string) (res systemSchema.SystemAuthAdminSelfResp, e error) {
	// 管理员信息
	var sysAdmin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("id = ? ", adminId).First(&sysAdmin).Error
	if e = response.CheckErr(err, "获取用户信息失败"); e != nil {
		return
	}
	// 角色权限
	var auths []string
	if adminId == config.AdminConfig.SuperAdminId {
		auths = append(auths, "*")
	} else {
		auths, e = PermService.GetAdminPerms(adminId)
		if e != nil {
			return
		}
	}
	var admin systemSchema.SystemAuthAdminSelfOneResp
	convert_util.Copy(&admin, sysAdmin)
	admin.Dept = sysAdmin.DeptId
	admin.Avatar = util.UrlUtil.ToAbsoluteUrl(sysAdmin.Avatar)
	return systemSchema.SystemAuthAdminSelfResp{User: admin, Permissions: auths}, nil
}

// 获取管理员列表-
func (adminSrv systemAuthAdminService) ListByUserIdOrDeptIdPostId(userId, deptId, postId string) (res []systemSchema.SystemAuthAdminResp, e error) {
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})

	adminModel := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin")
	if userId != "" {
		adminModel.Where("admin.id =?", userId)
	}
	if deptId != "" {
		adminModel.Where("admin.dept_id =?", deptId)
	}
	if postId != "" {
		adminModel.Where("admin.post_id =?", postId)
	}
	// 数据
	var adminResp []systemSchema.SystemAuthAdminResp
	err := adminModel.Find(&adminResp).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	for i := 0; i < len(adminResp); i++ {
		adminResp[i].Avatar = util.UrlUtil.ToAbsoluteUrl(adminResp[i].Avatar)
		if adminResp[i].ID == config.AdminConfig.SuperAdminId {
			adminResp[i].Role = "系统管理员"
		}
	}
	return adminResp, nil
}

// 导出
func (adminSrv systemAuthAdminService) ExportFile(listReq systemSchema.SystemAuthAdminListReq) (res []systemSchema.SystemAuthAdminResp, e error) {
	// 查询
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	roleTbName := core.DBTableName(&system_model.SystemAuthRole{})
	deptTbName := core.DBTableName(&system_model.SystemAuthDept{})
	adminModel := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin").Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.role_id = %s.id", roleTbName, roleTbName)).Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).Select(
		fmt.Sprintf("admin.*, %s.name as dept, %s.name as role", deptTbName, roleTbName))
	// 条件
	if listReq.Username != "" {
		adminModel = adminModel.Where("username like ?", "%"+listReq.Username+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.RoleId != "" {
		adminModel = adminModel.Where("role_id = ?", listReq.RoleId)
	}
	// 数据
	var adminResp []systemSchema.SystemAuthAdminResp
	err := adminModel.Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	for i := 0; i < len(adminResp); i++ {
		// adminResp[i].Avatar = util.UrlUtil.ToAbsoluteUrl(adminResp[i].Avatar)
		if adminResp[i].ID == config.AdminConfig.SuperAdminId {
			adminResp[i].Role = "系统管理员"
		}
	}
	return adminResp, nil
}

// 导入
func (adminSrv systemAuthAdminService) ImportFile(importReq []systemSchema.SystemAuthAdminResp) (e error) {
	var sysAdmin []system_model.SystemAuthAdmin
	convert_util.Copy(&sysAdmin, importReq)
	err := adminSrv.db.Create(&sysAdmin).Error
	e = response.CheckErr(err, "添加失败")
	return e
}

// 获取Excel的列
func (adminSrv systemAuthAdminService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "账号", Key: "Username", Width: 15, Decode: core.DecodeString},
		{Name: "昵称", Key: "Nickname", Width: 15, Decode: core.DecodeString},
		{Name: "头像", Key: "Avatar", Width: 15, Decode: core.DecodeString},
		{Name: "角色", Key: "Role", Width: 15, Decode: core.DecodeString},
		{Name: "部门ID", Key: "DeptId", Width: 15, Decode: core.DecodeString},
		{Name: "岗位ID", Key: "PostId", Width: 15, Decode: core.DecodeString},
		{Name: "角色ID", Key: "RoleId", Width: 15, Decode: core.DecodeString},
		{Name: "部门", Key: "Dept", Width: 15, Decode: core.DecodeString},
		{Name: "是否禁用", Key: "IsDisable", Width: 15, Decode: core.DecodeInt},
		{Name: "最后登录IP", Key: "LastLoginIp", Width: 15, Decode: core.DecodeString},
		{Name: "最后登录时间", Key: "LastLoginTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	return cols
}

// List 管理员列表
func (adminSrv systemAuthAdminService) List(page request.PageReq, listReq systemSchema.SystemAuthAdminListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	deptTbName := core.DBTableName(&system_model.SystemAuthDept{})
	adminRoleTbName := core.DBTableName(&system_model.SystemAuthAdminRole{})

	adminModel := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin").Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).Select(
		fmt.Sprintf("admin.*, %s.name as dept", deptTbName))
	if listReq.Username != "" {
		adminModel = adminModel.Where("username like ?", "%"+listReq.Username+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.RoleId != "" {
		adminModel = adminModel.Where("admin.id in (SELECT admin_id FROM "+adminRoleTbName+" WHERE role_id = ?)", listReq.RoleId)
	}
	var count int64
	err := adminModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	var adminResp []systemSchema.SystemAuthAdminResp
	err = adminModel.Limit(limit).Offset(offset).Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	for i := 0; i < len(adminResp); i++ {
		adminResp[i].Avatar = util.UrlUtil.ToAbsoluteUrl(adminResp[i].Avatar)
		if adminResp[i].ID == config.AdminConfig.SuperAdminId {
			adminResp[i].Role = "系统管理员"
		} else {
			roleIds, _ := AdminRoleService.GetRoleIdsByAdminId(adminResp[i].ID)
			adminResp[i].RoleIds = roleIds
			if len(roleIds) > 0 {
				var roles []system_model.SystemAuthRole
				adminSrv.db.Where("id in ?", roleIds).Find(&roles)
				var roleNames []string
				for _, r := range roles {
					roleNames = append(roleNames, r.Name)
				}
				adminResp[i].Role = strings.Join(roleNames, ",")
			}
		}
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    adminResp,
	}, nil
}

// ListAll 管理员列表
func (adminSrv systemAuthAdminService) ListAll(listReq systemSchema.SystemAuthAdminListReq) (res []systemSchema.SystemAuthAdminResp, e error) {
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	deptTbName := core.DBTableName(&system_model.SystemAuthDept{})
	adminRoleTbName := core.DBTableName(&system_model.SystemAuthAdminRole{})
	adminModel := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin").Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).Select(
		fmt.Sprintf("admin.*, %s.name as dept", deptTbName))
	if listReq.Username != "" {
		adminModel = adminModel.Where("username like ?", "%"+listReq.Username+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.RoleId != "" {
		adminModel = adminModel.Where("admin.id in (SELECT admin_id FROM "+adminRoleTbName+" WHERE role_id = ?)", listReq.RoleId)
	}
	var adminResp []systemSchema.SystemAuthAdminResp
	err := adminModel.Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	for i := 0; i < len(adminResp); i++ {
		adminResp[i].Avatar = util.UrlUtil.ToAbsoluteUrl(adminResp[i].Avatar)
		if adminResp[i].ID == config.AdminConfig.SuperAdminId {
			adminResp[i].Role = "系统管理员"
		} else {
			roleIds, _ := AdminRoleService.GetRoleIdsByAdminId(adminResp[i].ID)
			adminResp[i].RoleIds = roleIds
			if len(roleIds) > 0 {
				var roles []system_model.SystemAuthRole
				adminSrv.db.Where("id in ?", roleIds).Find(&roles)
				var roleNames []string
				for _, r := range roles {
					roleNames = append(roleNames, r.Name)
				}
				adminResp[i].Role = strings.Join(roleNames, ",")
			}
		}
	}
	return adminResp, nil
}

// Detail 管理员详细
func (adminSrv systemAuthAdminService) Detail(id string) (res systemSchema.SystemAuthAdminResp, e error) {
	var sysAdmin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("id = ?", id).First(&sysAdmin).Error
	if e = response.CheckDBNotRecord(err, "账号已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, sysAdmin)
	res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
	if res.Dept == "" {
		res.Dept = res.DeptId
	}
	if res.ID != config.AdminConfig.SuperAdminId {
		res.RoleIds, _ = AdminRoleService.GetRoleIdsByAdminId(id)
		if len(res.RoleIds) > 0 {
			var roles []system_model.SystemAuthRole
			adminSrv.db.Where("id in ?", res.RoleIds).Find(&roles)
			var roleNames []string
			for _, r := range roles {
				roleNames = append(roleNames, r.Name)
			}
			res.Role = strings.Join(roleNames, ",")
		}
	}
	return
}

// Add 管理员新增
func (adminSrv systemAuthAdminService) Add(addReq systemSchema.SystemAuthAdminAddReq) (e error) {
	var sysAdmin system_model.SystemAuthAdmin
	r := adminSrv.db.Where("username = ?", addReq.Username).Limit(1).Find(&sysAdmin)
	err := r.Error
	if e = response.CheckErr(err, "Add Find by username err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("账号已存在换一个吧！")
	}
	r = adminSrv.db.Where("nickname = ?", addReq.Nickname).Limit(1).Find(&sysAdmin)
	err = r.Error
	if e = response.CheckErr(err, "Add Find by nickname err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("昵称已存在换一个吧！")
	}
	for _, roleId := range addReq.RoleIds {
		var roleResp systemSchema.SystemAuthRoleResp
		if roleResp, e = RoleService.Detail(roleId); e != nil {
			return
		}
		if roleResp.IsDisable > 0 {
			return response.AssertArgumentError.SetMessage("当前角色已被禁用!")
		}
	}
	passwdLen := len(addReq.Password)
	if passwdLen != 32 {
		return response.Failed.SetMessage("密码格式不正确")
	}
	salt := util.ToolsUtil.RandomString(5)
	convert_util.Copy(&sysAdmin, addReq)
	sysAdmin.Salt = salt
	sysAdmin.Password = util.ToolsUtil.MakeMd5(strings.Trim(addReq.Password, " ") + salt)
	if addReq.Avatar == "" {
		addReq.Avatar = "/api/static/backend_avatar.png"
	}
	sysAdmin.Avatar = util.UrlUtil.ToRelativeUrl(addReq.Avatar)
	err = adminSrv.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&sysAdmin).Error; err != nil {
			return err
		}
		if len(addReq.RoleIds) > 0 {
			if err := AdminRoleService.SaveAdminRoles(sysAdmin.ID, addReq.RoleIds, tx); err != nil {
				return err
			}
		}
		return nil
	})
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 管理员编辑
func (adminSrv systemAuthAdminService) Edit(c *gin.Context, editReq systemSchema.SystemAuthAdminEditReq) (e error) {
	err := adminSrv.db.Where("id = ?", editReq.ID).First(&system_model.SystemAuthAdmin{}).Error
	if e = response.CheckDBNotRecord(err, "账号不存在了!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	var admin system_model.SystemAuthAdmin
	r := adminSrv.db.Where("username = ? AND id != ?", editReq.Username, editReq.ID).Find(&admin)
	err = r.Error
	if e = response.CheckErr(err, "Edit Find by username err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("账号已存在换一个吧！")
	}
	r = adminSrv.db.Where("nickname = ? AND id != ?", editReq.Nickname, editReq.ID).Find(&admin)
	err = r.Error
	if e = response.CheckErr(err, "Edit Find by nickname err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("昵称已存在换一个吧！")
	}
	for _, roleId := range editReq.RoleIds {
		if editReq.ID != config.AdminConfig.SuperAdminId {
			if _, e = RoleService.Detail(roleId); e != nil {
				return
			}
		}
	}
	adminMap := structs.Map(editReq)
	delete(adminMap, "ID")
	delete(adminMap, "RoleIds")
	adminMap["Avatar"] = util.UrlUtil.ToRelativeUrl(editReq.Avatar)
	if editReq.ID == config.AdminConfig.SuperAdminId {
		delete(adminMap, "Username")
	}
	if editReq.Password != "" {
		passwdLen := len(editReq.Password)
		if passwdLen != 32 {
			return response.Failed.SetMessage("密码格式不正确")
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(editReq.Password, "") + salt)
	} else {
		delete(adminMap, "Password")
	}
	err = adminSrv.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&admin).Where("id = ?", editReq.ID).Updates(adminMap).Error; err != nil {
			return err
		}
		if editReq.ID != config.AdminConfig.SuperAdminId {
			if err := AdminRoleService.SaveAdminRoles(editReq.ID, editReq.RoleIds, tx); err != nil {
				return err
			}
		}
		return nil
	})
	if e = response.CheckErr(err, "编辑失败"); e != nil {
		return
	}
	adminSrv.CacheAdminUserByUid(editReq.ID)
	PermService.RemoveAdminPermsCache(editReq.ID)
	AdminRoleService.RemoveAdminRoleCache(editReq.ID)
	adminId := config.AdminConfig.GetAdminId(c)
	if editReq.Password != "" && editReq.ID == adminId {
		token := c.Request.Header.Get("token")
		adminSrv.ClearOtherTokens(adminId, token)
	}
	return
}

// Update 管理员更新自己
func (adminSrv systemAuthAdminService) Update(c *gin.Context, updateReq systemSchema.SystemAuthAdminUpdateReq, adminId string) (e error) {
	// 检查id
	var admin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("id = ?", adminId).First(&admin).Error
	if e = response.CheckDBNotRecord(err, "账号不存在了!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Update First err"); e != nil {
		return
	}
	// 更新管理员信息
	adminMap := structs.Map(updateReq)
	delete(adminMap, "CurrPassword")
	avatar := "/api/static/backend_avatar.png"
	if updateReq.Avatar != "" {
		avatar = updateReq.Avatar
	}
	adminMap["Avatar"] = util.UrlUtil.ToRelativeUrl(avatar)

	if updateReq.Password != "" {
		currPass := util.ToolsUtil.MakeMd5(updateReq.CurrPassword + admin.Salt)
		if currPass != admin.Password {
			return response.Failed.SetMessage("当前密码不正确!")
		}
		passwdLen := len(updateReq.Password)
		if passwdLen != 32 {
			return response.Failed.SetMessage("新密码格式不正确")
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(updateReq.Password, " ") + salt)
	} else {
		delete(adminMap, "Password")
	}
	err = adminSrv.db.Model(&admin).Updates(adminMap).Error
	if e = response.CheckErr(err, "Update Updates err"); e != nil {
		return
	}
	adminSrv.CacheAdminUserByUid(adminId)
	// 如果更改自己的密码,则删除其他登录缓存
	if updateReq.Password != "" {
		token := c.Request.Header.Get("token")
		adminSrv.ClearOtherTokens(adminId, token)
	}
	return
}

// Del 管理员删除
func (adminSrv systemAuthAdminService) Del(c *gin.Context, id string) (e error) {
	if id == config.AdminConfig.SuperAdminId {
		return response.AssertArgumentError.SetMessage("系统管理员不允许删除!")
	}
	if id == config.AdminConfig.GetAdminId(c) {
		return response.AssertArgumentError.SetMessage("不能删除自己!")
	}

	err := adminSrv.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system_model.SystemAuthAdmin{}).Error; err != nil {
			return err
		}
		if err := AdminRoleService.DeleteByAdminId(id); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return response.CheckErr(err, "删除失败")
	}

	util.RedisUtil.HDel(config.AdminConfig.BackstageManageKey, id)
	util.RedisUtil.HDel(config.AdminConfig.BackstageAdminRolesKey, id)
	util.RedisUtil.HDel(config.AdminConfig.BackstageAdminPermsKey, id)
	adminSetKey := config.AdminConfig.BackstageTokenSet + id
	ts := util.RedisUtil.SGet(adminSetKey)
	if len(ts) > 0 {
		var tokenKeys []string
		for _, t := range ts {
			tokenKeys = append(tokenKeys, config.AdminConfig.BackstageTokenKey+t)
		}
		util.RedisUtil.Del(tokenKeys...)
	}
	util.RedisUtil.Del(adminSetKey)

	return
}

// Disable 管理员状态切换
func (adminSrv systemAuthAdminService) Disable(c *gin.Context, id string) (e error) {
	var admin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("id = ?", id).Limit(1).Find(&admin).Error
	if e = response.CheckErr(err, "Disable Find err"); e != nil {
		return
	}
	if admin.ID == "" {
		return response.AssertArgumentError.SetMessage("账号已不存在!")
	}
	if id == config.AdminConfig.GetAdminId(c) {
		return response.AssertArgumentError.SetMessage("不能禁用自己!")
	}
	var isDisable uint8
	if admin.IsDisable == 0 {
		isDisable = 1
	} else {
		isDisable = 0
	}
	err = adminSrv.db.Model(&admin).Updates(map[string]any{"is_disable": isDisable, "update_time": time.Now()}).Error
	e = response.CheckErr(err, "Disable Updates err")
	return
}

// CacheAdminUserByUid 缓存管理员
func (adminSrv systemAuthAdminService) CacheAdminUserByUid(id string) (err error) {
	var admin system_model.SystemAuthAdmin
	err = adminSrv.db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return err
	}
	// redis排除缓存
	admin.Password = ""

	str, err := util.ToolsUtil.ObjToJson(&admin)
	if err != nil {
		return err
	}
	util.RedisUtil.HSet(config.AdminConfig.BackstageManageKey, admin.ID, str, 0)
	return nil
}

// 清理用户其他登陆token
func (adminSrv systemAuthAdminService) ClearOtherTokens(id string, nowToken string) (err error) {
	// 账号token集合key
	adminSetKey := config.AdminConfig.BackstageTokenSet + id
	// 获取账号所有token
	tokens := util.RedisUtil.SGet(adminSetKey)
	if len(tokens) > 0 {
		var delTokens []string
		for _, token := range tokens {
			if token != nowToken {
				delTokens = append(delTokens, config.AdminConfig.BackstageTokenKey+token)
			}
		}
		// 清除其他token缓存
		util.RedisUtil.Del(delTokens...)
	}
	util.RedisUtil.Del(adminSetKey)
	// 添加当前token到集合
	util.RedisUtil.SSet(adminSetKey, nowToken)
	return nil
}
