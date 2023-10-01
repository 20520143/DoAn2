package service

import (
	"time"
)

type student struct{}

type ClassOfStudent struct {
	ID          string
	StartDate   time.Time
	EndDate     time.Time
	TeacherName string
	CourseName  string
	Credits     int
}

func NewStudent() IStudent {
	return &student{}
}

func (s *student) GetClass(studentID string) ([]ClassOfStudent, error) {
	var inforClasses []ClassOfStudent
	result := DB.Raw(`Select c.id,c.start_date,c.end_date,t.full_name as teacher_name,co.course_name,co.credits 
					from students as s 
					inner join student_classes as s_c on s.id = s_c.student_id
					inner join classes as c on c.id = s_c.class_id
					inner join courses as co on co.id = c.course_id
					left join teachers as t on t.id = c.teacher_id where s.id=?`, studentID).Scan(&inforClasses)
	if result.Error != nil {
		return inforClasses, result.Error
	}
	return inforClasses, nil
}
