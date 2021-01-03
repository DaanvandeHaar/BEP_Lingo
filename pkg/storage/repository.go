package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {

	var (
		dbUser                 = os.Getenv("USER")
		dbPwd                  = os.Getenv("PASSWORD")
		instanceConnectionName = os.Getenv("CONNECTION_NAME")
		dbName                 = os.Getenv("DB_NAME")
	)

	socketDir := "/cloudsql"

	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil
	}

	return &Storage{db}
}
