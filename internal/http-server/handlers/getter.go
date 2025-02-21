package handlers

import (
	"log"
	"net/http"
)

type URLGetter interface {
	URL(string) (string, error)
}

type Getter struct {
	urlGetter URLGetter
}

func NewGetter(urlGetter URLGetter) *Getter {
	return &Getter{urlGetter: urlGetter}
}

func (g *Getter) URL() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := g.urlGetter.URL(r.URL.Query().Get("shortUrl"))
		log.Println(url, err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
		return
	}
}
