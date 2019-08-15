package redis__test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/redis_"
)

const (
	Config_path = "../config/test.yaml"
)

//func TestGetRedisClusterItemFromConf(t *testing.T) {
//	c := config.ConfigEngine{}
//	var err error
//	err = c.Load(Config_path)
//	errors_.CheckCommonErr(err)
//	redisItem := new(redis_.RedisItem)
//	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
//	t.Log(redisItem)
//	redisCluster := new(redis_.RedisDbInfo)
//	redisCluster.GetRedisConnFromConf(&c, "Redis_cluster")
//	t.Log(redisCluster)
//	if redisCluster.RedisDataDb.Ping().Val() != "PONG" {
//		t.Error("can't connect to redis cluster db")
//	}
//	err = redisItem.ItemSet(redisCluster.RedisDataDb, "just test", "test1")
//	if err != nil {
//		fmt.Println(err)
//	}
//	cmd, err := redisItem.ItemGet(redisCluster.RedisDataDb, "test1")
//	if err != nil {
//		fmt.Println(err)
//	}
//	if cmd.Val() != "just test" {
//		t.Error("set redis item in redis db failed")
//	}
//}

func TestGetRedisItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errors_.CheckCommonErr(err)
	redisItem := new(redis_.RedisItem)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	t.Log(redisItem)
	redisdb := new(redis_.RedisDb)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	t.Log(redisdb)
	if redisdb.RedisDataDb.Ping().Val() != "PONG" {
		t.Error("can't connect to redis db")
	}
	err = redisItem.ItemHSet(redisdb.RedisDataDb, "ugender", "just test", "Hset:test1111")
	if err != nil {
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemHGet(redisdb.RedisDataDb, "ugender", "Hset:test1111")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just test" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}

	var kv map[string]string
	kv = make(map[string]string)
	kv["a"] = "just testaaaa"
	kv["b"] = "just testbbbbbbb"
	kv["c"] = "just asdfasdf"
	kv["b"] = "just bbbbbbbb"
	_, err = redisItem.ItemPHSet(redisdb.RedisDataDb, "ugender", kv)
	if err != nil {
		fmt.Println(err)
	}
	cmd1, err := redisItem.ItemPHGet(redisdb.RedisDataDb, "ugender", "a", "b")
	if err != nil {
		fmt.Println(err)
	}
	x := cmd1[0]
	if x.Val() != "just testaaaa" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}
	x = cmd1[1]
	if x.Val() != "just bbbbbbbb" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}

	// redis hset add expire
	err = redisItem.ItemHSet(redisdb.RedisDataDb, "ugender", "just test", "Hset:test123")
	if err != nil {
		fmt.Println(err)
	}
	cmd, err = redisItem.ItemHGet(redisdb.RedisDataDb, "ugender", "Hset:test123")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just test" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}
	time.Sleep(3 * time.Second)
	if cmd.Val() != "just test" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}

	fmt.Println("begin")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
}
