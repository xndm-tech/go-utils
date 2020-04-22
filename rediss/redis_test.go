package rediss_test

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errs"
	"github.com/xndm-recommend/go-utils/rediss"
)

const (
	Config_path = "../config/test.yaml"
)

func TestGetRedisClusterItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	redisItem := rediss.ItemInfo{}
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	t.Log(redisItem)
	redisCluster := new(rediss.RedisDbInfo)
	redisCluster.GetRedisConnFromConf(&c, "Redis_cluster")
	t.Log(redisCluster)
	if redisCluster.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis cluster db")
	}
	err = redisItem.ItemSet(redisCluster.GetDb(), "just redis_test", "test1")
	if err != nil {
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemGet(redisCluster.GetDb(), "test1")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just redis_test" {
		t.Error("set redis item in redis db failed")
	}
}

func TestGetRedisItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	t.Log(redisItem)
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	t.Log(redisdb)
	if redisdb.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis db")
	}
	err = redisItem.ItemSet(redisdb.GetDb(), []string{"ugender", "ok"}, "just redis_test")
	if err != nil {
		fmt.Println(err)
		fmt.Println("this error")
	}
	cmd, err := redisItem.ItemGet(redisdb.GetDb(), "just redis_test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cmd.Val())
	//if cmd.Val() !=[]string{"ugender","ok"}{
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
	//_, err = redisItem.itemPHSet(redisdb.redisDataDb, "ugender", fv)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//time.Sleep(2 * time.Second)
	//
	//cmd1, err := redisItem.itemPHGet(redisdb.redisDataDb, "ugender", "redis_test a", "redis_test b")
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
	//err = redisItem.itemHSet(redisdb.redisDataDb, "ugender", "just redis_test", "Hset:test123")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//cmd, err = redisItem.itemHGet(redisdb.redisDataDb, "ugender", "Hset:test123")
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

	//c := config.ConfigEngine{}
	//var err error
	//err = c.Load(Config_path)
	//errs.CheckCommonErr(err)
	//redisItem := new(rediss.ItemInfo)
	//redisItem.getRedisItemFromConf(&c, "Redis_items.test_item")
	//t.Log(redisItem)
	//redisdb := new(rediss.RedisDbInfo)
	//redisdb.getRedisConnFromConf(&c, "Redis")
	//t.Log(redisdb)
	//if redisdb.redisDataDb.Ping().Val() != "PONG" {
	//	t.Error("can't connect to redis db")
	//}
	//err = redisItem.itemHSet(redisdb.redisDataDb, "ugender", "just redis_test", "Hset:test1111")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//cmd, err := redisItem.itemHGet(redisdb.redisDataDb, "ugender", "Hset:test1111")
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
	//fv["redis_test d"] = "just bbbbbbbb"
	//_, err = redisItem.itemPHSetField(redisdb.redisDataDb, "Hset:test1111", fv)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////time.Sleep(2 * time.Second)
	//
	//cmd1, err := redisItem.itemHGet(redisdb.redisDataDb, "redis_test a", "Hset:test1111")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//x := cmd1
	//if x.Val() != "just testaaaa" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}
	//
	//cmd11, err := redisItem.itemPHGetField(redisdb.redisDataDb, "Hset:test1111", "redis_test b", "redis_test d")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//x = cmd11[0]
	//if x.Val() != "just testbbbbbbb" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}
	//x = cmd11[1]
	//if x.Val() != "just bbbbbbbb" {
	//	t.Error(cmd.Val())
	//	t.Error("set redis item in redis db failed")
	//} else {
	//	t.Log("success!!!")
	//}
}
