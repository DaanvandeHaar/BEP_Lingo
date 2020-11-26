package game

type Service interface {
	NewGame() Game
}

func InitGame() Game {
	game := Game{
		ID:         1,
		State:      GAME_STATE_NEW,
		CurrentTry: 0,
		Word:       "a",
	}
	return game
}
func makeGuess() {

}
