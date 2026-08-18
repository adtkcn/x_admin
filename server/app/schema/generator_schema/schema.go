package generator_schema

import (
	"time"

	"github.com/adtkcn/x_null"
)

// DbTablesReq 库表列表参数
type DbTablesReq struct {
	TableName    string `json:"table_name" form:"table_name"`           // 表名称
	TableComment string `json:"table_comment" form:"table_comment"`     // 表描述
}

// ListTableReq 生成列表参数
type ListTableReq struct {
	TableName    string    `json:"table_name" form:"table_name"`                                  // 表名称
	TableComment string    `json:"table_comment" form:"table_comment"`                            // 表描述
	StartTime    time.Time `json:"start_time" form:"start_time" time_format:"2006-01-02"`        // 开始时间
	EndTime      time.Time `json:"end_time" form:"end_time" time_format:"2006-01-02"`            // 结束时间
}

// DetailTableReq 生成详情参数
type DetailTableReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// ImportTableReq 导入表结构参数
type ImportTableReq struct {
	Tables string `json:"tables" form:"tables" binding:"required"` // 导入的表, 用","分隔
}

// SyncTableReq 同步表结构参数
type SyncTableReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// EditColumn 表编辑列
type EditColumn struct {
	ID      string `json:"id" form:"id" binding:"required"`      // 主键
	TableID string `json:"table_id" form:"table_id" binding:"required"` // 表ID

	ColumnName   string `json:"column_name" form:"column_name" binding:"required,max=200"` // 列名称
	ColumnLength uint   `json:"column_length" form:"column_length" binding:"required,max=5"` // 列长度
	ColumnType   string `json:"column_type" form:"column_type" binding:"required,max=100"` // 列类型

	GoField string `json:"go_field" form:"go_field" binding:"required,max=100"` // 字段
	GoType  string `json:"go_type" form:"go_type" binding:"required,max=100"`  // 字段类型

	ColumnComment string `json:"column_comment" form:"column_comment" binding:"required,max=200"` // 列描述

	IsPk        uint8       `json:"is_pk" form:"is_pk" binding:"oneof=0 1"`            // 是否主键: [0=否, 1=是]
	IsIncrement uint8       `json:"is_increment" form:"is_increment" binding:"oneof=0 1"`     // 是否自增: [0=否, 1=是]
	IsRequired  uint8       `json:"is_required" form:"is_required" binding:"oneof=0 1"`          // 是否必填: [0=否, 1=是]
	IsInsert    uint8       `json:"is_insert" form:"is_insert" binding:"oneof=0 1"`        // 是否新增字段: [0=否, 1=是]
	IsEdit      uint8       `json:"is_edit" form:"is_edit" binding:"oneof=0 1"`          // 是否编辑字段: [0=否, 1=是]
	IsList      uint8       `json:"is_list" form:"is_list" binding:"oneof=0 1"`          // 是否列表字段: [0=否, 1=是]
	IsQuery     uint8       `json:"is_query" form:"is_query" binding:"oneof=0 1"`         // 是否查询字段: [0=否, 1=是]
	QueryType   string      `json:"query_type" form:"query_type" binding:"required,max=30"` // 查询方式
	HtmlType    string      `json:"html_type" form:"html_type" binding:"required,max=30"`  // 表单类型
	DictType    string      `json:"dict_type" form:"dict_type" binding:"required,max=200"` // 字典类型
	ListAllApi  string      `json:"list_all_api" form:"list_all_api" binding:"max=200"`        // 下拉框数据来源listAll
	CreateTime  x_null.Time `json:"create_time" form:"create_time"`                          // 创建时间
	UpdateTime  x_null.Time `json:"update_time" form:"update_time"`                          // 更新时间
}

// EditTableReq 编辑表结构参数
type EditTableReq struct {
	ID           string `json:"id" form:"id" binding:"required"`                         // 主键
	TableName    string `json:"table_name" form:"table_name" binding:"required,min=1,max=200"`    // 表名称
	EntityName   string `json:"entity_name" form:"entity_name" binding:"required,min=1,max=200"`   // 实体名称
	TableComment string `json:"table_comment" form:"table_comment" binding:"required,min=1,max=200"` // 表描述
	AuthorName   string `json:"author_name" form:"author_name" binding:"max=100"`                  // 作者名称
	Remarks      string `json:"remarks" form:"remarks" binding:"max=60"`                      // 备注信息
	GenTpl       string `json:"gen_tpl" form:"gen_tpl" binding:"oneof=crud tree"`              // 生成模板方式: [crud=单表, tree=树表]
	ModuleName   string `json:"module_name" form:"module_name" binding:"required,min=1,max=60"`    // 生成模块名
	FunctionName string `json:"function_name" form:"function_name" binding:"required,min=1,max=60"`  // 生成功能名

	TreePrimary  string       `json:"tree_primary" form:"tree_primary"`                // 树表主键
	TreeParent   string       `json:"tree_parent" form:"tree_parent"`                 // 树表父键
	TreeName     string       `json:"tree_name" form:"tree_name"`                   // 树表名称
	SubTableName string       `json:"sub_table_name" form:"sub_table_name"`               // 子表名称
	SubTableFk   string       `json:"sub_table_fk" form:"sub_table_fk"`                 // 子表外键
	Columns      []EditColumn `json:"columns" form:"columns" binding:"required"` // 字段列表
}

