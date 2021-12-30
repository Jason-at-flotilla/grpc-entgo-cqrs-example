package redis

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"poolSize"`
}

const (
	maxRetries      = 5
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 500 * time.Millisecond
	dialTimeout     = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 3 * time.Second
	minIdleConns    = 20
	poolTimeout     = 6 * time.Second
	idleTimeout     = 12 * time.Second
)

type CacheType int

const (
	CACHE_LOGIN_SESSION CacheType = 0
	DB_2                CacheType = 1
	DB_3                CacheType = 2
	DB_CACHE            CacheType = 3
)

func (x CacheType) Number() protoreflect.EnumNumber {
	return (protoreflect.EnumNumber(x))
}

var (
	CacheType_name = map[int]string{
		0: "CACHE_LOGIN_SESSION",
		1: "DB_2",
		2: "DB_3",
		3: "DB_CACHE",
	}
	CacheType_value = map[string]int32{
		"CACHE_LOGIN_SESSION": 0,
		"DB_2":                1,
		"DB_3":                2,
		"DB_CACHE":            3,
	}
)

type MyRedis struct {
	client *redis.Client
}
type RedisInterFace interface {
	// TODO Review: value的型別是不是用string比較適合？
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
}

var redisClient map[int]*MyRedis
var port int
var host string
var cfg Config

func Init(cfg *Config) {
	cfg = &Config{
		Addr:     cfg.Addr,
		Password: "",
		PoolSize: cfg.PoolSize,
	}
}

func GetClient(db int) RedisInterFace {
	if redisClient == nil {
		redisClient = make(map[int]*MyRedis)
	}
	if redisClient[db] == nil {
		redisClient[db] = &MyRedis{}
		redisClient[db].client = redis.NewClient(&redis.Options{
			Addr:            cfg.Addr,
			Password:        "",
			DB:              db,
			MaxRetries:      maxRetries,
			MinRetryBackoff: minRetryBackoff,
			MaxRetryBackoff: maxRetryBackoff,
			DialTimeout:     dialTimeout,
			ReadTimeout:     readTimeout,
			WriteTimeout:    writeTimeout,
			PoolSize:        cfg.PoolSize,
			MinIdleConns:    minIdleConns,
			PoolTimeout:     poolTimeout,
			IdleTimeout:     idleTimeout,
		})
	}
	return redisClient[db]
}

func (r *MyRedis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.client.Context(), key, value, expiration).Err()
}

func (r *MyRedis) Get(key string) (string, error) {
	bb, err := r.client.Get(r.client.Context(), key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			fmt.Printf("%v", err)
		}
		return "", err
	}
	return bb, nil
}

func (r *MyRedis) Del(key string) error {
	err := r.client.Del(r.client.Context(), key).Err()
	if err != nil {
		return err
	}
	return nil
}
