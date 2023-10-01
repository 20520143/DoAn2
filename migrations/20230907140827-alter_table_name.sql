-- +migrate Up
ALTER TABLE students
CHANGE COLUMN student_id id VARCHAR(255) NOT NULL;
ALTER TABLE classes
CHANGE COLUMN class_id id VARCHAR(255) NOT NULL;
ALTER TABLE courses
CHANGE COLUMN course_id id VARCHAR(255) NOT NULL;
CHANGE COLUMN course_id id VARCHAR(255) NOT NULL;
CHANGE COLUMN course_id id VARCHAR(255) NOT NULL;
ALTER TABLE parameters
CHANGE COLUMN parameter_id id VARCHAR(255) NOT NULL;
ALTER TABLE teachers
CHANGE COLUMN teacher_id id VARCHAR(255) NOT NULL;

-- +migrate Down
ALTER TABLE students
CHANGE COLUMN id student_id VARCHAR(255) NOT NULL;
ALTER TABLE classes
CHANGE COLUMN id class_id VARCHAR(255) NOT NULL;
ALTER TABLE courses
CHANGE COLUMN id course_id VARCHAR(255) NOT NULL;
ALTER TABLE parameters
CHANGE COLUMN id parameter_id VARCHAR(255) NOT NULL;
ALTER TABLE teachers
CHANGE COLUMN id teacher_id VARCHAR(255) NOT NULL;
