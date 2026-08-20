package common_service

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json/v2"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"x_admin/config"
	"x_admin/core"
	"x_admin/util"
)

type AuthData struct {
	ExpireTime string `json:"expire_time"`
	Token      string `json:"token"`
}

// PushMessage 定义推送消息结构
type PushMessage struct {
	CID       string         `json:"cid"`
	RequestID string         `json:"request_id"` //10-32位之间；如果request_id重复，会导致消息丢失
	NotifyID  int            `json:"notify_id"`  //两条消息的notify_id相同，新的消息会覆盖老的消息.0-2147483647
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Payload   map[string]any `json:"payload"`
}

// PushResponse 定义推送响应
type PushResponse struct {
	Success bool
	Message string
	Data    any
}

const (
	// Redis key：个推 auth token（所有实例共享）
	geTuiTokenRedisKey = "getui:auth:token"
	// Redis key：个推 auth 刷新分布式锁
	geTuiAuthLockKey = "lock:getui:auth"
	// 本地缓存提前刷新时间（秒）：token 距过期不足此时间时主动刷新
	localRefreshAheadSec = 120
	// Redis token 存储 TTL（秒）：个推 token 有效期通常 24h，这里设 2h 留余量
	redisTokenTTLSec = 2 * 3600
)

// localTokenCache 本地 token 缓存（单实例内高速读取）
type localTokenCache struct {
	mu         sync.RWMutex
	token      string
	expireTime int64 // unix 毫秒
}

// get 从本地缓存获取 token，返回 token 和是否有效
func (c *localTokenCache) get() (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.token == "" {
		return "", false
	}
	// 提前 localRefreshAheadSec 秒认为过期，避免使用即将过期的 token
	if c.expireTime <= time.Now().UnixMilli()+int64(localRefreshAheadSec)*1000 {
		return "", false
	}
	return c.token, true
}

// set 写入本地缓存
func (c *localTokenCache) set(token string, expireTimeMs int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = token
	c.expireTime = expireTimeMs
}

var GeTuiService = NewGeTuiService()

// NewGeTuiService 初始化
func NewGeTuiService() *geTuiService {
	return &geTuiService{
		localCache:   &localTokenCache{},
		baseURL:      config.GeTuiConfig.Host + config.GeTuiConfig.AppID,
		appKey:       config.GeTuiConfig.AppKEY,
		masterSecret: config.GeTuiConfig.MasterSecret,
		packageName:  config.GeTuiConfig.PackName,
	}
}

// geTuiService 个推服务
type geTuiService struct {
	localCache   *localTokenCache // L1: 本地缓存（进程内高速读取）
	baseURL      string
	appKey       string
	masterSecret string
	packageName  string
}

// GetAuthToken 获取个推认证 token（两级缓存：本地 → Redis → 远程请求）
func (gt *geTuiService) GetAuthToken() (string, error) {
	// ---- L1: 本地缓存 ----
	if token, ok := gt.localCache.get(); ok {
		return token, nil
	}

	// ---- L2: Redis 共享缓存 ----
	if token := gt.getTokenFromRedis(); token != "" {
		return token, nil
	}

	// ---- L3: 分布式锁 + 远程请求 ----
	return gt.refreshTokenWithLock()
}

// getTokenFromRedis 从 Redis 读取 token 并回填本地缓存
func (gt *geTuiService) getTokenFromRedis() string {
	raw := util.RedisUtil.Get(geTuiTokenRedisKey)
	if raw == "" {
		return ""
	}

	var data AuthData
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		core.Logger.Warnf("geTuiService: Redis token 解析失败: %+v", err)
		return ""
	}

	expireMs, err := strconv.ParseInt(data.ExpireTime, 10, 64)
	if err != nil || expireMs <= time.Now().UnixMilli()+int64(localRefreshAheadSec)*1000 {
		return ""
	}

	// 回填本地缓存
	gt.localCache.set(data.Token, expireMs)
	core.Logger.Debugf("geTuiService: 从 Redis 缓存获取 token")
	return data.Token
}

// refreshTokenWithLock 通过分布式锁保证多实例只有一个请求远程刷新
func (gt *geTuiService) refreshTokenWithLock() (string, error) {
	lock := util.NewRedisLock(geTuiAuthLockKey, 10*time.Second)

	if !lock.Lock() {
		// 未拿到锁，说明其他实例正在刷新；等待后从 Redis 读取
		time.Sleep(200 * time.Millisecond)
		if token := gt.getTokenFromRedis(); token != "" {
			return token, nil
		}
		// 兜底：自己请求
		return gt.doAuthRequest()
	}
	defer func() {
		if err := lock.Unlock(); err != nil {
			core.Logger.Warnf("geTuiService: 释放分布式锁失败: %+v", err)
		}
	}()

	// 拿到锁后 double-check Redis（可能其他实例刚刷新完）
	if token := gt.getTokenFromRedis(); token != "" {
		return token, nil
	}

	// 远程请求新 token
	return gt.doAuthRequest()
}

