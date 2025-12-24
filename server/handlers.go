package server

import (
	"io"
	"net/http"
	"stateless/gemini"
	"stateless/gtranslate"
)

func handleGtranslate(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text != "" {
		resp := gtranslate.EnUk(text)
		io.WriteString(w, resp)
	}
}

func handleArtificial(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Query().Get("req")
	if req != "" {
		resp := gemini.GuesssKanji(req)
		io.WriteString(w, resp)
	}

}