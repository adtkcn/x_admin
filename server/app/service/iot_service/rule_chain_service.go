package iot_service

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

// RuleChainService 可视化规则链：既是管理端 CRUD 出入口，也向处理层提供执行树缓存读取，
// 替代原产品级三段 goja 脚本（解码/过滤/告警）。
var RuleChainService = NewRuleChainService()

func NewRuleChainService() *ruleChainService {
	return &ruleChainService{db: core.GetDB()}
}

type ruleChainService struct {
	db *gorm.DB
}

// chainCacheKey 规则链缓存键：按链 ID缓存执行树原文，多产品复用同一条链时共享一份。
func chainCacheKey(chainId string) string {
	return iotCacheKey("rulechain:" + chainId)
}

// List 规则链分页列表
func (s ruleChainService) List(page request.PageReq, listReq iot_schema.RuleChainListReq) (res response.PageResp, e error) {
	chain := s.db.Model(&iot_model.IotRuleChain{})
	if v := listReq.Name.ValueOrZero(); v != "" {
		chain = chain.Where("name like ?", "%"+v+"%")
	}
	if v := listReq.IsStop.ValueOrZero(); v != "" {
		chain = chain.Where("is_stop = ?", v)
	}
	var count int64
	if e = response.CheckErr(chain.Count(&count).Error, "统计规则链失败"); e != nil {
		return
	}
	var rows []iot_model.IotRuleChain
	if e = response.CheckErr(chain.Limit(page.PageSize).Offset(page.PageSize*(page.PageNo-1)).
		Order("id desc").Find(&rows).Error, "查询规则链失败"); e != nil {
		return
	}
	var list []iot_schema.RuleChainResp
	convert_util.Copy(&list, rows)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    list,
	}, nil
}

// All 全部启用规则链（供产品表单下拉选择）
func (s ruleChainService) All() (res []iot_schema.RuleChainResp, e error) {
	var rows []iot_model.IotRuleChain
	if e = response.CheckErr(s.db.Where("is_stop = 0").Order("id desc").Find(&rows).Error, "查询规则链失败"); e != nil {
		return
	}
	res = []iot_schema.RuleChainResp{}
	convert_util.Copy(&res, rows)
	return
}

// Detail 规则链详情
func (s ruleChainService) Detail(id string) (res iot_schema.RuleChainResp, e error) {
	var row iot_model.IotRuleChain
	if e = response.CheckDBErr(s.db.Where("id = ?", id).First(&row).Error, "规则链不存在!", "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, row)
	return
}

// checkTree 执行树必须能反序列化为 RuleTreeNode 且以「输入」为根。
// graph_data 仅供绘图、后端不执行；规模上限/防环由前端 buildRcTree 保证，后端保存与运行期均不再重复设限。
func checkTree(treeRaw string) error {
	var root iot_schema.RuleTreeNode
	if err := json.Unmarshal([]byte(treeRaw), &root); err != nil {
		return response.AssertArgumentError.SetMessage("规则链执行树结构不合法!")
	}
	if root.Type != iot_schema.RcNodeTrigger {
		return response.AssertArgumentError.SetMessage("规则链执行树必须以「输入」节点为根!")
	}
	return nil
}

// Add 新增规则链
func (s ruleChainService) Add(addReq iot_schema.RuleChainAddReq) (e error) {
	if e = checkTree(addReq.TreeData); e != nil {
		return
	}
	var row iot_model.IotRuleChain
	convert_util.Copy(&row, addReq)
	e = response.CheckErr(s.db.Create(&row).Error, "添加失败")
	return
}

// Edit 编辑规则链：变更后失效处理层执行树缓存，避免沿用旧拓扑
func (s ruleChainService) Edit(editReq iot_schema.RuleChainEditReq) (e error) {
	if e = checkTree(editReq.TreeData); e != nil {
		return
	}
	var row iot_model.IotRuleChain
	if e = response.CheckDBErr(s.db.Where("id = ?", editReq.ID).First(&row).Error, "规则链不存在!", "编辑失败"); e != nil {
		return
	}
	convert_util.Copy(&row, editReq)
	if e = response.CheckErr(s.db.Model(&row).Select("*").Omit("id", "create_time").Updates(row).Error, "编辑失败"); e != nil {
		return
	}
	InvalidateChainCache(editReq.ID)
	return
}

// Del 删除规则链：被产品引用时禁止删除
func (s ruleChainService) Del(id string) (e error) {
	var refCount int64
	if e = response.CheckErr(s.db.Model(&iot_model.IotProduct{}).
		Where("rule_chain_id = ? AND is_delete = 0", id).Count(&refCount).Error, "统计引用产品失败"); e != nil {
		return
	}
	if refCount > 0 {
		return response.AssertArgumentError.SetMessage("该规则链已被产品引用,请先解除引用!")
	}
	result := s.db.Where("id = ?", id).Delete(&iot_model.IotRuleChain{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("规则链不存在")
	}
	InvalidateChainCache(id)
	return
}

// CachedTreeRaw 取规则链执行树原文（Redis 缓存 ThingCacheTTL 秒），供处理层逐报文解析执行。
// 命中即返回；库侧不存在时返回空串（处理层按“未配置链”丢弃上报）。
func (s ruleChainService) CachedTreeRaw(chainId string) string {
	if chainId == "" {
		return ""
	}
	key := chainCacheKey(chainId)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if val, err := core.Redis.Get(ctx, key).Result(); err == nil {
		return val // 命中即返回（含空串的“链不存在”缓存）；redis.Nil 走回源
	} else if !isRedisMiss(err) {
		core.Logger.Errorf("读规则链缓存失败 chainId=%s err=%v", chainId, err)
	}

	var row iot_model.IotRuleChain
	raw := ""
	if s.db.Select("tree_data").Where("id = ? AND is_stop = 0", chainId).First(&row).Error == nil {
		raw = row.TreeData
	}
	// 回源即缓存（含空串的“链不存在/停用”标记），避免不存在的链被每条报文反复打库。
	// tree_data 本身已是 JSON 文本，必须原样存：对它再 json.Marshal 会得到带外层引号的
	// 字符串，命中缓存时解析出的是 string 而非对象，报 cannot unmarshal string into RuleTreeNode。
	if err := core.Redis.Set(ctx, key, raw, time.Duration(config.IotConfig.ThingCacheTTL)*time.Second).Err(); err != nil {
		core.Logger.Errorf("写规则链缓存失败 chainId=%s err=%v", chainId, err)
	}
	return raw
}

// InvalidateChainCache 规则链编辑/删除后失效处理层执行树缓存（包级函数，供 service 内部调用）
func InvalidateChainCache(chainId string) {
	if chainId == "" {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := core.Redis.Del(ctx, chainCacheKey(chainId)).Err(); err != nil {
		core.Logger.Errorf("失效规则链缓存失败 chainId=%s err=%v", chainId, err)
	}
}
