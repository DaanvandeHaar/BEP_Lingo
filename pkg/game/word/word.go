package word

type Word struct {
	Word string `json:"word"`
}

type Try struct {
	tryIndex int
	Letters  []Letter
}

type Letter struct {
	LetterString   string `json:"letterString"`
	LetterPosition int    `json:"letterPosition"`
	RightPlace     bool   `json:"rightPlace"`
	RightLetter    bool   `json:"rightLetter"`
}

type WordInterface interface {
	CheckIfAlpha(Word) (bool, error)
	CompareWords(string, string) (Try, error)
}
