package redis__test

import (
	"fmt"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/redis_"
	"testing"
)

const (
	Config_path = "../config/test.yaml"
)

func TestGetRedisItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errors_.CheckCommonErr(err)
	redisItem := new(redis_.RedisItem)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	t.Log(redisItem)
	redisCluster := new(redis_.RedisDbInfo)
	redisCluster.GetRedisConnFromConf(&c, "Redis_cluster")
	t.Log(redisCluster)
	redisdb := new(redis_.RedisDb)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	t.Log(redisdb)
	if redisCluster.RedisDataDb.Ping().Val() != "PONG"{
		t.Error("can't connect to redis cluster db")
	}
	if redisdb.RedisDataDb.Ping().Val() != "PONG"{
		t.Error("can't connect to redis db")
	}
	err= redisItem.ItemSet(redisdb.RedisDataDb, "just test", "test1")
	if err != nil{
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemGet(redisdb.RedisDataDb, "test1")
	if err != nil{
		fmt.Println(err)
	}
	if cmd.Val() != "just test"{
		t.Error("set redis item in redis db failed")
	}
	err = redisItem.ItemSet(redisCluster.RedisDataDb, "just test", "test1")
	if err != nil{
		fmt.Println(err)
	}
	cmd, err = redisItem.ItemGet(redisCluster.RedisDataDb, "test1")
	if err != nil{
		fmt.Println(err)
	}
	if cmd.Val() != "just test"{
		t.Error("set redis item in redis db failed")
	}
}
