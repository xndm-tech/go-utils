package rediss_test

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/types"

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
	redisCli := new(rediss.RedisDbInfo)
	redisCli.GetRedisConnFromConf(&c, "Redis")
	t.Log(redisCli)
	if redisCli.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis cluster db")
	}
	if err = redisItem.ItemSet(redisCli.GetDb(), "just redis_test", "test1"); err != nil {
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemGet(redisCli.GetDb(), "test1")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just redis_test" {
		t.Error("set redis item in redis db failed")
	}
}

//func TestGetRedisItemFromConf(t *testing.T) {
//	c := config.ConfigEngine{}
//	var err error
//	err = c.Load(Config_path)
//	errs.CheckCommonErr(err)
//	redisItem := new(rediss.ItemInfo)
//	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
//	t.Log(redisItem)
//	redisdb := new(rediss.RedisDbInfo)
//	redisdb.GetRedisConnFromConf(&c, "Redis")
//	t.Log(redisdb)
//	if redisdb.GetDb().Ping().Val() != "PONG" {
//		t.Error("can't connect to redis db")
//	}
//	err = redisItem.ItemSet(redisdb.GetDb(), []string{"ugender", "ok"}, "just redis_test")
//	if err != nil {
//		fmt.Println(err)
//		fmt.Println("this error")
//	}
//	cmd, err := redisItem.ItemGet(redisdb.GetDb(), "just redis_test")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(cmd.Val())
//}

func TestGetRedisItemRangeFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item2")
	t.Log(redisItem)
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	t.Log(redisdb)
	if redisdb.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis db")
	}
	err = redisItem.ItemZAdd(redisdb.GetDb(), types.IntsToStrs([]int{3, 4, 5, 6, 7}), "just redis_test")
	if err != nil {
		fmt.Println(err)
		fmt.Println("this error")
	}
	cmd, err := redisItem.ItemGetZRange(redisdb.GetDb(), "just redis_test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cmd)
}
