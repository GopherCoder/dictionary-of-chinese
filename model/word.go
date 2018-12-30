package model

type Word struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Explain string `json:"explain"`
}

type Words []Word
