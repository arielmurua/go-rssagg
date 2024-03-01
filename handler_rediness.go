package main

import "net/http"

func handlerRediness(w http.ResponseWriter, req *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
