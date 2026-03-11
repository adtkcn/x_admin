package system_model

import (
	"x_admin/core"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SystemAuthAdminRole 用户角色关联实体
type SystemAuthAdminRole struct {
	ID        string        `gorm:"primarykey;type:char(36);comment:'uuid'"`
	AdminId   string        `gorm:"not null;index:idx_admin_role;comment:'管理员ID'"`
	RoleId    string        `gorm:"not null;index:idx_admin_role;comment:'角色ID'"`
	CreatedAt core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemAuthAdminRole) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.ID = id.String()
	return nil
}

func (m *SystemAuthAdminRole) TableName() string {
	return "x_system_auth_admin_role"
}
