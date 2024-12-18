package excel2

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

type Test struct {
	Id       string
	Name     string
	Email    string
	Com      string
	Dept     string
	RoleKey  string
	RoleName string
	Remark   string
}

var cols = []Col{
	{Name: "用户账号", Key: "Id", Width: 15},
	{Name: "用户姓名", Key: "Name", Width: 15},
	{Name: "用户邮箱", Key: "Email", Width: 15},
	{Name: "所属公司", Key: "Com", Width: 20},
	{Name: "所在部门", Key: "Dept", Width: 21},
	{Name: "角色代码", Key: "RoleKey", Width: 20},
	{Name: "角色名称", Key: "RoleName", Width: 25, Replace: map[string]interface{}{"1": "1_超级管理员", "2": "2_普通用户"}},
	{Name: "备注", Key: "Remark", Width: 30},
}

// 导出
func TestExport(t *testing.T) {
	var testList = []Test{
		{"fuhua", "符华", "fuhua@123.com", "太虚剑派", "开发部", "CJGLY", "1", "备注备注"},
		{"baiye", "白夜", "baiye@123.com", "天命科技有限公司", "执行部", "PTYG", "2", ""},
		{"chiling", "炽翎", "chiling@123.com", "太虚剑派", "行政部", "PTYG", "2", "备注备注备注备注"},
		{"yunmo", "云墨", "yunmo@123.com", "太虚剑派", "财务部", "CJGLY", "1", ""},
		{"yuelun", "月轮", "yuelun@123.com", "天命科技有限公司", "执行部", "CJGLY", "1", ""},
		{"xunyu", "迅羽", "xunyu@123.com哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈", "天命科技有限公司", "开发部", "PTYG", "2",
			"备注备注备注备注com哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试"},
	}

	f, err := Export(testList, cols, "Sheet1", "用户信息")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Path = "./测试.xlsx"
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return
	}
}

// 导入
func TestImports(t *testing.T) {
	f, err := excelize.OpenFile("./测试.xlsx")
	if err != nil {
		fmt.Println("文件打开失败")
	}
	importList := []Test{}
	err = ImportExcel(f, &importList, 2, cols)
	if err != nil {
		fmt.Println(err)
	}
	for _, t := range importList {
		fmt.Println(t)
	}
}

func TestGetExcelColumnName(t *testing.T) {
	for i := 0; i < 100; i++ {
		var col = GetExcelColumnName(i)
		fmt.Println("col:", col)
	}

}
