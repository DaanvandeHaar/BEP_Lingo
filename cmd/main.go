package main

import (
	"awesomeProject/pkg/game/word"
	"awesomeProject/pkg/http/rest"
	"log"
	"net/http"
)

// Main function
func main() {
	var wordService word.Service

	router := rest.Handler(wordService)

	log.Fatal(http.ListenAndServe(":8000", router))

}
