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
type mockinWordStorage struct {
	word []word.Word
}

func (m mockinWordStorage) GetRandomWord(i int) string {
	switch {
	case i < 5:
		return ""
	case i == 5:
		return "woord"
	case i == 6:
		return "burger"
	case i == 7:
		return "knuffel"
	case i > 7:
		return ""
	}
	return ""
}

func (m *mockStorage) ResetTryCount(playerID int, gameID int) bool {
	if playerID == gameID && playerID != 0 {
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
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 3:
		return Game{
			ID:              3,
			PlayerID:        3,
			State:           1,
			CurrentTry:      6,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 4:
		return Game{
			ID:              4,
			PlayerID:        4,
			State:           2,
			CurrentTry:      0,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 5:
		return Game{
			ID:              5,
			PlayerID:        5,
			State:           2,
			CurrentTry:      1,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 6:
		return Game{
			ID:              6,
			PlayerID:        6,
			State:           2,
			CurrentTry:      6,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 7:
		return Game{
			ID:              7,
			PlayerID:        7,
			State:           3,
			CurrentTry:      0,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 8:
		return Game{
			ID:              8,
			PlayerID:        8,
			State:           3,
			CurrentTry:      1,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 9:
		return Game{
			ID:              9,
			PlayerID:        9,
			State:           3,
			CurrentTry:      6,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
			Score:           0,
			Time:            0,
		}, nil
	case 10:
		return Game{
			ID:              10,
			PlayerID:        10,
			State:           4,
			CurrentTry:      0,
			FiveLetterWord:  "woord",
			SixLetterWord:   "pandas",
			SevenLetterWord: "stoelen",
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
			State:           1,
			CurrentTry:      0,
			FiveLetterWord:  "wraps",
			SixLetterWord:   "dingen",
			SevenLetterWord: "woorden",
			Score:           0,
			Time:            0,
		}, nil
	case 3:
		return Game{
			ID:              3,
			PlayerID:        3,
			State:           1,
			CurrentTry:      6,
			FiveLetterWord:  "wraps",
			SixLetterWord:   "dingen",
			SevenLetterWord: "woorden",
			Score:           0,
			Time:            0,
		}, nil
	case 4:
		return Game{
			ID:              4,
			PlayerID:        4,
			State:           2,
			CurrentTry:      0,
			FiveLetterWord:  "wraps",
			SixLetterWord:   "dingen",
			SevenLetterWord: "woorden",
			Score:           0,
			Time:            0,
		}, nil
	case 5:
		return Game{
			ID:              5,
			PlayerID:        5,
			State:           2,
			CurrentTry:      1,
			FiveLetterWord:  "wraps",
			SixLetterWord:   "dingen",
			SevenLetterWord: "woorden",
			Score:           0,
			Time:            0,
		}, nil
	case 6:
		return Game{
			ID:              6,
			PlayerID:        6,
			State:           2,
			CurrentTry:      1,
			FiveLetterWord:  "wraps",
			SixLetterWord:   "dingen",
			SevenLetterWord: "woorden",
			Score:           0,
			Time:            0,
		}, nil
	case 7:
		return Game{
			ID:              7,
			PlayerID:        7,
			State:           2,
			CurrentTry:      6,
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
	case gameID == 1 && playerID == 1:
		return true
	case gameID == 2 && playerID == 2:
		return true
	case gameID == 3 && playerID == 3:
		return true
	case gameID == 4 && playerID == 4:
		return true
	case gameID == 5 && playerID == 5:
		return true
	case gameID == 6 && playerID == 6:
		return true
	case gameID == 7 && playerID == 7:
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
				words:    []string{"woord", "pandas", "stoelen"},
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
			name:   "RAISE_GAME_STATE_FAIL_1",
			fields: fields{mR},
			args: args{
				gameID:   0,
				playerID: 0,
			},
			want: false,
		},
		{
			name:   "RAISE_GAME_STATE_FAIL_2",
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
				FiveLetterWord:  "woord",
				SixLetterWord:   "pandas",
				SevenLetterWord: "stoelen",
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
	mR := new(mockStorage)
	mwR := new(mockinWordStorage)
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
		{
			name:   "TEST_GAME_RUNNER_PASS_1",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "woord",
				playerID: 1,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "w____",
				Correct:  false,
				Letters:  nil,
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_2",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "woord",
				playerID: 2,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "You have guessed the word correctly. Continuing to 6 letter game stage. Hint for 6 letter word: p_____",
				Correct:  true,
				Letters: []word.LetterInfo{{
					LetterString:   "w",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "o",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "o",
					LetterPosition: 2,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "r",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "d",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_3",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "waard",
				playerID: 2,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Correct:  false,
				Letters: []word.LetterInfo{{
					LetterString:   "w",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "a",
					LetterPosition: 1,
					RightPlace:     false,
					RightLetter:    false,
				}, {
					LetterString:   "a",
					LetterPosition: 2,
					RightPlace:     false,
					RightLetter:    false,
				}, {
					LetterString:   "r",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "d",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_4",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "waard",
				playerID: 2,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Correct:  false,
				Letters: []word.LetterInfo{{
					LetterString:   "w",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "a",
					LetterPosition: 1,
					RightPlace:     false,
					RightLetter:    false,
				}, {
					LetterString:   "a",
					LetterPosition: 2,
					RightPlace:     false,
					RightLetter:    false,
				}, {
					LetterString:   "r",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "d",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_4",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "waard",
				playerID: 3,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "You have not guessed the correct word in time. Continuing to 6 letter game stage. Hint for 6 letter word: p_____",
				Correct:  false,
				Letters:  nil,
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_5",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "pandas",
				playerID: 4,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "You have guessed the word correctly. Continuing to 7 letter game stage. Hint for 7 letter word: s______",
				Correct:  true,
				Letters: []word.LetterInfo{{
					LetterString:   "p",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "a",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "n",
					LetterPosition: 2,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "d",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "a",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "s",
					LetterPosition: 5,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_5",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "pondas",
				playerID: 4,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Correct:  false,
				Letters: []word.LetterInfo{{
					LetterString:   "p",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "o",
					LetterPosition: 1,
					RightPlace:     false,
					RightLetter:    false,
				}, {
					LetterString:   "n",
					LetterPosition: 2,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "d",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "a",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "s",
					LetterPosition: 5,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_6",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "pandas",
				playerID: 5,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "You have guessed the word correctly. Continuing to 7 letter game stage. Hint for 7 letter word: s______",
				Correct:  true,
				Letters: []word.LetterInfo{{
					LetterString:   "p",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "a",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "n",
					LetterPosition: 2,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "d",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "a",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "s",
					LetterPosition: 5,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_7",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "pandas",
				playerID: 6,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "You have not guessed the correct word in time. Continuing to 7 letter game stage. Hint for 7 letter word: s______",
				Correct:  false,
				Letters:  nil,
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_8",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "stoelen",
				playerID: 7,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "You have guessed the word correctly. The game is finished",
				Correct:  true,
				Letters: []word.LetterInfo{{
					LetterString:   "s",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "t",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "o",
					LetterPosition: 2,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "l",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 5,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "n",
					LetterPosition: 6,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_8",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "steelen",
				playerID: 7,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Correct:  false,
				Letters: []word.LetterInfo{{
					LetterString:   "s",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "t",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 2,
					RightPlace:     false,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "l",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 5,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "n",
					LetterPosition: 6,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_8",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "steelen",
				playerID: 8,
			},
			want: word.LingoMessage{
				TryIndex: 1,
				Correct:  false,
				Letters: []word.LetterInfo{{
					LetterString:   "s",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "t",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 2,
					RightPlace:     false,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "l",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "e",
					LetterPosition: 5,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "n",
					LetterPosition: 6,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_9",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "steelen",
				playerID: 9,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "You have not guessed the correct word in time. The game is finished",
				Correct:  false,
				Letters:  nil,
			},
			wantErr: false,
		},
		{
			name:   "TEST_GAME_RUNNER_PASS_10",
			fields: fields{mR},
			args: args{
				ws:       word.NewService(mwR),
				word:     "steelen",
				playerID: 10,
			},
			want: word.LingoMessage{
				TryIndex: 0,
				Info:     "This game is already finished!",
				Correct:  false,
				Letters:  nil,
			},
			wantErr: false,
		},
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

func TestRaiseTryCount(t *testing.T) {
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
