package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	latest := User{"userone", "userone@userone.ca"}
	json.NewEncoder(w).Encode(latest)
}
