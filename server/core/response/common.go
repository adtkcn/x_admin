package response

// PageResp 分页返回值
type PageResp struct {
	Count    int64       `json:"count"`    // 总数
	PageNo   int         `json:"pageNo"`   // 每页数量
	PageSize int         `json:"pageSize"` // 每页Size
	Lists    interface{} `json:"lists"`    // 数据
}

// Copy 拷贝结构体
// func Copy(toValue interface{}, fromValue interface{}) interface{} {
// 	if err := copier.Copy(toValue, fromValue); err != nil {
// 		core.Logger.Errorf("Copy err: err=[%+v]", err)
// 		panic(SystemError)
// 	}
// 	return toValue
// }
