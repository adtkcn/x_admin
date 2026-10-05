package iot_service

import (
	"context"
	"encoding/json"
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/redis/go-redis/v9"
)

// 本文件服务于设备接入层：鉴权信息与物模型的 Redis 缓存、在线状态的最小写入。
// key 统一为 {RedisPrefix}iot:<用途>:<标识>，与 pipeline 的运行态 key 同前缀不同用途。

// iotCacheKey 拼接 IoT 缓存键
func iotCacheKey(suffix string) string {
	return config.RedisConfig.RedisPrefix + "iot:" + suffix
}

// authCacheKey 鉴权缓存键：把 token 编进 key，
// 这样凭证变更自然形成新 key（旧 key 由 TTL 淘汰），无需依赖模糊匹配清理。
func authCacheKey(sn, token string) string {
	return iotCacheKey("auth:" + sn + ":" + token)
}

// authRow 鉴权回源时的联表查询载体：字段名按 GORM 的 snake_case 约定与 SELECT 别名对应。
// 独立于 iot_schema.DeviceAuthInfo，因为这里需要的是列的原始形态（is_stop 是 uint8 而非 bool）。
type authRow struct {
	DeviceId    string `gorm:"column:device_id"`
	Sn          string
	Name        string
	ProductId   string `gorm:"column:product_id"`
	SpaceId     string `gorm:"column:space_id"`
	Status      uint8
	Protocol    string
	ProductStop uint8 `gorm:"column:product_stop"`
	// 离线判定时长只读设备 offline_timeout 列（新增/编辑时已把产品值冗余到设备上）
	OfflineTimeout int `gorm:"column:offline_timeout"`
}

// DefaultOfflineTimeout 未配置时的兜底离线判定时长（秒）：设备的 offline_timeout
// 为 0 时生效，保证不会出现“无窗口可判”的设备（旧数据也能正常判离线）
const DefaultOfflineTimeout = 300

// EffectiveOfflineTimeout 回读单台设备当前生效的离线判定时长（供管理端运行态展示）。
// 不走鉴权缓存：管理端需要的是改完配置就能看到的实时值，且此接口低频。
func (s deviceService) EffectiveOfflineTimeout(id string) (timeout int, e error) {
	if id == "" {
		return 0, response.AssertArgumentError.SetMessage("设备主键不能为空!")
	}
	var device iot_model.IotDevice
	e = response.CheckDBErr(
		s.db.Select("offline_timeout").Where("id = ?", id).First(&device).Error,
		"设备不存在!", "离线时长查询失败")
	if e != nil {
		return
	}
	return device.OfflineTimeout, nil
}

// authQuery 设备+产品一次 LEFT JOIN 取回鉴权所需的全部列。
// 拆成两条查询时，缓存未命中（凭证轮换、TTL 集中到期）的那批报文会各自产生两次 DB 往返。
// 产品用 LEFT JOIN + IFNULL 而非 INNER JOIN：产品被删/未建时设备仍可鉴权，
// protocol 为空由中间件按“兼容历史数据”放行。
func (s deviceService) authQuery(sn, token string) (row authRow, e error) {
	e = response.CheckErr(
		s.db.Table("x_iot_device AS d").
			Select("d.id AS device_id, d.sn, d.name, d.product_id, d.space_id, d.status, d.offline_timeout, "+
				"IFNULL(p.protocol, '') AS protocol, IFNULL(p.is_stop, 0) AS product_stop").
			Joins("LEFT JOIN x_iot_product AS p ON p.id = d.product_id AND p.is_delete = 0").
			Where("d.sn = ? AND d.access_token = ? AND d.is_delete = 0", sn, token).
			Limit(1).Find(&row).Error, "鉴权查询失败")
	if e != nil || row.DeviceId == "" {
		// Find 不会因空结果集报错，这里补一次“记录不存在”判定，语义对齐 CheckDBErr
		if e == nil {
			e = response.SystemError.SetMessage("设备编号或接入凭证无效!")
		}
	}
	return
}

// Authenticate 设备接入鉴权：sn + access_token 校验，命中缓存则完全不查库。
// 高频上报下每条报文一次 DB 查询不可接受，故缓存 DeviceCacheTTL 秒；
// 设备编辑/删除会显式失效对应 key，避免删除后凭证在 TTL 内仍可用。
func (s deviceService) Authenticate(sn, token string) (info *iot_schema.DeviceAuthInfo, e error) {
	if sn == "" || token == "" {
		return nil, response.TokenEmpty.SetMessage("缺少设备编号或接入凭证!")
	}
	key := authCacheKey(sn, token)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if val, err := core.Redis.Get(ctx, key).Result(); err == nil && val != "" {
		var cached iot_schema.DeviceAuthInfo
		if err = json.Unmarshal([]byte(val), &cached); err == nil && cached.DeviceId != "" {
			return &cached, nil
		}
	} else if err != nil && !isRedisMiss(err) {
		// Redis 真故障不阻断鉴权，回源 DB，仅记日志
		core.Logger.Errorf("设备鉴权读缓存失败 sn=%s err=%v", sn, err)
	}

	var row authRow
	if row, e = s.authQuery(sn, token); e != nil {
		return
	}
	info = &iot_schema.DeviceAuthInfo{
		DeviceId:       row.DeviceId,
		Sn:             row.Sn,
		Name:           row.Name,
		ProductId:      row.ProductId,
		SpaceId:        row.SpaceId,
		Status:         row.Status,
		Protocol:       row.Protocol,
		ProductStop:    row.ProductStop == 1,
		OfflineTimeout: row.OfflineTimeout,
	}

	if bs, err := json.Marshal(info); err == nil {
		if err = core.Redis.Set(ctx, key, bs, time.Duration(config.IotConfig.DeviceCacheTTL)*time.Second).Err(); err != nil {
			core.Logger.Errorf("设备鉴权写缓存失败 sn=%s err=%v", sn, err)
		}
	}
	return
}

