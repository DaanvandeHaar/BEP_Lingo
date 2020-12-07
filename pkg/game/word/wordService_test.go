package word

import (
	"reflect"
	"testing"
)

type mockingStorage struct {
	word []Word
}

func (m mockingStorage) GetRandomWord(i int) string {
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
		{
			name:   "CHECK_IF_ALPHA_PASS",
			fields: fields{},
			args:   args{word: Word{"adfjkldfjkaigii"}},
			want:   true,
		},
		{
			name:   "CHECK_IF_ALPHA_FAIL",
			fields: fields{},
			args:   args{word: Word{"dkln8ddd"}},
			want:   false,
		},
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
	mR := new(mockingStorage)
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
		{
			name:   "COMPARE_WORDS_PASS",
			fields: fields{mR},
			args: args{
				word:        "knoop",
				correctWord: "knoop",
			},
			want: Try{
				tryIndex: 0,
				Letters: []Letter{{
					LetterString:   "k",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "n",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "o",
					LetterPosition: 2,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "o",
					LetterPosition: 3,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "p",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
		},
		{
			name:   "COMPARE_WORDS_FAIL",
			fields: fields{mR},
			args: args{
				word:        "knoop",
				correctWord: "kneep",
			},
			want: Try{
				tryIndex: 0,
				Letters: []Letter{{
					LetterString:   "k",
					LetterPosition: 0,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "n",
					LetterPosition: 1,
					RightPlace:     true,
					RightLetter:    true,
				}, {
					LetterString:   "o",
					LetterPosition: 2,
					RightPlace:     false,
					RightLetter:    false,
				}, {
					LetterString:   "o",
					LetterPosition: 3,
					RightPlace:     false,
					RightLetter:    false,
				}, {
					LetterString:   "p",
					LetterPosition: 4,
					RightPlace:     true,
					RightLetter:    true,
				}},
			},
		},
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
	mR := new(mockingStorage)
	type args struct {
		len int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "GET_RANDOM_WORD_PASS_1",
			fields: fields{mR},
			args:   args{5},
			want:   "woord",
		},
		{
			name:   "GET_RANDOM_WORD_PASS_2",
			fields: fields{mR},
			args:   args{7},
			want:   "knuffel",
		},
		{
			name:   "GET_RANDOM_WORD_FAIL_1",
			fields: fields{mR},
			args:   args{4},
			want:   "",
		},
		{
			name:   "GET_RANDOM_WORD_FAIL_2",
			fields: fields{mR},
			args:   args{8},
			want:   "",
		},
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
