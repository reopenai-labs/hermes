package redisx

import (
	"github.com/redis/go-redis/v9"
	"qiyibyte.com/hermes/pkg/configs"
	"qiyibyte.com/hermes/pkg/logger"
	"time"
)

var instance redis.Cmdable
var hashCommand *HashCommand
var listCommand *ListCommand
var stringCommand *StringCommand
var pubSubCommand *PubSubCommand
var setCommand *SetCommand
var zSetCommand *ZSetCommand
var pipelineCommand *PipelineCommand

func init() {
	configs.SetDefault("application.redis", "127.0.0.1:6379")
	configs.SetDefault("application.read-timeout", 5*time.Second)
	configs.SetDefault("application.write-timeout", 5*time.Second)
	initCommands()
}

func initCommands() {
	var config RedisConfig
	if err := configs.UnmarshalKey("application.redis", &config); err != nil {
		logger.Fatalf("redis: 配置文件解析失败: %v", err)
	}
	if len(config.Nodes) > 0 {
		client := NewCluster(config)
		setCommand = &SetCommand{client: client}
		zSetCommand = &ZSetCommand{client: client}
		hashCommand = &HashCommand{client: client}
		listCommand = &ListCommand{client: client}
		stringCommand = &StringCommand{client: client}
		pubSubCommand = &PubSubCommand{client: client}
		pipelineCommand = &PipelineCommand{client: client.Pipeline()}
		instance = client
	} else {
		client := New(config)
		setCommand = &SetCommand{client: client}
		zSetCommand = &ZSetCommand{client: client}
		hashCommand = &HashCommand{client: client}
		listCommand = &ListCommand{client: client}
		stringCommand = &StringCommand{client: client}
		pubSubCommand = &PubSubCommand{client: client}
		pipelineCommand = &PipelineCommand{client: client.Pipeline()}
		instance = client
	}
}

func Hash() *HashCommand {
	return hashCommand
}

func List() *ListCommand {
	return listCommand
}

func String() *StringCommand {
	return stringCommand
}

func Set() *SetCommand {
	return setCommand
}

func ZSet() *ZSetCommand {
	return zSetCommand
}

func PubSub() *PubSubCommand {
	return pubSubCommand
}

func Pipeline() *PipelineCommand {
	return pipelineCommand
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
