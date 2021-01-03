package main

import (
	"BEP_Lingo/pkg/game"
	"BEP_Lingo/pkg/game/player"
	"BEP_Lingo/pkg/game/word"
	"BEP_Lingo/pkg/http/rest"
	"BEP_Lingo/pkg/storage"
	"log"
	"net/http"
)

// Main function
func main() {
	var wordService word.Service
	var gameService game.Service
	var playerService player.Service
	s := storage.NewStorage()

	gameService = game.NewService(s)
	wordService = word.NewService(s)
	playerService = player.NewService(s)

	router := rest.Handler(wordService, gameService, playerService)

	log.Fatal(http.ListenAndServe(":8080", router))

}
