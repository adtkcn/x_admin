package system_service

import (
	"errors"
	"fmt"

	"strings"
	"time"

	"x_admin/app/model/system_model"
	"x_admin/app/schema/system_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"github.com/adtkcn/x_null"
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

// FindByEmail 根据邮箱(账号)查找管理员
func (adminSrv systemAuthAdminService) FindByEmail(email string) (admin system_model.SystemAuthAdmin, err error) {
	err = adminSrv.db.Where("email = ?", email).First(&admin).Error
	return
}

// Self 当前用户
func (adminSrv systemAuthAdminService) Self(adminId string) (res system_schema.SystemAuthAdminSelfResp, e error) {
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
	var admin system_schema.SystemAuthAdminSelfOneResp
	convert_util.Copy(&admin, sysAdmin)
	admin.Dept = sysAdmin.DeptId
	// Avatar 字段已是完整访问地址（/api/uploads/<id> 或 /api/static/...），由前端拼 ossDomain
	return system_schema.SystemAuthAdminSelfResp{User: admin, Permissions: auths}, nil
}

// 获取部门下用户列表
func (adminSrv systemAuthAdminService) ListByDeptId(dept_id string) (res []system_schema.SystemAuthAdminResp, e error) {
	if dept_id == "" {
		e = errors.New("部门ID不能为空")
		return
	}
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	deptTbName := core.DBTableName(&system_model.SystemAuthDept{})
	// 数据
	var adminResp []system_schema.SystemAuthAdminResp
	err := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).
		Table(adminTbName+" AS admin").
		Joins(fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).
		Select(fmt.Sprintf("admin.*, %s.name as dept", deptTbName)).
		Where("admin.dept_id = ?", dept_id).
		Find(&adminResp).Error

	if e = response.CheckDBNotRecord(err, "获取部门下用户列表失败"); e != nil {
		return
	}
	if err != nil {
		e = err
		return
	}

	return adminResp, nil
}

// 导出
func (adminSrv systemAuthAdminService) ExportFile(listReq system_schema.SystemAuthAdminListReq) (res []system_schema.SystemAuthAdminResp, e error) {
	// 查询
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	deptTbName := core.DBTableName(&system_model.SystemAuthDept{})
	postTbName := core.DBTableName(&system_model.SystemAuthPost{})
	adminRoleTbName := core.DBTableName(&system_model.SystemAuthAdminRole{})

	// 基本查询
	adminModel := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin").Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName),
	).Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.post_id = %s.id", postTbName, postTbName),
	).Select(
		fmt.Sprintf("admin.*, %s.name as dept, %s.name as post", deptTbName, postTbName),
	)

	// 条件
	if listReq.Email != "" {
		adminModel = adminModel.Where("email like ?", "%"+listReq.Email+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.RoleId != "" {
		adminModel = adminModel.Where("admin.id in (SELECT admin_id FROM "+adminRoleTbName+" WHERE role_id = ?)", listReq.RoleId)
	}

	// 数据
	var adminResp []system_schema.SystemAuthAdminResp
	err := adminModel.Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}

	// 收集所有管理员ID
	var adminIds []string
	for _, admin := range adminResp {
		if admin.ID != config.AdminConfig.SuperAdminId {
			adminIds = append(adminIds, admin.ID)
		}
	}

	// 一次性获取所有用户的角色关联
	adminRoleMap := make(map[string][]string) // adminId -> roleIds
	if len(adminIds) > 0 {
		var adminRoles []system_model.SystemAuthAdminRole
		adminSrv.db.Where("admin_id in ?", adminIds).Find(&adminRoles)
		for _, ar := range adminRoles {
			adminRoleMap[ar.AdminId] = append(adminRoleMap[ar.AdminId], ar.RoleId)
		}
	}

	// 收集所有角色ID
	var allRoleIds []string
	roleIdMap := make(map[string]bool)
	for _, userRoleIds := range adminRoleMap {
		for _, roleId := range userRoleIds {
			if !roleIdMap[roleId] {
				roleIdMap[roleId] = true
				allRoleIds = append(allRoleIds, roleId)
			}
		}
	}

	// 一次性获取所有角色信息
	roleMap := make(map[string]string) // roleId -> roleName
	if len(allRoleIds) > 0 {
		var roles []system_model.SystemAuthRole
		adminSrv.db.Where("id in ?", allRoleIds).Find(&roles)
		for _, role := range roles {
			roleMap[role.ID] = role.Name
		}
	}

	// 组装数据
	for i := 0; i < len(adminResp); i++ {
		if adminResp[i].ID == config.AdminConfig.SuperAdminId {
			adminResp[i].Role = "超管"
			adminResp[i].RoleIds = []string{}
		} else {
			roleIds := adminRoleMap[adminResp[i].ID]
			adminResp[i].RoleIds = roleIds
			if len(roleIds) > 0 {
				var roleNames []string
				for _, roleId := range roleIds {
					if roleName, ok := roleMap[roleId]; ok {
						roleNames = append(roleNames, roleName)
					}
				}
				adminResp[i].Role = strings.Join(roleNames, ",")
			}
		}
	}

	return adminResp, nil
}

