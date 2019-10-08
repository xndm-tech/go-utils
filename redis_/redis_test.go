package redis__test

import (
	"fmt"
	"testing"

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
//	err = redisItem.ItemSet(redisCluster.RedisDataDb, "just redis_test", "test1")
//	if err != nil {
//		fmt.Println(err)
//	}
//	cmd, err := redisItem.ItemGet(redisCluster.RedisDataDb, "test1")
//	if err != nil {
//		fmt.Println(err)
//	}
//	if cmd.Val() != "just redis_test" {
//		t.Error("set redis item in redis db failed")
//	}
//}

func TestGetRedisItemFromConf(t *testing.T) {
	//c := config.ConfigEngine{}
	//var err error
	//err = c.Load(Config_path)
	//errors_.CheckCommonErr(err)
	//redisItem := new(redis_.RedisItem)
	//redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	//t.Log(redisItem)
	//redisdb := new(redis_.RedisDb)
	//redisdb.GetRedisConnFromConf(&c, "Redis")
	//t.Log(redisdb)
	//if redisdb.RedisDataDb.Ping().Val() != "PONG" {
	//	t.Error("can't connect to redis db")
	//}
	//err = redisItem.ItemHSet(redisdb.RedisDataDb, "ugender", "just redis_test", "Hset:test1111")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//cmd, err := redisItem.ItemHGet(redisdb.RedisDataDb, "ugender", "Hset:test1111")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if cmd.Val() != "just redis_test" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}
	//
	//var fv map[string]string
	//fv = make(map[string]string)
	//fv["redis_test a"] = "just testaaaa"
	//fv["redis_test b"] = "just testbbbbbbb"
	//fv["redis_test c"] = "just asdfasdf"
	//fv["redis_test b"] = "just bbbbbbbb"
	//_, err = redisItem.ItemPHSet(redisdb.RedisDataDb, "ugender", fv)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//time.Sleep(2 * time.Second)
	//
	//cmd1, err := redisItem.ItemPHGet(redisdb.RedisDataDb, "ugender", "redis_test a", "redis_test b")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//x := cmd1[0]
	//if x.Val() != "just testaaaa" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}
	//x = cmd1[1]
	//if x.Val() != "just bbbbbbbb" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}

	//// redis hset add expire
	//err = redisItem.ItemHSet(redisdb.RedisDataDb, "ugender", "just redis_test", "Hset:test123")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//cmd, err = redisItem.ItemHGet(redisdb.RedisDataDb, "ugender", "Hset:test123")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if cmd.Val() != "just redis_test" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}
	//time.Sleep(3 * time.Second)
	//if cmd.Val() != "just redis_test" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}
	//
	//fmt.Println("begin")
	//time.Sleep(2 * time.Second)
	//fmt.Println("end")

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
	err = redisItem.ItemHSet(redisdb.RedisDataDb, "ugender", "just redis_test", "Hset:test1111")
	if err != nil {
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemHGet(redisdb.RedisDataDb, "ugender", "Hset:test1111")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just redis_test" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}

	var fv map[string]string
	fv = make(map[string]string)
	fv["redis_test a"] = "just testaaaa"
	fv["redis_test b"] = "just testbbbbbbb"
	fv["redis_test c"] = "just asdfasdf"
	fv["redis_test d"] = "just bbbbbbbb"
	_, err = redisItem.ItemPHSetField(redisdb.RedisDataDb, "Hset:test1111", fv)
	if err != nil {
		fmt.Println(err)
	}

	//time.Sleep(2 * time.Second)

	cmd1, err := redisItem.ItemHGet(redisdb.RedisDataDb, "redis_test a", "Hset:test1111")
	if err != nil {
		fmt.Println(err)
	}
	x := cmd1
	if x.Val() != "just testaaaa" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}

	cmd11, err := redisItem.ItemPHGetField(redisdb.RedisDataDb, "Hset:test1111", "redis_test b", "redis_test d")
	if err != nil {
		fmt.Println(err)
	}
	x = cmd11[0]
	if x.Val() != "just testbbbbbbb" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}
	x = cmd11[1]
	if x.Val() != "just bbbbbbbb" {
		t.Error(cmd.Val())
		t.Error("set redis item in redis db failed")
	} else {
		t.Log("success!!!")
	}
}
