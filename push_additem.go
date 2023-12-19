package main

import (
	"encoding/json"
	"net/http"
)

func addItem(q http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)
	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)
	json.NewEncoder(q).Encode(profiles)
}
