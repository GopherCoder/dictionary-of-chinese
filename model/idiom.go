package model

type Idiom struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	PinYin  string `json:"pinyin"`
	Explain string `json:"explain"`
	Source  string `json:"from"`
	Example string `json:"example"`
}

type Idioms []Idiom
