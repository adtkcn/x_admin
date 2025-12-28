package tpl_utils

import (
	"archive/zip"
	"bytes"
	"embed"
	"fmt"
	"io"
	"path"
	"text/template"
	"x_admin/core/response"
	"x_admin/model/gen_model"
	"x_admin/util"
	"x_admin/util/convert_util"
)

var TemplateUtil = templateUtil{

	tpl: template.New("").Delims("{{{", "}}}").Funcs(
		template.FuncMap{
			"sub":              sub,
			"slice":            slice,
			"toSnakeCase":      util.StringUtil.ToSnakeCase,
			"toCamelCase":      util.StringUtil.ToCamelCase,
			"toUpperCamelCase": util.StringUtil.ToUpperCamelCase,
			"contains":         util.ToolsUtil.Contains,
			"goToTsType":       GenUtil.GoToTsType,
			// "goToParamType":     GenUtil.GoToParamType,

			// "goWithRespType":   GenUtil.GoWithRespType,
			"getPageResp":      GenUtil.GetPageResp,
			"nameToPath":       GenUtil.NameToPath,
			"pathToName":       GenUtil.PathToName,
			"deletePathPrefix": GenUtil.DeletePathPrefix,
			"toSqlType":        GenUtil.ToSqlType,
			"makeID":           GenUtil.MakeID,
		}),
}

// sub 模板-减函数
func sub(a, b int) int {
	return a - b
}

// slice 模板-创建切片
func slice(items ...interface{}) []interface{} {
	return items
}

// zFile 待加入zip的文件
type zFile struct {
	Name string
	Body string
}

// ExtentGenTableColumn 扩展代码生成表列实体
type ExtentGenTableColumn struct {
	ID            string //主键ID
	ColumnName    string //列名称
	ColumnComment string //列描述
	ColumnLength  int    //列长度
	ColumnType    string //列类型

	IsPk        uint8  //是否主键: [1=是, 0=否]
	IsIncrement uint8  //是否自增: [1=是, 0=否]
	IsRequired  uint8  //是否必填: [1=是, 0=否]
	IsInsert    uint8  //是否为插入字段: [1=是, 0=否]
	IsEdit      uint8  //是否编辑字段: [1=是, 0=否]
	IsList      uint8  //是否列表字段: [1=是, 0=否]
	IsListShow  uint8  //表格列表是否显示字段: [1=是, 0=否]
	IsUid       uint8  //是否用户ID: [1=是, 0=否]
	IsQuery     uint8  //是否查询字段: [1=是, 0=否]
	QueryType   string //查询方式: [等于、不等于、大于、小于、范围]
	HtmlType    string //显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]
	DictType    string //字典类型
	ListAllApi  string //列表数据来源
	Sort        int    //排序编号

	GoType     string //go类型
	GoNullType string //go空类型
	GoField    string //go字段名

	TsType  string //TS类型
	TsField string //TS字段名

	SwagType string //swagger类型

	TableColumnProp string //表格列表prop属性,默认赋值TsField,用户名称赋值xxxUser.nickname、xxxUser.username
}

// TplVars 模板变量
type TplVars struct {
	GenTpl           string
	TableName        string
	AuthorName       string
	PackageName      string
	EntityName       string
	EntitySnakeName  string
	ModuleName       string
	FunctionName     string
	DateFields       []string
	PrimaryKey       string
	PrimaryField     string
	PrimaryKeyGoType string
	AllFields        []string
	SubPriCol        ExtentGenTableColumn
	SubPriField      string
	SubTableFields   []string
	ListFields       []string
	DetailFields     []string
	DictFields       []string
	ListAllFields    []string
	IsSearch         bool
	ModelOprMap      map[string]string
	Table            gen_model.GenTable
	Columns          []ExtentGenTableColumn
	SubColumns       []ExtentGenTableColumn
	//ModelTypeMap    map[string]string
}

// genUtil 模板工具
type templateUtil struct {
	tpl *template.Template
}

/**
 * PrepareVars 获取模板变量信息
 * @param table 表信息
 * @param columns 列信息
 * @param oriSubPriCol 子表主键列信息
 * @param oriSubCols 子表列信息
 * @returns 模板变量
 */
