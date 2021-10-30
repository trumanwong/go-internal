package cache

import (
	"time"
)

type Cache struct {
	redis  *Redis
	prefix string
}

func NewCache(addr, password string, prefix string, database int) *Cache {
	return &Cache{
		redis:  NewRedisClient(addr, password, database),
		prefix: prefix,
	}
}

// Get 读取缓存
func (this *Cache) Get(key string) (res string) {
	res, _ = this.redis.Get(key, this.prefix)
	return
}

func (this *Cache) GetNoPrefix(key string) (res string) {
	res, _ = this.redis.Get(key, "")
	return
}

// Put 写入缓存
func (this *Cache) Put(key string, value interface{}, expiration time.Duration) {
	this.redis.Set(key, value, expiration, this.prefix)
}

func (this *Cache) Remember(key string, callback Callback, expiration time.Duration) string {
	res := this.Get(key)
	if len(res) > 0 {
		return res
	}
	value := callback()
	this.Put(key, value, expiration)
	res = this.Get(key)
	return res
}

// Forget 删除key
func (this *Cache) Forget(key string) {
	this.redis.Del(key, this.prefix)
}
