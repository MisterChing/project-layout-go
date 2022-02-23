package mysqlx

import (
    "database/sql"
    "errors"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// DBConfig 数据库连接配置
type DBConfig struct {
    Dsn             string `mapstructure:"dsn"`             //Mysql连接DSN，比如：root:123456@tcp(127.0.0.1:3306)/project-demo?charset=utf8&parseTime=true
    MaxIdleConns    uint   `mapstructure:"maxidleconns"`    //最大空闲连接数
    MaxOpenConns    uint   `mapstructure:"maxopenconns"`    //最大连接数
    ConnMaxLifetime uint   `mapstructure:"connmaxlifetime"` //连接最长保持时间
}

// OpenDB 打开数据库连接
// @param dbConf *DBConfig 数据库连接配置
// @return *gorm.DB gorm.DB连接DB
// @return error 错误信息
func OpenDB(dbConf *DBConfig) (*gorm.DB, error) {
    dialer := mysql.New(mysql.Config{
        DSN: dbConf.Dsn,
    })
    gormDB, err := gorm.Open(dialer, &gorm.Config{})
    if err != nil {
        return nil, err
    }
    if gormDB == nil {
        return nil, errors.New("DB object is nil")
    }
    var rawDB *sql.DB
    rawDB, err = gormDB.DB()
    if err != nil {
        return nil, err
    }
    if err = rawDB.Ping(); err != nil {
        return nil, err
    }
    rawDB.SetMaxIdleConns(int(dbConf.MaxIdleConns))
    rawDB.SetMaxOpenConns(int(dbConf.MaxOpenConns))
    rawDB.SetConnMaxLifetime(time.Duration(dbConf.ConnMaxLifetime) * time.Second)
    return gormDB, nil
}
