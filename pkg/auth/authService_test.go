package auth

import (
	"BEP_Lingo/pkg/game/player"
	"net/http"
	"reflect"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	type args struct {
		player player.Player
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateJWT(tt.args.player); got != tt.want {
				t.Errorf("GenerateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsernameFromToken(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsernameFromToken(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsernameFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUsernameFromToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtVerify(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JwtVerify(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JwtVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}
