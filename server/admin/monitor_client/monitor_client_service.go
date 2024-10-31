package monitor_client

import (
	"errors"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/util"
	"x_admin/util/excel2"

	"gorm.io/gorm"
)

var MonitorClientService = NewMonitorClientService()
var cacheUtil = util.CacheUtil{
	Name: MonitorClientService.Name,
}

// NewMonitorClientService 初始化
func NewMonitorClientService() *monitorClientService {
	return &monitorClientService{
		db:   core.GetDB(),
		Name: "monitorClient",
	}
}

// monitorClientService 监控-客户端信息服务实现类
type monitorClientService struct {
	db   *gorm.DB
	Name string
}

// List 监控-客户端信息列表
func (service monitorClientService) GetModel(listReq MonitorClientListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.MonitorClient{})
	if listReq.ProjectKey != nil {
		dbModel = dbModel.Where("project_key = ?", *listReq.ProjectKey)
	}
	if listReq.ClientId != nil {
		dbModel = dbModel.Where("client_id = ?", *listReq.ClientId)
	}
	if listReq.UserId != nil {
		dbModel = dbModel.Where("user_id = ?", *listReq.UserId)
	}
	if listReq.Os != nil {
		dbModel = dbModel.Where("os = ?", *listReq.Os)
	}
	if listReq.Browser != nil {
		dbModel = dbModel.Where("browser = ?", *listReq.Browser)
	}
	if listReq.Country != nil {
		dbModel = dbModel.Where("country = ?", *listReq.Country)
	}
	if listReq.Province != nil {
		dbModel = dbModel.Where("province = ?", *listReq.Province)
	}
	if listReq.City != nil {
		dbModel = dbModel.Where("city = ?", *listReq.City)
	}
	if listReq.Operator != nil {
		dbModel = dbModel.Where("operator = ?", *listReq.Operator)
	}
	if listReq.Ip != nil {
		dbModel = dbModel.Where("ip = ?", *listReq.Ip)
	}
	if listReq.Width != nil {
		dbModel = dbModel.Where("width = ?", *listReq.Width)
	}
	if listReq.Height != nil {
		dbModel = dbModel.Where("height = ?", *listReq.Height)
	}
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
func (service monitorClientService) List(page request.PageReq, listReq MonitorClientListReq) (res response.PageResp, e error) {
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
	result := []MonitorClientResp{}
	util.ConvertUtil.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 监控-客户端信息列表
func (service monitorClientService) ListAll(listReq MonitorClientListReq) (res []MonitorClientResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.MonitorClient

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	util.ConvertUtil.Copy(&res, modelList)
	return res, nil
}

func (service monitorClientService) DetailByClientId(ClientId string) (res MonitorClientResp, e error) {
	if ClientId == "" {
		return res, errors.New("ClientId不能为空")
	}
	var obj = model.MonitorClient{}
	err := cacheUtil.GetCache("ClientId:"+ClientId, &obj)
	if err != nil {
		err := service.db.Where("client_id = ?", ClientId).Order("id DESC").Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
		cacheUtil.SetCache(obj.Id, obj)
		cacheUtil.SetCache("ClientId:"+obj.ClientId, obj)
	}

	util.ConvertUtil.Copy(&res, obj)
	return
}

// Detail 监控-客户端信息详情
func (service monitorClientService) Detail(Id int) (res MonitorClientResp, e error) {
	var obj = model.MonitorClient{}
	err := cacheUtil.GetCache(Id, &obj)
	if err != nil {
		err := service.db.Where("id = ?", Id).Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
		cacheUtil.SetCache(obj.Id, obj)
		cacheUtil.SetCache("ClientId:"+obj.ClientId, obj)
	}

	util.ConvertUtil.Copy(&res, obj)
	return
}

// ErrorUser 监控-客户端信息详情
func (service monitorClientService) ErrorUsers(error_id int) (res []MonitorClientResp, e error) {
	var obj = []model.MonitorClient{}
	service.db.Raw("SELECT client.*,list.create_time AS create_time from x_monitor_error_list as list right join x_monitor_client as client on client.id = list.client_id where list.error_id = ? Order by list.id DESC LIMIT 0,20", error_id).Scan(&obj)

	util.ConvertUtil.Copy(&res, obj)
	return
}

// Add 监控-客户端信息新增
func (service monitorClientService) Add(addReq MonitorClientAddReq) (createId int, e error) {
	var obj model.MonitorClient
	util.ConvertUtil.StructToStruct(addReq, &obj)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return 0, e
	}
	cacheUtil.SetCache(obj.Id, obj)
	cacheUtil.SetCache("ClientId:"+obj.ClientId, obj)
	createId = obj.Id
	return
}

