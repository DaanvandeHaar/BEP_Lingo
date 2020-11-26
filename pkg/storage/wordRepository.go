package storage

import (
	_ "github.com/lib/pq"
	"log"
)

func (s Storage) GetRandomWord(length int) string {
	type Result struct {
		Word string
	}

	var result Result

	rows, err := s.db.Query("SELECT * FROM words WHERE length(word) = $1 OFFSET random() * (SELECT count(*) FROM words WHERE length(word) = $1 )  LIMIT 1 ;", length)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&result.Word); err != nil {
			log.Fatal(err)
		} else {

		}
	}
	return result.Word
}
