package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	// w.Write([]byte("Display a specific snippet ..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST"{
		// below must be called for all errors codes except 200
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

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