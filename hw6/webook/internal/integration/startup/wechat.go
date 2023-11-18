package startup

import (
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/service/oauth2/wechat"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/pkg/logger"
)

func InitWechatService(l logger.LoggerV1) wechat.Service {
	return wechat.NewService("", "", l)
}
