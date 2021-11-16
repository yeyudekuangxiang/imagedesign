package db

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

var LogLevelMap = map[string]logger.LogLevel{
	"Silent": logger.Silent,
	"Error":  logger.Error,
	"Warn":   logger.Warn,
	"Info":   logger.Info,
}

func getLogLevel(level string) logger.LogLevel {
	logLevel, ok := LogLevelMap[level]
	if !ok {
		logLevel = logger.Error
	}
	return logLevel
}

type Config struct {
	Type         string
	Host         string
	UserName     string
	Password     string
	Database     string
	TablePrefix  string
	MaxOpenConns int    //最大连接数 <=0表示不限制连接数
	MaxIdleConns int    //最大空闲数 <=0表示不保留空闲连接
	MaxLifetime  int    //连接可重用时间 <=0表示永远可用(单位秒)
	LogLevel     string //控制台日志等级 Silent Error Warn Info
}

func NewDB(conf Config) (*gorm.DB, error) {
	switch strings.ToLower(conf.Type) {
	case "mysql":
		return NewMysqlDB(conf)
	default:
		return NewMysqlDB(conf)
	}
}

//创建Mysql数据库链接
func NewMysqlDB(conf Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,             // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		return nil, errors.Wrap(err, "open db failed")
	}
	sqlDb, err := db.DB()
	if err == nil {
		sqlDb.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDb.SetMaxOpenConns(conf.MaxOpenConns)
		sqlDb.SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Second)
	}
	//配置日志打印
	db.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 3 * time.Second,
		LogLevel:      getLogLevel(conf.LogLevel),
	})
	return db, nil
}
