package storage

import (
	"BEP_Lingo/pkg/game"
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

func (s *Storage) raiseGameState(gameID int, playerID int) bool {
	_, err := s.db.Query("UPDATE game SET state = state + 1 WHERE id == $1 && player_id ==$2", gameID, playerID)
	if err != nil {
		return false
	}
	return true
}

func (s *Storage) raiseTryState(gameID int, playerID int) bool {
	_, err := s.db.Query("UPDATE game SET current_try = current_try + 1 WHERE id == $1 && player_id ==$2", gameID, playerID)
	if err != nil {
		return false
	}
	return true
}
func (s *Storage) raiseGameScore(gameID int, playerID int) bool {
	_, err := s.db.Query("UPDATE game SET score = score + 1 WHERE id == $1 && player_id ==$2", gameID, playerID)
	if err != nil {
		return false
	}
	return true
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