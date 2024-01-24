package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"path/filepath"
)

// Message represents a chat message.
type Message struct {
	User User
	Body string
}

var messageTemplate = filepath.Join("templates", "message.html")

func (m *Message) Unmarshal(b []byte) error {
	return json.Unmarshal(b, m)
}

func (m *Message) Bytes() []byte {
	templ := template.Must(template.ParseFiles(messageTemplate))
	buffer := new(bytes.Buffer)
	templ.Execute(buffer, m)
	return buffer.Bytes()
}
