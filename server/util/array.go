package util

var ArrayUtil = arrayUtil{}

// arrayUtil 数组工具类
type arrayUtil struct{}

// ListToTree 列表转树形结构
func (au arrayUtil) ListToTree(arr []map[string]any, id string, pid string, child string) (mapList []any) {
	mapList = []any{}
	// 遍历以id_为key生成map
	idValMap := make(map[string]any)
	for _, m := range arr {
		if idVal, ok := m[id]; ok {
			idValMap[idVal.(string)] = m
		}
	}
	// 遍历
	for _, m := range arr {
		// 获取父节点
		if pidVal, ok := m[pid]; ok {
			if pNode, pok := idValMap[pidVal.(string)]; pok {
				// 有父节点则添加到父节点子集
				if cVal, cok := pNode.(map[string]any)[child]; cok {
					if cVal == nil {
						cVal = []any{m}
					} else {
						cVal = append(cVal.([]any), m)
					}
					pNode.(map[string]any)[child] = cVal
					continue
				} else {
					cVal := []any{m}
					pNode.(map[string]any)[child] = cVal
					continue
				}
			}
		}
		mapList = append(mapList, m)
	}
	return
}

// TreeItem 树节点接口
type TreeItem[T any] interface {
	GetID() string       // Unique identifier of the node
	GetParentID() string // Parent's identifier; root nodes usually have ParentID == rootParentID
	SetChildren([]T)     // Assign children to this node
}

// ListToTree 列表转树形结构 泛型版本
func ListToTree[T TreeItem[T]](items []T, rootParentID string) []T {
	if len(items) == 0 {
		return nil
	}

	// Map from node ID to node (for quick lookup)
	itemMap := make(map[string]T, len(items))
	for _, item := range items {
		itemMap[item.GetID()] = item
	}

	// Group children by parent ID
	childrenMap := make(map[string][]T)
	for _, item := range items {
		parentID := item.GetParentID()
		childrenMap[parentID] = append(childrenMap[parentID], item)
	}

	// Assign children to each node
	for _, item := range items {
		id := item.GetID()
		if children, exists := childrenMap[id]; exists {
			item.SetChildren(children)
		} else {
			item.SetChildren(nil)
		}
	}

	// Return all root nodes
	return childrenMap[rootParentID]
}
