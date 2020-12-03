package game

import (
	"errors"
)

type State string

type Game struct {
	ID              int    `json:"id"`
	PlayerID        int    `json:"player_id"`
	State           int    `json:"state"`
	CurrentTry      int    `json:"try"`
	FiveLetterWord  string `json:"five_letter_word"`
	SixLetterWord   string `json:"six_letter_word"`
	SevenLetterWord string `json:"seven_letter_word"`
	Score           int    `json:"score"`
	Time            int64  `json:"time"`
}

var ErrGameNotFound = errors.New("game not found")
var ErrGameOver = errors.New("game is already over")

const (
	GAME_STATE_NEW = iota
	GAME_STATE_5LETTER
	GAME_STATE_6LETTER
	GAME_STATE_7LETTER
	GAME_STATE_OVER
)
