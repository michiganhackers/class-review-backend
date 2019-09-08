/* Add course ID foreign key constraint to reviews table*/

ALTER TABLE reviews ADD CONSTRAINT fk_courseId FOREIGN KEY (courseId) REFERENCES courses(id);