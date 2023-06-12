package main

import (
	"fmt"
	"goweb/global"
	"goweb/initialize"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()

	Router := initialize.Routers()
	color.Cyan("goweb start...")
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误"))
	}
}
