package model

import (
	"x_admin/app/model/system_model"

	"github.com/adtkcn/x_null"
	"uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// SystemCorn 定时任务实体
type SystemCorn struct {
	Id            string                             `gorm:"column:id;type:char(36);primarykey;comment:'id'"`
	TaskName      x_null.String                      `gorm:"column:task_name;type:char(100);comment:'任务名称'"`
	TaskCode      x_null.String                      `gorm:"column:task_code;type:char(100);comment:'任务编码'"`
	CornExpr      x_null.String                      `gorm:"column:corn_expr;type:char(100);comment:'corn表达式'"`
	Status        x_null.Int64                       `gorm:"column:status;type:tinyint(1);comment:'状态'"`
	CreatedBy     x_null.String                      `gorm:"column:created_by;type:char(36);comment:'创建人'"`
	CreatedByUser system_model.SystemAuthAdminSimple `gorm:"foreignKey:CreatedBy"`
	IsDelete      soft_delete.DeletedAt              `gorm:"column:is_delete;type:tinyint(1);not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime    x_null.Time                        `gorm:"column:create_time;type:datetime;autoCreateTime;comment:'创建时间'"`
	UpdateTime    x_null.Time                        `gorm:"column:update_time;type:datetime;autoUpdateTime;comment:'更新时间'"`
	DeleteTime    x_null.Time                        `gorm:"column:delete_time;type:datetime;comment:'删除时间'"`
}

// 自动在创建时设置 UUIDv7
func (u *SystemCorn) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV7()
	u.Id = id.String()
	return nil
}
