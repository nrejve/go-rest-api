package main

import (
	"encoding/json"
	"net/http"
)

func getAllProfiles(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)
}
