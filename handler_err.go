package main

import "net/http"

func handlerErr(w http.ResponseWriter, req *http.Request) {
	respondWithError(w, 400, "something went wrong")
}
