package adapters

import (
	"context"
	"io"
	"log/slog"
)

// This implements the interfaces.Logger interface with the slog package.
// Other packages may implement this interface with other external packages.

type slogLogger struct {
	slog *slog.Logger
}

func NewSlogLogger(w io.Writer) *slogLogger {
	l := slogLogger{}

	jsonHandler := slog.NewTextHandler(w, &slog.HandlerOptions{})
	l.slog = slog.New(jsonHandler)
	return &l
}

func (l *slogLogger) Debug(ctx context.Context, msg string, args ...any) {
	l.slog.DebugContext(ctx, msg, enrichArgsWithContext(ctx, args)...)
}

func (l *slogLogger) Info(ctx context.Context, msg string, args ...any) {
	l.slog.InfoContext(ctx, msg, enrichArgsWithContext(ctx, args)...)
}

func (l *slogLogger) Warn(ctx context.Context, msg string, args ...any) {
	l.slog.WarnContext(ctx, msg, enrichArgsWithContext(ctx, args)...)
}

func (l *slogLogger) Error(ctx context.Context, msg string, args ...any) {
	l.slog.ErrorContext(ctx, msg, enrichArgsWithContext(ctx, args)...)
}

func enrichArgsWithContext(ctx context.Context, args []any) []any {
	REQUEST_ID_CTX_KEY := "request_id"
	USER_ID_CTX_KEY := "user_id"

	requestID := ctx.Value(REQUEST_ID_CTX_KEY)
	if requestID != nil {
		args = append(args, REQUEST_ID_CTX_KEY, requestID)
	}

	userID := ctx.Value(USER_ID_CTX_KEY)
	if userID != nil {
		args = append(args, USER_ID_CTX_KEY, userID)
	}

	return args
}
