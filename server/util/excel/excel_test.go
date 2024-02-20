// 测试源码的文件名以 _test.go 结尾。
// 测试函数的函数名以 Test 开头。
// 函数签名为 func (t *testing.T)。

// https://blog.csdn.net/weixin_43165220/article/details/132939884

package excel

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

type Test struct {
	Id       string `excel:"name:用户账号;"`
	Name     string `excel:"name:用户姓名;index:1;"`
	Email    string `excel:"name:用户邮箱;width:25;"`
	Com      string `excel:"name:所属公司;"`
	Dept     string `excel:"name:所在部门;"`
	RoleKey  string `excel:"name:角色代码;"`
	RoleName string `excel:"name:角色名称;replace:1_超级管理员,2_普通用户;"`
	Remark   string `excel:"name:备注;width:40;"`
}

// 导出
func TestExport(t *testing.T) {
	var testList = []Test{
		{"fuhua", "符华", "fuhua@123.com", "太虚剑派", "开发部", "CJGLY", "1", "备注备注"},
		{"baiye", "白夜", "baiye@123.com", "天命科技有限公司", "执行部", "PTYG", "2", ""},
		{"chiling", "炽翎", "chiling@123.com", "太虚剑派", "行政部", "PTYG", "2", "备注备注备注备注"},
		{"yunmo", "云墨", "yunmo@123.com", "太虚剑派", "财务部", "CJGLY", "1", ""},
		{"yuelun", "月轮", "yuelun@123.com", "天命科技有限公司", "执行部", "CJGLY", "1", ""},
		{"xunyu", "迅羽",
			"xunyu@123.com哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这11111111111里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试",
			"天命科技有限公司", "开发部", "PTYG", "2",
			"备注备注备注备注com哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试"},
	}
	changeHead := map[string]string{"Id": "账号", "Name": "真实姓名"}
	//f, err := excel.NormalExport(testList, "Sheet1", "用户信息", "Id,Email,", true, true, changeHead)
	f, err := NormalDynamicExport(testList, "Sheet1", "用户信息", "", true, false, changeHead)
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
	err = ImportExcel(f, &importList, 1, 2)
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
