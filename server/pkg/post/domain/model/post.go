package model

type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}
