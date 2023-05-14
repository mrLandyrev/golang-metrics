package rest

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BuildLoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		end := time.Since(start)

		logger.Info(
			"Request info",
			zap.String("URI", context.Request.RequestURI),
			zap.String("Method", context.Request.Method),
			zap.Duration("Time", end),
		)

		logger.Info(
			"Response info",
			zap.Int("Status", context.Writer.Status()),
			zap.Int("Size", context.Writer.Size()),
		)
	}
}
