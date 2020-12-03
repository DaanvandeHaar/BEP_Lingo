package storage

import (
	"awesomeProject/pkg/auth"
	"awesomeProject/pkg/game/player"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
func (s Storage) GetIDForPlayer(username string) (int, error) {
	var id int
	err := s.db.QueryRow("SELECT id FROM users where username = $1", username).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (s Storage) LoginWithHash(player player.Player) (bool, string) {

	var databaseUsername string
	var databasePassword string

	err := s.db.QueryRow("SELECT username, password FROM users WHERE username=$1;", player.UserName).Scan(&databaseUsername, &databasePassword)
	if sql.ErrNoRows == nil {
		fmt.Println(err)
		return false, ""
	}
	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(player.Password))
	if err != nil {
		return false, ""
	}
	return true, auth.GenerateJWT(player)

}

func (s Storage) SignUpWithHash(player player.Player) bool {
	var user string
	err := s.db.QueryRow("SELECT username FROM users WHERE username=$1", player.UserName).Scan(&user)
	switch {
	case err != nil:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			return false
		}
		_, err = s.db.Exec("INSERT INTO users(username, password) VALUES($1, $2)", player.UserName, string(hashedPassword))
		if err != nil {
			fmt.Println(err)
			return false
		}
	default:
		return false
	}
	return false
}
