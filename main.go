package main

// go mod init stateless
// go get google.golang.org/genai
// go get github.com/joho/godotenv

import (
	"stateless/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// gemini.Test()
	// gtranslate.JaUk("Hello, how are you?")

	server.Start()
}
