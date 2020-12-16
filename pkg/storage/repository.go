package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {

	//host     := os.Getenv("HOST")
	//port, err := strconv.Atoi(os.Getenv("PORT"))
	//user     := os.Getenv("USER")
	//password := os.Getenv("PASSWORD")
	//dbname   := os.Getenv("DB_NAME")

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "admin"
		dbname   = "lingo_db"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	return &Storage{db}
}
