package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Tasks    []Task `json:"tasks"`
	Password string `json:"password"`
}
