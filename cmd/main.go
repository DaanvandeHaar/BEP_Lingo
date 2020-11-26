package main

import (
	"awesomeProject/pkg/game"
	"awesomeProject/pkg/game/player"
	"awesomeProject/pkg/game/word"
	"awesomeProject/pkg/http/rest"
	"awesomeProject/pkg/storage"
	"log"
	"net/http"
)

// Main function
func main() {
	var wordService word.Service
	var gameService game.Service
	var playerService player.Service
	s := storage.NewStorage()

	wordService = word.NewService(s)
	playerService = player.NewService(s)

	router := rest.Handler(wordService, gameService, playerService)

	log.Fatal(http.ListenAndServe(":8000", router))

}
