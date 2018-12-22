package model

type Idiom struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Explain string `json:"explain"`
	Source  string `json:"source"`
	Example string `json:"example"`
}

type Idioms []Idiom
