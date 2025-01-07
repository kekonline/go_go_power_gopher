package handlers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "sync"
)

type Response struct {
    Message string `json:"message"`
}

var wg sync.WaitGroup

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    defer wg.Done()
    if r.Method == http.MethodGet {
        HandleGet(w, r)
    } else if r.Method == http.MethodPost {
        HandlePost(w, r)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile("data.json")
    if err != nil {
        http.Error(w, "Error reading file", http.StatusInternalServerError)
        return
    }
    var response Response
    if err := json.Unmarshal(data, &response); err != nil {
        http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
    var response Response
    if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
        http.Error(w, "Error decoding JSON", http.StatusBadRequest)
        return
    }
    if err := writeToFile(response); err != nil {
        http.Error(w, "Error writing to file", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func writeToFile(response Response) error {
    file, err := os.OpenFile("data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    data, err := json.Marshal(response)
    if err != nil {
        return err
    }
    
    if _, err := file.Write(data); err != nil {
        return err
    }
    return nil
}

func StartServer() {
    http.HandleFunc("/api", HandleRequest)
    wg.Add(1)
    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            fmt.Println("Server failed:", err)
        }
    }()
}