package tpl_utils

import (
	"archive/zip"
	"bytes"
	"embed"
	"io"
	"path"
	"strings"
	"text/template"
	"x_admin/core/response"
	"x_admin/model/gen_model"
	"x_admin/util"
)

var TemplateUtil = templateUtil{

	tpl: template.New("").Delims("{{{", "}}}").Funcs(
		template.FuncMap{
			"sub":         sub,
			"slice":       slice,
			"title":       strings.Title,
			"toSnakeCase": util.StringUtil.ToSnakeCase,
			"toCamelCase": util.StringUtil.ToCamelCase,
			"contains":    util.ToolsUtil.Contains,
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

// TplVars 模板变量
type TplVars struct {
	GenTpl          string
	TableName       string
	AuthorName      string
	PackageName     string
	EntityName      string
	EntitySnakeName string
	ModuleName      string
	FunctionName    string
	JavaCamelField  string
	DateFields      []string
	PrimaryKey      string
	PrimaryField    string
	AllFields       []string
	SubPriCol       gen_model.GenTableColumn
	SubPriField     string
	SubTableFields  []string
	ListFields      []string
	DetailFields    []string
	DictFields      []string
	IsSearch        bool
	ModelOprMap     map[string]string
	Table           gen_model.GenTable
	Columns         []gen_model.GenTableColumn
	SubColumns      []gen_model.GenTableColumn
	//ModelTypeMap    map[string]string
}

// genUtil 模板工具
type templateUtil struct {
	tpl *template.Template
}

// PrepareVars 获取模板变量信息
func (tu templateUtil) PrepareVars(table gen_model.GenTable, columns []gen_model.GenTableColumn,
	oriSubPriCol gen_model.GenTableColumn, oriSubCols []gen_model.GenTableColumn) TplVars {
	subPriField := "id"
	isSearch := false
	primaryKey := "id"
	primaryField := "id"
	functionName := "【请填写功能名称】"
	var allFields []string
	var subTableFields []string
	var listFields []string
	var detailFields []string
	var dictFields []string
	var subColumns []gen_model.GenTableColumn
	var oriSubColNames []string
	for _, column := range oriSubCols {
		oriSubColNames = append(oriSubColNames, column.ColumnName)
	}
	if oriSubPriCol.ID > 0 {
		subPriField = oriSubPriCol.ColumnName
		subColumns = append(subColumns, oriSubPriCol)
	}
	for _, column := range columns {
		allFields = append(allFields, column.ColumnName)
		if util.ToolsUtil.Contains(oriSubColNames, column.ColumnName) {
			subTableFields = append(subTableFields, column.ColumnName)
			subColumns = append(subColumns, column)
		}
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
		}
		if column.DictType != "" && !util.ToolsUtil.Contains(dictFields, column.DictType) {
			dictFields = append(dictFields, column.DictType)
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
		GenTpl:          table.GenTpl,
		TableName:       table.TableName,
		AuthorName:      table.AuthorName,
		PackageName:     table.ModuleName, //config.GenConfig.PackageName,
		EntityName:      table.EntityName,
		EntitySnakeName: util.StringUtil.ToSnakeCase(table.EntityName),
		ModuleName:      table.ModuleName,
		FunctionName:    functionName,
		DateFields:      SqlConstants.ColumnTimeName,
		PrimaryKey:      primaryKey,
		PrimaryField:    primaryField,
		AllFields:       allFields,
		SubPriCol:       oriSubPriCol,
		SubPriField:     subPriField,
		SubTableFields:  subTableFields,
		ListFields:      listFields,
		DetailFields:    detailFields,
		DictFields:      dictFields,
		IsSearch:        isSearch,
		ModelOprMap:     modelOprMap,
		Columns:         columns,
		SubColumns:      subColumns,
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
	}
	if genTpl == GenConstants.TplCrud {
		tplPaths = append(tplPaths, "vue/index.vue.tpl")
	} else if genTpl == GenConstants.TplTree {
		tplPaths = append(tplPaths, "vue/index-tree.vue.tpl")
	}
	return tplPaths
}

//go:embed templates/gocode
//go:embed templates/vue
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
	if e = response.CheckErr(err, "TemplateUtil.Render ParseFiles err"); e != nil {
		return "", e
	}
	buf := &bytes.Buffer{}
	fileName := path.Base(tplPath)
	err = tpl.ExecuteTemplate(buf, fileName, tplVars)
	if e = response.CheckErr(err, "TemplateUtil.Render Execute err"); e != nil {
		return "", e
	}
	return buf.String(), nil
}

// GetFilePaths 获取生成文件相对路径
func (tu templateUtil) GetFilePaths(tplCodeMap map[string]string, ModuleName string) map[string]string {
	//模板文件对应的输出文件
	fmtMap := map[string]string{
		"gocode/model.go.tpl":   strings.Join([]string{"server/model/", ModuleName, ".go"}, ""),
		"gocode/route.go.tpl":   strings.Join([]string{"server/admin/", ModuleName, "_route.go"}, ""),
		"gocode/schema.go.tpl":  strings.Join([]string{"server/admin/", ModuleName, "/", ModuleName, "_schema.go"}, ""),  //"server/admin/%s/%s_schema.go"
		"gocode/service.go.tpl": strings.Join([]string{"server/admin/", ModuleName, "/", ModuleName, "_service.go"}, ""), //"server/admin/%s/%s_service.go",
		// "server/admin/%s_route.go",
		"gocode/controller.go.tpl": strings.Join([]string{"server/admin/", ModuleName, "/", ModuleName, "_ctl.go"}, ""), //"server/admin/%s/%s_ctl.go",

		"vue/api.ts.tpl":         strings.Join([]string{"admin/src/api/", ModuleName, ".ts"}, ""),               // "admin/src/api/%s.ts",
		"vue/edit.vue.tpl":       strings.Join([]string{"admin/src/views/", ModuleName, "/edit.vue"}, ""),       // "admin/src/views/%s/edit.vue",
		"vue/index.vue.tpl":      strings.Join([]string{"admin/src/views/", ModuleName, "/index.vue"}, ""),      // "admin/src/views/%s/index.vue",
		"vue/index-tree.vue.tpl": strings.Join([]string{"admin/src/views/", ModuleName, "/index-tree.vue"}, ""), // "admin/src/views/%s/index-tree.vue",
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
