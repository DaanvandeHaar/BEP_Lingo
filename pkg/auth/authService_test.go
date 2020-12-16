package auth

import (
	"BEP_Lingo/pkg/game/player"
	"reflect"
	"testing"
)

type mockingStorage struct {
}

func TestGenerateJWT(t *testing.T) {
	type args struct {
		player player.Player
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST_GENERATE_JWT_PASS",
			args: args{player: player.Player{
				UserName:       "daan",
				Password:       "daan",
				HashedPassword: "",
			}},
		},
		{
			name: "TEST_GENERATE_JWT_FAIL",
			args: args{player: player.Player{
				UserName:       "",
				Password:       "",
				HashedPassword: "",
			}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateJWT(tt.args.player); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
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
		{
			name:    "TEST_GET_USERNAME_FROM_TOKEN_PASS",
			args:    args{tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDQxMzM1MDgsInBsYXllciI6ImRhYW4ifQ.vBE5QmzVfWLFdcRarElOYdweJjXkqhOtB_rquIhFNpg"},
			want:    "daan",
			wantErr: false,
		},
		{
			name:    "TEST_GET_USERNAME_FROM_TOKEN_PASS",
			args:    args{tokenStr: "eyJhbGciOiJIUzI1NiIsInRcCI6IkpXVCJ9.eyJhdXRoJpemVkIjp0cnVlLCJleHAiOjE2NDQxMzM1MDgsInBsYXllciI6ImRhYW4ifQ.vBE5QmzVfWLFdcRarElOYdweJjXkqhOtB_rquIhFNpg"},
			wantErr: true,
		},
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
