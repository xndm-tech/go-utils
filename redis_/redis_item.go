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

	ItemSetByte(redisClient redis.Cmdable, bytes []byte, items ...string) error
	ItemSet(redisClient redis.Cmdable, value interface{}, items ...string) error
	ItemGet(redisClient redis.Cmdable, items ...string) (*redis.StringCmd, error)

	ItemHSet(redisClient redis.Cmdable, key string, value interface{}, items ...string) error
	ItemHGet(redisClient redis.Cmdable, key string, items ...string) (*redis.StringCmd, error)

	ItemIncrExpire(redisClient redis.Cmdable, items ...string) (int, error)

	ItemZAdd(redisClient redis.Cmdable, ids []string, items ...string) error
	ItemGetZRange(redisClient redis.Cmdable, items ...string) ([]string, error)

	ItemSetSAdd(redisClient redis.Cmdable, ids []string, items ...string) error
	ItemGetSAdd(redisClient redis.Cmdable, items ...string) ([]string, error)

	// 批量获取
	ItemPGet(redisClient redis.Cmdable, ids []string) ([]*redis.StringCmd, error)
}

type RedisItem struct {
	Prefix string
	Expire time.Duration
	Len    int64
}

type RedisItemYaml struct {
	Prefix string `yaml:"prefix"`
	Expire int    `yaml:"expire"`
	Len    int    `yaml:"len"`
}

func (r *RedisItem) ItemGetKey(items ...string) string {
	return r.Prefix + "_" + strings.Join(items, "_")
}

// map
func (r *RedisItem) ItemSetByte(redisClient redis.Cmdable, bytes []byte, items ...string) error {
	return redisClient.Set(r.ItemGetKey(items...), bytes, r.Expire).Err()
}

func (r *RedisItem) ItemSet(redisClient redis.Cmdable, value interface{}, items ...string) error {
	return redisClient.Set(r.ItemGetKey(items...), value, r.Expire).Err()
}

func (r *RedisItem) ItemGet(redisClient redis.Cmdable, items ...string) (*redis.StringCmd, error) {
	stringCmd := redisClient.Get(r.ItemGetKey(items...))
	return stringCmd, stringCmd.Err()
}

func (r *RedisItem) ItemHSet(redisClient redis.Cmdable, field string, value interface{}, items ...string) error {
	return redisClient.HSet(r.ItemGetKey(items...), field, value).Err()
}

func (r *RedisItem) ItemHGet(redisClient redis.Cmdable, field string, items ...string) (*redis.StringCmd, error) {
	stringCmd := redisClient.HGet(r.ItemGetKey(items...), field)
	return stringCmd, stringCmd.Err()
}

// Incr
func (r *RedisItem) ItemIncrExpire(redisClient redis.Cmdable, items ...string) (int, error) {
	key := r.ItemGetKey(items...)
	p := redisClient.Pipeline()
	cmder := p.Incr(key)
	p.Expire(key, r.Expire)
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	val, err := cmder.Result()
	errors_.CheckCommonErr(err)
	return int(val), err
}

// 批量获取
func (r *RedisItem) ItemPGet(redisClient redis.Cmdable, ids []string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := redisClient.Pipeline()
	for _, id := range ids {
		key := r.ItemGetKey(id)
		cmd := p.Get(key)
		cmders = append(cmders, cmd)
	}
	_, err := p.Exec()
	return cmders, err
}

// ZAdd
func (r *RedisItem) ItemZAdd(redisClient redis.Cmdable, ids []string, items ...string) error {
	key := r.ItemGetKey(items...)
	zmembers := make([]redis.Z, 0, len(ids))
	for _, id := range ids {
		zmembers = append(zmembers, redis.Z{Score: float64(time.Now().UnixNano()), Member: id})
	}
	p := redisClient.Pipeline()
	err := p.ZAdd(key, zmembers...).Err()
	errors_.CheckCommonErr(err)
	cmdSetLen := p.ZCard(key)
	_, err = p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.Len {
		err := redisClient.ZRemRangeByRank(key, 0, setLen-r.Len).Err()
		errors_.CheckCommonErr(err)
	}
	return err
}

func (r *RedisItem) ItemGetZRange(redisClient redis.Cmdable, items ...string) ([]string, error) {
	key := r.ItemGetKey(items...)
	result, err := redisClient.ZRange(key, 0, -1).Result()
	errors_.CheckCommonErr(err)
	return result, err
}

// SAdd
func (r *RedisItem) ItemSetSAdd(redisClient redis.Cmdable, ids []string, items ...string) error {
	key := r.ItemGetKey(items...)
	p := redisClient.Pipeline()
	err := p.SAdd(key, ids).Err()
	errors_.CheckCommonErr(err)
	p.Expire(key, r.Expire)
	cmdSetLen := p.SCard(key)
	_, err = p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.Len {
		err = redisClient.SPopN(key, setLen-r.Len).Err()
		errors_.CheckCommonErr(err)
	}
	return err
}

func (r *RedisItem) ItemGetSAdd(redisClient redis.Cmdable, items ...string) ([]string, error) {
	key := r.ItemGetKey(items...)
	result, err := redisClient.SMembers(key).Result()
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
	errors_.CheckEmptyValue(this.Prefix)
}
