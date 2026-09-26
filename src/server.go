package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../website/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
