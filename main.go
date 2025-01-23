package main

import (
	"log"
	"net/http"
)

// Handler
func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet ..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet ..."))
}

// Router
func main() {
	// entry point of requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//! mux.HandleFunc("/snippet", showSnippet) can work without the initialized serve mux
	//! but it is not secure because it uses the dafault serve mux which is a global variable that can be accessed by 3rd parties
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	// "host:port"
	// if port is http or http-alt it will look for /etc/services
	err := http.ListenAndServe(":4000", mux)
	// log print errors with time and date and in the error stream(stack trace?)
	// fatal is  like print and an exit to the app
	log.Fatal(err)
}