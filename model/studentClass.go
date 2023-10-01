package model

import (
	"time"
)

type StudentClass struct {
	StudentID      string    `gorm:"primaryKey"`
	ClassID        string    `gorm:"primaryKey"`
	EnrollmentDate time.Time `gorm:"not null;column:enrollment_date"`
}
type StudentClassReq struct {
	StudentID   string
	ListClassID []string
}
type StudentClassRes struct {
	Success []string
	Fail    map[string]string
}
