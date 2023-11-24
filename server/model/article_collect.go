package model

//ArticleCollect 文章收藏实体
type ArticleCollect struct {
    Id  int  `gorm:"primarykey;comment:'主键'"` // 主键
    UserId  int  `gorm:"comment:'用户ID'"` // 用户ID
    ArticleId  int  `gorm:"comment:'文章ID'"` // 文章ID
    IsDelete  int  `gorm:"comment:'是否删除'"` // 是否删除
    CreateTime  int64  `gorm:"comment:'创建时间'"` // 创建时间
    UpdateTime  int64  `gorm:"comment:'更新时间'"` // 更新时间
    DeleteTime  int64  `gorm:"comment:'是否删除'"` // 是否删除
}
