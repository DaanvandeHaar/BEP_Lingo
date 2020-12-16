package player

import (
	"errors"
	"reflect"
	"testing"
)

type mockStorage struct {
}

func (m mockStorage) LoginWithHash(player Player) (bool, string) {
	switch player.UserName {
	case "daan":
		return true, "daantoken"
	case "henk":
		return true, "henktoken"
	case "frans":
		return false, ""

	}
	return false, ""
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
		return 0, ErrUserNotFound
	case s == "":
		return 0, errors.New("err, name field empty")
	}
	return 0, errors.New("err, other error")
}

//func TestLogin(t *testing.T) {
//	type fields struct {
//		r Repository
//	}
//	mR := new(mockStorage)
//	type args struct {
//		player Player
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//		want1  string
//	}{
//		{
//			name:   "LOGIN_TEST_PASS",
//			fields: fields{mR},
//			args: args{player: Player{
//				UserName: "daan",
//				Password: "daan",
//			},
//			},
//			want: true,
//			want1: "daantoken",
//		},
//		{
//			name:   "LOGIN_TEST_FAIL_1",
//			fields: fields{},
//			args: args{player: Player{
//				UserName: "henk",
//				Password: "henk",
//			},
//			},
//			want: false,
//		},
//		{
//			name:   "LOGIN_TEST_FAIL_2",
//			fields: fields{},
//			args:   args{player: Player{}},
//			want:   false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &service{
//				r: tt.fields.r,
//			}
//			got, got1 := s.Login(tt.args.player)
//			if got != tt.want {
//				t.Errorf("Login() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Login() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}

func TestSignUp(t *testing.T) {
	type fields struct {
		r Repository
	}
	mR := new(mockStorage)
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
			fields: fields{mR},
			args: args{player: Player{
				UserName: "daan",
				Password: "daan",
			},
			},
			want: true,
		},
		{
			name:   "SIGNUP_TEST_FAIL",
			fields: fields{mR},
			args: args{
				player: Player{
					UserName:       "daan",
					Password:       "henk",
					HashedPassword: "",
				},
			},
			want: false,
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

func TestNewService(t *testing.T) {
	type args struct {
		r Repository
	}
	mR := new(mockStorage)
	tests := []struct {
		name string
		args args
		want service
	}{
		{
			name: "TEST_NEW_SERVICE_PASS",
			args: args{mR},
			want: service{
				r: mR,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetIDForPlayer(t *testing.T) {
	type fields struct {
		r Repository
	}
	mR := new(mockStorage)
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "TEST_GET_ID_FOR_PLAYER_PASS_1",
			fields: fields{mR},
			args: args{
				username: "daan",
			},
			want:    1,
			wantErr: false,
		},
		{
			name:   "TEST_GET_ID_FOR_PLAYER_PASS_2",
			fields: fields{mR},
			args: args{
				username: "henk",
			},
			want:    2,
			wantErr: false,
		},
		{
			name:   "TEST_GET_ID_FOR_PLAYER_FAIL_1",
			fields: fields{mR},
			args: args{
				username: "frank",
			},
			want:    0,
			wantErr: true,
		},
		{
			name:   "TEST_GET_ID_FOR_PLAYER_FAIL_2",
			fields: fields{mR},
			args: args{
				username: "daan",
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, err := s.GetIDForPlayer(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIDForPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIDForPlayer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
