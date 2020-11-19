package main

import (
	"awesomeProject/pkg/http/rest"
	"log"
	"net/http"
)

// Main function
func main() {

	router := rest.Handler()

	log.Fatal(http.ListenAndServe(":8000", router))

}
