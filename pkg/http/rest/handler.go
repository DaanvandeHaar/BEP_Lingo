package rest

import (
	"awesomeProject/pkg/auth"
	"awesomeProject/pkg/game"
	"awesomeProject/pkg/game/player"
	"awesomeProject/pkg/game/word"
	"encoding/json"
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
		var player player.Player
		var token = r.Header.Get("x-access-token")
		player.UserName, _ = auth.GetUsernameFromToken(token)
		var words []string
		words = append(words, ws.GetRandomWord(5), ws.GetRandomWord(6), ws.GetRandomWord(7))
		game := gs.InitGame(words)
		json.NewEncoder(w).Encode(game)

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
