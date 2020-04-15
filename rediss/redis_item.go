package rediss

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
	ItemPSet(redisClient redis.Cmdable, kv map[string]string) ([]*redis.StatusCmd, error)
	ItemGet(redisClient redis.Cmdable, items ...string) (*redis.StringCmd, error)

	ItemHSet(redisClient redis.Cmdable, key string, value interface{}, items ...string) error
	ItemPHSet(redisClient redis.Cmdable, field string, kv map[string]string) ([]*redis.BoolCmd, error)
	ItemPHSetField(redisClient redis.Cmdable, key string, fv map[string]string) ([]*redis.BoolCmd, error)
	ItemHGet(redisClient redis.Cmdable, key string, items ...string) (*redis.StringCmd, error)
	ItemPHGet(redisClient redis.Cmdable, field string, keys ...string) ([]*redis.StringCmd, error)
	ItemPHGetField(redisClient redis.Cmdable, keys string, field ...string) ([]*redis.StringCmd, error)

	ItemIncrExpire(redisClient redis.Cmdable, items ...string) (int, error)

	ItemZAdd(redisClient redis.Cmdable, ids []string, items ...string) error
	ItemGetZRange(redisClient redis.Cmdable, items ...string) ([]string, error)

	ItemSetSAdd(redisClient redis.Cmdable, ids []string, items ...string) error
	ItemGetSAdd(redisClient redis.Cmdable, items ...string) ([]string, error)

	SetRedisItem(prefix string, len, expire int)

	// 批量获取
	ItemPGet(redisClient redis.Cmdable, ids []string) ([]*redis.StringCmd, error)
}

type ItemInfo struct {
	prefix string
	expire time.Duration
	size   int64
}

type RedisItemYaml struct {
	Prefix string `yaml:"prefix"`
	Expire int    `yaml:"expire"`
	Len    int    `yaml:"len"`
}

// key gen
func (r *ItemInfo) ItemGetKey(keys ...string) string {
	return r.prefix + "_" + strings.Join(keys, "_")
}

// set
func (r *ItemInfo) ItemSetByte(redisClient redis.Cmdable, values []byte, keys ...string) error {
	return redisClient.Set(r.ItemGetKey(keys...), values, r.expire).Err()
}

func (r *ItemInfo) ItemSet(redisClient redis.Cmdable, value interface{}, keys ...string) error {
	return redisClient.Set(r.ItemGetKey(keys...), value, r.expire).Err()
}

func (r *ItemInfo) ItemPSet(redisClient redis.Cmdable, kv map[string]string) ([]*redis.StatusCmd, error) {
	var cmders []*redis.StatusCmd
	p := redisClient.Pipeline()
	for k, v := range kv {
		cmders = append(cmders, p.Set(r.ItemGetKey(k), v, r.expire))
	}
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	return cmders, err
}

func (r *ItemInfo) ItemGet(redisClient redis.Cmdable, keys ...string) (*redis.StringCmd, error) {
	stringCmd := redisClient.Get(r.ItemGetKey(keys...))
	return stringCmd, stringCmd.Err()
}

func (r *ItemInfo) ItemPGet(redisClient redis.Cmdable, keys []string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := redisClient.Pipeline()
	for _, k := range keys {
		cmders = append(cmders, p.Get(r.ItemGetKey(k)))
	}
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	return cmders, err
}

// hset类型
func (r *ItemInfo) ItemHSet(redisClient redis.Cmdable, field string, values interface{}, keys ...string) error {
	if err := redisClient.HSet(r.ItemGetKey(keys...), field, values).Err(); nil == err {
		return redisClient.Expire(r.ItemGetKey(keys...), r.expire).Err()
	} else {
		errors_.CheckCommonErr(err)
		return err
	}
}

func (r *ItemInfo) ItemPHSet(redisClient redis.Cmdable, field string, kv map[string]string) ([]*redis.BoolCmd, error) {
	var cmders []*redis.BoolCmd
	p := redisClient.Pipeline()
	for k, v := range kv {
		cmders = append(cmders, p.HSet(r.ItemGetKey(k), field, v))
		cmders = append(cmders, p.Expire(r.ItemGetKey(k), r.expire))
	}
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	return cmders, err
}

