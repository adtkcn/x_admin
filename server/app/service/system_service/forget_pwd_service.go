package system_service

import (
	"errors"
	"strings"

	"x_admin/app/model/system_model"
	"x_admin/app/schema/system_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"gorm.io/gorm"
)

var ForgetPwdService = NewForgetPwdService()

func NewForgetPwdService() *forgetPwdService {
	return &forgetPwdService{db: core.GetDB()}
}

type forgetPwdService struct {
	db *gorm.DB
}

// SendResetCode 发送密码重置验证码到邮箱
func (srv forgetPwdService) SendResetCode(req *system_schema.SystemForgotPwdSendCodeReq) (e error) {
	// 1. 校验邮箱是否已注册
	var admin system_model.SystemAuthAdmin
	err := srv.db.Where("email = ?", req.Email).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该邮箱未注册")
		}
		return response.CheckErr(err, "查询邮箱失败")
	}

	// 2. 复用统一邮箱验证码逻辑（限流/每日上限/生成验证码/发邮件一体化）
	if err := util.EmailCodeUtil.SendCode(req.Email, util.CodeSceneReset, ""); err != nil {
		return err
	}

	return nil
}

// ResetPassword 校验验证码并重置密码
func (srv forgetPwdService) ResetPassword(req *system_schema.SystemForgotPwdResetReq) (e error) {
	// 1. 校验邮箱验证码（统一逻辑：过期/错误/一次性删除）
	if err := util.EmailCodeUtil.VerifyCode(req.Email, util.CodeSceneReset, req.Code, ""); err != nil {
		return err
	}

	// 2. 查找用户
	var admin system_model.SystemAuthAdmin
	err := srv.db.Where("email = ?", req.Email).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该邮箱未注册")
		}
		return response.CheckErr(err, "查询用户失败")
	}

	// 4. 生成新salt并加密新密码
	salt := util.ToolsUtil.RandomString(5)
	newPassword := util.ToolsUtil.MakeMd5(strings.Trim(req.Password, " ") + salt)

	// 5. 更新数据库
	updates := map[string]interface{}{
		"password": newPassword,
		"salt":     salt,
	}
	err = srv.db.Model(&admin).Updates(updates).Error
	if e = response.CheckErr(err, "重置密码失败"); e != nil {
		return
	}

	// 6. 自增 token_version 使该用户所有 token 失效（强制重新登录）
	err = srv.db.Model(&admin).Update("token_version", gorm.Expr("token_version + 1")).Error
	if e = response.CheckErr(err, "重置密码更新 token_version 失败"); e != nil {
		return
	}
	LoginService.InvalidateTokenVersionCache(admin.ID)

	// 8. 清除用户缓存
	util.RedisUtil.Del(config.AdminConfig.BackstageAdminKey + ":" + admin.ID)
	util.RedisUtil.Del(config.AdminConfig.BackstageAdminPermsKey + ":" + admin.ID)

	// 9. 记录日志
	core.Logger.Infof("用户 %s 通过邮箱验证码重置了密码", req.Email)

	return nil
}
