package util

import (
	"fmt"
	"strconv"
	"time"
	"x_admin/core"
	"x_admin/util/convert_util"
)

// 短信验证码场景
const (
	SmsSceneBind  = "sms_bind"  // 绑定手机
	SmsSceneLogin = "sms_login" // 短信验证码登录
	SmsSceneReset = "sms_reset" // 手机号重置密码
)

// Redis key 模板
const (
	smsCodeKeyTpl   = "user:sms:%s:%s"       // user:sms:{scene}:{phone}  验证码值
	smsCodeLimitTpl = "user:sms:limit:%s"    // user:sms:limit:{phone}    60s频率限制
	smsCodeDailyTpl = "user:sms:daily:%s:%s" // user:sms:daily:{phone}:{date} 每日上限
)

var SmsCodeUtil = &smsCodeUtil{}

type smsCodeUtil struct{}

// SendCode 发送短信验证码
func (s *smsCodeUtil) SendCode(phone, scene string) error {
	// 频率限制: 60s内不可重复发送
	limitKey := fmt.Sprintf(smsCodeLimitTpl, phone)
	if RedisUtil.Exists(limitKey) > 0 {
		return fmt.Errorf("验证码发送过于频繁，请60秒后再试")
	}

	// 每日上限: 同一手机号每天最多10次
	today := time.Now().Format("2006-01-02")
	dailyKey := fmt.Sprintf(smsCodeDailyTpl, phone, today)
	dailyCount := RedisUtil.Get(dailyKey)
	if dailyCount != "" {
		count, _ := strconv.Atoi(dailyCount)
		if count >= 10 {
			return fmt.Errorf("今日验证码发送次数已达上限")
		}
	}

	// 生成6位数字验证码
	code := ToolsUtil.Random(100000, 999999)
	codeStr := convert_util.ToString(code)

	// 存储验证码到 Redis，5分钟有效
	codeKey := fmt.Sprintf(smsCodeKeyTpl, scene, phone)
	RedisUtil.Set(codeKey, codeStr, 300)

	// 设置60s频率限制
	RedisUtil.Set(limitKey, "1", 60)

	// 增加每日计数
	if dailyCount == "" {
		RedisUtil.Set(dailyKey, "1", 86400) // 24h
	} else {
		RedisUtil.Incr(dailyKey)
	}

	// TODO: 对接短信服务商（阿里云/腾讯云等）发送短信
	// 当前仅记录日志，生产环境需替换为实际短信发送逻辑
	core.Logger.Infof("SendSmsCode: phone=%s scene=%s code=%s", phone, scene, codeStr)
	return nil
}

// VerifyCode 校验短信验证码（成功后删除）
func (s *smsCodeUtil) VerifyCode(phone, scene, code string) error {
	codeKey := fmt.Sprintf(smsCodeKeyTpl, scene, phone)
	stored := RedisUtil.Get(codeKey)
	if stored == "" {
		return fmt.Errorf("验证码已过期，请重新获取")
	}
	if stored != code {
		return fmt.Errorf("验证码错误")
	}
	// 验证成功，删除验证码（防止重复使用）
	RedisUtil.Del(codeKey)
	return nil
}
