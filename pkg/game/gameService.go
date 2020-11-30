package game

type Service interface {
	InitGame([]string) Game
}

type Repository interface {
}
type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) InitGame(words []string) Game {
	game := Game{
		State:      GAME_STATE_NEW,
		CurrentTry: 0,
		Score:      0,
		Time:       nil,
	}
	game.FiveLetterWord = words[0]
	game.SixLetterWord = words[1]
	game.SevenLetterWord = words[2]

	return game
}
