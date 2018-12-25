package redis_

import (
	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/conf_read"
)

/*
有关redis cluster连接的封装
*/
type RedisDbInfo struct {
	RedisDataDb *redis.ClusterClient
	PoolSize    int
	DataTime    int
}

func GetRedisConnFromConf(this *conf_read.ConfigEngine, SectionName string) *RedisDbInfo {
	RedisDbInfo := new(RedisDbInfo)
	redis_login := getRedisDataFromConf(this, SectionName)
	RedisDbInfo.RedisDataDb = createClusterClient(redis_login)
	RedisDbInfo.PoolSize = redis_login.Pool_size
	RedisDbInfo.DataTime = redis_login.Data_time
	return RedisDbInfo
}
