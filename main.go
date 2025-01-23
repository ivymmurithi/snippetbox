package main

import (
	"log"
	"net/http"
)

// Handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Router
func main() {
	// entry point of requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	// "host:port"
	// if port is http or http-alt it will look for /etc/services
	err := http.ListenAndServe(":4000", mux)
	// log print errors with time and date and in the error stream(stack trace?)
	// fatal is  like print and an exit to the app
	log.Fatal(err)
}