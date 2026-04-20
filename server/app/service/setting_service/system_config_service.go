package setting_service

import (
	"errors"
	"x_admin/app/model/system_model"
	"x_admin/util"

	"gorm.io/gorm"
)

var SystemConfigService = systemConfigService{}

// 数据库配置操作工具
type systemConfigService struct{}

// Get 根据类型和名称获取配置字典
func (cu systemConfigService) Get(db *gorm.DB, cnfType string, names ...string) (data map[string]string, err error) {
	chain := db.Where("type = ?", cnfType)
	if len(names) > 0 {
		chain.Where("name = ?", names[0])
	}

	var configs []system_model.SystemConfig
	err = chain.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	data = make(map[string]string)
	for i := 0; i < len(configs); i++ {
		data[configs[i].Name] = configs[i].Value
	}
	return data, nil
}

// GetVal 根据类型和名称获取配置值
func (cu systemConfigService) GetVal(db *gorm.DB, cnfType string, name string, defaultVal string) (data string, err error) {
	config, err := cu.Get(db, cnfType, name)
	if err != nil {
		return data, err
	}
	data, ok := config[name]
	if !ok {
		data = defaultVal
	}
	return data, nil
}

// GetMap 根据类型和名称获取配置值(Json字符串转dict)
func (cu systemConfigService) GetMap(db *gorm.DB, cnfType string, name string) (data map[string]string, err error) {
	val, err := cu.GetVal(db, cnfType, name, "")
	if err != nil {
		return data, err
	}
	if val == "" {
		return map[string]string{}, nil
	}
	err = util.ToolsUtil.JsonToObj(val, &data)
	return data, err
}

// Set 设置配置的值
// db *gorm.DB  GORM 数据库连接实例
// cnfType string  配置的类型
// name string  配置的名称
// val string  要设置的配置值
func (cu systemConfigService) Set(db *gorm.DB, cnfType string, name string, val string) (err error) {
	var config system_model.SystemConfig
	err = db.Where("type = ? AND name = ?", cnfType, name).First(&config).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		config.Type = cnfType
		config.Name = name
		config.Value = val
		if err = db.Create(&config).Error; err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	if err = db.Model(&config).Update("value", val).Error; err != nil {
		return err
	}
	return nil
}
