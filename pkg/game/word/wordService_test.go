package word

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

func Test_service_CheckIfAlpha(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		word Word
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
			if got := s.CheckIfAlpha(tt.args.word); got != tt.want {
				t.Errorf("CheckIfAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CompareWords(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		word        string
		correctWord string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Try
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.CompareWords(tt.args.word, tt.args.correctWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetRandomWord(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		len int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.GetRandomWord(tt.args.len); got != tt.want {
				t.Errorf("GetRandomWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
