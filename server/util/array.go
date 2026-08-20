package util

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
