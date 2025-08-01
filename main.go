package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    user, pass, ok := r.BasicAuth()
    if !ok || user != "admin" || pass != "password" {
        w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintln(w, "Unauthorized")
        return
    }
    fmt.Fprintln(w, "Hello, authorized user!")
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Listening on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
