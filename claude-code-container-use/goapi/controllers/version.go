package controllers

import (
	"encoding/json"
	"net/http"
	"os"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
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