package user_protocol
import (
	"x_admin/core"

)

//UserProtocolListReq 用户协议列表参数
type UserProtocolListReq struct {
            Title *string `` // 标题
            Content *string `` // 协议内容
            Sort *float64 `` // 排序
            CreateTimeStart *string `` // 开始创建时间
            CreateTimeEnd *string `` // 结束创建时间
            UpdateTimeStart *string `` // 开始更新时间
            UpdateTimeEnd *string `` // 结束更新时间
}



//UserProtocolAddReq 用户协议新增参数
type UserProtocolAddReq struct {
    Title  *string  `` // 标题
    Content  *string  `` // 协议内容
    Sort  core.NullFloat  `` // 排序
}

//UserProtocolEditReq 用户协议编辑参数
type UserProtocolEditReq struct {
        Id int `` // 
        Title  *string  `` // 标题
        Content  *string  `` // 协议内容
        Sort  core.NullFloat  `` // 排序
}

//UserProtocolDetailReq 用户协议详情参数
type UserProtocolDetailReq struct {
    Id int `` // 
}

//UserProtocolDelReq 用户协议删除参数
type UserProtocolDelReq struct {
    Id int `` // 
}

//UserProtocolResp 用户协议返回信息
type UserProtocolResp struct {
        Id int `` // 
        Title string `` // 标题
        Content string `` // 协议内容
        Sort core.NullFloat `` // 排序
        CreateTime core.NullTime `` // 创建时间
        UpdateTime core.NullTime `` // 更新时间
}
