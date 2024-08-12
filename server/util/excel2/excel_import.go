package excel2

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"x_admin/util"

	"github.com/xuri/excelize/v2"
)

func GetExcelData(file multipart.File, dst interface{}, cols []Col) (err error) {
	// 创建缓冲区
	buf := new(bytes.Buffer)

	// 将文件内容复制到缓冲区
	_, err = io.Copy(buf, file)
	if err != nil {
		fmt.Println(err)
		// c.String(http.StatusInternalServerError, "读取失败")
		err = errors.New("读取失败")
		return
	}
	// 创建Excel文件对象
	f, err := excelize.OpenReader(bytes.NewReader(buf.Bytes()))
	if err != nil {
		fmt.Println(err)
		// c.String(http.StatusInternalServerError, "Excel读取失败")
		err = errors.New("Excel读取失败")
		return
	}
	err = ImportExcel(f, dst, 2, cols)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return err
}

// ImportExcel 导入数据（单个sheet）
// 需要在传入的结构体中的字段加上tag：excel:"title:列头名称;"
// f 获取到的excel对象、dst 导入目标对象【传指针】
// headIndex 表头的索引，从0开始（用于获取表头名字）
// startRow 头行行数（从第startRow+1行开始扫）
func ImportExcel(f *excelize.File, dst interface{}, startRow int, cols []Col) (err error) {
	sheetName := f.GetSheetName(0) // 单个sheet时，默认读取第一个sheet
	err = importData(f, dst, sheetName, startRow, cols)
	return
}

// ImportBySheet 导入数据（读取指定sheet）sheetName Sheet名称
func ImportBySheet(f *excelize.File, dst interface{}, sheetName string, startRow int, cols []Col) (err error) {
	// 当需要读取多个sheet时，可以通过下面的方式，来调用 ImportBySheet 这个函数
	//sheetList := f.GetSheetList()
	//for _, sheetName := range sheetList {
	//	ImportBySheet(f,dst,sheetName,startRow)
	//}
	err = importData(f, dst, sheetName, startRow, cols)
	return
}

// 获取在数组中得下标
func GetIndex(items []string, item string) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

// 解析数据
func importData(f *excelize.File, dst interface{}, sheetName string, startRow int, cols []Col) (err error) {
	rows, err := f.GetRows(sheetName) // 获取所有行
	if err != nil {
		err = errors.New(sheetName + "工作表不存在")
		return
	}
	// heads := []string{}

	var data = []map[string]interface{}{}
	for i := 0; i < len(rows); i++ {

		if i < startRow { // 跳过头行
			continue
		}
		var rowMap = map[string]interface{}{}

		for j := 0; j < len(cols); j++ {
			col := cols[j]
			key := col.Key
			replace := col.Replace

			val := rows[i][j]
			// 将val替换为key
			for replaceKey, v := range replace {
				if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", val) {
					val = fmt.Sprintf("%v", replaceKey)
					break
				}
			}

			if col.Decode != nil {
				rowMap[key] = col.Decode(val)
			} else {
				rowMap[key] = val
			}
		}
		data = append(data, rowMap)

	}
	util.ConvertUtil.MapToStruct(data, dst)

	// 		fmt.Println("Type.Name:", field.Type.Name(), field.Type.Kind())
	// 		// 根据字段类型设置值
	// 		switch field.Type.Kind() {
	// 		case reflect.Int:
	// 			v, _ := strconv.ParseInt(cellValue, 10, 64)
	// 			newData.Field(i).SetInt(v)
	// 		case reflect.Int64:
	// 			v, _ := strconv.ParseInt(cellValue, 10, 64)
	// 			newData.Field(i).SetInt(v)
	// 		case reflect.Uint:
	// 			v, _ := strconv.ParseUint(cellValue, 10, 64)
	// 			newData.Field(i).SetUint(v)
	// 		case reflect.Uint8:
	// 			v, _ := strconv.ParseUint(cellValue, 10, 64)
	// 			newData.Field(i).SetUint(v)
	// 		case reflect.Uint16:
	// 			v, _ := strconv.ParseUint(cellValue, 10, 64)
	// 			newData.Field(i).SetUint(v)
	// 		case reflect.String:
	// 			newData.Field(i).SetString(cellValue)
	// 		}

	return
}
