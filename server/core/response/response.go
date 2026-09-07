package response

import (
	"errors"
	"fmt"
	"strconv"
	"x_admin/core"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// Response 统一响应结构
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// RespType 业务响应类型
//
// 该类型同时实现 error 接口，service 层可通过 return 直接上抛业务错误，
// 由 controller 层的统一出口（JSON / IsFail）翻译成 {code, message}。
type RespType struct {
	code    int
	message string
	data    any
}

// 预定义响应类型
var (
	Success             = RespType{code: 200, message: "成功"}
	Failed              = RespType{code: 300, message: "失败"}
	ParamsValidError    = RespType{code: 310, message: "参数校验错误"}
	ParamsTypeError     = RespType{code: 311, message: "参数类型错误"}
	RequestMethodError  = RespType{code: 312, message: "请求方法错误"}
	AssertArgumentError = RespType{code: 313, message: "断言参数错误"}
	LoginAccountError   = RespType{code: 330, message: "登录账号或密码错误"}
	LoginDisableError   = RespType{code: 331, message: "登录账号已被禁用了"}
	TokenEmpty          = RespType{code: 332, message: "token参数为空"}
	TokenInvalid        = RespType{code: 333, message: "登录失效"}
	NoPermission        = RespType{code: 403, message: "无相关权限"}
	Request404Error     = RespType{code: 404, message: "请求接口不存在"}
	Request405Error     = RespType{code: 405, message: "请求方法不允许"}
	SystemError         = RespType{code: 500, message: "系统错误"}
)

// Error 实现 error 接口
func (rt RespType) Error() string {
	return strconv.Itoa(rt.code) + ":" + rt.message
}

// SetMessage 设置消息
func (rt RespType) SetMessage(message string) RespType {
	rt.message = message
	return rt
}

// SetData 设置数据
func (rt RespType) SetData(data any) RespType {
	rt.data = data
	return rt
}

// Code 获取状态码
func (rt RespType) Code() int {
	return rt.code
}

// Msg 获取消息
func (rt RespType) Msg() string {
	return rt.message
}

// Data 获取数据
func (rt RespType) Data() any {
	return rt.data
}

// ========== 统一出口 ==========
//
// 约定一：业务响应的 HTTP 状态码恒为 200，成败一律以 body.code 表达。
// 约定二：失败响应不携带 data，避免把中间数据混入错误结果。
// 约定三：未识别的错误只落日志，对外统一返回系统错误文案，不暴露内部细节。

// JSON 统一出口：err == nil 时返回 data，否则把 err 翻译成业务响应。
//
//	data, err := XxxService.List(req)
//	response.JSON(c, data, err)
func JSON(c *gin.Context, data any, err error) {
	if err == nil {
		Send(c, Success.Code(), Success.Msg(), data)
		return
	}
	// 挂到 gin 的错误链上，供操作日志等中间件感知本次请求的失败原因
	_ = c.Error(err)
	code, msg, respData, needLog := Resolve(err)
	if needLog {
		core.Logger.Error("Response Error: " + err.Error())
	}
	Send(c, code, msg, respData)
}

// IsFail 守卫式出口：err != nil 时直接响应失败并返回 true。
//
//	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
//		return
//	}
func IsFail(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	JSON(c, nil, err)
	return true
}

// Fail 直接发送失败响应，接受 RespType 或普通 error，用于中间件等非 service 场景
func Fail(c *gin.Context, err error) {
	JSON(c, nil, err)
}

// FailMsg 以指定文案发送失败响应，使用默认业务失败码 300
func FailMsg(c *gin.Context, msg string) {
	fail := Failed.SetMessage(msg)
	_ = c.Error(fail)
	Send(c, Failed.Code(), msg, nil)
}

// Resolve 把 error 翻译成业务响应三元组。
// needLog 为 true 表示未识别的错误（需要记录日志、对外隐藏细节）。
//
// 翻译优先级：业务错误 RespType > MySQL 错误 > 记录不存在 > 其它
func Resolve(err error) (code int, msg string, data any, needLog bool) {
	// 业务错误：使用 errors.As 以支持 fmt.Errorf("xxx: %w", err) 的多层包裹
	var rt RespType
	if errors.As(err, &rt) {
		return rt.Code(), rt.Msg(), rt.Data(), false
	}
	// 数据库错误
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		resp, known := mysqlResp(mysqlErr)
		return resp.Code(), resp.Msg(), nil, !known
	}
	// 记录不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Failed.Code(), "数据不存在", nil, false
	}
	// 未知错误：对外统一文案，细节只进日志
	return SystemError.Code(), SystemError.Msg(), nil, true
}

// mysqlResp MySQL 错误码映射，known 表示是否为已识别的错误码
func mysqlResp(mysqlErr *mysql.MySQLError) (resp RespType, known bool) {
	switch mysqlErr.Number {
	case 1062:
		// 主键或唯一索引冲突
		return SystemError.SetMessage("数据已存在"), true
	case 1048:
		return SystemError.SetMessage("不能为空"), true
	case 1452:
		return SystemError.SetMessage("外键约束失败"), true
	case 1451:
		return SystemError.SetMessage("关联数据存在，不能删除"), true
	default:
		return SystemError.SetMessage("数据库错误"), false
	}
}

// ========== service 层错误构造辅助 ==========

// CheckErr 把内部错误转成业务错误上抛，template 支持 fmt 格式化参数
func CheckErr(err error, template string, args ...any) error {
	if err == nil {
		return nil
	}
	core.Logger.Error("CheckErr:", err)
	message := template
	if len(args) > 0 {
		message = fmt.Sprintf(template, args...)
	}
	return SystemError.SetMessage(message)
}

// CheckDBErr 数据库查询错误统一处理：
//
//   - 记录不存在 -> notFoundMsg（业务语义，不记日志）
//
//   - 其它错误   -> 记日志 + failMsg（对外隐藏内部细节）
//
//     if e = response.CheckDBErr(err, "岗位不存在!", "详情获取失败"); e != nil {
//     return
//     }
//
// 注意：仅用于「直接操作 DB」的场景。若 err 来自其它 service（内部已包装为业务错误），
// 应直接透传，不要再次包装。
func CheckDBErr(err error, notFoundMsg, failMsg string) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return SystemError.SetMessage(notFoundMsg)
	}
	core.Logger.Error("CheckDBErr: " + failMsg + ", err: " + err.Error())
	return SystemError.SetMessage(failMsg)
}
