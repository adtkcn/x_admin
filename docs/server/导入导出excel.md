

col结构体
```go
type Col struct {
	Name    string
	Key     string
	Width   int                          // 宽度
	Replace map[string]any               //实现值的替换
	Encode  func(value any) any          //编码函数-导出，（先Encode后Replace）
	Decode  func(value any) (any, error) //解码函数-导入，（先Replace后Decode）
}
```

使用示例
```go
var cols = []excel2.Col{
		{Name: "标题", Key: "Title", Width: 15},
		{Name: "协议内容", Key: "Content", Width: 15},
		{Name: "排序", Key: "Sort", Width: 15},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
```
 

<<< @/../server/util/excel2/excel_test.go#envConfig