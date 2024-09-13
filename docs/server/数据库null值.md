# 数据库null值

问题点：数据库在int,string等类型时，同时可能允许null值，但是go中int不允许null值，读取等操作可能报错。

注：数据库应尽量避免允许null值，避免不了时使用：https://github.com/guregu/null

相关解读：https://blog.csdn.net/qq_15437667/article/details/78780945

第三方库：https://pkg.go.dev/database/sql#NullInt64
```go
// 示例
type NullInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}
func (n *NullInt64) Scan(value any) error
func (n NullInt64) Value() (driver.Value, error)
```

第三方库（继承database/sql并补充json等）：https://github.com/guregu/null
```go
// 示例
type Int struct {
	sql.NullInt64
}
func (i *Int) UnmarshalJSON(data []byte) error {
	err := internal.UnmarshalIntJSON(data, &i.Int64, &i.Valid, 64, strconv.ParseInt)
	if err != nil {
		return err
	}
	i.Valid = i.Int64 != 0
	return nil
}
func (i Int) MarshalJSON() ([]byte, error) {
	n := i.Int64
	if !i.Valid {
		n = 0
	}
	return []byte(strconv.FormatInt(n, 10)), nil
}

```

### 自带了三个可为null类型：
1. 相对于guregu/null的优点
2. NullFloat、NullInt：支持前端传递null、字符串数字、数字
```go
core.NullFloat
core.NullInt
core.NullTime
```