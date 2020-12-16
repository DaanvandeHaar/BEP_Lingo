package word

import (
	"fmt"
	"strings"
)

type Service interface {
	CheckIfAlpha(string) bool
	CompareWords(string, string) (LingoMessage, error)
	GetRandomWord(int) string
	GetWordHelp(string) string
	GetEmptyMessage() LingoMessage
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

func (s *service) GetWordHelp(word string) string {
	firstChar := word[0]
	redactedString := strings.Repeat("_", len(word)-1)
	return string(firstChar) + redactedString
}

func (s *service) GetRandomWord(len int) string {
	return s.r.GetRandomWord(len)
}

func (s *service) GetEmptyMessage() LingoMessage {
	return LingoMessage{}
}

func (s *service) CheckIfAlpha(word string) bool {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	for _, char := range word {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}

	return true
}
func (s *service) CompareWords(word string, correctWord string) (LingoMessage, error) {
	var message LingoMessage
	if !s.CheckIfAlpha(word) || len(word) != len(correctWord) {
		return LingoMessage{}, ErrorNonValidWord
	}
	if word == correctWord {
		message.Correct = true
	}
	for pos, char := range word {
		if correctWord[pos] == word[pos] {
			fmt.Println(true, string(char))
			message.Letters = append(message.Letters, LetterInfo{
				LetterString:   string(char),
				LetterPosition: pos,
				RightPlace:     true,
				RightLetter:    true,
			})
		} else {
			fmt.Println(false, string(char))
			if strings.ContainsAny(correctWord, string(char)) {
				message.Letters = append(message.Letters, LetterInfo{
					LetterString:   string(char),
					LetterPosition: pos,
					RightPlace:     false,
					RightLetter:    true,
				})
			} else {
				message.Letters = append(message.Letters, LetterInfo{
					LetterString:   string(char),
					LetterPosition: pos,
					RightPlace:     false,
					RightLetter:    false,
				})
			}
		}
	}
	return message, nil
}
