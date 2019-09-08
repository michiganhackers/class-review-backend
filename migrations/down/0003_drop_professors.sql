/* Drop professors table and professor_uniqname foreign key contstraint */

ALTER TABLE reviews DROP FOREIGN KEY fk_professor_uniqname;
DROP TABLE professors;

