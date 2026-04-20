package plugin

import (
	"reflect"

	"github.com/adtkcn/x_null"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// ValidateValuer 将 NullInt等类型 转换为底层值（int64 或 nil）
func ValidateValuer(field reflect.Value) any {
	switch obj := field.Interface().(type) {
	case x_null.String:
		return obj.Val
	case x_null.Int64:
		return obj.Val
	case x_null.Float64:
		return obj.Val
	case x_null.Time:
		return obj.Val
	default:
		return nil
	}
}

// 注册null类型验证
func RegisterNullValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(ValidateValuer, x_null.String{}, x_null.Int64{}, x_null.Float64{}, x_null.Time{})
	}
}
