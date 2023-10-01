package model

import "time"

type Student struct {
	Id                string    `gorm:"primaryKey;column:id"`
	FullName          string    `gorm:"not null;column:full_name"`
	BirthDate         time.Time `gorm:"not null;column:birth_date"`
	CreditsRegistered int       `gorm:"column:credits_registered;default:0"`
	Classes           []Class   `gorm:"many2many:student_classes"`
}
