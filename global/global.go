package global

import (
	"github.com/dalefeng/chat-api-reverse/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Config   config.Server
	Viper    *viper.Viper
	Log      *zap.Logger
	SugarLog *zap.SugaredLogger
)
