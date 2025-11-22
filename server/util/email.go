package util

import (
	"fmt"
	"time"
	"x_admin/config"

	"github.com/wneessen/go-mail"
)

// EmailOptions 邮件内容选项
type EmailOptions struct {
	To          []string            // 收件人邮箱列表
	Cc          []string            // 可选:抄送收件人邮箱列表
	Bcc         []string            // 可选:密送收件人邮箱列表
	Subject     string              // 邮件主题
	HTMLBody    string              // 若提供，则优先使用 HTML；否则用 TextBody
	TextBody    string              // 可选:纯文本正文，若 HTMLBody 为空时使用
	Attachments []map[string]string // 本地文件路径列表
}

var EmailUtil = &emailUtil{
	Host:     config.Config.Email.Host,                                 // 如 "smtp.qq.com"
	Port:     config.Config.Email.Port,                                 // 如 465
	SSL:      config.Config.Email.SSL,                                  // 是否启用 SSL（465 用 true，587 用 false）
	Username: config.Config.Email.Username,                             // 发件人邮箱
	Password: config.Config.Email.Password,                             // 授权码或密码
	Timeout:  time.Duration(config.Config.Email.Timeout) * time.Second, // 超时时间，单位秒，默认 10s
}

type emailUtil struct {
	Host     string
	Port     int
	SSL      bool
	Username string
	Password string
	Timeout  time.Duration
}

// SendEmail 发送邮件的统一入口
func (cfg *emailUtil) SendEmail(opts EmailOptions) error {

	// 创建邮件消息
	m := mail.NewMsg()

	// 设置发件人
	if err := m.From(cfg.Username); err != nil {
		return fmt.Errorf("设置发件人失败: %w", err)
	}

	// 设置收件人
	if len(opts.To) == 0 {
		return fmt.Errorf("至少需要一个收件人")
	}
	if err := m.To(opts.To...); err != nil {
		return fmt.Errorf("设置收件人失败: %w", err)
	}

	// 设置抄送
	if len(opts.Cc) > 0 {
		if err := m.Cc(opts.Cc...); err != nil {
			return fmt.Errorf("设置抄送失败: %w", err)
		}
	}

	// 设置密送
	if len(opts.Bcc) > 0 {
		if err := m.Bcc(opts.Bcc...); err != nil {
			return fmt.Errorf("设置密送失败: %w", err)
		}
	}

	// 设置主题
	m.Subject(opts.Subject)

	// 设置正文（优先 HTML）
	if opts.HTMLBody != "" {
		m.SetBodyString(mail.TypeTextHTML, opts.HTMLBody)
	} else if opts.TextBody != "" {
		m.SetBodyString(mail.TypeTextPlain, opts.TextBody)
	} else {
		return fmt.Errorf("邮件正文不能为空")
	}

	// 添加附件
	for _, path := range opts.Attachments {
		m.AttachFile(path["path"], mail.WithFileName(path["filename"]))
	}

	// 配置客户端
	clientOpts := []mail.Option{
		mail.WithPort(cfg.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Username),
		mail.WithPassword(cfg.Password),
	}

	if cfg.SSL {
		clientOpts = append(clientOpts, mail.WithSSL())
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = 10 // 默认 10 秒
	}
	clientOpts = append(clientOpts, mail.WithTimeout(cfg.Timeout))

	// 创建客户端
	c, err := mail.NewClient(cfg.Host, clientOpts...)
	if err != nil {
		return fmt.Errorf("创建 SMTP 客户端失败: %w", err)
	}

	// 发送邮件
	if err := c.DialAndSend(m); err != nil {
		return fmt.Errorf("发送邮件失败: %w", err)
	}

	return nil
}

// func init() {

// 	opts := EmailOptions{
// 		To:       []string{"11675084@qq.com"},
// 		Cc:       []string{},
// 		Bcc:      []string{},
// 		Subject:  "【系统通知】订单已发货",
// 		HTMLBody: `<h2>您好！</h2><p>您的订单 <b>#222</b></p>`,
// 		Attachments: []map[string]string{
// 			{"path": "./uploads/f702cb14a76929b184e501db1fda72d3_9873.png", "filename": "发票.png"},
// 			{"path": "./uploads/f702cb14a76929b184e501db1fda72d3_9873.png", "filename": "发货标签.png"},
// 		},
// 	}

// 	if err := EmailUtil.SendEmail(opts); err != nil {
// 		log.Fatalf("❌ 邮件发送失败: %v", err)
// 	}

// 	log.Println("✅ 邮件发送成功！")
// }
