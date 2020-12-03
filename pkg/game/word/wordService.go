package word

import (
	"fmt"
	"strings"
)

type Service interface {
	CheckIfAlpha(Word) bool
	CompareWords(string, string) Try
	GetRandomWord(int) string
}
type Repository interface {
	GetRandomWord(int) string
}
type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetRandomWord(len int) string {
	return s.r.GetRandomWord(len)
}

func (s *service) CheckIfAlpha(word Word) bool {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	for _, char := range word.Word {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}

	return true
}
func (s *service) CompareWords(word string, correctWord string) Try {
	var try Try
	for pos, char := range word {
		if correctWord[pos] == word[pos] {
			fmt.Println(true, string(char))
			try.Letters = append(try.Letters, Letter{
				LetterString:   string(char),
				LetterPosition: pos,
				RightPlace:     true,
				RightLetter:    true,
			})
		} else {
			fmt.Println(false, string(char))
			if strings.ContainsAny(correctWord, string(char)) {
				try.Letters = append(try.Letters, Letter{
					LetterString:   string(char),
					LetterPosition: pos,
					RightPlace:     false,
					RightLetter:    true,
				})
			} else {
				try.Letters = append(try.Letters, Letter{
					LetterString:   string(char),
					LetterPosition: pos,
					RightPlace:     false,
					RightLetter:    false,
				})
			}
		}
	}
	return try
}
