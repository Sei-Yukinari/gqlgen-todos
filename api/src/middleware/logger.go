package middleware

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const loggerKey string = "logger"

func InjectLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logger.New().WithFields(logrus.Fields{})
		newCtx := context.WithValue(c.Request.Context(), loggerKey, logger)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}

func GetLogger(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey)

	if target, ok := logger.(*logrus.Entry); ok {
		return target
	} else {
		panic("cannot get logger from Context")
	}
}
