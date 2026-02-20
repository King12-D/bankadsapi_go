package main

import (
    "log"
    "net/http"

    "bankadsapi_go/internal/handlers"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    http.HandleFunc("/ads/create", handlers.CreateAd)

    log.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
