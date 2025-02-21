package reqID

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type ReqID string

const CtxKeyReqId = ReqID("reqID")

func NewReqIdMW(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			id := uuid.New()
			ctx = context.WithValue(ctx, CtxKeyReqId, id.String())
			log = slog.With(slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("request id", id.String()))
			next.ServeHTTP(w, r.WithContext(ctx))
			log.Info("finished handling reqID")
		})
	}
}
