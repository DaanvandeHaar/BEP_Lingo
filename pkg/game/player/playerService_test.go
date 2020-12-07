package player

import (
	"errors"
	"testing"
)

type mockStorage struct {
	player Player
}

func (m mockStorage) LoginWithHash(player Player) (bool, string) {
	if (Player{}) == player {
		return false, ""
	}
	return true, "iojdfainonidfdfaeatf"
}

func (m mockStorage) SignUpWithHash(player Player) bool {
	if (Player{}) == player {
		return false
	}
	if player.Password == "daan" && player.UserName == "daan" {
		return true
	}
	return false
}

func (m mockStorage) GetIDForPlayer(s string) (int, error) {
	switch {
	case s == "daan":
		return 1, nil
	case s == "henk":
		return 2, nil
	case s == "frank":
		return 0, ErrUerNotFound
	case s == "":
		return 0, errors.New("err, name field empty")
	}
	return 0, errors.New("err, other error")
}

func (s *service) Test_service_Login(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		player Player
	}
	mR := new(mockStorage)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  string
	}{
		{
			name:   "LOGIN_TEST_PASS",
			fields: fields{mR},
			args: args{player: Player{
				UserName: "daan",
				Password: "daan",
			},
			},
			want: true,
		},
		{
			name:   "LOGIN_TEST_FAIL_1",
			fields: fields{},
			args: args{player: Player{
				UserName: "fail",
				Password: "fail",
			},
			},
			want: false,
		},
		{
			name:   "LOGIN_TEST_FAIL_2",
			fields: fields{},
			args:   args{player: Player{}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, got1 := s.Login(tt.args.player)
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func (s *service) Test_service_SignUp(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		player Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "SIGNUP_TEST_PASS",
			fields: fields{s.r},
			args: args{player: Player{
				UserName: "test",
				Password: "test",
			},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.SignUp(tt.args.player); got != tt.want {
				t.Errorf("SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
