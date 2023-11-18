//go:build wireinject

package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository/cache"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository/dao"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/service"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/service/sms"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/service/sms/async"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/web"
	ijwt "github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/web/jwt"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/ioc"
)

var thirdPartySet = wire.NewSet( // 第三方依赖
	InitRedis, InitDB,
	InitLogger)

var userSvcProvider = wire.NewSet(
	dao.NewUserDAO,
	cache.NewUserCache,
	repository.NewCachedUserRepository,
	service.NewUserService)

var articlSvcProvider = wire.NewSet(
	repository.NewCachedArticleRepository,
	cache.NewArticleRedisCache,
	dao.NewArticleGORMDAO,
	service.NewArticleService)

func InitWebServer() *gin.Engine {
	wire.Build(
		thirdPartySet,
		userSvcProvider,
		articlSvcProvider,
		// cache 部分
		cache.NewCodeCache,

		// repository 部分
		repository.NewCodeRepository,

		// Service 部分
		ioc.InitSMSService,
		service.NewCodeService,
		InitWechatService,

		// handler 部分
		web.NewUserHandler,
		web.NewArticleHandler,
		web.NewOAuth2WechatHandler,
		ijwt.NewRedisJWTHandler,
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}

func InitAsyncSmsService(svc sms.Service) *async.Service {
	wire.Build(thirdPartySet, repository.NewAsyncSMSRepository,
		dao.NewGORMAsyncSmsDAO,
		async.NewService,
	)
	return &async.Service{}
}

func InitArticleHandler(dao dao.ArticleDAO) *web.ArticleHandler {
	wire.Build(
		thirdPartySet,
		userSvcProvider,
		repository.NewCachedArticleRepository,
		cache.NewArticleRedisCache,
		service.NewArticleService,
		web.NewArticleHandler)
	return &web.ArticleHandler{}
}
