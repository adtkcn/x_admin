package excel

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// GetExcelColumnName 根据列数生成 Excel 列名
func GetExcelColumnName(columnNumber int) string {
	columnName := ""
	for columnNumber > 0 {
		remainder := (columnNumber - 1) % 26
		columnName = string('A'+remainder) + columnName
		columnNumber = (columnNumber - 1) / 26
	}
	return columnName
}

// NormalDynamicExport 导出excel
// ** 需要在传入的结构体中的字段加上tag：excelize:"title:列头名称;index:列下标(从0开始);"
// list 需要导出的对象数组、sheet sheet名称、title 标题、isGhbj 是否设置隔行背景色
func NormalDynamicExport(list interface{}, sheet, title, fields string, isGhbj, isIgnore bool, changeHead map[string]string) (file *excelize.File, err error) {
	e := ExcelInit()
	err = ExportExcel(sheet, title, fields, isGhbj, isIgnore, list, changeHead, e)
	if err != nil {
		return
	}
	return e.F, err
}

// ExportExcel excel导出
func ExportExcel(sheet, title, fields string, isGhbj, isIgnore bool, list interface{}, changeHead map[string]string, e *Excel) (err error) {
	index, _ := e.F.GetSheetIndex(sheet)
	if index < 0 { // 如果sheet名称不存在
		e.F.NewSheet(sheet)
	}
	// 构造excel表格
	// 取目标对象的元素类型、字段类型和 tag
	dataValue := reflect.ValueOf(list)
	// 判断数据的类型
	if dataValue.Kind() != reflect.Slice {
		err = errors.New("invalid data type")
		return
	}
	// 构造表头
	endColName, dataRow, err := normalBuildTitle(e, sheet, title, fields, isIgnore, changeHead, dataValue)
	if err != nil {
		return
	}
	// 构造数据行
	err = normalBuildDataRow(e, sheet, endColName, fields, dataRow, isGhbj, isIgnore, dataValue)
	return
}

// 构造表头（endColName 最后一列的列名 dataRow 数据行开始的行号）
func normalBuildTitle(e *Excel, sheet, title, fields string, isIgnore bool, changeHead map[string]string, dataValue reflect.Value) (endColName string, dataRow int, err error) {
	dataType := dataValue.Type().Elem() // 获取导入目标对象的类型信息
	var exportTitle []ExcelTag          // 遍历目标对象的字段
	for i := 0; i < dataType.NumField(); i++ {
		var excelTag ExcelTag
		field := dataType.Field(i) // 获取字段信息和tag
		tag := field.Tag.Get(ExcelTagKey)
		if tag == "" { // 如果非导出则跳过
			continue
		}
		if fields != "" { // 选择要导出或要忽略的字段
			if isIgnore && strings.Contains(fields, field.Name+",") { // 忽略指定字段
				continue
			}
			if !isIgnore && !strings.Contains(fields, field.Name+",") { // 导出指定字段
				continue
			}
		}
		err = excelTag.GetTag(tag)
		if err != nil {
			return
		}
		// 更改指定字段的表头标题
		if changeHead != nil && changeHead[field.Name] != "" {
			excelTag.Name = changeHead[field.Name]
		}
		exportTitle = append(exportTitle, excelTag)
	}
	// 排序
	sort.Slice(exportTitle, func(i, j int) bool {
		return exportTitle[i].Index < exportTitle[j].Index
	})
	var titleRowData []interface{} // 列头行
	for i, colTitle := range exportTitle {
		endColName := GetExcelColumnName(i + 1)
		if colTitle.Width > 0 { // 根据给定的宽度设置列宽
			_ = e.F.SetColWidth(sheet, endColName, endColName, float64(colTitle.Width))
		} else {
			_ = e.F.SetColWidth(sheet, endColName, endColName, float64(20)) // 默认宽度为20
		}
		titleRowData = append(titleRowData, colTitle.Name)
	}
	endColName = GetExcelColumnName(len(titleRowData)) // 根据列数生成 Excel 列名
	if title != "" {
		dataRow = 3 // 如果有title，那么从第3行开始就是数据行，第1行是title，第2行是表头
		e.F.SetCellValue(sheet, "A1", title)
		e.F.MergeCell(sheet, "A1", endColName+"1") // 合并标题单元格
		e.F.SetCellStyle(sheet, "A1", endColName+"1", e.TitleStyle)
		e.F.SetRowHeight(sheet, 1, float64(30)) // 第一行行高
		e.F.SetRowHeight(sheet, 2, float64(30)) // 第二行行高
		e.F.SetCellStyle(sheet, "A2", endColName+"2", e.HeadStyle)
		if err = e.F.SetSheetRow(sheet, "A2", &titleRowData); err != nil {
			return
		}
	} else {
		dataRow = 2 // 如果没有title，那么从第2行开始就是数据行，第1行是表头
		e.F.SetRowHeight(sheet, 1, float64(30))
		e.F.SetCellStyle(sheet, "A1", endColName+"1", e.HeadStyle)
		if err = e.F.SetSheetRow(sheet, "A1", &titleRowData); err != nil {
			return
		}
	}
	return
}

