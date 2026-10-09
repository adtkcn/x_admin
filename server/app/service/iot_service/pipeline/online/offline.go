package online

import "gorm.io/gorm"

// 离线判定的共用常量与模型：快路径见 offline_fast.go，慢路径见 offline_slow.go。

// scanChunk 分批大小：既限住 IN 拼接出的 SQL 长度，也限住一次 Redis pipeline 里的命令数
const scanChunk = 500

// offlineCandidate 扫描候选：离线时长以设备列为准（新增/编辑时已冗余产品值），
// 扫描时现读库、不沿用上报侧的缓存值，改完配置下一轮就生效。
type offlineCandidate struct {
	ID         string
	Sn         string
	SpaceId    string `gorm:"column:space_id"`
	DevTimeout int    `gorm:"column:dev_timeout"`
}

// scanCandidates 按在线状态取候选设备（ids 为空表示不限定），只查设备表
func scanCandidates(db *gorm.DB, status uint8, ids []string) (rows []offlineCandidate, e error) {
	q := db.Table("x_iot_device AS d").
		Select("d.id, d.sn, d.space_id, d.offline_timeout AS dev_timeout").
		Where("d.status = ? AND d.is_delete = 0", status)
	if len(ids) > 0 {
		q = q.Where("d.id IN ?", ids)
	}
	e = q.Find(&rows).Error
	return
}
