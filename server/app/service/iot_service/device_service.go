package iot_service

import (
	"errors"
	"strings"
	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var DeviceService = NewDeviceService()

func NewDeviceService() *deviceService {
	return &deviceService{db: core.GetDB()}
}

type deviceService struct {
	db *gorm.DB
}

// List 设备分页列表
func (s deviceService) List(page request.PageReq, listReq iot_schema.DeviceListReq) (res response.PageResp, e error) {
	chain := s.db.Model(&iot_model.IotDevice{})
	if listReq.Name != "" {
		chain = chain.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Sn != "" {
		chain = chain.Where("sn like ?", "%"+listReq.Sn+"%")
	}
	if listReq.ProductId != "" {
		chain = chain.Where("product_id = ?", listReq.ProductId)
	}
	if listReq.SpaceId != "" {
		chain = chain.Where("space_id = ?", listReq.SpaceId)
	}
	if listReq.Status >= 0 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	var count int64
	if e = response.CheckErr(chain.Count(&count).Error, "统计设备失败"); e != nil {
		return
	}
	var devices []iot_model.IotDevice
	if e = response.CheckErr(chain.Limit(page.PageSize).Offset(page.PageSize*(page.PageNo-1)).
		Order("id desc").Find(&devices).Error, "查询设备失败"); e != nil {
		return
	}
	var list []iot_schema.DeviceResp
	convert_util.Copy(&list, devices)
	s.fillNames(list)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    list,
	}, nil
}

// fillNames 批量回填产品名称、产品节点类型、区域名称与网关设备名称
//
// 一律按“列表里实际出现的 id”做 IN 回查：先前是无条件全表 Find 产品与区域，
// 每次分页/详情请求都要把两张表读进内存，成本随元数据总量线性增长。
func (s deviceService) fillNames(list []iot_schema.DeviceResp) {
	if len(list) == 0 {
		return
	}
	// 一次遍历收集三类关联 id（去重），再按 id 集合回查
	productSet := map[string]struct{}{}
	spaceSet := map[string]struct{}{}
	gatewaySet := map[string]struct{}{}
	for _, d := range list {
		if d.ProductId != "" {
			productSet[d.ProductId] = struct{}{}
		}
		if d.SpaceId != "" {
			spaceSet[d.SpaceId] = struct{}{}
		}
		if d.GatewayId != "" {
			gatewaySet[d.GatewayId] = struct{}{}
		}
	}

	productMap := map[string]iot_model.IotProduct{}
	if ids := setToSlice(productSet); len(ids) > 0 {
		var products []iot_model.IotProduct
		s.db.Select("id, name, device_type").Where("id IN ?", ids).Find(&products)
		for _, p := range products {
			productMap[p.ID] = p
		}
	}
	spaceMap := map[string]string{}
	if ids := setToSlice(spaceSet); len(ids) > 0 {
		var spaces []iot_model.IotDeviceSpace
		s.db.Select("id, name").Where("id IN ?", ids).Find(&spaces)
		for _, sp := range spaces {
			spaceMap[sp.ID] = sp.Name
		}
	}
	gatewayMap := map[string]string{}
	if ids := setToSlice(gatewaySet); len(ids) > 0 {
		var gateways []iot_model.IotDevice
		s.db.Select("id, name").Where("id IN ?", ids).Find(&gateways)
		for _, g := range gateways {
			gatewayMap[g.ID] = g.Name
		}
	}
	for i := range list {
		product := productMap[list[i].ProductId]
		list[i].ProductName = product.Name
		list[i].ProductNodeType = product.DeviceType
		list[i].SpaceName = spaceMap[list[i].SpaceId]
		list[i].GatewayName = gatewayMap[list[i].GatewayId]
	}
}

// setToSlice 集合转切片（空集合返回空切片，供调用方直接跳过 IN 查询）
func setToSlice(set map[string]struct{}) []string {
	ids := make([]string, 0, len(set))
	for id := range set {
		ids = append(ids, id)
	}
	return ids
}

// Detail 设备详情
func (s deviceService) Detail(id string) (res iot_schema.DeviceResp, e error) {
	var device iot_model.IotDevice
	if e = response.CheckDBErr(s.db.Where("id = ?", id).First(&device).Error, "设备不存在!", "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, device)
	res_list := []iot_schema.DeviceResp{res}
	s.fillNames(res_list)
	res = res_list[0]
	return
}

// checkGateway 校验网关绑定：网关设备必须存在，其所属产品节点类型须具备网关能力(GATEWAY/EDGE)，且不得形成环形绑定
func (s deviceService) checkGateway(id string, gatewayId string) error {
	if gatewayId == "" {
		return nil
	}
	if gatewayId == id {
		return response.AssertArgumentError.SetMessage("网关设备不能是自己!")
	}
	var gateway iot_model.IotDevice
	if r := s.db.Where("id = ?", gatewayId).Limit(1).Find(&gateway); r.RowsAffected == 0 {
		return response.AssertArgumentError.SetMessage("所选网关设备不存在!")
	}
	var gatewayProduct iot_model.IotProduct
	s.db.Select("id, device_type").Where("id = ?", gateway.ProductId).First(&gatewayProduct)
	if !iot_schema.NodeTypeSupportsGateway(gatewayProduct.DeviceType) {
		return response.AssertArgumentError.SetMessage("所选设备所属产品不是网关类型，不能作为网关!")
	}
	// 环形绑定检测：从网关沿父链向上回溯，遇到当前设备即为环
	cur := gatewayId
	for depth := 0; cur != "" && depth < 100; depth++ {
		if cur == id {
			return response.AssertArgumentError.SetMessage("不能绑定到自身的下级设备，会形成环形层级!")
		}
		var parent iot_model.IotDevice
		if r := s.db.Select("gateway_id").Where("id = ?", cur).Limit(1).Find(&parent); r.RowsAffected == 0 {
			break
		}
		cur = parent.GatewayId
	}
	return nil
}

// fillOfflineTimeout 离线判定时长以设备记录为准（冗余+自定义）：未指定（0）时
// 把所属产品的值冗余到设备上，上报/扫描两条路径都只读设备列
func fillOfflineTimeout(timeout int, productTimeout int) int {
	if timeout > 0 {
		return timeout
	}
	return productTimeout
}

// Add 新增设备：校验产品/区域/网关绑定、sn 唯一，自动生成接入凭证
func (s deviceService) Add(addReq iot_schema.DeviceAddReq) (e error) {
	var product iot_model.IotProduct
	if r := s.db.Where("id = ?", addReq.ProductId).Limit(1).Find(&product); r.RowsAffected == 0 {
		return response.AssertArgumentError.SetMessage("所属产品不存在!")
	}
	if addReq.SpaceId != "" {
		var space iot_model.IotDeviceSpace
		if r := s.db.Where("id = ?", addReq.SpaceId).Limit(1).Find(&space); r.RowsAffected == 0 {
			return response.AssertArgumentError.SetMessage("所属区域不存在!")
		}
	}
	if r := s.db.Where("sn = ?", addReq.Sn).Limit(1).Find(&iot_model.IotDevice{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("设备编号已存在!")
	}
	if e = s.checkGateway("", addReq.GatewayId); e != nil {
		return
	}
	var device iot_model.IotDevice
	convert_util.Copy(&device, addReq)
	device.OfflineTimeout = fillOfflineTimeout(addReq.OfflineTimeout, product.OfflineTimeout)
	device.AccessToken = strings.ReplaceAll(util.ToolsUtil.MakeUuidV7(), "-", "")
	e = response.CheckErr(s.db.Create(&device).Error, "添加失败")
	return
}

// Edit 编辑设备（接入凭证不随编辑变更）
func (s deviceService) Edit(editReq iot_schema.DeviceEditReq) (e error) {
	var device iot_model.IotDevice
	if e = response.CheckDBErr(s.db.Where("id = ?", editReq.ID).First(&device).Error, "设备不存在!", "编辑失败"); e != nil {
		return
	}
	if r := s.db.Where("sn = ? AND id != ?", editReq.Sn, editReq.ID).Limit(1).Find(&iot_model.IotDevice{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("设备编号已存在!")
	}
	if e = s.checkGateway(editReq.ID, editReq.GatewayId); e != nil {
		return
	}
	// sn 可被改写：先按旧 sn 失效接入鉴权缓存，否则旧编号仍能上报
	s.InvalidateAuth(device.Sn, device.AccessToken)
	convert_util.Copy(&device, editReq)
	if device.OfflineTimeout <= 0 {
		// 编辑未指定时重新冗余新产品值（产品可能已换或产品时长已改）
		var product iot_model.IotProduct
		s.db.Select("offline_timeout").Where("id = ?", editReq.ProductId).First(&product)
		device.OfflineTimeout = fillOfflineTimeout(editReq.OfflineTimeout, product.OfflineTimeout)
	}
	// 在线状态与活跃时间由设备接入层维护，人工编辑不覆盖这几列
	e = response.CheckErr(s.db.Model(&device).Select("*").
		Omit("id", "create_time", "access_token", "status", "last_online_time", "last_offline_time").
		Updates(device).Error, "编辑失败")
	return
}

// UpdateOnlineStatus 更新设备在线状态（供设备接入层根据心跳/连接事件调用，非管理端人工设置）：
// status=1 上线并刷新最后上线时间；status=0 离线并刷新最后离线时间。
// 实际写入逻辑收敛到 MarkStatus，避免两处各维护一套“状态+时间戳”的更新语义。
func (s deviceService) UpdateOnlineStatus(id string, status uint8) (e error) {
	_, e = s.MarkStatus(id, status)
	return
}

// Del 删除设备：存在下级子设备（作为网关）时禁止删除
func (s deviceService) Del(id string) (e error) {
	if r := s.db.Where("gateway_id = ?", id).Limit(1).Find(&iot_model.IotDevice{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该设备下存在子设备,请先移除!")
	}
	var device iot_model.IotDevice
	s.db.Select("id, sn, access_token").Where("id = ?", id).First(&device)
	result := s.db.Where("id = ?", id).Delete(&iot_model.IotDevice{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("设备不存在")
	}
	// 设备已删除，立即失效鉴权缓存，避免 TTL 内旧凭证继续接入
	s.InvalidateAuth(device.Sn, device.AccessToken)
	return
}
