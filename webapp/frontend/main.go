package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	hostname := ":80"

	fmt.Printf("Starting Frontend-Server at: %s\n", hostname)

	http.Handle("/", http.FileServer(http.Dir("statics/")))
	err := http.ListenAndServe(hostname, nil)
	if err != nil {
		log.Fatal(err)
	}

}
