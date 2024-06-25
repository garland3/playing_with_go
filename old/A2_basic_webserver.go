package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// run with 
// go run A2_basic_webserver.go