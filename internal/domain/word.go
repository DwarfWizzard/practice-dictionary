package domain

type Word struct {
	Original string `json:"original" db:"original"`
	Translation string `json:"translation" db:"translation"`
}