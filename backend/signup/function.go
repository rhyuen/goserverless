package main

import (
	"encoding/json"
	"net/http"

	"github.com/rhyuen/go-fs-lambda/backend/arbitrary"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	latest := arbitrary.Arbitrary{"A title", "A Content"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(latest)

}
