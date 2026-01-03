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
	request := r.URL.Query().Get("request")
	if lang != "" && request != "" {
		resp := ""
		switch lang {
		case "en-uk":
			resp = gtranslate.EnUk(request)
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
		ct := "text/plain; charset=utf-8"

		switch task {
		case "guess-kanji":
			resp, err = gemini.GuesssKanji(req)
		case "translate-en-uk":
			ct = "text/html; charset=utf-8"
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
		case "jisho":
			resp, err = grabber.UseJisho(req)
		}

		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, resp)
	}
}
