package rediss_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xndm-tech/go-utils/tools/errs"

	"github.com/xndm-tech/go-utils/config"
	"github.com/xndm-tech/go-utils/dbs/rediss"
	"github.com/xndm-tech/go-utils/tools/converter"
)

const (
	Config_path = "../../config/test.yaml"
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
	err = redisItem.ItemZAdd(redisdb.GetDb(), converter.IntsToStrs([]int{3, 4, 5, 6, 7}), "just redis_test")
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

func TestGetRedisItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	if nil != err {
		panic(err)
	}
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	fmt.Println("redisItem.GetExpire().Seconds()", redisItem.GetExpire().Seconds())
	t.Log(redisItem)
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	//t.Log(redisdb)
	if redisdb.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis db")
	}
	err = redisItem.ItemZAdd(redisdb.GetDb(), []string{"ugender2", "ok1"}, "just redis_test")
	if err != nil {
		fmt.Println(err)
		fmt.Println("this error")
	}
	time.Sleep(5 * time.Second)
	cmd, err := redisItem.ItemGetZRange(redisdb.GetDb(), "just redis_test")
	if err != nil {
		fmt.Println(err)
	}
	for _, x := range cmd {
		fmt.Println(x)
	}
	fmt.Println("获取数据长度：", len(cmd))
}
