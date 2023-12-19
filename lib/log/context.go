package log

import (
	"context"
	"log/slog"
)

var loggerKey = struct{ name string }{
	name: "logger",
}

func NewContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	if value, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return value
	}

	return nil
}
