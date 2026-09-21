package notice_service

import (
	"fmt"
	"strings"
	"time"
	"x_admin/app/model"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/queue_schema"
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

// Send 发送通知：写DB + WebSocket推送。
// send=true 时还会向延迟队列投递「邮件补推」任务（延迟 EmailDelaySeconds 后触发）；
// 因此 IsEmailed 初值即决定该通知是否参与邮件推送：
//   - send=false -> EmailStatusNotSend(-1)：本次不推送邮件；
//   - send=true  -> EmailStatusPending(0)：待推送，延迟队列到期由 PushUserEmail 接管。
func (s *noticeService) Send(needSend bool, payload NoticePayload) error {
	IsEmailed := EmailStatusNotSend
	if needSend {
		IsEmailed = EmailStatusPending
	}
	notice := model.SystemNotice{
		Type:       payload.Type,
		Title:      payload.Title,
		Content:    payload.Content,
		ReceiverID: payload.ReceiverID,
		SenderID:   payload.SenderID,
		URL:        payload.URL,
		Extra:      payload.Extra,
		IsRead:     0,
		IsEmailed:  IsEmailed, // 邮件推送状态：send=false 为 -1(不发送)，send=true 为 0(待发送，延迟队列将接管)
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
	if needSend == true {
		// 投递到延迟队列：延迟 EmailDelaySeconds 后才补推邮件（WebSocket 即时送达，邮件延迟兜底）
		delayTask := queue_schema.NoticeEmailDelayTask{
			ReceiverID:   payload.ReceiverID,
			PreviewCount: emailPushPreviewCount,
		}
		if err := core.QueueDelay.EnqueueDelay(
			queue_schema.QueueNoticeEmailDelay,
			delayTask,
			time.Duration(config.NoticeConfig.EmailDelaySeconds)*time.Second,
		); err != nil {
			core.Logger.Errorf("通知邮件延迟入队失败: receiverID=%s err=%v", payload.ReceiverID, err)
		}
	}

	return nil
}

// 邮件推送状态（model.SystemNotice.IsEmailed 字段取值）
// 流转：-1 不发送 → 0 待发送 → 1 发送中（已入 notice:email 队列）→ 2 成功 / 3 失败
//
//	-1 EmailStatusNotSend：Send(send=false)，本次不参与邮件推送
//	 0 EmailStatusPending：Send(send=true) 初值，等待延迟队列到期
//	 1 EmailStatusSending：PushUserEmail 入队后立即标记（防重复推送）
//	 2/3 EmailStatusSuccess/Failed：ProcessNoticeEmail 发送后回写最终状态
const (
	EmailStatusNotSend int8 = -1 // 不发送
	EmailStatusPending int8 = 0  // 待发送
	EmailStatusSending int8 = 1  // 发送中（已进入队列）
	EmailStatusSuccess int8 = 2  // 发送成功
	EmailStatusFailed  int8 = 3  // 发送失败
)

// 邮件补推参数（不依赖外部配置，使用常量控制）
const (
	emailPushPreviewCount = 10 // 邮件正文展示的最新未读条数
)

// emailPushNotice 用户待补推的未读通知
type emailPushNotice struct {
	NoticeID string `gorm:"column:id"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
}

// PushUserEmail 取 receiverID 的「未读且未邮件推送」通知，合成一封邮件入队，并立即标记全部 is_emailed=1。
// 拦截规则：
//  1. 查询条件含 is_read = 0 —— 已读通知不补推（查询为空直接跳过，状态不变）；
//  2. 查询条件含 is_emailed = 0 —— 已推送过的不重复；
//  3. 校验 x_system_notice_setting（channel='email'）：用户需显式开启才补推；
//     未配置或 is_enabled=0 时，将这批通知标记为 EmailStatusNotSend(-1)，避免永远停在「待发送」被误触发。
//
// 因此延迟队列到期被消费时，已读的直接跳过，关闭邮件渠道的明确标记为不发送。
func (s *noticeService) PushUserEmail(db *gorm.DB, receiverID, email string, previewCount int) {
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

	// 校验用户是否允许 email 渠道推送：未配置 setting 或 is_enabled=0 则明确标记为「不发送」(-1)，
	// 避免这批通知永远停留在 is_emailed=0（待发送）却永不发送，也防止其他路径误触发补推。
	var setting model.SystemNoticeSetting
	hasSetting := db.Model(&model.SystemNoticeSetting{}).
		Where("admin_id = ? AND channel = 'email'", receiverID).
		First(&setting).Error == nil
	if !hasSetting || setting.IsEnabled != 1 {
		if err := db.Model(&model.SystemNotice{}).
			Where("id IN ?", noticeIDs).
			Update("is_emailed", EmailStatusNotSend).Error; err != nil {
			core.Logger.Error(fmt.Sprintf("标记通知不发送失败, receiverID=%s, err=%v", receiverID, err))
		}
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
	task := queue_schema.NoticeEmailTask{
		To:        email,
		Subject:   subject,
		HTMLBody:  sb.String(),
		NoticeIDs: noticeIDs,
		CreatedAt: time.Now().Unix(),
	}
	if err := core.Queue.Enqueue(queue_schema.QueueNoticeEmail, task); err != nil {
		core.Logger.Error(fmt.Sprintf("通知邮件入队失败, receiverID=%s, email=%s, err=%v", receiverID, email, err))
		return
	}

	// 入队即标记该批通知为「发送中」(EmailStatusSending)，状态 -1/0 -> 1，
	// 既防止延迟队列/重试重复推送，也确保最终由 ProcessNoticeEmail 回写为 2/3。
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
