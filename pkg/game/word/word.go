package word

import "errors"

type Word struct {
	Word string `json:"word"`
}

type LingoMessage struct {
	TryIndex int    `json:"index"`
	Info     string `json:"info"`
	Correct  bool   `json:"correct"`
	Letters  []LetterInfo
}

type LetterInfo struct {
	LetterString   string `json:"letterString"`
	LetterPosition int    `json:"letterPosition"`
	RightPlace     bool   `json:"rightPlace"`
	RightLetter    bool   `json:"rightLetter"`
}

var ErrorNonValidWord = errors.New("Error, words can only contain lower case letters. must contain the correct amount of characters")
