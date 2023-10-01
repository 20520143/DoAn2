package model

type Cours struct {
	Id         string `gorm:"primaryKey;column:id"`
	CourseName string `gorm:"not null;column:course_name"`
	Credits    int    `gorm:"not null;column:credits"`
	Classes    []Class
}
