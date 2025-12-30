package server

import (
	"io"
	"log"
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
		var resp string
		var err error
		ct := "text/plain"

		switch task {
		case "guess-kanji":
			resp, err = gemini.GuesssKanji(req)
		case "translate-en-uk":
			ct = "text/html"
			resp, err = gemini.TranslateEnUk(req)
		}

		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", ct)
		io.WriteString(w, resp)
	}

}

func handleGrabber(w http.ResponseWriter, r *http.Request) {
	task := r.PathValue("task")
	req := r.URL.Query().Get("request")
	if task != "" && req != "" {
		var resp string
		var err error

		switch task {
		case "e2u":
			resp, err = grabber.UseE2u(req)
		case "slovnyk":
			resp, err = grabber.UseSlovnyk(req)
		}

		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, resp)
	}
}
