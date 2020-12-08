package rest

import (
	"BEP_Lingo/pkg/auth"
	"BEP_Lingo/pkg/game"
	"BEP_Lingo/pkg/game/player"
	"BEP_Lingo/pkg/game/word"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Handler(ws word.Service, gs game.Service, ps player.Service) *mux.Router {

	//create new mux router
	router := mux.NewRouter()

	//create subrouter for JWT auth requests
	authRouter := router.PathPrefix("/api").Subrouter()

	authRouter.Use(auth.JwtVerify)

	// Route handles & endpoints
	authRouter.HandleFunc("/word/getrandom", getRandomWord(ws)).Methods("GET")
	authRouter.HandleFunc("/game/current/guess", guessWord(ws)).Methods("POST")
	router.HandleFunc("/jwt", getJwt).Methods("GET")
	authRouter.HandleFunc("/game/new", newGame(gs, ws, ps)).Methods("POST")
	router.HandleFunc("/auth/login", login(ps)).Methods("POST")
	router.HandleFunc("/auth/signup", signUp(ps)).Methods("POST")

	// return router to main.go
	return router
}
func newGame(gs game.Service, ws word.Service, ps player.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var token = r.Header.Get("x-access-token")
		var words []string
		username, err := auth.GetUsernameFromToken(token)
		if err != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode("Error, Could not find or process token")
			return
		}
		id, err := ps.GetIDForPlayer(username)
		fmt.Println(id)
		if id == 0 {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode("Error, Could not find user")
			return
		}
		words = append(words, ws.GetRandomWord(5), ws.GetRandomWord(6), ws.GetRandomWord(7))
		game, err := gs.InitGame(words, id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode("Error, Could not save new game")

		}
		json.NewEncoder(w).Encode(game.ID)

	}
}
func getRandomWord(ws word.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content_Type", "application/json")
		json.NewEncoder(w).Encode(ws.GetRandomWord(7))
	}

}

func signUp(p player.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var player player.Player
		_ = json.NewDecoder(r.Body).Decode(&player)
		bool := p.SignUp(player)
		if bool == false {
			json.NewEncoder(w).Encode("could not create new user")
			return
		}
		json.NewEncoder(w).Encode("new user with username: " + player.UserName + " created")
	}
}

func getJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	player := player.Player{
		UserName: "us",
		Password: "dddd",
	}
	json.NewEncoder(w).Encode(auth.GenerateJWT(player))

}

func login(p player.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var player player.Player
		_ = json.NewDecoder(r.Body).Decode(&player)
		if player.UserName == "" || player.Password == "" {
			json.NewEncoder(w).Encode("err: username or password was not filled in correctly")
		}
		bool, token := p.Login(player)
		if bool == false {
			json.NewEncoder(w).Encode("err: Wrong username or password")
		} else {
			json.NewEncoder(w).Encode(token)
		}
	}
}

func guessWord(s word.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//var token = r.Header.Get("x-access-token")
		var input word.Word
		_ = json.NewDecoder(r.Body).Decode(&input)
		if s.CheckIfAlpha(input) {
			if len(input.Word) == len("hallo") {
				json.NewEncoder(w).Encode(s.CompareWords("hallo", "hallo"))
			} else {
				json.NewEncoder(w).Encode(false)
			}
		} else {
			json.NewEncoder(w).Encode(false)
		}
	}
}
