
-- +migrate Up
-- Insert data into students table
INSERT INTO students (student_id, full_name, birth_date, credits_registered)
VALUES ('S001', 'John Doe', '2000-01-15', 12),
       ('S002', 'Jane Smith', '2001-03-22', 15),
       ('S003', 'Michael Johnson', '1999-08-10', 9);

-- Insert data into courses table
INSERT INTO courses (course_id, course_name, credits)
VALUES ('C001', 'Mathematics', 4),
       ('C002', 'History', 3),
       ('C003', 'Physics', 5);

-- Insert data into teachers table
INSERT INTO teachers (teacher_id, full_name, birth_date)
VALUES ('T001', 'Professor Smith', '1975-05-20'),
       ('T002', 'Dr. Johnson', '1982-11-12'),
       ('T003', 'Ms. Anderson', '1990-03-28');

-- Insert data into classes table
INSERT INTO classes (class_id, course_id, teacher_id, start_date, end_date, student_count)
VALUES ('CL001', 'C001', 'T001', '2023-09-01', '2023-12-15', 25),
       ('CL002', 'C001', 'T002', '2023-09-10', '2023-12-20', 18),
       ('CL003', 'C002', 'T001', '2023-09-05', '2023-12-18', 20),
       ('CL004', 'C002', 'T003', '2023-09-15', '2023-12-25', 22),
       ('CL005', 'C003', 'T002', '2023-09-08', '2023-12-22', 23),
       ('CL006', 'C003', 'T003', '2023-09-18', '2023-12-28', 18);

-- Insert data into student_classes table
INSERT INTO student_classes (student_id, class_id, enrollment_date)
VALUES ('S001', 'CL001', '2023-09-02'),
       ('S002', 'CL001', '2023-09-02'),
       ('S002', 'CL002', '2023-09-12'),
       ('S003', 'CL002', '2023-09-12'), 
       ('S001', 'CL003', '2023-09-06'),
       ('S002', 'CL003', '2023-09-06'),
       ('S003', 'CL004', '2023-09-16'),
       ('S001', 'CL005', '2023-09-10'),
       ('S003', 'CL005', '2023-09-10'),
       ('S002', 'CL006', '2023-09-20'),
       ('S003', 'CL006', '2023-09-20');

-- Insert data into parameters table
INSERT INTO parameters (parameter_id, parameter_name, value)
VALUES ('STCTD', 'Số tín chỉ tối đa được đăng ký', 20);


-- +migrate Down
-- Delete data from STUDENT_CLASS table
DELETE FROM student_classes;

-- Delete data from CLASS table
DELETE FROM classes;

-- Delete data from STUDENT table
DELETE FROM students;

-- Delete data from COURSE table
DELETE FROM courses;

-- Delete data from TEACHER table
DELETE FROM teachers;

-- Delete data from PARAMETERS table
DELETE FROM parameters;
