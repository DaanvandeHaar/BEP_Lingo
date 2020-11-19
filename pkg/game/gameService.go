package game

import "awesomeProject/pkg/persistence"

func newGame() Game {
	game := Game{
		ID:         nil,
		State:      GAME_STATE_5LETTER,
		CurrentTry: 0,
		Word:       persistence.GetRandomWord(5),
	}
	return game
}
func makeGuess() {

}
