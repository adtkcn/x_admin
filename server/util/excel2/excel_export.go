package excel2

import (
	"fmt"
	"x_admin/util/convert_util"

	"github.com/xuri/excelize/v2"
)

// GetExcelColumnName 根据列数生成 Excel 列名
func GetExcelColumnName(columnNumber int) string {
	columnName := ""
	for columnNumber > 0 {
		remainder := (columnNumber - 1) % 26
		columnName = string(rune('A'+remainder)) + columnName
		columnNumber = (columnNumber - 1) / 26
	}
	return columnName
}

// Export 导出excel
// lists		要导出的数据
// cols 		列
// sheet		文档
// title		标题
func Export(lists any, cols []Col, sheet string, title string) (file *excelize.File, err error) {
	e := ExcelInit()

	data := convert_util.ShallowStructsToMaps(lists)

	err = ExportExcel(sheet, title, data, cols, e)
	if err != nil {
		return
	}
	return e.F, err
}

// ExportExcel excel导出
func ExportExcel(sheet, title string, lists []map[string]interface{}, cols []Col, e *Excel) (err error) {
	index, _ := e.F.GetSheetIndex(sheet)
	if index < 0 { // 如果sheet名称不存在
		e.F.NewSheet(sheet)
	}
	// 构造表头
	endColName, startDataRow, err := buildTitle(e, sheet, title, cols)
	if err != nil {
		return
	}
	// 构造数据行
	err = buildDataRow(e, sheet, endColName, startDataRow, lists, cols)
	return
}

// 构造表头（endColName 最后一列的列名 startDataRow 数据行开始的行号）
func buildTitle(e *Excel, sheet, title string, cols []Col) (endColName string, startDataRow int, err error) {
	var titleRowData []interface{} // 列头行
	for i, colTitle := range cols {
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
		startDataRow = 3 // 如果有title，那么从第3行开始就是数据行，第1行是title，第2行是表头
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
		startDataRow = 2 // 如果没有title，那么从第2行开始就是数据行，第1行是表头
		e.F.SetRowHeight(sheet, 1, float64(30))
		e.F.SetCellStyle(sheet, "A1", endColName+"1", e.HeadStyle)
		if err = e.F.SetSheetRow(sheet, "A1", &titleRowData); err != nil {
			return
		}
	}
	return
}

// 构造数据行
func buildDataRow(e *Excel, sheet, endColName string, startDataRow int, lists []map[string]any, cols []Col) (err error) {
	//实时写入数据
	for i := 0; i < len(lists); i++ {
		startCol := fmt.Sprintf("A%d", startDataRow)
		endCol := fmt.Sprintf("%s%d", endColName, startDataRow)

		var rowData []interface{} // 数据列

		list := lists[i]
		for j := 0; j < len(cols); j++ {
			col := cols[j]
			replace := col.Replace

			val := list[col.Key]

			// 先编码
			if col.Encode != nil {
				val = col.Encode(val)
			} else if v, ok := val.(Encode); ok {
				val = v.String()
			}

			// 再替换
			for replaceKey, v := range replace {

				if replaceKey == fmt.Sprintf("%v", val) {
					val = fmt.Sprintf("%v", v)
					break
				}
			}

			rowData = append(rowData, val)
		}

		if startDataRow%2 == 0 {
			_ = e.F.SetCellStyle(sheet, startCol, endCol, e.ContentStyle2)
		} else {
			_ = e.F.SetCellStyle(sheet, startCol, endCol, e.ContentStyle1)
		}

		_ = e.F.SetRowHeight(sheet, startDataRow, float64(25)) // 默认行高25

		if err = e.F.SetSheetRow(sheet, startCol, &rowData); err != nil {
			return
		}
		startDataRow++
	}
	return

}
