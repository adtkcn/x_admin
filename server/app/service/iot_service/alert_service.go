package iot_service

import (
	"encoding/json"
	"time"
	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"gorm.io/gorm"
)

// AlertService 告警历史读写：DB 侧唯一出入口，处理层（pipeline）只负责 Redis 运行态与广播，
// 触发/恢复的落库动作统一收敛到这里，保持与 DeviceService.MarkStatus 相同的分层方向。
var AlertService = NewAlertService()

func NewAlertService() *alertService {
	return &alertService{db: core.GetDB()}
}

type alertService struct {
	db *gorm.DB
}

// Raise 触发一条新告警事件（status=1），返回记录主键供处理层写入 Redis 活动态。
// values 传解码后的属性快照，整包 JSON 存库；severity 为告警级别（1/2/3）。
// 恢复前的持续命中不再逐条写库。
func (s alertService) Raise(deviceId, sn, productId, spaceId string, values map[string]any, tsMs int64, severity uint8) (string, error) {
	bs, _ := json.Marshal(values)
	now := time.UnixMilli(tsMs)
	record := iot_model.IotAlertHistory{
		DeviceId:    deviceId,
		Sn:          sn,
		ProductId:   productId,
		SpaceId:     spaceId,
		Status:      iot_model.AlertStatusFiring,
		Severity:    severity,
		Values:      string(bs),
		TriggeredAt: now,
	}
	if e := response.CheckErr(s.db.Create(&record).Error, "创建告警记录失败"); e != nil {
		return "", e
	}
	return record.ID, nil
}

// UpdateSeverity 升级已有告警事件的级别（单事件可升级：同一条记录只把 severity 拉高，不新建）。
// 仅按主键更新，幂等；不改 status（仍在告警中）。
func (s alertService) UpdateSeverity(id string, severity uint8) error {
	if id == "" {
		return nil
	}
	return response.CheckErr(s.db.Model(&iot_model.IotAlertHistory{}).
		Where("id = ?", id).Update("severity", severity).Error, "升级告警级别失败")
}

// Recover 结束一次告警事件：置已恢复并回填恢复时间。仅按主键更新，幂等。
func (s alertService) Recover(id string, tsMs int64) error {
	if id == "" {
		return nil
	}
	return response.CheckErr(s.db.Model(&iot_model.IotAlertHistory{}).
		Where("id = ? AND status = ?", id, iot_model.AlertStatusFiring).
		Updates(map[string]any{
			"status":       iot_model.AlertStatusRecovered,
			"recovered_at": util.NullTimeUtil.ParseTime(time.UnixMilli(tsMs)),
		}).Error, "恢复告警记录失败")
}

// CloseActiveByDevice 设备删除时收尾：把该设备仍处于“告警中”的记录一并置恢复，
// 避免 Redis 活动态被清理后库里残留永不恢复的悬空记录。
func (s alertService) CloseActiveByDevice(deviceId string) error {
	if deviceId == "" {
		return nil
	}
	return response.CheckErr(s.db.Model(&iot_model.IotAlertHistory{}).
		Where("device_id = ? AND status = ?", deviceId, iot_model.AlertStatusFiring).
		Updates(map[string]any{
			"status":       iot_model.AlertStatusRecovered,
			"recovered_at": util.NullTimeUtil.Now(),
		}).Error, "收尾设备告警失败")
}

// List 告警历史分页：按触发时间倒序，回填空缺的产品/区域名称，属性 JSON 还原为 map
func (s alertService) List(page request.PageReq, listReq iot_schema.AlertListReq) (res response.PageResp, e error) {
	chain := s.db.Model(&iot_model.IotAlertHistory{})
	if listReq.DeviceId != "" {
		chain = chain.Where("device_id = ?", listReq.DeviceId)
	}
	if listReq.Sn != "" {
		chain = chain.Where("sn like ?", "%"+listReq.Sn+"%")
	}
	if listReq.ProductId != "" {
		chain = chain.Where("product_id = ?", listReq.ProductId)
	}
	if listReq.SpaceId != "" {
		chain = chain.Where("space_id = ?", listReq.SpaceId)
	}
	if listReq.Status >= 0 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	if listReq.Severity >= 0 {
		chain = chain.Where("severity = ?", listReq.Severity)
	}
	if listReq.StartTime != "" {
		chain = chain.Where("triggered_at >= ?", listReq.StartTime)
	}
	if listReq.EndTime != "" {
		chain = chain.Where("triggered_at <= ?", listReq.EndTime)
	}
	var count int64
	if e = response.CheckErr(chain.Count(&count).Error, "统计告警失败"); e != nil {
		return
	}
	var rows []iot_model.IotAlertHistory
	if e = response.CheckErr(chain.Limit(page.PageSize).Offset(page.PageSize*(page.PageNo-1)).
		Order("triggered_at desc").Find(&rows).Error, "查询告警失败"); e != nil {
		return
	}
	list := make([]iot_schema.AlertResp, 0, len(rows))
	for _, r := range rows {
		item := iot_schema.AlertResp{
			ID:          r.ID,
			DeviceId:    r.DeviceId,
			Sn:          r.Sn,
			ProductId:   r.ProductId,
			SpaceId:     r.SpaceId,
			Status:      r.Status,
			Severity:    r.Severity,
			TriggeredAt: r.TriggeredAt,
		}
		var vals map[string]any
		if r.Values != "" {
			_ = json.Unmarshal([]byte(r.Values), &vals)
		}
		item.Values = vals
		if r.RecoveredAt.Exist && r.RecoveredAt.Val != nil {
			item.RecoveredAt = r.RecoveredAt.Val.Format(util.NullTimeUtil.TimeFormat)
		}
		list = append(list, item)
	}
	s.fillNames(list)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    list,
	}, nil
}

// fillNames 按列表实际出现的 id 批量回填产品名、区域名（复用 deviceService 的 setToSlice）
func (s alertService) fillNames(list []iot_schema.AlertResp) {
	if len(list) == 0 {
		return
	}
	productSet := map[string]struct{}{}
	spaceSet := map[string]struct{}{}
	for _, a := range list {
		if a.ProductId != "" {
			productSet[a.ProductId] = struct{}{}
		}
		if a.SpaceId != "" {
			spaceSet[a.SpaceId] = struct{}{}
		}
	}
	productMap := map[string]string{}
	if ids := setToSlice(productSet); len(ids) > 0 {
		var products []iot_model.IotProduct
		s.db.Select("id, name").Where("id IN ?", ids).Find(&products)
		for _, p := range products {
			productMap[p.ID] = p.Name
		}
	}
	spaceMap := map[string]string{}
	if ids := setToSlice(spaceSet); len(ids) > 0 {
		var spaces []iot_model.IotDeviceSpace
		s.db.Select("id, name").Where("id IN ?", ids).Find(&spaces)
		for _, sp := range spaces {
			spaceMap[sp.ID] = sp.Name
		}
	}
	for i := range list {
		list[i].ProductName = productMap[list[i].ProductId]
		list[i].SpaceName = spaceMap[list[i].SpaceId]
	}
}

// Del 批量删除告警历史（硬删：机器生成的历史数据，不做软删）
func (s alertService) Del(ids []string) error {
	if len(ids) == 0 {
		return response.AssertArgumentError.SetMessage("请选择要删除的记录!")
	}
	return response.CheckErr(s.db.Where("id IN ?", ids).Delete(&iot_model.IotAlertHistory{}).Error, "删除失败")
}
