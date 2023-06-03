package resonse_controller

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Read(w http.ResponseWriter, r *http.Request) {
	log.Println("READS,", "PATH:", r.URL.Path)

	// create response controller...
	rc := http.NewResponseController(w)
	// ... and set deadline of 2 seconds for reading request body
	if err := rc.SetReadDeadline(deadline()); err != nil {
		log.Fatal(err)
	}

	// start reading request body
	var b bytes.Buffer

	if _, err := b.ReadFrom(r.Body); err != nil {
		// because data transmitting is slow we cannot finish reading request body in time
		msg := fmt.Sprintf("ERROR WHILE READING REQUEST BODY: %v", err)
		log.Println(msg)
		_, _ = w.Write([]byte(msg))
	}
}

func CustomRead(w http.ResponseWriter, r *http.Request) {
	log.Println("CUSTOM_READS,", "PATH:", r.URL.Path)

	// create response controller...
	rc := http.NewResponseController(w)
	// ... and set deadline of 2 seconds for reading request body
	if err := rc.SetReadDeadline(deadline()); err != nil {
		log.Fatal(err)
	}

	// start reading request body manually
	for {
		buff := make([]byte, 4096)
		if _, err := r.Body.Read(buff); err != nil {
			if err == io.EOF {
				break
			}

			// skip type assertion for simplicity
			msg := fmt.Sprintf("ERROR WHILE READING REQUEST BODY: %v", err)
			log.Println(msg)
			_, _ = w.Write([]byte(msg))
			break
		}

		// imitate 'long work` which causes deadline to exceed
		time.Sleep(time.Second)
	}
}
