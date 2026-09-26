package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting client...")
	resp, err := http.Get("http://localhost:8080/get")
	if err != nil {
		log.Fatal("Error: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: %v", err)
		return
	}
	fmt.Println(string(body))
}
