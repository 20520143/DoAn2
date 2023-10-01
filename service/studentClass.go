package service

import (
	"fmt"
	"projectDemo/model"
	"strings"
	"time"

	"gorm.io/gorm"
)

type studentClass struct{}

func NewStudentClass() IStudentClass {
	return &studentClass{}
}
func (s *studentClass) RegisterClass(req model.StudentClassReq) error {
	var parameter model.Parameter
	var items []string
	var credits int
	var studentClasses []model.StudentClass
	//check if there are available slots in the class
	err := DB.Raw(`Select id from classes where id in ? and student_count = 0`, req.ListClassID).Scan(&items).Error
	if err != nil {
		return err
	}
	if len(items) != 0 {
		return fmt.Errorf("There are no slots in %s", strings.Join(items, " "))
	}
	// check classes have been registered
	err = DB.Raw(`Select class_id from student_classes where student_id = ? and class_id in ?`, req.StudentID, req.ListClassID).Scan(&items).Error
	if err != nil {
		return err
	}
	if len(items) != 0 {
		return fmt.Errorf("Classes have been registered: %s", strings.Join(items, " "))
	}
	// check that the class does not have duplicate course
	err = DB.Raw(`SELECT c.course_id
				FROM classes c
				WHERE c.id IN (SELECT class_id FROM student_classes as s_c WHERE s_c.student_id= ? ) OR c.id IN ?
				GROUP BY c.course_id
				HAVING COUNT(*) > 1;`, req.StudentID, req.ListClassID).Scan(&items).Error
	if err != nil {
		return err
	}
	if len(items) != 0 {
		return fmt.Errorf("The courses are duplicated: %s", strings.Join(items, " "))
	}
	//check if the student's credit limit allows registration
	err = DB.Where("id = ?", "STCTD").First(&parameter).Error
	if err != nil {
		return err
	}
	err = DB.Raw(` SELECT SUM(total_credits) 
	FROM (
		SELECT SUM(co.credits) AS total_credits
		FROM courses AS co
		INNER JOIN classes AS c ON c.course_id = co.id
		WHERE c.id IN (?)
		UNION ALL
		SELECT credits_registered
		FROM students
		WHERE id = ?
	) AS subquery`, req.ListClassID, req.StudentID).Scan(&credits).Error
	if err != nil {
		return err
	}
	if credits > parameter.Value {
		return fmt.Errorf("The maximum number of credits registered is %d", parameter.Value)
	}
	// transaction register classes
	return DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Exec(`update classes set student_count = student_count - 1 where id in ?`, req.ListClassID).Error
		if err != nil {
			return err
		}
		err = tx.Exec(`update students set credits_registered = ? where id=?`, credits, req.StudentID).Error
		if err != nil {
			return err
		}
		for _, classID := range req.ListClassID {
			studentClasses = append(studentClasses, model.StudentClass{
				StudentID:      req.StudentID,
				ClassID:        classID,
				EnrollmentDate: time.Now()})
		}
		err = tx.Save(&studentClasses).Error
		if err != nil {
			return err
		}
		return nil
	})
}
func (c *studentClass) DeleteClass(req model.StudentClassReq) error {
	var items []string
	err := DB.Raw(`SELECT id FROM classes 
					WHERE id IN ? AND id 
					NOT IN (SELECT class_id FROM student_classes WHERE student_id = ?) `, req.ListClassID, req.StudentID).Scan(&items).Error
	if err != nil {
		return err
	}
	if len(items) != 0 {
		return fmt.Errorf("These classes have not yet been registered: %s", strings.Join(items, " "))
	}

	return DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Exec(`UPDATE classes AS c
						SET c.student_count = c.student_count + 1
						WHERE c.id IN (?)`, req.ListClassID).Error
		if err != nil {
			return err
		}
		err = tx.Exec(`UPDATE students AS s
						SET s.credits_registered = s.credits_registered - (
							SELECT SUM(co.credits)
							FROM classes AS c
							INNER JOIN courses AS co ON c.course_id = co.id
							WHERE c.id IN ?
						)
						WHERE s.id = ?;`, req.ListClassID, req.StudentID).Error
		if err != nil {
			return err
		}
		err = tx.Exec("delete from student_classes where class_id in ? and student_id=?", req.ListClassID, req.StudentID).Error
		if err != nil {
			return err
		}
		return nil
	})
}
