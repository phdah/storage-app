package api

import (
    "encoding/json"
    "net/http"
    "sync"
)

type Response struct {
    Message string `json:"message"`
}


var response = Response{Message: "Hello World!"} // Variable accessible from API
var mu sync.Mutex // Lock for altering `response`

// GET function to retrieve the response variable. See `UpdateString`
// on how to alter the response variable
func GetString(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock() // unlock then return

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// PUT function to update the response variable, which can be
// retrieved by `GetString`
func UpdateString(w http.ResponseWriter, r *http.Request) {
    var newResponse Response
    if err := json.NewDecoder(r.Body).Decode(&newResponse); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    mu.Lock()
    response.Message = newResponse.Message
    mu.Unlock()
    json.NewEncoder(w).Encode(response)
}
