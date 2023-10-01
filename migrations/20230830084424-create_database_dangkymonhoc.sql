-- +migrate Up
CREATE TABLE students (
    student_id VARCHAR(255) PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    credits_registered INT DEFAULT 0
);
CREATE TABLE courses (
    course_id VARCHAR(255) PRIMARY KEY,
    course_name VARCHAR(255) NOT NULL,
    credits INT NOT NULL
);
CREATE TABLE teachers (
    teacher_id VARCHAR(255) PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL
);
CREATE TABLE classes (
    class_id VARCHAR(255) PRIMARY KEY,
    course_id VARCHAR(255),
    teacher_id VARCHAR(255),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    student_count INT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(course_id),
    FOREIGN KEY (teacher_id) REFERENCES teachers(teacher_id)
);
CREATE TABLE student_classes (
    student_id VARCHAR(255),
    class_id VARCHAR(255),
    enrollment_date DATE NOT NULL,
    PRIMARY KEY (student_id, class_id),
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (class_id) REFERENCES classes(class_id)
);
CREATE TABLE parameters (
    parameter_id VARCHAR(255) PRIMARY KEY,
    parameter_name VARCHAR(255) NOT NULL,
    value INT NOT NULL
);
-- +migrate Down
DROP TABLE student_classes, classes, students, courses, teachers, parameters;
