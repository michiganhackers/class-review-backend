ALTER TABLE IF EXISTS courses (
    ADD CONSTRAINT fk_courseId FOREIGN KEY (courseId) REFERENCES reviews(courseId)
)