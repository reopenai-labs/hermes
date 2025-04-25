package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"log"
	"os"
	"qiyibyte.com/hermes/pkg/configs"
	"qiyibyte.com/hermes/pkg/logger"
	"runtime"
	"sync"
	"time"
)

func init() {
	configs.SetDefault("application.datasource.host", "127.0.0.1")
	configs.SetDefault("application.datasource.port", "5432")
	configs.SetDefault("application.datasource.username", "postgres")
	configs.SetDefault("application.datasource.password", "postgres")
	configs.SetDefault("application.datasource.pool.max-idle-conns", runtime.NumCPU())
	configs.SetDefault("application.datasource.pool.max-open-conns", runtime.NumCPU())
	configs.SetDefault("application.datasource.pool.conn-max-lifetime", 30*time.Minute)
}

var instance *gorm.DB
var initializer sync.Once

func Instance() *gorm.DB {
	initializer.Do(func() {
		config := &gorm.Config{
			Logger: gormlog.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				gormlog.Config{
					LogLevel:                  gormlog.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			),
		}
		host := configs.GetString("application.datasource.host")
		port := configs.GetString("application.datasource.port")
		username := configs.GetString("application.datasource.username")
		password := configs.GetString("application.datasource.password")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=anbernic sslmode=disable TimeZone=UTC", host, port, username, password)
		db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), config)
		if err != nil {
			logger.Fatalf("GORM: 连接数据库失败.err=%s", err.Error())
		}
		sqlDB, err := db.DB()
		if err != nil {
			logger.Fatalf("GORM: 获取数据库连接失败.err=%s", err.Error())
		}
		sqlDB.SetMaxIdleConns(configs.GetInt("application.datasource.pool.max-idle-conns"))
		sqlDB.SetMaxOpenConns(configs.GetInt("application.datasource.pool.max-open-conns"))
		sqlDB.SetConnMaxLifetime(configs.GetDuration("application.datasource.pool.conn-max-lifetime"))
		instance = db
		logger.Infof("Postgres: 初始化数据库成功.host=%s", host)
	})
	return instance
}
