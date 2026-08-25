package system_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// SystemConfig 系统配置实体
type SystemConfig struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'uuid'"`
	Type       string      `gorm:"default:'';comment:'类型''"`
	Name       string      `gorm:"not null;default:'';comment:'键'"`
	Value      string      `gorm:"type:text;not null;default:'';comment:'值'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemConfig) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemAuthMenu 系统菜单实体
type SystemAuthMenu struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'uuid'"`
	Pid        string      `gorm:"not null;default:'';comment:'上级菜单'"`
	MenuType   string      `gorm:"not null;default:'';comment:'权限类型: M=目录，C=菜单，A=按钮''"`
	MenuName   string      `gorm:"not null;default:'';comment:'菜单名称'"`
	MenuIcon   string      `gorm:"not null;default:'';comment:'菜单图标'"`
	MenuSort   uint16      `gorm:"not null;default:0;comment:'菜单排序'"`
	Perms      string      `gorm:"not null;default:'';comment:'权限标识'"`
	Paths      string      `gorm:"not null;default:'';comment:'路由地址'"`
	Component  string      `gorm:"not null;default:'';comment:'前端组件'"`
	Selected   string      `gorm:"not null;default:'';comment:'选中路径'"`
	Params     string      `gorm:"not null;default:'';comment:'路由参数'"`
	IsCache    uint8       `gorm:"not null;default:0;comment:'是否缓存: 0=否, 1=是''"`
	IsShow     uint8       `gorm:"not null;default:1;comment:'是否显示: 0=否, 1=是'"`
	IsDisable  uint8       `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemAuthMenu) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemAuthPerm 系统角色菜单实体
type SystemAuthPerm struct {
	ID     string `gorm:"primarykey;type:char(36);comment:'uuid'"`
	RoleId string `gorm:"not null;comment:'角色ID'"`
	MenuId string `gorm:"not null;comment:'菜单ID'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemAuthPerm) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemAuthRole 系统角色实体
type SystemAuthRole struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'uuid'"`
	Name       string      `gorm:"not null;default:'';comment:'角色名称''"`
	Remark     string      `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable  uint8       `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort       uint16      `gorm:"not null;default:0;comment:'角色排序'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemAuthRole) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemAuthDept 系统部门实体
type SystemAuthDept struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'uuid'"`
	Pid        string                `gorm:"not null;default:'';comment:'上级主键'"`
	Name       string                `gorm:"not null;default:'';comment:'部门名称''"`
	DutyId     string                `gorm:"null;comment:'负责人id'"`
	Duty       string                `gorm:"null;default:'';comment:'负责人名'"`
	Mobile     string                `gorm:"null;default:'';comment:'联系电话'"`
	Sort       uint16                `gorm:"not null;default:0;comment:'排序编号'"`
	IsStop     uint8                 `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime x_null.Time           `gorm:"default:0;comment:'删除时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemAuthDept) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemAuthPost 系统岗位管理
type SystemAuthPost struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'uuid'"`
	Code       string                `gorm:"not null;default:'';comment:'岗位编码''"`
	Name       string                `gorm:"not null;default:'';comment:'岗位名称''"`
	Remarks    string                `gorm:"not null;default:'';comment:'岗位备注''"`
	Sort       uint16                `gorm:"not null;default:0;comment:'岗位排序'"`
	IsStop     uint8                 `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemAuthPost) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemLogLogin 系统登录日志实体
type SystemLogLogin struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'uuid'"`
	AdminId    string      `gorm:"not null;comment:'管理员ID'"`
	Email      string      `gorm:"not null;default:'';comment:'登录邮箱'"`
	Ip         string      `gorm:"not null;default:'';comment:'登录地址'"`
	Os         string      `gorm:"not null;default:'';comment:'操作系统'"`
	Browser    string      `gorm:"not null;default:'';comment:'浏览器'"`
	Status     uint8       `gorm:"not null;default:0;comment:'操作状态: 1=成功, 0=失败'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemLogLogin) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemLogOperate 系统操作日志实体
type SystemLogOperate struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'uuid'"`
	AdminId    string      `gorm:"not null;comment:'操作人ID'"`
	Type       string      `gorm:"not null;default:'';comment:'请求类型: GET/POST/PUT'"`
	Title      string      `gorm:"default:'';comment:'操作标题'"`
	Ip         string      `gorm:"not null;default:'';comment:'请求IP'"`
	Url        string      `gorm:"not null;default:'';comment:'请求接口'"`
	Method     string      `gorm:"not null;default:'';comment:'请求方法'"`
	Args       string      `gorm:"comment:'请求参数'"`
	Error      string      `gorm:"comment:'错误信息'"`
	Status     uint8       `gorm:"not null;default:0;comment:'执行状态: 1=成功, 2=失败'"`
	StartTime  x_null.Time `gorm:"not null;default:0;comment:'开始时间'"`
	EndTime    x_null.Time `gorm:"not null;default:0;comment:'结束时间'"`
	TaskTime   int64       `gorm:"not null;default:0;comment:'执行耗时'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *SystemLogOperate) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}
