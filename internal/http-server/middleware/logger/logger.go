package logger

import (
	"github.com/gwkeo/url-shortener/internal/http-server/middleware/reqID"
	"log/slog"
	"net/http"
	"time"
)

func NewLoggerMW(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log = log.With(
			slog.String("component", "middleware/logger"))
		log.Info("Logger middleware enabled")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			reqIdRaw := ctx.Value(reqID.CtxKeyReqId)
			reqId, ok := reqIdRaw.(string)
			if !ok {
				log.Info("Unable to parse req id")
				return
			}
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", reqId))

			t := time.Now()
			defer entry.Info("request completed", slog.String("duration", time.Since(t).String()))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
