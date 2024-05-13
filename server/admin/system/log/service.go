package log

import (
	"fmt"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/system_model"

	"gorm.io/gorm"
)

type ISystemLogsServer interface {
	Operate(page request.PageReq, logReq SystemLogOperateReq) (res response.PageResp, e error)
	Login(page request.PageReq, logReq SystemLogLoginReq) (res response.PageResp, e error)
}

var Service = NewSystemLogsServer()

// NewSystemLogsServer 初始化
func NewSystemLogsServer() ISystemLogsServer {
	db := core.GetDB()
	return &systemLogsServer{db: db}
}

// systemLogsServer 系统日志服务实现类
type systemLogsServer struct {
	db *gorm.DB
}

// Operate 系统操作日志
func (logSrv systemLogsServer) Operate(page request.PageReq, logReq SystemLogOperateReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	logTbName := core.DBTableName(&system_model.SystemLogOperate{})
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})
	logModel := logSrv.db.Table(logTbName + " AS log").Joins(
		fmt.Sprintf("LEFT JOIN %s AS admin ON log.admin_id = admin.id", adminTbName)).Select(
		"log.*, admin.username, admin.nickname")
	// 条件
	if logReq.Title != "" {
		logModel = logModel.Where("title like ?", "%"+logReq.Title+"%")
	}
	if logReq.Username != "" {
		logModel = logModel.Where("username like ?", "%"+logReq.Username+"%")
	}
	if logReq.Ip != "" {
		logModel = logModel.Where("ip like ?", "%"+logReq.Ip+"%")
	}
	if logReq.Type != "" {
		logModel = logModel.Where("type = ?", logReq.Type)
	}
	if logReq.Status > 0 {
		logModel = logModel.Where("status = ?", logReq.Status)
	}
	if logReq.Url != "" {
		logModel = logModel.Where("url = ?", logReq.Url)
	}
	if logReq.StartTime != "" {
		logModel = logModel.Where("log.create_time >= ?", logReq.StartTime)
	}
	if logReq.EndTime != "" {
		logModel = logModel.Where("log.create_time <= ?", logReq.EndTime)
	}
	// 总数
	var count int64
	err := logModel.Count(&count).Error
	if e = response.CheckErr(err, "Operate Count err"); e != nil {
		return
	}
	// 数据
	var logResp []SystemLogOperateResp
	err = logModel.Limit(limit).Offset(offset).Order("id desc").Find(&logResp).Error
	if e = response.CheckErr(err, "Operate Find err"); e != nil {
		return
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    logResp,
	}, nil
}

// Login 系统登录日志
func (logSrv systemLogsServer) Login(page request.PageReq, logReq SystemLogLoginReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	logModel := logSrv.db.Model(&system_model.SystemLogLogin{})
	// 条件
	if logReq.Username != "" {
		logModel = logModel.Where("username like ?", "%"+logReq.Username+"%")
	}
	if logReq.Status > 0 {
		logModel = logModel.Where("status = ?", logReq.Status)
	}
	if logReq.StartTime != "" {
		logModel = logModel.Where("create_time >= now(?)", logReq.StartTime)
	}
	if logReq.EndTime != "" {
		logModel = logModel.Where("create_time <= now(?)", logReq.EndTime)
	}
	// 总数
	var count int64
	err := logModel.Count(&count).Error
	if e = response.CheckErr(err, "Login Count err"); e != nil {
		return
	}
	// 数据
	var logResp []SystemLogLoginResp
	err = logModel.Limit(limit).Offset(offset).Order("id desc").Find(&logResp).Error
	if e = response.CheckErr(err, "Login Find err"); e != nil {
		return
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    logResp,
	}, nil
}
