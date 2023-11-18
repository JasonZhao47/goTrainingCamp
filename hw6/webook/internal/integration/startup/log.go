package startup

import "github.com/jasonzhao47/goTrainingCamp/hw5/webook/pkg/logger"

func InitLogger() logger.LoggerV1 {
	return logger.NewNopLogger()
}
