package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /gtranslate/{lang}", handleGtranslate) // gtranslate/en-uk?request=...
	mux.HandleFunc("GET /artificial/{task}", handleArtificial) // artificial/guess-kanji?request=...
	mux.HandleFunc("GET /grabber/{task}", handleGrabber)       // grabber/e2u?request=...
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here we go!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Listening on port: %s\n", port)

	handler := cors.Default().Handler(mux)
	// err := http.ListenAndServe(":"+port, mux)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
