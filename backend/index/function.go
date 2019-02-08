package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Node struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, go stuff. deux")

	data := Node{"This is a Title", "This is an author"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