func (tu templateUtil) PrepareVars(table gen_model.GenTable, columns []gen_model.GenTableColumn,
	oriSubPriCol gen_model.GenTableColumn, oriSubCols []gen_model.GenTableColumn) TplVars {
	// subPriField := "id"
	isSearch := false
	primaryKey := "id"
	primaryKeyGoType := "string"
	primaryField := "id"
	functionName := "【请填写功能名称】"
	var allFields []string
	// var subTableFields []string
	var listFields []string
	var detailFields []string
	var dictFields []string
	var listAllFields []string

	var newColumns = []ExtentGenTableColumn{}
	convert_util.Copy(&newColumns, columns)

	// var newSubColumns = []ExtentGenTableColumn{}
	// convert_util.Copy(subColumns, &newSubColumns)

	// var oriSubColNames []string
	// for _, column := range oriSubCols {
	// 	oriSubColNames = append(oriSubColNames, column.ColumnName)
	// }
	// var subColumns []ExtentGenTableColumn
	// if oriSubPriCol.ID != "" {
	// 	// subPriField = oriSubPriCol.ColumnName
	// 	// subColumns = append(subColumns, oriSubPriCol)
	// 	convert_util.Copy(&subColumns, oriSubPriCol)
	// }

	var userFiled = []ExtentGenTableColumn{}

	for i, column := range newColumns {
		newColumns[i].IsListShow = 1

		// 获取用户字段
		for _, columnName := range SqlConstants.ColumnNameUserFiled {
			if column.ColumnName == columnName {
				userFiled = append(userFiled, column)
				// 列表不显示用户id
				newColumns[i].IsListShow = 0
				// 是用户ID
				newColumns[i].IsUid = 1
			}
		}
		newColumns[i].GoNullType = GenUtil.GoTypeToNullType(column.GoType)

		newColumns[i].TsType = GenUtil.GoToTsType(column.GoType)
		// newColumns[i].TsField = column.ColumnName
		newColumns[i].TsField = column.GoField
		newColumns[i].TableColumnProp = newColumns[i].TsField
		newColumns[i].SwagType = GenUtil.GoTypeToSwagType(column.GoType)
	}

	if len(userFiled) > 0 {
		// 添加可查询字段
		for _, column := range userFiled {
			CreatedByColumns := []ExtentGenTableColumn{

				{
					// 查询用户名称
					ColumnName:    column.ColumnName + "_nickname",
					ColumnComment: column.ColumnComment + "名称",
					ColumnLength:  32,
					ColumnType:    "char",
					GoType:        GoConstants.TypeString,
					GoNullType:    GenUtil.GoTypeToNullType(column.GoType),

					GoField:         column.GoField + "Nickname",
					TsType:          "string",
					TsField:         column.GoField + "Nickname",
					TableColumnProp: "",
					SwagType:        SwagTypeConstants.String,

					IsPk:        0,
					IsIncrement: 0,
					IsRequired:  0,
					IsInsert:    0,
					IsEdit:      0,
					IsList:      0,
					IsListShow:  0,
					IsQuery:     1,
					QueryType:   "LIKE",
					HtmlType:    "input",
					DictType:    "",
					ListAllApi:  "",
					Sort:        6,
				}, {
					// 查询用户账号（一般不用）
					ColumnName:    column.ColumnName + "_username",
					ColumnComment: column.ColumnComment + "账号",
					ColumnLength:  32,
					ColumnType:    "char",
					GoType:        GoConstants.TypeString,
					GoNullType:    GenUtil.GoTypeToNullType(column.GoType),

					GoField: column.GoField + "Username",

					TsType:   "string",
					TsField:  column.GoField + "Username",
					SwagType: SwagTypeConstants.String,

					IsPk:        0,
					IsIncrement: 0,
					IsRequired:  0,
					IsInsert:    0,
					IsEdit:      0,
					IsList:      0,
					IsListShow:  0,
					IsQuery:     1,
					QueryType:   "=",
					HtmlType:    "input",
					DictType:    "",
					ListAllApi:  "",
					Sort:        6,
				}, {
					// 列表用户名称
					ColumnName:    column.ColumnName + "_user",
					ColumnComment: column.ColumnComment,
					ColumnLength:  32,
					ColumnType:    "char",
					GoType:        "systemSchema.SystemAuthAdminSimpleInfo",

					GoNullType: "systemSchema.SystemAuthAdminSimpleInfo",
					GoField:    column.GoField + "User",

					TsType:          "object",
					TsField:         column.GoField + "User",
					TableColumnProp: column.GoField + "User" + ".nickname",
					SwagType:        SwagTypeConstants.Object,
					IsPk:            0,
					IsIncrement:     0,
					IsRequired:      0,
					IsInsert:        0,
					IsEdit:          0,
					IsListShow:      1,
					IsList:          1,
					IsQuery:         0,
					QueryType:       "=",
					HtmlType:        "input",
					DictType:        "",
					ListAllApi:      "",
					Sort:            6,
				},
			}
			newColumns = append(newColumns, CreatedByColumns...)
		}

	}

	for _, column := range newColumns {
		allFields = append(allFields, column.ColumnName)
		// if util.ToolsUtil.Contains(oriSubColNames, column.ColumnName) {
		// 	subTableFields = append(subTableFields, column.ColumnName)
		// 	subColumns = append(subColumns, column)
		// }
		if column.IsList == 1 {
			listFields = append(listFields, column.ColumnName)
		}
		if column.IsEdit == 1 {
			detailFields = append(detailFields, column.ColumnName)
		}
		if column.IsQuery == 1 {
			isSearch = true
		}
		if column.IsPk == 1 {
			primaryKey = column.GoField
			primaryField = column.ColumnName
			primaryKeyGoType = column.GoType
		}
		if column.DictType != "" && !util.ToolsUtil.Contains(dictFields, column.DictType) {
			dictFields = append(dictFields, column.DictType)
		}
		if column.ListAllApi != "" && !util.ToolsUtil.Contains(listAllFields, column.ListAllApi) {
			listAllFields = append(listAllFields, column.ListAllApi)
		}
	}
	//QueryType转换查询比较运算符
	modelOprMap := map[string]string{
		"=":    "==",
		"LIKE": "like",
	}
	if table.FunctionName != "" {
		functionName = table.FunctionName
	}
	return TplVars{
		GenTpl:           table.GenTpl,
		TableName:        table.TableName,
		AuthorName:       table.AuthorName,
		PackageName:      table.ModuleName,
		EntityName:       table.EntityName,
		EntitySnakeName:  util.StringUtil.ToSnakeCase(table.EntityName),
		ModuleName:       table.ModuleName,
		FunctionName:     functionName,
		DateFields:       SqlConstants.ColumnTimeName,
		PrimaryKey:       primaryKey,
		PrimaryField:     primaryField,
		PrimaryKeyGoType: primaryKeyGoType,
		AllFields:        allFields,
		// SubPriCol:        oriSubPriCol,
		// SubPriField:      subPriField,
		// SubTableFields:   subTableFields,
		ListFields:    listFields,
		DetailFields:  detailFields,
		DictFields:    dictFields,
		ListAllFields: listAllFields,
		IsSearch:      isSearch,
		ModelOprMap:   modelOprMap,
		Table:         table,
		Columns:       newColumns,
		// SubColumns:    subColumns,
	}
}

