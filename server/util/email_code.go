package util

import (
	"fmt"
	"strconv"
	"time"
	"x_admin/app/schema/queue_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/util/convert_util"
)

// 邮箱验证码场景
const (
	CodeSceneRegister = "register" // 注册
	CodeSceneReset    = "reset"    // 重置密码
	CodeSceneBind     = "bind"     // 绑定邮箱
	CodeSceneUnbind   = "unbind"   // 解绑手机（邮箱验证码确认身份）
)

// EmailCodeTask 邮箱验证码发送任务载荷（推入队列异步发送）
type EmailCodeTask struct {
	To        string `json:"to"`
	Subject   string `json:"subject"`
	HTMLBody  string `json:"html_body"`
	CreatedAt int64  `json:"created_at"` // 入队时间戳（秒），用于过期判定
}

// Redis key 模板
const (
	// codeKeyTpl   =       // user:code:{scene}:{email}  验证码值
	codeLimitTpl = "user:code:limit:%s"    // user:code:limit:{email}    60s频率限制
	codeDailyTpl = "user:code:daily:%s:%s" // user:code:daily:{email}:{date} 每日上限
)

var EmailCodeUtil = &emailCodeUtil{}

type emailCodeUtil struct{}

// SendCode 发送邮箱验证码
// uid 为可选参数，非空时验证码将与具体账号绑定（key 变为 user:code:{scene}:{uid}:{email}），
// 防止跨账号滥用；为空则保持原有 key 格式，向后兼容。
func (e *emailCodeUtil) SendCode(email, scene, uid string) error {
	// 频率限制: 60s内不可重复发送
	limitKey := fmt.Sprintf(codeLimitTpl, email)
	if RedisUtil.Exists(limitKey) > 0 {
		return fmt.Errorf("验证码发送过于频繁，请60秒后再试")
	}

	// 每日上限: 同一邮箱每日最多10次
	today := time.Now().Format("2006-01-02")
	dailyKey := fmt.Sprintf(codeDailyTpl, email, today)
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
	codeKey := e.buildCodeKey(scene, email, uid)
	RedisUtil.Set(codeKey, codeStr, 300)

	// 设置60s频率限制
	RedisUtil.Set(limitKey, "1", 60)

	// 增加每日计数（如不存在则初始化，否则自增）
	if dailyCount == "" {
		RedisUtil.Set(dailyKey, "1", 86400) // 24h
	} else {
		RedisUtil.Incr(dailyKey)
	}

	// 发送邮件：推入队列异步执行，避免阻塞请求
	sceneMap := map[string]string{
		CodeSceneRegister: "注册账号",
		CodeSceneReset:    "重置密码",
		CodeSceneBind:     "绑定邮箱",
		CodeSceneUnbind:   "解绑手机",
	}
	sceneName := sceneMap[scene]
	if sceneName == "" {
		sceneName = "验证操作"
	}

	opts := EmailCodeTask{
		To:        email,
		CreatedAt: time.Now().Unix(),
		Subject:   fmt.Sprintf("【%s】%s验证码", config.AppConfig.AppName, sceneName),
		HTMLBody: fmt.Sprintf(`
			<h3>%s</h3>
			<p>您的验证码是：<b style="font-size:28px;color:#409eff;letter-spacing:4px">%s</b></p>
			<p>验证码 5 分钟内有效，请勿泄露给他人。</p>
			<p style="color:#999;font-size:12px">如非本人操作，请忽略此邮件。</p>`, sceneName, codeStr),
	}
	if err := core.Queue.Enqueue(queue_schema.QueueEmailCode, opts); err != nil {
		core.Logger.Errorf("SendCode 验证码入队失败: email=%s scene=%s err=%v", email, scene, err)
		return fmt.Errorf("验证码发送失败，请稍后重试")
	}
	return nil
}

// VerifyCode 校验邮箱验证码（成功后删除）
// uid 必须与 SendCode 调用时传入的一致，否则取不到对应验证码。
func (e *emailCodeUtil) VerifyCode(email, scene, code, uid string) error {
	codeKey := e.buildCodeKey(scene, email, uid)
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

// buildCodeKey 构造验证码 Redis key
// uid 为空：user:code:{scene}:{email}（向后兼容）
// uid 非空：user:code:{scene}:{uid}:{email}（验证码与具体账号绑定）
func (e *emailCodeUtil) buildCodeKey(scene, email, uid string) string {
	if uid != "" {
		return fmt.Sprintf("user:code:%s:%s:%s", scene, uid, email)
	}
	return fmt.Sprintf("user:code:%s:%s", scene, email)
}
