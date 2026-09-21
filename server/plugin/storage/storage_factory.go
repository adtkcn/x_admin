package storage

import "sync"

// ---- 全局存储引擎实例（单例，仅支持本地存储） ----

var (
	_storageEngine     StorageEngine
	_storageEngineOnce sync.Once
)

// GetStorageEngine 获取全局存储引擎实例（懒加载，仅本地存储）
func GetStorageEngine() StorageEngine {
	_storageEngineOnce.Do(func() {
		_storageEngine = newLocalStorageEngine()
	})
	return _storageEngine
}
