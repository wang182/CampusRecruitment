package config

import (
	"CampusRecruitment/pkg/utils"
	"fmt"

	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Listen    string      `mapstructure:"listen" yaml:"listen"`
	SecretKey string      `mapstructure:"secretKey" yaml:"secretKey"`
	Storage   string      `mapstructure:"storage" yaml:"storage"`
	Mysql     MysqlConfig `mapstructure:"mysql" yaml:"mysql"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	DB       string `mapstructure:"db" yaml:"db"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}

var (
	loadConfigLock sync.RWMutex
	gConfig        = Config{}
)

func Default() Config {
	return Config{
		Listen:    "0.0.0.0:8080",
		SecretKey: "",
		Mysql: MysqlConfig{
			Host:     "localhost",
			Port:     3306,
			DB:       "db_name",
			Username: "root",
		},
	}
}

func init() {
	gConfig = Default()
}

func Load(filename string) error {
	loadConfigLock.Lock()
	defer loadConfigLock.Unlock()

	viper.SetConfigFile(filename)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&gConfig); err != nil {
		return err
	}

	// check jwt secret key
	if gConfig.SecretKey == "" {
		return fmt.Errorf("jwt secret key is not configed")
	}
	// 总是保证密钥长度为 32 位，如果传入的不是 32 位密钥则进行一次 md5
	if len(gConfig.SecretKey) != 32 {
		gConfig.SecretKey = utils.Md5([]byte(gConfig.SecretKey))
	}
	return nil
}

func Get() Config {
	loadConfigLock.RLock()
	defer loadConfigLock.RUnlock()
	return gConfig
}
