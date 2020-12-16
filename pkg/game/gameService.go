package game

import (
	"BEP_Lingo/pkg/game/word"
	"errors"
	"fmt"
	"time"
)

type Service interface {
	InitGame([]string, int) (Game, error)
	RaiseGameState(int, int) bool
	RaiseTryCount(int, int) bool
	RaiseGameScore(int, int, int) bool
	GameRunner(word.Service, string, int) (word.LingoMessage, error)
	GetCurrentGame(int) (Game, error)
}

type Repository interface {
	NewGame(Game) (int, error)
	RaiseGameState(int, int) bool
	RaiseTryCount(int, int) bool
	RaiseGameScore(int, int, int) bool
	ResetTryCount(int, int) bool
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
	if s.r == nil {
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
	if s.r == nil {
		return ws.GetEmptyMessage(), errors.New("Error, nil pointer")
	}
	game, err := s.r.GetCurrentGame(playerID)
	if err != nil {
		result, _ := ws.CompareWords("", "")
		return result, ErrGameNotFound
	}
	switch game.State {
	case GAME_STATE_NEW:
		{
			updated := s.r.RaiseGameState(game.ID, game.PlayerID)
			if updated == false {
				return ws.GetEmptyMessage(), errors.New("Error, could not update game state")
			}
			help := ws.GetWordHelp(game.FiveLetterWord)
			message := ws.GetEmptyMessage()
			message.Info = help
			return message, nil
		}
	case GAME_STATE_5LETTER:
		if game.CurrentTry <= 5 {
			message, err := ws.CompareWords(word, game.FiveLetterWord)
			if err != nil {
				return ws.GetEmptyMessage(), errors.New("Error, Something went wrong while comparing words")
			}
			if &message.Correct != nil && message.Correct {
				message.Info = "U have guessed the word correctly. Continuing to 6 letter game stage. Hint for 6 letter word: " + ws.GetWordHelp(game.SixLetterWord)
				s.r.RaiseGameScore(game.ID, game.PlayerID, (50)-(game.CurrentTry*10))
				s.r.ResetTryCount(game.ID, game.PlayerID)
				s.r.RaiseGameState(game.ID, game.PlayerID)
				return message, nil
			} else {
				s.r.RaiseTryCount(game.ID, game.PlayerID)
				message.TryIndex = game.CurrentTry
				return message, nil
			}
		} else {
			message := ws.GetEmptyMessage()
			message.Info = "You have not guessed the correct word in time. Continuing to 6 letter game stage. Hint for 6 letter word: " + ws.GetWordHelp(game.SixLetterWord)
			s.r.ResetTryCount(game.ID, game.PlayerID)
			s.r.RaiseGameState(game.ID, game.PlayerID)
			return message, nil
		}
	case GAME_STATE_6LETTER:
		if game.CurrentTry <= 5 {
			message, err := ws.CompareWords(word, game.SixLetterWord)
			if err != nil {
				return ws.GetEmptyMessage(), errors.New("Error, Something went wrong while comparing words")
			}
			if &message.Correct != nil && message.Correct {
				message.Info = "U have guessed the word correctly. Continuing to 7 letter game stage. Hint for 7 letter word: " + ws.GetWordHelp(game.SevenLetterWord)
				s.r.RaiseGameScore(game.ID, game.PlayerID, (50)-(game.CurrentTry*10))
				s.r.ResetTryCount(game.ID, game.PlayerID)
				s.r.RaiseGameState(game.ID, game.PlayerID)
				return message, nil
			} else {
				s.r.RaiseTryCount(game.ID, game.PlayerID)
				message.TryIndex = game.CurrentTry
				return message, nil
			}
		} else {
			message := ws.GetEmptyMessage()
			message.Info = "You have not guessed the correct word in time. Continuing to 7 letter game stage. Hint for 7 letter word: " + ws.GetWordHelp(game.SevenLetterWord)
			s.r.ResetTryCount(game.ID, game.PlayerID)
			s.r.RaiseGameState(game.ID, game.PlayerID)
			return message, nil
		}
	case GAME_STATE_7LETTER:
		if game.CurrentTry <= 5 {
			message, err := ws.CompareWords(word, game.SevenLetterWord)
			if err != nil {
				return ws.GetEmptyMessage(), errors.New("Error, Something went wrong while comparing words")
			}
			if &message.Correct != nil && message.Correct {
				message.Info = "U have guessed the word correctly. The game is finished "
				s.r.RaiseGameScore(game.ID, game.PlayerID, (50)-(game.CurrentTry*10))
				s.r.ResetTryCount(game.ID, game.PlayerID)
				s.r.RaiseGameState(game.ID, game.PlayerID)
				return message, nil
			} else {
				s.r.RaiseTryCount(game.ID, game.PlayerID)
				message.TryIndex = game.CurrentTry
				return message, nil
			}
		} else {
			message := ws.GetEmptyMessage()
			message.Info = "You have not guessed the correct word in time. The game is finished"
			s.r.RaiseGameState(game.ID, game.PlayerID)
			return message, nil
		}
	case GAME_STATE_OVER:
		message := ws.GetEmptyMessage()
		message.Info = "This game is already finished!"
		return message, nil
	}
	return ws.GetEmptyMessage(), nil
}

func (s *service) RaiseGameScore(gameID int, playerID int, score int) bool {
	if s.r == nil {
		fmt.Println("Service not found")
		return false
	}
	if s.r.RaiseGameScore(gameID, playerID, score) {
		return true
	}
	return false
}

func (s *service) RaiseGameState(gameID int, playerID int) bool {
	if s.r == nil {
		fmt.Println("Service not found")
		return false
	}
	if s.r.RaiseGameState(gameID, playerID) {
		return true
	}
	return false
}

func (s *service) RaiseTryCount(gameID int, playerID int) bool {
	if s.r == nil {
		fmt.Println("Service not found")
		return false
	}
	if s.r.RaiseTryCount(gameID, playerID) {
		return true
	}
	return false
}
