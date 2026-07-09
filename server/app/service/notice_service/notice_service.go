package notice_service

import (
	"fmt"
	"strings"
	"time"
	"x_admin/app/model"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/system_schema"
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
	core.Ws.SendToUser(payload.ReceiverID, "notice", map[string]any{
		"id":         notice.ID,
		"noticeType": payload.Type,
		"title":      payload.Title,
		"content":    payload.Content,
		"url":        payload.URL,
		"createTime": notice.CreateTime,
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

// List 通知列表
func (s *noticeService) List(receiverID string, pageNo, pageSize int, listReq *system_schema.SystemNoticeListReq) ([]system_schema.SystemNoticeResp, int64, error) {
	db := core.GetDB()
	model := db.Model(&model.SystemNotice{}).Where("receiver_id = ?", receiverID)

	if listReq.Type != "" {
		model = model.Where("type = ?", listReq.Type)
	}
	if listReq.IsRead == 0 {
		model = model.Where("is_read = 0")
	} else if listReq.IsRead == 1 {
		model = model.Where("is_read = 1")
	}

	var count int64
	if err := model.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var list []system_schema.SystemNoticeResp
	offset := pageSize * (pageNo - 1)
	err := model.Order("create_time DESC").Limit(pageSize).Offset(offset).Find(&list).Error
	return list, count, err
}

// UnreadCount 未读数量
func (s *noticeService) UnreadCount(receiverID string) (int64, error) {
	db := core.GetDB()
	var count int64
	err := db.Model(&model.SystemNotice{}).Where("receiver_id = ? AND is_read = 0", receiverID).Count(&count).Error
	return count, err
}

// Read 标记单条已读（同时设置阅读时间）
func (s *noticeService) Read(noticeID, receiverID string) error {
	db := core.GetDB()
	now := util.NullTimeUtil.Now()
	return db.Model(&model.SystemNotice{}).
		Where("id = ? AND receiver_id = ? AND is_read = 0", noticeID, receiverID).
		Updates(map[string]any{"is_read": 1, "read_time": now}).Error
}

// ReadAll 全部标为已读
func (s *noticeService) ReadAll(receiverID string) error {
	db := core.GetDB()
	now := util.NullTimeUtil.Now()
	return db.Model(&model.SystemNotice{}).
		Where("receiver_id = ? AND is_read = 0", receiverID).
		Updates(map[string]any{"is_read": 1, "read_time": now}).Error
}

// Del 删除通知
func (s *noticeService) Del(noticeID, receiverID string) error {
	db := core.GetDB()
	return db.Where("id = ? AND receiver_id = ?", noticeID, receiverID).Delete(&model.SystemNotice{}).Error
}

// GetSetting 获取通知偏好
func (s *noticeService) GetSetting(adminID string) (*system_schema.SystemNoticeSettingResp, error) {
	db := core.GetDB()
	resp := &system_schema.SystemNoticeSettingResp{SiteEnabled: 1, EmailEnabled: 1}
	// 默认开启
	var settings []model.SystemNoticeSetting
	if err := db.Where("admin_id = ?", adminID).Find(&settings).Error; err != nil {
		return nil, err
	}
	for _, st := range settings {
		switch st.Channel {
		case "site":
			resp.SiteEnabled = st.IsEnabled
		case "email":
			resp.EmailEnabled = st.IsEnabled
		}
	}
	return resp, nil
}

// SaveSetting 保存通知偏好
func (s *noticeService) SaveSetting(adminID string, saveReq *system_schema.SystemNoticeSettingSaveReq) error {
	db := core.GetDB()
	channels := []struct {
		Channel string
		Enabled uint8
	}{
		{"site", saveReq.SiteEnabled},
		{"email", saveReq.EmailEnabled},
	}
	for _, ch := range channels {
		var setting model.SystemNoticeSetting
		result := db.Where("admin_id = ? AND channel = ?", adminID, ch.Channel).First(&setting)
		if result.Error != nil {
			// 不存在则创建（Select 强制写入零值字段）
			setting = model.SystemNoticeSetting{
				AdminID:   adminID,
				Channel:   ch.Channel,
				IsEnabled: ch.Enabled,
			}
			if err := db.Create(&setting).Error; err != nil {
				return err
			}
		} else {
			db.Model(&setting).Update("is_enabled", ch.Enabled)
		}
	}
	return nil
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
