package model
import (
    "github.com/google/uuid"
    "gorm.io/gorm"
	"x_admin/core"
	"gorm.io/plugin/soft_delete"
)

//SystemCorn 定时任务实体
type SystemCorn struct {
        Id string `gorm:"column:id;type:char(36);primarykey;comment:'taskid'"`
        TaskName core.NullString `gorm:"column:task_name;type:char(100);comment:'任务名称'"`
        TaskCode core.NullString `gorm:"column:task_code;type:char(100);comment:'任务编码'"`
        CornExpr core.NullString `gorm:"column:corn_expr;type:char(100);comment:'corn表达式'"`
        Disabled core.NullInt `gorm:"column:disabled;type:tinyint(1);comment:'禁用'"`
        CreatedBy core.NullString `gorm:"column:created_by;type:char(36);comment:'创建人'"`
        IsDelete soft_delete.DeletedAt `gorm:"column:is_delete;type:tinyint(1);not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
        CreateTime core.NullTime `gorm:"column:create_time;type:datetime;autoCreateTime;comment:'创建时间'"`
        UpdateTime core.NullTime `gorm:"column:update_time;type:datetime;autoUpdateTime;comment:'更新时间'"`
        DeleteTime core.NullTime `gorm:"column:delete_time;type:datetime;comment:'删除时间'"`
}
// 自动在创建时设置 UUIDv7
func (u *SystemCorn) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.Id = id.String()
	return nil
}