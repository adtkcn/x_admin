package validation

import (
	"reflect"

	"github.com/adtkcn/x_null"
	"github.com/go-playground/validator/v10"
)

// validateValuer 将 x_null 类型转换为其底层值（对应类型或 nil），供 validator 校验
func validateValuer(field reflect.Value) any {
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
func register_null(v *validator.Validate) {
	v.RegisterCustomTypeFunc(validateValuer, x_null.String{}, x_null.Int64{}, x_null.Float64{}, x_null.Time{})
}
