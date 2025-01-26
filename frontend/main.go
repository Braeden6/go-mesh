package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

type BackendResponse struct {
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}

type FrontendResponse struct {
    FrontendMessage string          `json:"frontend_message"`
    BackendData     BackendResponse `json:"backend_data"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        resp, err := http.Get("http://backend/api/data")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var backendResp BackendResponse
        if err := json.Unmarshal(body, &backendResp); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        frontendResp := FrontendResponse{
            FrontendMessage: "Frontend received data asd",
            BackendData:     backendResp,
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(frontendResp)
    })

    log.Println("Frontend service starting on :80")
    log.Fatal(http.ListenAndServe(":80", nil))
}