package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {

	const (
		dbUser					= "postgres"
		dbPwd 					= "admin"
		instanceConnectionName			= "bep-lingo:europe-west1:lingodb"
		dbName					= "lingo_db"
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil
	}

	return &Storage{db}
}
