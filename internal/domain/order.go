package domain

type Order struct {
	Id      string `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
}
