package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AppName        string
	MySQL          MySQLConfig
	Log            LogConfig
	StaticPath     PathConfig
	MsgChannelType MsgChannelType
	Redis          RedisConfig
	ElasticSearch  ElasticSearchConfig
}

// MySQL相关配置
type MySQLConfig struct {
	Name         string
	Host         string
	Password     string
	Port         int
	TablePrefix  string
	User         string
	Loc          string
	Charset      string
	Time         int
	MaxOpenConns int
	MaxIdelConns int
}

// 日志保存地址
type LogConfig struct {
	Path  string
	Level string
}

// 相关地址信息，例如静态文件地址
type PathConfig struct {
	FilePath string
}

type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int
	PoolTimeout  int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	Prefix       redisPrefix
	Expire       redisExpire
}
type redisExpire struct {
	AuthCode int
}

type redisPrefix struct {
	AuthCode string
}

// 消息队列类型及其消息队列相关信息
// gochannel为单机使用go默认的channel进行消息传递
// kafka是使用kafka作为消息队列，可以分布式扩展消息聊天程序
type MsgChannelType struct {
	ChannelType string
	KafkaHosts  string
	KafkaTopic  string
}

type ElasticSearchConfig struct {
	Repositories struct {
		Enabled bool
	}
	ClusterNodes string
	ClusterName  string
}

var c Config

func init() {

	//workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	//viper.AddConfigPath(workDir+"\\global\\config")
	viper.AddConfigPath("G:\\GoProject\\mall\\global\\config")
	err := viper.ReadInConfig()

	if err != nil {
		logrus.Error("An error occurred while reading the configuration file.")
	}
	viper.Unmarshal(&c)

}

func GetConfig() Config {
	return c
}
