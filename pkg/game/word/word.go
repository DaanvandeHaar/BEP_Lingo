package word

type Word struct {
	Word string `json:"word"`
}

type LingoMessage struct {
	tryIndex int
	Letters  []LetterInfo
}

type LetterInfo struct {
	LetterString   string `json:"letterString"`
	LetterPosition int    `json:"letterPosition"`
	RightPlace     bool   `json:"rightPlace"`
	RightLetter    bool   `json:"rightLetter"`
}
