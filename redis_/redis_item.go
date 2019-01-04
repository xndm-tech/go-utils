package redis_

import (
	"strings"
	"time"

	"github.com/xndm-recommend/go-utils/config"

	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/errors_"
)

type RedisItemMethod interface {
	GetRedisItemFromConf(c *config.ConfigEngine, name string)

	ItemSetByte(redisClient *RedisDbInfo, bytes []byte, items ...string) error

	ItemSet(redisClient *RedisDbInfo, value interface{}, items ...string) error

	ItemGet(redisClient *RedisDbInfo, items ...string) (string, error)

	ItemIncr(redisClient *RedisDbInfo, items ...string) (int, error)

	ItemZAdd(redisClient *RedisDbInfo, ids []string, items ...string) error

	ItemGetZRange(redisClient *RedisDbInfo, items ...string) ([]string, error)

	ItemPGet(redisClient *RedisDbInfo, ids []string) ([]string, error)
}

type RedisItem struct {
	KeyPrefix  string
	ExpireTime time.Duration
	Len        int64
}

func (r *RedisItem) GetKey(items ...string) string {
	return r.KeyPrefix + "_" + strings.Join(items, "_")
}

func (r *RedisItem) ItemSetByte(redisClient *RedisDbInfo, bytes []byte, items ...string) error {
	return redisClient.RedisDataDb.Set(r.GetKey(items...), bytes, r.ExpireTime).Err()
}

func (r *RedisItem) ItemSet(redisClient *RedisDbInfo, value interface{}, items ...string) error {
	return redisClient.RedisDataDb.Set(r.GetKey(items...), value, r.ExpireTime).Err()
}

func (r *RedisItem) ItemGet(redisClient *RedisDbInfo, items ...string) (*redis.StringCmd, error) {
	stringCmd := redisClient.RedisDataDb.Get(r.GetKey(items...))
	return stringCmd, stringCmd.Err()
}

func (r *RedisItem) ItemIncrExpire(redisClient *RedisDbInfo, items ...string) (int, error) {
	key := r.GetKey(items...)
	p := redisClient.RedisDataDb.Pipeline()
	cmder := p.Incr(key)
	p.Expire(key, r.ExpireTime)
	p.Exec()
	val, err := cmder.Result()
	errors_.CheckCommonErr(err)
	return int(val), err
}

func (r *RedisItem) ItemZAdd(redisClient *RedisDbInfo, ids []string, items ...string) error {
	key := r.GetKey(items...)
	zmembers := make([]redis.Z, 0, len(ids))
	for _, id := range ids {
		zmembers = append(zmembers, redis.Z{Score: float64(time.Now().UnixNano()), Member: id})
	}
	p := redisClient.RedisDataDb.Pipeline()
	err := p.ZAdd(key, zmembers...).Err()
	errors_.CheckCommonErr(err)
	cmdSetLen := p.ZCard(key)
	p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.Len {
		err := redisClient.RedisDataDb.ZRemRangeByRank(key, 0, setLen-r.Len).Err()
		errors_.CheckCommonErr(err)
	}
	return err
}

func (r *RedisItem) ItemGetZRange(redisClient *RedisDbInfo, items ...string) ([]string, error) {
	key := r.GetKey(items...)
	result, err := redisClient.RedisDataDb.ZRange(key, 0, -1).Result()
	errors_.CheckCommonErr(err)
	return result, err
}

func (r *RedisItem) ItemPGet(redisClient *RedisDbInfo, ids []string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := redisClient.RedisDataDb.Pipeline()
	for _, id := range ids {
		key := r.GetKey(id)
		cmd := p.Get(key)
		cmders = append(cmders, cmd)
	}
	_, err := p.Exec()
	return cmders, err
}

func (r *RedisItem) GetRedisItemFromConf(c *config.ConfigEngine, name string) {
	login := new(RedisItem)
	redisLogin := c.GetStruct(name, login)
	r = redisLogin.(*RedisItem)
}
