package game

import (
	"reflect"
	"testing"
	"time"
)

type mockStorage struct {
	game []Game
}

func (m *mockStorage) ResetTryCount(i int, i2 int) bool {
	panic("implement me")
}

func (m *mockStorage) GetCurrentGame(i int) (Game, error) {
	panic("implement me")
}

func (m *mockStorage) GetGameForID(i int, i2 int) (Game, error) {
	panic("implement me")
}

func (m *mockStorage) NewGame(game Game) (int, error) {
	return 1, nil
}

func (m *mockStorage) RaiseGameState(gameID int, playerID int) bool {
	switch {
	case &gameID == nil || &playerID == nil:
		return false
	case gameID == 0 && playerID == 0:
		return false
	case gameID == 1 && playerID == 0:
		return false
	case gameID == 0 && playerID == 1:
		return false
	case gameID == 1 && playerID == 1:
		return true
	case gameID == 2 && playerID == 2:
		return true
	}
	return false
}

func (m *mockStorage) RaiseTryCount(gameID int, playerID int) bool {
	if gameID != 0 && playerID != 0 {
		return true
	}
	return false
}

func (m *mockStorage) RaiseGameScore(gameID int, playerID int, score int) bool {
	if gameID != 0 && playerID != 0 {
		return true
	}
	return false
}

func Test_service_InitGame(t *testing.T) {
	type fields struct {
		r Repository
	}
	mR := new(mockStorage)
	type args struct {
		words    []string
		playerID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Game
		wantErr bool
	}{
		{
			name:   "INIT_GAME_PASS",
			fields: fields{mR},
			args: args{
				words:    []string{"knoop", "schiet", "schepen"},
				playerID: 1,
			},
			want: Game{
				ID:              1,
				PlayerID:        1,
				State:           0,
				CurrentTry:      0,
				FiveLetterWord:  "knoop",
				SixLetterWord:   "schiet",
				SevenLetterWord: "schepen",
				Score:           0,
				Time:            time.Now().Unix(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, err := s.InitGame(tt.args.words, tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitGame() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_RaiseGameScore(t *testing.T) {
	type fields struct {
		r Repository
	}
	mR := new(mockStorage)
	type args struct {
		gameID   int
		playerID int
		score    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "RAISE_GAME_SCORE_PASS",
			fields: fields{mR},
			args: args{
				gameID:   1,
				playerID: 1,
				score:    10,
			},
			want: true,
		},
		{
			name:   "RAISE_GAME_SCORE_FAIL",
			fields: fields{mR},
			args: args{
				gameID:   0,
				playerID: 0,
				score:    10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.RaiseGameScore(tt.args.gameID, tt.args.playerID, tt.args.score); got != tt.want {
				t.Errorf("RaiseGameScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_RaiseGameState(t *testing.T) {
	type fields struct {
		r Repository
	}
	mR := new(mockStorage)
	type args struct {
		gameID   int
		playerID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "RAISE_GAME_STATE_PASS",
			fields: fields{mR},
			args: args{
				gameID:   1,
				playerID: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.RaiseGameState(tt.args.gameID, tt.args.playerID); got != tt.want {
				t.Errorf("RaiseGameState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *service) Test_service_RaiseTryCount(t *testing.T) {
	type fields struct {
		r Repository
	}
	mR := new(mockStorage)

	type args struct {
		gameID   int
		playerID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "TEST_RAISE_TRY_COUNT_FAIL",
			fields: fields{mR},
			args: args{
				gameID:   5,
				playerID: 2,
			},
			want: false,
		},
		{
			name:   "TEST_RAISE_TRY_COUNT_PASS",
			fields: fields{mR},
			args: args{
				gameID:   1,
				playerID: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.RaiseTryCount(tt.args.gameID, tt.args.playerID); got != tt.want {
				t.Errorf("RaiseTryCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
