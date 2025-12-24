package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/health", healthHandler)

    log.Println("starting http server on :8080")

    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}


