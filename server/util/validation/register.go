package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Register 统一配置 gin 绑定校验器：注册 x_null 自定义类型 + 中文错误翻译。
// 显示注册，需在启动阶段调用一次；对同一引擎重复调用幂等。
func Register() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}
	register_null(v)
	register_translations(v)
}
