//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository/cache"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository/dao"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/service"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/web"
	ijwt "github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/web/jwt"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		ioc.InitRedis, ioc.InitDB,
		ioc.InitLogger,
		// DAO 部分
		dao.NewUserDAO,
		dao.NewArticleGORMDAO,

		// cache 部分
		cache.NewCodeCache, cache.NewUserCache,
		cache.NewArticleRedisCache,

		// repository 部分
		repository.NewCachedUserRepository,
		repository.NewCodeRepository,
		repository.NewCachedArticleRepository,

		// Service 部分
		ioc.InitSMSService,
		ioc.InitWechatService,
		service.NewUserService,
		service.NewCodeService,
		service.NewArticleService,

		// handler 部分
		web.NewUserHandler,
		web.NewArticleHandler,
		ijwt.NewRedisJWTHandler,
		web.NewOAuth2WechatHandler,
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
