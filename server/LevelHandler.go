package main

import (
	"net/http"
	"simulator/core/level"
	"strings"
)

type LevelHandler struct {
}

func (h *LevelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		switch head {
		case "":
			h.GetMaps().ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}

	default:
		http.NotFound(w, r)
	}
}

func (h *LevelHandler) GetMaps() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level.GetMaps(PATH_TO_LEVELS,
			func(level string) {
				w.Write([]byte(strings.TrimPrefix(level, PATH_TO_LEVELS+"/") + "\n"))
			})
	})
}
