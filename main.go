package main

import (
	"log"
	"net/http"

	"quoteapi/router"
)

func main() {

	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
