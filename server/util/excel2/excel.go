package excel2

import (
	"net/http"

	"github.com/xuri/excelize/v2"
)

type Excel struct {
	F             *excelize.File // excel 对象
	TitleStyle    int            // 表头样式
	HeadStyle     int            // 表头样式
	ContentStyle1 int            // 主体样式1，无背景色
	ContentStyle2 int            // 主体样式2，有背景色
}
type Col struct {
	Name    string
	Key     string
	Width   int                          // 宽度
	Replace map[string]any               //实现值的替换
	Encode  func(value any) any          //编码函数-导出，（先Encode后Replace）
	Decode  func(value any) (any, error) //解码函数-导入，（先Replace后Decode）
}

// 定义编码接口
type Encode interface {
	String() string
}

// // 定义解码接口
// type Decode interface {
// 	MarshalJSON() ([]byte, error)
// }

// 下载
func DownLoadExcel(fileName string, res http.ResponseWriter, file *excelize.File) {
	// 设置响应头
	// res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Content-Disposition", "attachment; filename="+fileName+".xlsx")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	err := file.Write(res) // 写入Excel文件内容到响应体
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 初始化
func ExcelInit() (e *Excel) {
	e = &Excel{}
	// excel构建
	e.F = excelize.NewFile()
	// 初始化样式
	e.getTitleRowStyle()
	e.getHeadRowStyle()
	e.getDataRowStyle()
	return e
}

// 获取边框样式
func getBorder() []excelize.Border {
	return []excelize.Border{ // 边框
		{Type: "top", Color: "000000", Style: 1},
		{Type: "bottom", Color: "000000", Style: 1},
		{Type: "left", Color: "000000", Style: 1},
		{Type: "right", Color: "000000", Style: 1},
	}
}

// 标题样式
func (e *Excel) getTitleRowStyle() {
	e.TitleStyle, _ = e.F.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{ // 对齐方式
			Horizontal: "center", // 水平对齐居中
			Vertical:   "center", // 垂直对齐居中
		},
		Fill: excelize.Fill{ // 背景颜色
			Type:    "pattern",
			Color:   []string{"#fff2cc"},
			Pattern: 1,
		},
		Font: &excelize.Font{ // 字体
			Bold: true,
			Size: 16,
		},
		Border: getBorder(),
	})
}

// 列头行样式
func (e *Excel) getHeadRowStyle() {
	e.HeadStyle, _ = e.F.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{ // 对齐方式
			Horizontal: "center", // 水平对齐居中
			Vertical:   "center", // 垂直对齐居中
			WrapText:   true,     // 自动换行
		},
		Fill: excelize.Fill{ // 背景颜色
			Type:    "pattern",
			Color:   []string{"#FDE9D8"},
			Pattern: 1,
		},
		Font: &excelize.Font{ // 字体
			Bold: true,
			Size: 14,
		},
		Border: getBorder(),
	})
}

// 数据行样式
func (e *Excel) getDataRowStyle() {
	style := excelize.Style{}
	style.Border = getBorder()
	style.Alignment = &excelize.Alignment{
		Horizontal: "center", // 水平对齐居中
		Vertical:   "center", // 垂直对齐居中
		WrapText:   true,     // 自动换行
	}
	style.Font = &excelize.Font{
		Size: 12,
	}
	e.ContentStyle1, _ = e.F.NewStyle(&style)
	style.Fill = excelize.Fill{ // 背景颜色
		Type:    "pattern",
		Color:   []string{"#cce7f5"},
		Pattern: 1,
	}
	e.ContentStyle2, _ = e.F.NewStyle(&style)
}
