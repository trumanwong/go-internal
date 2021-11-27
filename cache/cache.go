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
func (this *Cache) Get(key string) (res string, err error) {
	res, err = this.redis.Get(key, this.prefix)
	return
}

func (this *Cache) GetNoPrefix(key string) (res string, err error) {
	res, err = this.redis.Get(key, "")
	return
}

// Put 写入缓存
func (this *Cache) Put(key string, value interface{}, expiration time.Duration) (string, error) {
	res, err := this.redis.Set(key, value, expiration, this.prefix)
	return res, err
}

func (this *Cache) Remember(key string, callback Callback, expiration time.Duration) (string, error) {
	res, err := this.Get(key)
	if err != nil {
		return "", err
	}
	if len(res) > 0 {
		return res, nil
	}
	value := callback()
	res, err = this.Put(key, value, expiration)
	if err != nil {
		return "", err
	}
	res, err = this.Get(key)
	return res, err
}

// Forget 删除key
func (this *Cache) Forget(key string) (int64, error) {
	count, err := this.redis.Del(key, this.prefix)
	return count, err
}
