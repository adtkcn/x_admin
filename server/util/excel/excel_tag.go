package excel

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// 定义正则表达式模式
const (
	ExcelTagKey = "excel"
	Pattern     = "name:(.*?);|index:(.*?);|width:(.*?);|needMerge:(.*?);|replace:(.*?);"
)

type ExcelTag struct {
	Value     interface{}
	Name      string // 表头标题
	Index     int    // 列下标(从0开始)
	Width     int    // 列宽
	NeedMerge bool   // 是否需要合并
	Replace   string // 替换（需要替换的内容_替换后的内容。比如：1_未开始 ==> 表示1替换为未开始）
}

// 构造函数，返回一个带有默认值的 ExcelTag 实例
func NewExcelTag() ExcelTag {
	return ExcelTag{
		// 导入时会根据这个下标来拿单元格的值，当目标结构体字段没有设置index时，
		// 解析字段tag值时Index没读到就一直默认为0，拿单元格的值时，就始终拿的是第一列的值
		Index: -1, // 设置 Index 的默认值为 -1
	}
}

// 读取字段tag值
func (e *ExcelTag) GetTag(tag string) (err error) {
	// 编译正则表达式
	re := regexp.MustCompile(Pattern)
	matches := re.FindAllStringSubmatch(tag, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			for i, val := range match {
				if i != 0 && val != "" {
					e.setValue(match, val)
				}
			}
		}
	} else {
		err = errors.New("未匹配到值")
		return
	}
	return
}

// 设置ExcelTag 对应字段的值
func (e *ExcelTag) setValue(tag []string, value string) {
	if strings.Contains(tag[0], "name") {
		e.Name = value
	}
	if strings.Contains(tag[0], "index") {
		v, _ := strconv.ParseInt(value, 10, 8)
		e.Index = int(v)
	}
	if strings.Contains(tag[0], "width") {
		v, _ := strconv.ParseInt(value, 10, 8)
		e.Width = int(v)
	}
	if strings.Contains(tag[0], "needMerge") {
		v, _ := strconv.ParseBool(value)
		e.NeedMerge = v
	}
	if strings.Contains(tag[0], "replace") {
		e.Replace = value
	}
}
