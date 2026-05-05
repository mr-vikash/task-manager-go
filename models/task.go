package models

type Task struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user"`
}
