package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithErr(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		w.WriteHeader(status)
		return
	}
	type errorType struct {
		Error string `json:"error"`
	}
	responseWithJson(w, status, errorType{
		Error: msg,
	})
}
func responseWithJson(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error %v", payload)
		w.WriteHeader(500)
		return

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
