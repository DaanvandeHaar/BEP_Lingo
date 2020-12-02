package game

import (
	"fmt"
	"time"
)

type Service interface {
	InitGame([]string, int) (Game, error)
	RaiseGameState(int) bool
}

type Repository interface {
	NewGame(Game) (int, error)
}
type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) InitGame(words []string, playerID int) (Game, error) {
	game := Game{
		State:      GAME_STATE_NEW,
		CurrentTry: 0,
		Score:      0,
		Time:       time.Now().Unix(),
	}
	game.FiveLetterWord = words[0]
	game.SixLetterWord = words[1]
	game.SevenLetterWord = words[2]
	game.PlayerID = playerID
	gameID, err := s.r.NewGame(game)
	if err != nil {
		fmt.Println(err)
		return game, err
	}
	game.ID = gameID
	return game, nil
}
func (s *service) RaiseGameState(id int) bool {

	return true
}