func (r *ItemInfo) ItemPHSetField(redisClient redis.Cmdable, key string, fv map[string]string) ([]*redis.BoolCmd, error) {
	var cmders []*redis.BoolCmd
	p := redisClient.Pipeline()
	for f, v := range fv {
		cmders = append(cmders, p.HSet(r.ItemGetKey(key), f, v))
		cmders = append(cmders, p.Expire(r.ItemGetKey(key), r.expire))
	}
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	return cmders, err
}

func (r *ItemInfo) ItemHGet(redisClient redis.Cmdable, field string, keys ...string) (*redis.StringCmd, error) {
	stringCmd := redisClient.HGet(r.ItemGetKey(keys...), field)
	return stringCmd, stringCmd.Err()
}

func (r *ItemInfo) ItemPHGet(redisClient redis.Cmdable, field string, keys ...string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := redisClient.Pipeline()
	for _, k := range keys {
		cmders = append(cmders, p.HGet(r.ItemGetKey(k), field))
	}
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	return cmders, err
}

func (r *ItemInfo) ItemPHGetField(redisClient redis.Cmdable, keys string, field ...string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := redisClient.Pipeline()
	for _, f := range field {
		cmders = append(cmders, p.HGet(r.ItemGetKey(keys), f))
	}
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	return cmders, err
}

// Incr
func (r *ItemInfo) ItemIncrExpire(redisClient redis.Cmdable, keys ...string) (int, error) {
	key := r.ItemGetKey(keys...)
	p := redisClient.Pipeline()
	cmder := p.Incr(key)
	p.Expire(key, r.expire)
	_, err := p.Exec()
	errors_.CheckCommonErr(err)
	val, err := cmder.Result()
	errors_.CheckCommonErr(err)
	return int(val), err
}

// ZAdd
func (r *ItemInfo) ItemZAdd(redisClient redis.Cmdable, values []string, keys ...string) error {
	key := r.ItemGetKey(keys...)
	zmembers := make([]redis.Z, 0, len(values))
	for _, id := range values {
		zmembers = append(zmembers, redis.Z{Score: float64(time.Now().UnixNano()), Member: id})
	}
	p := redisClient.Pipeline()
	err := p.ZAdd(key, zmembers...).Err()
	errors_.CheckCommonErr(err)
	cmdSetLen := p.ZCard(key)
	_, err = p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.size {
		err := redisClient.ZRemRangeByRank(key, 0, setLen-r.size).Err()
		errors_.CheckCommonErr(err)
	}
	return err
}

func (r *ItemInfo) ItemGetZRange(redisClient redis.Cmdable, keys ...string) ([]string, error) {
	key := r.ItemGetKey(keys...)
	result, err := redisClient.ZRange(key, 0, -1).Result()
	errors_.CheckCommonErr(err)
	return result, err
}

// SAdd
func (r *ItemInfo) ItemSetSAdd(redisClient redis.Cmdable, values []string, keys ...string) error {
	key := r.ItemGetKey(keys...)
	p := redisClient.Pipeline()
	err := p.SAdd(key, values).Err()
	errors_.CheckCommonErr(err)
	p.Expire(key, r.expire)
	cmdSetLen := p.SCard(key)
	_, err = p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.size {
		err = redisClient.SPopN(key, setLen-r.size).Err()
		errors_.CheckCommonErr(err)
	}
	return err
}

func (r *ItemInfo) ItemGetSAdd(redisClient redis.Cmdable, keys ...string) ([]string, error) {
	result, err := redisClient.SMembers(r.ItemGetKey(keys...)).Result()
	errors_.CheckCommonErr(err)
	return result, err
}

func (this *ItemInfo) SetRedisItem(prefix string, len, expire int) {
	this.prefix = prefix
	this.size = int64(len)
	this.expire = time.Duration(expire) * time.Second
}

// connection
func (this *ItemInfo) GetRedisItemFromConf(c *config.ConfigEngine, name string) {
	r := new(RedisItemYaml)
	ret := c.GetStruct(name, r).(*RedisItemYaml)
	this.prefix = ret.Prefix
	this.size = int64(ret.Len)
	this.expire = time.Duration(ret.Expire) * time.Second
	errors_.CheckEmptyValue(this.prefix)
}
