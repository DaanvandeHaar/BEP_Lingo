package player

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(Player) (bool, string)
	SignUp(Player) bool
}
type Repository interface {
	LoginWithHash(Player) (bool, string)
	SignUpWithHash(Player) bool
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Login(player Player) (bool, string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
	if err != nil {
		fmt.Print(err)
	}
	player.HashedPassword = string(hashedPassword)
	valid, token := s.r.LoginWithHash(player)

	if valid == false {
		fmt.Println(err)
		return false, "err: Wrong username or passowrd"
	} else {
		return true, token
	}

}

func (s *service) SignUp(player Player) bool {
	valid := s.r.SignUpWithHash(player)
	if valid == false {
		fmt.Println("Error at signUp service")
		return false
	}
	return true

}