// 导入
func (adminSrv systemAuthAdminService) ImportFile(importReq []system_schema.SystemAuthAdminResp) (e error) {
	if len(importReq) == 0 {
		return nil
	}

	// 批量查询已存在的邮箱
	var emails []string
	for _, item := range importReq {
		emails = append(emails, item.Email)
	}
	var existingAdmins []system_model.SystemAuthAdmin
	if err := adminSrv.db.Where("email IN ?", emails).Find(&existingAdmins).Error; err != nil {
		return response.CheckErr(err, "检查用户是否存在失败")
	}
	// 构建已存在的邮箱 map
	existingMap := make(map[string]bool)
	for _, admin := range existingAdmins {
		existingMap[admin.Email] = true
	}

	for _, importItem := range importReq {
		// 检查用户是否已存在（内存判断）
		if existingMap[importItem.Email] {
			continue
		}

		// 创建用户
		var sysAdmin system_model.SystemAuthAdmin
		convert_util.Copy(&sysAdmin, importItem)

		// 初始化密码
		sysAdmin.Salt = "WFdiD"
		sysAdmin.Password = "81a13dd8e25644a8823082573ca973f7"

		// 设置默认头像
		if sysAdmin.Avatar == "" {
			sysAdmin.Avatar = "/api/static/backend_avatar.png"
		}

		// 单个用户开启事务
		err := adminSrv.db.Transaction(func(tx *gorm.DB) error {
			// 插入用户
			if err := tx.Create(&sysAdmin).Error; err != nil {
				return err
			}

			// 处理角色关联（批量写入）
			if len(importItem.RoleIds) > 0 {
				var adminRoles []system_model.SystemAuthAdminRole
				for _, roleId := range importItem.RoleIds {
					adminRoles = append(adminRoles, system_model.SystemAuthAdminRole{
						AdminId: sysAdmin.ID,
						RoleId:  roleId,
					})
				}
				if err := tx.Create(&adminRoles).Error; err != nil {
					return err
				}
			}
			return nil
		})

		if err != nil {
			e = response.CheckErr(err, "添加用户失败: "+importItem.Email)
			return
		}
	}

	return nil
}

// 获取Excel的列
func (adminSrv systemAuthAdminService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "账号", Key: "Email", Width: 15, Decode: x_null.DecodeString},
		{Name: "名称", Key: "Nickname", Width: 15, Decode: x_null.DecodeString},
		{Name: "头像", Key: "Avatar", Width: 15, Decode: x_null.DecodeString},

		{Name: "角色ID", Key: "RoleIds", Width: 15, Decode: func(value any) (any, error) {
			return strings.Split(value.(string), ","), nil
		}},
		{Name: "角色", Key: "Role", Width: 15, Decode: x_null.DecodeString},
		{Name: "部门ID", Key: "DeptId", Width: 15, Decode: x_null.DecodeString},
		{Name: "部门", Key: "Dept", Width: 15, Decode: x_null.DecodeString},

		{Name: "岗位ID", Key: "PostId", Width: 15, Decode: x_null.DecodeString},
		{Name: "岗位", Key: "Post", Width: 15, Decode: x_null.DecodeString},

		{Name: "是否禁用", Key: "IsDisable", Width: 15, Decode: x_null.DecodeInt64},
		{Name: "最后登录IP", Key: "LastLoginIp", Width: 15, Decode: x_null.DecodeString},
		{Name: "最后登录时间", Key: "LastLoginTime", Width: 15, Decode: x_null.DecodeTime},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: x_null.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: x_null.DecodeTime},
	}
	return cols
}

