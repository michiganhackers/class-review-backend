/* Add professors table and professor_uniqname foreign key contstraint */

CREATE TABLE IF NOT EXISTS professors(
    professor_uniqname VARCHAR(255) PRIMARY KEY,
    professor_name VARCHAR(255)
);


ALTER TABLE reviews ADD CONSTRAINT fk_professor_uniqname FOREIGN KEY (professor_uniqname) REFERENCES professors(professor_uniqname);