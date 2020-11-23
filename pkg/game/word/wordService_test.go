package word

import (
	"reflect"
	"testing"
)

func TestCheckIfAlpha(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIfAlpha(tt.args.s); got != tt.want {
				t.Errorf("CheckIfAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareWords(t *testing.T) {
	type args struct {
		word        string
		correctWord string
	}
	tests := []struct {
		name string
		args args
		want Try
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareWords(tt.args.word, tt.args.correctWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
