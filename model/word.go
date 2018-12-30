package model

type Word struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Explain string `json:"explain"`
}

type Words []Word

type Serialize struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Explain string `json:"explain"`
}

func (w *Word) BasicSerialize() Serialize {
	return Serialize{
		ID:      w.ID,
		Name:    w.Name,
		Explain: w.Explain,
	}
}
