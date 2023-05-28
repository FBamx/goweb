package initialize

import (
	"fmt"
	"goweb/global"
	"goweb/utils"

	"go.uber.org/zap"
)

func InitLogger() {
	ofg := zap.NewDevelopmentConfig()
	ofg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsLocation, utils.GetNowFormatTodayTime()),
		"stdout",
	}

	logg, _ := ofg.Build()
	zap.ReplaceGlobals(logg)
	global.Lg = logg
}
