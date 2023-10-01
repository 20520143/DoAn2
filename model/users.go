package model

// Model
type Users struct {
	Id       string `gorm:"primarykey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
