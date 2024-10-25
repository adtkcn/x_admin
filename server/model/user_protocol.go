package model
import (
	"x_admin/core"
	"gorm.io/plugin/soft_delete"
)

//UserProtocol 用户协议实体
type UserProtocol struct {
                Id int `gorm:"primarykey;comment:''"` // 
                Title string `gorm:"comment:'标题'"` // 标题
                Content string `gorm:"comment:'协议内容'"` // 协议内容
                Sort core.NullFloat `gorm:"comment:'排序'"` // 排序
                IsDelete soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
                CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
                UpdateTime core.NullTime `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
                DeleteTime core.NullTime `gorm:"comment:'删除时间'"` // 删除时间
}
