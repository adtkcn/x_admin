package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// requestType 请求参数类
type requestType string

const (
	RequestFile    requestType = "file"    // 文件类型
	RequestDefault requestType = "default" // 默认数据类型
)

// QueueOperateLog 操作日志队列名，消费者见 app/task
const QueueOperateLog = "operate_log"

// RecordLog 记录系统日志信息中间件
func RecordLog(title string, reqTypes ...requestType) gin.HandlerFunc {
	reqType := RequestDefault
	if len(reqTypes) > 0 {
		reqType = reqTypes[0]
	}
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 异常信息
		errStr := ""
		var status uint8 = 1 // 1=成功, 2=失败
		args := ""
		// 请求方式
		reqMethod := c.Request.Method
		// 获取请求参数
		switch reqMethod {
		case "POST":
			// POST请求
			if reqType == RequestFile {
				// 文件类型
				var filenames []string
				form, err := c.MultipartForm()
				// 校验错误
				if response.IsFailWithResp(c, response.CheckErr(err, "RecordLog MultipartForm err")) {
					c.Abort()
					return
				}
				// 获取文件列表
				for _, files := range form.File {
					for _, file := range files {
						filenames = append(filenames, file.Filename)
					}
				}
				args = strings.Join(filenames, ",")
			} else {
				//默认类型
				body, _ := io.ReadAll(c.Request.Body)
				c.Request.Body = io.NopCloser(bytes.NewReader(body))
				args = string(body)
			}
		case "GET":
			// GET请求
			query := c.Request.URL.RawQuery
			if query != "" {
				args, _ = url.QueryUnescape(query)
			}
		}

		// 写入操作日志（投递到 Redis 队列，由后台消费者落库，避免阻塞主请求）
		writeLog := func() {
			// 结束时间
			endTime := time.Now()
			// 执行时间(毫秒)
			taskTime := endTime.UnixMilli() - startTime.UnixMilli()
			// 获取当前的用户
			adminId := config.AdminConfig.GetAdminId(c)
			urlPath := c.Request.URL.Path
			ip := c.ClientIP()
			method := c.HandlerName()
			payload := OperateLogPayload{
				AdminId:   adminId,
				Type:      reqMethod,
				Title:     title,
				Ip:        ip,
				Url:       urlPath,
				Method:    method,
				Args:      args,
				Error:     errStr,
				Status:    status,
				StartTime: startTime,
				EndTime:   endTime,
				TaskTime:  taskTime,
			}
			if err := core.Queue.Enqueue(QueueOperateLog, payload); err != nil {
				core.Logger.Errorf("RecordLog Enqueue err: %v", err)
			}
		}

		// 处理异常
		defer func() {
			if r := recover(); r != nil {
				errStr = fmt.Sprintf("%+v", r)
				status = 2
				// 记录失败日志后继续抛出
				writeLog()
				core.Logger.WithOptions(zap.AddCallerSkip(2)).Infof(
					"RecordLog recover: err=[%+v]", r)
				panic(r)
			}
		}()
		// 执行方法
		c.Next()
		if len(c.Errors) > 0 {
			errStr = c.Errors.String()
			status = 2
		}
		// 写入操作日志
		writeLog()
	}
}

// OperateLogPayload 操作日志队列载荷，仅承载可 JSON 序列化的基础字段，
// 消费者侧再转换为 system_model.SystemLogOperate 落库。
type OperateLogPayload struct {
	AdminId   string
	Type      string
	Title     string
	Ip        string
	Url       string
	Method    string
	Args      string
	Error     string
	Status    uint8
	StartTime time.Time
	EndTime   time.Time
	TaskTime  int64
}
