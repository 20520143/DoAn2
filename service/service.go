package service

import (
	"projectDemo/model"
)

//	type IStudent interface {
//		PostStudents(students []model.Students) error
//		GetStudents(student model.StudentReq) ([]model.Students, error)
//	}
type IAuthen interface {
	UserAuthentication(req model.Users) (string, error)
	CreateUser(req model.Users) (string, error)
}
type IStudentClass interface {
	RegisterClass(req model.StudentClassReq) error
	DeleteClass(req model.StudentClassReq) error
}
type IClass interface {
	GetTeacherStudent(classID string) (model.Class, error)
	GetTeacherStudentAll() ([]model.Class, error)
	Create(class model.Class) error
}
type IStudent interface {
	GetClass(studentID string) ([]ClassOfStudent, error)
}
