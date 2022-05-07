package db

import (
	"fmt"
	"project-layout-go/internal/infrastructure/bootstrap"
	"project-layout-go/internal/infrastructure/common/config"
	"project-layout-go/pkg/mysqlx"
	"sync"

	"gorm.io/gorm"
)

var (
	// mysql多实例map
	mysqlInstancePool = sync.Map{}
)

type mysqlInstance struct {
	name   string
	writer *gorm.DB
	reader *gorm.DB
}

type mysqlInstanceCfg struct {
	Name   string
	Master *mysqlx.DBConfig
	Slave  *mysqlx.DBConfig
}

func (ins *mysqlInstance) GetName() string {
	if ins != nil {
		return ins.name
	}
	return ""
}

func (ins *mysqlInstance) GetWriter() *gorm.DB {
	if ins != nil {
		return ins.writer
	}
	return nil
}

func (ins *mysqlInstance) GetReader() *gorm.DB {
	if ins != nil {
		return ins.reader
	}
	return nil
}

func SetUp() bootstrap.BeforeServerFunc {
	return func() error {
		if err := InitMysql(); err != nil {
			return err
		}
		return nil
	}
}

// InitMysql 初始化Mysql连接，依赖先加载配置信息
func InitMysql() error {
	if config.AppCfg.IsSet("mysql2") {
		configList := make([]*mysqlInstanceCfg, 0)
		if err := config.AppCfg.UnmarshalKey("mysql", &configList); err != nil {
			return err
		}
		for _, item := range configList {
			mysqlIns := &mysqlInstance{}
			mysqlIns.name = item.Name
			writerClient, err := mysqlx.OpenDB(item.Master)
			if err != nil {
				return fmt.Errorf("open mysql error, name: %s, err: %+v", item.Name, err)
			}
			mysqlIns.writer = writerClient

			readerClient, err := mysqlx.OpenDB(item.Slave)
			if err != nil {
				return fmt.Errorf("open mysql error, name: %s, err: %+v", item.Name, err)
			}
			mysqlIns.reader = readerClient
			mysqlInstancePool.Store(item.Name, mysqlIns)
		}
	}
	return nil
}

func GetInstance(name string) *mysqlInstance {
	if obj, ok := mysqlInstancePool.Load(name); ok {
		return obj.(*mysqlInstance)
	}
	return nil
}

func Destroy() bootstrap.AfterServerFunc {
	return func() {
		walkFunc := func(key, value interface{}) bool {
			rawConnPoolW, _ := value.(*mysqlInstance).writer.DB()
			rawConnPoolR, _ := value.(*mysqlInstance).reader.DB()
			if rawConnPoolW != nil {
				_ = rawConnPoolW.Close()
			}
			if rawConnPoolR != nil {
				_ = rawConnPoolR.Close()
			}
			return true
		}
		mysqlInstancePool.Range(walkFunc)
	}

}
