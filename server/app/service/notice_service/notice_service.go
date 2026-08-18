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

	"gorm.io/gorm"
)

var NoticeService = &noticeService{}

type noticeService struct{}

// NoticePayload 通知内容
type NoticePayload struct {
	Type       string // 通知类型 'success' | 'danger' | 'primary' | 'info' | 'warning'
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
		IsEmailed:  EmailStatusPending, // 新通知默认待发送（延迟补推任务会扫描）
	}

	db := core.GetDB()
	if err := db.Create(&notice).Error; err != nil {
		return fmt.Errorf("创建通知失败: %w", err)
	}

	// WebSocket实时推送
	core.Ws.SendToUser(payload.ReceiverID, "notice", map[string]any{
		"id":          notice.ID,
		"type":        notice.Type,
		"title":       payload.Title,
		"content":     payload.Content,
		"url":         payload.URL,
		"create_time": notice.CreateTime,
	})

	return nil
}

// 通知邮件补推队列名
const QueueNoticeEmail = "notice:email"

// 邮件推送状态
const (
	EmailStatusNotSend  int8 = -1 // 不发送
	EmailStatusPending  int8 = 0  // 待发送
	EmailStatusSending  int8 = 1  // 发送中（已进入队列）
	EmailStatusSuccess  int8 = 2  // 发送成功
	EmailStatusFailed   int8 = 3  // 发送失败
)

// NoticeEmailTask 通知邮件补推任务载荷（推入队列异步发送）
type NoticeEmailTask struct {
	To        string   `json:"to"`
	Subject   string   `json:"subject"`
	HTMLBody  string   `json:"html_body"`
	NoticeIDs []string `json:"notice_ids"` // 发送成功后标记 is_emailed 的通知ID
	CreatedAt int64    `json:"created_at"` // 入队时间戳（秒），用于过期判定
}

// 邮件补推参数（不依赖外部配置，使用常量控制）
const (
	emailPushBatchSize    = 200 // 每轮处理的用户数（distinct 用户分页大小）
	emailPushPreviewCount = 10  // 邮件正文展示的最新未读条数
	emailPushMaxRounds    = 50  // 单轮任务最多处理的用户批次数，防止极端情况无限扫描
)

// emailPushUser 待补推用户（含邮箱）
type emailPushUser struct {
	ReceiverID string
	Email      string
}

