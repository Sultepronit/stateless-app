package server

import (
	"io"
	"net/http"
	"stateless/gemini"
	"stateless/grabber"
	"stateless/gtranslate"
)

func handleGtranslate(w http.ResponseWriter, r *http.Request) {
	lang := r.PathValue("lang")
	text := r.URL.Query().Get("text")
	if lang != "" && text != "" {
		resp := ""
		switch lang {
		case "en-uk":
			resp = gtranslate.EnUk(text)
		}
		io.WriteString(w, resp)
	}
}

func handleArtificial(w http.ResponseWriter, r *http.Request) {
	task := r.PathValue("task")
	req := r.URL.Query().Get("request")
	if task != "" && req != "" {
		resp := ""
		switch task {
		case "guess-kanji":
			resp = gemini.GuesssKanji(req)
		case "translate-en-uk":
			resp = gemini.TranslateEnUk(req)
		}
		io.WriteString(w, resp)
	}

}

func handleGrabber(w http.ResponseWriter, r *http.Request) {
	task := r.PathValue("task")
	req := r.URL.Query().Get("request")
	if task != "" && req != "" {
		resp := ""
		switch task {
		case "e2u":
			grabber.UseE2u(req)
		}

		io.WriteString(w, resp)
	}
}
