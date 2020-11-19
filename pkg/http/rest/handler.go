package rest

import (
	"awesomeProject/pkg/auth"
	"awesomeProject/pkg/game/player"
	"awesomeProject/pkg/game/word"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Init books var as a slice Book struct

func Handler() *mux.Router {

	router := mux.NewRouter()
	authRouter := router.PathPrefix("/api").Subrouter()

	authRouter.Use(auth.JwtVerify)

	// Route handles & endpoints
	authRouter.HandleFunc("/game/current/guess", guessWord).Methods("POST")
	router.HandleFunc("/jwt", getJwt).Methods("GET")
	//router.HandleFunc("/game/new", newGame).Methods("POST")
	//router.HandleFunc("/auth/login", login).Methods("POST")
	//router.HandleFunc("/auth/logout", logout).Methods("POST")
	//router.HandleFunc("/game/mygames"

	// return router to main.go
	return router
}

func getJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	player := player.Player{
		UserName: "us",
		Password: "dddd",
	}
	json.NewEncoder(w).Encode(auth.GenerateJWT(player))

}

//get data for old games

func getWordForGame(gameId string) bool {
	return false
}

func guessWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input word.Word
	_ = json.NewDecoder(r.Body).Decode(&input)
	fmt.Println(input.Word)
	if word.CheckIfAlpha(input.Word) {
		if len(input.Word) == len("testtes") {
			json.NewEncoder(w).Encode(word.CompareWords("hallo", "hallo"))
		} else {
			json.NewEncoder(w).Encode(false)
		}
	} else {
		json.NewEncoder(w).Encode(false)
	}
}
