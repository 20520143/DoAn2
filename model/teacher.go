package model

import "time"

type Teacher struct {
	Id        string    `gorm:"primaryKey;column:id"`
	FullName  string    `gorm:"not null;column:full_name"`
	BirthDate time.Time `gorm:"not null;column:birth_date"`
}
