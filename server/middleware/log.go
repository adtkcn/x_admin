package middleware

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"
	"x_admin/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// requestType 请求参数类
type requestType string

const (
	RequestFile    requestType = "file"    // 文件类型
	RequestDefault requestType = "default" // 默认数据类型
)

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
		if reqMethod == "POST" {
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
				var formParams map[string]interface{}
				err := c.ShouldBindBodyWith(&formParams, binding.JSON)
				if err == nil {
					val, err := util.ToolsUtil.ObjToJson(&formParams)
					// 校验错误
					if response.IsFailWithResp(c, response.CheckErr(err, "RecordLog POST Marshal err")) {
						c.Abort()
						return
					}
					args = val
				}
			}
		} else if reqMethod == "GET" {
			// GET请求
			query := c.Request.URL.RawQuery
			if query != "" {
				args, _ = url.QueryUnescape(query)
			}
		}
		// 处理异常
		defer func() {
			if r := recover(); r != nil {
				errStr = fmt.Sprintf("%+v", r)
				status = 2
				// 结束时间
				endTime := time.Now()
				// 执行时间(毫秒)
				taskTime := endTime.UnixMilli() - startTime.UnixMilli()
				// 获取当前的用户
				adminId := config.AdminConfig.GetAdminId(c)
				urlPath := c.Request.URL.Path
				ip := c.ClientIP()
				method := c.HandlerName()
				err := core.GetDB().Create(&system_model.SystemLogOperate{
					AdminId: adminId, Type: reqMethod, Title: title, Ip: ip,
					Url: urlPath, Method: method, Args: args, Error: errStr, Status: status,
					StartTime: core.ParseTimeToTsTime(startTime),
					EndTime:   core.ParseTimeToTsTime(endTime),
					TaskTime:  taskTime,
				}).Error
				response.CheckErr(err, "RecordLog recover Create err")
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
		// 结束时间
		endTime := time.Now()
		// 执行时间(毫秒)
		taskTime := endTime.UnixMilli() - startTime.UnixMilli()
		// 获取当前的用户
		adminId := config.AdminConfig.GetAdminId(c)
		urlPath := c.Request.URL.Path
		ip := c.ClientIP()
		method := c.HandlerName()
		err := core.GetDB().Create(&system_model.SystemLogOperate{
			AdminId: adminId, Type: reqMethod, Title: title, Ip: ip,
			Url: urlPath, Method: method, Args: args, Error: errStr, Status: status,
			StartTime: core.ParseTimeToTsTime(startTime),
			EndTime:   core.ParseTimeToTsTime(endTime),
			TaskTime:  taskTime,
		}).Error
		response.CheckErr(err, "RecordLog Create err")
	}
}
