package model

import (
	"x_admin/core"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MonitorErrorList 错误对应的用户记录实体
type MonitorErrorList struct {
	Id         string        `gorm:"primarykey;type:char(36);comment:'id'"` // id
	Eid        string        `gorm:"comment:'错误id'"`                        // 错误表id
	Cid        string        `gorm:"comment:'客户端id'"`                       // 客户端表id
	Width      core.NullInt  `gorm:"comment:'屏幕'"`                          // 屏幕
	Height     core.NullInt  `gorm:"comment:'屏幕高度'"`                        // 屏幕高度
	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"`         // 创建时间
}

// BeforeCreate 在创建前生成UUIDv7
func (m *MonitorErrorList) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.Id = id.String()
	return nil
}
