package model

import "time"

type Class struct {
	Id           string    `gorm:"primaryKey;column:id"`
	CourseID     string    `gorm:"not null;column:course_id"`
	TeacherID    *string   `gorm:"column:teacher_id;default:null"`
	StartDate    time.Time `gorm:"not null;column:start_date"`
	EndDate      time.Time `gorm:"not null;column:end_date"`
	StudentCount int       `gorm:";column:student_count;default:0"`
	Teacher      Teacher
	Students     []Student `gorm:"many2many:student_classes"`
}
