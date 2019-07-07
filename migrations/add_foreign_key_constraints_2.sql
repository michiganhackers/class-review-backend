/* Add course ID foreign key constraint to reviews table

/* Up migration*/

ALTER TABLE reviews ADD CONSTRAINT fk_courseId FOREIGN KEY (courseId) REFERENCES courses(id);


/* Down migration -- comment out the up migration and uncomment this to run */
/*
ALTER TABLE reviews DROP FOREIGN KEY fk_courseId;
*/