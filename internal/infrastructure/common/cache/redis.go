package cache

import (
	"errors"
	"fmt"
	"project-layout-go/internal/infrastructure/bootstrap"
	"project-layout-go/internal/infrastructure/common/config"
	"project-layout-go/pkg/redisx"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	// redis多实例map
	redisInstancePool = sync.Map{}
)

type redisInstance struct {
	name   string
	writer *redis.Client
	reader *redis.Client
}

type redisInstanceCfg struct {
	Name   string
	Master *redisx.RedisConfig
	Slave  *redisx.RedisConfig
}

func (ins *redisInstance) GetName() string {
	if ins != nil {
		return ins.name
	}
	return ""
}

func (ins *redisInstance) GetWriter() *redis.Client {
	if ins != nil {
		return ins.writer
	}
	return nil
}

func (ins *redisInstance) GetReader() *redis.Client {
	if ins != nil {
		return ins.reader
	}
	return nil
}

func SetUp() bootstrap.BeforeServerFunc {
	return func() error {
		if err := InitRedis(); err != nil {
			return err
		}
		return nil
	}
}

// InitRedis 初始化Redis连接，依赖先加载配置信息
func InitRedis() error {
	if config.AppCfg.IsSet("redis") {
		configList := make([]*redisInstanceCfg, 0)
		if err := config.AppCfg.UnmarshalKey("redis", &configList); err != nil {
			return err
		}
		for _, item := range configList {
			redisIns := &redisInstance{}
			redisIns.name = item.Name
			writerClient, err := redisx.OpenRedis(item.Master)

			if err != nil {
				return fmt.Errorf("open redis error, name: %s, err: %+v", item.Name, err)
			}
			redisIns.writer = writerClient

			readerClient, err := redisx.OpenRedis(item.Slave)
			if err != nil {
				return fmt.Errorf("open redis error, name: %s, err: %+v", item.Name, err)
			}
			redisIns.reader = readerClient
			redisInstancePool.Store(item.Name, redisIns)
		}
	}

	return nil
}

func GetInstance(name string) *redisInstance {
	if obj, ok := redisInstancePool.Load(name); ok {
		return obj.(*redisInstance)
	}
	return nil
}

func IsKeyNotExists(err error) bool {
	return errors.Is(redis.Nil, err)
}

func Destroy() bootstrap.AfterServerFunc {
	return func() {
		walkFunc := func(key, value interface{}) bool {
			rawConnW := value.(*redisInstance).writer
			rawConnR := value.(*redisInstance).reader
			if rawConnW != nil {
				_ = rawConnW.Close()
			}
			if rawConnR != nil {
				_ = rawConnR.Close()
			}
			return true
		}
		redisInstancePool.Range(walkFunc)
	}
}
