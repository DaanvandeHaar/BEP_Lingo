package auth

import (
	"awesomeProject/pkg/game/player"
	"encoding/json"
	"net/http"
	"strings"

	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(player player.Player) string {
	var mySignInKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["player"] = player.UserName
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignInKey)

	if err != nil {
		fmt.Errorf("something went wrong: %s", err)
	}
	return tokenString
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var mySignInKey = []byte("secretkey")
		var header = r.Header.Get("x-access-token")

		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header != "" {
			token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
				if err, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("JWT error: %e", err)
				}
				return mySignInKey, nil
			})
			if err != nil {
				fmt.Errorf(err.Error())
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				json.NewEncoder(w).Encode("Error, not authorized")
			}
		} else {
			json.NewEncoder(w).Encode("Error, could not find JWT token")
		}
	})
}
