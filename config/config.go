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
	//GC    *GrpcServerConfig
	LC *logs.LogConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

//type GrpcServerConfig struct {
//	Name string
//	Addr string
//}

var Conf = InitConfig()

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
	//config.ReadGrpcServerConfig()
	logs.InitConfig(config.LC)
	return config
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{
		Name: c.viper.GetString("server.name"),
		Addr: c.viper.GetString("server.addr"),
	}
	c.SC = sc
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
