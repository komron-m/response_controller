package resonse_controller

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Write(w http.ResponseWriter, r *http.Request) {
	log.Println("WRITES,", "PATH:", r.URL.Path)

	// create response controller...
	rc := http.NewResponseController(w)
	// ... and set deadline of 2 seconds for writing response body
	if err := rc.SetWriteDeadline(deadline()); err != nil {
		log.Fatal(err)
	}

	// imitating `long work` which causes deadline to exceed
	time.Sleep(time.Second * 3)

	// this response has no chance to be written
	response := bytes.Repeat([]byte("x"), 4096)

	if _, err := w.Write(response); err != nil {
		log.Println("ERROR WHILE WRITING RESPONSE:", err)
		return
	}
}

func CustomWrite(w http.ResponseWriter, r *http.Request) {
	log.Println("CUSTOM_WRITES,", "PATH:", r.URL.Path)

	// create response controller...
	rc := http.NewResponseController(w)
	// ... and set deadline of 2 seconds for writing response body
	if err := rc.SetWriteDeadline(deadline()); err != nil {
		log.Fatal(err)
	}

	// set appropriate headers for partial response data
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// try writing response, with our custom response writing implementation
	// and on each iteration we write only 1024 bytes
	// => we will need 4 seconds to finish the job
	// => but deadline should occur first
	response := bytes.Repeat([]byte("x"), 4096)
	responseSize := 1024
	for i := 0; i < len(response); i += responseSize {
		if _, err := fmt.Fprintf(w, string(response[i:i+responseSize])); err != nil {
			log.Println("ERROR WHILE WRITING RESPONSE (CUSTOM):", err)
			return
		}

		if err := rc.Flush(); err != nil {
			log.Println("ERROR WHILE FLUSHING RESPONSE (CUSTOM):", err)
			return
		}

		time.Sleep(time.Second) // imitate 'long work`
	}
}