// // Edit 监控-客户端信息编辑
// func (service monitorClientService) Edit(editReq MonitorClientEditReq) (e error) {
// 	var obj model.MonitorClient
// 	err := service.db.Where("id = ?", editReq.Id).Limit(1).First(&obj).Error
// 	// 校验
// 	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
// 		return
// 	}
// 	if e = response.CheckErr(err, "查询失败"); e != nil {
// 		return
// 	}
// 	util.ConvertUtil.Copy(&obj, editReq)

// 	err = service.db.Model(&obj).Select("*").Updates(obj).Error
// 	if e = response.CheckErr(err, "编辑失败"); e != nil {
// 		return
// 	}
// 	cacheUtil.RemoveCache(obj.Id)
// 	service.Detail(obj.Id)
// 	return
// }

// Del 监控-客户端信息删除
func (service monitorClientService) Del(Id int) (e error) {
	var obj model.MonitorClient
	err := service.db.Where("id = ?", Id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询数据失败"); e != nil {
		return
	}
	// 删除
	err = service.db.Delete(&obj).Error
	e = response.CheckErr(err, "删除失败")
	cacheUtil.RemoveCache(obj.Id)
	cacheUtil.RemoveCache("ClientId:" + obj.ClientId)
	return
}

// DelBatch 用户协议-批量删除
func (service monitorClientService) DelBatch(Ids []string) (e error) {
	var obj []model.MonitorClient
	// 查询Ids对应的数据
	err := service.db.Where("id in (?)", Ids).Find(&obj).Error
	if err != nil {
		return err
	}
	if len(obj) == 0 {
		return errors.New("数据不存在")
	}
	err = service.db.Where("id in (?)", Ids).Delete(model.MonitorClient{}).Error
	if err != nil {
		return err
	}
	// md5集合
	var Clients []string
	for _, v := range obj {
		Clients = append(Clients, "ClientId:"+v.ClientId)
	}

	// 删除缓存
	cacheUtil.RemoveCache(Ids)
	cacheUtil.RemoveCache(Clients)
	return nil
}

// 获取Excel的列
func (service monitorClientService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "项目key", Key: "ProjectKey", Width: 15},
		{Name: "sdk生成的客户端id", Key: "ClientId", Width: 15},
		{Name: "用户id", Key: "UserId", Width: 15},
		{Name: "系统", Key: "Os", Width: 15},
		{Name: "浏览器", Key: "Browser", Width: 15},
		{Name: "城市", Key: "City", Width: 15},
		{Name: "屏幕", Key: "Width", Width: 15, Decode: core.DecodeInt},
		{Name: "屏幕高度", Key: "Height", Width: 15, Decode: core.DecodeInt},
		{Name: "ua记录", Key: "Ua", Width: 15},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 监控-客户端信息导出
func (service monitorClientService) ExportFile(listReq MonitorClientListReq) (res []MonitorClientResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.MonitorClient
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorClientResp{}
	util.ConvertUtil.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service monitorClientService) ImportFile(importReq []MonitorClientResp) (e error) {
	var importData []model.MonitorClient
	util.ConvertUtil.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
