package game

import (
	"errors"
)

type State string

type Game struct {
	ID         int
	State      int
	CurrentTry int
	Word       string
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
