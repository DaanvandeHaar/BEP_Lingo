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

func (s Storage) RaiseGameState(gameID int, playerID int) bool {
	_, err := s.db.Query("UPDATE games SET state = state + 1 WHERE id = $1 AND player_id = $2", gameID, playerID)
	if err != nil {
		return false
	}
	return true
}

func (s Storage) RaiseTryCount(gameID int, playerID int) bool {
	_, err := s.db.Query("UPDATE games SET current_try = current_try + 1 WHERE id = $1 AND player_id = $2", gameID, playerID)
	if err != nil {
		return false
	}
	return true
}

func (s Storage) ResetTryCount(gameID int, playerID int) bool {
	_, err := s.db.Query("UPDATE games SET current_try = 0 WHERE id = $1 AND player_id = $2", gameID, playerID)
	if err != nil {
		return false
	}
	return true
}

func (s Storage) RaiseGameScore(gameID int, playerID int, score int) bool {
	_, err := s.db.Query("UPDATE games SET score = score + $1 WHERE id = $2 AND player_id = $3", score, gameID, playerID)
	if err != nil {
		return false
	}
	return true
}

func (s Storage) GetGameForID(gameID int, playerID int) (game.Game, error) {
	var game game.Game
	err := s.db.QueryRow(`
		SELECT 	id, 
		        player_id, 
		        five_letter_word, 
		        six_letter_word, 
		        seven_letter_word, 
		        state, 
		        time_epoch, 
		        score, 
		        current_try
		FROM games 
		WHERE id = $1 && player_id = $2`, gameID, playerID).Scan(&game.ID, &game.PlayerID, &game.FiveLetterWord, &game.SixLetterWord, &game.SevenLetterWord, &game.State, &game.Time, &game.Score, &game.CurrentTry)
	if err != nil {
		fmt.Println(err)
		return game, err
	}
	return game, nil
}

func (s Storage) GetCurrentGame(playerID int) (game.Game, error) {
	var game game.Game
	err := s.db.QueryRow(`
		SELECT	id, 
		        player_id, 
		        five_letter_word, 
		        six_letter_word, 
		        seven_letter_word, 
		        state, 
		        time_epoch, 
		        score, 
		        current_try
		FROM games 
		WHERE player_id = $1 
		ORDER BY id DESC 
		LIMIT 1`, playerID).Scan(&game.ID, &game.PlayerID, &game.FiveLetterWord, &game.SixLetterWord, &game.SevenLetterWord, &game.State, &game.Time, &game.Score, &game.CurrentTry)
	if err != nil {
		fmt.Println(err)
		return game, err
	}
	return game, nil
}
