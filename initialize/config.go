package initialize

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"goweb/config"
	"goweb/global"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	// 配置文件位置
	v.SetConfigFile("./settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.Settings = serverConfig
	color.Blue("goweb start...")
}
