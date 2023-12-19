package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func updateProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID Can not be converted"))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("No Profile Found"))
		return
	}

	var updateProfile Profile
	json.NewDecoder(r.Body).Decode(&updateProfile)

	profiles[id] = updateProfile
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(updateProfile)

}
