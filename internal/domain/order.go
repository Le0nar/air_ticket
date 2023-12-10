package domain

type Order struct {
	Id      string `db:"id"`
	Content string `db:"content"`
}
