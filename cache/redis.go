package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	ctx context.Context
	rdb *redis.Client
}

type Callback func() interface{}

func NewRedisClient(addr, password string, db int) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ctx := context.Background()
	return &Redis{rdb: rdb, ctx: ctx}
}

func CloseRedis(redis *Redis) {
	redis.rdb.Close()
}

func (this *Redis) GetRdb() *redis.Client {
	return this.rdb
}

func (this *Redis) GetCtx() context.Context {
	return this.ctx
}

func getKey(key string, prefix string) string {
	if len(prefix) > 0 {
		return prefix + ":" + key
	}
	return key
}

func (this *Redis) Get(key string, prefix string) (res string, err error) {
	res, err = this.rdb.Get(this.ctx, getKey(key, prefix)).Result()
	return
}

func (this *Redis) Set(key string, value interface{}, expiration time.Duration, prefix string) (res string, err error) {
	res, err = this.rdb.Set(
		this.ctx,
		getKey(key, prefix),
		value,
		expiration*time.Second,
	).Result()
	return
}

// Exists 指定的key是否存在
func (this *Redis) Exists(key string, prefix string) (bool, error) {
	count, err := this.rdb.Exists(this.ctx, getKey(key, prefix)).Result()
	return count > 0, err
}

// Del 删除指定的key
func (this *Redis) Del(key string, prefix string) (int64, error) {
	count, err := this.rdb.Del(this.ctx, getKey(key, prefix)).Result()
	return count, err
}
