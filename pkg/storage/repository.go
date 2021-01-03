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

	socketDir := "/cloudsql"

	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	db, err := sql.Open("postges", dbURI)
	if err != nil {
		return nil
	}

	return &Storage{db}
}
