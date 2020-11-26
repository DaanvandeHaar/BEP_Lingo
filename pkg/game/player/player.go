package player

type Player struct {
	UserName       string `json:"username"`
	Password       string `json:"password"`
	HashedPassword string
}
