package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

var counter int = 0

func increment(w http.ResponseWriter, r *http.Request) {
    counter++
	w.WriteHeader(http.StatusOK)
}

func show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"counter\": %s}", strconv.FormatInt(int64(counter), 10))
}

func healthz(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func main() {
    http.HandleFunc("/increment", increment)
    http.HandleFunc("/show", show)
    http.HandleFunc("/healthz", healthz)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
