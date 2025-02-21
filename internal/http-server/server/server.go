package server

import (
	"github.com/gorilla/mux"
	"github.com/gwkeo/url-shortener/internal/config"
	"github.com/gwkeo/url-shortener/internal/http-server/handlers"
	"github.com/gwkeo/url-shortener/internal/http-server/middleware/logger"
	"github.com/gwkeo/url-shortener/internal/http-server/middleware/reqID"
	"log/slog"
	"net/http"
)

type Repo interface {
	Create(string, string) error
	URL(string) (string, error)
}

type Server struct {
	repo   Repo
	router *mux.Router
	logger *slog.Logger
	cfg    *config.Config
}

func New(repo Repo, log *slog.Logger, cfg *config.Config) *Server {
	router := mux.NewRouter()

	server := &Server{
		repo:   repo,
		router: router,
		logger: log,
		cfg:    cfg,
	}
	return server
}

func (s *Server) Start() error {

	s.router.Use(reqID.NewReqIdMW(s.logger))
	s.router.Use(logger.NewLoggerMW(s.logger))

	s.setRoutes()
	if err := http.ListenAndServe(s.cfg.Addr, s.router); err != nil {
		return err
	}
	return nil
}

func (s *Server) setRoutes() {
	shortenerHandler := handlers.NewShortener(s.repo)
	s.router.HandleFunc("/shorten", shortenerHandler.Shorten()).Methods("POST")
	getterURLHandler := handlers.NewGetter(s.repo)
	s.router.HandleFunc("/redirect", getterURLHandler.URL()).Methods("GET")
}
