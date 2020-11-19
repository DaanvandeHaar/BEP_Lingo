package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "admin"
		password = "admin"
		dbname   = "lingo_db"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d player=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
func GetRandomWord(length int) string {
	type Result struct {
		Word string
	}
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "admin"
		dbname   = "lingo_db"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d player=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result Result

	rows, err := db.Query("SELECT * FROM words WHERE length(word) = $1 OFFSET random() * (SELECT count(*) FROM words WHERE length(word) = $1 )  LIMIT 1 ;", length)
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
