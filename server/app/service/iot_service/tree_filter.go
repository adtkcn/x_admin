package iot_service

import (
	"x_admin/app/model/iot_model"

	"gorm.io/gorm"
)

// idPid 树节点的 id 与其上级 id（区域树/分类树都用 pid 邻接表建模）
type idPid struct {
	ID  string
	Pid string
}

// collectSubtree 从扁平的 (id,pid) 关系收集 rootId 及其全部后代 id（含自身）。
// 一次读入邻接表后在内存做 BFS，避免为求子树反复查库。
func collectSubtree(pairs []idPid, rootId string) []string {
	children := make(map[string][]string, len(pairs))
	for _, p := range pairs {
		children[p.Pid] = append(children[p.Pid], p.ID)
	}
	out := []string{rootId}
	for i := 0; i < len(out); i++ {
		out = append(out, children[out[i]]...)
	}
	return out
}

// spaceSubtreeIds 区域 id 及其整棵子树的 id：设备/告警列表按区域筛选时含子区域
func spaceSubtreeIds(db *gorm.DB, spaceId string) []string {
	var spaces []iot_model.IotDeviceSpace
	db.Select("id, pid").Find(&spaces)
	pairs := make([]idPid, 0, len(spaces))
	for _, sp := range spaces {
		pairs = append(pairs, idPid{ID: sp.ID, Pid: sp.Pid})
	}
	return collectSubtree(pairs, spaceId)
}

// categorySubtreeIds 分类 id 及其整棵子树的 id：产品列表按分类筛选时含子分类
func categorySubtreeIds(db *gorm.DB, categoryId string) []string {
	var cats []iot_model.IotProductCategory
	db.Select("id, pid").Find(&cats)
	pairs := make([]idPid, 0, len(cats))
	for _, c := range cats {
		pairs = append(pairs, idPid{ID: c.ID, Pid: c.Pid})
	}
	return collectSubtree(pairs, categoryId)
}
