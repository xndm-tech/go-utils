package rediss

import (
	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

type RedisDbInfo struct {
	redisDataDb *redis.Client
	poolSize    int
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

func (this *RedisDbInfo) GetRedisConnFromConf(c *config.ConfigEngine, name string) {
	redis_login := c.GetRedisDataFromConf(name)
	this.redisDataDb = createSingleClient(redis_login)
	this.poolSize = redis_login.Pool_size
}

func (this *RedisDbInfo) CreateSingleClient(addr, password string, poolSize int) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     addr,
		PoolSize: poolSize,
		Password: password,
	})
	_, err := redisDB.Ping().Result()
	errors_.CheckFatalErr(err)
	this.redisDataDb = redisDB
	this.poolSize = poolSize
}

func (this *RedisDbInfo) GetDb() *redis.Client {
	return this.redisDataDb
}
