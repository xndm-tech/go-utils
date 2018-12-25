package rediss

import (
	"github.com/go-redis/redis"
	"github.com/zhanglanhui/go-utils/utils/conf_utils"
)

/*
有关redis cluster连接的封装
*/
type RedisDbInfo struct {
	RedisDataDb *redis.ClusterClient
	PoolSize    int
	DataTime    int
}

func GetRedisConnFromConf(this *conf_utils.ConfigEngine, SectionName string) *RedisDbInfo {
	RedisDbInfo := new(RedisDbInfo)
	redis_login := getRedisDataFromConf(this, SectionName)
	RedisDbInfo.RedisDataDb = createClusterClient(redis_login)
	RedisDbInfo.PoolSize = redis_login.Pool_size
	RedisDbInfo.DataTime = redis_login.Data_time
	return RedisDbInfo
}
