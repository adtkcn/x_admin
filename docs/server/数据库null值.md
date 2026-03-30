## 数据库null值

问题点：数据库在int,string等类型时，同时可能允许null值，但是go中int,string等类型不允许null值，数据类型不一致就会报错。

## 解决方法，自定义结构体实现以下接口
用于Gorm等数据库插件的扫描和写入数据库。
```
func (i *NullString) Scan(value any) error
func (i NullString) Value() (driver.Value, error)
```
可以实现JSON序列化和反序列化。
```
func (i NullString) MarshalJSON() ([]byte, error)
func (i *NullString) UnmarshalJSON(data []byte) error
```
用于gin框架的地址栏param参数绑定接口
```
func (i *NullString) UnmarshalParam(param string) error 
```
用于fmt等格式化输出（可选）
```
func (i NullString) String() string
```


### 扩展了四个可为null类型：
1. NullFloat、NullInt：支持前端传递null、字符串数字、数字
```go
x_null.Float64
x_null.Int64
x_null.String
x_null.Time
```

### 参考：
1. https://blog.csdn.net/qq_15437667/article/details/78780945
2. https://pkg.go.dev/database/sql#NullInt64
3. https://github.com/guregu/null