package router

import (
	"github.com/gorilla/mux"
	"github.com/gwkeo/url-shortener/internal/http-server/handlers"
	"github.com/gwkeo/url-shortener/internal/http-server/middleware/logger"
	"github.com/gwkeo/url-shortener/internal/http-server/middleware/reqID"
	"log/slog"
)

func New(log *slog.Logger) *mux.Router {
	router := mux.NewRouter()

	router.Use(reqID.NewReqIdMW(log))
	router.Use(logger.NewLoggerMW(log))
	router.HandleFunc("/", handlers.Shorten())

	return router
}
