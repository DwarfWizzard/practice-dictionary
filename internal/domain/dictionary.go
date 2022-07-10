package domain

type Dictionary struct {
	Id          int      `json:"id" db:"id"`
	Original    string   `json:"original" db:"original"`
	Translation []string `json:"translation" db:"translation"`
}
