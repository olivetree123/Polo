package utils

import "time"

type CacheValue struct {
	value  interface{}
	expire int64 // 过期时间
}

type Cache map[string]*CacheValue

func NewCacheValue(value interface{}, timeout int64) *CacheValue {
	return &CacheValue{
		value:  value,
		expire: timeout,
	}
}

func NewCache() *Cache {
	m := make(map[string]*CacheValue)
	c := Cache(m)
	return &c
}

func (c *Cache) Add(key string, value interface{}) {
	map[string]*CacheValue(*c)[key] = NewCacheValue(value, 0)
}

func (c *Cache) AddWithExpire(key string, value interface{}, timeout int64) {
	expire := time.Now().Unix() + timeout
	map[string]*CacheValue(*c)[key] = NewCacheValue(value, expire)
}

func (c *Cache) Get(key string) interface{} {
	m := map[string]*CacheValue(*c)
	if r, found := m[key]; found {
		if r.expire > 0 && r.expire < time.Now().Unix() {
			return nil
		}
		return r.value
	}
	return nil
}

var cache *Cache

func init() {
	cache = NewCache()
}

func SetCache(key string, value interface{}) {
	cache.Add(key, value)
}

func SetCacheWithExpire(key string, value interface{}, timeout int64) {
	cache.AddWithExpire(key, value, timeout)
}

func GetCache(key string) {
	cache.Get(key)
}
