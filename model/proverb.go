package model

type Proverb struct {
	ID     string `json:"id"`
	Riddle string `json:"riddle"`
	Answer string `json:"answer"`
}

type Proverbs []Proverb
