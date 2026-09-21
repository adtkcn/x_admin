package pubsub

import "sync"

// handlerEntry 以自增 id 标识一次订阅，规避 Go 中函数不可直接比较的限制。
type handlerEntry struct {
	id int
	h  Handler
}

// registry 是按事件类型组织的线程安全 handler 注册表，
// 供各后端（当前 redis）共用（负责订阅存储与本地分发）。
type registry struct {
	mu       sync.RWMutex
	handlers map[string][]handlerEntry
	nextID   int
}

// newRegistry 创建 handler 注册表。
func newRegistry() *registry {
	return &registry{handlers: make(map[string][]handlerEntry)}
}

// on 注册一个 handler，返回用于取消该次订阅的函数。
func (r *registry) on(eventType string, h Handler) func() {
	r.mu.Lock()
	r.nextID++
	id := r.nextID
	r.handlers[eventType] = append(r.handlers[eventType], handlerEntry{id: id, h: h})
	r.mu.Unlock()
	return func() { r.remove(eventType, id) }
}

// off 移除某事件类型下的全部 handler。
func (r *registry) off(eventType string) {
	r.mu.Lock()
	delete(r.handlers, eventType)
	r.mu.Unlock()
}

// remove 按 id 移除单个 handler。
func (r *registry) remove(eventType string, id int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	hs := r.handlers[eventType]
	for i, x := range hs {
		if x.id == id {
			r.handlers[eventType] = append(hs[:i], hs[i+1:]...)
			break
		}
	}
	if len(r.handlers[eventType]) == 0 {
		delete(r.handlers, eventType)
	}
}

// dispatch 将 payload 同步分发给该类型及通配符("*")的订阅者。
// 快照后再调用，避免 handler 内再次订阅/取消造成死锁或迭代问题。
func (r *registry) dispatch(eventType string, payload []byte) {
	r.mu.RLock()
	hs := make([]Handler, 0, len(r.handlers[eventType])+len(r.handlers[All]))
	for _, he := range r.handlers[eventType] {
		hs = append(hs, he.h)
	}
	if eventType != All {
		for _, he := range r.handlers[All] {
			hs = append(hs, he.h)
		}
	}
	r.mu.RUnlock()

	for _, h := range hs {
		h(payload)
	}
}

// reset 清空全部订阅（关闭时调用）。
func (r *registry) reset() {
	r.mu.Lock()
	r.handlers = make(map[string][]handlerEntry)
	r.mu.Unlock()
}
