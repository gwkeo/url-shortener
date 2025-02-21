package handlers

import "net/http"

type Creator interface {
	Create(url, shortened string) (int64, error)
}

type Shortener struct {
	creator Creator
}

func NewShortener(c Creator) *Shortener {
	return &Shortener{
		creator: c,
	}
}

func Shorten() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
