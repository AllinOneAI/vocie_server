package main

import (
    "log"
    "net/http"
)

func main() {
    // Serve files from the "voices" directory
    fs := http.FileServer(http.Dir("./voices"))
    http.Handle("/voices/", http.StripPrefix("/voices/", fs))

    // Start the HTTP server on port 8080
    log.Println("Hosting files on http://localhost:8080/voices/")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}
