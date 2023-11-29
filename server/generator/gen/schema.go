package gen

import (
	"time"
	"x_admin/core"
)

// DbTablesReq 库表列表参数
type DbTablesReq struct {
	TableName    string `form:"tableName"`    // 表名称
	TableComment string `form:"tableComment"` // 表描述
}

// ListTableReq 生成列表参数
type ListTableReq struct {
	TableName    string    `form:"tableName"`                          // 表名称
	TableComment string    `form:"tableComment"`                       // 表描述
	StartTime    time.Time `form:"startTime" time_format:"2006-01-02"` // 开始时间
	EndTime      time.Time `form:"endTime" time_format:"2006-01-02"`   // 结束时间
}

// DetailTableReq 生成详情参数
type DetailTableReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// ImportTableReq 导入表结构参数
type ImportTableReq struct {
	Tables string `form:"tables" binding:"required"` // 导入的表, 用","分隔
}

// SyncTableReq 同步表结构参数
type SyncTableReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// EditColumn 表编辑列
type EditColumn struct {
	ID      uint `form:"id" binding:"required,gt=0"`      // 主键
	TableID uint `form:"tableId" binding:"required,gt=0"` // 表ID

	ColumnName   string `form:"columnName" binding:"required,max=200"` // 列名称
	ColumnLength uint   `form:"columnLength" binding:"required,max=5"` // 列长度
	ColumnType   string `form:"columnType" binding:"required,max=100"` // 列类型

	GoField string `form:"goField" binding:"required,max=100"` // 字段
	GoType  string `form:"goType" binding:"required,max=100"`  // 字段类型

	ColumnComment string `form:"columnComment" binding:"required,max=200"` // 列描述

	IsPk        uint8  `form:"isPk" binding:"oneof=0 1"`            // 是否主键: [0=否, 1=是]
	IsIncrement uint8  `form:"isIncrement" binding:"oneof=0 1"`     // 是否自增: [0=否, 1=是]
	IsRequired  uint8  `form:"isStop" binding:"oneof=0 1"`          // 是否必填: [0=否, 1=是]
	IsInsert    uint8  `form:"isInsert" binding:"oneof=0 1"`        // 是否新增字段: [0=否, 1=是]
	IsEdit      uint8  `form:"isEdit" binding:"oneof=0 1"`          // 是否编辑字段: [0=否, 1=是]
	IsList      uint8  `form:"isList" binding:"oneof=0 1"`          // 是否列表字段: [0=否, 1=是]
	IsQuery     uint8  `form:"isQuery" binding:"oneof=0 1"`         // 是否查询字段: [0=否, 1=是]
	QueryType   string `form:"queryType" binding:"required,max=30"` // 查询方式
	HtmlType    string `form:"htmlType" binding:"required,max=30"`  // 表单类型
	DictType    string `form:"dictType" binding:"required,max=200"` // 字典类型
}

// EditTableReq 编辑表结构参数
type EditTableReq struct {
	ID           uint   `form:"id" binding:"required,gt=0"`                    // 主键
	TableName    string `form:"tableName" binding:"required,min=1,max=200"`    // 表名称
	EntityName   string `form:"entityName" binding:"required,min=1,max=200"`   // 实体名称
	TableComment string `form:"tableComment" binding:"required,min=1,max=200"` // 表描述
	AuthorName   string `form:"authorName" binding:"max=100"`                  // 作者名称
	Remarks      string `form:"remarks" binding:"max=60"`                      // 备注信息
	GenTpl       string `form:"genTpl" binding:"oneof=crud tree"`              // 生成模板方式: [crud=单表, tree=树表]
	ModuleName   string `form:"moduleName" binding:"required,min=1,max=60"`    // 生成模块名
	FunctionName string `form:"functionName" binding:"required,min=1,max=60"`  // 生成功能名

	TreePrimary  string       `form:"treePrimary"`                // 树表主键
	TreeParent   string       `form:"treeParent"`                 // 树表父键
	TreeName     string       `form:"treeName"`                   // 树表名称
	SubTableName string       `form:"subTableName"`               // 子表名称
	SubTableFk   string       `form:"subTableFk"`                 // 子表外键
	Columns      []EditColumn `form:"columns" binding:"required"` // 字段列表
}

// DelTableReq 删除表结构参数
type DelTableReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键
}

// PreviewCodeReq 预览代码参数
type PreviewCodeReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// GenCodeReq 生成代码参数
type GenCodeReq struct {
	Tables string `form:"tables" binding:"required"` // 生成的表, 用","分隔
}

// DownloadReq 下载代码参数
type DownloadReq struct {
	Tables string `form:"tables" binding:"required"` // 下载的表, 用","分隔
}

