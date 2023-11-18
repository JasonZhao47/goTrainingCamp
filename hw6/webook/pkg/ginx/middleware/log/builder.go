package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/pkg/logger"
)

type Builder struct {
	l logger.Logger
}

func NewLoggerBuilder(logger logger.Logger) *Builder {
	return &Builder{
		l: logger,
	}
}

func (b *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 递归处理后面的请求
		ctx.Next()
		// 处理完了，开始打印错误
		if len(ctx.Errors) > 0 {
			for _, e := range ctx.Errors {
				logger.Error(e)
			}
		}
	}
}
