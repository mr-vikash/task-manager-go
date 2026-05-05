package models

type User struct {
	ID    uint
	Name  string
	Email string
	Tasks []Task
}
