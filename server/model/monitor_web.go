package model

//MonitorWeb 错误收集error实体
type MonitorWeb struct {
    
        Id int `gorm:"primarykey;comment:'uuid'" excel:"name:uuid;"` // uuid
    
        ProjectKey string `gorm:"comment:'项目key'" excel:"name:项目key;"` // 项目key
    
        ClientId string `gorm:"comment:'sdk生成的客户端id'" excel:"name:sdk生成的客户端id;"` // sdk生成的客户端id
    
        EventType string `gorm:"comment:'事件类型'" excel:"name:事件类型;"` // 事件类型
    
        Page string `gorm:"comment:'URL地址'" excel:"name:URL地址;"` // URL地址
    
        Message string `gorm:"comment:'错误消息'" excel:"name:错误消息;"` // 错误消息
    
        Stack string `gorm:"comment:'错误堆栈'" excel:"name:错误堆栈;"` // 错误堆栈
    
        ClientTime int `gorm:"comment:'客户端时间'" excel:"name:客户端时间;"` // 客户端时间
    
        CreateTime int64 `gorm:"autoCreateTime;comment:'创建时间'" excel:"name:创建时间;"` // 创建时间
    
}
