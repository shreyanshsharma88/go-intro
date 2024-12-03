package models

type Todo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
}
