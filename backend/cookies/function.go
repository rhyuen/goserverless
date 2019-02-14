package main

import (
	"encoding/json"
	"net/http"
)

type Arbitrary struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Node struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Message struct {
	First  Arbitrary `json:"arbitrary"`
	Second Node      `json:"node"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	sampleToSend := Arbitrary{"SampleOne", "SampleContent"}
	NodeSample := Node{"MyNameIsNode", "TheTextIsNodeAsWell"}
	ResponseSample := Message{sampleToSend, NodeSample}
	json.NewEncoder(w).Encode(ResponseSample)
}
