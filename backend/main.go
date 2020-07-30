package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: migrate to https://github.com/valyala/fasthttp
	http.HandleFunc("/session", Session)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
