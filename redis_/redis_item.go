package redis_

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

type RedisItemMethod interface {
	GetRedisItemFromConf(c *config.ConfigEngine, name string)

	ItemSetByte(redisClient *RedisDbInfo, bytes []byte, items ...string) error
	ItemSet(redisClient *RedisDbInfo, value interface{}, items ...string) error
	ItemGet(redisClient *RedisDbInfo, items ...string) (*redis.StringCmd, error)

	ItemIncrExpire(redisClient *RedisDbInfo, items ...string) (int, error)

	ItemZAdd(redisClient *RedisDbInfo, ids []string, items ...string) error
	ItemGetZRange(redisClient *RedisDbInfo, items ...string) ([]string, error)

	ItemSetSAdd(redisClient *RedisDbInfo, ids []string, items ...string) error
	ItemGetSAdd(redisClient *RedisDbInfo, items ...string) ([]string, error)

	// 批量获取
	ItemPGet(redisClient *RedisDbInfo, ids []string) ([]*redis.StringCmd, error)
}

type RedisItem struct {
	Prefix string
	Expire time.Duration
	Len    int64
}

type RedisItemYaml struct {
	Prefix string
	Expire int
	Len    int
}

func (r *RedisItem) getKey(items ...string) string {
	return r.Prefix + "_" + strings.Join(items, "_")
}

// map
func (r *RedisItem) ItemSetByte(redisClient *RedisDbInfo, bytes []byte, items ...string) error {
	return redisClient.RedisDataDb.Set(r.getKey(items...), bytes, r.Expire).Err()
}

func (r *RedisItem) ItemSet(redisClient *RedisDbInfo, value interface{}, items ...string) error {
	return redisClient.RedisDataDb.Set(r.getKey(items...), value, r.Expire).Err()
}

func (r *RedisItem) ItemGet(redisClient *RedisDbInfo, items ...string) (*redis.StringCmd, error) {
	stringCmd := redisClient.RedisDataDb.Get(r.getKey(items...))
	return stringCmd, stringCmd.Err()
}

// Incr
func (r *RedisItem) ItemIncrExpire(redisClient *RedisDbInfo, items ...string) (int, error) {
	key := r.getKey(items...)
	p := redisClient.RedisDataDb.Pipeline()
	cmder := p.Incr(key)
	p.Expire(key, r.Expire)
	p.Exec()
	val, err := cmder.Result()
	errors_.CheckCommonErr(err)
	return int(val), err
}

// 批量获取
func (r *RedisItem) ItemPGet(redisClient *RedisDbInfo, ids []string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := redisClient.RedisDataDb.Pipeline()
	for _, id := range ids {
		key := r.getKey(id)
		cmd := p.Get(key)
		cmders = append(cmders, cmd)
	}
	_, err := p.Exec()
	return cmders, err
}

// ZAdd
func (r *RedisItem) ItemZAdd(redisClient *RedisDbInfo, ids []string, items ...string) error {
	key := r.getKey(items...)
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
	key := r.getKey(items...)
	result, err := redisClient.RedisDataDb.ZRange(key, 0, -1).Result()
	errors_.CheckCommonErr(err)
	return result, err
}

// SAdd
func (r *RedisItem) ItemSetSAdd(redisClient *RedisDbInfo, ids []string, items ...string) error {
	key := r.getKey(items...)
	p := redisClient.RedisDataDb.Pipeline()
	err := p.SAdd(key, ids).Err()
	errors_.CheckCommonErr(err)
	p.Expire(key, r.Expire)
	cmdSetLen := p.SCard(key)
	_, err = p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.Len {
		err = redisClient.RedisDataDb.SPopN(key, setLen-r.Len).Err()
		errors_.CheckCommonErr(err)
	}
	return err
}

func (r *RedisItem) ItemGetSAdd(redisClient *RedisDbInfo, items ...string) ([]string, error) {
	key := r.getKey(items...)
	result, err := redisClient.RedisDataDb.SMembers(key).Result()
	errors_.CheckCommonErr(err)
	return result, err
}

// connection
func (this *RedisItem) GetRedisItemFromConf(c *config.ConfigEngine, name string) {
	r := new(RedisItemYaml)
	ret := c.GetStruct(name, r).(*RedisItemYaml)
	this.Prefix = ret.Prefix
	this.Len = int64(ret.Len)
	this.Expire = time.Duration(ret.Expire) * time.Second
}
