package handlers

import (
	"github.com/gwkeo/url-shortener/internal/utils"
	"net/http"
)

type Creator interface {
	Create(string, string) error
}

type Shortener struct {
	creator Creator
}

func NewShortener(c Creator) *Shortener {
	return &Shortener{
		creator: c,
	}
}

func (s *Shortener) Shorten() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing url parameter", http.StatusBadRequest)
			return
		}

		shortened := utils.GenerateRandomString(len(url))
		err := s.creator.Create(url, shortened)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", shortened)
		w.WriteHeader(http.StatusCreated)
		return
	}
}
