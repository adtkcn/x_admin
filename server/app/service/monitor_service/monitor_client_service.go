package monitor_service

import (
	"errors"
	"x_admin/app/model"
	"x_admin/app/schema/monitor_schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var MonitorClientService = NewMonitorClientService()

// NewMonitorClientService 初始化
func NewMonitorClientService() *monitorClientService {
	return &monitorClientService{
		db: core.GetDB(),
		CacheUtil: util.CacheUtil{
			Name: "monitorClient",
		},
	}
}

// monitorClientService 监控-客户端信息服务实现类
type monitorClientService struct {
	db        *gorm.DB
	CacheUtil util.CacheUtil
}

// List 监控-客户端信息列表
func (service monitorClientService) GetModel(listReq monitor_schema.MonitorClientListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.MonitorClient{})
	if listReq.ProjectKey != nil {
		dbModel = dbModel.Where("project_key = ?", *listReq.ProjectKey)
	}
	if listReq.ClientId != nil {
		dbModel = dbModel.Where("client_id = ?", *listReq.ClientId)
	}
	// if listReq.UserId != nil {
	// 	dbModel = dbModel.Where("user_id = ?", *listReq.UserId)
	// }
	if listReq.Os != nil {
		dbModel = dbModel.Where("os = ?", *listReq.Os)
	}
	if listReq.Browser != nil {
		dbModel = dbModel.Where("browser = ?", *listReq.Browser)
	}
	// if listReq.Country != nil {
	// 	dbModel = dbModel.Where("country = ?", *listReq.Country)
	// }
	// if listReq.Province != nil {
	// 	dbModel = dbModel.Where("province = ?", *listReq.Province)
	// }
	// if listReq.City != nil {
	// 	dbModel = dbModel.Where("city = ?", *listReq.City)
	// }
	// if listReq.Operator != nil {
	// 	dbModel = dbModel.Where("operator = ?", *listReq.Operator)
	// }
	// if listReq.Ip != nil {
	// 	dbModel = dbModel.Where("ip = ?", *listReq.Ip)
	// }

	if listReq.Ua != nil {
		dbModel = dbModel.Where("ua = ?", *listReq.Ua)
	}
	if listReq.CreateTimeStart != nil {
		dbModel = dbModel.Where("create_time >= ?", *listReq.CreateTimeStart)
	}
	if listReq.CreateTimeEnd != nil {
		dbModel = dbModel.Where("create_time <= ?", *listReq.CreateTimeEnd)
	}

	return dbModel
}

// List 监控-客户端信息列表
func (service monitorClientService) List(page request.PageReq, listReq monitor_schema.MonitorClientListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	dbModel := service.GetModel(listReq)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.MonitorClient
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []monitor_schema.MonitorClientResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 监控-客户端信息列表
func (service monitorClientService) ListAll(listReq monitor_schema.MonitorClientListReq) (res []monitor_schema.MonitorClientResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.MonitorClient

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

func (service monitorClientService) DetailByClientId(ClientId string) (res monitor_schema.MonitorClientResp, e error) {
	if ClientId == "" {
		return res, errors.New("ClientId不能为空")
	}
	var obj = model.MonitorClient{}
	err := service.CacheUtil.GetCache("ClientId:"+ClientId, &obj)
	if err != nil {
		err := service.db.Where("client_id = ?", ClientId).Order("id DESC").First(&obj).Error
		if e = response.CheckDBErr(err, "数据不存在!", "获取详情失败"); e != nil {
			return
		}
		service.CacheUtil.SetCache("ClientId:"+obj.ClientId, obj)
	}

	convert_util.Copy(&res, obj)
	return
}

// Detail 监控-客户端信息详情
func (service monitorClientService) Detail(Id string) (res monitor_schema.MonitorClientResp, e error) {
	var obj = model.MonitorClient{}
	err := service.CacheUtil.GetCache(Id, &obj)
	if err != nil {
		err := service.db.Where("id = ?", Id).First(&obj).Error
		if e = response.CheckDBErr(err, "数据不存在!", "获取详情失败"); e != nil {
			return
		}

		service.CacheUtil.SetCache("ClientId:"+obj.ClientId, obj)
	}
	convert_util.Copy(&res, obj)
	return
}

// ErrorUser 监控-客户端信息详情
func (service monitorClientService) ErrorUsers(error_id string) (res []monitor_schema.MonitorClientResp, e error) {
	var obj = []monitor_schema.MonitorClientResp{}
	service.db.Raw("SELECT client.os,client.browser,client.ua,list.* from x_monitor_error_list as list left join x_monitor_client as client on client.client_id = list.client_id where list.error_id = ? Order by list.id DESC LIMIT 0,20", error_id).Scan(&obj)

	convert_util.Copy(&res, obj)
	return
}

// Add 监控-客户端信息新增
func (service monitorClientService) Add(addReq monitor_schema.MonitorClientAddReq) (createId string, e error) {
	var obj model.MonitorClient
	convert_util.Copy(&obj, addReq)

	// 基于 Redis 去重短路：client_id 近期已上报过则直接命中缓存返回，
	// 避免高频重复上报每次都打 DB 写（缓存 TTL 1h，过期后回源 DB 重新校验）
	cacheKey := "ClientId:" + obj.ClientId
	var cached model.MonitorClient
	if err := service.CacheUtil.GetCache(cacheKey, &cached); err == nil && cached.Id != "" {
		return cached.Id, nil
	}

	// 缓存未命中：回源 DB，依赖 client_id 唯一索引保证幂等（冲突不更新任何字段）
	err := service.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "client_id"}, // 指定以 client_id 作为冲突判断字段（唯一索引）
		},
		DoNothing: true, // client_id 已存在时不更新任何字段
	}).Create(&obj).Error
	if e = response.CheckMysqlErr(err); e != nil {
		return "", e
	}

	// DoNothing 命中冲突时 obj.Id 为 BeforeCreate 生成的临时 uuid（未落库），
	// 需按 client_id 回查真实记录，避免返回不存在的假 id
	err = service.db.Where("client_id = ?", obj.ClientId).Order("id DESC").First(&obj).Error
	if e = response.CheckErr(err, "写入后回查失败"); e != nil {
		return "", e
	}

	// 写入/刷新缓存（TTL 由 CacheUtil.SetCache 控制，过期后下次回源 DB 校验）
	service.CacheUtil.SetCache(cacheKey, obj)
	createId = obj.Id
	return
}

