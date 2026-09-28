package iot_service

import (
	"errors"
	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var DeviceSpaceService = NewDeviceSpaceService()

func NewDeviceSpaceService() *deviceSpaceService {
	return &deviceSpaceService{db: core.GetDB()}
}

type deviceSpaceService struct {
	db *gorm.DB
}

// All 全部区域（不分页，前端依据 id/pid 构建树）
func (s deviceSpaceService) All() (res []iot_schema.DeviceSpaceResp, e error) {
	var list []iot_model.IotDeviceSpace
	if e = response.CheckErr(s.db.Order("sort desc, id asc").Find(&list).Error, "查询区域失败"); e != nil {
		return
	}
	res = []iot_schema.DeviceSpaceResp{}
	convert_util.Copy(&res, list)
	return
}

// Detail 区域详情
func (s deviceSpaceService) Detail(id string) (res iot_schema.DeviceSpaceResp, e error) {
	var space iot_model.IotDeviceSpace
	if e = response.CheckDBErr(s.db.Where("id = ?", id).First(&space).Error, "区域不存在!", "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, space)
	return
}

// getDescendantIds 收集某区域的所有后代ID（防止上级设为自身或子孙形成环路）
func (s deviceSpaceService) getDescendantIds(id string) (descendantIds []string, e error) {
	var list []iot_model.IotDeviceSpace
	if e = s.db.Select("id, pid").Find(&list).Error; e != nil {
		return
	}
	childrenMap := make(map[string][]string)
	for _, d := range list {
		childrenMap[d.Pid] = append(childrenMap[d.Pid], d.ID)
	}
	queue := []string{id}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, childId := range childrenMap[current] {
			descendantIds = append(descendantIds, childId)
			queue = append(queue, childId)
		}
	}
	return
}

// Add 新增区域
func (s deviceSpaceService) Add(addReq iot_schema.DeviceSpaceAddReq) (e error) {
	if addReq.Pid != "" {
		var parent iot_model.IotDeviceSpace
		r := s.db.Where("id = ?", addReq.Pid).Limit(1).Find(&parent)
		if e = response.CheckErr(r.Error, "上级区域查询失败"); e != nil {
			return
		}
		if r.RowsAffected == 0 {
			return response.AssertArgumentError.SetMessage("上级区域不存在!")
		}
	}
	var space iot_model.IotDeviceSpace
	convert_util.Copy(&space, addReq)
	e = response.CheckErr(s.db.Create(&space).Error, "添加失败")
	return
}

// Edit 编辑区域
func (s deviceSpaceService) Edit(editReq iot_schema.DeviceSpaceEditReq) (e error) {
	var space iot_model.IotDeviceSpace
	if e = response.CheckDBErr(s.db.Where("id = ?", editReq.ID).First(&space).Error, "区域不存在!", "编辑失败"); e != nil {
		return
	}
	if editReq.ID == editReq.Pid {
		return response.AssertArgumentError.SetMessage("上级区域不能是自己!")
	}
	if editReq.Pid != "" {
		descendantIds, dErr := s.getDescendantIds(editReq.ID)
		if dErr != nil {
			return dErr
		}
		for _, dId := range descendantIds {
			if dId == editReq.Pid {
				return response.AssertArgumentError.SetMessage("上级区域不能是自己的子级!")
			}
		}
	}
	convert_util.Copy(&space, editReq)
	e = response.CheckErr(s.db.Model(&space).Select("*").Omit("id", "create_time").Updates(space).Error, "编辑失败")
	return
}

// Del 删除区域：存在子区域或区域下已有设备时禁止删除
func (s deviceSpaceService) Del(id string) (e error) {
	if r := s.db.Where("pid = ?", id).Limit(1).Find(&iot_model.IotDeviceSpace{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("请先删除子级区域!")
	}
	if r := s.db.Where("space_id = ?", id).Limit(1).Find(&iot_model.IotDevice{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该区域下存在设备,请先移除!")
	}
	result := s.db.Where("id = ?", id).Delete(&iot_model.IotDeviceSpace{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("区域不存在")
	}
	return
}

// Sort 同级拖拽排序
func (s deviceSpaceService) Sort(ids []string) (e error) {
	if len(ids) == 0 {
		return
	}
	caseExpr := "CASE id"
	args := make([]interface{}, 0, len(ids)*2)
	for i, id := range ids {
		caseExpr += " WHEN ? THEN ?"
		args = append(args, id, len(ids)-i)
	}
	caseExpr += " ELSE sort END"
	if err := s.db.Model(&iot_model.IotDeviceSpace{}).
		Where("id in ?", ids).
		Update("sort", gorm.Expr(caseExpr, args...)).Error; err != nil {
		return response.CheckErr(err, "排序更新失败")
	}
	return
}
