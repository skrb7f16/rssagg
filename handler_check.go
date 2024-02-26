package main

import "net/http"

func handlerCheck(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}
