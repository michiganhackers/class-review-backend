CREATE TABLE IF NOT EXISTS professors(
    professor_name VARCHAR(255),
    professor_uniqname VARCHAR(255);
);

ALTER TABLE IF EXISTS reviews(
    CHANGE professor_name professor_uniqname VARCHAR(255)
    ADD CONSTRAINT fk_professor_uniqname FOREIGN KEY (professor_uniqname) REFERENCES professors(professor_uniqname)
);