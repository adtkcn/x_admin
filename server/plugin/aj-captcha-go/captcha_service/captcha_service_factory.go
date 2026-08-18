package captcha_service

import (
	"fmt"
	"sync"
	"x_admin/plugin/aj-captcha-go/captcha_config"
)

func NewCaptchaServiceFactory(config *captcha_config.Config) *CaptchaServiceFactory {

	factory := &CaptchaServiceFactory{
		ServiceMap: make(map[string]CaptchaInterface),
		CacheMap:   make(map[string]CacheCaptchaInterface),
		config:     config,
	}
	return factory
}

// CaptchaServiceFactory 验证码服务工厂
type CaptchaServiceFactory struct {
	config      *captcha_config.Config
	ServiceMap  map[string]CaptchaInterface
	ServiceLock sync.RWMutex

	CacheMap  map[string]CacheCaptchaInterface
	CacheLock sync.RWMutex
}

func (c *CaptchaServiceFactory) GetCache() (CacheCaptchaInterface, error) {
	key := c.config.CacheType
	c.CacheLock.RLock()
	defer c.CacheLock.RUnlock()
	cache, ok := c.CacheMap[key]
	if !ok {
		return nil, fmt.Errorf("未注册 %s 类型的 Cache", key)
	}
	return cache, nil
}

func (c *CaptchaServiceFactory) RegisterCache(key string, cacheInterface CacheCaptchaInterface) {
	c.CacheLock.Lock()
	defer c.CacheLock.Unlock()
	c.CacheMap[key] = cacheInterface
}

func (c *CaptchaServiceFactory) RegisterService(key string, service CaptchaInterface) {
	c.ServiceLock.Lock()
	defer c.ServiceLock.Unlock()
	c.ServiceMap[key] = service
}

func (c *CaptchaServiceFactory) GetService(key string) (CaptchaInterface, error) {
	c.ServiceLock.RLock()
	defer c.ServiceLock.RUnlock()
	service, ok := c.ServiceMap[key]
	if !ok {
		return nil, fmt.Errorf("未注册 %s 类型的 Service", key)
	}
	return service, nil
}
