package configs

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"qiyibyte.com/hermes/pkg/builtin/eventbus"
	"time"
)

func init() {
	// 允许从环境变量中加载配置
	viper.AutomaticEnv()
	// 如果环境变量中存在key，也将覆盖配置
	viper.AllowEmptyEnv(true)
	// 从flag中加载配置
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	eventbus.Publish(eventbus.TopicConfigReady)
}

// SetDefault 给配置设置一个默认值。
// 此默认值的优先级最低，如果通过其他各种渠道都无法获取到配置信息时，才会使用此默认值
// value 还支持map结构的参数。
func SetDefault(key string, value any) {
	viper.SetDefault(key, value)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetUint8(key string) uint8 {
	return viper.GetUint8(key)
}

func GetUint(key string) uint {
	return viper.GetUint(key)
}

func GetUint16(key string) uint16 {
	return viper.GetUint16(key)
}

func GetUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func GetUint64(key string) uint64 {
	return viper.GetUint64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetStringMap(key string) map[string]any {
	return viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

func GetSizeInBytes(key string) uint {
	return viper.GetSizeInBytes(key)
}

func Unmarshal(p any) error {
	return viper.Unmarshal(p)
}

func UnmarshalKey(key string, rawVal any) error {
	return viper.UnmarshalKey(key, rawVal)
}
