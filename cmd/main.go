package main

import (
	"log"
	"net/http"

	"github.com/komron-m/response_controller"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/read", http.HandlerFunc(resonse_controller.Read))

	mux.Handle("/custom_read", http.HandlerFunc(resonse_controller.CustomRead))

	mux.Handle("/write", http.HandlerFunc(resonse_controller.Write))

	mux.Handle("/custom_write", http.HandlerFunc(resonse_controller.CustomWrite))

	log.Println("START LISTENING FOR REQUESTS...")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
