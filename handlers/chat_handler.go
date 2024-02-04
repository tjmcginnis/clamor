package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/tjmcginnis/clamor"
)

type ChatHandler struct {
	Counter *clamor.UserCounter
}

var templateFiles = []string{
	filepath.Join("templates", "index.html"),
	filepath.Join("templates", "user_counter.html"),
}

func (h ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles(templateFiles...))
	templ.ExecuteTemplate(w, "index", struct {
		Host    string
		Counter *clamor.UserCounter
	}{
		Host:    r.Host,
		Counter: h.Counter,
	})
}
