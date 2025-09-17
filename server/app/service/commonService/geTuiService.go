package commonService

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"x_admin/config"
	"x_admin/util"
)

type AuthData struct {
	ExpireTime string `json:"expire_time"`
	Token      string `json:"token"`
}

// PushMessage 定义推送消息结构
type PushMessage struct {
	CID       string                 `json:"cid"`
	RequestID string                 `json:"request_id"` //10-32位之间；如果request_id重复，会导致消息丢失
	NotifyID  int                    `json:"notify_id"`  //两条消息的notify_id相同，新的消息会覆盖老的消息.0-2147483647
	Title     string                 `json:"title"`
	Body      string                 `json:"body"`
	Payload   map[string]interface{} `json:"payload"`
}

// PushResponse 定义推送响应
type PushResponse struct {
	Success bool
	Message string
	Data    interface{}
}

var GeTuiService = NewGeTuiService()

// NewGeTuiService 初始化
func NewGeTuiService() *geTuiService {
	return &geTuiService{
		authData: AuthData{
			ExpireTime: "0",
			Token:      "",
		},
		authLock: sync.Mutex{},

		// appSecret    : config.GeTuiConfig.APPSECRET,
		baseURL:      config.GeTuiConfig.HOST + config.GeTuiConfig.APPID,
		appKey:       config.GeTuiConfig.APPKEY,
		masterSecret: config.GeTuiConfig.MASTERSECRET,
		packageName:  config.GeTuiConfig.PackName,
	}
}

// indexService 主页服务实现类
type geTuiService struct {
	authData AuthData
	authLock sync.Mutex

	// appSecret    string
	baseURL      string
	appKey       string
	masterSecret string
	packageName  string
}

// GetAuthToken 获取个推认证token
func (gt *geTuiService) GetAuthToken() (string, error) {

	gt.authLock.Lock()
	defer gt.authLock.Unlock()

	// 检查缓存token是否有效（提前1分钟过期）
	if gt.authData.Token != "" {
		expireTime, err := strconv.ParseInt(gt.authData.ExpireTime, 10, 64)
		if err == nil && expireTime > time.Now().UnixNano()/1e6+60*1000 {
			fmt.Println("获取缓存token", gt.authData)
			return gt.authData.Token, nil
		}
	}

	// 生成签名:将 appkey、timestamp、mastersecret 对应的字符串按此固定顺序拼接后，使用 SHA256 算法加密。
	timestamp := time.Now().UnixNano() / 1e6
	signStr := gt.appKey + strconv.FormatInt(timestamp, 10) + gt.masterSecret
	hash := sha256.Sum256([]byte(signStr))
	sign := hex.EncodeToString(hash[:])

	// 构建请求数据
	reqData := map[string]interface{}{
		"sign":      sign,
		"timestamp": timestamp,
		"appkey":    gt.appKey,
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return "", fmt.Errorf("JSON编码失败: %v", err)
	}

	// 发送认证请求
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

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	if result.Code != 0 {
		return "", fmt.Errorf("认证失败: %s", result.Msg)
	}

	// 更新缓存
	gt.authData = AuthData{
		ExpireTime: result.Data.ExpireTime,
		Token:      result.Data.Token,
	}

	fmt.Println("新的auth", gt.authData)
	return gt.authData.Token, nil
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
	var batches [][]map[string]interface{}

	for i := 0; i < len(messages); i += batchSize {
		end := i + batchSize
		if end > len(messages) {
			end = len(messages)
		}

		var batch []map[string]interface{}
		for _, msg := range messages[i:end] {
			payload := msg.Payload
			if payload == nil {
				payload = make(map[string]interface{})
			}

			intent := fmt.Sprintf("intent://io.dcloud.unipush/?#Intent;scheme=unipush;launchFlags=0x4000000;component=%s/io.dcloud.PandoraEntry;S.UP-OL-SU=true;S.title=%s;S.content=%s;S.payload=%s;end", gt.packageName, msg.Title, msg.Body, toJSONString(payload))
			fmt.Println("intent:", intent)

			ios := map[string]interface{}{
				"type":    "notify",
				"payload": toJSONString(payload),

				"aps": map[string]interface{}{
					"alert": map[string]interface{}{
						"title": msg.Title,
						"body":  msg.Body,
					},
					"content-available": 0, // 0表示普通通知消息,1表示静默推送(无通知栏消息
					"category":          "ACTIONABLE",
				},
				// "apns-collapse-id": notify_id, // 使用相同的apns-collapse-id可以覆盖之前的消息
			}
			notification := map[string]interface{}{
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
			android := map[string]interface{}{
				"ups": map[string]interface{}{
					"notification": notification,
				},
			}
			info := map[string]interface{}{
				"request_id": msg.RequestID,
				"settings": map[string]interface{}{
					"ttl": 3600000, // 消息离线时间设置，单位毫秒
					"strategy": map[string]interface{}{
						"default": 4, // 优先走厂商通道
					},
				},
				"audience": map[string]interface{}{
					"cid": []string{msg.CID},
				},
				"push_message": map[string]interface{}{
					"notification": notification,
				},
				"push_channel": map[string]interface{}{
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
		go func(index int, batch []map[string]interface{}) {
			defer wg.Done()

			reqData := map[string]interface{}{
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
				Code int         `json:"code"`
				Msg  string      `json:"msg"`
				Data interface{} `json:"data"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
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
func toJSONString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}