// InvalidateAuth 失效设备鉴权缓存：设备编辑（sn 变更）与删除、产品离线时长变更同步设备时调用，
// 否则已删设备的凭证在 DeviceCacheTTL 内仍能通过缓存接入。
// 注意：这份缓存里也编了设备的离线判定时长，配置变更后需失效缓存才会在上报侧即时生效；
// 离线判定扫描本身按库侧设备值现算，不受缓存滞后影响。
func (s deviceService) InvalidateAuth(sn, token string) {
	if sn == "" || token == "" {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := core.Redis.Del(ctx, authCacheKey(sn, token)).Err(); err != nil {
		core.Logger.Errorf("失效设备鉴权缓存失败 sn=%s err=%v", sn, err)
	}
}

// ProductModel 处理层需要的物模型与引用的规则链ID：
// 同属产品配置，合并成一条缓存共同回源，避免每条报文按字段各查一次库。
// 解码/过滤/告警逻辑不再落在产品列，而是由 RuleChainId 指向的可视化规则链图描述。
type ProductModel struct {
	ThingModel  string `json:"thing_model"`
	RuleChainId string `json:"rule_chain_id"`
}

// CachedProductModel 取产品物模型+引用的规则链ID（缓存 ThingCacheTTL 秒），供处理层解码与白名单裁剪。
// 一次命中即返回；库侧记录不存在时返回零值（物模型为空 → 处理层按“未配置”丢弃上报）。
func (s deviceService) CachedProductModel(productId string) ProductModel {
	if productId == "" {
		return ProductModel{}
	}
	key := iotCacheKey("model:" + productId)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if val, err := core.Redis.Get(ctx, key).Result(); err == nil {
		var m ProductModel
		if json.Unmarshal([]byte(val), &m) == nil {
			return m // 命中即返回（含“配置为空”的 {}）；未命中 err 为 redis.Nil，回源 DB
		}
	} else if !isRedisMiss(err) {
		core.Logger.Errorf("读产品模型缓存失败 productId=%s err=%v", productId, err)
	}

	var product iot_model.IotProduct
	if s.db.Select("id, thing_model, rule_chain_id").Where("id = ?", productId).First(&product).Error != nil {
		return ProductModel{}
	}
	m := ProductModel{
		ThingModel:  product.ThingModel,
		RuleChainId: product.RuleChainId,
	}
	if bs, err := json.Marshal(m); err == nil {
		if err := core.Redis.Set(ctx, key, bs, time.Duration(config.IotConfig.ThingCacheTTL)*time.Second).Err(); err != nil {
			core.Logger.Errorf("写产品模型缓存失败 productId=%s err=%v", productId, err)
		}
	}
	return m
}

// InvalidateProductCache 产品编辑/删除后失效物模型与各脚本缓存，避免处理层沿用旧配置
func (s deviceService) InvalidateProductCache(productId string) {
	if productId == "" {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := core.Redis.Del(ctx, iotCacheKey("model:"+productId)).Err(); err != nil {
		core.Logger.Errorf("失效产品模型缓存失败 productId=%s err=%v", productId, err)
	}
}

// MarkStatus 更新设备在线状态，仅在状态确实发生变化时写库（接入层高频调用，避免每条报文 UPDATE）。
// 返回 changed=true 表示发生了上线/下线转换，调用方可据此决定是否推送状态变更。
//
// 用 “WHERE id=? AND status<>?” 的条件更新一次完成判断与写入：先 SELECT 再 UPDATE 时，
// 并发的上报与离线复查可能同时读到旧状态，各写一遍库、各推一次上下线广播。
func (s deviceService) MarkStatus(id string, status uint8) (changed bool, e error) {
	if id == "" {
		return false, response.AssertArgumentError.SetMessage("设备主键不能为空!")
	}
	now := util.NullTimeUtil.Now()
	updates := map[string]any{"status": status}
	if status == 1 {
		updates["last_online_time"] = now
	} else {
		updates["last_offline_time"] = now
	}
	res := s.db.Model(&iot_model.IotDevice{}).
		Where("id = ? AND status <> ?", id, status).Updates(updates)
	if e = response.CheckErr(res.Error, "状态更新失败"); e != nil {
		return
	}
	if res.RowsAffected > 0 {
		return true, nil
	}
	// 0 行只有两种可能：状态本来就一致（并发下别人已改，正是不该重复广播的场景），
	// 或设备已被删除（软删除条件已过滤）。两者都按“未变化”返回，不必再补一次查询。
	return false, nil
}

// isRedisMiss 区分“键不存在”与 Redis 真实故障：前者回源 DB，后者也回源但记日志
func isRedisMiss(err error) bool {
	return err == nil || err == redis.Nil
}
