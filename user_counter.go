package main

import (
	"bytes"
	"html/template"
	"path/filepath"
)

type UserCounter int

var userCounterTemplate = filepath.Join("templates", "user_counter.html")

func (uc UserCounter) Bytes() []byte {
	templ := template.Must(template.ParseFiles(userCounterTemplate))
	buffer := new(bytes.Buffer)
	templ.ExecuteTemplate(buffer, "user_counter", uc)
	return buffer.Bytes()
}

func (uc UserCounter) Increment() UserCounter {
	return uc + 1
}
