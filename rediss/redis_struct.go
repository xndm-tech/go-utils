package rediss

import (
	"github.com/go-redis/redis"
	"github.com/zhanglanhui/go-utils/utils/conf_utils"
	"github.com/zhanglanhui/go-utils/utils/err_utils"
)

type redisYamlData struct {
	Master_host []string `yaml:"Master_host"`
	Master_port []string `yaml:"Master_port"`
	Slave_host  []string `yaml:"Slave_host"`
	Slave_port  []string `yaml:"Slave_port"`
	Password    string   `yaml:"Password"`
	Nodes       int      `yaml:"Nodes"`
	Data_time   int      `yaml:"Data_time"`
	Pool_size   int      `yaml:"Pool_size"`
}

func getRedisNodes(redisConf *redisYamlData, i int) (clusterNodes []redis.ClusterNode) {
	clusterNodes = append(clusterNodes, redis.ClusterNode{
		Addr: redisConf.Master_host[i] + ":" + redisConf.Master_port[i]})
	clusterNodes = append(clusterNodes, redis.ClusterNode{
		Addr: redisConf.Slave_host[i] + ":" + redisConf.Slave_port[i]})
	return clusterNodes
}

func readRedisCluster(redisConf *redisYamlData) (redisClusters [][]redis.ClusterNode) {
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

func getRedisDataFromConf(this *conf_utils.ConfigEngine, sectionName string) *redisYamlData {
	login := new(redisYamlData)
	redisLogin := this.GetStruct(sectionName, login)
	return redisLogin.(*redisYamlData)
}

func createClusterClient(redisConf *redisYamlData) *redis.ClusterClient {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots: getRedisSlots(readRedisCluster(redisConf)),
		PoolSize:     redisConf.Pool_size,
		Password:     redisConf.Password,
	})
	_, err := redisdb.Ping().Result()
	err_utils.CheckFatalErr(err)
	return redisdb
}