// DelTableReq 删除表结构参数
type DelTableReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 主键
}

// PreviewCodeReq 预览代码参数
type PreviewCodeReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// GenCodeReq 生成代码参数
type GenCodeReq struct {
	Tables string `json:"tables" form:"tables" binding:"required"` // 生成的表, 用","分隔
}

// DownloadReq 下载代码参数
type DownloadReq struct {
	Tables string `json:"tables" form:"tables" binding:"required"` // 下载的表, 用","分隔
}

// DbTableResp 数据表返回信息
type DbTableResp struct {
	TableName    string      `json:"table_name"`    // 表的名称
	TableComment string      `json:"table_comment"` // 表的描述
	CreateTime   x_null.Time `json:"create_time"`   // 创建时间
	UpdateTime   x_null.Time `json:"update_time"`   // 更新时间
}

// GenTableResp 生成表返回信息
type GenTableResp struct {
	ID string `json:"id"` // 主键

	TableName    string      `json:"table_name"`    // 表名称
	TableComment string      `json:"table_comment"` // 表描述
	CreateTime   x_null.Time `json:"create_time"`   // 创建时间
	UpdateTime   x_null.Time `json:"update_time"`   // 更新时间
}

// GenTableBaseResp 生成表基本返回信息
type GenTableBaseResp struct {
	ID           string      `json:"id"`           // 主键
	TableName    string      `json:"table_name"`    // 表的名称
	TableComment string      `json:"table_comment"` // 表的描述
	EntityName   string      `json:"entity_name"`   // 实体名称
	AuthorName   string      `json:"author_name"`   // 作者名称
	Remarks      string      `json:"remarks"`      // 备注信息
	CreateTime   x_null.Time `json:"create_time"`   // 创建时间
	UpdateTime   x_null.Time `json:"update_time"`   // 更新时间
}

// GenTableGenResp 生成表生成返回信息
type GenTableGenResp struct {
	GenTpl string `json:"gen_tpl"` // 生成模板方式: [crud=单表, tree=树表]

	ModuleName   string `json:"module_name"`   // 生成模块名
	FunctionName string `json:"function_name"` // 生成功能名
	TreePrimary  string `json:"tree_primary"`  // 树主键字段
	TreeParent   string `json:"tree_parent"`   // 树父级字段
	TreeName     string `json:"tree_name"`     // 树显示字段
	SubTableName string `json:"sub_table_name"` // 关联表名称
	SubTableFk   string `json:"sub_table_fk"`   // 关联表外键
}

// GenColumnResp 生成列返回信息
type GenColumnResp struct {
	ID      string `json:"id"`      // 字段主键
	TableID string `json:"table_id"` // 归属表主键

	ColumnName    string `json:"column_name"`    // 字段名称
	ColumnComment string `json:"column_comment"` // 字段描述
	ColumnLength  int    `json:"column_length"`  // 字段长度
	ColumnType    string `json:"column_type"`    // 字段类型
	GoType        string `json:"go_type"`        // Go类型
	GoField       string `json:"go_field"`       // Go字段

	IsPk        uint8       `json:"is_pk"`
	IsIncrement uint8       `json:"is_increment"`
	IsRequired  uint8       `json:"is_required"` // 是否必填
	IsInsert    uint8       `json:"is_insert"`   // 是否为插入字段
	IsEdit      uint8       `json:"is_edit"`     // 是否编辑字段
	IsList      uint8       `json:"is_list"`     // 是否列表字段
	IsQuery     uint8       `json:"is_query"`    // 是否查询字段
	QueryType   string      `json:"query_type"`  // 查询方式: [等于、不等于、大于、小于、范围]
	HtmlType    string      `json:"html_type"`   // 显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]
	DictType    string      `json:"dict_type"`   // 字典类型
	ListAllApi  string      `json:"list_all_api"` // 下拉框数据来源listAll
	CreateTime  x_null.Time `json:"create_time"` // 创建时间
	UpdateTime  x_null.Time `json:"update_time"` // 更新时间
}

// GenTableDetailResp 生成表详情返回信息
type GenTableDetailResp struct {
	Base   GenTableBaseResp `json:"base"`   // 基本信息
	Gen    GenTableGenResp  `json:"gen"`    // 生成信息
	Column []GenColumnResp  `json:"column"` // 字段列表
}
