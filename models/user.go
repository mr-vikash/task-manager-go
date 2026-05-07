package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"-"`
	Age      uint   `json:"age" binding:"required,gte=18"`

	Tasks []Task `json:"tasks" gorm:"foreignKey:UserID"`
}
