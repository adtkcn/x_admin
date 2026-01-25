package captcha_service

import (
	"log"
	"sync"
	"x_admin/util/aj-captcha-go/captcha_config"
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

func (c *CaptchaServiceFactory) GetCache() CacheCaptchaInterface {
	key := c.config.CacheType
	c.CacheLock.RLock()
	defer c.CacheLock.RUnlock()
	if _, ok := c.CacheMap[key]; !ok {
		log.Printf("未注册%s类型的Cache", key)
	}
	return c.CacheMap[key]
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

func (c *CaptchaServiceFactory) GetService(key string) CaptchaInterface {
	c.ServiceLock.RLock()
	defer c.ServiceLock.RUnlock()
	if _, ok := c.ServiceMap[key]; !ok {
		log.Printf("未注册%s类型的Service", key)
	}
	return c.ServiceMap[key]
}
