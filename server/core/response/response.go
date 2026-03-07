package response

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RespType 响应类型
type RespType struct {
	code    int
	message string
	data    interface{}
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

// ========== 兼容旧代码的方法 ==========

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
func (rt RespType) SetData(data interface{}) RespType {
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
func (rt RespType) Data() interface{} {
	return rt.data
}

// IsFailWithResp 判断是否错误并响应
func IsFailWithResp(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	switch v := err.(type) {
	case RespType:
		Send(c, v.Code(), v.Msg(), v.Data())
	default:
		Send(c, 500, err.Error(), nil)
	}
	return true
}

// CheckAndRespWithData 检查错误并响应带数据
func CheckAndRespWithData(c *gin.Context, data interface{}, err error) {
	if err != nil {
		switch v := err.(type) {
		case RespType:
			Send(c, v.Code(), v.Msg(), data)
		default:
			Send(c, 500, err.Error(), data)
		}
		return
	}
	Send(c, 200, "success", data)
}

// CheckErr 检查错误
func CheckErr(err error, template string, args ...interface{}) error {
	if err != nil {
		return SystemError.SetMessage(template)
	}
	return nil
}

// CheckMysqlErr 检查 MySQL 错误
func CheckMysqlErr(err error) error {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062:
			// 主键或唯一索引冲突
			return SystemError.SetMessage("数据已存在")
		case 1048:
			return SystemError.SetMessage("不能为空")
		case 1452:
			return SystemError.SetMessage("外键约束失败")
		case 1451:
			return SystemError.SetMessage("关联数据存在，不能删除")

		default:
			return err
		}
	}
	return err
}

// CheckDBNotRecord 检查记录不存在，返回错误
func CheckDBNotRecord(err error, message string) error {
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return SystemError.SetMessage(message)
	}
	return nil
}
