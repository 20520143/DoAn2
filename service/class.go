package service

import (
	"projectDemo/model"
)

type class struct{}

func NewClass() IClass {
	return &class{}
}

func (*class) GetTeacherStudent(classID string) (model.Class, error) {
	var class model.Class
	err := DB.Preload("Teacher").Preload("Students").Where("id=?", classID).First(&class).Error
	if err != nil {
		return class, err
	}
	return class, nil
}

func (*class) GetTeacherStudentAll() ([]model.Class, error) {
	var classes []model.Class
	err := DB.Preload("Teacher").Preload("Students").Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}
func (*class) Create(class model.Class) error {
	err := DB.Save(&class).Error
	if err != nil {
		return err
	}
	return nil
}
