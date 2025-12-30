package main

// go mod init stateless
// go get google.golang.org/genai
// go get github.com/joho/godotenv

import (
	"fmt"
	"math/rand/v2"
	"stateless/server"

	"github.com/joho/godotenv"
)

func main() {
	test()
	godotenv.Load()
	server.Start()
}

func test() {
	// models := []string{
	// 	"gemini-flash-latest", // "gemini-2.5-flash"
	// 	"gemini-3-flash-preview",
	// 	"gemini-2.5-flash-lite",
	// }
	// // slices.Delete()
	// fmt.Println(len(models))
	count := []int{0, 0, 0}
	for range 60 {
		// model := models[rand.IntN(len(models))]
		// fmt.Println(model)
		ri := rand.IntN(len(count))
		count[ri]++
		if count[ri] > 20 {
			break
		}
	}
	fmt.Println(count)
}