// 构造数据行
func normalBuildDataRow(e *Excel, sheet, endColName, fields string, row int, isGhbj, isIgnore bool, dataValue reflect.Value) (err error) {
	//实时写入数据
	for i := 0; i < dataValue.Len(); i++ {
		startCol := fmt.Sprintf("A%d", row)
		endCol := fmt.Sprintf("%s%d", endColName, row)
		item := dataValue.Index(i)
		typ := item.Type()
		num := item.NumField()
		var exportRow []ExcelTag
		maxLen := 0 // 记录这一行中，数据最多的单元格的值的长度
		//遍历结构体的所有字段
		for j := 0; j < num; j++ {
			dataField := typ.Field(j) //获取到struct标签，需要通过reflect.Type来获取tag标签的值
			tagVal := dataField.Tag.Get(ExcelTagKey)
			if tagVal == "" { // 如果非导出则跳过
				continue
			}
			if fields != "" { // 选择要导出或要忽略的字段
				if isIgnore && strings.Contains(fields, dataField.Name+",") { // 忽略指定字段
					continue
				}
				if !isIgnore && !strings.Contains(fields, dataField.Name+",") { // 导出指定字段
					continue
				}
			}
			var dataCol ExcelTag
			err = dataCol.GetTag(tagVal)
			fieldData := item.FieldByName(dataField.Name) // 取字段值
			if fieldData.Type().String() == "string" {    // string类型的才计算长度
				rwsTemp := fieldData.Len() // 当前单元格内容的长度
				if rwsTemp > maxLen {      //这里取每一行中的每一列字符长度最大的那一列的字符
					maxLen = rwsTemp
				}
			}
			// 替换
			if dataCol.Replace != "" {
				split := strings.Split(dataCol.Replace, ",")
				for j := range split {
					s := strings.Split(split[j], "_") // 根据下划线进行分割，格式：需要替换的内容_替换后的内容
					value := fieldData.String()
					if strings.Contains(fieldData.Type().String(), "int") {
						value = strconv.Itoa(int(fieldData.Int()))
					} else if fieldData.Type().String() == "bool" {
						value = strconv.FormatBool(fieldData.Bool())
					} else if strings.Contains(fieldData.Type().String(), "float") {
						value = strconv.FormatFloat(fieldData.Float(), 'f', -1, 64)
					}
					if s[0] == value {
						dataCol.Value = s[1]
					}
				}
			} else {
				dataCol.Value = fieldData
			}
			if err != nil {
				return
			}
			exportRow = append(exportRow, dataCol)
		}
		// 排序
		sort.Slice(exportRow, func(i, j int) bool {
			return exportRow[i].Index < exportRow[j].Index
		})
		var rowData []interface{} // 数据列
		for _, colTitle := range exportRow {
			rowData = append(rowData, colTitle.Value)
		}
		if isGhbj && row%2 == 0 {
			_ = e.F.SetCellStyle(sheet, startCol, endCol, e.ContentStyle2)
		} else {
			_ = e.F.SetCellStyle(sheet, startCol, endCol, e.ContentStyle1)
		}
		if maxLen > 25 { // 自适应行高
			d := maxLen / 25
			f := 25 * d
			_ = e.F.SetRowHeight(sheet, row, float64(f))
		} else {
			_ = e.F.SetRowHeight(sheet, row, float64(25)) // 默认行高25
		}
		if err = e.F.SetSheetRow(sheet, startCol, &rowData); err != nil {
			return
		}
		row++
	}
	return
}

// 下载
func DownLoadExcel(fileName string, res http.ResponseWriter, file *excelize.File) {
	// 设置响应头
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Content-Disposition", "attachment; filename="+fileName+".xlsx")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	err := file.Write(res) // 写入Excel文件内容到响应体
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
