package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type GameData struct {
	ID       string
	TryCount int
	Tries    []Try
}

type Try struct {
	Letters []Letter
}

type Letter struct {
	LetterString   string
	LetterPosition int
	RightPlace     bool
	RightLetter    bool
}
type Word struct {
	Word string `json:"word"`
}

const guesword = "schepen"

// Init books var as a slice Book struct

var gameData []GameData

func getGameData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range gameData {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&GameData{})
}
func getWordForGame(gameId string) bool {
	return false
}

func checkIfValidWord(s string) bool {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

// Add new book
func checkIfValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var word Word
	_ = json.NewDecoder(r.Body).Decode(&word)
	fmt.Println(word.Word)
	if checkIfValidWord(word.Word) {
		if len(word.Word) == len(guesword) {
			json.NewEncoder(w).Encode(compareWords(word.Word, guesword))
		} else {
			json.NewEncoder(w).Encode(false)
		}
	} else {
		json.NewEncoder(w).Encode(false)
	}
}

func compareWords(word string, comepareWord string) Try {
	var try Try
	for pos, char := range word {
		if guesword[pos] == word[pos] {
			fmt.Println(true, string(char))
			try.Letters = append(try.Letters, Letter{
				LetterString:   string(char),
				LetterPosition: pos,
				RightPlace:     true,
				RightLetter:    true,
			})
		} else {
			fmt.Println(false, string(char))
			if strings.ContainsAny(guesword, string(char)) {
				try.Letters = append(try.Letters, Letter{
					LetterString:   string(char),
					LetterPosition: pos,
					RightPlace:     false,
					RightLetter:    true,
				})
			} else {
				try.Letters = append(try.Letters, Letter{
					LetterString:   string(char),
					LetterPosition: pos,
					RightPlace:     false,
					RightLetter:    false,
				})
			}
		}
	}
	return try
}

func StartHttpServer() {
	r := mux.NewRouter()

	// Route handles & endpoints
	//r.Handle("game/", getGamesData).Methods("GET")
	r.HandleFunc("/game/{id}", getGameData).Methods("GET")
	r.HandleFunc("/word", checkIfValid).Methods("POST")

	// Start server
	fmt.Println("Starting server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