// GetTemplatePaths 获取模板路径
func (tu templateUtil) GetTemplatePaths(genTpl string) []string {
	tplPaths := []string{
		"gocode/model.go.tpl",
		"gocode/schema.go.tpl",
		"gocode/service.go.tpl",
		"gocode/route.go.tpl",
		"gocode/controller.go.tpl",

		"vue/api.ts.tpl",
		"vue/edit.vue.tpl",
		"vue/details.vue.tpl",

		"uniapp/api.ts.tpl",
		"uniapp/edit.vue.tpl",
		"uniapp/index.vue.tpl",
		"uniapp/details.vue.tpl",
		"uniapp/search.vue.tpl",
		"uniapp/pages.json.tpl",
	}
	switch genTpl {
	case GenConstants.TplCrud:
		tplPaths = append(tplPaths, "vue/index.vue.tpl")
	case GenConstants.TplTree:
		tplPaths = append(tplPaths, "vue/index-tree.vue.tpl")
	}
	return tplPaths
}

//go:embed templates/gocode
//go:embed templates/vue
//go:embed templates/uniapp
var templatesFs embed.FS

/**
* Render 渲染模板
* @Description:
* @param tplPath 模板路径
* @param tplVars 模板变量
* @return string 渲染后内容
* @return error
 */
func (tu templateUtil) Render(tplPath string, tplVars TplVars) (res string, e error) {

	tpl, err := tu.tpl.ParseFS(templatesFs, "templates/"+tplPath)
	if e = response.CheckErr(err, "tu.tpl.ParseFS err"); e != nil {
		return "", e
	}
	buf := &bytes.Buffer{}
	fileName := path.Base(tplPath)
	err = tpl.ExecuteTemplate(buf, fileName, tplVars)

	if e = response.CheckErr(err, "tpl.ExecuteTemplate err"); e != nil {
		return "", e
	}
	return buf.String(), nil
}

