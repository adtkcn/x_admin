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

var ProductCategoryService = NewProductCategoryService()

func NewProductCategoryService() *productCategoryService {
	return &productCategoryService{db: core.GetDB()}
}

type productCategoryService struct {
	db *gorm.DB
}

// All 全部分类（不分页，前端 vxe-table 依据 id/pid 构建树）
func (s productCategoryService) All() (res []iot_schema.ProductCategoryResp, e error) {
	var list []iot_model.IotProductCategory
	if e = response.CheckErr(s.db.Order("sort desc, id asc").Find(&list).Error, "查询分类失败"); e != nil {
		return
	}
	res = []iot_schema.ProductCategoryResp{}
	convert_util.Copy(&res, list)
	return
}

// Detail 分类详情
func (s productCategoryService) Detail(id string) (res iot_schema.ProductCategoryResp, e error) {
	var cate iot_model.IotProductCategory
	if e = response.CheckDBErr(s.db.Where("id = ?", id).First(&cate).Error, "分类不存在!", "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, cate)
	return
}

// getDescendantIds 收集某分类的所有后代ID（防止上级设为自身或子孙形成环路）
func (s productCategoryService) getDescendantIds(id string) (descendantIds []string, e error) {
	var list []iot_model.IotProductCategory
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

// Add 新增分类
func (s productCategoryService) Add(addReq iot_schema.ProductCategoryAddReq) (e error) {
	if addReq.Pid != "" {
		var parent iot_model.IotProductCategory
		r := s.db.Where("id = ?", addReq.Pid).Limit(1).Find(&parent)
		if e = response.CheckErr(r.Error, "上级分类查询失败"); e != nil {
			return
		}
		if r.RowsAffected == 0 {
			return response.AssertArgumentError.SetMessage("上级分类不存在!")
		}
	}
	var cate iot_model.IotProductCategory
	convert_util.Copy(&cate, addReq)
	e = response.CheckErr(s.db.Create(&cate).Error, "添加失败")
	return
}

// Edit 编辑分类
func (s productCategoryService) Edit(editReq iot_schema.ProductCategoryEditReq) (e error) {
	var cate iot_model.IotProductCategory
	if e = response.CheckDBErr(s.db.Where("id = ?", editReq.ID).First(&cate).Error, "分类不存在!", "编辑失败"); e != nil {
		return
	}
	if editReq.ID == editReq.Pid {
		return response.AssertArgumentError.SetMessage("上级分类不能是自己!")
	}
	if editReq.Pid != "" {
		descendantIds, dErr := s.getDescendantIds(editReq.ID)
		if dErr != nil {
			return dErr
		}
		for _, dId := range descendantIds {
			if dId == editReq.Pid {
				return response.AssertArgumentError.SetMessage("上级分类不能是自己的子级!")
			}
		}
	}
	convert_util.Copy(&cate, editReq)
	e = response.CheckErr(s.db.Model(&cate).Select("*").Omit("id", "create_time").Updates(cate).Error, "编辑失败")
	return
}

// Del 删除分类：存在子分类或已被产品引用时禁止删除
func (s productCategoryService) Del(id string) (e error) {
	if r := s.db.Where("pid = ?", id).Limit(1).Find(&iot_model.IotProductCategory{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("请先删除子级分类!")
	}
	if r := s.db.Where("category_id = ?", id).Limit(1).Find(&iot_model.IotProduct{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("该分类已被产品使用,请先移除!")
	}
	result := s.db.Where("id = ?", id).Delete(&iot_model.IotProductCategory{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("分类不存在")
	}
	return
}

// Sort 同级拖拽排序，按传入 id 顺序重排 sort（单条 CASE SQL 批量更新）
func (s productCategoryService) Sort(ids []string) (e error) {
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
	if err := s.db.Model(&iot_model.IotProductCategory{}).
		Where("id in ?", ids).
		Update("sort", gorm.Expr(caseExpr, args...)).Error; err != nil {
		return response.CheckErr(err, "排序更新失败")
	}
	return
}
