package game

import (
	"BEP_Lingo/pkg/game/word"
	"reflect"
	"testing"
	"time"
)

type mockStorage struct {
	game []Game
}

func (m *mockStorage) ResetTryCount(playerID int, gameID int) bool {
	if playerID == 1 && gameID == 1 {
		return true
	}
	return false
}

func (m *mockStorage) GetCurrentGame(playerID int) (Game, error) {
	switch playerID {
	case 1:
		return Game{
			ID:              1,
			PlayerID:        1,
			State:           0,
			CurrentTry:      0,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 2:
		return Game{
			ID:              2,
			PlayerID:        2,
			State:           1,
			CurrentTry:      0,
			FiveLetterWord:  "wraps",
			SixLetterWord:   "dingen",
			SevenLetterWord: "woorden",
			Score:           0,
			Time:            0,
		}, nil
	}
	return Game{}, nil
}

func (m *mockStorage) GetGameForID(playerID int, gameID int) (Game, error) {
	switch playerID {
	case 1:
		return Game{
			ID:              1,
			PlayerID:        1,
			State:           0,
			CurrentTry:      0,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 2:
		return Game{
			ID:              2,
			PlayerID:        2,
			State:           0,
			CurrentTry:      0,
			FiveLetterWord:  "wraps",
			SixLetterWord:   "dingen",
			SevenLetterWord: "woorden",
			Score:           0,
			Time:            0,
		}, nil
	}
	return Game{}, nil
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

func TestInitGame(t *testing.T) {
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

func TestRaiseGameScore(t *testing.T) {
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
		{
			name:   "RAISE_GAME_SCORE_FAIL_2",
			fields: fields{nil},
			args: args{
				gameID:   0,
				playerID: 0,
				score:    0,
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

func TestRaiseGameState(t *testing.T) {
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
		{
			name:   "RAISE_GAME_STATE_PASS",
			fields: fields{mR},
			args: args{
				gameID:   0,
				playerID: 0,
			},
			want: false,
		},
		{
			name:   "RAISE_GAME_STATE_PASS",
			fields: fields{nil},
			args: args{
				gameID:   0,
				playerID: 0,
			},
			want: false,
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

func TestGetCurrentGame(t *testing.T) {
	type fields struct {
		r Repository
	}
	mR := new(mockStorage)
	type args struct {
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
			name:   "TEST_GET_CURRENT_GAME_PASS",
			fields: fields{mR},
			args: args{
				playerID: 1,
			},
			want: Game{
				ID:              1,
				PlayerID:        1,
				State:           0,
				CurrentTry:      0,
				FiveLetterWord:  "woord",
				SixLetterWord:   "pandas",
				SevenLetterWord: "stoelen",
				Score:           0,
				Time:            0,
			},
			wantErr: false,
		},
		{
			name:   "TEST_GET_CURRENT_GAME_PASS_2",
			fields: fields{mR},
			args: args{
				playerID: 2,
			},
			want: Game{
				ID:              2,
				PlayerID:        2,
				State:           1,
				CurrentTry:      0,
				FiveLetterWord:  "wraps",
				SixLetterWord:   "dingen",
				SevenLetterWord: "woorden",
				Score:           0,
				Time:            0,
			},
			wantErr: false,
		},
		{
			name: "TEST_GET_CURRENT_GAME_FAIL",
			fields: fields{
				r: mR,
			},
			args: args{
				playerID: 0,
			},
			want:    Game{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, err := s.GetCurrentGame(tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentGame() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameRunner(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		ws       word.Service
		word     string
		playerID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    word.LingoMessage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, err := s.GameRunner(tt.args.ws, tt.args.word, tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameRunner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameRunner() got = %v, want %v", got, tt.want)
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
			args: args{
				r: mR,
			},
			want: service{mR},
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

func Test_service_RaiseTryCount(t *testing.T) {
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
				gameID:   0,
				playerID: 0,
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
		{
			name:   "TEST_RAISE_TRY_COUNT_PASS",
			fields: fields{nil},
			args: args{
				gameID:   1,
				playerID: 1,
			},
			want: false,
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
