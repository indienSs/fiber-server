package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":3000"
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received", r.Method)
	w.Write([]byte("hello from server"))
}

func main() {
	http.HandleFunc("/", itemsHandler)
	fmt.Printf("server started on PORT%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("server start error", err)
	}
}