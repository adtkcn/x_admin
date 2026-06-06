package notice_service

import (
	"fmt"
	"strings"
	"time"
	"x_admin/app/model"
	"x_admin/app/model/system_model"
	"x_admin/config"
	"x_admin/core"
	"x_admin/util"
)

var NoticeService = &noticeService{}

type noticeService struct{}

// NoticePayload 通知内容
type NoticePayload struct {
	Type       string // 通知类型
	Title      string // 标题
	Content    string // 正文
	ReceiverID string // 接收人ID
	SenderID   string // 发送人ID
	URL        string // 跳转URL
	Extra      string // 扩展数据JSON
}

// Send 发送通知：写DB + WebSocket推送
func (s *noticeService) Send(payload NoticePayload) error {
	notice := model.SystemNotice{
		Type:       payload.Type,
		Title:      payload.Title,
		Content:    payload.Content,
		ReceiverID: payload.ReceiverID,
		SenderID:   payload.SenderID,
		URL:        payload.URL,
		Extra:      payload.Extra,
		IsRead:     0,
	}

	db := core.GetDB()
	if err := db.Create(&notice).Error; err != nil {
		return fmt.Errorf("创建通知失败: %w", err)
	}

	// WebSocket实时推送
	core.Ws.SendToUser(payload.ReceiverID, map[string]any{
		"type": "notice",
		"data": map[string]any{
			"id":         notice.ID,
			"noticeType": payload.Type,
			"title":      payload.Title,
			"content":    payload.Content,
			"url":        payload.URL,
			"createTime": notice.CreateTime,
		},
	})

	return nil
}

// ProcessEmailDelayPush 邮件延迟补推 — 由定时任务调用
// 扫描创建超过 EmailDelaySeconds 秒且未读的通知，对设置了邮箱的用户补发邮件
func (s *noticeService) ProcessEmailDelayPush() {
	db := core.GetDB()
	delaySeconds := config.NoticeConfig.EmailDelaySeconds
	threshold := time.Now().Add(-time.Duration(delaySeconds) * time.Second)

	// 1. 查询需要邮件补推的未读通知（通知有对应接收人且有邮箱）
	type NoticeWithEmail struct {
		NoticeID   string
		ReceiverID string
		Email      string
		Title      string
		Content    string
	}

	var notices []NoticeWithEmail
	err := db.Raw(`
		SELECT n.id AS notice_id, n.receiver_id, a.email, n.title, n.content
		FROM x_system_notice n
		INNER JOIN x_system_auth_admin a ON n.receiver_id = a.id
		LEFT JOIN x_system_notice_setting ns ON ns.admin_id = a.id AND ns.channel = 'email'
		WHERE n.is_read = 0
		  AND n.is_emailed = 0
		  AND n.create_time <= ?
		  AND a.email != ''
		  AND (ns.id IS NULL OR ns.is_enabled = 1)
		ORDER BY n.create_time ASC
		LIMIT 100
	`, threshold).Scan(&notices).Error

	if err != nil {
		core.Logger.Error("ProcessEmailDelayPush 查询失败:", err)
		return
	}

	if len(notices) == 0 {
		return
	}

	core.Logger.Info(fmt.Sprintf("邮件延迟补推: 扫描到 %d 条待补推通知", len(notices)))

	// 2. 按用户聚合，每人只发一封摘要邮件
	userNotices := make(map[string][]NoticeWithEmail)
	for _, n := range notices {
		userNotices[n.ReceiverID] = append(userNotices[n.ReceiverID], n)
	}

	for receiverID, userNoteList := range userNotices {
		if len(userNoteList) == 0 {
			continue
		}
		email := userNoteList[0].Email
		if email == "" {
			continue
		}

		// 构建邮件内容
		subject := fmt.Sprintf("【系统通知】您有 %d 条未读通知", len(userNoteList))
		var sb strings.Builder
		sb.WriteString("<h2>您好！</h2><p>您有以下未读通知：</p><ul>")
		for _, n := range userNoteList {
			sb.WriteString(fmt.Sprintf("<li><b>%s</b>：%s</li>", n.Title, n.Content))
		}
		sb.WriteString("</ul><p>请登录系统查看详情。</p>")

		// 发送邮件
		opts := util.EmailOptions{
			To:       []string{email},
			Subject:  subject,
			HTMLBody: sb.String(),
		}
		if err := util.EmailUtil.SendEmail(opts); err != nil {
			core.Logger.Error(fmt.Sprintf("邮件延迟补推失败, receiverID=%s, email=%s, err=%v", receiverID, email, err))
			continue
		}

		// 标记已发送邮件，防止重复推送
		noticeIDs := make([]string, 0, len(userNoteList))
		for _, n := range userNoteList {
			noticeIDs = append(noticeIDs, n.NoticeID)
		}
		db.Model(&model.SystemNotice{}).Where("id IN ?", noticeIDs).Update("is_emailed", 1)

		core.Logger.Info(fmt.Sprintf("邮件延迟补推成功: receiverID=%s, email=%s, 通知数=%d", receiverID, email, len(userNoteList)))
	}
}

// GetAdminsByIDs 批量获取管理员信息（含邮箱）
func (s *noticeService) GetAdminsByIDs(ids []string) (map[string]system_model.SystemAuthAdmin, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	db := core.GetDB()
	var admins []system_model.SystemAuthAdmin
	if err := db.Where("id IN ?", ids).Find(&admins).Error; err != nil {
		return nil, err
	}
	result := make(map[string]system_model.SystemAuthAdmin)
	for _, a := range admins {
		result[a.ID] = a
	}
	return result, nil
}
