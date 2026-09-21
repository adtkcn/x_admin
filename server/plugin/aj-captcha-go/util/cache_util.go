package util

import (
	"log"
	"sync"
	"time"
)

type cacheEntry struct {
	value    string
	expireAt int64 // 0 表示永不过期
}

type CacheUtil struct {
	Data                  map[string]cacheEntry
	DataRWLock            sync.RWMutex
	CaptchaCacheMaxNumber int
}

func NewCacheUtil(captchaCacheMaxNumber int) *CacheUtil {
	return &CacheUtil{
		Data:                  make(map[string]cacheEntry),
		CaptchaCacheMaxNumber: captchaCacheMaxNumber,
	}
}

func (l *CacheUtil) Exists(key string) bool {
	l.DataRWLock.RLock()
	e, ok := l.Data[key]
	l.DataRWLock.RUnlock()
	if !ok {
		return false
	}
	if e.expireAt == 0 {
		return true
	}
	if e.expireAt < time.Now().Unix() {
		l.Delete(key)
		return false
	}
	return true
}

func (l *CacheUtil) Get(key string) string {
	l.DataRWLock.RLock()
	e, ok := l.Data[key]
	l.DataRWLock.RUnlock()
	if !ok {
		return ""
	}
	if e.expireAt != 0 && e.expireAt < time.Now().Unix() {
		l.Delete(key)
		return ""
	}
	return e.value
}

func (l *CacheUtil) Set(key string, val string, expiresInSeconds int) {
	// 设置阈值，达到即 clear 缓存
	if len(l.Data) >= l.CaptchaCacheMaxNumber {
		log.Println("CACHE_MAP达到阈值，clear map")
		l.Clear()
	}

	var expireAt int64
	if expiresInSeconds > 0 {
		expireAt = time.Now().Unix() + int64(expiresInSeconds)
	}

	l.DataRWLock.Lock()
	l.Data[key] = cacheEntry{value: val, expireAt: expireAt}
	l.DataRWLock.Unlock()
}

func (l *CacheUtil) Delete(key string) {
	l.DataRWLock.Lock()
	defer l.DataRWLock.Unlock()
	delete(l.Data, key)
}

func (l *CacheUtil) Clear() {
	l.DataRWLock.Lock()
	defer l.DataRWLock.Unlock()
	l.Data = make(map[string]cacheEntry)
}
