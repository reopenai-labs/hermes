package redisx

import (
	"github.com/redis/go-redis/v9"
	"qiyibyte.com/hermes/pkg/configs"
	"qiyibyte.com/hermes/pkg/logger"
	"sync"
	"time"
)

var instance redis.Cmdable
var instanceInit sync.Once

func init() {
	configs.SetDefault("application.redis", "127.0.0.1:6379")
	configs.SetDefault("application.read-timeout", 5*time.Second)
	configs.SetDefault("application.write-timeout", 5*time.Second)
}

func New(config RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
}

func NewCluster(config RedisConfig) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    config.Nodes,
		Username: config.Username,
		Password: config.Password,
		PoolSize: config.PoolSize,
	})
}

type RedisConfig struct {
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	DB              int           `mapstructure:"db"`
	Addr            string        `mapstructure:"addr"`
	Nodes           []string      `mapstructure:"nodes"`
	PoolSize        int           `mapstructure:"pool-size"`
	ReadTimeout     time.Duration `mapstructure:"read-timeout"`
	WriteTimeout    time.Duration `mapstructure:"write-timeout"`
	MinIdleConns    int           `mapstructure:"min-idle-conns"`
	MaxIdleConns    int           `mapstructure:"max-idle-conns"`
	MaxActiveConns  int           `mapstructure:"max-active-conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn-max-idle-time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn-max-lifetime"`
}

func getInstance() redis.Cmdable {
	instanceInit.Do(func() {
		var config RedisConfig
		if err := configs.UnmarshalKey("application.redis", &config); err != nil {
			logger.Fatalf("redis: 配置文件解析失败: %v", err)
		}
		if len(config.Nodes) > 0 {
			instance = NewCluster(config)
			logger.Info("Redis Cluster: 实例初始化成功")
		} else {
			instance = New(config)
			logger.Infof("Redis: 实例初始化成功.host=%s,db=%d", config.Addr, config.DB)
		}
	})
	return instance
}
