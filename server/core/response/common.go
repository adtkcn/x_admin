package response

//PageResp 分页返回值
type PageResp struct {
	Count    int64       `json:"count"`    // 总数
	PageNo   int         `json:"pageNo"`   // 每页数量
	PageSize int         `json:"pageSize"` // 每页Size
	Lists    interface{} `json:"lists"`    // 数据
}
