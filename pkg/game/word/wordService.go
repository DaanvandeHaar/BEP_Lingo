package word

import (
	"fmt"
	"strings"
)

func CheckIfAlpha(s string) bool {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
func CompareWords(word string, correctWord string) Try {
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
