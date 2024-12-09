package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"violin-home.cn/retail/common/logs"
)

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	MC    *MongoConfig
	RC    *RedisConfig
	//GC    *GrpcServerConfig
	LC *logs.LogConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type RedisConfig struct {
	Network  string
	Addr     string
	Password string
	DB       int
}

type MongoConfig struct {
	Host string
	Port string
}

//type GrpcServerConfig struct {
//	Name string
//	Addr string
//}

var Conf *Config

func InitConfig() *Config {
	v := viper.New()
	config := &Config{viper: v}
	workDir, _ := os.Getwd()
	config.viper.SetConfigName("application")
	config.viper.SetConfigType("yaml")
	config.viper.AddConfigPath(workDir + "/config")
	err := v.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}

	config.ReadServerConfig()
	config.ReadLogsConfig()
	config.ReadMongoConfig()
	config.ReadRedisConfig()
	//config.ReadGrpcServerConfig()
	err = logs.InitConfig(config.LC)
	if err != nil {
		return nil
	}
	Conf = config
	return config
}

// ReadServerConfig 获取服务器配置
func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{
		Name: c.viper.GetString("server.name"),
		Addr: c.viper.GetString("server.addr"),
	}
	c.SC = sc
}

// ReadMongoConfig 获取mongo数据库配置
func (c *Config) ReadMongoConfig() {
	c.MC = &MongoConfig{
		Host: c.viper.GetString("mongo.host"),
		Port: c.viper.GetString("mongo.port"),
	}
}

// ReadRedisConfig 获取mongo数据库配置
func (c *Config) ReadRedisConfig() {
	c.RC = &RedisConfig{
		Network:  c.viper.GetString("redis.network"),
		Addr:     c.viper.GetString("mongo.addr"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("mongo.db"),
	}
}

//func (c *Config) ReadGrpcServerConfig() {
//	gsc := &GrpcServerConfig{
//		Name: c.viper.GetString("grpc.name"),
//		Addr: c.viper.GetString("grpc.addr"),
//	}
//	c.GC = gsc
//}

func (c *Config) ReadLogsConfig() {
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("zap.maxSize"),
		MaxAge:        c.viper.GetInt("zap.maxAge"),
		MaxBackups:    c.viper.GetInt("zap.maxBackups"),
	}
	c.LC = lc
}