// emailPushNotice 用户待补推的未读通知
type emailPushNotice struct {
	NoticeID string `gorm:"column:id"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
}

// ProcessEmailDelayPush 邮件延迟补推 — 由定时任务调用
// 口径：扫描创建超过 EmailDelaySeconds 秒且未读、未邮件推送、且开启邮件渠道的用户，
// 每个用户每轮只发送【一封】邮件，正文展示其最新的 emailPushPreviewCount 条未读，
// 并附“未发送总数”；入队即把该用户全部待补推通知标记为 is_emailed=1（防重复）。
// 采用“先查 distinct 用户列表 + 游标分页，再逐用户取全量未读”的方式，避免 LIMIT 截断导致漏推。
func (s *noticeService) ProcessEmailDelayPush() {
	db := core.GetDB()
	delaySeconds := config.NoticeConfig.EmailDelaySeconds
	threshold := time.Now().Add(-time.Duration(delaySeconds) * time.Second)
	batchSize := emailPushBatchSize
	previewCount := emailPushPreviewCount

	lastReceiver := ""

	for round := 0; round < emailPushMaxRounds; round++ {
		// 1. 取一批待补推用户（去重），按 receiver_id 游标分页
		var users []emailPushUser
		userSQL := `
			SELECT DISTINCT n.receiver_id AS receiver_id, a.email AS email
			FROM x_system_notice n
			INNER JOIN x_system_auth_admin a ON n.receiver_id = a.id
			LEFT JOIN x_system_notice_setting ns ON ns.admin_id = a.id AND ns.channel = 'email'
			WHERE n.is_read = 0
			  AND n.is_emailed = 0
			  AND n.create_time <= ?
			  AND a.email != ''
			  AND (ns.id IS NULL OR ns.is_enabled = 1)`
		userArgs := []any{threshold}
		if round > 0 {
			// 游标：receiver_id 字典序大于上一页最后一条，避免漏用户
			userSQL += ` AND n.receiver_id > ?`
			userArgs = append(userArgs, lastReceiver)
		}
		userSQL += ` ORDER BY n.receiver_id ASC LIMIT ?`
		userArgs = append(userArgs, batchSize)

		if err := db.Raw(userSQL, userArgs...).Scan(&users).Error; err != nil {
			core.Logger.Error("ProcessEmailDelayPush 查询用户失败:", err)
			return
		}
		if len(users) == 0 {
			break
		}
		lastReceiver = users[len(users)-1].ReceiverID

		// 2. 逐用户取全部未读，合成一封邮件并入队
		for _, u := range users {
			if u.Email == "" {
				continue
			}
			s.pushUserEmail(db, u.ReceiverID, u.Email, previewCount)
		}

		if len(users) < batchSize {
			break // 不足一页，已是最后一批
		}
	}
}

// pushUserEmail 取 receiverID 的全部待补推未读，合成一封邮件入队，并立即标记全部 is_emailed=1
func (s *noticeService) pushUserEmail(db *gorm.DB, receiverID, email string, previewCount int) {
	var notices []emailPushNotice
	if err := db.Model(&model.SystemNotice{}).
		Where("receiver_id = ? AND is_read = 0 AND is_emailed = 0", receiverID).
		Order("create_time DESC").
		Scan(&notices).Error; err != nil {
		core.Logger.Error(fmt.Sprintf("查询用户未读通知失败, receiverID=%s, err=%v", receiverID, err))
		return
	}
	if len(notices) == 0 {
		return
	}

	// 收集全部通知ID（用于入队即标记）
	noticeIDs := make([]string, 0, len(notices))
	for _, n := range notices {
		noticeIDs = append(noticeIDs, n.NoticeID)
	}
	if len(noticeIDs) == 0 {
		return
	}

	// 正文只展示最新的 previewCount 条（notices 已按 create_time DESC）
	shown := notices
	if len(shown) > previewCount {
		shown = shown[:previewCount]
	}
	var sb strings.Builder
	sb.WriteString("<h2>您好！</h2>")
	sb.WriteString(fmt.Sprintf("<p>您有 <b>%d</b> 条未读通知", len(notices)))
	if len(notices) > previewCount {
		sb.WriteString(fmt.Sprintf("，以下展示最新的 %d 条：", previewCount))
	}
	sb.WriteString("</p><ul>")
	for _, n := range shown {
		sb.WriteString(fmt.Sprintf("<li><b>%s</b>：%s</li>", n.Title, n.Content))
	}
	sb.WriteString("</ul><p>请登录系统查看全部详情。</p>")

	subject := fmt.Sprintf("【系统通知】您有 %d 条未读通知", len(notices))
	task := NoticeEmailTask{
		To:        email,
		Subject:   subject,
		HTMLBody:  sb.String(),
		NoticeIDs: noticeIDs,
		CreatedAt: time.Now().Unix(),
	}
	if err := core.Queue.Enqueue(QueueNoticeEmail, task); err != nil {
		core.Logger.Error(fmt.Sprintf("通知邮件入队失败, receiverID=%s, email=%s, err=%v", receiverID, email, err))
		return
	}

	// 入队即标记该用户全部待补推通知为“发送中”，防止重复推送
	if err := db.Model(&model.SystemNotice{}).
		Where("id IN ?", noticeIDs).
		Update("is_emailed", EmailStatusSending).Error; err != nil {
		core.Logger.Error(fmt.Sprintf("标记通知已邮件失败, receiverID=%s, err=%v", receiverID, err))
		return
	}

	core.Logger.Info(fmt.Sprintf("通知邮件已入队: receiverID=%s, email=%s, 未读数=%d, 展示=%d",
		receiverID, email, len(notices), len(shown)))
}

// MarkEmailStatus 批量更新通知邮件推送状态（供异步发送消费者回调）
// status 取值见 EmailStatus* 常量：-1不发送 0待发送 1发送中 2发送成功 3发送失败
func (s *noticeService) MarkEmailStatus(noticeIDs []string, status int8) error {
	if len(noticeIDs) == 0 {
		return nil
	}
	db := core.GetDB()
	return db.Model(&model.SystemNotice{}).
		Where("id IN ?", noticeIDs).
		Update("is_emailed", status).Error
}

// List 通知列表
func (s *noticeService) List(receiverID string, pageNo, pageSize int, listReq *system_schema.SystemNoticeListReq) ([]system_schema.SystemNoticeResp, int64, error) {
	db := core.GetDB()
	model := db.Model(&model.SystemNotice{}).Where("receiver_id = ?", receiverID)

	if listReq.Type != "" {
		model = model.Where("type = ?", listReq.Type)
	}
	switch listReq.IsRead {
	case 0:
		model = model.Where("is_read = 0")
	case 1:
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

// GetSetting 获取通知偏好（后端定义的渠道清单 + 当前开关状态）
func (s *noticeService) GetSetting(adminId string) (*system_schema.SystemNoticeSettingResp, error) {
	db := core.GetDB()
	// 当前用户已保存的渠道开关
	enabledMap := make(map[string]uint8)
	var settings []model.SystemNoticeSetting
	if err := db.Where("admin_id = ?", adminId).Find(&settings).Error; err != nil {
		return nil, err
	}
	for _, st := range settings {
		enabledMap[st.Channel] = st.IsEnabled
	}

	// 按后端定义的渠道清单返回，未配置的使用默认值
	channels := make([]system_schema.NoticeChannelSetting, 0, len(config.NoticeConfig.Channels))
	for _, ch := range config.NoticeConfig.Channels {
		enabled, ok := enabledMap[ch.Key]
		if !ok {
			enabled = ch.DefaultEnabled
		}
		channels = append(channels, system_schema.NoticeChannelSetting{
			Key:     ch.Key,
			Label:   ch.Label,
			Enabled: enabled,
		})
	}

	return &system_schema.SystemNoticeSettingResp{Channels: channels}, nil
}

// SaveSetting 保存通知偏好（按 channel 逐个 upsert）
func (s *noticeService) SaveSetting(adminId string, saveReq *system_schema.SystemNoticeSettingSaveReq) error {
	db := core.GetDB()
	for channel, enabled := range saveReq.Settings {
		var setting model.SystemNoticeSetting
		result := db.Where("admin_id = ? AND channel = ?", adminId, channel).First(&setting)
		if result.Error != nil {
			// 不存在则创建
			setting = model.SystemNoticeSetting{
				AdminID:   adminId,
				Channel:   channel,
				IsEnabled: enabled,
			}
			if err := db.Create(&setting).Error; err != nil {
				return err
			}
		} else {
			db.Model(&setting).Update("is_enabled", enabled)
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