// Del 监控-客户端信息删除
func (service monitorClientService) Del(Id string) (e error) {
	var obj model.MonitorClient
	err := service.db.Where("id = ?", Id).First(&obj).Error
	// 校验
	if e = response.CheckDBErr(err, "数据不存在!", "查询数据失败"); e != nil {
		return
	}
	// 删除
	err = service.db.Delete(&obj).Error
	e = response.CheckErr(err, "删除失败")
	service.CacheUtil.RemoveCache(obj.Id)
	service.CacheUtil.RemoveCache("ClientId:" + obj.ClientId)
	return
}

// DelBatch 批量删除
func (service monitorClientService) DelBatch(Ids []string) (e error) {
	// 删除前取出缓存键（client_id），用于删除后清理缓存
	var objs []model.MonitorClient
	service.db.Select("client_id").Where("id in (?)", Ids).Find(&objs)
	var clients []string
	for _, v := range objs {
		clients = append(clients, "ClientId:"+v.ClientId)
	}
	// 直接删除 + 影响行数判断数据不存在
	result := service.db.Where("id in (?)", Ids).Delete(&model.MonitorClient{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("数据不存在")
	}
	// 删除缓存
	service.CacheUtil.RemoveCache(Ids...)
	service.CacheUtil.RemoveCache(clients...)
	return nil
}

// 获取Excel的列
func (service monitorClientService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "项目key", JsonTag: "project_key", Width: 15},
		{Name: "sdk生成的客户端id", JsonTag: "client_id", Width: 15},
		{Name: "用户id", JsonTag: "user_id", Width: 15},
		{Name: "系统", JsonTag: "os", Width: 15},
		{Name: "浏览器", JsonTag: "browser", Width: 15},
		{Name: "城市", JsonTag: "city", Width: 15},
		{Name: "屏幕", JsonTag: "width", Width: 15, Decode: x_null.DecodeInt64},
		{Name: "屏幕高度", JsonTag: "height", Width: 15, Decode: x_null.DecodeInt64},
		{Name: "ua记录", JsonTag: "ua", Width: 15},
		{Name: "创建时间", JsonTag: "create_time", Width: 15, Decode: x_null.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 监控-客户端信息导出
func (service monitorClientService) ExportFile(listReq monitor_schema.MonitorClientListReq) (res []monitor_schema.MonitorClientResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.MonitorClient
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []monitor_schema.MonitorClientResp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service monitorClientService) ImportFile(importReq []monitor_schema.MonitorClientResp) (e error) {
	var importData []model.MonitorClient
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
