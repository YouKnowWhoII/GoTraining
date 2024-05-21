package main

import (
	"fmt"
	"net/http"
)

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my website!</h1>")
}

func startWebServer() {
	http.HandleFunc("/", serveHomePage)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
