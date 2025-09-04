-- Sample students data
INSERT INTO students (student_id, first_name, last_name, email, phone, date_of_birth, gender, address) VALUES
('STU001', 'John', 'Doe', 'john.doe@student.swadiq.com', '+256700123456', '2010-05-15', 'male', 'Kampala, Uganda'),
('STU002', 'Jane', 'Smith', 'jane.smith@student.swadiq.com', '+256700123457', '2009-08-22', 'female', 'Entebbe, Uganda'),
('STU003', 'Michael', 'Johnson', 'michael.johnson@student.swadiq.com', '+256700123458', '2010-12-03', 'male', 'Jinja, Uganda'),
('STU004', 'Sarah', 'Williams', 'sarah.williams@student.swadiq.com', '+256700123459', '2009-03-18', 'female', 'Mbarara, Uganda'),
('STU005', 'David', 'Brown', 'david.brown@student.swadiq.com', '+256700123460', '2010-07-09', 'male', 'Gulu, Uganda')
ON CONFLICT (student_id) DO NOTHING;
