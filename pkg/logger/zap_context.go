package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const loggerKey = iota

var Logger *zap.Logger

// Default initializes the logger with the default configuration.
func Default() *zap.Logger {
	config := zap.NewProductionConfig()
	l, err := config.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		panic(err)
	}
	Logger = l
	return Logger
}

// WithLogger returns a new context with the provided logger.
func WithLogger(ctx context.Context, fields ...zapcore.Field) context.Context {
	return context.WithValue(ctx, loggerKey, WithContext(ctx).With(fields...))
}

// WithContext returns a new logger with the provided context.
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}

	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}

	return Logger
}
