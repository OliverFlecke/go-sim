package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
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
		err := filepath.Walk(PATH_TO_LEVELS, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}

			if !info.IsDir() {
				w.Write([]byte(strings.TrimPrefix(path, PATH_TO_LEVELS+"/") + "\n"))
			}

			return nil
		})

		if err != nil {
			http.Error(w, "Could not read maps", http.StatusInternalServerError)
			return
		}
	})
}
