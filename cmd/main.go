package main

import (
    "log"
    "net/http"

    "github.com/phdah/storage-app/internals/api"
)

func main() {
    http.HandleFunc("/api/string", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            api.GetString(w, r)
        case http.MethodPut:
            api.UpdateString(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server staring on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
