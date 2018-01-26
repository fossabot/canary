package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("listen", ":8080", "http listening port")

// defaultHandler gets all requests not going to another handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// statusHandler respondes to /health
func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"status": "OK"}`)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/health", statusHandler)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
