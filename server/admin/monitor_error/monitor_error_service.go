package monitor_error

import (
	"errors"
	"strconv"
	"x_admin/admin/monitor_client"
	"x_admin/admin/monitor_error_list"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"gorm.io/gorm"
)

var MonitorErrorService = NewMonitorErrorService()
var cacheUtil = util.CacheUtil{
	Name: MonitorErrorService.Name,
}

// NewMonitorErrorService 初始化
func NewMonitorErrorService() *monitorErrorService {
	return &monitorErrorService{
		db:   core.GetDB(),
		Name: "monitorError",
	}
}

// monitorErrorService 监控-错误列服务实现类
type monitorErrorService struct {
	db   *gorm.DB
	Name string
}

// List 监控-错误列列表
func (service monitorErrorService) GetModel(listReq MonitorErrorListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.MonitorError{})
	if listReq.ProjectKey != nil {
		dbModel = dbModel.Where("project_key = ?", *listReq.ProjectKey)
	}
	if listReq.EventType != nil {
		dbModel = dbModel.Where("event_type = ?", *listReq.EventType)
	}
	if listReq.Path != nil {
		dbModel = dbModel.Where("path = ?", *listReq.Path)
	}
	if listReq.Message != nil {
		dbModel = dbModel.Where("message = ?", *listReq.Message)
	}
	if listReq.Stack != nil {
		dbModel = dbModel.Where("stack = ?", *listReq.Stack)
	}
	if listReq.Md5 != nil {
		dbModel = dbModel.Where("md5 = ?", *listReq.Md5)
	}
	if listReq.CreateTimeStart != nil {
		dbModel = dbModel.Where("create_time >= ?", *listReq.CreateTimeStart)
	}
	if listReq.CreateTimeEnd != nil {
		dbModel = dbModel.Where("create_time <= ?", *listReq.CreateTimeEnd)
	}

	return dbModel
}

// List 监控-错误列列表
func (service monitorErrorService) List(page request.PageReq, listReq MonitorErrorListReq) (res response.PageResp, e error) {
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
	var modelList []model.MonitorError
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorErrorResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 监控-错误列列表
func (service monitorErrorService) ListAll(listReq MonitorErrorListReq) (res []MonitorErrorResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.MonitorError

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail 监控-错误列详情
func (service monitorErrorService) Detail(Id int) (res MonitorErrorResp, e error) {
	var obj = model.MonitorError{}
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
	}

	convert_util.Copy(&res, obj)
	return
}

// DetailByMD5 监控-错误列详情
func (service monitorErrorService) DetailByMD5(md5 string) (res MonitorErrorResp, e error) {
	var obj = model.MonitorError{}
	err := cacheUtil.GetCache("md5:"+md5, &obj)
	if err != nil {
		err := service.db.Where("md5 = ?", md5).Order("id DESC").Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
		cacheUtil.SetCache("md5:"+md5, obj)
	}

	convert_util.Copy(&res, obj)
	return
}

// Add 监控-错误列新增
func (service monitorErrorService) Add(addReq MonitorErrorAddReq) (createId int, err error) {

	var obj model.MonitorError
	convert_util.StructToStruct(addReq, &obj)

	Md5 := util.ToolsUtil.MakeMd5(obj.ProjectKey + obj.EventType + obj.Message + obj.Path + obj.Stack)

	errorDetails, err := service.DetailByMD5(Md5)
	if err != nil {

		obj.Md5 = Md5

		err := service.db.Create(&obj).Error
		err = response.CheckMysqlErr(err)
		if err != nil {
			return 0, err
		}
		createId = obj.Id
		cacheUtil.SetCache(createId, obj)
		cacheUtil.SetCache("md5:"+Md5, obj)
	} else {
		createId = errorDetails.Id
	}
	client, err := monitor_client.MonitorClientService.DetailByClientId(addReq.ClientId)
	if err != nil {
		return 0, err
	}

	_, err = monitor_error_list.MonitorErrorListService.Add(monitor_error_list.MonitorErrorListAddReq{
		Eid: strconv.Itoa(createId),
		Cid: strconv.Itoa(client.Id),
		// ClientId:   addReq.ClientId,
		// ProjectKey: addReq.ProjectKey,
	})

	return createId, err
}

// Del 监控-错误列删除
func (service monitorErrorService) Del(Id int) (e error) {
	var obj model.MonitorError
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
	cacheUtil.RemoveCache("md5:" + obj.Md5)
	return
}

// DelBatch 用户协议-批量删除
func (service monitorErrorService) DelBatch(Ids []string) (e error) {
	var obj []model.MonitorError
	// 查询Ids对应的数据
	err := service.db.Where("id in (?)", Ids).Find(&obj).Error
	if err != nil {
		return err
	}
	if len(obj) == 0 {
		return errors.New("数据不存在")
	}
	err = service.db.Where("id in (?)", Ids).Delete(model.MonitorError{}).Error
	if err != nil {
		return err
	}
	// md5集合
	var md5s []string
	for _, v := range obj {
		md5s = append(md5s, "md5:"+v.Md5)
	}
	// 删除缓存
	cacheUtil.RemoveCache(Ids)
	cacheUtil.RemoveCache(md5s)
	return nil
}

// 获取Excel的列
func (service monitorErrorService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "项目key", Key: "ProjectKey", Width: 15},
		{Name: "事件类型", Key: "EventType", Width: 15},
		{Name: "URL地址", Key: "Path", Width: 15},
		{Name: "错误消息", Key: "Message", Width: 15},
		{Name: "错误堆栈", Key: "Stack", Width: 15},
		{Name: "md5", Key: "Md5", Width: 15},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "更新时间", Key: "ClientTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 监控-错误列导出
func (service monitorErrorService) ExportFile(listReq MonitorErrorListReq) (res []MonitorErrorResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.MonitorError
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorErrorResp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service monitorErrorService) ImportFile(importReq []MonitorErrorResp) (e error) {
	var importData []model.MonitorError
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
