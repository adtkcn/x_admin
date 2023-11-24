package article_collect

import "x_admin/core"

//ArticleCollectListReq 文章收藏列表参数
type ArticleCollectListReq struct {
	UserId    int `form:"userId"`    // 用户ID
	ArticleId int `form:"articleId"` // 文章ID
}

//ArticleCollectDetailReq 文章收藏详情参数
type ArticleCollectDetailReq struct {
	Id int `form:"id"` // 主键
}

//ArticleCollectAddReq 文章收藏新增参数
type ArticleCollectAddReq struct {
	UserId    int `form:"userId"`    // 用户ID
	ArticleId int `form:"articleId"` // 文章ID
}

//ArticleCollectEditReq 文章收藏新增参数
type ArticleCollectEditReq struct {
	Id        int `form:"id"`        // 主键
	UserId    int `form:"userId"`    // 用户ID
	ArticleId int `form:"articleId"` // 文章ID
}

//ArticleCollectDelReq 文章收藏新增参数
type ArticleCollectDelReq struct {
	Id int `form:"id"` // 主键
}

//ArticleCollectResp 文章收藏返回信息
type ArticleCollectResp struct {
	Id         int         `json:"id" structs:"id"`                 // 主键
	UserId     int         `json:"userId" structs:"userId"`         // 用户ID
	ArticleId  int         `json:"articleId" structs:"articleId"`   // 文章ID
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
