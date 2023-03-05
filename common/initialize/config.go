package initialize

import (
	"fmt"
	"go-blog/config"

	"github.com/spf13/viper"
)

// github.com/spf13/viper
func InitAppConfig() *config.AppConfig {
	viper.SetConfigFile("appconfig.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name

	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置出现错误:", err.Error())
	}

	var _config *config.AppConfig
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println("配置文件解析出现错误:", err.Error())
	}
	return _config
}
