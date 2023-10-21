package entity

type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

type PostRepository interface {
	FindAll() ([]*Post, error)
	FindById(id string) (*Post, error)
	Create(post *Post) error
}
