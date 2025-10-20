package main

import (
    "fmt"
    "net/http"
    "go-url-shortener/shortener"
)

func main() {
    http.HandleFunc("/shorten", shortener.ShortenHandler)
    http.HandleFunc("/", shortener.RedirectHandler)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
