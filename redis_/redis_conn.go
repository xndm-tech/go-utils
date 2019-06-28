package redis_

import (
	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

type RedisDb struct {
	RedisDataDb *redis.Client
	PoolSize    int
}

func createSingleClient(redisConf *config.RedisData) *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		PoolSize: redisConf.Pool_size,
		Password: redisConf.Password,
	})
	_, err := redisdb.Ping().Result()
	errors_.CheckFatalErr(err)
	return redisdb
}

func (this *RedisDb) GetRedisConnFromConf(c *config.ConfigEngine, name string) {
	redis_login := c.GetRedisDataFromConf(name)
	this.RedisDataDb = createSingleClient(redis_login)
	this.PoolSize = redis_login.Pool_size
}
