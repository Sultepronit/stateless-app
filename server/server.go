package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /gtranslate", handleGtranslate)
	mux.HandleFunc("GET /artificial", handleArtificial)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here we go!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Listening on port: %s\n", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
