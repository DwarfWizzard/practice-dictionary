package domain

type Word struct {
	Id          int    `json:"id" db:"id"`
	Original    string `json:"original" db:"original"`
	Translation string `json:"translation" db:"translation"`
}
