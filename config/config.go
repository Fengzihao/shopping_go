package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var cfgReader *configReader

type (
	Configuration struct {
		DatabaseSettings
		JwtSettings
	}
	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}
	JwtSettings struct {
		SecretKey string
	}
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

// 实例化
func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}

}

func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	newConfigReader(configFile)
	if cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件失败：%s", err)
		return nil, err
	}
	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("解析配置文件到结构体失败: %s", err)
		return nil, err
	}
	return configuration, err
}