// GetFilePaths 获取生成文件相对路径,返回 {文件路径:文件内容}
func (tu templateUtil) GetFilePaths(tplCodeMap map[string]string, ModuleName string) map[string]string {
	//模板文件对应的输出文件
	fmtMap := map[string]string{
		"gocode/model.go.tpl": fmt.Sprintf("server/model/%s.go", ModuleName),                   //strings.Join([]string{"server/model/", ModuleName, ".go"}, ""),
		"gocode/route.go.tpl": fmt.Sprintf("server/routes/adminRoute/%s_route.go", ModuleName), //strings.Join([]string{"server/routes/adminRoute/", ModuleName, "_route.go"}, ""),

		"gocode/schema.go.tpl":     fmt.Sprintf("server/app/schema/%s_schema.go", ModuleName),            //"server/app/schema/%s_schema.go"
		"gocode/service.go.tpl":    fmt.Sprintf("server/app/service/%s_service.go", ModuleName),          //"server/app/service/%s_service.go",
		"gocode/controller.go.tpl": fmt.Sprintf("server/app/controller/admin_ctl/%s_ctl.go", ModuleName), //"server/app/controller/admin_ctl/%s_ctl.go",

		"vue/api.ts.tpl":         fmt.Sprintf("admin/src/api/%s.ts", GenUtil.NameToPath(ModuleName)),            // "admin/src/api/%s.ts",
		"vue/edit.vue.tpl":       fmt.Sprintf("admin/src/views/%s/edit.vue", GenUtil.NameToPath(ModuleName)),    // "admin/src/views/%s/edit.vue",
		"vue/details.vue.tpl":    fmt.Sprintf("admin/src/views/%s/details.vue", GenUtil.NameToPath(ModuleName)), // "admin/src/views/%s/details.vue",
		"vue/index.vue.tpl":      fmt.Sprintf("admin/src/views/%s/index.vue", GenUtil.NameToPath(ModuleName)),   // "admin/src/views/%s/index.vue",
		"vue/index-tree.vue.tpl": fmt.Sprintf("admin/src/views/%s/index.vue", GenUtil.NameToPath(ModuleName)),   // "admin/src/views/%s/index-tree.vue",

		"uniapp/api.ts.tpl":      fmt.Sprintf("x_admin_app/api/%s.ts", GenUtil.NameToPath(ModuleName)),          // "x_admin_app/api/%s.ts",
		"uniapp/edit.vue.tpl":    fmt.Sprintf("x_admin_app/pages/%s/edit.vue", GenUtil.NameToPath(ModuleName)),  // "x_admin_app/pages/%s/edit.vue",
		"uniapp/index.vue.tpl":   fmt.Sprintf("x_admin_app/pages/%s/index.vue", GenUtil.NameToPath(ModuleName)), // "x_admin_app/pages/%s/index.vue",
		"uniapp/search.vue.tpl":  fmt.Sprintf("x_admin_app/pages/%s/search.vue", GenUtil.NameToPath(ModuleName)),
		"uniapp/details.vue.tpl": fmt.Sprintf("x_admin_app/pages/%s/details.vue", GenUtil.NameToPath(ModuleName)),
		"uniapp/pages.json.tpl":  fmt.Sprintf("x_admin_app/pages/%s/pages.json", GenUtil.NameToPath(ModuleName)),
	}
	filePath := make(map[string]string)
	for tplPath, tplCode := range tplCodeMap {
		file := fmtMap[tplPath]
		filePath[file] = tplCode
	}
	return filePath
}

func addFileToZip(zipWriter *zip.Writer, file zFile) error {
	header := &zip.FileHeader{
		Name:   file.Name,
		Method: zip.Deflate,
	}
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return response.CheckErr(err, "TemplateUtil.addFileToZip CreateHeader err")
	}
	_, err = io.WriteString(writer, file.Body)
	if err != nil {
		return response.CheckErr(err, "TemplateUtil.addFileToZip WriteString err")
	}
	return nil
}

// GenZip 生成代码压缩包
func (tu templateUtil) GenZip(zipWriter *zip.Writer, tplCodeMap map[string]string, ModuleName string) error {
	filePaths := tu.GetFilePaths(tplCodeMap, ModuleName)
	files := make([]zFile, 0)
	for file, tplCode := range filePaths {
		files = append(files, zFile{
			Name: file,
			Body: tplCode,
		})
	}
	for _, file := range files {
		err := addFileToZip(zipWriter, file)
		if err != nil {
			return response.CheckErr(err, "TemplateUtil.GenZip zipFiles err")
		}
	}
	return nil
}
