package ioc

import (
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitLogger() logger.LoggerV1 {
	cfg := zap.NewDevelopmentConfig()
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
