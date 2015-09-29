package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Change log output
	log.SetOutput(os.Stdout)

	// Set handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Access to %s%s", r.Host, r.URL.String())
		fmt.Fprint(w, "Hello world\n")
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		// Set default port
		port = "3000"
	}

	log.Printf("Start to listen on :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
