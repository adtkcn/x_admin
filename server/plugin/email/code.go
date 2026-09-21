package email

import (
	"fmt"
	"strconv"
	"time"
	"x_admin/app/schema/queue_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/util"
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

// Redis key 模板与验证码相关常量
const (
	codeLimitTpl = "user:code:limit:%s"    // user:code:limit:{email}    60s频率限制
	codeDailyTpl = "user:code:daily:%s:%s" // user:code:daily:{email}:{date} 每日上限
	codeFailTpl  = "user:code:fail:%s"     // user:code:fail:{codeKey}   校验失败计数
)

const (
	codeExpireSec    = 300   // 验证码有效期（秒）
	codeLimitSec     = 60    // 发送频率限制窗口（秒）
	codeDailyTTL     = 86400 // 每日计数保留时长（秒）
	codeDailyLimit   = 10    // 同一邮箱每日最多发送次数
	codeMaxFailTimes = 10    // 同一验证码最多校验失败次数
)

var sceneMap = map[string]string{
	CodeSceneRegister: "注册账号",
	CodeSceneReset:    "重置密码",
	CodeSceneBind:     "绑定邮箱",
	CodeSceneUnbind:   "解绑手机",
}

// SendCode 发送邮箱验证码
// uid 为可选参数，非空时验证码将与具体账号绑定（key 变为 user:code:{scene}:{uid}:{email}），
// 防止跨账号滥用；为空则保持原有 key 格式，向后兼容。
func SendCode(email, scene, uid string) error {
	// 频率限制: SetNX 原子占用 60s 窗口，避免并发请求同时通过检查
	limitKey := fmt.Sprintf(codeLimitTpl, email)
	if !util.RedisUtil.SetNX(limitKey, "1", codeLimitSec) {
		return fmt.Errorf("验证码发送过于频繁，请60秒后再试")
	}

	// 每日上限: Incr 原子自增并返回计数，首次自增时设置过期，
	// 避免 Get+Set 非原子导致并发重置计数；超过上限后拒绝发送
	today := time.Now().Format("2006-01-02")
	dailyKey := fmt.Sprintf(codeDailyTpl, email, today)
	dailyCount := util.RedisUtil.Incr(dailyKey)
	if dailyCount == 1 {
		util.RedisUtil.Expire(dailyKey, codeDailyTTL)
	}
	if dailyCount > codeDailyLimit {
		return fmt.Errorf("今日验证码发送次数已达上限")
	}

	// 生成6位数字验证码
	code := util.ToolsUtil.Random(100000, 999999)
	codeStr := strconv.Itoa(code)

	// 存储验证码到 Redis，5分钟有效；同时清掉旧验证码遗留的失败计数
	codeKey := buildCodeKey(scene, email, uid)
	util.RedisUtil.Set(codeKey, codeStr, codeExpireSec)
	util.RedisUtil.Del(fmt.Sprintf(codeFailTpl, codeKey))

	// 发送邮件：推入队列异步执行，避免阻塞请求

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
func VerifyCode(email, scene, code, uid string) error {
	codeKey := buildCodeKey(scene, email, uid)
	stored := util.RedisUtil.Get(codeKey)
	if stored == "" {
		return fmt.Errorf("验证码已过期，请重新获取")
	}
	if stored != code {
		// 失败计数: 同一验证码累计错误 codeMaxFailTimes 次即作废，防止暴力猜测
		failKey := fmt.Sprintf(codeFailTpl, codeKey)
		fails := util.RedisUtil.Incr(failKey)
		if fails == 1 {
			util.RedisUtil.Expire(failKey, codeExpireSec)
		}
		if fails >= codeMaxFailTimes {
			util.RedisUtil.Del(codeKey, failKey)
			return fmt.Errorf("验证码错误次数过多，已失效，请重新获取")
		}
		return fmt.Errorf("验证码错误")
	}
	// 验证成功，删除验证码与失败计数（防止重复使用）
	util.RedisUtil.Del(codeKey, fmt.Sprintf(codeFailTpl, codeKey))
	return nil
}

// buildCodeKey 构造验证码 Redis key
// uid 为空：user:code:{scene}:{email}（向后兼容）
// uid 非空：user:code:{scene}:{uid}:{email}（验证码与具体账号绑定）
func buildCodeKey(scene, email, uid string) string {
	if uid != "" {
		return fmt.Sprintf("user:code:%s:%s:%s", scene, uid, email)
	}
	return fmt.Sprintf("user:code:%s:%s", scene, email)
}
