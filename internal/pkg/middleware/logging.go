package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/ryanpenn/go-realworld/pkg/logger"
)

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// traceId for every request
		traceId := uuid.NewString()
		logger.WithLogger(ctx, zap.String("traceId", traceId))

		// log request info
		logger.WithLogger(ctx, zap.String("request.method", ctx.Request.Method))
		headers, _ := json.Marshal(ctx.Request.Header)
		logger.WithLogger(ctx, zap.String("request.headers", string(headers)))
		logger.WithLogger(ctx, zap.String("request.url", ctx.Request.URL.String()))

		// log request params
		if ctx.Request.Form == nil {
			ctx.Request.ParseMultipartForm(32 << 20)
		}
		form, _ := json.Marshal(ctx.Request.Form)
		logger.WithLogger(ctx, zap.String("request.params", string(form)))

		ctx.Next()
	}
}
