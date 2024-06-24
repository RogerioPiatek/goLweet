package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Hello world"
	io.WriteString(w, message)
}

func main() {
	http.HandleFunc("/helloworld", helloWorld)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server!")
	}
}
