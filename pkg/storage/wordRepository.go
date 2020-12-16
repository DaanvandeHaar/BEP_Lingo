package storage

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

func (s Storage) GetRandomWord(length int) string {
	type Result struct {
		Word string
	}
	var result Result
	err := s.db.QueryRow("SELECT * FROM words WHERE length(word) = $1 OFFSET random() * (SELECT count(*) FROM words WHERE length(word) = $1 )  LIMIT 1 ;", length).Scan(&result.Word)
	if err != nil {
		fmt.Println(errors.New("Error, could not get word from database"))
	}
	return result.Word
}