// doAuthRequest 请求个推 auth 接口获取新 token，并写入 Redis + 本地缓存
func (gt *geTuiService) doAuthRequest() (string, error) {
	// 生成签名
	timestamp := time.Now().UnixMilli()
	signStr := gt.appKey + strconv.FormatInt(timestamp, 10) + gt.masterSecret
	hash := sha256.Sum256([]byte(signStr))
	sign := hex.EncodeToString(hash[:])

	reqData := map[string]any{
		"sign":      sign,
		"timestamp": timestamp,
		"appkey":    gt.appKey,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return "", fmt.Errorf("JSON编码失败: %v", err)
	}

	resp, err := http.Post(gt.baseURL+"/auth", "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		return "", fmt.Errorf("认证请求失败: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			ExpireTime string `json:"expire_time"`
			Token      string `json:"token"`
		} `json:"data"`
	}
	if err := json.UnmarshalRead(resp.Body, &result); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}
	if result.Code != 0 {
		return "", fmt.Errorf("认证失败: %s", result.Msg)
	}

	token := result.Data.Token
	expireMs, _ := strconv.ParseInt(result.Data.ExpireTime, 10, 64)

	// 写入本地缓存
	gt.localCache.set(token, expireMs)

	// 写入 Redis（所有实例共享）
	cacheData, _ := json.Marshal(AuthData{
		ExpireTime: result.Data.ExpireTime,
		Token:      token,
	})
	util.RedisUtil.Set(geTuiTokenRedisKey, string(cacheData), redisTokenTTLSec)

	core.Logger.Infof("geTuiService: 获取新 token 成功，已写入 Redis + 本地缓存")
	return token, nil
}

// PushToSingleBatchCID 批量单推消息
func (gt *geTuiService) PushToSingleBatchCID(messages []PushMessage) ([]PushResponse, error) {

	// 限制标题和内容长度
	for i := range messages {
		if len(messages[i].Title) > 32 {
			messages[i].Title = messages[i].Title[:32]
		}
		if len(messages[i].Body) > 100 {
			messages[i].Body = messages[i].Body[:100]
		}
		if messages[i].RequestID == "" {
			messages[i].RequestID = util.ToolsUtil.RandomString(32)
		}
	}

	// 获取认证token
	token, err := gt.GetAuthToken()
	if err != nil {
		return nil, fmt.Errorf("获取token失败: %v", err)
	}

	// 将消息分批处理，每批最多200条
	batchSize := 100
	var batches [][]map[string]any

	for i := 0; i < len(messages); i += batchSize {
		end := i + batchSize
		if end > len(messages) {
			end = len(messages)
		}

		var batch []map[string]any
		for _, msg := range messages[i:end] {
			payload := msg.Payload
			if payload == nil {
				payload = make(map[string]any)
			}

			intent := fmt.Sprintf("intent://io.dcloud.unipush/?#Intent;scheme=unipush;launchFlags=0x4000000;component=%s/io.dcloud.PandoraEntry;S.UP-OL-SU=true;S.title=%s;S.content=%s;S.payload=%s;end", gt.packageName, msg.Title, msg.Body, toJSONString(payload))
			fmt.Println("intent:", intent)

			ios := map[string]any{
				"type":    "notify",
				"payload": toJSONString(payload),

				"aps": map[string]any{
					"alert": map[string]any{
						"title": msg.Title,
						"body":  msg.Body,
					},
					"content-available": 0, // 0表示普通通知消息,1表示静默推送(无通知栏消息
					"category":          "ACTIONABLE",
				},
				// "apns-collapse-id": notify_id, // 使用相同的apns-collapse-id可以覆盖之前的消息
			}
			notification := map[string]any{
				"title":      msg.Title,
				"body":       msg.Body,
				"click_type": "intent",
				"intent":     intent,
				// "notify_id":  notify_id, // 两条消息的notify_id相同，新的消息会覆盖老的消息
			}

			if msg.NotifyID != 0 {
				ios["apns-collapse-id"] = msg.NotifyID
				notification["notify_id"] = msg.NotifyID
			}
			android := map[string]any{
				"ups": map[string]any{
					"notification": notification,
				},
			}
			info := map[string]any{
				"request_id": msg.RequestID,
				"settings": map[string]any{
					"ttl": 3600000, // 消息离线时间设置，单位毫秒
					"strategy": map[string]any{
						"default": 4, // 优先走厂商通道
					},
				},
				"audience": map[string]any{
					"cid": []string{msg.CID},
				},
				"push_message": map[string]any{
					"notification": notification,
				},
				"push_channel": map[string]any{
					"ios":     ios,
					"android": android,
				},
			}
			batch = append(batch, info)
		}
		batches = append(batches, batch)
	}
	str, _ := json.Marshal(batches)
	fmt.Printf("batches: %+v\n", string(str))

	// 并发发送所有批次
	var wg sync.WaitGroup
	results := make([]PushResponse, len(batches))
	client := &http.Client{}

	for i, batch := range batches {
		wg.Add(1)
		go func(index int, batch []map[string]any) {
			defer wg.Done()

			reqData := map[string]any{
				"is_async": true,
				"msg_list": batch,
			}

			reqBody, err := json.Marshal(reqData)
			if err != nil {
				results[index] = PushResponse{
					Success: false,
					Message: fmt.Sprintf("JSON编码失败: %v", err),
				}
				return
			}

			req, err := http.NewRequest("POST", gt.baseURL+"/push/single/batch/cid", strings.NewReader(string(reqBody)))
			if err != nil {
				results[index] = PushResponse{
					Success: false,
					Message: fmt.Sprintf("创建请求失败: %v", err),
				}
				return
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("token", token)

			resp, err := client.Do(req)
			if err != nil {
				results[index] = PushResponse{
					Success: false,
					Message: fmt.Sprintf("请求失败: %v", err),
				}
				return
			}
			defer resp.Body.Close()

			var result struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
				Data any    `json:"data"`
			}

			if err := json.UnmarshalRead(resp.Body, &result); err != nil {
				results[index] = PushResponse{
					Success: false,
					Message: fmt.Sprintf("解析响应失败: %v", err),
				}
				return
			}

			success := result.Code == 0
			results[index] = PushResponse{
				Success: success,
				Message: result.Msg,
				Data:    result.Data,
			}

			fmt.Printf("推送结果: %+v\n", result)
		}(i, batch)
	}

	wg.Wait()

	return results, nil
}

// toJSONString 将对象转换为JSON字符串
func toJSONString(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}
