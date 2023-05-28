package global

import (
	"go.uber.org/zap"
	"goweb/config"
)

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
)
