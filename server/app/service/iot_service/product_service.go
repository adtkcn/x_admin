package iot_service

import (
	"encoding/json"
	"errors"
	"strings"
	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var ProductService = NewProductService()

func NewProductService() *productService {
	return &productService{db: core.GetDB()}
}

type productService struct {
	db *gorm.DB
}

// List 产品分页列表
func (s productService) List(page request.PageReq, listReq iot_schema.ProductListReq) (res response.PageResp, e error) {
	chain := s.db.Model(&iot_model.IotProduct{})
	if listReq.Name != "" {
		chain = chain.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.CategoryId != "" {
		chain = chain.Where("category_id = ?", listReq.CategoryId)
	}
	var count int64
	if e = response.CheckErr(chain.Count(&count).Error, "统计产品失败"); e != nil {
		return
	}
	var products []iot_model.IotProduct
	if e = response.CheckErr(chain.Limit(page.PageSize).Offset(page.PageSize*(page.PageNo-1)).
		Order("id desc").Find(&products).Error, "查询产品失败"); e != nil {
		return
	}
	var list []iot_schema.ProductResp
	convert_util.Copy(&list, products)
	s.fillCategoryName(list)
	s.fillRuleChainName(list)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    list,
	}, nil
}

// All 全部产品（供设备表单下拉选择）
func (s productService) All() (res []iot_schema.ProductResp, e error) {
	var products []iot_model.IotProduct
	if e = response.CheckErr(s.db.Where("is_stop = 0").Order("id desc").Find(&products).Error, "查询产品失败"); e != nil {
		return
	}
	res = []iot_schema.ProductResp{}
	convert_util.Copy(&res, products)
	s.fillCategoryName(res)
	s.fillRuleChainName(res)
	return
}

// fillCategoryName 批量回填分类名称
func (s productService) fillCategoryName(list []iot_schema.ProductResp) {
	if len(list) == 0 {
		return
	}
	nameMap := map[string]string{}
	var cats []iot_model.IotProductCategory
	s.db.Select("id, name").Find(&cats)
	for _, c := range cats {
		nameMap[c.ID] = c.Name
	}
	for i := range list {
		list[i].CategoryName = nameMap[list[i].CategoryId]
	}
}

// Detail 产品详情
func (s productService) Detail(id string) (res iot_schema.ProductResp, e error) {
	var product iot_model.IotProduct
	if e = response.CheckDBErr(s.db.Where("id = ?", id).First(&product).Error, "产品不存在!", "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, product)
	var cate iot_model.IotProductCategory
	if s.db.Where("id = ?", product.CategoryId).First(&cate).Error == nil {
		res.CategoryName = cate.Name
	}
	if product.RuleChainId != "" {
		var rc iot_model.IotRuleChain
		if s.db.Select("name").Where("id = ?", product.RuleChainId).First(&rc).Error == nil {
			res.RuleChainName = rc.Name
		}
	}
	return
}

// checkThingModel 物模型必填且必须是合法 JSON：处理层按物模型白名单裁剪属性，
// 未配置物模型的产品无法安全落地上报数据，故写入侧就要求必须配置（存非法 JSON 比静默降级更难排查）
func checkThingModel(thing string) error {
	if strings.TrimSpace(thing) == "" {
		return response.AssertArgumentError.SetMessage("物模型必须配置!")
	}
	if !json.Valid([]byte(thing)) {
		return response.AssertArgumentError.SetMessage("物模型必须是合法的 JSON!")
	}
	return nil
}

// Add 新增产品
func (s productService) Add(addReq iot_schema.ProductAddReq) (e error) {
	if e = checkThingModel(addReq.ThingModel); e != nil {
		return
	}
	// 分类已必填，保留空值分支：接口直调时不至于拿着野 category_id 建出无分类产品
	if addReq.CategoryId != "" {
		var cate iot_model.IotProductCategory
		r := s.db.Where("id = ?", addReq.CategoryId).Limit(1).Find(&cate)
		if e = response.CheckErr(r.Error, "分类查询失败"); e != nil {
			return
		}
		if r.RowsAffected == 0 {
			return response.AssertArgumentError.SetMessage("所属分类不存在!")
		}
	}
	var product iot_model.IotProduct
	convert_util.Copy(&product, addReq)
	e = response.CheckErr(s.db.Create(&product).Error, "添加失败")
	return
}

// Edit 编辑产品：节点类型从网关能力(GATEWAY/EDGE)变更为非网关时，需先解除其下设备的子设备绑定
func (s productService) Edit(editReq iot_schema.ProductEditReq) (e error) {
	if e = checkThingModel(editReq.ThingModel); e != nil {
		return
	}
	var product iot_model.IotProduct
	if e = response.CheckDBErr(s.db.Where("id = ?", editReq.ID).First(&product).Error, "产品不存在!", "编辑失败"); e != nil {
		return
	}
	if iot_schema.NodeTypeSupportsGateway(product.DeviceType) && !iot_schema.NodeTypeSupportsGateway(editReq.DeviceType) {
		var subCount int64
		if e = response.CheckErr(s.db.Model(&iot_model.IotDevice{}).
			Where("gateway_id IN (SELECT id FROM x_iot_device WHERE product_id = ? AND delete_time = 0)", editReq.ID).
			Where("delete_time = 0").Count(&subCount).Error, "统计子设备失败"); e != nil {
			return
		}
		if subCount > 0 {
			return response.AssertArgumentError.SetMessage("该产品下的设备已作为网关被子设备绑定，请先解除绑定再修改节点类型!")
		}
	}
	// 先留存旧值：Convert 会原地覆盖 product，同步设备时需要旧产品值判定“未自定义”
	oldTimeout := product.OfflineTimeout
	newTimeout := editReq.OfflineTimeout
	convert_util.Copy(&product, editReq)
	e = response.CheckErr(s.db.Model(&product).Select("*").Omit("id", "create_time").Updates(product).Error, "编辑失败")
	if e == nil {
		// 物模型/规则链引用可能已变更，失效处理层缓存，避免继续沿用旧配置
		DeviceService.InvalidateProductCache(editReq.ID)
		// 离线时长以设备记录为准，产品变更后及时同步未自定义的设备
		s.syncDevicesTimeout(editReq.ID, oldTimeout, newTimeout)
	}
	return
}

// syncDevicesTimeout 产品离线时长变更后，把其下未自定义的设备（值为 0 或等于旧产品值）
// 刷成新值，并失效这些设备的鉴权缓存，让上报侧立刻用上新窗口；
// 设过自定义值的设备（不等于旧产品值）保持不动。
func (s productService) syncDevicesTimeout(productId string, oldTimeout, newTimeout int) {
	if newTimeout <= 0 || oldTimeout == newTimeout {
		return
	}
	var devices []iot_model.IotDevice
	if e := s.db.Select("sn, access_token").
		Where("product_id = ? AND (offline_timeout = 0 OR offline_timeout = ?)", productId, oldTimeout).
		Find(&devices).Error; e != nil {
		core.Logger.Errorf("同步设备离线时长失败 productId=%s err=%v", productId, e)
		return
	}
	if len(devices) == 0 {
		return
	}
	if e := s.db.Model(&iot_model.IotDevice{}).
		Where("product_id = ? AND (offline_timeout = 0 OR offline_timeout = ?)", productId, oldTimeout).
		Update("offline_timeout", newTimeout).Error; e != nil {
		core.Logger.Errorf("同步设备离线时长失败 productId=%s err=%v", productId, e)
		return
	}
	for _, d := range devices {
		DeviceService.InvalidateAuth(d.Sn, d.AccessToken)
	}
}

// fillRuleChainName 批量回填产品引用的规则链名称
func (s productService) fillRuleChainName(list []iot_schema.ProductResp) {
	if len(list) == 0 {
		return
	}
	idSet := map[string]struct{}{}
	for i := range list {
		if list[i].RuleChainId != "" {
			idSet[list[i].RuleChainId] = struct{}{}
		}
	}
	if len(idSet) == 0 {
		return
	}
	ids := make([]string, 0, len(idSet))
	for id := range idSet {
		ids = append(ids, id)
	}
	nameMap := map[string]string{}
	var chains []iot_model.IotRuleChain
	s.db.Select("id, name").Where("id IN ?", ids).Find(&chains)
	for _, c := range chains {
		nameMap[c.ID] = c.Name
	}
	for i := range list {
		list[i].RuleChainName = nameMap[list[i].RuleChainId]
	}
}

// Del 删除产品：产品下存在设备时禁止删除
func (s productService) Del(id string) (e error) {
	if r := s.db.Where("product_id = ?", id).Limit(1).Find(&iot_model.IotDevice{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该产品下存在设备,请先移除!")
	}
	result := s.db.Where("id = ?", id).Delete(&iot_model.IotProduct{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("产品不存在")
	}
	DeviceService.InvalidateProductCache(id)
	return
}
