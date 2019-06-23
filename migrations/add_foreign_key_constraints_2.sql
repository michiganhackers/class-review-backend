/* Add foreign key constraints to */

ALTER TABLE IF EXISTS reviews (
    ADD CONSTRAINT fk_courseId FOREIGN KEY (courseId) REFERENCES courses(courseId)
)