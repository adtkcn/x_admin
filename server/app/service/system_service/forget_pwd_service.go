package system_service

import (
	"errors"
	"fmt"
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

	// 2. 限流：同一邮箱60秒内最多发送2次
	limitKey := "forgot_pwd:limit:" + req.Email
	count := util.RedisUtil.Incr(limitKey)
	if count == 1 {
		util.RedisUtil.Expire("forgot_pwd:limit:"+req.Email, 60)
	}
	if count > 2 {
		return errors.New("发送频率过高，请稍后再试")
	}

	// 3. 生成6位数字验证码
	code := util.ToolsUtil.Random(100000, 999999)

	// 4. 存入Redis，key: forgot_pwd:code:{email}，有效期300秒
	codeKey := "forgot_pwd:code:" + req.Email
	util.RedisUtil.Set(codeKey, fmt.Sprintf("%d", code), 300)

	// 5. 发送邮件
	opts := util.EmailOptions{
		To:      []string{req.Email},
		Subject: fmt.Sprintf(`【%s】密码重置验证码`, config.AppConfig.AppName),
		HTMLBody: fmt.Sprintf(`
			<div style="font-family:Helvetica,Arial,sans-serif;font-size:14px;color:#333;max-width:600px;margin:0 auto;padding:20px;">
				<p>您正在申请重置密码，验证码如下：</p>
				<p style="font-size:32px;font-weight:bold;color:#409EFF;letter-spacing:4px;margin:20px 0;">%d</p>
				<p>验证码 <b>5分钟内</b> 有效，请勿泄露给他人。</p>
				<p style="color:#999;font-size:12px;margin-top:30px;">如非本人操作，请忽略此邮件。</p>
			</div>`, code),
	}
	if err := util.EmailUtil.SendEmail(opts); err != nil {
		core.Logger.Error("忘记密码邮件发送失败:", err)
		// 清理已写入Redis的验证码
		util.RedisUtil.Del(codeKey)
		return errors.New("验证码发送失败，请稍后重试")
	}

	return nil
}

// ResetPassword 校验验证码并重置密码
func (srv forgetPwdService) ResetPassword(req *system_schema.SystemForgotPwdResetReq) (e error) {
	// 1. 从Redis获取验证码
	codeKey := "forgot_pwd:code:" + req.Email
	storedCode := util.RedisUtil.Get(codeKey)
	if storedCode == "" {
		return errors.New("验证码已过期，请重新获取")
	}

	// 2. 验证码匹配校验
	if storedCode != req.Code {
		return errors.New("验证码错误")
	}

	// 3. 查找用户
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

	// 6. 删除Redis中的验证码（一次性使用）
	util.RedisUtil.Del(codeKey)
	// 清理限流key
	util.RedisUtil.Del("forgot_pwd:limit:" + req.Email)

	// 7. 清除该用户所有登录token（强制重新登录）
	var adminSetKey = config.AdminConfig.BackstageTokenSet + admin.ID
	tokens := util.RedisUtil.SGet(adminSetKey)
	if len(tokens) > 0 {
		var tokenKeys []string
		for _, t := range tokens {
			tokenKeys = append(tokenKeys, config.AdminConfig.BackstageTokenKey+t)
		}
		util.RedisUtil.Del(tokenKeys...)
	}
	util.RedisUtil.Del(adminSetKey)

	// 8. 清除用户缓存
	util.RedisUtil.HDel(config.AdminConfig.BackstageAdminKey, admin.ID)
	util.RedisUtil.HDel(config.AdminConfig.BackstageAdminPermsKey, admin.ID)

	// 9. 记录日志
	core.Logger.Infof("用户 %s 通过邮箱验证码重置了密码", req.Email)

	return nil
}
