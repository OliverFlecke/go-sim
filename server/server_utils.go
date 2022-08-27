package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func SendSSE(w http.ResponseWriter, event string, data []byte) {
	fmt.Fprintf(w, "event: %s\n", event)
	fmt.Fprintf(w, "data: %s\n\n", data)
	w.(http.Flusher).Flush()
}
