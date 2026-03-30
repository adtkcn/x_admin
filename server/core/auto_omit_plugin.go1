package core

import (
	"reflect"

	"gorm.io/gorm"
)

// ================== 1. 定义接口 ==================
// IsExists 接口
// 如果结构体字段实现了此接口，Update 插件会调用 Exist() 方法
// 返回 true  -> 更新该字段
// 返回 false -> 跳过该字段 (不更新)
type IsExists interface {
	IsExists() bool
}

// ================== 2. 定义插件 ==================
type ExistPlugin struct{}

func (e *ExistPlugin) Name() string {
	return "existPlugin"
}

func (e *ExistPlugin) Initialize(db *gorm.DB) error {
	// 注册到 Update 和 Create 的 Before 阶段
	// 放在 "gorm:assign_record" 之前，确保在赋值前剔除字段
	if err := db.Callback().Update().Before("gorm:assign_record").Register("exist_check", processFields); err != nil {
		return err
	}
	// 如果你也希望 Create 时生效，取消下面的注释
	// if err := db.Callback().Create().Before("gorm:assign_record").Register("exist_check", processFields); err != nil {
	//     return err
	// }
	return nil
}

// ================== 3. 核心处理逻辑 ==================
func processFields(db *gorm.DB) {
	// 获取当前操作的实例反射值
	value := db.Statement.ReflectValue
	if !value.IsValid() {
		return
	}

	// 处理指针
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// 只处理结构体
	if value.Kind() != reflect.Struct {
		return
	}

	// 遍历所有字段
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		structField := db.Statement.Schema.Fields[i]

		// 跳过未改变的字段 (优化性能)
		// if !db.Statement.Changed(structField.Name) {
		// 	continue
		// }

		// 核心逻辑: 检查字段是否实现了 IsExists 接口
		if exister, ok := field.Interface().(IsExists); ok {
			if !exister.IsExists() {
				// 如果 Exist() 返回 false，从更新列表中移除
				db.Statement.Omit(structField.Name)
				Logger.Debugf("字段 %s 被标记为不存在，将被忽略", structField.Name)
			}
			// 如果返回 true，不做任何处理，GORM 会正常更新
		}

		// 如果字段没有实现 Exister 接口，不做任何处理 (保持默认行为)
	}
}
