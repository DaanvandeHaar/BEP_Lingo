package storage

import (
	"awesomeProject/pkg/game"
	"fmt"
)

func (s Storage) NewGame(game game.Game) (int, error) {
	var id int
	err := s.db.QueryRow(
		`INSERT INTO games(
								player_id, 
								current_try, 
								five_letter_word, 
								six_letter_word, 
								seven_letter_word, 
								state, 
								time_epoch, 
								score) 
			  VALUES ($1,$2,$3,$4,$5,$6,$7, $8)
			  returning id`,
		game.PlayerID,
		game.CurrentTry,
		game.FiveLetterWord,
		game.SixLetterWord,
		game.SevenLetterWord,
		game.State,
		game.Time,
		game.Score).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}
func (s Storage) GetGame(playerID int, gameID int) interface{} {
	var game game.Game
	err := s.db.QueryRow(`
		SELECT (
		        id, 
		        player_id, 
		        five_letter_word, 
		        six_letter_word, 
		        seven_letter_word, 
		        state, 
		        time_epoch, 
		        score, 
		        current_try)
		FROM games 
		WHERE id == $1 && player_id == $2`, gameID, playerID).Scan(
		&game.ID,
		&game.PlayerID,
		&game.FiveLetterWord,
		&game.SixLetterWord,
		&game.SevenLetterWord,
		&game.State,
		&game.Time,
		&game.Score,
		&game.CurrentTry)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return game
}
