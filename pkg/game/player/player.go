package player

import "errors"

type Player struct {
	UserName       string `json:"username"`
	Password       string `json:"password"`
	HashedPassword string
}

var ErrUerNotFound = errors.New("err, could not find user")
