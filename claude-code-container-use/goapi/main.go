package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("version.json")
	if err != nil {
		http.Error(w, "Could not read version file", http.StatusInternalServerError)
		return
	}

	var versionData VersionResponse
	if err := json.Unmarshal(data, &versionData); err != nil {
		http.Error(w, "Could not parse version file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(versionData)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/version", versionHandler)
	http.ListenAndServe(":8080", nil)
}
