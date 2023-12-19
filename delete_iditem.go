package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteProfile(q http.ResponseWriter, r *http.Request) {
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

	profiles = append(profiles[:id], profiles[:id+1]...)

	q.WriteHeader(200)

}
