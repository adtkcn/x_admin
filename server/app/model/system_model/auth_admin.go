package system_model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// SystemAuthAdmin 系统管理员实体
type SystemAuthAdmin struct {
	ID       string `gorm:"primarykey;type:char(36);comment:'uuid'"`
	DeptId   string `gorm:"not null;comment:'部门ID'"`
	PostId   string `gorm:"not null;comment:'岗位ID'"`
	Username string `gorm:"not null;default:'';comment:'用户账号''"`
	Nickname string `gorm:"not null;default:'';comment:'用户昵称'"`
	Password string `gorm:"not null;default:'';comment:'用户密码'"`
	Avatar   string `gorm:"not null;default:'';comment:'用户头像'"`
	// RoleId        string                `gorm:"not null;default:'';comment:'角色主键'"`
	Salt          string                `gorm:"not null;default:'';comment:'加密盐巴'"`
	Sort          uint16                `gorm:"not null;default:0;comment:'排序编号'"`
	IsDisable     uint8                 `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	IsDelete      soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	Email         string                `gorm:"not null;default:'';comment:'邮箱'" json:"email"`
	LastLoginIp   string                `gorm:"not null;default:'';comment:'最后登录IP'"`
	LastLoginTime x_null.Time           `gorm:"default:null;comment:'最后登录时间'"`
	CreateTime    x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime    x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime    x_null.Time           `gorm:"default:null;comment:'删除时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemAuthAdmin) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.ID = id.String()
	return nil
}

// SystemAuthAdminSimple 简化版系统管理员实体,用于关联查询
type SystemAuthAdminSimple struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'uuid'"`
	Username   string                `gorm:"not null;default:'';comment:'用户账号''"`
	Nickname   string                `gorm:"not null;default:'';comment:'用户昵称'"`
	Avatar     string                `gorm:"not null;default:'';comment:'用户头像'"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'"`
}

func (m *SystemAuthAdminSimple) TableName() string {
	return "x_system_auth_admin"
}
