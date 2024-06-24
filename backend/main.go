package main

import (
	"fmt"
	"io"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rogeriopiatek/goLweet/backend/db"
	"github.com/rogeriopiatek/goLweet/backend/handlers/home"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Hello world"
	io.WriteString(w, message)
}

func main() {
	db.Init()

	http.HandleFunc("/helloworld", helloWorld)

	http.HandleFunc("/", home.HomeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server!")
	}
}
