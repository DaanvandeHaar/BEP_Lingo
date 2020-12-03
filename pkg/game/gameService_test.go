package game

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		r Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_InitGame(t *testing.T) {
	type fields struct {
		r Repository
	}
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.RaiseGameScore(tt.args.gameID, tt.args.playerID); got != tt.want {
				t.Errorf("RaiseGameScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_RaiseGameState(t *testing.T) {
	type fields struct {
		r Repository
	}
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
		// TODO: Add test cases.
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

func Test_service_RaiseTryCount(t *testing.T) {
	type fields struct {
		r Repository
	}
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
		// TODO: Add test cases.
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
