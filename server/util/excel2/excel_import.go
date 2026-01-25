package excel2

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"x_admin/util/convert_util"

	"github.com/xuri/excelize/v2"
)

/**
 * @Description:  获取导入数据
 * @param file 上传的文件
 * @param dst 导入目标对象【传指针】
 * @param cols 列信息
 * @return err
 */
func GetExcelData[T any](file multipart.File, dst T, cols []Col) (err error) {
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
		err = errors.New("Excel读取失败")
		return
	}
	err = ImportExcel(f, dst, 2, cols)

	return err
}

// ImportExcel 导入数据（单个sheet）
// 需要在传入的结构体中的字段加上tag：excel:"title:列头名称;"
// f 获取到的excel对象、dst 导入目标对象【传指针】
// headIndex 表头的索引，从0开始（用于获取表头名字）
// startRow 头行行数（从第startRow+1行开始扫）
func ImportExcel[T any](f *excelize.File, dst T, startRow int, cols []Col) (err error) {
	sheetName := f.GetSheetName(0) // 单个sheet时，默认读取第一个sheet
	err = importData(f, dst, sheetName, startRow, cols)
	return
}

// 解析数据
func importData[T any](f *excelize.File, dst T, sheetName string, startRow int, cols []Col) (err error) {
	rows, err := f.GetRows(sheetName) // 获取所有行
	if err != nil {
		err = errors.New(sheetName + "工作表不存在")
		return
	}

	var data = []map[string]any{}
	for i := 0; i < len(rows); i++ {

		if i < startRow { // 跳过头行
			continue
		}
		var rowMap = map[string]any{}

		for j := 0; j < len(cols); j++ {
			col := cols[j]
			key := col.Key
			replace := col.Replace

			colVal := rows[i][j]
			// 先替换，将val替换为key
			for replaceKey, replaceVal := range replace {
				if fmt.Sprintf("%v", replaceVal) == colVal {
					colVal = replaceKey
					break
				}
			}
			// 再解码
			if col.Decode != nil {
				v, e := col.Decode(colVal)
				if e == nil {
					rowMap[key] = v
				}
			} else {
				rowMap[key] = colVal
			}
		}
		data = append(data, rowMap)

	}
	convert_util.MapToStruct(data, dst)
	return
}