// List 管理员列表
func (adminSrv systemAuthAdminService) List(page request.PageReq, listReq system_schema.SystemAuthAdminListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	deptTbName := core.DBTableName(&system_model.SystemAuthDept{})
	adminRoleTbName := core.DBTableName(&system_model.SystemAuthAdminRole{})

	adminModel := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin").Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).Select(
		fmt.Sprintf("admin.*, %s.name as dept", deptTbName))
	if listReq.Email != "" {
		adminModel = adminModel.Where("email like ?", "%"+listReq.Email+"%")
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
	var adminResp []system_schema.SystemAuthAdminResp
	err = adminModel.Limit(limit).Offset(offset).Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}

	// 收集所有管理员ID
	var adminIds []string
	for _, admin := range adminResp {
		if admin.ID != config.AdminConfig.SuperAdminId {
			adminIds = append(adminIds, admin.ID)
		}
	}

	// 一次性获取所有用户的角色关联
	adminRoleMap := make(map[string][]string) // adminId -> roleIds
	if len(adminIds) > 0 {
		var adminRoles []system_model.SystemAuthAdminRole
		adminSrv.db.Where("admin_id in ?", adminIds).Find(&adminRoles)
		for _, ar := range adminRoles {
			adminRoleMap[ar.AdminId] = append(adminRoleMap[ar.AdminId], ar.RoleId)
		}
	}

	// 收集所有角色ID
	var allRoleIds []string
	roleIdMap := make(map[string]bool)
	for _, userRoleIds := range adminRoleMap {
		for _, roleId := range userRoleIds {
			if !roleIdMap[roleId] {
				roleIdMap[roleId] = true
				allRoleIds = append(allRoleIds, roleId)
			}
		}
	}

	// 一次性获取所有角色信息
	roleMap := make(map[string]string) // roleId -> roleName
	if len(allRoleIds) > 0 {
		var roles []system_model.SystemAuthRole
		adminSrv.db.Where("id in ?", allRoleIds).Find(&roles)
		for _, role := range roles {
			roleMap[role.ID] = role.Name
		}
	}

	// 组装数据
	for i := 0; i < len(adminResp); i++ {
		if adminResp[i].ID == config.AdminConfig.SuperAdminId {
			adminResp[i].Role = "超管"
			adminResp[i].RoleIds = []string{}
		} else {
			roleIds := adminRoleMap[adminResp[i].ID]
			adminResp[i].RoleIds = roleIds
			if len(roleIds) > 0 {
				var roleNames []string
				for _, roleId := range roleIds {
					if roleName, ok := roleMap[roleId]; ok {
						roleNames = append(roleNames, roleName)
					}
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
func (adminSrv systemAuthAdminService) ListAll(listReq system_schema.SystemAuthAdminListReq) (res []system_schema.SystemAuthAdminResp, e error) {
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	deptTbName := core.DBTableName(&system_model.SystemAuthDept{})
	adminRoleTbName := core.DBTableName(&system_model.SystemAuthAdminRole{})
	adminModel := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin").Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).Select(
		fmt.Sprintf("admin.*, %s.name as dept", deptTbName))
	if listReq.Email != "" {
		adminModel = adminModel.Where("email like ?", "%"+listReq.Email+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.RoleId != "" {
		adminModel = adminModel.Where("admin.id in (SELECT admin_id FROM "+adminRoleTbName+" WHERE role_id = ?)", listReq.RoleId)
	}
	var adminResp []system_schema.SystemAuthAdminResp
	err := adminModel.Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	return adminResp, nil
}

// Detail 管理员详细
func (adminSrv systemAuthAdminService) Detail(id string) (res system_schema.SystemAuthAdminResp, e error) {
	var sysAdmin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("id = ?", id).First(&sysAdmin).Error
	if e = response.CheckDBNotRecord(err, "账号已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, sysAdmin)
	if res.Dept == "" {
		// res.Dept = res.DeptId
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
func (adminSrv systemAuthAdminService) Add(addReq system_schema.SystemAuthAdminAddReq) (e error) {
	// 检查 email 是否已存在
	var existAdmin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("email = ?", addReq.Email).First(&existAdmin).Error
	if err == nil {
		return errors.New("账号已存在换一个吧！")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.CheckErr(err, "Add Find err")
	}

	// 批量验证角色是否存在及是否被禁用
	if len(addReq.RoleIds) > 0 {
		var roles []system_model.SystemAuthRole
		adminSrv.db.Where("id IN ?", addReq.RoleIds).Find(&roles)
		if len(roles) != len(addReq.RoleIds) {
			return errors.New("包含无效的角色!")
		}
		for _, role := range roles {
			if role.IsDisable > 0 {
				return errors.New("角色[" + role.Name + "]已被禁用!")
			}
		}
	}

	var sysAdmin system_model.SystemAuthAdmin
	passwdLen := len(addReq.Password)
	if passwdLen != 32 {
		return errors.New("密码格式不正确")
	}
	salt := util.ToolsUtil.RandomString(5)
	convert_util.Copy(&sysAdmin, addReq)
	sysAdmin.Salt = salt
	sysAdmin.Password = util.ToolsUtil.MakeMd5(strings.Trim(addReq.Password, " ") + salt)
	if addReq.Avatar == "" {
		addReq.Avatar = "/api/static/backend_avatar.png"
	}
	sysAdmin.Avatar = addReq.Avatar // 直接保存完整访问地址（/api/uploads/<id> 或内置默认头像路径）
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
func (adminSrv systemAuthAdminService) Edit(c *gin.Context, editReq system_schema.SystemAuthAdminEditReq) (e error) {
	err := adminSrv.db.Where("id = ?", editReq.ID).First(&system_model.SystemAuthAdmin{}).Error
	if e = response.CheckDBNotRecord(err, "账号不存在了!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}

	// 检查 email 是否已被其他用户使用
	var existAdmin system_model.SystemAuthAdmin
	err = adminSrv.db.Where("email = ? AND id != ?", editReq.Email, editReq.ID).First(&existAdmin).Error
	if err == nil {
		return errors.New("账号已存在换一个吧！")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.CheckErr(err, "Edit Find err")
	}

	// 批量验证角色是否存在
	if editReq.ID != config.AdminConfig.SuperAdminId && len(editReq.RoleIds) > 0 {
		var count int64
		adminSrv.db.Model(&system_model.SystemAuthRole{}).Where("id IN ?", editReq.RoleIds).Count(&count)
		if int(count) != len(editReq.RoleIds) {
			return errors.New("包含无效的角色!")
		}
	}

	adminMap := map[string]interface{}{
		"DeptId":    editReq.DeptId,
		"PostId":    editReq.PostId,
		"Email":     editReq.Email,
		"Nickname":  editReq.Nickname,
		"Avatar":    editReq.Avatar, // 直接保存完整访问地址（/api/uploads/<id> 或内置默认头像路径）
		"Sort":      editReq.Sort,
		"IsDisable": editReq.IsDisable,
	}

	if editReq.Password != "" {
		passwdLen := len(editReq.Password)
		if passwdLen != 32 {
			return response.Failed.SetMessage("密码格式不正确")
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(editReq.Password, "") + salt)
	}
	err = adminSrv.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system_model.SystemAuthAdmin{}).Where("id = ?", editReq.ID).Updates(adminMap).Error; err != nil {
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
	adminSrv.CacheAdminById(editReq.ID)
	PermService.RemoveAdminPermsCache(editReq.ID)
	adminId := config.AdminConfig.GetAdminId(c)
	if editReq.Password != "" {
		if editReq.ID == adminId {
			// 改的是当前登录用户自己的密码：给当前设备重新签发 token 续签，其他设备失效
			if e = LoginService.IssueNewTokenForAdmin(c, adminId); e != nil {
				return
			}
		} else {
			// 管理员重置他人密码：使对方所有设备 token 失效（全设备踢下线）
			if e = adminSrv.ClearOtherTokens(editReq.ID); e != nil {
				return
			}
		}
	}
	return
}

// SendBindEmailCode 发送邮箱绑定验证码
func (adminSrv systemAuthAdminService) SendBindEmailCode(adminId string, email string) (e error) {
	// 检查邮箱占用情况：按邮箱查唯一记录，再判断归属
	var existAdmin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("email = ?", email).First(&existAdmin).Error
	if err == nil {
		if existAdmin.ID == adminId {
			return errors.New("该邮箱已绑定，无需重复绑定")
		}
		return errors.New("该邮箱已被其他账号使用")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.CheckErr(err, "检查邮箱失败")
	}
	// 2. 复用统一邮箱验证码逻辑（限流/每日上限/生成验证码/发邮件一体化）
	// 传入 adminId，使验证码与具体账号绑定，防止跨账号滥用
	if err := util.EmailCodeUtil.SendCode(email, util.CodeSceneBind, adminId); err != nil {
		return err
	}

	return nil
}

// Update 管理员更新自己
func (adminSrv systemAuthAdminService) Update(c *gin.Context, updateReq system_schema.SystemAuthAdminUpdateReq, adminId string) (e error) {
	// 检查id
	var admin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("id = ?", adminId).First(&admin).Error
	if e = response.CheckDBNotRecord(err, "账号不存在了!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Update First err"); e != nil {
		return
	}

	// 邮箱变更校验
	if updateReq.Email != "" && updateReq.Email != admin.Email {
		if updateReq.EmailCode == "" {
			return response.Failed.SetMessage("修改邮箱需要输入验证码")
		}
		// 1. 校验邮箱验证码（统一逻辑：过期/错误/一次性删除）
		// 传入 adminId，与 SendBindEmailCode 发送时保持一致，验证码与账号绑定
		if err := util.EmailCodeUtil.VerifyCode(updateReq.Email, util.CodeSceneBind, updateReq.EmailCode, adminId); err != nil {
			return err
		}

		// codeKey := config.AdminConfig.BackstageAdminKey + ":email_code:" + adminId
		// stored := util.RedisUtil.Get(codeKey)
		// if stored == "" {
		// 	return response.Failed.SetMessage("验证码已过期，请重新获取")
		// }
		// expected := updateReq.Email + ":" + updateReq.EmailCode
		// if stored != expected {
		// 	return response.Failed.SetMessage("验证码错误")
		// }
		// 检查邮箱是否被占用
		var existAdmin system_model.SystemAuthAdmin
		err = adminSrv.db.Where("email = ? AND id != ?", updateReq.Email, adminId).First(&existAdmin).Error
		if err == nil {
			return response.Failed.SetMessage("该邮箱已被其他账号使用")
		}
		// util.RedisUtil.Del(codeKey) // 验证通过，删除验证码
	}

	// 更新管理员信息
	avatar := "/api/static/backend_avatar.png"
	if updateReq.Avatar != "" {
		avatar = updateReq.Avatar
	}
	adminMap := map[string]interface{}{
		"Nickname": updateReq.Nickname,
		"Avatar":   avatar, // 直接保存完整访问地址（/api/uploads/<id> 或内置默认头像路径）
	}
	// 不传邮箱则不更新，避免覆盖为空
	if updateReq.Email != "" {
		adminMap["Email"] = updateReq.Email
	}

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
	}
	err = adminSrv.db.Model(&admin).Updates(adminMap).Error
	if e = response.CheckErr(err, "Update Updates err"); e != nil {
		return
	}
	adminSrv.CacheAdminById(adminId)
	// 如果更改自己的密码,给当前设备重新签发 token 续签（其他设备失效）
	if updateReq.Password != "" {
		if e = LoginService.IssueNewTokenForAdmin(c, adminId); e != nil {
			return
		}
	}
	return
}

// Del 管理员删除
func (adminSrv systemAuthAdminService) Del(c *gin.Context, id string) (e error) {
	if id == config.AdminConfig.SuperAdminId {
		return errors.New("系统管理员不允许删除!")
	}
	if id == config.AdminConfig.GetAdminId(c) {
		return errors.New("不能删除自己!")
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

	util.RedisUtil.Del(config.AdminConfig.BackstageAdminKey + ":" + id)
	util.RedisUtil.Del(config.AdminConfig.BackstageAdminPermsKey + ":" + id)

	// 自增 token_version 使其所有 token 失效（无需维护 token 集合）
	if err := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).
		Where("id = ?", id).
		Update("token_version", gorm.Expr("token_version + 1")).Error; err != nil {
		core.Logger.Errorf("Del 更新 token_version 失败: id=%s err=%v", id, err)
	}
	LoginService.InvalidateTokenVersionCache(id)

	return
}

// Disable 管理员状态切换
func (adminSrv systemAuthAdminService) Disable(c *gin.Context, id string) (e error) {
	var admin system_model.SystemAuthAdmin
	err := adminSrv.db.Where("id = ?", id).Limit(1).Find(&admin).Error
	if e = response.CheckErr(err, "禁用失败"); e != nil {
		return
	}
	if admin.ID == "" {
		return errors.New("账号已不存在!")
	}
	if id == config.AdminConfig.GetAdminId(c) {
		return errors.New("不能禁用自己!")
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

// AdminCache 管理员信息缓存结构。
// TokenVersion 单独声明（json:"token_version"）覆盖 SystemAuthAdmin 中 json:"-" 的 TokenVersion，
// 使 auth 中间件只需读取 admin:users 缓存即可同时拿到管理员信息与 token_version。
type AdminCache struct {
	system_model.SystemAuthAdmin
	TokenVersion int64 `json:"token_version"`
}

// CacheAdminById 缓存管理员（含 token_version）
func (adminSrv systemAuthAdminService) CacheAdminById(id string) (user system_model.SystemAuthAdmin, err error) {
	var admin system_model.SystemAuthAdmin
	result := adminSrv.db.Where("id = ?", id).First(&admin)
	if result.Error != nil {
		core.Logger.Error("CacheAdminById First err", result.Error)
		err = errors.New("查询失败")
		return
	}
	if result.RowsAffected == 0 {
		err = errors.New("管理员不存在")
		return
	}
	// 密码不进缓存
	admin.Password = ""
	cache := AdminCache{SystemAuthAdmin: admin, TokenVersion: admin.TokenVersion}
	str, err := util.ToolsUtil.ObjToJson(&cache)
	if err != nil {
		return
	}
	util.RedisUtil.Set(config.AdminConfig.BackstageAdminKey+":"+admin.ID, str, 0)
	return admin, nil
}

// ClearOtherTokens 改密/重置密码后使该管理员所有 token 失效（JWT 方案下为全设备踢下线）
func (adminSrv systemAuthAdminService) ClearOtherTokens(id string) (err error) {
	if err = adminSrv.db.Model(&system_model.SystemAuthAdmin{}).
		Where("id = ?", id).
		Update("token_version", gorm.Expr("token_version + 1")).Error; err != nil {
		return response.CheckErr(err, "ClearOtherTokens Update token_version err")
	}
	LoginService.InvalidateTokenVersionCache(id)
	return nil
}

// 获取今日新增用户数量和总用户数量
func (adminSrv systemAuthAdminService) GetTodayCount() (res system_schema.SystemAuthAdminTodayCountResp, e error) {
	var totalCount int64
	var todayCount int64
	err := adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Count(&totalCount).Error
	if e = response.CheckErr(err, "GetTodayCount Count err"); e != nil {
		return
	}
	err = adminSrv.db.Model(&system_model.SystemAuthAdmin{}).Where("create_time >= ?", util.NullTimeUtil.TodayZero()).Count(&todayCount).Error
	if e = response.CheckErr(err, "GetTodayCount Count err"); e != nil {
		return
	}
	res.TotalUsers = totalCount
	res.TodayUsers = todayCount
	return
}
