package redisx

import (
    "context"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

// RedisConfig redis连接配置
type RedisConfig struct {
    Host           string `mapstructure:"host"`           //Redis服务器Host
    Port           string `mapstructure:"port"`           //Redis服务器Port
    Password       string `mapstructure:"password"`       //连接密码
    Database       int    `mapstructure:"database"`       //数据库，默认为0
    MaxConns       int    `mapstructure:"maxconns"`       //Redis连接池最大连接数
    MinIdle        int    `mapstructure:"minidle"`        //Redis连接池最小空闲连接数
    ConnTimeoutMs  int    `mapstructure:"conntimeoutms"`  //Redis连接超时(毫秒)
    ReadTimeoutMs  int    `mapstructure:"readtimeoutms"`  //Redis读超时(毫秒)
    WriteTimeoutMs int    `mapstructure:"writetimeoutms"` //Redis写超时(毫秒)
}

// OpenRedis 打开redis连接
// @param redisConf *RedisConfig redis连接配置
// @return *redis.Client go-redis连接client
// @return error 错误信息
func OpenRedis(redisConf *RedisConfig) (*redis.Client, error) {
    option := &redis.Options{
        Addr:         redisConf.Host + ":" + redisConf.Port,
        Password:     redisConf.Password,
        DB:           redisConf.Database,
        PoolSize:     redisConf.MaxConns,
        MinIdleConns: redisConf.MinIdle,
        DialTimeout:  time.Duration(redisConf.ConnTimeoutMs) * time.Millisecond,
        ReadTimeout:  time.Duration(redisConf.ReadTimeoutMs) * time.Millisecond,
        WriteTimeout: time.Duration(redisConf.WriteTimeoutMs) * time.Millisecond,
    }
    client := redis.NewClient(option)
    ret, err := client.Ping(context.Background()).Result()
    if err != nil || ret != "PONG" {
        return nil, fmt.Errorf("connect redis failed. error:%+v", err)
    }
    return client, nil
}
