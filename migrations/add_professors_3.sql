/* Add professors table and professor_uniqname foreign key contstraint */
/* Up migration*/


CREATE TABLE IF NOT EXISTS professors(
    professor_uniqname VARCHAR(255) PRIMARY KEY,
    professor_name VARCHAR(255)
);


ALTER TABLE reviews ADD CONSTRAINT fk_professor_uniqname FOREIGN KEY (professor_uniqname) REFERENCES professors(professor_uniqname);



/* Down migration -- comment out the up migration and uncomment this to run */

/*
ALTER TABLE reviews DROP FOREIGN KEY fk_professor_uniqname;
DROP TABLE professors
*/

