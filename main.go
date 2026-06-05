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
	// gemini.TestInst()
	// gtranslate.JaUk("Hello, how are you?")
	// fmt.Println(gemini.TranslateEnUk("This is a simple test!"))
	// fmt.Println(gemini.GuesssKanji("念"))
	// fmt.Println(gemini.TranslateJaUk("チューターとは、個人指導の教師、あるいは家庭教師のこと"))

	server.Start()
}
