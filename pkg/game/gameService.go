package game

import (
	"BEP_Lingo/pkg/game/word"
	"fmt"
	"time"
)

type Service interface {
	InitGame([]string, int) (Game, error)
	RaiseGameState(int, int) bool
	RaiseTryCount(int, int) bool
	RaiseGameScore(int, int) bool
	GameRunner(word.Service, string, int) (word.LingoMessage, error)
	GetCurrentGame(int) (Game, error)
}

type Repository interface {
	NewGame(Game) (int, error)
	RaiseGameState(int, int) bool
	RaiseTryCount(int, int) bool
	RaiseGameScore(int, int) bool
	GetCurrentGame(int) (Game, error)
	GetGameForID(int, int) (Game, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) InitGame(words []string, playerID int) (Game, error) {
	if s == nil {
		fmt.Println("Service not found")
		return Game{}, fmt.Errorf("Err: service not found")
	}
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
func (s *service) GetCurrentGame(playerID int) (Game, error) {
	game, err := s.r.GetCurrentGame(playerID)
	if err != nil {
		return game, ErrGameNotFound
	}
	return game, nil
}
func (s *service) GameRunner(ws word.Service, word string, playerID int) (word.LingoMessage, error) {
	game, err := s.r.GetCurrentGame(playerID)
	if err != nil {
		return ws.CompareWords("", ""), ErrGameNotFound
	}
	switch game.State {
	case GAME_STATE_NEW:
		{
			s.r.RaiseGameState(game.ID, game.PlayerID)
			return ws.CompareWords(game.FiveLetterWord, ws.GetWordHelp(game.SixLetterWord)), nil
		}
	case GAME_STATE_5LETTER:
		if game.CurrentTry >= 5 {
			s.r.RaiseGameState(game.ID, game.PlayerID)
			return ws.CompareWords(game.SixLetterWord, ws.GetWordHelp(game.SixLetterWord)), nil
		} else {

		}
	case GAME_STATE_6LETTER:
		if game.CurrentTry != 5 {

		} else {

		}
	case GAME_STATE_7LETTER:
		if game.CurrentTry != 5 {

		} else {

		}
	case GAME_STATE_OVER:
		if game.CurrentTry != 5 {

		} else {

		}

	}
	return ws.CompareWords("", ""), nil
}

func (s *service) RaiseGameScore(gameID int, playerID int) bool {
	if s == nil {
		fmt.Println("Service not found")
		return false
	}
	if s.r.RaiseGameScore(gameID, playerID) {
		return true
	}
	return false
}

func (s *service) RaiseGameState(gameID int, playerID int) bool {
	if s == nil {
		fmt.Println("Service not found")
		return false
	}
	if s.r.RaiseGameState(gameID, playerID) {
		return true
	}
	return false
}

func (s *service) RaiseTryCount(gameID int, playerID int) bool {
	if s == nil {
		fmt.Println("Service not found")
		return false
	}
	if s.r.RaiseTryCount(gameID, playerID) {
		return true
	}
	return false
}
