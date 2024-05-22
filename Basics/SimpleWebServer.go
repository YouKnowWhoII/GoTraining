package main

import (
	"fmt"
	"net/http"
)

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	serve, err := fmt.Fprint(w, "<h1>Welcome to my website!</h1>")
	if err != nil {
		fmt.Println("Error serving home page:", err)
	} else {
		fmt.Println("Served home page successfully:", serve)
	}
}

func startWebServer() {
	http.HandleFunc("/", serveHomePage)
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
