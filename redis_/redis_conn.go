package redis_

import (
	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

type RedisConnMethod interface {
	GetRedisConnFromConf(c *config.ConfigEngine, name string)
}

/*
有关redis cluster连接的封装
*/
type RedisDbInfo struct {
	RedisDataDb *redis.ClusterClient
	PoolSize    int
	DataTime    int
}

func getRedisNodes(redisConf *config.RedisData, i int) (clusterNodes []redis.ClusterNode) {
	clusterNodes = append(clusterNodes, redis.ClusterNode{
		Addr: redisConf.Master_host[i] + ":" + redisConf.Master_port[i]})
	clusterNodes = append(clusterNodes, redis.ClusterNode{
		Addr: redisConf.Slave_host[i] + ":" + redisConf.Slave_port[i]})
	return clusterNodes
}

func readRedisCluster(redisConf *config.RedisData) (redisClusters [][]redis.ClusterNode) {
	for i := 0; i < redisConf.Nodes; i++ {
		redisClusters = append(redisClusters, getRedisNodes(redisConf, i))
	}
	return redisClusters
}

func getRedisSlots(redisClusters [][]redis.ClusterNode) func() ([]redis.ClusterSlot, error) {
	return func() ([]redis.ClusterSlot, error) {
		var slots []redis.ClusterSlot
		nodes := len(redisClusters)
		var gap int
		gap = 16383 / nodes
		for i, redisClusterNodes := range redisClusters {
			slots = append(slots, redis.ClusterSlot{
				Start: i * gap,
				End:   16383 - (nodes-i-1)*gap,
				Nodes: redisClusterNodes,
			})
		}
		return slots, nil
	}
}

func createClusterClient(redisConf *config.RedisData) *redis.ClusterClient {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots: getRedisSlots(readRedisCluster(redisConf)),
		PoolSize:     redisConf.Pool_size,
		Password:     redisConf.Password,
	})
	_, err := redisdb.Ping().Result()
	errors_.CheckFatalErr(err)
	return redisdb
}

func (this *RedisDbInfo) GetRedisConnFromConf(c *config.ConfigEngine, name string) {
	redis_login := c.GetRedisDataFromConf(name)
	this.RedisDataDb = createClusterClient(redis_login)
	this.PoolSize = redis_login.Pool_size
	this.DataTime = redis_login.Data_time
}
