package rediss

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/zhanglanhui/go-utils/utils/conf_utils"
	"github.com/zhanglanhui/go-utils/utils/err_utils"
)

type RedisItem struct {
	KeyPrefix  string
	ExpireTime time.Duration
	Len        int64
}

func (r *RedisItem) getKey(items ...string) string {
	return r.KeyPrefix + "_" + strings.Join(items, "_")
}

func (r *RedisItem) ItemSetByte(redisClient *RedisDbInfo, bytes []byte, items ...string) error {
	return redisClient.RedisDataDb.Set(r.getKey(items...), bytes, r.ExpireTime).Err()
}

func (r *RedisItem) ItemSet(redisClient *RedisDbInfo, value interface{}, items ...string) error {
	return redisClient.RedisDataDb.Set(r.getKey(items...), value, r.ExpireTime).Err()
}

func (r *RedisItem) ItemGet(redisClient *RedisDbInfo, items ...string) (string, error) {
	stringCmd := redisClient.RedisDataDb.Get(r.getKey(items...))
	return stringCmd.Val(), stringCmd.Err()
}

func (r *RedisItem) ItemIncr(redisClient *RedisDbInfo, items ...string) (int, error) {
	key := r.getKey(items...)
	p := redisClient.RedisDataDb.Pipeline()
	cmder := p.Incr(key)
	p.Expire(key, r.ExpireTime)
	p.Exec()
	val, err := cmder.Result()
	err_utils.CheckCommonErr(err)
	return int(val), err
}

func (r *RedisItem) ItemZAdd(redisClient *RedisDbInfo, ids []string, items ...string) error {
	key := r.getKey(items...)
	zmembers := make([]redis.Z, 0, len(ids))
	for _, id := range ids {
		zmembers = append(zmembers, redis.Z{Score: float64(time.Now().UnixNano()), Member: id})
	}
	p := redisClient.RedisDataDb.Pipeline()
	err := p.ZAdd(key, zmembers...).Err()
	err_utils.CheckCommonErr(err)
	cmdSetLen := p.ZCard(key)
	p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.Len {
		err := redisClient.RedisDataDb.ZRemRangeByRank(key, 0, setLen-r.Len).Err()
		err_utils.CheckCommonErr(err)
	}
	return err
}

func (r *RedisItem) ItemGetZRange(redisClient *RedisDbInfo, items ...string) ([]string, error) {
	key := r.getKey(items...)
	result, err := redisClient.RedisDataDb.ZRange(key, 0, -1).Result()
	err_utils.CheckCommonErr(err)
	return result, err
}

func (r *RedisItem) ItemPGet(redisClient *RedisDbInfo, ids []string) ([]string, error) {
	var cmders []string
	p := redisClient.RedisDataDb.Pipeline()
	for _, id := range ids {
		key := r.getKey(id)
		cmd := p.Get(key)
		cmders = append(cmders, cmd.Val())
	}
	_, err := p.Exec()
	return cmders, err
}

func GetRedisItemInfoFromConf(c *conf_utils.ConfigEngine, sectionName string) *RedisItem {
	//interfaceMap := c.Get(name)
	//itemMap := interfaceMap.(map[interface{}]interface{})
	//ritem.KeyPrefix = itemMap["Prefix"].(string)
	//errors.CheckEmptyValue(ritem.KeyPrefix)
	//ExpireSeconds := itemMap["Expire"].(int)
	//ritem.ExpireTime = time.Duration(ExpireSeconds) * time.Second
	//length := itemMap["Len"].(int)
	//ritem.Len = int64(length)
	//return ritem
	login := new(RedisItem)
	redisLogin := c.GetStruct(sectionName, login)
	return redisLogin.(*RedisItem)
}
