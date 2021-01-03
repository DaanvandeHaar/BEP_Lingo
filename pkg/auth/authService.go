package auth

import (
	"BEP_Lingo/pkg/game/player"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

type Claims struct {
	Player string `json:"player"`
	Exp    int    `json:"exp"`
	jwt.StandardClaims
}

func GenerateJWT(player player.Player) string {
	var mySignInKey = []byte("f34a54b3c45ca43c05bb")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["player"] = player.UserName
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(mySignInKey)

	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}
func GetUsernameFromToken(tokenStr string) (string, error) {
	mySignInKey := []byte("f34a54b3c45ca43c05bb")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return mySignInKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["player"].(string)
	return userName, nil

}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var mySignInKey = []byte(os.Getenv("KEY"))
		var header = r.Header.Get("x-access-token")

		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)
		if header != "" {
			token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("JWT error")
				}
				return mySignInKey, nil
			})
			if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode("Error, Found token but expired or formatted incorrectly")
				return
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode("Error, not authorized")
				return
			}
		} else {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("Error, could not find JWT token")
			return
		}
	})
}
