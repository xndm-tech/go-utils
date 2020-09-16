package bitmap

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/dbs/rediss"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

const length = 30000

func TestBitMapFilterByRedis(t *testing.T) {
	fmt.Println("11")
	f := []int64{1, 2, 3}
	can := []int64{1, 2, 3, 4, 5, 6}
	c := config.ConfigEngine{}
	var err error
	err = c.Load("../../config/test.yaml")
	errs.CheckCommonErr(err)
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	kk := RedisBitMap{
		RedisCli:  redisdb,
		RedisItem: redisItem,
	}
	err = kk.MultiSetBit("ddd", true, f...)
	fmt.Println("11")
	xx, _ := kk.FilterByRedis("ddd", can...)
	t.Log("22")
	fmt.Println("xx", xx)
}

//
//func BenchmarkFilterByRedis(b *testing.B) {
//	b.StopTimer()
//	can := maths.GenNRandInt64(length, 100)
//	//canstr := converter.IntsToStrs(can)
//	c := config.ConfigEngine{}
//	var err error
//	err = c.Load("../../config/test.yaml")
//	errs.CheckCommonErr(err)
//	redisItem := new(rediss.ItemInfo)
//	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
//	redisdb := new(rediss.RedisDbInfo)
//	redisdb.GetRedisConnFromConf(&c, "Redis")
//	kk := RedisBitMap{
//		RedisCli:  redisdb,
//		RedisItem: redisItem,
//	}
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		_, _ = kk.FilterByRedis("ddd", can...)
//	}
//}