// DbTableResp 数据表返回信息
type DbTableResp struct {
	TableName    string              `json:"tableName" structs:"tableName"`       // 表的名称
	TableComment string              `json:"tableComment" structs:"tableComment"` // 表的描述
	CreateTime   core.OnlyRespTsTime `json:"createTime" structs:"createTime"`     // 创建时间
	UpdateTime   core.OnlyRespTsTime `json:"updateTime" structs:"updateTime"`     // 更新时间
}

// GenTableResp 生成表返回信息
type GenTableResp struct {
	ID uint `json:"id" structs:"id"` // 主键

	TableName    string      `json:"tableName" structs:"tableName"`       // 表名称
	TableComment string      `json:"tableComment" structs:"tableComment"` // 表描述
	CreateTime   core.TsTime `json:"createTime" structs:"createTime"`     // 创建时间
	UpdateTime   core.TsTime `json:"updateTime" structs:"updateTime"`     // 更新时间
}

// GenTableBaseResp 生成表基本返回信息
type GenTableBaseResp struct {
	ID           uint        `json:"id" structs:"id"`                     // 主键
	TableName    string      `json:"tableName" structs:"tableName"`       // 表的名称
	TableComment string      `json:"tableComment" structs:"tableComment"` // 表的描述
	EntityName   string      `json:"entityName" structs:"entityName"`     // 实体名称
	AuthorName   string      `json:"authorName" structs:"authorName"`     // 作者名称
	Remarks      string      `json:"remarks" structs:"remarks"`           // 备注信息
	CreateTime   core.TsTime `json:"createTime" structs:"createTime"`     // 创建时间
	UpdateTime   core.TsTime `json:"updateTime" structs:"updateTime"`     // 更新时间
}

// GenTableGenResp 生成表生成返回信息
type GenTableGenResp struct {
	GenTpl string `json:"genTpl" structs:"genTpl"` // 生成模板方式: [crud=单表, tree=树表]

	ModuleName   string `json:"moduleName" structs:"moduleName"`     // 生成模块名
	FunctionName string `json:"functionName" structs:"functionName"` // 生成功能名
	TreePrimary  string `json:"treePrimary" structs:"treePrimary"`   // 树主键字段
	TreeParent   string `json:"treeParent" structs:"treeParent"`     // 树父级字段
	TreeName     string `json:"treeName" structs:"treeName"`         // 树显示字段
	SubTableName string `json:"subTableName" structs:"subTableName"` // 关联表名称
	SubTableFk   string `json:"subTableFk" structs:"subTableFk"`     // 关联表外键
}

// GenColumnResp 生成列返回信息
type GenColumnResp struct {
	ID      uint `json:"id" structs:"id"`           // 字段主键
	TableID uint `json:"tableId" structs:"tableId"` // 归属表主键

	ColumnName    string `json:"columnName" structs:"columnName"`       // 字段名称
	ColumnComment string `json:"columnComment" structs:"columnComment"` // 字段描述
	ColumnLength  int    `json:"columnLength" structs:"columnLength"`   // 字段长度
	ColumnType    string `json:"columnType" structs:"columnType"`       // 字段类型
	GoType        string `json:"goType" structs:"goType"`               // Go类型
	GoField       string `json:"goField" structs:"goField"`             // Go字段

	IsPk        uint8       `json:"isPk" structs:"isPk"`
	IsIncrement uint8       `json:"isIncrement" structs:"isIncrement"`
	IsRequired  uint8       `json:"isRequired" structs:"isRequired"` // 是否必填
	IsInsert    uint8       `json:"isInsert" structs:"isInsert"`     // 是否为插入字段
	IsEdit      uint8       `json:"isEdit" structs:"isEdit"`         // 是否编辑字段
	IsList      uint8       `json:"isList" structs:"isList"`         // 是否列表字段
	IsQuery     uint8       `json:"isQuery" structs:"isQuery"`       // 是否查询字段
	QueryType   string      `json:"queryType" structs:"queryType"`   // 查询方式: [等于、不等于、大于、小于、范围]
	HtmlType    string      `json:"htmlType" structs:"htmlType"`     // 显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]
	DictType    string      `json:"dictType" structs:"dictType"`     // 字典类型
	CreateTime  core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime  core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// GenTableDetailResp 生成表详情返回信息
type GenTableDetailResp struct {
	Base   GenTableBaseResp `json:"base" structs:"base"`     // 基本信息
	Gen    GenTableGenResp  `json:"gen" structs:"gen"`       // 生成信息
	Column []GenColumnResp  `json:"column" structs:"column"` // 字段列表
}
